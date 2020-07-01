// Code generated by go-swagger; DO NOT EDIT.

package auto_discovery

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

// NewPostAutoDiscoveryIpmiParams creates a new PostAutoDiscoveryIpmiParams object
// with the default values initialized.
func NewPostAutoDiscoveryIpmiParams() *PostAutoDiscoveryIpmiParams {
	var (
		discoveryTypeDefault = string("IPMI")
	)
	return &PostAutoDiscoveryIpmiParams{
		DiscoveryType: &discoveryTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAutoDiscoveryIpmiParamsWithTimeout creates a new PostAutoDiscoveryIpmiParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAutoDiscoveryIpmiParamsWithTimeout(timeout time.Duration) *PostAutoDiscoveryIpmiParams {
	var (
		discoveryTypeDefault = string("IPMI")
	)
	return &PostAutoDiscoveryIpmiParams{
		DiscoveryType: &discoveryTypeDefault,

		timeout: timeout,
	}
}

// NewPostAutoDiscoveryIpmiParamsWithContext creates a new PostAutoDiscoveryIpmiParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAutoDiscoveryIpmiParamsWithContext(ctx context.Context) *PostAutoDiscoveryIpmiParams {
	var (
		discoveryTypeDefault = string("IPMI")
	)
	return &PostAutoDiscoveryIpmiParams{
		DiscoveryType: &discoveryTypeDefault,

		Context: ctx,
	}
}

// NewPostAutoDiscoveryIpmiParamsWithHTTPClient creates a new PostAutoDiscoveryIpmiParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAutoDiscoveryIpmiParamsWithHTTPClient(client *http.Client) *PostAutoDiscoveryIpmiParams {
	var (
		discoveryTypeDefault = string("IPMI")
	)
	return &PostAutoDiscoveryIpmiParams{
		DiscoveryType: &discoveryTypeDefault,
		HTTPClient:    client,
	}
}

/*PostAutoDiscoveryIpmiParams contains all the parameters to send to the API endpoint
for the post auto discovery ipmi operation typically these are written to a http.Request
*/
type PostAutoDiscoveryIpmiParams struct {

	/*BmcPassword
	  password for discovery

	*/
	BmcPassword string
	/*BmcUser
	  username for discovery

	*/
	BmcUser string
	/*ClearExistingSchedule*/
	ClearExistingSchedule *string
	/*DebugLevel*/
	DebugLevel *string
	/*DiscoveryType*/
	DiscoveryType *string
	/*Groups
	  name of one or more groups separated by commas

	*/
	Groups *string
	/*HostnameToUse*/
	HostnameToUse string
	/*IPEnd
	  ending IP address, use same as start for single address

	*/
	IPEnd string
	/*IPStart
	  starting IP address

	*/
	IPStart string
	/*Name
	  name of the job

	*/
	Name string
	/*ObjectCategory*/
	ObjectCategory *string
	/*OverwriteObjectCategories*/
	OverwriteObjectCategories *string
	/*RemoteCollectorID*/
	RemoteCollectorID *string
	/*RunAsOperator*/
	RunAsOperator *string
	/*ScheduleDays
	  Comma separated days of week, where Monday = 0. e.g. 0,1,2 wil set the job for Mon, Tue and Wed. For multiple schedules, separate with a slash (/).

	*/
	ScheduleDays *string
	/*ScheduleTime
	  Time in HH:MM format if you want to schedule the job. Note: Must be formatted as text NOT date. For multiple schedules, separate with a slash (/).

	*/
	ScheduleTime *string
	/*UpdateModelIfFound*/
	UpdateModelIfFound *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithTimeout(timeout time.Duration) *PostAutoDiscoveryIpmiParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithContext(ctx context.Context) *PostAutoDiscoveryIpmiParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithHTTPClient(client *http.Client) *PostAutoDiscoveryIpmiParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBmcPassword adds the bmcPassword to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithBmcPassword(bmcPassword string) *PostAutoDiscoveryIpmiParams {
	o.SetBmcPassword(bmcPassword)
	return o
}

// SetBmcPassword adds the bmcPassword to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetBmcPassword(bmcPassword string) {
	o.BmcPassword = bmcPassword
}

