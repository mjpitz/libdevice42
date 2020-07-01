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
)

// NewPostUpdateSoftwareLicensesParams creates a new PostUpdateSoftwareLicensesParams object
// with the default values initialized.
func NewPostUpdateSoftwareLicensesParams() *PostUpdateSoftwareLicensesParams {
	var ()
	return &PostUpdateSoftwareLicensesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUpdateSoftwareLicensesParamsWithTimeout creates a new PostUpdateSoftwareLicensesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUpdateSoftwareLicensesParamsWithTimeout(timeout time.Duration) *PostUpdateSoftwareLicensesParams {
	var ()
	return &PostUpdateSoftwareLicensesParams{

		timeout: timeout,
	}
}

// NewPostUpdateSoftwareLicensesParamsWithContext creates a new PostUpdateSoftwareLicensesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUpdateSoftwareLicensesParamsWithContext(ctx context.Context) *PostUpdateSoftwareLicensesParams {
	var ()
	return &PostUpdateSoftwareLicensesParams{

		Context: ctx,
	}
}

// NewPostUpdateSoftwareLicensesParamsWithHTTPClient creates a new PostUpdateSoftwareLicensesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUpdateSoftwareLicensesParamsWithHTTPClient(client *http.Client) *PostUpdateSoftwareLicensesParams {
	var ()
	return &PostUpdateSoftwareLicensesParams{
		HTTPClient: client,
	}
}

/*PostUpdateSoftwareLicensesParams contains all the parameters to send to the API endpoint
for the post update software licenses operation typically these are written to a http.Request
*/
type PostUpdateSoftwareLicensesParams struct {

	/*ID
	  The ID of the software_license_key object (use if updating)

	*/
	ID *string
	/*Key
	  software license key

	*/
	Key *string
	/*SoftwareID
	  The id of the software component

	*/
	SoftwareID *string
	/*SoftwareName
	  software component name

	*/
	SoftwareName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post update software licenses params
func (o *PostUpdateSoftwareLicensesParams) WithTimeout(timeout time.Duration) *PostUpdateSoftwareLicensesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post update software licenses params
func (o *PostUpdateSoftwareLicensesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post update software licenses params
func (o *PostUpdateSoftwareLicensesParams) WithContext(ctx context.Context) *PostUpdateSoftwareLicensesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post update software licenses params
func (o *PostUpdateSoftwareLicensesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post update software licenses params
func (o *PostUpdateSoftwareLicensesParams) WithHTTPClient(client *http.Client) *PostUpdateSoftwareLicensesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post update software licenses params
func (o *PostUpdateSoftwareLicensesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post update software licenses params
func (o *PostUpdateSoftwareLicensesParams) WithID(id *string) *PostUpdateSoftwareLicensesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post update software licenses params
func (o *PostUpdateSoftwareLicensesParams) SetID(id *string) {
	o.ID = id
}

// WithKey adds the key to the post update software licenses params
func (o *PostUpdateSoftwareLicensesParams) WithKey(key *string) *PostUpdateSoftwareLicensesParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the post update software licenses params
func (o *PostUpdateSoftwareLicensesParams) SetKey(key *string) {
	o.Key = key
}

// WithSoftwareID adds the softwareID to the post update software licenses params
func (o *PostUpdateSoftwareLicensesParams) WithSoftwareID(softwareID *string) *PostUpdateSoftwareLicensesParams {
	o.SetSoftwareID(softwareID)
	return o
}

// SetSoftwareID adds the softwareId to the post update software licenses params
func (o *PostUpdateSoftwareLicensesParams) SetSoftwareID(softwareID *string) {
	o.SoftwareID = softwareID
}

// WithSoftwareName adds the softwareName to the post update software licenses params
func (o *PostUpdateSoftwareLicensesParams) WithSoftwareName(softwareName *string) *PostUpdateSoftwareLicensesParams {
	o.SetSoftwareName(softwareName)
	return o
}

// SetSoftwareName adds the softwareName to the post update software licenses params
func (o *PostUpdateSoftwareLicensesParams) SetSoftwareName(softwareName *string) {
	o.SoftwareName = softwareName
}

// WriteToRequest writes these params to a swagger request
func (o *PostUpdateSoftwareLicensesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// form param id
		var frID string
		if o.ID != nil {
			frID = *o.ID
		}
		fID := frID
		if fID != "" {
			if err := r.SetFormParam("id", fID); err != nil {
				return err
			}
		}

	}

	if o.Key != nil {

		// form param key
		var frKey string
		if o.Key != nil {
			frKey = *o.Key
		}
		fKey := frKey
		if fKey != "" {
			if err := r.SetFormParam("key", fKey); err != nil {
				return err
			}
		}

	}

	if o.SoftwareID != nil {

		// form param software_id
		var frSoftwareID string
		if o.SoftwareID != nil {
			frSoftwareID = *o.SoftwareID
		}
		fSoftwareID := frSoftwareID
		if fSoftwareID != "" {
			if err := r.SetFormParam("software_id", fSoftwareID); err != nil {
				return err
			}
		}

	}

	if o.SoftwareName != nil {

		// form param software_name
		var frSoftwareName string
		if o.SoftwareName != nil {
			frSoftwareName = *o.SoftwareName
		}
		fSoftwareName := frSoftwareName
		if fSoftwareName != "" {
			if err := r.SetFormParam("software_name", fSoftwareName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
