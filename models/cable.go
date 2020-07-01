// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Cable cable
//
// swagger:model Cable
type Cable struct {

	// cable id
	CableID interface{} `json:"cable_id,omitempty"`

	// cable length
	CableLength interface{} `json:"cable_length,omitempty"`

	// cable length units
	CableLengthUnits interface{} `json:"cable_length_units,omitempty"`

	// custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// end back patch panel
	EndBackPatchPanel interface{} `json:"end_back_patch_panel,omitempty"`

	// end cable color
	EndCableColor interface{} `json:"end_cable_color,omitempty"`

	// end cable type
	EndCableType interface{} `json:"end_cable_type,omitempty"`

	// end connector type
	EndConnectorType interface{} `json:"end_connector_type,omitempty"`

	// end content type
	EndContentType interface{} `json:"end_content_type,omitempty"`

	// end optic type
	EndOpticType interface{} `json:"end_optic_type,omitempty"`

	// end point multiple
	EndPointMultiple interface{} `json:"end_point_multiple,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`

	// notes
	Notes interface{} `json:"notes,omitempty"`

	// origin back patch panel
	OriginBackPatchPanel interface{} `json:"origin_back_patch_panel,omitempty"`

	// origin cable color
	OriginCableColor interface{} `json:"origin_cable_color,omitempty"`

	// origin cable type
	OriginCableType interface{} `json:"origin_cable_type,omitempty"`

	// origin connector type
	OriginConnectorType interface{} `json:"origin_connector_type,omitempty"`

	// origin content type
	OriginContentType interface{} `json:"origin_content_type,omitempty"`

	// origin netport id
	OriginNetportID interface{} `json:"origin_netport_id,omitempty"`

	// origin netport name
	OriginNetportName interface{} `json:"origin_netport_name,omitempty"`

	// origin optic type
	OriginOpticType interface{} `json:"origin_optic_type,omitempty"`

	// room
	Room interface{} `json:"room,omitempty"`

	// tags
	Tags interface{} `json:"tags,omitempty"`

	// vendor
	Vendor interface{} `json:"vendor,omitempty"`
}

// Validate validates this cable
func (m *Cable) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Cable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cable) UnmarshalBinary(b []byte) error {
	var res Cable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}