// WithBmcUser adds the bmcUser to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithBmcUser(bmcUser string) *PostAutoDiscoveryIpmiParams {
	o.SetBmcUser(bmcUser)
	return o
}

// SetBmcUser adds the bmcUser to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetBmcUser(bmcUser string) {
	o.BmcUser = bmcUser
}

// WithClearExistingSchedule adds the clearExistingSchedule to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithClearExistingSchedule(clearExistingSchedule *string) *PostAutoDiscoveryIpmiParams {
	o.SetClearExistingSchedule(clearExistingSchedule)
	return o
}

// SetClearExistingSchedule adds the clearExistingSchedule to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetClearExistingSchedule(clearExistingSchedule *string) {
	o.ClearExistingSchedule = clearExistingSchedule
}

// WithDebugLevel adds the debugLevel to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithDebugLevel(debugLevel *string) *PostAutoDiscoveryIpmiParams {
	o.SetDebugLevel(debugLevel)
	return o
}

// SetDebugLevel adds the debugLevel to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetDebugLevel(debugLevel *string) {
	o.DebugLevel = debugLevel
}

// WithDiscoveryType adds the discoveryType to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithDiscoveryType(discoveryType *string) *PostAutoDiscoveryIpmiParams {
	o.SetDiscoveryType(discoveryType)
	return o
}

// SetDiscoveryType adds the discoveryType to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetDiscoveryType(discoveryType *string) {
	o.DiscoveryType = discoveryType
}

// WithGroups adds the groups to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithGroups(groups *string) *PostAutoDiscoveryIpmiParams {
	o.SetGroups(groups)
	return o
}

// SetGroups adds the groups to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetGroups(groups *string) {
	o.Groups = groups
}

// WithHostnameToUse adds the hostnameToUse to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithHostnameToUse(hostnameToUse string) *PostAutoDiscoveryIpmiParams {
	o.SetHostnameToUse(hostnameToUse)
	return o
}

// SetHostnameToUse adds the hostnameToUse to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetHostnameToUse(hostnameToUse string) {
	o.HostnameToUse = hostnameToUse
}

// WithIPEnd adds the iPEnd to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithIPEnd(iPEnd string) *PostAutoDiscoveryIpmiParams {
	o.SetIPEnd(iPEnd)
	return o
}

// SetIPEnd adds the ipEnd to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetIPEnd(iPEnd string) {
	o.IPEnd = iPEnd
}

// WithIPStart adds the iPStart to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithIPStart(iPStart string) *PostAutoDiscoveryIpmiParams {
	o.SetIPStart(iPStart)
	return o
}

// SetIPStart adds the ipStart to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetIPStart(iPStart string) {
	o.IPStart = iPStart
}

// WithName adds the name to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithName(name string) *PostAutoDiscoveryIpmiParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetName(name string) {
	o.Name = name
}

// WithObjectCategory adds the objectCategory to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithObjectCategory(objectCategory *string) *PostAutoDiscoveryIpmiParams {
	o.SetObjectCategory(objectCategory)
	return o
}

// SetObjectCategory adds the objectCategory to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetObjectCategory(objectCategory *string) {
	o.ObjectCategory = objectCategory
}

// WithOverwriteObjectCategories adds the overwriteObjectCategories to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithOverwriteObjectCategories(overwriteObjectCategories *string) *PostAutoDiscoveryIpmiParams {
	o.SetOverwriteObjectCategories(overwriteObjectCategories)
	return o
}

// SetOverwriteObjectCategories adds the overwriteObjectCategories to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetOverwriteObjectCategories(overwriteObjectCategories *string) {
	o.OverwriteObjectCategories = overwriteObjectCategories
}

