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

// Partmodels partmodels
//
// swagger:model Partmodels
type Partmodels struct {

	// limit
	Limit interface{} `json:"limit,omitempty"`

	// offset
	Offset interface{} `json:"offset,omitempty"`

	// partmodels
	Partmodels []*PartmodelsPartmodelsItems0 `json:"partmodels"`

	// total count
	TotalCount interface{} `json:"total_count,omitempty"`
}

// Validate validates this partmodels
func (m *Partmodels) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePartmodels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Partmodels) validatePartmodels(formats strfmt.Registry) error {

	if swag.IsZero(m.Partmodels) { // not required
		return nil
	}

	for i := 0; i < len(m.Partmodels); i++ {
		if swag.IsZero(m.Partmodels[i]) { // not required
			continue
		}

		if m.Partmodels[i] != nil {
			if err := m.Partmodels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("partmodels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Partmodels) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Partmodels) UnmarshalBinary(b []byte) error {
	var res Partmodels
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PartmodelsPartmodelsItems0 partmodels partmodels items0
//
// swagger:model PartmodelsPartmodelsItems0
type PartmodelsPartmodelsItems0 struct {

	// available
	Available interface{} `json:"available,omitempty"`

	// cores
	Cores interface{} `json:"cores,omitempty"`

	// cpuspeed
	Cpuspeed interface{} `json:"cpuspeed,omitempty"`

	// description
	Description interface{} `json:"description,omitempty"`

	// in devices
	InDevices interface{} `json:"in_devices,omitempty"`

	// in rma
	InRma interface{} `json:"in_rma,omitempty"`

	// in storage racks
	InStorageRacks interface{} `json:"in_storage_racks,omitempty"`

	// in storage rooms
	InStorageRooms interface{} `json:"in_storage_rooms,omitempty"`

	// in transit
	InTransit interface{} `json:"in_transit,omitempty"`

	// length
	Length interface{} `json:"length,omitempty"`

	// location
	Location interface{} `json:"location,omitempty"`

	// manufacturer
	Manufacturer interface{} `json:"manufacturer,omitempty"`

	// modelno
	Modelno interface{} `json:"modelno,omitempty"`

	// name
	Name interface{} `json:"name,omitempty"`

	// notes
	Notes interface{} `json:"notes,omitempty"`

	// partmodel id
	PartmodelID interface{} `json:"partmodel_id,omitempty"`

	// partno
	Partno interface{} `json:"partno,omitempty"`

	// threads
	Threads interface{} `json:"threads,omitempty"`

	// total count
	TotalCount interface{} `json:"total_count,omitempty"`

	// type
	Type interface{} `json:"type,omitempty"`

	// used
	Used interface{} `json:"used,omitempty"`
}

// Validate validates this partmodels partmodels items0
func (m *PartmodelsPartmodelsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PartmodelsPartmodelsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartmodelsPartmodelsItems0) UnmarshalBinary(b []byte) error {
	var res PartmodelsPartmodelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
