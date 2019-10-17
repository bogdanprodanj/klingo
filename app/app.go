package app

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/klingo/app/stapi/client"
	"github.com/klingo/app/stapi/client/character"
	"github.com/klingo/app/stapi/models"
	"github.com/pkg/errors"
)

var (
	// UnknownSpecies is assigned to a name if no species provided
	UnknownSpecies  = "unknown"
	allowedInput    = regexp.MustCompile(`^[a-zA-z0-9,.'\s]+$`)
	errInvalidInput = errors.New("input contains characters that have corresponding value in pIqaD")
)

// StartTrack responsible for translation and finding species
type StartTrack struct {
	client *client.ST
}

// New creates new StartTrack with stapi http client
func New() *StartTrack {
	return &StartTrack{client: client.NewHTTPClient(nil)}
}

// Validate checks if provided input is valid
func (st *StartTrack) Validate(englishName string) error {
	if !allowedInput.MatchString(englishName) {
		return errInvalidInput
	}
	return nil
}

// Translate translates given input into pIqad
func (st *StartTrack) Translate(englishName string) string {
	var piqadName []string
	for i := 0; i < len(englishName); i++ {
		piqadName = append(piqadName, strings.ReplaceAll(fmt.Sprintf("%U", englishToPiqadMap[strings.ToLower(string(englishName[i]))].unicode), "U+", "0x"))
	}
	return strings.Join(piqadName, " ")
}

// FindSpecies returns name(s) mapped to the corresponding species
func (st *StartTrack) FindSpecies(input string) (map[string]map[string]struct{}, error) {
	characterSearchParams := character.NewPostCharacterSearchParams()
	characterSearchParams.SetName(&input)
	resp, err := st.client.Character.PostCharacterSearch(characterSearchParams)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to find character with name containing %s", input)
	}
	uids := make([]string, len(resp.Payload.Characters))
	for i, v := range resp.Payload.Characters {
		uids[i] = v.UID
	}
	crs, err := st.findCharacters(uids)
	if err != nil {
		return nil, err
	}
	namesToSpecies := make(map[string]map[string]struct{})
	for _, cr := range crs {
		species, ok := namesToSpecies[cr.Name]
		if !ok {
			species = make(map[string]struct{})
		}
		for _, cs := range cr.CharacterSpecies {
			species[cs.Name] = struct{}{}
		}
		namesToSpecies[cr.Name] = species
	}
	for k, v := range namesToSpecies {
		if len(v) == 0 {
			namesToSpecies[k] = map[string]struct{}{UnknownSpecies: {}}
		}
	}
	return namesToSpecies, nil
}

type characterResult struct {
	resp *character.GetCharacterOK
	uID  string
	err  error
}

type characterError struct {
	uIDs map[string]error
}

func (ce *characterError) Error() string {
	if len(ce.uIDs) == 0 {
		return ""
	}
	return fmt.Sprintf("some of the characters' species weren't detected due to errors: %+v", ce.uIDs)
}

func (st *StartTrack) findCharacters(uIDs []string) ([]*models.CharacterFull, error) {
	characterResults := make(chan characterResult, len(uIDs))
	var wg sync.WaitGroup
	wg.Add(len(uIDs))
	for _, v := range uIDs {
		go func(uID string) {
			defer wg.Done()
			cp := character.NewGetCharacterParams()
			cp.SetUID(uID)
			resp, err := st.client.Character.GetCharacter(cp)
			characterResults <- characterResult{
				uID:  uID,
				resp: resp,
				err:  err,
			}
		}(v)
	}
	wg.Wait()
	close(characterResults)
	chErrs := &characterError{uIDs: make(map[string]error)}
	var characters []*models.CharacterFull
	for cr := range characterResults {
		if cr.err != nil {
			chErrs.uIDs[cr.uID] = cr.err
			continue
		}
		characters = append(characters, cr.resp.Payload.Character)
	}
	if len(chErrs.uIDs) != 0 {
		return characters, chErrs
	}
	return characters, nil
}
