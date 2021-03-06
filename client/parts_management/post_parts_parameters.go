// Code generated by go-swagger; DO NOT EDIT.

package parts_management

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

// NewPostPartsParams creates a new PostPartsParams object
// with the default values initialized.
func NewPostPartsParams() *PostPartsParams {
	var ()
	return &PostPartsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPartsParamsWithTimeout creates a new PostPartsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPartsParamsWithTimeout(timeout time.Duration) *PostPartsParams {
	var ()
	return &PostPartsParams{

		timeout: timeout,
	}
}

// NewPostPartsParamsWithContext creates a new PostPartsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPartsParamsWithContext(ctx context.Context) *PostPartsParams {
	var ()
	return &PostPartsParams{

		Context: ctx,
	}
}

// NewPostPartsParamsWithHTTPClient creates a new PostPartsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPartsParamsWithHTTPClient(client *http.Client) *PostPartsParams {
	var ()
	return &PostPartsParams{
		HTTPClient: client,
	}
}

/*PostPartsParams contains all the parameters to send to the API endpoint
for the post parts operation typically these are written to a http.Request
*/
type PostPartsParams struct {

	/*Assignment
	  room, device, rma - required if assigning device

	*/
	Assignment *string
	/*Cores
	  number of cores

	*/
	Cores *string
	/*Count
	  number of parts

	*/
	Count *int64
	/*Cpuspeed
	  enter in MHZ, e.g.: 3.5 GHZ use 3500

	*/
	Cpuspeed *string
	/*DateChanged
	  Update the Date Changed field, using format YYYY-MM-DD HH:MM:SS

	*/
	DateChanged *string
	/*Description*/
	Description *string
	/*Device
	  Room name - required if assigned to device

	*/
	Device *string
	/*Firmware*/
	Firmware *string
	/*Hddrpm
	  new or existing

	*/
	Hddrpm *string
	/*Hddsize
	  enter in GB, e.g.: 250 GB enter 250

	*/
	Hddsize *string
	/*Hddtype
	  new or existing

	*/
	Hddtype *string
	/*Name
	  name of part model - new or existing

	*/
	Name *string
	/*PartID
	  Use for updating existing part

	*/
	PartID *string
	/*PartmodelID*/
	PartmodelID *string
	/*RaidGroup
	  RAID group name

	*/
	RaidGroup *string
	/*RaidType
	  type of RAID

	*/
	RaidType *string
	/*Ramsize
	  enter in MB, e.g.: 8 GB enter 8192

	*/
	Ramsize *string
	/*Ramspeed
	  e.g.: 1600

	*/
	Ramspeed *string
	/*Ramtype
	  e.g.: DDR3

	*/
	Ramtype *string
	/*Room
	  Room name - required if assigned to room

	*/
	Room *string
	/*SerialNo
	  Use for updating existing part. Caution: will update first matching serial if multiple parts with same serial exist. Use part_id or partmodel_id to uniquely identify.

	*/
	SerialNo *string
	/*Threads
	  number of threads

	*/
	Threads *string
	/*Type
	  Part type - new or existing. Must be hdd or harddisk to update HDD model parameters (hddsize, hddtype, etc)

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post parts params
func (o *PostPartsParams) WithTimeout(timeout time.Duration) *PostPartsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post parts params
func (o *PostPartsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post parts params
func (o *PostPartsParams) WithContext(ctx context.Context) *PostPartsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post parts params
func (o *PostPartsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post parts params
func (o *PostPartsParams) WithHTTPClient(client *http.Client) *PostPartsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post parts params
func (o *PostPartsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssignment adds the assignment to the post parts params
func (o *PostPartsParams) WithAssignment(assignment *string) *PostPartsParams {
	o.SetAssignment(assignment)
	return o
}

// SetAssignment adds the assignment to the post parts params
func (o *PostPartsParams) SetAssignment(assignment *string) {
	o.Assignment = assignment
}

// WithCores adds the cores to the post parts params
func (o *PostPartsParams) WithCores(cores *string) *PostPartsParams {
	o.SetCores(cores)
	return o
}

// SetCores adds the cores to the post parts params
func (o *PostPartsParams) SetCores(cores *string) {
	o.Cores = cores
}

// WithCount adds the count to the post parts params
func (o *PostPartsParams) WithCount(count *int64) *PostPartsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the post parts params
func (o *PostPartsParams) SetCount(count *int64) {
	o.Count = count
}

// WithCpuspeed adds the cpuspeed to the post parts params
func (o *PostPartsParams) WithCpuspeed(cpuspeed *string) *PostPartsParams {
	o.SetCpuspeed(cpuspeed)
	return o
}

// SetCpuspeed adds the cpuspeed to the post parts params
func (o *PostPartsParams) SetCpuspeed(cpuspeed *string) {
	o.Cpuspeed = cpuspeed
}

// WithDateChanged adds the dateChanged to the post parts params
func (o *PostPartsParams) WithDateChanged(dateChanged *string) *PostPartsParams {
	o.SetDateChanged(dateChanged)
	return o
}

// SetDateChanged adds the dateChanged to the post parts params
func (o *PostPartsParams) SetDateChanged(dateChanged *string) {
	o.DateChanged = dateChanged
}

// WithDescription adds the description to the post parts params
func (o *PostPartsParams) WithDescription(description *string) *PostPartsParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the post parts params
func (o *PostPartsParams) SetDescription(description *string) {
	o.Description = description
}

// WithDevice adds the device to the post parts params
func (o *PostPartsParams) WithDevice(device *string) *PostPartsParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the post parts params
func (o *PostPartsParams) SetDevice(device *string) {
	o.Device = device
}

// WithFirmware adds the firmware to the post parts params
func (o *PostPartsParams) WithFirmware(firmware *string) *PostPartsParams {
	o.SetFirmware(firmware)
	return o
}

// SetFirmware adds the firmware to the post parts params
func (o *PostPartsParams) SetFirmware(firmware *string) {
	o.Firmware = firmware
}

// WithHddrpm adds the hddrpm to the post parts params
func (o *PostPartsParams) WithHddrpm(hddrpm *string) *PostPartsParams {
	o.SetHddrpm(hddrpm)
	return o
}

// SetHddrpm adds the hddrpm to the post parts params
func (o *PostPartsParams) SetHddrpm(hddrpm *string) {
	o.Hddrpm = hddrpm
}

// WithHddsize adds the hddsize to the post parts params
func (o *PostPartsParams) WithHddsize(hddsize *string) *PostPartsParams {
	o.SetHddsize(hddsize)
	return o
}

// SetHddsize adds the hddsize to the post parts params
func (o *PostPartsParams) SetHddsize(hddsize *string) {
	o.Hddsize = hddsize
}

// WithHddtype adds the hddtype to the post parts params
func (o *PostPartsParams) WithHddtype(hddtype *string) *PostPartsParams {
	o.SetHddtype(hddtype)
	return o
}

// SetHddtype adds the hddtype to the post parts params
func (o *PostPartsParams) SetHddtype(hddtype *string) {
	o.Hddtype = hddtype
}

// WithName adds the name to the post parts params
func (o *PostPartsParams) WithName(name *string) *PostPartsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post parts params
func (o *PostPartsParams) SetName(name *string) {
	o.Name = name
}

// WithPartID adds the partID to the post parts params
func (o *PostPartsParams) WithPartID(partID *string) *PostPartsParams {
	o.SetPartID(partID)
	return o
}

// SetPartID adds the partId to the post parts params
func (o *PostPartsParams) SetPartID(partID *string) {
	o.PartID = partID
}

// WithPartmodelID adds the partmodelID to the post parts params
func (o *PostPartsParams) WithPartmodelID(partmodelID *string) *PostPartsParams {
	o.SetPartmodelID(partmodelID)
	return o
}

// SetPartmodelID adds the partmodelId to the post parts params
func (o *PostPartsParams) SetPartmodelID(partmodelID *string) {
	o.PartmodelID = partmodelID
}

// WithRaidGroup adds the raidGroup to the post parts params
func (o *PostPartsParams) WithRaidGroup(raidGroup *string) *PostPartsParams {
	o.SetRaidGroup(raidGroup)
	return o
}

// SetRaidGroup adds the raidGroup to the post parts params
func (o *PostPartsParams) SetRaidGroup(raidGroup *string) {
	o.RaidGroup = raidGroup
}

// WithRaidType adds the raidType to the post parts params
func (o *PostPartsParams) WithRaidType(raidType *string) *PostPartsParams {
	o.SetRaidType(raidType)
	return o
}

// SetRaidType adds the raidType to the post parts params
func (o *PostPartsParams) SetRaidType(raidType *string) {
	o.RaidType = raidType
}

// WithRamsize adds the ramsize to the post parts params
func (o *PostPartsParams) WithRamsize(ramsize *string) *PostPartsParams {
	o.SetRamsize(ramsize)
	return o
}

// SetRamsize adds the ramsize to the post parts params
func (o *PostPartsParams) SetRamsize(ramsize *string) {
	o.Ramsize = ramsize
}

// WithRamspeed adds the ramspeed to the post parts params
func (o *PostPartsParams) WithRamspeed(ramspeed *string) *PostPartsParams {
	o.SetRamspeed(ramspeed)
	return o
}

// SetRamspeed adds the ramspeed to the post parts params
func (o *PostPartsParams) SetRamspeed(ramspeed *string) {
	o.Ramspeed = ramspeed
}

// WithRamtype adds the ramtype to the post parts params
func (o *PostPartsParams) WithRamtype(ramtype *string) *PostPartsParams {
	o.SetRamtype(ramtype)
	return o
}

// SetRamtype adds the ramtype to the post parts params
func (o *PostPartsParams) SetRamtype(ramtype *string) {
	o.Ramtype = ramtype
}

// WithRoom adds the room to the post parts params
func (o *PostPartsParams) WithRoom(room *string) *PostPartsParams {
	o.SetRoom(room)
	return o
}

// SetRoom adds the room to the post parts params
func (o *PostPartsParams) SetRoom(room *string) {
	o.Room = room
}

// WithSerialNo adds the serialNo to the post parts params
func (o *PostPartsParams) WithSerialNo(serialNo *string) *PostPartsParams {
	o.SetSerialNo(serialNo)
	return o
}

// SetSerialNo adds the serialNo to the post parts params
func (o *PostPartsParams) SetSerialNo(serialNo *string) {
	o.SerialNo = serialNo
}

// WithThreads adds the threads to the post parts params
func (o *PostPartsParams) WithThreads(threads *string) *PostPartsParams {
	o.SetThreads(threads)
	return o
}

// SetThreads adds the threads to the post parts params
func (o *PostPartsParams) SetThreads(threads *string) {
	o.Threads = threads
}

// WithType adds the typeVar to the post parts params
func (o *PostPartsParams) WithType(typeVar *string) *PostPartsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the post parts params
func (o *PostPartsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *PostPartsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Assignment != nil {

		// form param assignment
		var frAssignment string
		if o.Assignment != nil {
			frAssignment = *o.Assignment
		}
		fAssignment := frAssignment
		if fAssignment != "" {
			if err := r.SetFormParam("assignment", fAssignment); err != nil {
				return err
			}
		}

	}

	if o.Cores != nil {

		// form param cores
		var frCores string
		if o.Cores != nil {
			frCores = *o.Cores
		}
		fCores := frCores
		if fCores != "" {
			if err := r.SetFormParam("cores", fCores); err != nil {
				return err
			}
		}

	}

	if o.Count != nil {

		// form param count
		var frCount int64
		if o.Count != nil {
			frCount = *o.Count
		}
		fCount := swag.FormatInt64(frCount)
		if fCount != "" {
			if err := r.SetFormParam("count", fCount); err != nil {
				return err
			}
		}

	}

	if o.Cpuspeed != nil {

		// form param cpuspeed
		var frCpuspeed string
		if o.Cpuspeed != nil {
			frCpuspeed = *o.Cpuspeed
		}
		fCpuspeed := frCpuspeed
		if fCpuspeed != "" {
			if err := r.SetFormParam("cpuspeed", fCpuspeed); err != nil {
				return err
			}
		}

	}

	if o.DateChanged != nil {

		// form param date_changed
		var frDateChanged string
		if o.DateChanged != nil {
			frDateChanged = *o.DateChanged
		}
		fDateChanged := frDateChanged
		if fDateChanged != "" {
			if err := r.SetFormParam("date_changed", fDateChanged); err != nil {
				return err
			}
		}

	}

	if o.Description != nil {

		// form param description
		var frDescription string
		if o.Description != nil {
			frDescription = *o.Description
		}
		fDescription := frDescription
		if fDescription != "" {
			if err := r.SetFormParam("description", fDescription); err != nil {
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

	if o.Firmware != nil {

		// form param firmware
		var frFirmware string
		if o.Firmware != nil {
			frFirmware = *o.Firmware
		}
		fFirmware := frFirmware
		if fFirmware != "" {
			if err := r.SetFormParam("firmware", fFirmware); err != nil {
				return err
			}
		}

	}

	if o.Hddrpm != nil {

		// form param hddrpm
		var frHddrpm string
		if o.Hddrpm != nil {
			frHddrpm = *o.Hddrpm
		}
		fHddrpm := frHddrpm
		if fHddrpm != "" {
			if err := r.SetFormParam("hddrpm", fHddrpm); err != nil {
				return err
			}
		}

	}

	if o.Hddsize != nil {

		// form param hddsize
		var frHddsize string
		if o.Hddsize != nil {
			frHddsize = *o.Hddsize
		}
		fHddsize := frHddsize
		if fHddsize != "" {
			if err := r.SetFormParam("hddsize", fHddsize); err != nil {
				return err
			}
		}

	}

	if o.Hddtype != nil {

		// form param hddtype
		var frHddtype string
		if o.Hddtype != nil {
			frHddtype = *o.Hddtype
		}
		fHddtype := frHddtype
		if fHddtype != "" {
			if err := r.SetFormParam("hddtype", fHddtype); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// form param name
		var frName string
		if o.Name != nil {
			frName = *o.Name
		}
		fName := frName
		if fName != "" {
			if err := r.SetFormParam("name", fName); err != nil {
				return err
			}
		}

	}

	if o.PartID != nil {

		// form param part_id
		var frPartID string
		if o.PartID != nil {
			frPartID = *o.PartID
		}
		fPartID := frPartID
		if fPartID != "" {
			if err := r.SetFormParam("part_id", fPartID); err != nil {
				return err
			}
		}

	}

	if o.PartmodelID != nil {

		// form param partmodel_id
		var frPartmodelID string
		if o.PartmodelID != nil {
			frPartmodelID = *o.PartmodelID
		}
		fPartmodelID := frPartmodelID
		if fPartmodelID != "" {
			if err := r.SetFormParam("partmodel_id", fPartmodelID); err != nil {
				return err
			}
		}

	}

	if o.RaidGroup != nil {

		// form param raid_group
		var frRaidGroup string
		if o.RaidGroup != nil {
			frRaidGroup = *o.RaidGroup
		}
		fRaidGroup := frRaidGroup
		if fRaidGroup != "" {
			if err := r.SetFormParam("raid_group", fRaidGroup); err != nil {
				return err
			}
		}

	}

	if o.RaidType != nil {

		// form param raid_type
		var frRaidType string
		if o.RaidType != nil {
			frRaidType = *o.RaidType
		}
		fRaidType := frRaidType
		if fRaidType != "" {
			if err := r.SetFormParam("raid_type", fRaidType); err != nil {
				return err
			}
		}

	}

	if o.Ramsize != nil {

		// form param ramsize
		var frRamsize string
		if o.Ramsize != nil {
			frRamsize = *o.Ramsize
		}
		fRamsize := frRamsize
		if fRamsize != "" {
			if err := r.SetFormParam("ramsize", fRamsize); err != nil {
				return err
			}
		}

	}

	if o.Ramspeed != nil {

		// form param ramspeed
		var frRamspeed string
		if o.Ramspeed != nil {
			frRamspeed = *o.Ramspeed
		}
		fRamspeed := frRamspeed
		if fRamspeed != "" {
			if err := r.SetFormParam("ramspeed", fRamspeed); err != nil {
				return err
			}
		}

	}

	if o.Ramtype != nil {

		// form param ramtype
		var frRamtype string
		if o.Ramtype != nil {
			frRamtype = *o.Ramtype
		}
		fRamtype := frRamtype
		if fRamtype != "" {
			if err := r.SetFormParam("ramtype", fRamtype); err != nil {
				return err
			}
		}

	}

	if o.Room != nil {

		// form param room
		var frRoom string
		if o.Room != nil {
			frRoom = *o.Room
		}
		fRoom := frRoom
		if fRoom != "" {
			if err := r.SetFormParam("room", fRoom); err != nil {
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

	if o.Threads != nil {

		// form param threads
		var frThreads string
		if o.Threads != nil {
			frThreads = *o.Threads
		}
		fThreads := frThreads
		if fThreads != "" {
			if err := r.SetFormParam("threads", fThreads); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
