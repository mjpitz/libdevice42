// Code generated by go-swagger; DO NOT EDIT.

package operating_systems

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

// NewGetOperatingSystemsParams creates a new GetOperatingSystemsParams object
// with the default values initialized.
func NewGetOperatingSystemsParams() *GetOperatingSystemsParams {
	var ()
	return &GetOperatingSystemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOperatingSystemsParamsWithTimeout creates a new GetOperatingSystemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOperatingSystemsParamsWithTimeout(timeout time.Duration) *GetOperatingSystemsParams {
	var ()
	return &GetOperatingSystemsParams{

		timeout: timeout,
	}
}

// NewGetOperatingSystemsParamsWithContext creates a new GetOperatingSystemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOperatingSystemsParamsWithContext(ctx context.Context) *GetOperatingSystemsParams {
	var ()
	return &GetOperatingSystemsParams{

		Context: ctx,
	}
}

// NewGetOperatingSystemsParamsWithHTTPClient creates a new GetOperatingSystemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOperatingSystemsParamsWithHTTPClient(client *http.Client) *GetOperatingSystemsParams {
	var ()
	return &GetOperatingSystemsParams{
		HTTPClient: client,
	}
}

/*GetOperatingSystemsParams contains all the parameters to send to the API endpoint
for the get operating systems operation typically these are written to a http.Request
*/
type GetOperatingSystemsParams struct {

	/*Category
	  Filter by OS category (ie: Linux, Windows)

	*/
	Category *string
	/*LicensedCount
	  Number of licensed instances of operating system

	*/
	LicensedCount *string
	/*NotLicensedCount
	  Number of discovered instances of operating system not set to licensed

	*/
	NotLicensedCount *string
	/*TotalCount
	  Count of IPs returned (use with offset as max results are limited to 1000)

	*/
	TotalCount *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get operating systems params
func (o *GetOperatingSystemsParams) WithTimeout(timeout time.Duration) *GetOperatingSystemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get operating systems params
func (o *GetOperatingSystemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get operating systems params
func (o *GetOperatingSystemsParams) WithContext(ctx context.Context) *GetOperatingSystemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get operating systems params
func (o *GetOperatingSystemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get operating systems params
func (o *GetOperatingSystemsParams) WithHTTPClient(client *http.Client) *GetOperatingSystemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get operating systems params
func (o *GetOperatingSystemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategory adds the category to the get operating systems params
func (o *GetOperatingSystemsParams) WithCategory(category *string) *GetOperatingSystemsParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the get operating systems params
func (o *GetOperatingSystemsParams) SetCategory(category *string) {
	o.Category = category
}

// WithLicensedCount adds the licensedCount to the get operating systems params
func (o *GetOperatingSystemsParams) WithLicensedCount(licensedCount *string) *GetOperatingSystemsParams {
	o.SetLicensedCount(licensedCount)
	return o
}

// SetLicensedCount adds the licensedCount to the get operating systems params
func (o *GetOperatingSystemsParams) SetLicensedCount(licensedCount *string) {
	o.LicensedCount = licensedCount
}

// WithNotLicensedCount adds the notLicensedCount to the get operating systems params
func (o *GetOperatingSystemsParams) WithNotLicensedCount(notLicensedCount *string) *GetOperatingSystemsParams {
	o.SetNotLicensedCount(notLicensedCount)
	return o
}

// SetNotLicensedCount adds the notLicensedCount to the get operating systems params
func (o *GetOperatingSystemsParams) SetNotLicensedCount(notLicensedCount *string) {
	o.NotLicensedCount = notLicensedCount
}

// WithTotalCount adds the totalCount to the get operating systems params
func (o *GetOperatingSystemsParams) WithTotalCount(totalCount *string) *GetOperatingSystemsParams {
	o.SetTotalCount(totalCount)
	return o
}

// SetTotalCount adds the totalCount to the get operating systems params
func (o *GetOperatingSystemsParams) SetTotalCount(totalCount *string) {
	o.TotalCount = totalCount
}

// WriteToRequest writes these params to a swagger request
func (o *GetOperatingSystemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Category != nil {

		// query param category
		var qrCategory string
		if o.Category != nil {
			qrCategory = *o.Category
		}
		qCategory := qrCategory
		if qCategory != "" {
			if err := r.SetQueryParam("category", qCategory); err != nil {
				return err
			}
		}

	}

	if o.LicensedCount != nil {

		// query param licensed_count
		var qrLicensedCount string
		if o.LicensedCount != nil {
			qrLicensedCount = *o.LicensedCount
		}
		qLicensedCount := qrLicensedCount
		if qLicensedCount != "" {
			if err := r.SetQueryParam("licensed_count", qLicensedCount); err != nil {
				return err
			}
		}

	}

	if o.NotLicensedCount != nil {

		// query param not_licensed_count
		var qrNotLicensedCount string
		if o.NotLicensedCount != nil {
			qrNotLicensedCount = *o.NotLicensedCount
		}
		qNotLicensedCount := qrNotLicensedCount
		if qNotLicensedCount != "" {
			if err := r.SetQueryParam("not_licensed_count", qNotLicensedCount); err != nil {
				return err
			}
		}

	}

	if o.TotalCount != nil {

		// query param total_count
		var qrTotalCount string
		if o.TotalCount != nil {
			qrTotalCount = *o.TotalCount
		}
		qTotalCount := qrTotalCount
		if qTotalCount != "" {
			if err := r.SetQueryParam("total_count", qTotalCount); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