// WithRemoteCollectorID adds the remoteCollectorID to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithRemoteCollectorID(remoteCollectorID *string) *PostAutoDiscoveryIpmiParams {
	o.SetRemoteCollectorID(remoteCollectorID)
	return o
}

// SetRemoteCollectorID adds the remoteCollectorId to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetRemoteCollectorID(remoteCollectorID *string) {
	o.RemoteCollectorID = remoteCollectorID
}

// WithRunAsOperator adds the runAsOperator to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithRunAsOperator(runAsOperator *string) *PostAutoDiscoveryIpmiParams {
	o.SetRunAsOperator(runAsOperator)
	return o
}

// SetRunAsOperator adds the runAsOperator to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetRunAsOperator(runAsOperator *string) {
	o.RunAsOperator = runAsOperator
}

// WithScheduleDays adds the scheduleDays to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithScheduleDays(scheduleDays *string) *PostAutoDiscoveryIpmiParams {
	o.SetScheduleDays(scheduleDays)
	return o
}

// SetScheduleDays adds the scheduleDays to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetScheduleDays(scheduleDays *string) {
	o.ScheduleDays = scheduleDays
}

// WithScheduleTime adds the scheduleTime to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithScheduleTime(scheduleTime *string) *PostAutoDiscoveryIpmiParams {
	o.SetScheduleTime(scheduleTime)
	return o
}

// SetScheduleTime adds the scheduleTime to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetScheduleTime(scheduleTime *string) {
	o.ScheduleTime = scheduleTime
}

// WithUpdateModelIfFound adds the updateModelIfFound to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) WithUpdateModelIfFound(updateModelIfFound *string) *PostAutoDiscoveryIpmiParams {
	o.SetUpdateModelIfFound(updateModelIfFound)
	return o
}

// SetUpdateModelIfFound adds the updateModelIfFound to the post auto discovery ipmi params
func (o *PostAutoDiscoveryIpmiParams) SetUpdateModelIfFound(updateModelIfFound *string) {
	o.UpdateModelIfFound = updateModelIfFound
}

