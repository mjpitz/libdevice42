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

// NewGetServiceDetailsParams creates a new GetServiceDetailsParams object
// with the default values initialized.
func NewGetServiceDetailsParams() *GetServiceDetailsParams {
	var ()
	return &GetServiceDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceDetailsParamsWithTimeout creates a new GetServiceDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServiceDetailsParamsWithTimeout(timeout time.Duration) *GetServiceDetailsParams {
	var ()
	return &GetServiceDetailsParams{

		timeout: timeout,
	}
}

// NewGetServiceDetailsParamsWithContext creates a new GetServiceDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServiceDetailsParamsWithContext(ctx context.Context) *GetServiceDetailsParams {
	var ()
	return &GetServiceDetailsParams{

		Context: ctx,
	}
}

// NewGetServiceDetailsParamsWithHTTPClient creates a new GetServiceDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServiceDetailsParamsWithHTTPClient(client *http.Client) *GetServiceDetailsParams {
	var ()
	return &GetServiceDetailsParams{
		HTTPClient: client,
	}
}

/*GetServiceDetailsParams contains all the parameters to send to the API endpoint
for the get service details operation typically these are written to a http.Request
*/
type GetServiceDetailsParams struct {

	/*Device
	  Device name

	*/
	Device *string
	/*DeviceID
	  Device ID

	*/
	DeviceID string
	/*ServiceDetailID
	  filter by id of the service in use

	*/
	ServiceDetailID *string
	/*ServiceID
	  filter by id of the service

	*/
	ServiceID *string
	/*UserID
	  filter by id of the user

	*/
	UserID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service details params
func (o *GetServiceDetailsParams) WithTimeout(timeout time.Duration) *GetServiceDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service details params
func (o *GetServiceDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service details params
func (o *GetServiceDetailsParams) WithContext(ctx context.Context) *GetServiceDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service details params
func (o *GetServiceDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service details params
func (o *GetServiceDetailsParams) WithHTTPClient(client *http.Client) *GetServiceDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service details params
func (o *GetServiceDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevice adds the device to the get service details params
func (o *GetServiceDetailsParams) WithDevice(device *string) *GetServiceDetailsParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the get service details params
func (o *GetServiceDetailsParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the get service details params
func (o *GetServiceDetailsParams) WithDeviceID(deviceID string) *GetServiceDetailsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get service details params
func (o *GetServiceDetailsParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithServiceDetailID adds the serviceDetailID to the get service details params
func (o *GetServiceDetailsParams) WithServiceDetailID(serviceDetailID *string) *GetServiceDetailsParams {
	o.SetServiceDetailID(serviceDetailID)
	return o
}

// SetServiceDetailID adds the serviceDetailId to the get service details params
func (o *GetServiceDetailsParams) SetServiceDetailID(serviceDetailID *string) {
	o.ServiceDetailID = serviceDetailID
}

// WithServiceID adds the serviceID to the get service details params
func (o *GetServiceDetailsParams) WithServiceID(serviceID *string) *GetServiceDetailsParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the get service details params
func (o *GetServiceDetailsParams) SetServiceID(serviceID *string) {
	o.ServiceID = serviceID
}

// WithUserID adds the userID to the get service details params
func (o *GetServiceDetailsParams) WithUserID(userID *string) *GetServiceDetailsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get service details params
func (o *GetServiceDetailsParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param device_id
	qrDeviceID := o.DeviceID
	qDeviceID := qrDeviceID
	if qDeviceID != "" {
		if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
			return err
		}
	}

	if o.ServiceDetailID != nil {

		// query param service_detail_id
		var qrServiceDetailID string
		if o.ServiceDetailID != nil {
			qrServiceDetailID = *o.ServiceDetailID
		}
		qServiceDetailID := qrServiceDetailID
		if qServiceDetailID != "" {
			if err := r.SetQueryParam("service_detail_id", qServiceDetailID); err != nil {
				return err
			}
		}

	}

	if o.ServiceID != nil {

		// query param service_id
		var qrServiceID string
		if o.ServiceID != nil {
			qrServiceID = *o.ServiceID
		}
		qServiceID := qrServiceID
		if qServiceID != "" {
			if err := r.SetQueryParam("service_id", qServiceID); err != nil {
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
