// Code generated by go-swagger; DO NOT EDIT.

package operating_systems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operating systems API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operating systems API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteDeviceOs(params *DeleteDeviceOsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDeviceOsOK, error)

	DeleteOperatingSystems(params *DeleteOperatingSystemsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOperatingSystemsOK, error)

	GetDeviceOs(params *GetDeviceOsParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceOsOK, error)

	GetOperatingSystems(params *GetOperatingSystemsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOperatingSystemsOK, error)

	PostDeviceOs(params *PostDeviceOsParams, authInfo runtime.ClientAuthInfoWriter) (*PostDeviceOsOK, error)

	PostOperatingSystems(params *PostOperatingSystemsParams, authInfo runtime.ClientAuthInfoWriter) (*PostOperatingSystemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteDeviceOs deletes operating system by o s ID

  This API is used to delete the operating system with the operating system id supplied as the required argument.
*/
func (a *Client) DeleteDeviceOs(params *DeleteDeviceOsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDeviceOsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeviceOsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDevice_os",
		Method:             "DELETE",
		PathPattern:        "/api/1.0/device_os/{device_os_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDeviceOsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDeviceOsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDevice_os: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteOperatingSystems deletes operating system

  This API is used to delete the operating system with the operating system id supplied as the required argument.
*/
func (a *Client) DeleteOperatingSystems(params *DeleteOperatingSystemsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOperatingSystemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOperatingSystemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOperatingSystems",
		Method:             "DELETE",
		PathPattern:        "/api/1.0/operatingsystems/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteOperatingSystemsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteOperatingSystemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteOperatingSystems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeviceOs gets operating systems by devices

  This call will get information about operating systems and the devices they’re discovered on.
*/
func (a *Client) GetDeviceOs(params *GetDeviceOsParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceOsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceOsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDevice_os",
		Method:             "GET",
		PathPattern:        "/api/1.0/device_os/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDeviceOsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceOsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDevice_os: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOperatingSystems this call will get information about operating systems

  Get all operating systems
*/
func (a *Client) GetOperatingSystems(params *GetOperatingSystemsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOperatingSystemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOperatingSystemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOperatingSystems",
		Method:             "GET",
		PathPattern:        "/api/1.0/operatingsystems/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOperatingSystemsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOperatingSystemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOperatingSystems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostDeviceOs this call will create or update operating systems and assign them to a device

  Create/Update operating systems on devices
*/
func (a *Client) PostDeviceOs(params *PostDeviceOsParams, authInfo runtime.ClientAuthInfoWriter) (*PostDeviceOsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDeviceOsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postDevice_os",
		Method:             "POST",
		PathPattern:        "/api/1.0/device_os/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostDeviceOsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostDeviceOsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postDevice_os: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostOperatingSystems this call will create update information about operating systems

  Create/update OS
*/
func (a *Client) PostOperatingSystems(params *PostOperatingSystemsParams, authInfo runtime.ClientAuthInfoWriter) (*PostOperatingSystemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOperatingSystemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postOperatingSystems",
		Method:             "POST",
		PathPattern:        "/api/1.0/operatingsystems/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostOperatingSystemsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostOperatingSystemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postOperatingSystems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
