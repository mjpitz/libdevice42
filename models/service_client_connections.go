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

// ServiceClientConnections service client connections
//
// swagger:model Service_client_connections
type ServiceClientConnections struct {

	// client connections
	ClientConnections []*ServiceClientConnectionsClientConnectionsItems0 `json:"client_connections"`

	// limit
	Limit interface{} `json:"limit,omitempty"`

	// offset
	Offset interface{} `json:"offset,omitempty"`

	// total count
	TotalCount interface{} `json:"total_count,omitempty"`
}

// Validate validates this service client connections
func (m *ServiceClientConnections) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientConnections(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceClientConnections) validateClientConnections(formats strfmt.Registry) error {

	if swag.IsZero(m.ClientConnections) { // not required
		return nil
	}

	for i := 0; i < len(m.ClientConnections); i++ {
		if swag.IsZero(m.ClientConnections[i]) { // not required
			continue
		}

		if m.ClientConnections[i] != nil {
			if err := m.ClientConnections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("client_connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceClientConnections) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceClientConnections) UnmarshalBinary(b []byte) error {
	var res ServiceClientConnections
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceClientConnectionsClientConnectionsItems0 service client connections client connections items0
//
// swagger:model ServiceClientConnectionsClientConnectionsItems0
type ServiceClientConnectionsClientConnectionsItems0 struct {

	// client device
	ClientDevice string `json:"client_device,omitempty"`

	// client device id
	ClientDeviceID int64 `json:"client_device_id,omitempty"`

	// client ip
	ClientIP string `json:"client_ip,omitempty"`

	// client process display name
	ClientProcessDisplayName string `json:"client_process_display_name,omitempty"`

	// client process name
	ClientProcessName string `json:"client_process_name,omitempty"`

	// client service
	ClientService string `json:"client_service,omitempty"`

	// client service id
	ClientServiceID int64 `json:"client_service_id,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// listener device
	ListenerDevice string `json:"listener_device,omitempty"`

	// listener device id
	ListenerDeviceID int64 `json:"listener_device_id,omitempty"`

	// listener discovered serivce
	ListenerDiscoveredSerivce string `json:"listener_discovered_serivce,omitempty"`

	// listener discovered service id
	ListenerDiscoveredServiceID int64 `json:"listener_discovered_service_id,omitempty"`

	// listener ip
	ListenerIP string `json:"listener_ip,omitempty"`

	// listener mapped service
	ListenerMappedService string `json:"listener_mapped_service,omitempty"`

	// listener mapped service id
	ListenerMappedServiceID int64 `json:"listener_mapped_service_id,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`

	// service port ip id
	ServicePortIPID int64 `json:"service_port_ip_id,omitempty"`

	// stats
	Stats interface{} `json:"stats,omitempty"`

	// type
	Type int64 `json:"type,omitempty"`
}

// Validate validates this service client connections client connections items0
func (m *ServiceClientConnectionsClientConnectionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceClientConnectionsClientConnectionsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceClientConnectionsClientConnectionsItems0) UnmarshalBinary(b []byte) error {
	var res ServiceClientConnectionsClientConnectionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
