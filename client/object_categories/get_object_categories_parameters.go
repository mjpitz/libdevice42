// Code generated by go-swagger; DO NOT EDIT.

package object_categories

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

// NewGetObjectCategoriesParams creates a new GetObjectCategoriesParams object
// with the default values initialized.
func NewGetObjectCategoriesParams() *GetObjectCategoriesParams {
	var ()
	return &GetObjectCategoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetObjectCategoriesParamsWithTimeout creates a new GetObjectCategoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetObjectCategoriesParamsWithTimeout(timeout time.Duration) *GetObjectCategoriesParams {
	var ()
	return &GetObjectCategoriesParams{

		timeout: timeout,
	}
}

// NewGetObjectCategoriesParamsWithContext creates a new GetObjectCategoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetObjectCategoriesParamsWithContext(ctx context.Context) *GetObjectCategoriesParams {
	var ()
	return &GetObjectCategoriesParams{

		Context: ctx,
	}
}

// NewGetObjectCategoriesParamsWithHTTPClient creates a new GetObjectCategoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetObjectCategoriesParamsWithHTTPClient(client *http.Client) *GetObjectCategoriesParams {
	var ()
	return &GetObjectCategoriesParams{
		HTTPClient: client,
	}
}

/*GetObjectCategoriesParams contains all the parameters to send to the API endpoint
for the get object categories operation typically these are written to a http.Request
*/
type GetObjectCategoriesParams struct {

	/*Name
	  filter by name (Added in v6.0.0)

	*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get object categories params
func (o *GetObjectCategoriesParams) WithTimeout(timeout time.Duration) *GetObjectCategoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get object categories params
func (o *GetObjectCategoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get object categories params
func (o *GetObjectCategoriesParams) WithContext(ctx context.Context) *GetObjectCategoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get object categories params
func (o *GetObjectCategoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get object categories params
func (o *GetObjectCategoriesParams) WithHTTPClient(client *http.Client) *GetObjectCategoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get object categories params
func (o *GetObjectCategoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get object categories params
func (o *GetObjectCategoriesParams) WithName(name *string) *GetObjectCategoriesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get object categories params
func (o *GetObjectCategoriesParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetObjectCategoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}