// Code generated by go-swagger; DO NOT EDIT.

package service_levels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service levels API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service levels API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetServiceLevel(params *GetServiceLevelParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceLevelOK, error)

	PostServiceLevel(params *PostServiceLevelParams, authInfo runtime.ClientAuthInfoWriter) (*PostServiceLevelOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetServiceLevel gets all service levels devices

  Get all Service Levels
*/
func (a *Client) GetServiceLevel(params *GetServiceLevelParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceLevelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceLevelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getService_level",
		Method:             "GET",
		PathPattern:        "/api/1.0/service_level/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServiceLevelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceLevelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getService_level: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostServiceLevel creates service level

  Create Service Level
*/
func (a *Client) PostServiceLevel(params *PostServiceLevelParams, authInfo runtime.ClientAuthInfoWriter) (*PostServiceLevelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostServiceLevelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postService_level",
		Method:             "POST",
		PathPattern:        "/api/1.0/service_level/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostServiceLevelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostServiceLevelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postService_level: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
