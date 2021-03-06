// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Racks racks
//
// swagger:model racks
type Racks struct {

	// asset no
	AssetNo interface{} `json:"asset_no,omitempty"`

	// available u
	Availableu interface{} `json:"available_u,omitempty"`

	// building
	Building interface{} `json:"building,omitempty"`

	// col size
	ColSize interface{} `json:"col_size,omitempty"`

	// custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// first number
	FirstNumber interface{} `json:"first_number,omitempty"`

	// groups
	Groups interface{} `json:"groups,omitempty"`

	// manufacturer
	Manufacturer interface{} `json:"manufacturer,omitempty"`

	// name
	Name interface{} `json:"name,omitempty"`

	// notes
	Notes interface{} `json:"notes,omitempty"`

	// number between us
	NumberBetweenUs interface{} `json:"number_between_us,omitempty"`

	// numbering start from bottom
	NumberingStartFromBottom interface{} `json:"numbering_start_from_bottom,omitempty"`

	// rack id
	RackID interface{} `json:"rack_id,omitempty"`

	// rack middle option
	RackMiddleOption interface{} `json:"rack_middle_option,omitempty"`

	// rack url
	RackURL interface{} `json:"rack_url,omitempty"`

	// room
	Room interface{} `json:"room,omitempty"`

	// row
	Row interface{} `json:"row,omitempty"`

	// row size
	RowSize interface{} `json:"row_size,omitempty"`

	// size
	Size interface{} `json:"size,omitempty"`

	// start col
	StartCol interface{} `json:"start_col,omitempty"`

	// start row
	StartRow interface{} `json:"start_row,omitempty"`

	// tags
	Tags interface{} `json:"tags,omitempty"`
}

// Validate validates this racks
func (m *Racks) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Racks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Racks) UnmarshalBinary(b []byte) error {
	var res Racks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