// WriteToRequest writes these params to a swagger request
func (o *PostAutoDiscoveryIpmiParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param bmc_password
	frBmcPassword := o.BmcPassword
	fBmcPassword := frBmcPassword
	if fBmcPassword != "" {
		if err := r.SetFormParam("bmc_password", fBmcPassword); err != nil {
			return err
		}
	}

	// form param bmc_user
	frBmcUser := o.BmcUser
	fBmcUser := frBmcUser
	if fBmcUser != "" {
		if err := r.SetFormParam("bmc_user", fBmcUser); err != nil {
			return err
		}
	}

	if o.ClearExistingSchedule != nil {

		// form param clear_existing_schedule
		var frClearExistingSchedule string
		if o.ClearExistingSchedule != nil {
			frClearExistingSchedule = *o.ClearExistingSchedule
		}
		fClearExistingSchedule := frClearExistingSchedule
		if fClearExistingSchedule != "" {
			if err := r.SetFormParam("clear_existing_schedule", fClearExistingSchedule); err != nil {
				return err
			}
		}

	}

	if o.DebugLevel != nil {

		// form param debug_level
		var frDebugLevel string
		if o.DebugLevel != nil {
			frDebugLevel = *o.DebugLevel
		}
		fDebugLevel := frDebugLevel
		if fDebugLevel != "" {
			if err := r.SetFormParam("debug_level", fDebugLevel); err != nil {
				return err
			}
		}

	}

	if o.DiscoveryType != nil {

		// form param discovery_type
		var frDiscoveryType string
		if o.DiscoveryType != nil {
			frDiscoveryType = *o.DiscoveryType
		}
		fDiscoveryType := frDiscoveryType
		if fDiscoveryType != "" {
			if err := r.SetFormParam("discovery_type", fDiscoveryType); err != nil {
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

	// query param hostname_to_use
	qrHostnameToUse := o.HostnameToUse
	qHostnameToUse := qrHostnameToUse
	if qHostnameToUse != "" {
		if err := r.SetQueryParam("hostname_to_use", qHostnameToUse); err != nil {
			return err
		}
	}

	// form param ip_end
	frIPEnd := o.IPEnd
	fIPEnd := frIPEnd
	if fIPEnd != "" {
		if err := r.SetFormParam("ip_end", fIPEnd); err != nil {
			return err
		}
	}

	// form param ip_start
	frIPStart := o.IPStart
	fIPStart := frIPStart
	if fIPStart != "" {
		if err := r.SetFormParam("ip_start", fIPStart); err != nil {
			return err
		}
	}

	// form param name
	frName := o.Name
	fName := frName
	if fName != "" {
		if err := r.SetFormParam("name", fName); err != nil {
			return err
		}
	}

	if o.ObjectCategory != nil {

		// form param object_category
		var frObjectCategory string
		if o.ObjectCategory != nil {
			frObjectCategory = *o.ObjectCategory
		}
		fObjectCategory := frObjectCategory
		if fObjectCategory != "" {
			if err := r.SetFormParam("object_category", fObjectCategory); err != nil {
				return err
			}
		}

	}

	if o.OverwriteObjectCategories != nil {

		// form param overwrite_object_categories
		var frOverwriteObjectCategories string
		if o.OverwriteObjectCategories != nil {
			frOverwriteObjectCategories = *o.OverwriteObjectCategories
		}
		fOverwriteObjectCategories := frOverwriteObjectCategories
		if fOverwriteObjectCategories != "" {
			if err := r.SetFormParam("overwrite_object_categories", fOverwriteObjectCategories); err != nil {
				return err
			}
		}

	}

	if o.RemoteCollectorID != nil {

		// form param remote_collector_id
		var frRemoteCollectorID string
		if o.RemoteCollectorID != nil {
			frRemoteCollectorID = *o.RemoteCollectorID
		}
		fRemoteCollectorID := frRemoteCollectorID
		if fRemoteCollectorID != "" {
			if err := r.SetFormParam("remote_collector_id", fRemoteCollectorID); err != nil {
				return err
			}
		}

	}

	if o.RunAsOperator != nil {

		// form param run_as_operator
		var frRunAsOperator string
		if o.RunAsOperator != nil {
			frRunAsOperator = *o.RunAsOperator
		}
		fRunAsOperator := frRunAsOperator
		if fRunAsOperator != "" {
			if err := r.SetFormParam("run_as_operator", fRunAsOperator); err != nil {
				return err
			}
		}

	}

	if o.ScheduleDays != nil {

		// form param schedule_days
		var frScheduleDays string
		if o.ScheduleDays != nil {
			frScheduleDays = *o.ScheduleDays
		}
		fScheduleDays := frScheduleDays
		if fScheduleDays != "" {
			if err := r.SetFormParam("schedule_days", fScheduleDays); err != nil {
				return err
			}
		}

	}

	if o.ScheduleTime != nil {

		// form param schedule_time
		var frScheduleTime string
		if o.ScheduleTime != nil {
			frScheduleTime = *o.ScheduleTime
		}
		fScheduleTime := frScheduleTime
		if fScheduleTime != "" {
			if err := r.SetFormParam("schedule_time", fScheduleTime); err != nil {
				return err
			}
		}

	}

	if o.UpdateModelIfFound != nil {

		// form param update_model_if_found
		var frUpdateModelIfFound string
		if o.UpdateModelIfFound != nil {
			frUpdateModelIfFound = *o.UpdateModelIfFound
		}
		fUpdateModelIfFound := frUpdateModelIfFound
		if fUpdateModelIfFound != "" {
			if err := r.SetFormParam("update_model_if_found", fUpdateModelIfFound); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}