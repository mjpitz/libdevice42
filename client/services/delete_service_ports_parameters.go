// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteServicePortsParams creates a new DeleteServicePortsParams object
// with the default values initialized.
func NewDeleteServicePortsParams() *DeleteServicePortsParams {
	var ()
	return &DeleteServicePortsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteServicePortsParamsWithTimeout creates a new DeleteServicePortsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteServicePortsParamsWithTimeout(timeout time.Duration) *DeleteServicePortsParams {
	var ()
	return &DeleteServicePortsParams{

		timeout: timeout,
	}
}

// NewDeleteServicePortsParamsWithContext creates a new DeleteServicePortsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteServicePortsParamsWithContext(ctx context.Context) *DeleteServicePortsParams {
	var ()
	return &DeleteServicePortsParams{

		Context: ctx,
	}
}

// NewDeleteServicePortsParamsWithHTTPClient creates a new DeleteServicePortsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteServicePortsParamsWithHTTPClient(client *http.Client) *DeleteServicePortsParams {
	var ()
	return &DeleteServicePortsParams{
		HTTPClient: client,
	}
}

/*DeleteServicePortsParams contains all the parameters to send to the API endpoint
for the delete service ports operation typically these are written to a http.Request
*/
type DeleteServicePortsParams struct {

	/*ID
	  service port id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete service ports params
func (o *DeleteServicePortsParams) WithTimeout(timeout time.Duration) *DeleteServicePortsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete service ports params
func (o *DeleteServicePortsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete service ports params
func (o *DeleteServicePortsParams) WithContext(ctx context.Context) *DeleteServicePortsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete service ports params
func (o *DeleteServicePortsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete service ports params
func (o *DeleteServicePortsParams) WithHTTPClient(client *http.Client) *DeleteServicePortsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete service ports params
func (o *DeleteServicePortsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete service ports params
func (o *DeleteServicePortsParams) WithID(id strfmt.UUID) *DeleteServicePortsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete service ports params
func (o *DeleteServicePortsParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteServicePortsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}