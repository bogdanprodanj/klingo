// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ResponsePage Object describing response page
// swagger:model ResponsePage
type ResponsePage struct {

	// Whether it is the first page
	FirstPage bool `json:"firstPage,omitempty"`

	// Whether it is the last page
	LastPage bool `json:"lastPage,omitempty"`

	// Number of elements in page
	NumberOfElements int32 `json:"numberOfElements,omitempty"`

	// Zero-based page number
	PageNumber int32 `json:"pageNumber,omitempty"`

	// Page size
	PageSize int32 `json:"pageSize,omitempty"`

	// Total elements found
	TotalElements int32 `json:"totalElements,omitempty"`

	// Total pages found
	TotalPages int32 `json:"totalPages,omitempty"`
}

// Validate validates this response page
func (m *ResponsePage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponsePage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponsePage) UnmarshalBinary(b []byte) error {
	var res ResponsePage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
