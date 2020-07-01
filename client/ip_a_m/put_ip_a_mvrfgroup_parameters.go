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

// NewPutIPAMvrfgroupParams creates a new PutIPAMvrfgroupParams object
// with the default values initialized.
func NewPutIPAMvrfgroupParams() *PutIPAMvrfgroupParams {
	var ()
	return &PutIPAMvrfgroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutIPAMvrfgroupParamsWithTimeout creates a new PutIPAMvrfgroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutIPAMvrfgroupParamsWithTimeout(timeout time.Duration) *PutIPAMvrfgroupParams {
	var ()
	return &PutIPAMvrfgroupParams{

		timeout: timeout,
	}
}

// NewPutIPAMvrfgroupParamsWithContext creates a new PutIPAMvrfgroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutIPAMvrfgroupParamsWithContext(ctx context.Context) *PutIPAMvrfgroupParams {
	var ()
	return &PutIPAMvrfgroupParams{

		Context: ctx,
	}
}

// NewPutIPAMvrfgroupParamsWithHTTPClient creates a new PutIPAMvrfgroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutIPAMvrfgroupParamsWithHTTPClient(client *http.Client) *PutIPAMvrfgroupParams {
	var ()
	return &PutIPAMvrfgroupParams{
		HTTPClient: client,
	}
}

/*PutIPAMvrfgroupParams contains all the parameters to send to the API endpoint
for the put IP a mvrfgroup operation typically these are written to a http.Request
*/
type PutIPAMvrfgroupParams struct {

	/*Buildings
	  list of building names for the VRF Group

	*/
	Buildings *string
	/*Description*/
	Description *string
	/*Groups
	  If multitenancy is on, admin groups that have access to this object are specified here, e.g. Prod_East:no,Corp:yes specifies that the admin groups for this object are Prod_East with view only permission and Corp with change permission. If this parameter is present with no value, all groups are deleted.

	*/
	Groups *string
	/*Name
	  Name of the VRF Group

	*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put IP a mvrfgroup params
func (o *PutIPAMvrfgroupParams) WithTimeout(timeout time.Duration) *PutIPAMvrfgroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put IP a mvrfgroup params
func (o *PutIPAMvrfgroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put IP a mvrfgroup params
func (o *PutIPAMvrfgroupParams) WithContext(ctx context.Context) *PutIPAMvrfgroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put IP a mvrfgroup params
func (o *PutIPAMvrfgroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put IP a mvrfgroup params
func (o *PutIPAMvrfgroupParams) WithHTTPClient(client *http.Client) *PutIPAMvrfgroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put IP a mvrfgroup params
func (o *PutIPAMvrfgroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildings adds the buildings to the put IP a mvrfgroup params
func (o *PutIPAMvrfgroupParams) WithBuildings(buildings *string) *PutIPAMvrfgroupParams {
	o.SetBuildings(buildings)
	return o
}

// SetBuildings adds the buildings to the put IP a mvrfgroup params
func (o *PutIPAMvrfgroupParams) SetBuildings(buildings *string) {
	o.Buildings = buildings
}

// WithDescription adds the description to the put IP a mvrfgroup params
func (o *PutIPAMvrfgroupParams) WithDescription(description *string) *PutIPAMvrfgroupParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the put IP a mvrfgroup params
func (o *PutIPAMvrfgroupParams) SetDescription(description *string) {
	o.Description = description
}

// WithGroups adds the groups to the put IP a mvrfgroup params
func (o *PutIPAMvrfgroupParams) WithGroups(groups *string) *PutIPAMvrfgroupParams {
	o.SetGroups(groups)
	return o
}

// SetGroups adds the groups to the put IP a mvrfgroup params
func (o *PutIPAMvrfgroupParams) SetGroups(groups *string) {
	o.Groups = groups
}

// WithName adds the name to the put IP a mvrfgroup params
func (o *PutIPAMvrfgroupParams) WithName(name *string) *PutIPAMvrfgroupParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put IP a mvrfgroup params
func (o *PutIPAMvrfgroupParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PutIPAMvrfgroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Buildings != nil {

		// form param buildings
		var frBuildings string
		if o.Buildings != nil {
			frBuildings = *o.Buildings
		}
		fBuildings := frBuildings
		if fBuildings != "" {
			if err := r.SetFormParam("buildings", fBuildings); err != nil {
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

	if o.Groups != nil {

		// form param groups
		var frGroups string
		if o.Groups != nil {
			frGroups = *o.Groups
		}
		fGroups := frGroups
		if fGroups != "" {
			if err := r.SetFormParam("groups", fGroups); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
