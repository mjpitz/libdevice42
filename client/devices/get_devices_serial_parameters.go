// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewGetDevicesSerialParams creates a new GetDevicesSerialParams object
// with the default values initialized.
func NewGetDevicesSerialParams() *GetDevicesSerialParams {
	var ()
	return &GetDevicesSerialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesSerialParamsWithTimeout creates a new GetDevicesSerialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDevicesSerialParamsWithTimeout(timeout time.Duration) *GetDevicesSerialParams {
	var ()
	return &GetDevicesSerialParams{

		timeout: timeout,
	}
}

// NewGetDevicesSerialParamsWithContext creates a new GetDevicesSerialParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDevicesSerialParamsWithContext(ctx context.Context) *GetDevicesSerialParams {
	var ()
	return &GetDevicesSerialParams{

		Context: ctx,
	}
}

// NewGetDevicesSerialParamsWithHTTPClient creates a new GetDevicesSerialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDevicesSerialParamsWithHTTPClient(client *http.Client) *GetDevicesSerialParams {
	var ()
	return &GetDevicesSerialParams{
		HTTPClient: client,
	}
}

/*GetDevicesSerialParams contains all the parameters to send to the API endpoint
for the get devices serial operation typically these are written to a http.Request
*/
type GetDevicesSerialParams struct {

	/*DeviceSerial*/
	DeviceSerial strfmt.UUID
	/*IncludeCols
	  do not return all columns just the ones specified. For example, ?include_cols=name, device_id, rack will only result in name, device_id, and rack included in the output. The following column names can be part of include_cols: name, device_id, rack, name, device_id, serial_no, asset_no, uuid, notes, in_service, service_level, type, id, last_updated, tags, customer_id, customer, hw_model, hw_size, manufacturer, hw_depth, rack, start_at, rack_id, orientation, row, room, building, blade_host_name, blade_host_id, slot_number, virtual_host_name, location, device_sub_type, os, osarch, osver, osverno, custom_fields, device_purchase_line_items, device_external_links, ip_addresses, mac_addresses, cpucount, cpucore, cpuspeed, ram, hddcount, hddsize, hddraid, hddraid_type, hdd_details, pdu_mapping_url,modules, vms, devices, aliases, xpos, ucs_manager

	*/
	IncludeCols *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get devices serial params
func (o *GetDevicesSerialParams) WithTimeout(timeout time.Duration) *GetDevicesSerialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices serial params
func (o *GetDevicesSerialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices serial params
func (o *GetDevicesSerialParams) WithContext(ctx context.Context) *GetDevicesSerialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices serial params
func (o *GetDevicesSerialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices serial params
func (o *GetDevicesSerialParams) WithHTTPClient(client *http.Client) *GetDevicesSerialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices serial params
func (o *GetDevicesSerialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceSerial adds the deviceSerial to the get devices serial params
func (o *GetDevicesSerialParams) WithDeviceSerial(deviceSerial strfmt.UUID) *GetDevicesSerialParams {
	o.SetDeviceSerial(deviceSerial)
	return o
}

// SetDeviceSerial adds the deviceSerial to the get devices serial params
func (o *GetDevicesSerialParams) SetDeviceSerial(deviceSerial strfmt.UUID) {
	o.DeviceSerial = deviceSerial
}

// WithIncludeCols adds the includeCols to the get devices serial params
func (o *GetDevicesSerialParams) WithIncludeCols(includeCols *string) *GetDevicesSerialParams {
	o.SetIncludeCols(includeCols)
	return o
}

// SetIncludeCols adds the includeCols to the get devices serial params
func (o *GetDevicesSerialParams) SetIncludeCols(includeCols *string) {
	o.IncludeCols = includeCols
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesSerialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param device-serial
	if err := r.SetPathParam("device-serial", o.DeviceSerial.String()); err != nil {
		return err
	}

	if o.IncludeCols != nil {

		// query param include_cols
		var qrIncludeCols string
		if o.IncludeCols != nil {
			qrIncludeCols = *o.IncludeCols
		}
		qIncludeCols := qrIncludeCols
		if qIncludeCols != "" {
			if err := r.SetQueryParam("include_cols", qIncludeCols); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
