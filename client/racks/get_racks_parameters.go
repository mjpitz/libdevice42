// Code generated by go-swagger; DO NOT EDIT.

package racks

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

// NewGetRacksParams creates a new GetRacksParams object
// with the default values initialized.
func NewGetRacksParams() *GetRacksParams {
	var ()
	return &GetRacksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRacksParamsWithTimeout creates a new GetRacksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRacksParamsWithTimeout(timeout time.Duration) *GetRacksParams {
	var ()
	return &GetRacksParams{

		timeout: timeout,
	}
}

// NewGetRacksParamsWithContext creates a new GetRacksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRacksParamsWithContext(ctx context.Context) *GetRacksParams {
	var ()
	return &GetRacksParams{

		Context: ctx,
	}
}

// NewGetRacksParamsWithHTTPClient creates a new GetRacksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRacksParamsWithHTTPClient(client *http.Client) *GetRacksParams {
	var ()
	return &GetRacksParams{
		HTTPClient: client,
	}
}

/*GetRacksParams contains all the parameters to send to the API endpoint
for the get racks operation typically these are written to a http.Request
*/
type GetRacksParams struct {

	/*AssetNo
	  filter by asset number

	*/
	AssetNo *string
	/*Building
	  filter by building name

	*/
	Building *string
	/*BuildingID
	  filter by building ID (Added in v5.9.0)

	*/
	BuildingID *string
	/*Manufacturer
	  filter by manufacturer

	*/
	Manufacturer *string
	/*Name
	  filter by name (Added in v6.0.0)

	*/
	Name *string
	/*Room
	  filter by room name. Only works if room ID is not present (Added in v5.9.0)

	*/
	Room *string
	/*RoomID
	  filter by room ID (Added in v5.9.0)

	*/
	RoomID *string
	/*Row
	  filter by row name

	*/
	Row *string
	/*Size
	  filter by rack size in U

	*/
	Size *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get racks params
func (o *GetRacksParams) WithTimeout(timeout time.Duration) *GetRacksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get racks params
func (o *GetRacksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get racks params
func (o *GetRacksParams) WithContext(ctx context.Context) *GetRacksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get racks params
func (o *GetRacksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get racks params
func (o *GetRacksParams) WithHTTPClient(client *http.Client) *GetRacksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get racks params
func (o *GetRacksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssetNo adds the assetNo to the get racks params
func (o *GetRacksParams) WithAssetNo(assetNo *string) *GetRacksParams {
	o.SetAssetNo(assetNo)
	return o
}

// SetAssetNo adds the assetNo to the get racks params
func (o *GetRacksParams) SetAssetNo(assetNo *string) {
	o.AssetNo = assetNo
}

// WithBuilding adds the building to the get racks params
func (o *GetRacksParams) WithBuilding(building *string) *GetRacksParams {
	o.SetBuilding(building)
	return o
}

// SetBuilding adds the building to the get racks params
func (o *GetRacksParams) SetBuilding(building *string) {
	o.Building = building
}

// WithBuildingID adds the buildingID to the get racks params
func (o *GetRacksParams) WithBuildingID(buildingID *string) *GetRacksParams {
	o.SetBuildingID(buildingID)
	return o
}

// SetBuildingID adds the buildingId to the get racks params
func (o *GetRacksParams) SetBuildingID(buildingID *string) {
	o.BuildingID = buildingID
}

// WithManufacturer adds the manufacturer to the get racks params
func (o *GetRacksParams) WithManufacturer(manufacturer *string) *GetRacksParams {
	o.SetManufacturer(manufacturer)
	return o
}

// SetManufacturer adds the manufacturer to the get racks params
func (o *GetRacksParams) SetManufacturer(manufacturer *string) {
	o.Manufacturer = manufacturer
}

// WithName adds the name to the get racks params
func (o *GetRacksParams) WithName(name *string) *GetRacksParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get racks params
func (o *GetRacksParams) SetName(name *string) {
	o.Name = name
}

// WithRoom adds the room to the get racks params
func (o *GetRacksParams) WithRoom(room *string) *GetRacksParams {
	o.SetRoom(room)
	return o
}

// SetRoom adds the room to the get racks params
func (o *GetRacksParams) SetRoom(room *string) {
	o.Room = room
}

// WithRoomID adds the roomID to the get racks params
func (o *GetRacksParams) WithRoomID(roomID *string) *GetRacksParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the get racks params
func (o *GetRacksParams) SetRoomID(roomID *string) {
	o.RoomID = roomID
}

// WithRow adds the row to the get racks params
func (o *GetRacksParams) WithRow(row *string) *GetRacksParams {
	o.SetRow(row)
	return o
}

// SetRow adds the row to the get racks params
func (o *GetRacksParams) SetRow(row *string) {
	o.Row = row
}

// WithSize adds the size to the get racks params
func (o *GetRacksParams) WithSize(size *int64) *GetRacksParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get racks params
func (o *GetRacksParams) SetSize(size *int64) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetRacksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AssetNo != nil {

		// query param asset_no
		var qrAssetNo string
		if o.AssetNo != nil {
			qrAssetNo = *o.AssetNo
		}
		qAssetNo := qrAssetNo
		if qAssetNo != "" {
			if err := r.SetQueryParam("asset_no", qAssetNo); err != nil {
				return err
			}
		}

	}

	if o.Building != nil {

		// query param building
		var qrBuilding string
		if o.Building != nil {
			qrBuilding = *o.Building
		}
		qBuilding := qrBuilding
		if qBuilding != "" {
			if err := r.SetQueryParam("building", qBuilding); err != nil {
				return err
			}
		}

	}

	if o.BuildingID != nil {

		// query param building_id
		var qrBuildingID string
		if o.BuildingID != nil {
			qrBuildingID = *o.BuildingID
		}
		qBuildingID := qrBuildingID
		if qBuildingID != "" {
			if err := r.SetQueryParam("building_id", qBuildingID); err != nil {
				return err
			}
		}

	}

	if o.Manufacturer != nil {

		// query param manufacturer
		var qrManufacturer string
		if o.Manufacturer != nil {
			qrManufacturer = *o.Manufacturer
		}
		qManufacturer := qrManufacturer
		if qManufacturer != "" {
			if err := r.SetQueryParam("manufacturer", qManufacturer); err != nil {
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

	if o.Room != nil {

		// query param room
		var qrRoom string
		if o.Room != nil {
			qrRoom = *o.Room
		}
		qRoom := qrRoom
		if qRoom != "" {
			if err := r.SetQueryParam("room", qRoom); err != nil {
				return err
			}
		}

	}

	if o.RoomID != nil {

		// query param room_id
		var qrRoomID string
		if o.RoomID != nil {
			qrRoomID = *o.RoomID
		}
		qRoomID := qrRoomID
		if qRoomID != "" {
			if err := r.SetQueryParam("room_id", qRoomID); err != nil {
				return err
			}
		}

	}

	if o.Row != nil {

		// query param row
		var qrRow string
		if o.Row != nil {
			qrRow = *o.Row
		}
		qRow := qrRow
		if qRow != "" {
			if err := r.SetQueryParam("row", qRow); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize int64
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt64(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
