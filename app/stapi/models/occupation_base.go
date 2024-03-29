// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OccupationBase Base occupations, returned in search results
// swagger:model OccupationBase
type OccupationBase struct {

	// Whether it's a legal occupation
	LegalOccupation bool `json:"legalOccupation,omitempty"`

	// Whether it's a medical occupation
	MedicalOccupation bool `json:"medicalOccupation,omitempty"`

	// Occupation name
	Name string `json:"name,omitempty"`

	// Whether it's a scientific occupation
	ScientificOccupation bool `json:"scientificOccupation,omitempty"`

	// Occupation unique ID
	UID string `json:"uid,omitempty"`
}

// Validate validates this occupation base
func (m *OccupationBase) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OccupationBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OccupationBase) UnmarshalBinary(b []byte) error {
	var res OccupationBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
