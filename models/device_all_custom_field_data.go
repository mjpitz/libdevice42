// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceAllCustomFieldData device all custom field data
//
// swagger:model deviceAllCustomFieldData
type DeviceAllCustomFieldData struct {

	// key
	Key interface{} `json:"key,omitempty"`

	// notes
	Notes interface{} `json:"notes,omitempty"`

	// value
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this device all custom field data
func (m *DeviceAllCustomFieldData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceAllCustomFieldData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceAllCustomFieldData) UnmarshalBinary(b []byte) error {
	var res DeviceAllCustomFieldData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
