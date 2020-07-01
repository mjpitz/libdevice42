// Code generated by go-swagger; DO NOT EDIT.

package purchasing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new purchasing API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for purchasing API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeletePurchases(params *DeletePurchasesParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePurchasesOK, error)

	GetPurchases(params *GetPurchasesParams, authInfo runtime.ClientAuthInfoWriter) (*GetPurchasesOK, error)

	PostPurchases(params *PostPurchasesParams, authInfo runtime.ClientAuthInfoWriter) (*PostPurchasesOK, error)

	PutCustomFieldPurchases(params *PutCustomFieldPurchasesParams, authInfo runtime.ClientAuthInfoWriter) (*PutCustomFieldPurchasesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeletePurchases deletes purchase

  This API is used to delete the purchase order with the purchase order id supplied as the required argument.
*/
func (a *Client) DeletePurchases(params *DeletePurchasesParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePurchasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePurchasesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePurchases",
		Method:             "DELETE",
		PathPattern:        "/api/1.0/purchases/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePurchasesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePurchasesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePurchases: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPurchases gs e t method retrieves all purchases

  Get All Purchases
*/
func (a *Client) GetPurchases(params *GetPurchasesParams, authInfo runtime.ClientAuthInfoWriter) (*GetPurchasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPurchasesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPurchases",
		Method:             "GET",
		PathPattern:        "/api/1.0/purchases/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPurchasesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPurchasesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPurchases: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostPurchases creates update purchases

  Create / Update Purchases. Required parameters: <ul><li>order_no <b>OR</b> purchase_id</li>
*/
func (a *Client) PostPurchases(params *PostPurchasesParams, authInfo runtime.ClientAuthInfoWriter) (*PostPurchasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPurchasesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postPurchases",
		Method:             "POST",
		PathPattern:        "/api/1.0/purchases/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPurchasesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostPurchasesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postPurchases: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutCustomFieldPurchases creates update custom fields for purchases

  Custom Fields
*/
func (a *Client) PutCustomFieldPurchases(params *PutCustomFieldPurchasesParams, authInfo runtime.ClientAuthInfoWriter) (*PutCustomFieldPurchasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCustomFieldPurchasesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putCustom_FieldPurchases",
		Method:             "PUT",
		PathPattern:        "/api/1.0/custom_fields/purchases/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCustomFieldPurchasesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutCustomFieldPurchasesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putCustom_FieldPurchases: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
