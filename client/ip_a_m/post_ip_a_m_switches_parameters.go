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

// NewPostIPAMSwitchesParams creates a new PostIPAMSwitchesParams object
// with the default values initialized.
func NewPostIPAMSwitchesParams() *PostIPAMSwitchesParams {
	var ()
	return &PostIPAMSwitchesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIPAMSwitchesParamsWithTimeout creates a new PostIPAMSwitchesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIPAMSwitchesParamsWithTimeout(timeout time.Duration) *PostIPAMSwitchesParams {
	var ()
	return &PostIPAMSwitchesParams{

		timeout: timeout,
	}
}

// NewPostIPAMSwitchesParamsWithContext creates a new PostIPAMSwitchesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIPAMSwitchesParamsWithContext(ctx context.Context) *PostIPAMSwitchesParams {
	var ()
	return &PostIPAMSwitchesParams{

		Context: ctx,
	}
}

// NewPostIPAMSwitchesParamsWithHTTPClient creates a new PostIPAMSwitchesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIPAMSwitchesParamsWithHTTPClient(client *http.Client) *PostIPAMSwitchesParams {
	var ()
	return &PostIPAMSwitchesParams{
		HTTPClient: client,
	}
}

/*PostIPAMSwitchesParams contains all the parameters to send to the API endpoint
for the post IP a m switches operation typically these are written to a http.Request
*/
type PostIPAMSwitchesParams struct {

	/*AssetIds
	  comma separated values of existing assets.

	*/
	AssetIds *string
	/*Assets
	  comma separated names of new assets.

	*/
	Assets *string
	/*Device
	  Name of new or existing device. Existing device must be a network switch. If stacked switches, must be of type 'cluster'

	*/
	Device *string
	/*DeviceID
	  ID of existing device. Existing device must be a network switch. IF stacked switches, must be of type 'cluster'

	*/
	DeviceID *string
	/*DeviceIds
	  comma separated values of existing devices.

	*/
	DeviceIds *string
	/*Devices
	  comma separated names of new devices.

	*/
	Devices *string
	/*SwitchTemplateID
	  GET all Switch Templates

	*/
	SwitchTemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post IP a m switches params
func (o *PostIPAMSwitchesParams) WithTimeout(timeout time.Duration) *PostIPAMSwitchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post IP a m switches params
func (o *PostIPAMSwitchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post IP a m switches params
func (o *PostIPAMSwitchesParams) WithContext(ctx context.Context) *PostIPAMSwitchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post IP a m switches params
func (o *PostIPAMSwitchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post IP a m switches params
func (o *PostIPAMSwitchesParams) WithHTTPClient(client *http.Client) *PostIPAMSwitchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post IP a m switches params
func (o *PostIPAMSwitchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssetIds adds the assetIds to the post IP a m switches params
func (o *PostIPAMSwitchesParams) WithAssetIds(assetIds *string) *PostIPAMSwitchesParams {
	o.SetAssetIds(assetIds)
	return o
}

// SetAssetIds adds the assetIds to the post IP a m switches params
func (o *PostIPAMSwitchesParams) SetAssetIds(assetIds *string) {
	o.AssetIds = assetIds
}

// WithAssets adds the assets to the post IP a m switches params
func (o *PostIPAMSwitchesParams) WithAssets(assets *string) *PostIPAMSwitchesParams {
	o.SetAssets(assets)
	return o
}

// SetAssets adds the assets to the post IP a m switches params
func (o *PostIPAMSwitchesParams) SetAssets(assets *string) {
	o.Assets = assets
}

// WithDevice adds the device to the post IP a m switches params
func (o *PostIPAMSwitchesParams) WithDevice(device *string) *PostIPAMSwitchesParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the post IP a m switches params
func (o *PostIPAMSwitchesParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the post IP a m switches params
func (o *PostIPAMSwitchesParams) WithDeviceID(deviceID *string) *PostIPAMSwitchesParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the post IP a m switches params
func (o *PostIPAMSwitchesParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithDeviceIds adds the deviceIds to the post IP a m switches params
func (o *PostIPAMSwitchesParams) WithDeviceIds(deviceIds *string) *PostIPAMSwitchesParams {
	o.SetDeviceIds(deviceIds)
	return o
}

// SetDeviceIds adds the deviceIds to the post IP a m switches params
func (o *PostIPAMSwitchesParams) SetDeviceIds(deviceIds *string) {
	o.DeviceIds = deviceIds
}

// WithDevices adds the devices to the post IP a m switches params
func (o *PostIPAMSwitchesParams) WithDevices(devices *string) *PostIPAMSwitchesParams {
	o.SetDevices(devices)
	return o
}

// SetDevices adds the devices to the post IP a m switches params
func (o *PostIPAMSwitchesParams) SetDevices(devices *string) {
	o.Devices = devices
}

// WithSwitchTemplateID adds the switchTemplateID to the post IP a m switches params
func (o *PostIPAMSwitchesParams) WithSwitchTemplateID(switchTemplateID string) *PostIPAMSwitchesParams {
	o.SetSwitchTemplateID(switchTemplateID)
	return o
}

// SetSwitchTemplateID adds the switchTemplateId to the post IP a m switches params
func (o *PostIPAMSwitchesParams) SetSwitchTemplateID(switchTemplateID string) {
	o.SwitchTemplateID = switchTemplateID
}

// WriteToRequest writes these params to a swagger request
func (o *PostIPAMSwitchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AssetIds != nil {

		// form param asset_ids
		var frAssetIds string
		if o.AssetIds != nil {
			frAssetIds = *o.AssetIds
		}
		fAssetIds := frAssetIds
		if fAssetIds != "" {
			if err := r.SetFormParam("asset_ids", fAssetIds); err != nil {
				return err
			}
		}

	}

	if o.Assets != nil {

		// form param assets
		var frAssets string
		if o.Assets != nil {
			frAssets = *o.Assets
		}
		fAssets := frAssets
		if fAssets != "" {
			if err := r.SetFormParam("assets", fAssets); err != nil {
				return err
			}
		}

	}

	if o.Device != nil {

		// form param device
		var frDevice string
		if o.Device != nil {
			frDevice = *o.Device
		}
		fDevice := frDevice
		if fDevice != "" {
			if err := r.SetFormParam("device", fDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// form param device_id
		var frDeviceID string
		if o.DeviceID != nil {
			frDeviceID = *o.DeviceID
		}
		fDeviceID := frDeviceID
		if fDeviceID != "" {
			if err := r.SetFormParam("device_id", fDeviceID); err != nil {
				return err
			}
		}

	}

	if o.DeviceIds != nil {

		// form param device_ids
		var frDeviceIds string
		if o.DeviceIds != nil {
			frDeviceIds = *o.DeviceIds
		}
		fDeviceIds := frDeviceIds
		if fDeviceIds != "" {
			if err := r.SetFormParam("device_ids", fDeviceIds); err != nil {
				return err
			}
		}

	}

	if o.Devices != nil {

		// form param devices
		var frDevices string
		if o.Devices != nil {
			frDevices = *o.Devices
		}
		fDevices := frDevices
		if fDevices != "" {
			if err := r.SetFormParam("devices", fDevices); err != nil {
				return err
			}
		}

	}

	// form param switch_template_id
	frSwitchTemplateID := o.SwitchTemplateID
	fSwitchTemplateID := frSwitchTemplateID
	if fSwitchTemplateID != "" {
		if err := r.SetFormParam("switch_template_id", fSwitchTemplateID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
