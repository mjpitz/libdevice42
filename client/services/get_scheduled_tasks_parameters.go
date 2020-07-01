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

// NewGetScheduledTasksParams creates a new GetScheduledTasksParams object
// with the default values initialized.
func NewGetScheduledTasksParams() *GetScheduledTasksParams {
	var ()
	return &GetScheduledTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScheduledTasksParamsWithTimeout creates a new GetScheduledTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScheduledTasksParamsWithTimeout(timeout time.Duration) *GetScheduledTasksParams {
	var ()
	return &GetScheduledTasksParams{

		timeout: timeout,
	}
}

// NewGetScheduledTasksParamsWithContext creates a new GetScheduledTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScheduledTasksParamsWithContext(ctx context.Context) *GetScheduledTasksParams {
	var ()
	return &GetScheduledTasksParams{

		Context: ctx,
	}
}

// NewGetScheduledTasksParamsWithHTTPClient creates a new GetScheduledTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScheduledTasksParamsWithHTTPClient(client *http.Client) *GetScheduledTasksParams {
	var ()
	return &GetScheduledTasksParams{
		HTTPClient: client,
	}
}

/*GetScheduledTasksParams contains all the parameters to send to the API endpoint
for the get scheduled tasks operation typically these are written to a http.Request
*/
type GetScheduledTasksParams struct {

	/*Device
	  device

	*/
	Device *string
	/*DeviceID
	  device id

	*/
	DeviceID *int64
	/*ID
	  filter results by Service schedule ID

	*/
	ID *int64
	/*UserID
	  user id

	*/
	UserID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scheduled tasks params
func (o *GetScheduledTasksParams) WithTimeout(timeout time.Duration) *GetScheduledTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scheduled tasks params
func (o *GetScheduledTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scheduled tasks params
func (o *GetScheduledTasksParams) WithContext(ctx context.Context) *GetScheduledTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scheduled tasks params
func (o *GetScheduledTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scheduled tasks params
func (o *GetScheduledTasksParams) WithHTTPClient(client *http.Client) *GetScheduledTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scheduled tasks params
func (o *GetScheduledTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevice adds the device to the get scheduled tasks params
func (o *GetScheduledTasksParams) WithDevice(device *string) *GetScheduledTasksParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the get scheduled tasks params
func (o *GetScheduledTasksParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the get scheduled tasks params
func (o *GetScheduledTasksParams) WithDeviceID(deviceID *int64) *GetScheduledTasksParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get scheduled tasks params
func (o *GetScheduledTasksParams) SetDeviceID(deviceID *int64) {
	o.DeviceID = deviceID
}

// WithID adds the id to the get scheduled tasks params
func (o *GetScheduledTasksParams) WithID(id *int64) *GetScheduledTasksParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get scheduled tasks params
func (o *GetScheduledTasksParams) SetID(id *int64) {
	o.ID = id
}

// WithUserID adds the userID to the get scheduled tasks params
func (o *GetScheduledTasksParams) WithUserID(userID *string) *GetScheduledTasksParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get scheduled tasks params
func (o *GetScheduledTasksParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScheduledTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID int64
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := swag.FormatInt64(qrDeviceID)
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID int64
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt64(qrID)
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID string
		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {
			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
