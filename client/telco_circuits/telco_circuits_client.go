// Code generated by go-swagger; DO NOT EDIT.

package telco_circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new telco circuits API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for telco circuits API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteCircuit(params *DeleteCircuitParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCircuitOK, error)

	GetCircuits(params *GetCircuitsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCircuitsOK, error)

	PostUpdateCircuits(params *PostUpdateCircuitsParams, authInfo runtime.ClientAuthInfoWriter) (*PostUpdateCircuitsOK, error)

	PutCustomFields(params *PutCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter) (*PutCustomFieldsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteCircuit deletes circuit

  This API is used to delete the circuit with the circuit id supplied as the required argument.
*/
func (a *Client) DeleteCircuit(params *DeleteCircuitParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCircuitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCircuitParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCircuit",
		Method:             "DELETE",
		PathPattern:        "/api/1.0/circuits/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCircuitReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCircuitOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCircuit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCircuits gets all circuits

  Get all Circuits
*/
func (a *Client) GetCircuits(params *GetCircuitsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCircuitsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCircuitsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCircuits",
		Method:             "GET",
		PathPattern:        "/api/1.0/circuits/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCircuitsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCircuitsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCircuits: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUpdateCircuits creates update circuits

  Create / Update Circuits
*/
func (a *Client) PostUpdateCircuits(params *PostUpdateCircuitsParams, authInfo runtime.ClientAuthInfoWriter) (*PostUpdateCircuitsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUpdateCircuitsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "post_Update_Circuits",
		Method:             "POST",
		PathPattern:        "/api/1.0/circuits/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostUpdateCircuitsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUpdateCircuitsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_Update_Circuits: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutCustomFields creates update custom fields for circuits added in v5 9 3

  Custom Fields for circuits. Either ID or circuit_id is required.
*/
func (a *Client) PutCustomFields(params *PutCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter) (*PutCustomFieldsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCustomFieldsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "put_Custom_Fields",
		Method:             "PUT",
		PathPattern:        "/api/1.0/custom_fields/circuit/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCustomFieldsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutCustomFieldsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for put_Custom_Fields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
