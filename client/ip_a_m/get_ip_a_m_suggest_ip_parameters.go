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

// NewGetIPAMSuggestIPParams creates a new GetIPAMSuggestIPParams object
// with the default values initialized.
func NewGetIPAMSuggestIPParams() *GetIPAMSuggestIPParams {
	var ()
	return &GetIPAMSuggestIPParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPAMSuggestIPParamsWithTimeout creates a new GetIPAMSuggestIPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPAMSuggestIPParamsWithTimeout(timeout time.Duration) *GetIPAMSuggestIPParams {
	var ()
	return &GetIPAMSuggestIPParams{

		timeout: timeout,
	}
}

// NewGetIPAMSuggestIPParamsWithContext creates a new GetIPAMSuggestIPParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPAMSuggestIPParamsWithContext(ctx context.Context) *GetIPAMSuggestIPParams {
	var ()
	return &GetIPAMSuggestIPParams{

		Context: ctx,
	}
}

// NewGetIPAMSuggestIPParamsWithHTTPClient creates a new GetIPAMSuggestIPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPAMSuggestIPParamsWithHTTPClient(client *http.Client) *GetIPAMSuggestIPParams {
	var ()
	return &GetIPAMSuggestIPParams{
		HTTPClient: client,
	}
}

/*GetIPAMSuggestIPParams contains all the parameters to send to the API endpoint
for the get IP a m suggest ip operation typically these are written to a http.Request
*/
type GetIPAMSuggestIPParams struct {

	/*MaskBits
	  mask bits (added in v7.2.0)

	*/
	MaskBits string
	/*Name
	  filter by name (Added in v6.0.0)

	*/
	Name *string
	/*Number
	  vlan number

	*/
	Number *string
	/*ReserveIP
	  If value of yes is passed, the suggested IP is reserved. Return value also adds reserved as yes or no. (added in v7.2.0)

	*/
	ReserveIP *string
	/*Subnet
	  name of the subnet

	*/
	Subnet *string
	/*SubnetID
	  ID of the subnet (added in v7.2.0)

	*/
	SubnetID *string
	/*VrfGroup
	  VRF group name

	*/
	VrfGroup *string
	/*VrfGroupID
	  ID of the VRF group

	*/
	VrfGroupID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) WithTimeout(timeout time.Duration) *GetIPAMSuggestIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) WithContext(ctx context.Context) *GetIPAMSuggestIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) WithHTTPClient(client *http.Client) *GetIPAMSuggestIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaskBits adds the maskBits to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) WithMaskBits(maskBits string) *GetIPAMSuggestIPParams {
	o.SetMaskBits(maskBits)
	return o
}

// SetMaskBits adds the maskBits to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) SetMaskBits(maskBits string) {
	o.MaskBits = maskBits
}

// WithName adds the name to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) WithName(name *string) *GetIPAMSuggestIPParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) SetName(name *string) {
	o.Name = name
}

// WithNumber adds the number to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) WithNumber(number *string) *GetIPAMSuggestIPParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) SetNumber(number *string) {
	o.Number = number
}

// WithReserveIP adds the reserveIP to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) WithReserveIP(reserveIP *string) *GetIPAMSuggestIPParams {
	o.SetReserveIP(reserveIP)
	return o
}

// SetReserveIP adds the reserveIp to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) SetReserveIP(reserveIP *string) {
	o.ReserveIP = reserveIP
}

// WithSubnet adds the subnet to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) WithSubnet(subnet *string) *GetIPAMSuggestIPParams {
	o.SetSubnet(subnet)
	return o
}

// SetSubnet adds the subnet to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) SetSubnet(subnet *string) {
	o.Subnet = subnet
}

// WithSubnetID adds the subnetID to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) WithSubnetID(subnetID *string) *GetIPAMSuggestIPParams {
	o.SetSubnetID(subnetID)
	return o
}

// SetSubnetID adds the subnetId to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) SetSubnetID(subnetID *string) {
	o.SubnetID = subnetID
}

// WithVrfGroup adds the vrfGroup to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) WithVrfGroup(vrfGroup *string) *GetIPAMSuggestIPParams {
	o.SetVrfGroup(vrfGroup)
	return o
}

// SetVrfGroup adds the vrfGroup to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) SetVrfGroup(vrfGroup *string) {
	o.VrfGroup = vrfGroup
}

// WithVrfGroupID adds the vrfGroupID to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) WithVrfGroupID(vrfGroupID *string) *GetIPAMSuggestIPParams {
	o.SetVrfGroupID(vrfGroupID)
	return o
}

// SetVrfGroupID adds the vrfGroupId to the get IP a m suggest ip params
func (o *GetIPAMSuggestIPParams) SetVrfGroupID(vrfGroupID *string) {
	o.VrfGroupID = vrfGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPAMSuggestIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param mask_bits
	qrMaskBits := o.MaskBits
	qMaskBits := qrMaskBits
	if qMaskBits != "" {
		if err := r.SetQueryParam("mask_bits", qMaskBits); err != nil {
			return err
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

	if o.Number != nil {

		// query param number
		var qrNumber string
		if o.Number != nil {
			qrNumber = *o.Number
		}
		qNumber := qrNumber
		if qNumber != "" {
			if err := r.SetQueryParam("number", qNumber); err != nil {
				return err
			}
		}

	}

	if o.ReserveIP != nil {

		// query param reserve_ip
		var qrReserveIP string
		if o.ReserveIP != nil {
			qrReserveIP = *o.ReserveIP
		}
		qReserveIP := qrReserveIP
		if qReserveIP != "" {
			if err := r.SetQueryParam("reserve_ip", qReserveIP); err != nil {
				return err
			}
		}

	}

	if o.Subnet != nil {

		// query param subnet
		var qrSubnet string
		if o.Subnet != nil {
			qrSubnet = *o.Subnet
		}
		qSubnet := qrSubnet
		if qSubnet != "" {
			if err := r.SetQueryParam("subnet", qSubnet); err != nil {
				return err
			}
		}

	}

	if o.SubnetID != nil {

		// query param subnet_id
		var qrSubnetID string
		if o.SubnetID != nil {
			qrSubnetID = *o.SubnetID
		}
		qSubnetID := qrSubnetID
		if qSubnetID != "" {
			if err := r.SetQueryParam("subnet_id", qSubnetID); err != nil {
				return err
			}
		}

	}

	if o.VrfGroup != nil {

		// query param vrf_group
		var qrVrfGroup string
		if o.VrfGroup != nil {
			qrVrfGroup = *o.VrfGroup
		}
		qVrfGroup := qrVrfGroup
		if qVrfGroup != "" {
			if err := r.SetQueryParam("vrf_group", qVrfGroup); err != nil {
				return err
			}
		}

	}

	if o.VrfGroupID != nil {

		// query param vrf_group_id
		var qrVrfGroupID string
		if o.VrfGroupID != nil {
			qrVrfGroupID = *o.VrfGroupID
		}
		qVrfGroupID := qrVrfGroupID
		if qVrfGroupID != "" {
			if err := r.SetQueryParam("vrf_group_id", qVrfGroupID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
