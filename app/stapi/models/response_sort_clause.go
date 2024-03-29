// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ResponseSortClause Single response sort clause
// swagger:model ResponseSortClause
type ResponseSortClause struct {

	// Order in which this clause was applied
	ClauseOrder int64 `json:"clauseOrder,omitempty"`

	// Sort direction
	Direction ResponseSortDirection `json:"direction,omitempty"`

	// Field name results are sorted by
	Name string `json:"name,omitempty"`
}

// Validate validates this response sort clause
func (m *ResponseSortClause) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseSortClause) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	if err := m.Direction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("direction")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponseSortClause) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseSortClause) UnmarshalBinary(b []byte) error {
	var res ResponseSortClause
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
