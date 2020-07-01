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
)

// NewPostServicesParams creates a new PostServicesParams object
// with the default values initialized.
func NewPostServicesParams() *PostServicesParams {
	var ()
	return &PostServicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostServicesParamsWithTimeout creates a new PostServicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostServicesParamsWithTimeout(timeout time.Duration) *PostServicesParams {
	var ()
	return &PostServicesParams{

		timeout: timeout,
	}
}

// NewPostServicesParamsWithContext creates a new PostServicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostServicesParamsWithContext(ctx context.Context) *PostServicesParams {
	var ()
	return &PostServicesParams{

		Context: ctx,
	}
}

// NewPostServicesParamsWithHTTPClient creates a new PostServicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostServicesParamsWithHTTPClient(client *http.Client) *PostServicesParams {
	var ()
	return &PostServicesParams{
		HTTPClient: client,
	}
}

/*PostServicesParams contains all the parameters to send to the API endpoint
for the post services operation typically these are written to a http.Request
*/
type PostServicesParams struct {

	/*Category
	  name of the category

	*/
	Category *string
	/*Description*/
	Description *string
	/*DisplayName
	  if not provided, the name is used as display name

	*/
	DisplayName string
	/*Name
	  filter by name (Added in v6.0.0)

	*/
	Name *string
	/*Notes
	  Any additional notes

	*/
	Notes *string
	/*ServiceType
	  could be ignored or tracked. Default is tracked.

	*/
	ServiceType *string
	/*Vendor
	  The cloud vendor

	*/
	Vendor *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post services params
func (o *PostServicesParams) WithTimeout(timeout time.Duration) *PostServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post services params
func (o *PostServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post services params
func (o *PostServicesParams) WithContext(ctx context.Context) *PostServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post services params
func (o *PostServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post services params
func (o *PostServicesParams) WithHTTPClient(client *http.Client) *PostServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post services params
func (o *PostServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategory adds the category to the post services params
func (o *PostServicesParams) WithCategory(category *string) *PostServicesParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the post services params
func (o *PostServicesParams) SetCategory(category *string) {
	o.Category = category
}

// WithDescription adds the description to the post services params
func (o *PostServicesParams) WithDescription(description *string) *PostServicesParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the post services params
func (o *PostServicesParams) SetDescription(description *string) {
	o.Description = description
}

// WithDisplayName adds the displayName to the post services params
func (o *PostServicesParams) WithDisplayName(displayName string) *PostServicesParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the post services params
func (o *PostServicesParams) SetDisplayName(displayName string) {
	o.DisplayName = displayName
}

// WithName adds the name to the post services params
func (o *PostServicesParams) WithName(name *string) *PostServicesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post services params
func (o *PostServicesParams) SetName(name *string) {
	o.Name = name
}

// WithNotes adds the notes to the post services params
func (o *PostServicesParams) WithNotes(notes *string) *PostServicesParams {
	o.SetNotes(notes)
	return o
}

// SetNotes adds the notes to the post services params
func (o *PostServicesParams) SetNotes(notes *string) {
	o.Notes = notes
}

// WithServiceType adds the serviceType to the post services params
func (o *PostServicesParams) WithServiceType(serviceType *string) *PostServicesParams {
	o.SetServiceType(serviceType)
	return o
}

// SetServiceType adds the serviceType to the post services params
func (o *PostServicesParams) SetServiceType(serviceType *string) {
	o.ServiceType = serviceType
}

// WithVendor adds the vendor to the post services params
func (o *PostServicesParams) WithVendor(vendor *string) *PostServicesParams {
	o.SetVendor(vendor)
	return o
}

// SetVendor adds the vendor to the post services params
func (o *PostServicesParams) SetVendor(vendor *string) {
	o.Vendor = vendor
}

// WriteToRequest writes these params to a swagger request
func (o *PostServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Description != nil {

		// form param description
		var frDescription string
		if o.Description != nil {
			frDescription = *o.Description
		}
		fDescription := frDescription
		if fDescription != "" {
			if err := r.SetFormParam("description", fDescription); err != nil {
				return err
			}
		}

	}

	// query param display_name
	qrDisplayName := o.DisplayName
	qDisplayName := qrDisplayName
	if qDisplayName != "" {
		if err := r.SetQueryParam("display_name", qDisplayName); err != nil {
			return err
		}
	}

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

	if o.Notes != nil {

		// query param notes
		var qrNotes string
		if o.Notes != nil {
			qrNotes = *o.Notes
		}
		qNotes := qrNotes
		if qNotes != "" {
			if err := r.SetQueryParam("notes", qNotes); err != nil {
				return err
			}
		}

	}

	if o.ServiceType != nil {

		// query param service_type
		var qrServiceType string
		if o.ServiceType != nil {
			qrServiceType = *o.ServiceType
		}
		qServiceType := qrServiceType
		if qServiceType != "" {
			if err := r.SetQueryParam("service_type", qServiceType); err != nil {
				return err
			}
		}

	}

	if o.Vendor != nil {

		// form param vendor
		var frVendor string
		if o.Vendor != nil {
			frVendor = *o.Vendor
		}
		fVendor := frVendor
		if fVendor != "" {
			if err := r.SetFormParam("vendor", fVendor); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
