// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// NewPutCustomFieldsParams creates a new PutCustomFieldsParams object
// with the default values initialized.
func NewPutCustomFieldsParams() *PutCustomFieldsParams {
	var ()
	return &PutCustomFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomFieldsParamsWithTimeout creates a new PutCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomFieldsParamsWithTimeout(timeout time.Duration) *PutCustomFieldsParams {
	var ()
	return &PutCustomFieldsParams{

		timeout: timeout,
	}
}

// NewPutCustomFieldsParamsWithContext creates a new PutCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomFieldsParamsWithContext(ctx context.Context) *PutCustomFieldsParams {
	var ()
	return &PutCustomFieldsParams{

		Context: ctx,
	}
}

// NewPutCustomFieldsParamsWithHTTPClient creates a new PutCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomFieldsParamsWithHTTPClient(client *http.Client) *PutCustomFieldsParams {
	var ()
	return &PutCustomFieldsParams{
		HTTPClient: client,
	}
}

/*PutCustomFieldsParams contains all the parameters to send to the API endpoint
for the put custom fields operation typically these are written to a http.Request
*/
type PutCustomFieldsParams struct {

	/*BulkFields
	  comma separated key value pairs, with key and value separated by colon. e.g.key1:value1, key2:value2

	*/
	BulkFields *string
	/*ClearValue
	  yes to clear existing value for that field

	*/
	ClearValue *string
	/*Key
	  Can be new or existing. This is the custom field name.

	*/
	Key string
	/*Name
	  Name of the customer

	*/
	Name string
	/*Notes
	  Any additional notes

	*/
	Notes *string
	/*RelatedFieldName
	  Required if type = related field.

	*/
	RelatedFieldName *string
	/*Type
	  this is the custom field type. If left blank, default is text. Date should be formatted as YYYY-MM-DD

	*/
	Type *string
	/*Value
	  This will set the value of the custom field for the specific object.

	*/
	Value *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put custom fields params
func (o *PutCustomFieldsParams) WithTimeout(timeout time.Duration) *PutCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put custom fields params
func (o *PutCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put custom fields params
func (o *PutCustomFieldsParams) WithContext(ctx context.Context) *PutCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put custom fields params
func (o *PutCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put custom fields params
func (o *PutCustomFieldsParams) WithHTTPClient(client *http.Client) *PutCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put custom fields params
func (o *PutCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBulkFields adds the bulkFields to the put custom fields params
func (o *PutCustomFieldsParams) WithBulkFields(bulkFields *string) *PutCustomFieldsParams {
	o.SetBulkFields(bulkFields)
	return o
}

// SetBulkFields adds the bulkFields to the put custom fields params
func (o *PutCustomFieldsParams) SetBulkFields(bulkFields *string) {
	o.BulkFields = bulkFields
}

// WithClearValue adds the clearValue to the put custom fields params
func (o *PutCustomFieldsParams) WithClearValue(clearValue *string) *PutCustomFieldsParams {
	o.SetClearValue(clearValue)
	return o
}

// SetClearValue adds the clearValue to the put custom fields params
func (o *PutCustomFieldsParams) SetClearValue(clearValue *string) {
	o.ClearValue = clearValue
}

// WithKey adds the key to the put custom fields params
func (o *PutCustomFieldsParams) WithKey(key string) *PutCustomFieldsParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the put custom fields params
func (o *PutCustomFieldsParams) SetKey(key string) {
	o.Key = key
}

// WithName adds the name to the put custom fields params
func (o *PutCustomFieldsParams) WithName(name string) *PutCustomFieldsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put custom fields params
func (o *PutCustomFieldsParams) SetName(name string) {
	o.Name = name
}

// WithNotes adds the notes to the put custom fields params
func (o *PutCustomFieldsParams) WithNotes(notes *string) *PutCustomFieldsParams {
	o.SetNotes(notes)
	return o
}

// SetNotes adds the notes to the put custom fields params
func (o *PutCustomFieldsParams) SetNotes(notes *string) {
	o.Notes = notes
}

// WithRelatedFieldName adds the relatedFieldName to the put custom fields params
func (o *PutCustomFieldsParams) WithRelatedFieldName(relatedFieldName *string) *PutCustomFieldsParams {
	o.SetRelatedFieldName(relatedFieldName)
	return o
}

// SetRelatedFieldName adds the relatedFieldName to the put custom fields params
func (o *PutCustomFieldsParams) SetRelatedFieldName(relatedFieldName *string) {
	o.RelatedFieldName = relatedFieldName
}

// WithType adds the typeVar to the put custom fields params
func (o *PutCustomFieldsParams) WithType(typeVar *string) *PutCustomFieldsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the put custom fields params
func (o *PutCustomFieldsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithValue adds the value to the put custom fields params
func (o *PutCustomFieldsParams) WithValue(value *string) *PutCustomFieldsParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the put custom fields params
func (o *PutCustomFieldsParams) SetValue(value *string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BulkFields != nil {

		// form param bulk_fields
		var frBulkFields string
		if o.BulkFields != nil {
			frBulkFields = *o.BulkFields
		}
		fBulkFields := frBulkFields
		if fBulkFields != "" {
			if err := r.SetFormParam("bulk_fields", fBulkFields); err != nil {
				return err
			}
		}

	}

	if o.ClearValue != nil {

		// form param clear_value
		var frClearValue string
		if o.ClearValue != nil {
			frClearValue = *o.ClearValue
		}
		fClearValue := frClearValue
		if fClearValue != "" {
			if err := r.SetFormParam("clear_value", fClearValue); err != nil {
				return err
			}
		}

	}

	// form param key
	frKey := o.Key
	fKey := frKey
	if fKey != "" {
		if err := r.SetFormParam("key", fKey); err != nil {
			return err
		}
	}

	// form param name
	frName := o.Name
	fName := frName
	if fName != "" {
		if err := r.SetFormParam("name", fName); err != nil {
			return err
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

	if o.RelatedFieldName != nil {

		// form param related_field_name
		var frRelatedFieldName string
		if o.RelatedFieldName != nil {
			frRelatedFieldName = *o.RelatedFieldName
		}
		fRelatedFieldName := frRelatedFieldName
		if fRelatedFieldName != "" {
			if err := r.SetFormParam("related_field_name", fRelatedFieldName); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// form param type
		var frType string
		if o.Type != nil {
			frType = *o.Type
		}
		fType := frType
		if fType != "" {
			if err := r.SetFormParam("type", fType); err != nil {
				return err
			}
		}

	}

	if o.Value != nil {

		// form param value
		var frValue string
		if o.Value != nil {
			frValue = *o.Value
		}
		fValue := frValue
		if fValue != "" {
			if err := r.SetFormParam("value", fValue); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
