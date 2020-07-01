// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppcompsCustomFields appcomps custom fields
//
// swagger:model Appcomps_Custom_fields
type AppcompsCustomFields struct {

	// key
	Key interface{} `json:"key,omitempty"`

	// notes
	Notes interface{} `json:"notes,omitempty"`

	// value
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this appcomps custom fields
func (m *AppcompsCustomFields) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppcompsCustomFields) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppcompsCustomFields) UnmarshalBinary(b []byte) error {
	var res AppcompsCustomFields
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
