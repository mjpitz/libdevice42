// Code generated by go-swagger; DO NOT EDIT.

package vendors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vendors API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vendors API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteVendors(params *DeleteVendorsParams) (*DeleteVendorsOK, error)

	GetVendors(params *GetVendorsParams) (*GetVendorsOK, error)

	PostVendors(params *PostVendorsParams) (*PostVendorsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteVendors deletes vendor

  This API is used to delete the vendor with the vendor id supplied as the required argument.
*/
func (a *Client) DeleteVendors(params *DeleteVendorsParams) (*DeleteVendorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVendorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteVendors",
		Method:             "DELETE",
		PathPattern:        "/api/1.0/vendors/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteVendorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVendorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteVendors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVendors gets all vendors

  Get all Vendors
*/
func (a *Client) GetVendors(params *GetVendorsParams) (*GetVendorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVendorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVendors",
		Method:             "GET",
		PathPattern:        "/api/1.0/vendors/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetVendorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVendorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVendors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostVendors Create / Update Vendors
*/
func (a *Client) PostVendors(params *PostVendorsParams) (*PostVendorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVendorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postVendors",
		Method:             "POST",
		PathPattern:        "/api/1.0/vendors/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostVendorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostVendorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postVendors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}