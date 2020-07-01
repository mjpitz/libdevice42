// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetHistory get history
//
// swagger:model getHistory
type GetHistory struct {

	// action
	Action interface{} `json:"action,omitempty"`

	// action time
	ActionTime interface{} `json:"action_time,omitempty"`

	// change message
	ChangeMessage interface{} `json:"change_message,omitempty"`

	// content type
	ContentType interface{} `json:"content_type,omitempty"`

	// obj repr
	ObjRepr interface{} `json:"obj_repr,omitempty"`

	// user
	User interface{} `json:"user,omitempty"`
}

// Validate validates this get history
func (m *GetHistory) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetHistory) UnmarshalBinary(b []byte) error {
	var res GetHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
