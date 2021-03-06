// Code generated by go-swagger; DO NOT EDIT.

package software

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

// NewDeleteSoftwareLicenseKeysParams creates a new DeleteSoftwareLicenseKeysParams object
// with the default values initialized.
func NewDeleteSoftwareLicenseKeysParams() *DeleteSoftwareLicenseKeysParams {
	var ()
	return &DeleteSoftwareLicenseKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSoftwareLicenseKeysParamsWithTimeout creates a new DeleteSoftwareLicenseKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSoftwareLicenseKeysParamsWithTimeout(timeout time.Duration) *DeleteSoftwareLicenseKeysParams {
	var ()
	return &DeleteSoftwareLicenseKeysParams{

		timeout: timeout,
	}
}

// NewDeleteSoftwareLicenseKeysParamsWithContext creates a new DeleteSoftwareLicenseKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSoftwareLicenseKeysParamsWithContext(ctx context.Context) *DeleteSoftwareLicenseKeysParams {
	var ()
	return &DeleteSoftwareLicenseKeysParams{

		Context: ctx,
	}
}

// NewDeleteSoftwareLicenseKeysParamsWithHTTPClient creates a new DeleteSoftwareLicenseKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSoftwareLicenseKeysParamsWithHTTPClient(client *http.Client) *DeleteSoftwareLicenseKeysParams {
	var ()
	return &DeleteSoftwareLicenseKeysParams{
		HTTPClient: client,
	}
}

/*DeleteSoftwareLicenseKeysParams contains all the parameters to send to the API endpoint
for the delete software license keys operation typically these are written to a http.Request
*/
type DeleteSoftwareLicenseKeysParams struct {

	/*ID
	  software detail id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete software license keys params
func (o *DeleteSoftwareLicenseKeysParams) WithTimeout(timeout time.Duration) *DeleteSoftwareLicenseKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete software license keys params
func (o *DeleteSoftwareLicenseKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete software license keys params
func (o *DeleteSoftwareLicenseKeysParams) WithContext(ctx context.Context) *DeleteSoftwareLicenseKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete software license keys params
func (o *DeleteSoftwareLicenseKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete software license keys params
func (o *DeleteSoftwareLicenseKeysParams) WithHTTPClient(client *http.Client) *DeleteSoftwareLicenseKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete software license keys params
func (o *DeleteSoftwareLicenseKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete software license keys params
func (o *DeleteSoftwareLicenseKeysParams) WithID(id int64) *DeleteSoftwareLicenseKeysParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete software license keys params
func (o *DeleteSoftwareLicenseKeysParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSoftwareLicenseKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
