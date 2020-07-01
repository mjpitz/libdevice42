// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetAllCircuits get all circuits
//
// swagger:model Get_all_Circuits
type GetAllCircuits struct {

	// ID
	ID interface{} `json:"ID,omitempty"`

	// bandwidth
	Bandwidth interface{} `json:"bandwidth,omitempty"`

	// circuit id
	CircuitID interface{} `json:"circuit_id,omitempty"`

	// custom fields
	CustomFields []*GetAllCircuitsCustomFieldsItems0 `json:"custom_fields"`

	// customer
	Customer interface{} `json:"customer,omitempty"`

	// dmarc
	Dmarc interface{} `json:"dmarc,omitempty"`

	// dmarc id
	DmarcID interface{} `json:"dmarc_id,omitempty"`

	// end point id
	EndPointID interface{} `json:"end_point_id,omitempty"`

	// end point type
	EndPointType interface{} `json:"end_point_type,omitempty"`

	// notes
	Notes interface{} `json:"notes,omitempty"`

	// order date
	OrderDate interface{} `json:"order_date,omitempty"`

	// origin id
	OriginID interface{} `json:"origin_id,omitempty"`

	// origin type
	OriginType interface{} `json:"origin_type,omitempty"`

	// provision date
	ProvisionDate interface{} `json:"provision_date,omitempty"`

	// turn on date
	TurnOnDate interface{} `json:"turn_on_date,omitempty"`

	// type
	Type interface{} `json:"type,omitempty"`

	// vendor
	Vendor interface{} `json:"vendor,omitempty"`
}

// Validate validates this get all circuits
func (m *GetAllCircuits) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetAllCircuits) validateCustomFields(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomFields) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomFields); i++ {
		if swag.IsZero(m.CustomFields[i]) { // not required
			continue
		}

		if m.CustomFields[i] != nil {
			if err := m.CustomFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetAllCircuits) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAllCircuits) UnmarshalBinary(b []byte) error {
	var res GetAllCircuits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetAllCircuitsCustomFieldsItems0 get all circuits custom fields items0
//
// swagger:model GetAllCircuitsCustomFieldsItems0
type GetAllCircuitsCustomFieldsItems0 struct {

	// key
	Key interface{} `json:"key,omitempty"`

	// notes
	Notes interface{} `json:"notes,omitempty"`

	// value
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this get all circuits custom fields items0
func (m *GetAllCircuitsCustomFieldsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetAllCircuitsCustomFieldsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAllCircuitsCustomFieldsItems0) UnmarshalBinary(b []byte) error {
	var res GetAllCircuitsCustomFieldsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
