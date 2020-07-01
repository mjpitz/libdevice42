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

// NewDeleteDeviceMountpointsParams creates a new DeleteDeviceMountpointsParams object
// with the default values initialized.
func NewDeleteDeviceMountpointsParams() *DeleteDeviceMountpointsParams {
	var ()
	return &DeleteDeviceMountpointsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeviceMountpointsParamsWithTimeout creates a new DeleteDeviceMountpointsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDeviceMountpointsParamsWithTimeout(timeout time.Duration) *DeleteDeviceMountpointsParams {
	var ()
	return &DeleteDeviceMountpointsParams{

		timeout: timeout,
	}
}

// NewDeleteDeviceMountpointsParamsWithContext creates a new DeleteDeviceMountpointsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDeviceMountpointsParamsWithContext(ctx context.Context) *DeleteDeviceMountpointsParams {
	var ()
	return &DeleteDeviceMountpointsParams{

		Context: ctx,
	}
}

// NewDeleteDeviceMountpointsParamsWithHTTPClient creates a new DeleteDeviceMountpointsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDeviceMountpointsParamsWithHTTPClient(client *http.Client) *DeleteDeviceMountpointsParams {
	var ()
	return &DeleteDeviceMountpointsParams{
		HTTPClient: client,
	}
}

/*DeleteDeviceMountpointsParams contains all the parameters to send to the API endpoint
for the delete device mountpoints operation typically these are written to a http.Request
*/
type DeleteDeviceMountpointsParams struct {

	/*ID
	  id of the mountpoint to delete

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete device mountpoints params
func (o *DeleteDeviceMountpointsParams) WithTimeout(timeout time.Duration) *DeleteDeviceMountpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete device mountpoints params
func (o *DeleteDeviceMountpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete device mountpoints params
func (o *DeleteDeviceMountpointsParams) WithContext(ctx context.Context) *DeleteDeviceMountpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete device mountpoints params
func (o *DeleteDeviceMountpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete device mountpoints params
func (o *DeleteDeviceMountpointsParams) WithHTTPClient(client *http.Client) *DeleteDeviceMountpointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete device mountpoints params
func (o *DeleteDeviceMountpointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete device mountpoints params
func (o *DeleteDeviceMountpointsParams) WithID(id int64) *DeleteDeviceMountpointsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete device mountpoints params
func (o *DeleteDeviceMountpointsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeviceMountpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
