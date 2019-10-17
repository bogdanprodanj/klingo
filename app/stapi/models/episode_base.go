// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EpisodeBase Base episode, returned in search results
// swagger:model EpisodeBase
type EpisodeBase struct {

	// Episode number in season
	EpisodeNumber int64 `json:"episodeNumber,omitempty"`

	// Whether it's a feature length episode
	FeatureLength bool `json:"featureLength,omitempty"`

	// Date the episode script was completed
	// Format: date
	FinalScriptDate strfmt.Date `json:"finalScriptDate,omitempty"`

	// Production serial number
	ProductionSerialNumber string `json:"productionSerialNumber,omitempty"`

	// Season this episode belongs to
	Season *SeasonHeader `json:"season,omitempty"`

	// Season number
	SeasonNumber int64 `json:"seasonNumber,omitempty"`

	// Series this episode belongs to
	Series *SeriesHeader `json:"series,omitempty"`

	// Starting stardate of episode story
	StardateFrom float32 `json:"stardateFrom,omitempty"`

	// Ending stardate of episode story
	StardateTo float32 `json:"stardateTo,omitempty"`

	// Episode title
	Title string `json:"title,omitempty"`

	// Episode title in German
	TitleGerman string `json:"titleGerman,omitempty"`

	// Episode title in Italian
	TitleItalian string `json:"titleItalian,omitempty"`

	// Episode title in Japanese
	TitleJapanese string `json:"titleJapanese,omitempty"`

	// Episode unique ID
	UID string `json:"uid,omitempty"`

	// Date the episode was first aired in the United States
	// Format: date
	UsAirDate strfmt.Date `json:"usAirDate,omitempty"`

	// Starting year of episode story
	YearFrom int64 `json:"yearFrom,omitempty"`

	// Ending year of episode story
	YearTo int64 `json:"yearTo,omitempty"`
}

// Validate validates this episode base
func (m *EpisodeBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinalScriptDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsAirDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EpisodeBase) validateFinalScriptDate(formats strfmt.Registry) error {

	if swag.IsZero(m.FinalScriptDate) { // not required
		return nil
	}

	if err := validate.FormatOf("finalScriptDate", "body", "date", m.FinalScriptDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EpisodeBase) validateSeason(formats strfmt.Registry) error {

	if swag.IsZero(m.Season) { // not required
		return nil
	}

	if m.Season != nil {
		if err := m.Season.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("season")
			}
			return err
		}
	}

	return nil
}

func (m *EpisodeBase) validateSeries(formats strfmt.Registry) error {

	if swag.IsZero(m.Series) { // not required
		return nil
	}

	if m.Series != nil {
		if err := m.Series.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("series")
			}
			return err
		}
	}

	return nil
}

func (m *EpisodeBase) validateUsAirDate(formats strfmt.Registry) error {

	if swag.IsZero(m.UsAirDate) { // not required
		return nil
	}

	if err := validate.FormatOf("usAirDate", "body", "date", m.UsAirDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EpisodeBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EpisodeBase) UnmarshalBinary(b []byte) error {
	var res EpisodeBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
