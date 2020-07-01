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

// NewPostIPAMSubnetsCreateChildParams creates a new PostIPAMSubnetsCreateChildParams object
// with the default values initialized.
func NewPostIPAMSubnetsCreateChildParams() *PostIPAMSubnetsCreateChildParams {
	var ()
	return &PostIPAMSubnetsCreateChildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIPAMSubnetsCreateChildParamsWithTimeout creates a new PostIPAMSubnetsCreateChildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIPAMSubnetsCreateChildParamsWithTimeout(timeout time.Duration) *PostIPAMSubnetsCreateChildParams {
	var ()
	return &PostIPAMSubnetsCreateChildParams{

		timeout: timeout,
	}
}

// NewPostIPAMSubnetsCreateChildParamsWithContext creates a new PostIPAMSubnetsCreateChildParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIPAMSubnetsCreateChildParamsWithContext(ctx context.Context) *PostIPAMSubnetsCreateChildParams {
	var ()
	return &PostIPAMSubnetsCreateChildParams{

		Context: ctx,
	}
}

// NewPostIPAMSubnetsCreateChildParamsWithHTTPClient creates a new PostIPAMSubnetsCreateChildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIPAMSubnetsCreateChildParamsWithHTTPClient(client *http.Client) *PostIPAMSubnetsCreateChildParams {
	var ()
	return &PostIPAMSubnetsCreateChildParams{
		HTTPClient: client,
	}
}

/*PostIPAMSubnetsCreateChildParams contains all the parameters to send to the API endpoint
for the post IP a m subnets create child operation typically these are written to a http.Request
*/
type PostIPAMSubnetsCreateChildParams struct {

	/*IPV6
	  Required if creating an ipv6 subnet

	*/
	IPV6 *string
	/*MaskBits
	  e.g. 24

	*/
	MaskBits string
	/*Network*/
	Network *string
	/*ParentMaskBits
	  only if searching within a VRF and you want to restrict to certain parents with particular mask bits (added in v9.0.0)

	*/
	ParentMaskBits *string
	/*ParentSubnetID
	  ID of the parent subnet. Can be obtained via /api/api/1.0/subnets/ or Tools > Export > Subnet. Required if vrf_group and vrf_group_id are not present.

	*/
	ParentSubnetID *string
	/*SubnetID*/
	SubnetID *string
	/*VrfGroup
	  Name of the VRF group. Required if parent_subnet_id or vrf_group_id are not present.

	*/
	VrfGroup *string
	/*VrfGroupID
	  ID of the VRF group. Required if parent_subnet_id or vrf_group are not present.

	*/
	VrfGroupID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) WithTimeout(timeout time.Duration) *PostIPAMSubnetsCreateChildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) WithContext(ctx context.Context) *PostIPAMSubnetsCreateChildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) WithHTTPClient(client *http.Client) *PostIPAMSubnetsCreateChildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIPV6 adds the iPV6 to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) WithIPV6(iPV6 *string) *PostIPAMSubnetsCreateChildParams {
	o.SetIPV6(iPV6)
	return o
}

// SetIPV6 adds the ipv6 to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) SetIPV6(iPV6 *string) {
	o.IPV6 = iPV6
}

// WithMaskBits adds the maskBits to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) WithMaskBits(maskBits string) *PostIPAMSubnetsCreateChildParams {
	o.SetMaskBits(maskBits)
	return o
}

// SetMaskBits adds the maskBits to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) SetMaskBits(maskBits string) {
	o.MaskBits = maskBits
}

// WithNetwork adds the network to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) WithNetwork(network *string) *PostIPAMSubnetsCreateChildParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) SetNetwork(network *string) {
	o.Network = network
}

// WithParentMaskBits adds the parentMaskBits to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) WithParentMaskBits(parentMaskBits *string) *PostIPAMSubnetsCreateChildParams {
	o.SetParentMaskBits(parentMaskBits)
	return o
}

// SetParentMaskBits adds the parentMaskBits to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) SetParentMaskBits(parentMaskBits *string) {
	o.ParentMaskBits = parentMaskBits
}

// WithParentSubnetID adds the parentSubnetID to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) WithParentSubnetID(parentSubnetID *string) *PostIPAMSubnetsCreateChildParams {
	o.SetParentSubnetID(parentSubnetID)
	return o
}

