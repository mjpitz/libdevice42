// Code generated by go-swagger; DO NOT EDIT.

package ip_a_m

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

// NewGetIPAMIpnatParams creates a new GetIPAMIpnatParams object
// with the default values initialized.
func NewGetIPAMIpnatParams() *GetIPAMIpnatParams {

	return &GetIPAMIpnatParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPAMIpnatParamsWithTimeout creates a new GetIPAMIpnatParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPAMIpnatParamsWithTimeout(timeout time.Duration) *GetIPAMIpnatParams {

	return &GetIPAMIpnatParams{

		timeout: timeout,
	}
}

// NewGetIPAMIpnatParamsWithContext creates a new GetIPAMIpnatParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPAMIpnatParamsWithContext(ctx context.Context) *GetIPAMIpnatParams {

	return &GetIPAMIpnatParams{

		Context: ctx,
	}
}

// NewGetIPAMIpnatParamsWithHTTPClient creates a new GetIPAMIpnatParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPAMIpnatParamsWithHTTPClient(client *http.Client) *GetIPAMIpnatParams {

	return &GetIPAMIpnatParams{
		HTTPClient: client,
	}
}

/*GetIPAMIpnatParams contains all the parameters to send to the API endpoint
for the get IP a m ipnat operation typically these are written to a http.Request
*/
type GetIPAMIpnatParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get IP a m ipnat params
func (o *GetIPAMIpnatParams) WithTimeout(timeout time.Duration) *GetIPAMIpnatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get IP a m ipnat params
func (o *GetIPAMIpnatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get IP a m ipnat params
func (o *GetIPAMIpnatParams) WithContext(ctx context.Context) *GetIPAMIpnatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get IP a m ipnat params
func (o *GetIPAMIpnatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get IP a m ipnat params
func (o *GetIPAMIpnatParams) WithHTTPClient(client *http.Client) *GetIPAMIpnatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get IP a m ipnat params
func (o *GetIPAMIpnatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPAMIpnatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}