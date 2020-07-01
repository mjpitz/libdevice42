// Code generated by go-swagger; DO NOT EDIT.

package cables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cables API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cables API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteCable(params *DeleteCableParams) (*DeleteCableOK, error)

	GetCables(params *GetCablesParams) (*GetCablesOK, error)

	PostCables(params *PostCablesParams) (*PostCablesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteCable deletes cable

  This API is used to delete the cable with the ID supplied as the required argument.
*/
func (a *Client) DeleteCable(params *DeleteCableParams) (*DeleteCableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCable",
		Method:             "DELETE",
		PathPattern:        "/api/1.0/cables/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCableReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCables retrieves information about all cables

  Get All Cables
*/
func (a *Client) GetCables(params *GetCablesParams) (*GetCablesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCablesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCables",
		Method:             "GET",
		PathPattern:        "/api/1.0/cables/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCablesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCablesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCables: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostCables creates or update cables

  Create/Update Cable
*/
func (a *Client) PostCables(params *PostCablesParams) (*PostCablesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCablesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postCables",
		Method:             "POST",
		PathPattern:        "/api/1.0/cables/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostCablesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCablesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postCables: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
