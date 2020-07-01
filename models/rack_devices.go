// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RackDevices rack devices
//
// swagger:model RackDevices
type RackDevices struct {

	// depth
	Depth interface{} `json:"depth,omitempty"`

	// device
	Device *RackDevicesDevice `json:"device,omitempty"`

	// file names
	FileNames interface{} `json:"file_names,omitempty"`

	// orientation
	Orientation interface{} `json:"orientation,omitempty"`

	// reversed
	Reversed interface{} `json:"reversed,omitempty"`

	// size
	Size interface{} `json:"size,omitempty"`

	// start at
	StartAt interface{} `json:"start_at,omitempty"`

	// where
	Where interface{} `json:"where,omitempty"`

	// width
	Width interface{} `json:"width,omitempty"`

	// xpos
	Xpos interface{} `json:"xpos,omitempty"`
}

// Validate validates this rack devices
func (m *RackDevices) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackDevices) validateDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.Device) { // not required
		return nil
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RackDevices) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackDevices) UnmarshalBinary(b []byte) error {
	var res RackDevices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackDevicesDevice rack devices device
//
// swagger:model RackDevicesDevice
type RackDevicesDevice struct {

	// asset no
	AssetNo interface{} `json:"asset_no,omitempty"`

	// device id
	DeviceID interface{} `json:"device_id,omitempty"`

	// device url
	DeviceURL interface{} `json:"device_url,omitempty"`

	// is it blade host
	IsItBladeHost interface{} `json:"is_it_blade_host,omitempty"`

	// is it switch
	IsItSwitch interface{} `json:"is_it_switch,omitempty"`

	// is it virtual host
	IsItVirtualHost interface{} `json:"is_it_virtual_host,omitempty"`

	// name
	Name interface{} `json:"name,omitempty"`

	// serial no
	SerialNo interface{} `json:"serial_no,omitempty"`

	// type
	Type interface{} `json:"type,omitempty"`

	// uuid
	UUID interface{} `json:"uuid,omitempty"`
}

// Validate validates this rack devices device
func (m *RackDevicesDevice) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RackDevicesDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackDevicesDevice) UnmarshalBinary(b []byte) error {
	var res RackDevicesDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
