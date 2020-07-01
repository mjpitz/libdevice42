// Code generated by go-swagger; DO NOT EDIT.

package ip_a_m

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

// NewGetIPAMMacsParams creates a new GetIPAMMacsParams object
// with the default values initialized.
func NewGetIPAMMacsParams() *GetIPAMMacsParams {
	var ()
	return &GetIPAMMacsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPAMMacsParamsWithTimeout creates a new GetIPAMMacsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPAMMacsParamsWithTimeout(timeout time.Duration) *GetIPAMMacsParams {
	var ()
	return &GetIPAMMacsParams{

		timeout: timeout,
	}
}

// NewGetIPAMMacsParamsWithContext creates a new GetIPAMMacsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPAMMacsParamsWithContext(ctx context.Context) *GetIPAMMacsParams {
	var ()
	return &GetIPAMMacsParams{

		Context: ctx,
	}
}

// NewGetIPAMMacsParamsWithHTTPClient creates a new GetIPAMMacsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPAMMacsParamsWithHTTPClient(client *http.Client) *GetIPAMMacsParams {
	var ()
	return &GetIPAMMacsParams{
		HTTPClient: client,
	}
}

/*GetIPAMMacsParams contains all the parameters to send to the API endpoint
for the get IP a m macs operation typically these are written to a http.Request
*/
type GetIPAMMacsParams struct {

	/*Device
	  Device name

	*/
	Device *string
	/*DeviceID
	  Device ID

	*/
	DeviceID *string
	/*LastUpdatedGt
	  last updated greater than date YYYY-MM-DD format

	*/
	LastUpdatedGt *string
	/*LastUpdatedLt
	  last updated less than date YYYY-MM-DD format

	*/
	LastUpdatedLt *string
	/*Mac
	  mac address

	*/
	Mac *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get IP a m macs params
func (o *GetIPAMMacsParams) WithTimeout(timeout time.Duration) *GetIPAMMacsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get IP a m macs params
func (o *GetIPAMMacsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get IP a m macs params
func (o *GetIPAMMacsParams) WithContext(ctx context.Context) *GetIPAMMacsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get IP a m macs params
func (o *GetIPAMMacsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get IP a m macs params
func (o *GetIPAMMacsParams) WithHTTPClient(client *http.Client) *GetIPAMMacsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get IP a m macs params
func (o *GetIPAMMacsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevice adds the device to the get IP a m macs params
func (o *GetIPAMMacsParams) WithDevice(device *string) *GetIPAMMacsParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the get IP a m macs params
func (o *GetIPAMMacsParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the get IP a m macs params
func (o *GetIPAMMacsParams) WithDeviceID(deviceID *string) *GetIPAMMacsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get IP a m macs params
func (o *GetIPAMMacsParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithLastUpdatedGt adds the lastUpdatedGt to the get IP a m macs params
func (o *GetIPAMMacsParams) WithLastUpdatedGt(lastUpdatedGt *string) *GetIPAMMacsParams {
	o.SetLastUpdatedGt(lastUpdatedGt)
	return o
}

// SetLastUpdatedGt adds the lastUpdatedGt to the get IP a m macs params
func (o *GetIPAMMacsParams) SetLastUpdatedGt(lastUpdatedGt *string) {
	o.LastUpdatedGt = lastUpdatedGt
}

// WithLastUpdatedLt adds the lastUpdatedLt to the get IP a m macs params
func (o *GetIPAMMacsParams) WithLastUpdatedLt(lastUpdatedLt *string) *GetIPAMMacsParams {
	o.SetLastUpdatedLt(lastUpdatedLt)
	return o
}

// SetLastUpdatedLt adds the lastUpdatedLt to the get IP a m macs params
func (o *GetIPAMMacsParams) SetLastUpdatedLt(lastUpdatedLt *string) {
	o.LastUpdatedLt = lastUpdatedLt
}

// WithMac adds the mac to the get IP a m macs params
func (o *GetIPAMMacsParams) WithMac(mac *string) *GetIPAMMacsParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the get IP a m macs params
func (o *GetIPAMMacsParams) SetMac(mac *string) {
	o.Mac = mac
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPAMMacsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
		var qrDeviceID string
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}

	}

	if o.LastUpdatedGt != nil {

		// query param last_updated_gt
		var qrLastUpdatedGt string
		if o.LastUpdatedGt != nil {
			qrLastUpdatedGt = *o.LastUpdatedGt
		}
		qLastUpdatedGt := qrLastUpdatedGt
		if qLastUpdatedGt != "" {
			if err := r.SetQueryParam("last_updated_gt", qLastUpdatedGt); err != nil {
				return err
			}
		}

	}

	if o.LastUpdatedLt != nil {

		// query param last_updated_lt
		var qrLastUpdatedLt string
		if o.LastUpdatedLt != nil {
			qrLastUpdatedLt = *o.LastUpdatedLt
		}
		qLastUpdatedLt := qrLastUpdatedLt
		if qLastUpdatedLt != "" {
			if err := r.SetQueryParam("last_updated_lt", qLastUpdatedLt); err != nil {
				return err
			}
		}

	}

	if o.Mac != nil {

		// query param mac
		var qrMac string
		if o.Mac != nil {
			qrMac = *o.Mac
		}
		qMac := qrMac
		if qMac != "" {
			if err := r.SetQueryParam("mac", qMac); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
