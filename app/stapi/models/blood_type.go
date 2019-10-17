// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// BloodType Blood type
// swagger:model BloodType
type BloodType string

const (

	// BloodTypeBNEGATIVE captures enum value "B_NEGATIVE"
	BloodTypeBNEGATIVE BloodType = "B_NEGATIVE"

	// BloodTypeONEGATIVE captures enum value "O_NEGATIVE"
	BloodTypeONEGATIVE BloodType = "O_NEGATIVE"

	// BloodTypeTNEGATIVE captures enum value "T_NEGATIVE"
	BloodTypeTNEGATIVE BloodType = "T_NEGATIVE"
)

// for schema
var bloodTypeEnum []interface{}

func init() {
	var res []BloodType
	if err := json.Unmarshal([]byte(`["B_NEGATIVE","O_NEGATIVE","T_NEGATIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bloodTypeEnum = append(bloodTypeEnum, v)
	}
}

func (m BloodType) validateBloodTypeEnum(path, location string, value BloodType) error {
	if err := validate.Enum(path, location, value, bloodTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this blood type
func (m BloodType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBloodTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
