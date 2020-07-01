// Code generated by go-swagger; DO NOT EDIT.

package power_circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new power circuits API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for power circuits API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeletePowerCircuit(params *DeletePowerCircuitParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePowerCircuitOK, error)

	GetAllPowerCircuits(params *GetAllPowerCircuitsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllPowerCircuitsOK, error)

	PostUpdatePowerCircuits(params *PostUpdatePowerCircuitsParams, authInfo runtime.ClientAuthInfoWriter) (*PostUpdatePowerCircuitsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeletePowerCircuit powers circuit

  This API is used to delete the power circuit with the ID supplied as the required argument.
*/
func (a *Client) DeletePowerCircuit(params *DeletePowerCircuitParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePowerCircuitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePowerCircuitParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePowerCircuit",
		Method:             "DELETE",
		PathPattern:        "/api/1.0/powercircuits/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePowerCircuitReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePowerCircuitOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePowerCircuit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllPowerCircuits gets all power circuits

  Get All Power Circuits
*/
func (a *Client) GetAllPowerCircuits(params *GetAllPowerCircuitsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllPowerCircuitsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllPowerCircuitsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_All_Power_Circuits",
		Method:             "GET",
		PathPattern:        "/api/1.0/powercircuits/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAllPowerCircuitsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllPowerCircuitsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_All_Power_Circuits: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUpdatePowerCircuits creates update power circuit

  Create/Update Power Circuit
*/
func (a *Client) PostUpdatePowerCircuits(params *PostUpdatePowerCircuitsParams, authInfo runtime.ClientAuthInfoWriter) (*PostUpdatePowerCircuitsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUpdatePowerCircuitsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "post_Update_PowerCircuits",
		Method:             "POST",
		PathPattern:        "/api/1.0/powercircuits/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostUpdatePowerCircuitsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUpdatePowerCircuitsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_Update_PowerCircuits: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
