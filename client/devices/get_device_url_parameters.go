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
)

// NewGetDeviceURLParams creates a new GetDeviceURLParams object
// with the default values initialized.
func NewGetDeviceURLParams() *GetDeviceURLParams {
	var ()
	return &GetDeviceURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceURLParamsWithTimeout creates a new GetDeviceURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceURLParamsWithTimeout(timeout time.Duration) *GetDeviceURLParams {
	var ()
	return &GetDeviceURLParams{

		timeout: timeout,
	}
}

// NewGetDeviceURLParamsWithContext creates a new GetDeviceURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceURLParamsWithContext(ctx context.Context) *GetDeviceURLParams {
	var ()
	return &GetDeviceURLParams{

		Context: ctx,
	}
}

// NewGetDeviceURLParamsWithHTTPClient creates a new GetDeviceURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceURLParamsWithHTTPClient(client *http.Client) *GetDeviceURLParams {
	var ()
	return &GetDeviceURLParams{
		HTTPClient: client,
	}
}

/*GetDeviceURLParams contains all the parameters to send to the API endpoint
for the get device Url operation typically these are written to a http.Request
*/
type GetDeviceURLParams struct {

	/*Device
	  name of the device

	*/
	Device *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device Url params
func (o *GetDeviceURLParams) WithTimeout(timeout time.Duration) *GetDeviceURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device Url params
func (o *GetDeviceURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device Url params
func (o *GetDeviceURLParams) WithContext(ctx context.Context) *GetDeviceURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device Url params
func (o *GetDeviceURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device Url params
func (o *GetDeviceURLParams) WithHTTPClient(client *http.Client) *GetDeviceURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device Url params
func (o *GetDeviceURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevice adds the device to the get device Url params
func (o *GetDeviceURLParams) WithDevice(device *string) *GetDeviceURLParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the get device Url params
func (o *GetDeviceURLParams) SetDevice(device *string) {
	o.Device = device
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}