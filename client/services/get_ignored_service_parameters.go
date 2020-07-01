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

// NewGetIgnoredServiceParams creates a new GetIgnoredServiceParams object
// with the default values initialized.
func NewGetIgnoredServiceParams() *GetIgnoredServiceParams {
	var ()
	return &GetIgnoredServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIgnoredServiceParamsWithTimeout creates a new GetIgnoredServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIgnoredServiceParamsWithTimeout(timeout time.Duration) *GetIgnoredServiceParams {
	var ()
	return &GetIgnoredServiceParams{

		timeout: timeout,
	}
}

// NewGetIgnoredServiceParamsWithContext creates a new GetIgnoredServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIgnoredServiceParamsWithContext(ctx context.Context) *GetIgnoredServiceParams {
	var ()
	return &GetIgnoredServiceParams{

		Context: ctx,
	}
}

// NewGetIgnoredServiceParamsWithHTTPClient creates a new GetIgnoredServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIgnoredServiceParamsWithHTTPClient(client *http.Client) *GetIgnoredServiceParams {
	var ()
	return &GetIgnoredServiceParams{
		HTTPClient: client,
	}
}

/*GetIgnoredServiceParams contains all the parameters to send to the API endpoint
for the get ignored service operation typically these are written to a http.Request
*/
type GetIgnoredServiceParams struct {

	/*IgnoredID
	  service id

	*/
	IgnoredID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ignored service params
func (o *GetIgnoredServiceParams) WithTimeout(timeout time.Duration) *GetIgnoredServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ignored service params
func (o *GetIgnoredServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ignored service params
func (o *GetIgnoredServiceParams) WithContext(ctx context.Context) *GetIgnoredServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ignored service params
func (o *GetIgnoredServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ignored service params
func (o *GetIgnoredServiceParams) WithHTTPClient(client *http.Client) *GetIgnoredServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ignored service params
func (o *GetIgnoredServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIgnoredID adds the ignoredID to the get ignored service params
func (o *GetIgnoredServiceParams) WithIgnoredID(ignoredID *int64) *GetIgnoredServiceParams {
	o.SetIgnoredID(ignoredID)
	return o
}

// SetIgnoredID adds the ignoredId to the get ignored service params
func (o *GetIgnoredServiceParams) SetIgnoredID(ignoredID *int64) {
	o.IgnoredID = ignoredID
}

// WriteToRequest writes these params to a swagger request
func (o *GetIgnoredServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IgnoredID != nil {

		// query param ignored_id
		var qrIgnoredID int64
		if o.IgnoredID != nil {
			qrIgnoredID = *o.IgnoredID
		}
		qIgnoredID := swag.FormatInt64(qrIgnoredID)
		if qIgnoredID != "" {
			if err := r.SetQueryParam("ignored_id", qIgnoredID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}