// Code generated by go-swagger; DO NOT EDIT.

package p_d_u

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

// NewGetPdusParams creates a new GetPdusParams object
// with the default values initialized.
func NewGetPdusParams() *GetPdusParams {
	var ()
	return &GetPdusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPdusParamsWithTimeout creates a new GetPdusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPdusParamsWithTimeout(timeout time.Duration) *GetPdusParams {
	var ()
	return &GetPdusParams{

		timeout: timeout,
	}
}

// NewGetPdusParamsWithContext creates a new GetPdusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPdusParamsWithContext(ctx context.Context) *GetPdusParams {
	var ()
	return &GetPdusParams{

		Context: ctx,
	}
}

// NewGetPdusParamsWithHTTPClient creates a new GetPdusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPdusParamsWithHTTPClient(client *http.Client) *GetPdusParams {
	var ()
	return &GetPdusParams{
		HTTPClient: client,
	}
}

/*GetPdusParams contains all the parameters to send to the API endpoint
for the get pdus operation typically these are written to a http.Request
*/
type GetPdusParams struct {

	/*BuildingID
	  filter by building id

	*/
	BuildingID *int64
	/*DeviceID
	  filter by device id

	*/
	DeviceID *int64
	/*Name
	  filter by name

	*/
	Name *string
	/*PduModel
	  filter by PDU model name

	*/
	PduModel *string
	/*PduModelID
	  filter by PDU model ID

	*/
	PduModelID *int64
	/*RackID
	  filter by rack id

	*/
	RackID *int64
	/*RoomID
	  filter by room id

	*/
	RoomID *int64
	/*SerialNo
	  filter by PDU serial_no

	*/
	SerialNo *string
	/*Type
	  filter by type

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pdus params
func (o *GetPdusParams) WithTimeout(timeout time.Duration) *GetPdusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pdus params
func (o *GetPdusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pdus params
func (o *GetPdusParams) WithContext(ctx context.Context) *GetPdusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pdus params
func (o *GetPdusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pdus params
func (o *GetPdusParams) WithHTTPClient(client *http.Client) *GetPdusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pdus params
func (o *GetPdusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildingID adds the buildingID to the get pdus params
func (o *GetPdusParams) WithBuildingID(buildingID *int64) *GetPdusParams {
	o.SetBuildingID(buildingID)
	return o
}

// SetBuildingID adds the buildingId to the get pdus params
func (o *GetPdusParams) SetBuildingID(buildingID *int64) {
	o.BuildingID = buildingID
}

// WithDeviceID adds the deviceID to the get pdus params
func (o *GetPdusParams) WithDeviceID(deviceID *int64) *GetPdusParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get pdus params
func (o *GetPdusParams) SetDeviceID(deviceID *int64) {
	o.DeviceID = deviceID
}

// WithName adds the name to the get pdus params
func (o *GetPdusParams) WithName(name *string) *GetPdusParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get pdus params
func (o *GetPdusParams) SetName(name *string) {
	o.Name = name
}

// WithPduModel adds the pduModel to the get pdus params
func (o *GetPdusParams) WithPduModel(pduModel *string) *GetPdusParams {
	o.SetPduModel(pduModel)
	return o
}

// SetPduModel adds the pduModel to the get pdus params
func (o *GetPdusParams) SetPduModel(pduModel *string) {
	o.PduModel = pduModel
}

// WithPduModelID adds the pduModelID to the get pdus params
func (o *GetPdusParams) WithPduModelID(pduModelID *int64) *GetPdusParams {
	o.SetPduModelID(pduModelID)
	return o
}

// SetPduModelID adds the pduModelId to the get pdus params
func (o *GetPdusParams) SetPduModelID(pduModelID *int64) {
	o.PduModelID = pduModelID
}

// WithRackID adds the rackID to the get pdus params
func (o *GetPdusParams) WithRackID(rackID *int64) *GetPdusParams {
	o.SetRackID(rackID)
	return o
}

// SetRackID adds the rackId to the get pdus params
func (o *GetPdusParams) SetRackID(rackID *int64) {
	o.RackID = rackID
}

// WithRoomID adds the roomID to the get pdus params
func (o *GetPdusParams) WithRoomID(roomID *int64) *GetPdusParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the get pdus params
func (o *GetPdusParams) SetRoomID(roomID *int64) {
	o.RoomID = roomID
}

// WithSerialNo adds the serialNo to the get pdus params
func (o *GetPdusParams) WithSerialNo(serialNo *string) *GetPdusParams {
	o.SetSerialNo(serialNo)
	return o
}

// SetSerialNo adds the serialNo to the get pdus params
func (o *GetPdusParams) SetSerialNo(serialNo *string) {
	o.SerialNo = serialNo
}

// WithType adds the typeVar to the get pdus params
func (o *GetPdusParams) WithType(typeVar *string) *GetPdusParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get pdus params
func (o *GetPdusParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetPdusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BuildingID != nil {

		// query param building_id
		var qrBuildingID int64
		if o.BuildingID != nil {
			qrBuildingID = *o.BuildingID
		}
		qBuildingID := swag.FormatInt64(qrBuildingID)
		if qBuildingID != "" {
			if err := r.SetQueryParam("building_id", qBuildingID); err != nil {
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

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.PduModel != nil {

		// query param pdu_model
		var qrPduModel string
		if o.PduModel != nil {
			qrPduModel = *o.PduModel
		}
		qPduModel := qrPduModel
		if qPduModel != "" {
			if err := r.SetQueryParam("pdu_model", qPduModel); err != nil {
				return err
			}
		}

	}

	if o.PduModelID != nil {

		// query param pdu_model_id
		var qrPduModelID int64
		if o.PduModelID != nil {
			qrPduModelID = *o.PduModelID
		}
		qPduModelID := swag.FormatInt64(qrPduModelID)
		if qPduModelID != "" {
			if err := r.SetQueryParam("pdu_model_id", qPduModelID); err != nil {
				return err
			}
		}

	}

	if o.RackID != nil {

		// query param rack_id
		var qrRackID int64
		if o.RackID != nil {
			qrRackID = *o.RackID
		}
		qRackID := swag.FormatInt64(qrRackID)
		if qRackID != "" {
			if err := r.SetQueryParam("rack_id", qRackID); err != nil {
				return err
			}
		}

	}

	if o.RoomID != nil {

		// query param room_id
		var qrRoomID int64
		if o.RoomID != nil {
			qrRoomID = *o.RoomID
		}
		qRoomID := swag.FormatInt64(qrRoomID)
		if qRoomID != "" {
			if err := r.SetQueryParam("room_id", qRoomID); err != nil {
				return err
			}
		}

	}

	if o.SerialNo != nil {

		// query param serial_no
		var qrSerialNo string
		if o.SerialNo != nil {
			qrSerialNo = *o.SerialNo
		}
		qSerialNo := qrSerialNo
		if qSerialNo != "" {
			if err := r.SetQueryParam("serial_no", qSerialNo); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
