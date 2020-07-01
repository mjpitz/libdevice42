// Code generated by go-swagger; DO NOT EDIT.

package vendors

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

// NewDeleteVendorsParams creates a new DeleteVendorsParams object
// with the default values initialized.
func NewDeleteVendorsParams() *DeleteVendorsParams {
	var ()
	return &DeleteVendorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVendorsParamsWithTimeout creates a new DeleteVendorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteVendorsParamsWithTimeout(timeout time.Duration) *DeleteVendorsParams {
	var ()
	return &DeleteVendorsParams{

		timeout: timeout,
	}
}

// NewDeleteVendorsParamsWithContext creates a new DeleteVendorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteVendorsParamsWithContext(ctx context.Context) *DeleteVendorsParams {
	var ()
	return &DeleteVendorsParams{

		Context: ctx,
	}
}

// NewDeleteVendorsParamsWithHTTPClient creates a new DeleteVendorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteVendorsParamsWithHTTPClient(client *http.Client) *DeleteVendorsParams {
	var ()
	return &DeleteVendorsParams{
		HTTPClient: client,
	}
}

/*DeleteVendorsParams contains all the parameters to send to the API endpoint
for the delete vendors operation typically these are written to a http.Request
*/
type DeleteVendorsParams struct {

	/*ID
	  IP Address id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete vendors params
func (o *DeleteVendorsParams) WithTimeout(timeout time.Duration) *DeleteVendorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vendors params
func (o *DeleteVendorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vendors params
func (o *DeleteVendorsParams) WithContext(ctx context.Context) *DeleteVendorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vendors params
func (o *DeleteVendorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vendors params
func (o *DeleteVendorsParams) WithHTTPClient(client *http.Client) *DeleteVendorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vendors params
func (o *DeleteVendorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete vendors params
func (o *DeleteVendorsParams) WithID(id int64) *DeleteVendorsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete vendors params
func (o *DeleteVendorsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVendorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
