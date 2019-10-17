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

// MovieBase Base movie, returned in search results
// swagger:model MovieBase
type MovieBase struct {

	// Director of the movie
	MainDirector *StaffHeader `json:"mainDirector,omitempty"`

	// Starting stardate of movie story
	StardateFrom float32 `json:"stardateFrom,omitempty"`

	// Ending stardate of movie story
	StardateTo float32 `json:"stardateTo,omitempty"`

	// Movie title
	Title string `json:"title,omitempty"`

	// Movie title in Bulgarian
	TitleBulgarian string `json:"titleBulgarian,omitempty"`

	// Movie title in Catalan
	TitleCatalan string `json:"titleCatalan,omitempty"`

	// Movie title in Chinese traditional
	TitleChineseTraditional string `json:"titleChineseTraditional,omitempty"`

	// Movie title in German
	TitleGerman string `json:"titleGerman,omitempty"`

	// Movie title in Italian
	TitleItalian string `json:"titleItalian,omitempty"`

	// Movie title in Japanese
	TitleJapanese string `json:"titleJapanese,omitempty"`

	// Movie title in Polish
	TitlePolish string `json:"titlePolish,omitempty"`

	// Movie title in Russian
	TitleRussian string `json:"titleRussian,omitempty"`

	// Movie title in Serbian
	TitleSerbian string `json:"titleSerbian,omitempty"`

	// Movie title in Spanish
	TitleSpanish string `json:"titleSpanish,omitempty"`

	// Movie unique ID
	UID string `json:"uid,omitempty"`

	// Date the movie was first released in the United States
	// Format: date
	UsReleaseDate strfmt.Date `json:"usReleaseDate,omitempty"`

	// Starting year of movie story
	YearFrom int64 `json:"yearFrom,omitempty"`

	// Ending year of movie story
	YearTo int64 `json:"yearTo,omitempty"`
}

// Validate validates this movie base
func (m *MovieBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMainDirector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsReleaseDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MovieBase) validateMainDirector(formats strfmt.Registry) error {

	if swag.IsZero(m.MainDirector) { // not required
		return nil
	}

	if m.MainDirector != nil {
		if err := m.MainDirector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mainDirector")
			}
			return err
		}
	}

	return nil
}

func (m *MovieBase) validateUsReleaseDate(formats strfmt.Registry) error {

	if swag.IsZero(m.UsReleaseDate) { // not required
		return nil
	}

	if err := validate.FormatOf("usReleaseDate", "body", "date", m.UsReleaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MovieBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MovieBase) UnmarshalBinary(b []byte) error {
	var res MovieBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}