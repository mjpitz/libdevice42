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

// NewPostServices2Params creates a new PostServices2Params object
// with the default values initialized.
func NewPostServices2Params() *PostServices2Params {
	var ()
	return &PostServices2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostServices2ParamsWithTimeout creates a new PostServices2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostServices2ParamsWithTimeout(timeout time.Duration) *PostServices2Params {
	var ()
	return &PostServices2Params{

		timeout: timeout,
	}
}

// NewPostServices2ParamsWithContext creates a new PostServices2Params object
// with the default values initialized, and the ability to set a context for a request
func NewPostServices2ParamsWithContext(ctx context.Context) *PostServices2Params {
	var ()
	return &PostServices2Params{

		Context: ctx,
	}
}

// NewPostServices2ParamsWithHTTPClient creates a new PostServices2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostServices2ParamsWithHTTPClient(client *http.Client) *PostServices2Params {
	var ()
	return &PostServices2Params{
		HTTPClient: client,
	}
}

/*PostServices2Params contains all the parameters to send to the API endpoint
for the post services2 operation typically these are written to a http.Request
*/
type PostServices2Params struct {

	/*Category
	  Name of category

	*/
	Category *string
	/*Description
	  description of service

	*/
	Description *string
	/*DisplayName
	  if not provided, the name is used as display name

	*/
	DisplayName *string
	/*Name
	  name for service

	*/
	Name *string
	/*Notes
	  notes

	*/
	Notes *string
	/*Vendor
	  Names of vendor

	*/
	Vendor *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post services2 params
func (o *PostServices2Params) WithTimeout(timeout time.Duration) *PostServices2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post services2 params
func (o *PostServices2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post services2 params
func (o *PostServices2Params) WithContext(ctx context.Context) *PostServices2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post services2 params
func (o *PostServices2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post services2 params
func (o *PostServices2Params) WithHTTPClient(client *http.Client) *PostServices2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post services2 params
func (o *PostServices2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategory adds the category to the post services2 params
func (o *PostServices2Params) WithCategory(category *string) *PostServices2Params {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the post services2 params
func (o *PostServices2Params) SetCategory(category *string) {
	o.Category = category
}

// WithDescription adds the description to the post services2 params
func (o *PostServices2Params) WithDescription(description *string) *PostServices2Params {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the post services2 params
func (o *PostServices2Params) SetDescription(description *string) {
	o.Description = description
}

// WithDisplayName adds the displayName to the post services2 params
func (o *PostServices2Params) WithDisplayName(displayName *string) *PostServices2Params {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the post services2 params
func (o *PostServices2Params) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WithName adds the name to the post services2 params
func (o *PostServices2Params) WithName(name *string) *PostServices2Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the post services2 params
func (o *PostServices2Params) SetName(name *string) {
	o.Name = name
}

// WithNotes adds the notes to the post services2 params
func (o *PostServices2Params) WithNotes(notes *string) *PostServices2Params {
	o.SetNotes(notes)
	return o
}

// SetNotes adds the notes to the post services2 params
func (o *PostServices2Params) SetNotes(notes *string) {
	o.Notes = notes
}

// WithVendor adds the vendor to the post services2 params
func (o *PostServices2Params) WithVendor(vendor *string) *PostServices2Params {
	o.SetVendor(vendor)
	return o
}

// SetVendor adds the vendor to the post services2 params
func (o *PostServices2Params) SetVendor(vendor *string) {
	o.Vendor = vendor
}

// WriteToRequest writes these params to a swagger request
func (o *PostServices2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Category != nil {

		// form param category
		var frCategory string
		if o.Category != nil {
			frCategory = *o.Category
		}
		fCategory := frCategory
		if fCategory != "" {
			if err := r.SetFormParam("category", fCategory); err != nil {
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

	if o.DisplayName != nil {

		// form param display_name
		var frDisplayName string
		if o.DisplayName != nil {
			frDisplayName = *o.DisplayName
		}
		fDisplayName := frDisplayName
		if fDisplayName != "" {
			if err := r.SetFormParam("display_name", fDisplayName); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// form param name
		var frName string
		if o.Name != nil {
			frName = *o.Name
		}
		fName := frName
		if fName != "" {
			if err := r.SetFormParam("name", fName); err != nil {
				return err
			}
		}

	}

	if o.Notes != nil {

		// form param notes
		var frNotes string
		if o.Notes != nil {
			frNotes = *o.Notes
		}
		fNotes := frNotes
		if fNotes != "" {
			if err := r.SetFormParam("notes", fNotes); err != nil {
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
