// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkshareDetails networkshare details
//
// swagger:model Networkshare_details
type NetworkshareDetails struct {

	// caption
	Caption interface{} `json:"caption,omitempty"`

	// description
	Description interface{} `json:"description,omitempty"`

	// device name
	DeviceName interface{} `json:"device_name,omitempty"`

	// first detected
	FirstDetected interface{} `json:"first_detected,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`

	// install date
	InstallDate interface{} `json:"install_date,omitempty"`

	// last updated
	LastUpdated interface{} `json:"last_updated,omitempty"`

	// name
	Name interface{} `json:"name,omitempty"`

	// path
	Path interface{} `json:"path,omitempty"`

	// status
	Status interface{} `json:"status,omitempty"`

	// type
	Type interface{} `json:"type,omitempty"`
}

// Validate validates this networkshare details
func (m *NetworkshareDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkshareDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkshareDetails) UnmarshalBinary(b []byte) error {
	var res NetworkshareDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
