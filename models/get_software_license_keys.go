// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetSoftwareLicenseKeys get software license keys
//
// swagger:model Get_Software_License_Keys
type GetSoftwareLicenseKeys struct {

	// count
	Count interface{} `json:"count,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`

	// key
	Key interface{} `json:"key,omitempty"`

	// notes
	Notes interface{} `json:"notes,omitempty"`

	// software id
	SoftwareID interface{} `json:"software_id,omitempty"`

	// software name
	SoftwareName interface{} `json:"software_name,omitempty"`
}

// Validate validates this get software license keys
func (m *GetSoftwareLicenseKeys) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetSoftwareLicenseKeys) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSoftwareLicenseKeys) UnmarshalBinary(b []byte) error {
	var res GetSoftwareLicenseKeys
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
