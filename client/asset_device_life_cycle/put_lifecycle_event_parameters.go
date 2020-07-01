// Code generated by go-swagger; DO NOT EDIT.

package asset_device_life_cycle

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

// NewPutLifecycleEventParams creates a new PutLifecycleEventParams object
// with the default values initialized.
func NewPutLifecycleEventParams() *PutLifecycleEventParams {
	var ()
	return &PutLifecycleEventParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLifecycleEventParamsWithTimeout creates a new PutLifecycleEventParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLifecycleEventParamsWithTimeout(timeout time.Duration) *PutLifecycleEventParams {
	var ()
	return &PutLifecycleEventParams{

		timeout: timeout,
	}
}

// NewPutLifecycleEventParamsWithContext creates a new PutLifecycleEventParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLifecycleEventParamsWithContext(ctx context.Context) *PutLifecycleEventParams {
	var ()
	return &PutLifecycleEventParams{

		Context: ctx,
	}
}

// NewPutLifecycleEventParamsWithHTTPClient creates a new PutLifecycleEventParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLifecycleEventParamsWithHTTPClient(client *http.Client) *PutLifecycleEventParams {
	var ()
	return &PutLifecycleEventParams{
		HTTPClient: client,
	}
}

/*PutLifecycleEventParams contains all the parameters to send to the API endpoint
for the put lifecycle event operation typically these are written to a http.Request
*/
type PutLifecycleEventParams struct {

	/*AssetID
	  ID of the asset that the event is for

	*/
	AssetID *string
	/*AssetNo
	  Asset number of the device that the event is for

	*/
	AssetNo *string
	/*Date
	  in YYYY-MM-DD or YYYY-MM-DD-HH:MM format.

	*/
	Date string
	/*Device
	  Name of the device that the event is for

	*/
	Device *string
	/*DeviceID
	  ID of the device that the event is for

	*/
	DeviceID *string
	/*Notes
	  Any additional notes

	*/
	Notes *string
	/*SerialNo
	  Serial number of the device that the event is for

	*/
	SerialNo *string
	/*Type
	  must be defined already in device42.

	*/
	Type string
	/*User
	  enduser name

	*/
	User *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put lifecycle event params
func (o *PutLifecycleEventParams) WithTimeout(timeout time.Duration) *PutLifecycleEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put lifecycle event params
func (o *PutLifecycleEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put lifecycle event params
func (o *PutLifecycleEventParams) WithContext(ctx context.Context) *PutLifecycleEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put lifecycle event params
func (o *PutLifecycleEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put lifecycle event params
func (o *PutLifecycleEventParams) WithHTTPClient(client *http.Client) *PutLifecycleEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put lifecycle event params
func (o *PutLifecycleEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssetID adds the assetID to the put lifecycle event params
func (o *PutLifecycleEventParams) WithAssetID(assetID *string) *PutLifecycleEventParams {
	o.SetAssetID(assetID)
	return o
}

// SetAssetID adds the assetId to the put lifecycle event params
func (o *PutLifecycleEventParams) SetAssetID(assetID *string) {
	o.AssetID = assetID
}

// WithAssetNo adds the assetNo to the put lifecycle event params
func (o *PutLifecycleEventParams) WithAssetNo(assetNo *string) *PutLifecycleEventParams {
	o.SetAssetNo(assetNo)
	return o
}

// SetAssetNo adds the assetNo to the put lifecycle event params
func (o *PutLifecycleEventParams) SetAssetNo(assetNo *string) {
	o.AssetNo = assetNo
}

// WithDate adds the date to the put lifecycle event params
func (o *PutLifecycleEventParams) WithDate(date string) *PutLifecycleEventParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the put lifecycle event params
func (o *PutLifecycleEventParams) SetDate(date string) {
	o.Date = date
}

// WithDevice adds the device to the put lifecycle event params
func (o *PutLifecycleEventParams) WithDevice(device *string) *PutLifecycleEventParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the put lifecycle event params
func (o *PutLifecycleEventParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the put lifecycle event params
func (o *PutLifecycleEventParams) WithDeviceID(deviceID *string) *PutLifecycleEventParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the put lifecycle event params
func (o *PutLifecycleEventParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithNotes adds the notes to the put lifecycle event params
func (o *PutLifecycleEventParams) WithNotes(notes *string) *PutLifecycleEventParams {
	o.SetNotes(notes)
	return o
}

// SetNotes adds the notes to the put lifecycle event params
func (o *PutLifecycleEventParams) SetNotes(notes *string) {
	o.Notes = notes
}

// WithSerialNo adds the serialNo to the put lifecycle event params
func (o *PutLifecycleEventParams) WithSerialNo(serialNo *string) *PutLifecycleEventParams {
	o.SetSerialNo(serialNo)
	return o
}

// SetSerialNo adds the serialNo to the put lifecycle event params
func (o *PutLifecycleEventParams) SetSerialNo(serialNo *string) {
	o.SerialNo = serialNo
}

// WithType adds the typeVar to the put lifecycle event params
func (o *PutLifecycleEventParams) WithType(typeVar string) *PutLifecycleEventParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the put lifecycle event params
func (o *PutLifecycleEventParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WithUser adds the user to the put lifecycle event params
func (o *PutLifecycleEventParams) WithUser(user *string) *PutLifecycleEventParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the put lifecycle event params
func (o *PutLifecycleEventParams) SetUser(user *string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *PutLifecycleEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AssetID != nil {

		// form param asset_id
		var frAssetID string
		if o.AssetID != nil {
			frAssetID = *o.AssetID
		}
		fAssetID := frAssetID
		if fAssetID != "" {
			if err := r.SetFormParam("asset_id", fAssetID); err != nil {
				return err
			}
		}

	}

	if o.AssetNo != nil {

		// form param asset_no
		var frAssetNo string
		if o.AssetNo != nil {
			frAssetNo = *o.AssetNo
		}
		fAssetNo := frAssetNo
		if fAssetNo != "" {
			if err := r.SetFormParam("asset_no", fAssetNo); err != nil {
				return err
			}
		}

	}

	// form param date
	frDate := o.Date
	fDate := frDate
	if fDate != "" {
		if err := r.SetFormParam("date", fDate); err != nil {
			return err
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

	if o.SerialNo != nil {

		// form param serial_no
		var frSerialNo string
		if o.SerialNo != nil {
			frSerialNo = *o.SerialNo
		}
		fSerialNo := frSerialNo
		if fSerialNo != "" {
			if err := r.SetFormParam("serial_no", fSerialNo); err != nil {
				return err
			}
		}

	}

	// form param type
	frType := o.Type
	fType := frType
	if fType != "" {
		if err := r.SetFormParam("type", fType); err != nil {
			return err
		}
	}

	if o.User != nil {

		// form param user
		var frUser string
		if o.User != nil {
			frUser = *o.User
		}
		fUser := frUser
		if fUser != "" {
			if err := r.SetFormParam("user", fUser); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