// SetParentSubnetID adds the parentSubnetId to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) SetParentSubnetID(parentSubnetID *string) {
	o.ParentSubnetID = parentSubnetID
}

// WithSubnetID adds the subnetID to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) WithSubnetID(subnetID *string) *PostIPAMSubnetsCreateChildParams {
	o.SetSubnetID(subnetID)
	return o
}

// SetSubnetID adds the subnetId to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) SetSubnetID(subnetID *string) {
	o.SubnetID = subnetID
}

// WithVrfGroup adds the vrfGroup to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) WithVrfGroup(vrfGroup *string) *PostIPAMSubnetsCreateChildParams {
	o.SetVrfGroup(vrfGroup)
	return o
}

// SetVrfGroup adds the vrfGroup to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) SetVrfGroup(vrfGroup *string) {
	o.VrfGroup = vrfGroup
}

// WithVrfGroupID adds the vrfGroupID to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) WithVrfGroupID(vrfGroupID *string) *PostIPAMSubnetsCreateChildParams {
	o.SetVrfGroupID(vrfGroupID)
	return o
}

// SetVrfGroupID adds the vrfGroupId to the post IP a m subnets create child params
func (o *PostIPAMSubnetsCreateChildParams) SetVrfGroupID(vrfGroupID *string) {
	o.VrfGroupID = vrfGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *PostIPAMSubnetsCreateChildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IPV6 != nil {

		// form param ipv6
		var frIPV6 string
		if o.IPV6 != nil {
			frIPV6 = *o.IPV6
		}
		fIPV6 := frIPV6
		if fIPV6 != "" {
			if err := r.SetFormParam("ipv6", fIPV6); err != nil {
				return err
			}
		}

	}

	// form param mask_bits
	frMaskBits := o.MaskBits
	fMaskBits := frMaskBits
	if fMaskBits != "" {
		if err := r.SetFormParam("mask_bits", fMaskBits); err != nil {
			return err
		}
	}

	if o.Network != nil {

		// form param network
		var frNetwork string
		if o.Network != nil {
			frNetwork = *o.Network
		}
		fNetwork := frNetwork
		if fNetwork != "" {
			if err := r.SetFormParam("network", fNetwork); err != nil {
				return err
			}
		}

	}

	if o.ParentMaskBits != nil {

		// form param parent_mask_bits
		var frParentMaskBits string
		if o.ParentMaskBits != nil {
			frParentMaskBits = *o.ParentMaskBits
		}
		fParentMaskBits := frParentMaskBits
		if fParentMaskBits != "" {
			if err := r.SetFormParam("parent_mask_bits", fParentMaskBits); err != nil {
				return err
			}
		}

	}

	if o.ParentSubnetID != nil {

		// form param parent_subnet_id
		var frParentSubnetID string
		if o.ParentSubnetID != nil {
			frParentSubnetID = *o.ParentSubnetID
		}
		fParentSubnetID := frParentSubnetID
		if fParentSubnetID != "" {
			if err := r.SetFormParam("parent_subnet_id", fParentSubnetID); err != nil {
				return err
			}
		}

	}

	if o.SubnetID != nil {

		// form param subnet_id
		var frSubnetID string
		if o.SubnetID != nil {
			frSubnetID = *o.SubnetID
		}
		fSubnetID := frSubnetID
		if fSubnetID != "" {
			if err := r.SetFormParam("subnet_id", fSubnetID); err != nil {
				return err
			}
		}

	}

	if o.VrfGroup != nil {

		// form param vrf_group
		var frVrfGroup string
		if o.VrfGroup != nil {
			frVrfGroup = *o.VrfGroup
		}
		fVrfGroup := frVrfGroup
		if fVrfGroup != "" {
			if err := r.SetFormParam("vrf_group", fVrfGroup); err != nil {
				return err
			}
		}

	}

	if o.VrfGroupID != nil {

		// form param vrf_group_id
		var frVrfGroupID string
		if o.VrfGroupID != nil {
			frVrfGroupID = *o.VrfGroupID
		}
		fVrfGroupID := frVrfGroupID
		if fVrfGroupID != "" {
			if err := r.SetFormParam("vrf_group_id", fVrfGroupID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}