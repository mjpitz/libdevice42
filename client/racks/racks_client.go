// Code generated by go-swagger; DO NOT EDIT.

package racks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new racks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for racks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteRacksID(params *DeleteRacksIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRacksIDOK, error)

	GetRacks(params *GetRacksParams, authInfo runtime.ClientAuthInfoWriter) (*GetRacksOK, error)

	GetRacksID(params *GetRacksIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetRacksIDOK, error)

	PostRacks(params *PostRacksParams, authInfo runtime.ClientAuthInfoWriter) (*PostRacksOK, error)

	PutCustomFieldsRack(params *PutCustomFieldsRackParams, authInfo runtime.ClientAuthInfoWriter) (*PutCustomFieldsRackOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteRacksID this API is used to delete the rack with the rack id supplied as the required argument

  Delete Rack
*/
func (a *Client) DeleteRacksID(params *DeleteRacksIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRacksIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRacksIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRacksID",
		Method:             "DELETE",
		PathPattern:        "/api/1.0/racks/{ID}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRacksIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteRacksIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteRacksID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRacks gets all racks

  This API will retrieve basic information about all racks.
*/
func (a *Client) GetRacks(params *GetRacksParams, authInfo runtime.ClientAuthInfoWriter) (*GetRacksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRacksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRacks",
		Method:             "GET",
		PathPattern:        "/api/1.0/racks/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRacksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRacksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRacks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRacksID retrieves detailed information about a specific rack including all racked devices assets and p d us

  Get a Specific Rack
*/
func (a *Client) GetRacksID(params *GetRacksIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetRacksIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRacksIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRacksID",
		Method:             "GET",
		PathPattern:        "/api/1.0/racks/{ID}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRacksIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRacksIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRacksID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostRacks creates a rack

  Create / Update Racks. Creating a new rack requires both <ul><li>name</li><li>size</li></ul><p> However, if updating a rack, you can use either</p> <ul><li>rack_id</li></ul> <p>or all of</p> <ul><li>name</li><li>room</li> <li>building</li></ul><p> If using room/building name, first combination of room name or room and building name will be used.</p>
*/
func (a *Client) PostRacks(params *PostRacksParams, authInfo runtime.ClientAuthInfoWriter) (*PostRacksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRacksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postRacks",
		Method:             "POST",
		PathPattern:        "/api/1.0/racks/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRacksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostRacksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postRacks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutCustomFieldsRack creates update custom fields for racks

  Create or update custom fields for racks. "ID" or "name" of rack is needed even when value is not being changed.
*/
func (a *Client) PutCustomFieldsRack(params *PutCustomFieldsRackParams, authInfo runtime.ClientAuthInfoWriter) (*PutCustomFieldsRackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCustomFieldsRackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putCustom_fieldsRack",
		Method:             "PUT",
		PathPattern:        "/api/1.0/custom_fields/rack/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCustomFieldsRackReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutCustomFieldsRackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putCustom_fieldsRack: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
