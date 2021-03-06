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

// NewPostIPAMMacsParams creates a new PostIPAMMacsParams object
// with the default values initialized.
func NewPostIPAMMacsParams() *PostIPAMMacsParams {
	var ()
	return &PostIPAMMacsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIPAMMacsParamsWithTimeout creates a new PostIPAMMacsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIPAMMacsParamsWithTimeout(timeout time.Duration) *PostIPAMMacsParams {
	var ()
	return &PostIPAMMacsParams{

		timeout: timeout,
	}
}

// NewPostIPAMMacsParamsWithContext creates a new PostIPAMMacsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIPAMMacsParamsWithContext(ctx context.Context) *PostIPAMMacsParams {
	var ()
	return &PostIPAMMacsParams{

		Context: ctx,
	}
}

// NewPostIPAMMacsParamsWithHTTPClient creates a new PostIPAMMacsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIPAMMacsParamsWithHTTPClient(client *http.Client) *PostIPAMMacsParams {
	var ()
	return &PostIPAMMacsParams{
		HTTPClient: client,
	}
}

/*PostIPAMMacsParams contains all the parameters to send to the API endpoint
for the post IP a m macs operation typically these are written to a http.Request
*/
type PostIPAMMacsParams struct {

	/*Device
	  name of the device

	*/
	Device *string
	/*Macaddress
	  MAC address – can be new or existing

	*/
	Macaddress *string
	/*Override
	  smart – will detect if the port_name passed exist or not, if not – it is added to the current port name. Helpful, if you want to track all the port names for that mac address (e.g. eth0 & bond0).<br>yes – change the port name. This is default behavior even if you don’t pass this parameter<br>no – will not change the port name

	*/
	Override *string
	/*Port
	  Refers to the switchport name (not the interface name) - Use with parameter switch

	*/
	Port *string
	/*PortID
	  Use this parameter or a combination of port and switch to specify the port.

	*/
	PortID *string
	/*PortName
	  Interface name. (Please note: This is NOT the switchport name.)

	*/
	PortName *string
	/*Switch
	  Refers to the device name of the switch

	*/
	Switch *string
	/*VlanID
	  ID of the vlan

	*/
	VlanID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post IP a m macs params
func (o *PostIPAMMacsParams) WithTimeout(timeout time.Duration) *PostIPAMMacsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post IP a m macs params
func (o *PostIPAMMacsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post IP a m macs params
func (o *PostIPAMMacsParams) WithContext(ctx context.Context) *PostIPAMMacsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post IP a m macs params
func (o *PostIPAMMacsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post IP a m macs params
func (o *PostIPAMMacsParams) WithHTTPClient(client *http.Client) *PostIPAMMacsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post IP a m macs params
func (o *PostIPAMMacsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevice adds the device to the post IP a m macs params
func (o *PostIPAMMacsParams) WithDevice(device *string) *PostIPAMMacsParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the post IP a m macs params
func (o *PostIPAMMacsParams) SetDevice(device *string) {
	o.Device = device
}

// WithMacaddress adds the macaddress to the post IP a m macs params
func (o *PostIPAMMacsParams) WithMacaddress(macaddress *string) *PostIPAMMacsParams {
	o.SetMacaddress(macaddress)
	return o
}

// SetMacaddress adds the macaddress to the post IP a m macs params
func (o *PostIPAMMacsParams) SetMacaddress(macaddress *string) {
	o.Macaddress = macaddress
}

// WithOverride adds the override to the post IP a m macs params
func (o *PostIPAMMacsParams) WithOverride(override *string) *PostIPAMMacsParams {
	o.SetOverride(override)
	return o
}

// SetOverride adds the override to the post IP a m macs params
func (o *PostIPAMMacsParams) SetOverride(override *string) {
	o.Override = override
}

// WithPort adds the port to the post IP a m macs params
func (o *PostIPAMMacsParams) WithPort(port *string) *PostIPAMMacsParams {
	o.SetPort(port)
	return o
}

// SetPort adds the port to the post IP a m macs params
func (o *PostIPAMMacsParams) SetPort(port *string) {
	o.Port = port
}

// WithPortID adds the portID to the post IP a m macs params
func (o *PostIPAMMacsParams) WithPortID(portID *string) *PostIPAMMacsParams {
	o.SetPortID(portID)
	return o
}

// SetPortID adds the portId to the post IP a m macs params
func (o *PostIPAMMacsParams) SetPortID(portID *string) {
	o.PortID = portID
}

// WithPortName adds the portName to the post IP a m macs params
func (o *PostIPAMMacsParams) WithPortName(portName *string) *PostIPAMMacsParams {
	o.SetPortName(portName)
	return o
}

// SetPortName adds the portName to the post IP a m macs params
func (o *PostIPAMMacsParams) SetPortName(portName *string) {
	o.PortName = portName
}

// WithSwitch adds the switchVar to the post IP a m macs params
func (o *PostIPAMMacsParams) WithSwitch(switchVar *string) *PostIPAMMacsParams {
	o.SetSwitch(switchVar)
	return o
}

// SetSwitch adds the switch to the post IP a m macs params
func (o *PostIPAMMacsParams) SetSwitch(switchVar *string) {
	o.Switch = switchVar
}

// WithVlanID adds the vlanID to the post IP a m macs params
func (o *PostIPAMMacsParams) WithVlanID(vlanID *string) *PostIPAMMacsParams {
	o.SetVlanID(vlanID)
	return o
}

// SetVlanID adds the vlanId to the post IP a m macs params
func (o *PostIPAMMacsParams) SetVlanID(vlanID *string) {
	o.VlanID = vlanID
}

// WriteToRequest writes these params to a swagger request
func (o *PostIPAMMacsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Macaddress != nil {

		// form param macaddress
		var frMacaddress string
		if o.Macaddress != nil {
			frMacaddress = *o.Macaddress
		}
		fMacaddress := frMacaddress
		if fMacaddress != "" {
			if err := r.SetFormParam("macaddress", fMacaddress); err != nil {
				return err
			}
		}

	}

	if o.Override != nil {

		// form param override
		var frOverride string
		if o.Override != nil {
			frOverride = *o.Override
		}
		fOverride := frOverride
		if fOverride != "" {
			if err := r.SetFormParam("override", fOverride); err != nil {
				return err
			}
		}

	}

	if o.Port != nil {

		// form param port
		var frPort string
		if o.Port != nil {
			frPort = *o.Port
		}
		fPort := frPort
		if fPort != "" {
			if err := r.SetFormParam("port", fPort); err != nil {
				return err
			}
		}

	}

	if o.PortID != nil {

		// form param port_id
		var frPortID string
		if o.PortID != nil {
			frPortID = *o.PortID
		}
		fPortID := frPortID
		if fPortID != "" {
			if err := r.SetFormParam("port_id", fPortID); err != nil {
				return err
			}
		}

	}

	if o.PortName != nil {

		// form param port_name
		var frPortName string
		if o.PortName != nil {
			frPortName = *o.PortName
		}
		fPortName := frPortName
		if fPortName != "" {
			if err := r.SetFormParam("port_name", fPortName); err != nil {
				return err
			}
		}

	}

	if o.Switch != nil {

		// form param switch
		var frSwitch string
		if o.Switch != nil {
			frSwitch = *o.Switch
		}
		fSwitch := frSwitch
		if fSwitch != "" {
			if err := r.SetFormParam("switch", fSwitch); err != nil {
				return err
			}
		}

	}

	if o.VlanID != nil {

		// form param vlan_id
		var frVlanID string
		if o.VlanID != nil {
			frVlanID = *o.VlanID
		}
		fVlanID := frVlanID
		if fVlanID != "" {
			if err := r.SetFormParam("vlan_id", fVlanID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
