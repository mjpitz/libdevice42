// Code generated by go-swagger; DO NOT EDIT.

package audit_log

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

// NewGetAuditlogsParams creates a new GetAuditlogsParams object
// with the default values initialized.
func NewGetAuditlogsParams() *GetAuditlogsParams {
	var ()
	return &GetAuditlogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuditlogsParamsWithTimeout creates a new GetAuditlogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuditlogsParamsWithTimeout(timeout time.Duration) *GetAuditlogsParams {
	var ()
	return &GetAuditlogsParams{

		timeout: timeout,
	}
}

// NewGetAuditlogsParamsWithContext creates a new GetAuditlogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuditlogsParamsWithContext(ctx context.Context) *GetAuditlogsParams {
	var ()
	return &GetAuditlogsParams{

		Context: ctx,
	}
}

// NewGetAuditlogsParamsWithHTTPClient creates a new GetAuditlogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuditlogsParamsWithHTTPClient(client *http.Client) *GetAuditlogsParams {
	var ()
	return &GetAuditlogsParams{
		HTTPClient: client,
	}
}

/*GetAuditlogsParams contains all the parameters to send to the API endpoint
for the get auditlogs operation typically these are written to a http.Request
*/
type GetAuditlogsParams struct {

	/*ActionTimeGt
	  Filters actions that have happened past the time entered (ie, greater than 2 weeks) in YYYY-MM-DDTHH:MM:ss.uuuuuu (ie 2016-10-27T13:52:01.213416)

	*/
	ActionTimeGt *string
	/*ActionTimeLt
	  Returns actions within the last X amount of days in YYYY-MM-DDTHH:MM:ss.uuuuuu (ie 2016-10-27T13:52:01.213416)

	*/
	ActionTimeLt *string
	/*ContentType
	  Returns changes done to a particular content type

	*/
	ContentType *string
	/*ObjectID
	  Filters by object ID (ie, device ID, asset ID)

	*/
	ObjectID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get auditlogs params
func (o *GetAuditlogsParams) WithTimeout(timeout time.Duration) *GetAuditlogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get auditlogs params
func (o *GetAuditlogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get auditlogs params
func (o *GetAuditlogsParams) WithContext(ctx context.Context) *GetAuditlogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get auditlogs params
func (o *GetAuditlogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get auditlogs params
func (o *GetAuditlogsParams) WithHTTPClient(client *http.Client) *GetAuditlogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get auditlogs params
func (o *GetAuditlogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionTimeGt adds the actionTimeGt to the get auditlogs params
func (o *GetAuditlogsParams) WithActionTimeGt(actionTimeGt *string) *GetAuditlogsParams {
	o.SetActionTimeGt(actionTimeGt)
	return o
}

// SetActionTimeGt adds the actionTimeGt to the get auditlogs params
func (o *GetAuditlogsParams) SetActionTimeGt(actionTimeGt *string) {
	o.ActionTimeGt = actionTimeGt
}

// WithActionTimeLt adds the actionTimeLt to the get auditlogs params
func (o *GetAuditlogsParams) WithActionTimeLt(actionTimeLt *string) *GetAuditlogsParams {
	o.SetActionTimeLt(actionTimeLt)
	return o
}

// SetActionTimeLt adds the actionTimeLt to the get auditlogs params
func (o *GetAuditlogsParams) SetActionTimeLt(actionTimeLt *string) {
	o.ActionTimeLt = actionTimeLt
}

// WithContentType adds the contentType to the get auditlogs params
func (o *GetAuditlogsParams) WithContentType(contentType *string) *GetAuditlogsParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the get auditlogs params
func (o *GetAuditlogsParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithObjectID adds the objectID to the get auditlogs params
func (o *GetAuditlogsParams) WithObjectID(objectID *string) *GetAuditlogsParams {
	o.SetObjectID(objectID)
	return o
}

// SetObjectID adds the objectId to the get auditlogs params
func (o *GetAuditlogsParams) SetObjectID(objectID *string) {
	o.ObjectID = objectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuditlogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActionTimeGt != nil {

		// query param action_time_gt
		var qrActionTimeGt string
		if o.ActionTimeGt != nil {
			qrActionTimeGt = *o.ActionTimeGt
		}
		qActionTimeGt := qrActionTimeGt
		if qActionTimeGt != "" {
			if err := r.SetQueryParam("action_time_gt", qActionTimeGt); err != nil {
				return err
			}
		}

	}

	if o.ActionTimeLt != nil {

		// query param action_time_lt
		var qrActionTimeLt string
		if o.ActionTimeLt != nil {
			qrActionTimeLt = *o.ActionTimeLt
		}
		qActionTimeLt := qrActionTimeLt
		if qActionTimeLt != "" {
			if err := r.SetQueryParam("action_time_lt", qActionTimeLt); err != nil {
				return err
			}
		}

	}

	if o.ContentType != nil {

		// query param content_type
		var qrContentType string
		if o.ContentType != nil {
			qrContentType = *o.ContentType
		}
		qContentType := qrContentType
		if qContentType != "" {
			if err := r.SetQueryParam("content_type", qContentType); err != nil {
				return err
			}
		}

	}

	if o.ObjectID != nil {

		// query param object_id
		var qrObjectID string
		if o.ObjectID != nil {
			qrObjectID = *o.ObjectID
		}
		qObjectID := qrObjectID
		if qObjectID != "" {
			if err := r.SetQueryParam("object_id", qObjectID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
