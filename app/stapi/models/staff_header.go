// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// StaffHeader Header staff, embedded in other objects
// swagger:model StaffHeader
type StaffHeader struct {

	// Staff name
	Name string `json:"name,omitempty"`

	// Staff unique ID
	UID string `json:"uid,omitempty"`
}

// Validate validates this staff header
func (m *StaffHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StaffHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StaffHeader) UnmarshalBinary(b []byte) error {
	var res StaffHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
