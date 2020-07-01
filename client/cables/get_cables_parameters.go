// Code generated by go-swagger; DO NOT EDIT.

package cables

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

// NewGetCablesParams creates a new GetCablesParams object
// with the default values initialized.
func NewGetCablesParams() *GetCablesParams {
	var ()
	return &GetCablesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCablesParamsWithTimeout creates a new GetCablesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCablesParamsWithTimeout(timeout time.Duration) *GetCablesParams {
	var ()
	return &GetCablesParams{

		timeout: timeout,
	}
}

// NewGetCablesParamsWithContext creates a new GetCablesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCablesParamsWithContext(ctx context.Context) *GetCablesParams {
	var ()
	return &GetCablesParams{

		Context: ctx,
	}
}

// NewGetCablesParamsWithHTTPClient creates a new GetCablesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCablesParamsWithHTTPClient(client *http.Client) *GetCablesParams {
	var ()
	return &GetCablesParams{
		HTTPClient: client,
	}
}

/*GetCablesParams contains all the parameters to send to the API endpoint
for the get cables operation typically these are written to a http.Request
*/
type GetCablesParams struct {

	/*CableID
	  filter by cable_id

	*/
	CableID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cables params
func (o *GetCablesParams) WithTimeout(timeout time.Duration) *GetCablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cables params
func (o *GetCablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cables params
func (o *GetCablesParams) WithContext(ctx context.Context) *GetCablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cables params
func (o *GetCablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cables params
func (o *GetCablesParams) WithHTTPClient(client *http.Client) *GetCablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cables params
func (o *GetCablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCableID adds the cableID to the get cables params
func (o *GetCablesParams) WithCableID(cableID *int64) *GetCablesParams {
	o.SetCableID(cableID)
	return o
}

// SetCableID adds the cableId to the get cables params
func (o *GetCablesParams) SetCableID(cableID *int64) {
	o.CableID = cableID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CableID != nil {

		// query param cable_id
		var qrCableID int64
		if o.CableID != nil {
			qrCableID = *o.CableID
		}
		qCableID := swag.FormatInt64(qrCableID)
		if qCableID != "" {
			if err := r.SetQueryParam("cable_id", qCableID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
