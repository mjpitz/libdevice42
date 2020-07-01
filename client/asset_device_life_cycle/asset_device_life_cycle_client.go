// Code generated by go-swagger; DO NOT EDIT.

package asset_device_life_cycle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new asset device life cycle API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for asset device life cycle API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetLifecycleEvent(params *GetLifecycleEventParams) (*GetLifecycleEventOK, error)

	PutLifecycleEvent(params *PutLifecycleEventParams) (*PutLifecycleEventOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetLifecycleEvent retrieves life cycle events using filters introduced in version 5 5 7

  Get Life Cycle Events
*/
func (a *Client) GetLifecycleEvent(params *GetLifecycleEventParams) (*GetLifecycleEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLifecycleEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLifecycle_event",
		Method:             "GET",
		PathPattern:        "/api/1.0/lifecycle_event/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLifecycleEventReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLifecycleEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLifecycle_event: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutLifecycleEvent creates life cycle events

  Use this API to create life cycle events for devices or assets. Use device, device_id, asset_no, serial_no, or asset_id to indicate the device or asset the event is to be created for.
*/
func (a *Client) PutLifecycleEvent(params *PutLifecycleEventParams) (*PutLifecycleEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutLifecycleEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putLifecycle_event",
		Method:             "PUT",
		PathPattern:        "/api/1.0/lifecycle_event/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutLifecycleEventReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutLifecycleEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putLifecycle_event: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
