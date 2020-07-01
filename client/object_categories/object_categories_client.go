// Code generated by go-swagger; DO NOT EDIT.

package object_categories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new object categories API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for object categories API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetObjectCategories(params *GetObjectCategoriesParams) (*GetObjectCategoriesOK, error)

	PostObjectCategories(params *PostObjectCategoriesParams) (*PostObjectCategoriesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetObjectCategories this call will get information about object categories

  Get all Object Categories
*/
func (a *Client) GetObjectCategories(params *GetObjectCategoriesParams) (*GetObjectCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetObjectCategoriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getObject_categories",
		Method:             "GET",
		PathPattern:        "/api/1.0/object_categories/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetObjectCategoriesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetObjectCategoriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getObject_categories: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostObjectCategories this call will create update information about object categories

  Create/update Object Categories
*/
func (a *Client) PostObjectCategories(params *PostObjectCategoriesParams) (*PostObjectCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostObjectCategoriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postObject_categories",
		Method:             "POST",
		PathPattern:        "/api/1.0/object_categories/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostObjectCategoriesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostObjectCategoriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postObject_categories: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
