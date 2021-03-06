// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewDeleteDeviceURLParams creates a new DeleteDeviceURLParams object
// with the default values initialized.
func NewDeleteDeviceURLParams() *DeleteDeviceURLParams {
	var ()
	return &DeleteDeviceURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeviceURLParamsWithTimeout creates a new DeleteDeviceURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDeviceURLParamsWithTimeout(timeout time.Duration) *DeleteDeviceURLParams {
	var ()
	return &DeleteDeviceURLParams{

		timeout: timeout,
	}
}

// NewDeleteDeviceURLParamsWithContext creates a new DeleteDeviceURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDeviceURLParamsWithContext(ctx context.Context) *DeleteDeviceURLParams {
	var ()
	return &DeleteDeviceURLParams{

		Context: ctx,
	}
}

// NewDeleteDeviceURLParamsWithHTTPClient creates a new DeleteDeviceURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDeviceURLParamsWithHTTPClient(client *http.Client) *DeleteDeviceURLParams {
	var ()
	return &DeleteDeviceURLParams{
		HTTPClient: client,
	}
}

/*DeleteDeviceURLParams contains all the parameters to send to the API endpoint
for the delete device Url operation typically these are written to a http.Request
*/
type DeleteDeviceURLParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete device Url params
func (o *DeleteDeviceURLParams) WithTimeout(timeout time.Duration) *DeleteDeviceURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete device Url params
func (o *DeleteDeviceURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete device Url params
func (o *DeleteDeviceURLParams) WithContext(ctx context.Context) *DeleteDeviceURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete device Url params
func (o *DeleteDeviceURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete device Url params
func (o *DeleteDeviceURLParams) WithHTTPClient(client *http.Client) *DeleteDeviceURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete device Url params
func (o *DeleteDeviceURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete device Url params
func (o *DeleteDeviceURLParams) WithID(id int64) *DeleteDeviceURLParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete device Url params
func (o *DeleteDeviceURLParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeviceURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
