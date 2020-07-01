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

// NewDeleteDevicesIDParams creates a new DeleteDevicesIDParams object
// with the default values initialized.
func NewDeleteDevicesIDParams() *DeleteDevicesIDParams {
	var ()
	return &DeleteDevicesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDevicesIDParamsWithTimeout creates a new DeleteDevicesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDevicesIDParamsWithTimeout(timeout time.Duration) *DeleteDevicesIDParams {
	var ()
	return &DeleteDevicesIDParams{

		timeout: timeout,
	}
}

// NewDeleteDevicesIDParamsWithContext creates a new DeleteDevicesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDevicesIDParamsWithContext(ctx context.Context) *DeleteDevicesIDParams {
	var ()
	return &DeleteDevicesIDParams{

		Context: ctx,
	}
}

// NewDeleteDevicesIDParamsWithHTTPClient creates a new DeleteDevicesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDevicesIDParamsWithHTTPClient(client *http.Client) *DeleteDevicesIDParams {
	var ()
	return &DeleteDevicesIDParams{
		HTTPClient: client,
	}
}

/*DeleteDevicesIDParams contains all the parameters to send to the API endpoint
for the delete devices Id operation typically these are written to a http.Request
*/
type DeleteDevicesIDParams struct {

	/*ID
	  Device id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete devices Id params
func (o *DeleteDevicesIDParams) WithTimeout(timeout time.Duration) *DeleteDevicesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete devices Id params
func (o *DeleteDevicesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete devices Id params
func (o *DeleteDevicesIDParams) WithContext(ctx context.Context) *DeleteDevicesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete devices Id params
func (o *DeleteDevicesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete devices Id params
func (o *DeleteDevicesIDParams) WithHTTPClient(client *http.Client) *DeleteDevicesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete devices Id params
func (o *DeleteDevicesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete devices Id params
func (o *DeleteDevicesIDParams) WithID(id int64) *DeleteDevicesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete devices Id params
func (o *DeleteDevicesIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDevicesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
