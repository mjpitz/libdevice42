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

// NewGetDevicesImpactlistParams creates a new GetDevicesImpactlistParams object
// with the default values initialized.
func NewGetDevicesImpactlistParams() *GetDevicesImpactlistParams {
	var ()
	return &GetDevicesImpactlistParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesImpactlistParamsWithTimeout creates a new GetDevicesImpactlistParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDevicesImpactlistParamsWithTimeout(timeout time.Duration) *GetDevicesImpactlistParams {
	var ()
	return &GetDevicesImpactlistParams{

		timeout: timeout,
	}
}

// NewGetDevicesImpactlistParamsWithContext creates a new GetDevicesImpactlistParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDevicesImpactlistParamsWithContext(ctx context.Context) *GetDevicesImpactlistParams {
	var ()
	return &GetDevicesImpactlistParams{

		Context: ctx,
	}
}

// NewGetDevicesImpactlistParamsWithHTTPClient creates a new GetDevicesImpactlistParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDevicesImpactlistParamsWithHTTPClient(client *http.Client) *GetDevicesImpactlistParams {
	var ()
	return &GetDevicesImpactlistParams{
		HTTPClient: client,
	}
}

/*GetDevicesImpactlistParams contains all the parameters to send to the API endpoint
for the get devices impactlist operation typically these are written to a http.Request
*/
type GetDevicesImpactlistParams struct {

	/*DeviceID*/
	DeviceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get devices impactlist params
func (o *GetDevicesImpactlistParams) WithTimeout(timeout time.Duration) *GetDevicesImpactlistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices impactlist params
func (o *GetDevicesImpactlistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices impactlist params
func (o *GetDevicesImpactlistParams) WithContext(ctx context.Context) *GetDevicesImpactlistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices impactlist params
func (o *GetDevicesImpactlistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices impactlist params
func (o *GetDevicesImpactlistParams) WithHTTPClient(client *http.Client) *GetDevicesImpactlistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices impactlist params
func (o *GetDevicesImpactlistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get devices impactlist params
func (o *GetDevicesImpactlistParams) WithDeviceID(deviceID int64) *GetDevicesImpactlistParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get devices impactlist params
func (o *GetDevicesImpactlistParams) SetDeviceID(deviceID int64) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesImpactlistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param device-id
	if err := r.SetPathParam("device-id", swag.FormatInt64(o.DeviceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
