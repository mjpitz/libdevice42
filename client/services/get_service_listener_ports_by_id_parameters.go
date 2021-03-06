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
	"github.com/go-openapi/swag"
)

// NewGetServiceListenerPortsByIDParams creates a new GetServiceListenerPortsByIDParams object
// with the default values initialized.
func NewGetServiceListenerPortsByIDParams() *GetServiceListenerPortsByIDParams {
	var ()
	return &GetServiceListenerPortsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceListenerPortsByIDParamsWithTimeout creates a new GetServiceListenerPortsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServiceListenerPortsByIDParamsWithTimeout(timeout time.Duration) *GetServiceListenerPortsByIDParams {
	var ()
	return &GetServiceListenerPortsByIDParams{

		timeout: timeout,
	}
}

// NewGetServiceListenerPortsByIDParamsWithContext creates a new GetServiceListenerPortsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServiceListenerPortsByIDParamsWithContext(ctx context.Context) *GetServiceListenerPortsByIDParams {
	var ()
	return &GetServiceListenerPortsByIDParams{

		Context: ctx,
	}
}

// NewGetServiceListenerPortsByIDParamsWithHTTPClient creates a new GetServiceListenerPortsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServiceListenerPortsByIDParamsWithHTTPClient(client *http.Client) *GetServiceListenerPortsByIDParams {
	var ()
	return &GetServiceListenerPortsByIDParams{
		HTTPClient: client,
	}
}

/*GetServiceListenerPortsByIDParams contains all the parameters to send to the API endpoint
for the get service listener ports by ID operation typically these are written to a http.Request
*/
type GetServiceListenerPortsByIDParams struct {

	/*ID
	  service port id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service listener ports by ID params
func (o *GetServiceListenerPortsByIDParams) WithTimeout(timeout time.Duration) *GetServiceListenerPortsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service listener ports by ID params
func (o *GetServiceListenerPortsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service listener ports by ID params
func (o *GetServiceListenerPortsByIDParams) WithContext(ctx context.Context) *GetServiceListenerPortsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service listener ports by ID params
func (o *GetServiceListenerPortsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service listener ports by ID params
func (o *GetServiceListenerPortsByIDParams) WithHTTPClient(client *http.Client) *GetServiceListenerPortsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service listener ports by ID params
func (o *GetServiceListenerPortsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get service listener ports by ID params
func (o *GetServiceListenerPortsByIDParams) WithID(id int64) *GetServiceListenerPortsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get service listener ports by ID params
func (o *GetServiceListenerPortsByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceListenerPortsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
