// Code generated by go-swagger; DO NOT EDIT.

package services

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

// NewPostServiceDetailsParams creates a new PostServiceDetailsParams object
// with the default values initialized.
func NewPostServiceDetailsParams() *PostServiceDetailsParams {
	var ()
	return &PostServiceDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostServiceDetailsParamsWithTimeout creates a new PostServiceDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostServiceDetailsParamsWithTimeout(timeout time.Duration) *PostServiceDetailsParams {
	var ()
	return &PostServiceDetailsParams{

		timeout: timeout,
	}
}

// NewPostServiceDetailsParamsWithContext creates a new PostServiceDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostServiceDetailsParamsWithContext(ctx context.Context) *PostServiceDetailsParams {
	var ()
	return &PostServiceDetailsParams{

		Context: ctx,
	}
}

// NewPostServiceDetailsParamsWithHTTPClient creates a new PostServiceDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostServiceDetailsParamsWithHTTPClient(client *http.Client) *PostServiceDetailsParams {
	var ()
	return &PostServiceDetailsParams{
		HTTPClient: client,
	}
}

/*PostServiceDetailsParams contains all the parameters to send to the API endpoint
for the post service details operation typically these are written to a http.Request
*/
type PostServiceDetailsParams struct {

	/*Appcomp
	  The application component that depends on this service

	*/
	Appcomp *string
	/*AtLogon
	  only for schedule based services where startmode = ‘Scheduled’

	*/
	AtLogon *string
	/*AtStartup
	  only for schedule based services where startmode = ‘Scheduled’

	*/
	AtStartup *string
	/*Category
	  name of the category

	*/
	Category *string
	/*DayOfMonth
	  only for schedule based services where startmode = ‘Scheduled’

	*/
	DayOfMonth *string
	/*DayOfWeek
	  only for schedule based services where startmode = ‘Scheduled’

	*/
	DayOfWeek *string
	/*Days
	  only for schedule based services where startmode = ‘Scheduled’

	*/
	Days *string
	/*Description*/
	Description *string
	/*Device
	  Device name

	*/
	Device *string
	/*EventBased
	  only for schedule based services where startmode = ‘Scheduled’

	*/
	EventBased *string
	/*Hours
	  only for schedule based services where startmode = ‘Scheduled’

	*/
	Hours *string
	/*IdleTime
	  only for schedule based services where startmode = ‘Scheduled’

	*/
	IdleTime *string
	/*InstallDate
	  The date that the software was installed

	*/
	InstallDate *string
	/*Minutes
	  only for schedule based services where startmode = ‘Scheduled’

	*/
	Minutes *string
	/*MonthOfYear
	  only for schedule based services where startmode = ‘Scheduled’

	*/
	MonthOfYear *string
	/*OtherTrigger
	  only for schedule based services where startmode = ‘Scheduled’

	*/
	OtherTrigger *string
	/*OtherType
	  only for schedule based services where startmode = ‘Scheduled’

	*/
	OtherType *string
	/*ServiceDisplayName
	  The user freindly display name of the service

	*/
	ServiceDisplayName string
	/*ServiceName
	  The executable name of the service

	*/
	ServiceName string
	/*ServiceType
	  could be ignored or tracked. Default is tracked.

	*/
	ServiceType *string
	/*Startmode
	  The start mode of this service - valid values are ‘Automatic’, ‘Manual’, ‘Disabled’ and ‘Unknown’

	*/
	Startmode *string
	/*State
	  The current running state of this service. Valid values are ‘Running’, ‘Started’, ‘Paused’, ‘Stopped’ and ‘Unknown’

	*/
	State *string
	/*Status
	  Instance status (ie, running, stopped)

	*/
	Status *string
	/*User
	  enduser name

	*/
	User *string
	/*Vendor
	  The cloud vendor

	*/
	Vendor *string
	/*Weeks
	  only for schedule based services where startmode = ‘Scheduled’

	*/
	Weeks *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post service details params
func (o *PostServiceDetailsParams) WithTimeout(timeout time.Duration) *PostServiceDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post service details params
func (o *PostServiceDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post service details params
func (o *PostServiceDetailsParams) WithContext(ctx context.Context) *PostServiceDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post service details params
func (o *PostServiceDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post service details params
func (o *PostServiceDetailsParams) WithHTTPClient(client *http.Client) *PostServiceDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post service details params
func (o *PostServiceDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppcomp adds the appcomp to the post service details params
func (o *PostServiceDetailsParams) WithAppcomp(appcomp *string) *PostServiceDetailsParams {
	o.SetAppcomp(appcomp)
	return o
}

// SetAppcomp adds the appcomp to the post service details params
func (o *PostServiceDetailsParams) SetAppcomp(appcomp *string) {
	o.Appcomp = appcomp
}

// WithAtLogon adds the atLogon to the post service details params
func (o *PostServiceDetailsParams) WithAtLogon(atLogon *string) *PostServiceDetailsParams {
	o.SetAtLogon(atLogon)
	return o
}

// SetAtLogon adds the atLogon to the post service details params
func (o *PostServiceDetailsParams) SetAtLogon(atLogon *string) {
	o.AtLogon = atLogon
}

// WithAtStartup adds the atStartup to the post service details params
func (o *PostServiceDetailsParams) WithAtStartup(atStartup *string) *PostServiceDetailsParams {
	o.SetAtStartup(atStartup)
	return o
}

// SetAtStartup adds the atStartup to the post service details params
func (o *PostServiceDetailsParams) SetAtStartup(atStartup *string) {
	o.AtStartup = atStartup
}

// WithCategory adds the category to the post service details params
func (o *PostServiceDetailsParams) WithCategory(category *string) *PostServiceDetailsParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the post service details params
func (o *PostServiceDetailsParams) SetCategory(category *string) {
	o.Category = category
}

// WithDayOfMonth adds the dayOfMonth to the post service details params
func (o *PostServiceDetailsParams) WithDayOfMonth(dayOfMonth *string) *PostServiceDetailsParams {
	o.SetDayOfMonth(dayOfMonth)
	return o
}

// SetDayOfMonth adds the dayOfMonth to the post service details params
func (o *PostServiceDetailsParams) SetDayOfMonth(dayOfMonth *string) {
	o.DayOfMonth = dayOfMonth
}

// WithDayOfWeek adds the dayOfWeek to the post service details params
func (o *PostServiceDetailsParams) WithDayOfWeek(dayOfWeek *string) *PostServiceDetailsParams {
	o.SetDayOfWeek(dayOfWeek)
	return o
}

// SetDayOfWeek adds the dayOfWeek to the post service details params
func (o *PostServiceDetailsParams) SetDayOfWeek(dayOfWeek *string) {
	o.DayOfWeek = dayOfWeek
}

// WithDays adds the days to the post service details params
func (o *PostServiceDetailsParams) WithDays(days *string) *PostServiceDetailsParams {
	o.SetDays(days)
	return o
}

// SetDays adds the days to the post service details params
func (o *PostServiceDetailsParams) SetDays(days *string) {
	o.Days = days
}

// WithDescription adds the description to the post service details params
func (o *PostServiceDetailsParams) WithDescription(description *string) *PostServiceDetailsParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the post service details params
func (o *PostServiceDetailsParams) SetDescription(description *string) {
	o.Description = description
}

// WithDevice adds the device to the post service details params
func (o *PostServiceDetailsParams) WithDevice(device *string) *PostServiceDetailsParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the post service details params
func (o *PostServiceDetailsParams) SetDevice(device *string) {
	o.Device = device
}

// WithEventBased adds the eventBased to the post service details params
func (o *PostServiceDetailsParams) WithEventBased(eventBased *string) *PostServiceDetailsParams {
	o.SetEventBased(eventBased)
	return o
}

// SetEventBased adds the eventBased to the post service details params
func (o *PostServiceDetailsParams) SetEventBased(eventBased *string) {
	o.EventBased = eventBased
}

// WithHours adds the hours to the post service details params
func (o *PostServiceDetailsParams) WithHours(hours *string) *PostServiceDetailsParams {
	o.SetHours(hours)
	return o
}

// SetHours adds the hours to the post service details params
func (o *PostServiceDetailsParams) SetHours(hours *string) {
	o.Hours = hours
}

// WithIdleTime adds the idleTime to the post service details params
func (o *PostServiceDetailsParams) WithIdleTime(idleTime *string) *PostServiceDetailsParams {
	o.SetIdleTime(idleTime)
	return o
}

// SetIdleTime adds the idleTime to the post service details params
func (o *PostServiceDetailsParams) SetIdleTime(idleTime *string) {
	o.IdleTime = idleTime
}

// WithInstallDate adds the installDate to the post service details params
func (o *PostServiceDetailsParams) WithInstallDate(installDate *string) *PostServiceDetailsParams {
	o.SetInstallDate(installDate)
	return o
}

// SetInstallDate adds the installDate to the post service details params
func (o *PostServiceDetailsParams) SetInstallDate(installDate *string) {
	o.InstallDate = installDate
}

// WithMinutes adds the minutes to the post service details params
func (o *PostServiceDetailsParams) WithMinutes(minutes *string) *PostServiceDetailsParams {
	o.SetMinutes(minutes)
	return o
}

// SetMinutes adds the minutes to the post service details params
func (o *PostServiceDetailsParams) SetMinutes(minutes *string) {
	o.Minutes = minutes
}

// WithMonthOfYear adds the monthOfYear to the post service details params
func (o *PostServiceDetailsParams) WithMonthOfYear(monthOfYear *string) *PostServiceDetailsParams {
	o.SetMonthOfYear(monthOfYear)
	return o
}

// SetMonthOfYear adds the monthOfYear to the post service details params
func (o *PostServiceDetailsParams) SetMonthOfYear(monthOfYear *string) {
	o.MonthOfYear = monthOfYear
}

// WithOtherTrigger adds the otherTrigger to the post service details params
func (o *PostServiceDetailsParams) WithOtherTrigger(otherTrigger *string) *PostServiceDetailsParams {
	o.SetOtherTrigger(otherTrigger)
	return o
}

// SetOtherTrigger adds the otherTrigger to the post service details params
func (o *PostServiceDetailsParams) SetOtherTrigger(otherTrigger *string) {
	o.OtherTrigger = otherTrigger
}

// WithOtherType adds the otherType to the post service details params
func (o *PostServiceDetailsParams) WithOtherType(otherType *string) *PostServiceDetailsParams {
	o.SetOtherType(otherType)
	return o
}

// SetOtherType adds the otherType to the post service details params
func (o *PostServiceDetailsParams) SetOtherType(otherType *string) {
	o.OtherType = otherType
}

// WithServiceDisplayName adds the serviceDisplayName to the post service details params
func (o *PostServiceDetailsParams) WithServiceDisplayName(serviceDisplayName string) *PostServiceDetailsParams {
	o.SetServiceDisplayName(serviceDisplayName)
	return o
}

// SetServiceDisplayName adds the serviceDisplayName to the post service details params
func (o *PostServiceDetailsParams) SetServiceDisplayName(serviceDisplayName string) {
	o.ServiceDisplayName = serviceDisplayName
}

// WithServiceName adds the serviceName to the post service details params
func (o *PostServiceDetailsParams) WithServiceName(serviceName string) *PostServiceDetailsParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post service details params
func (o *PostServiceDetailsParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WithServiceType adds the serviceType to the post service details params
func (o *PostServiceDetailsParams) WithServiceType(serviceType *string) *PostServiceDetailsParams {
	o.SetServiceType(serviceType)
	return o
}

// SetServiceType adds the serviceType to the post service details params
func (o *PostServiceDetailsParams) SetServiceType(serviceType *string) {
	o.ServiceType = serviceType
}

// WithStartmode adds the startmode to the post service details params
func (o *PostServiceDetailsParams) WithStartmode(startmode *string) *PostServiceDetailsParams {
	o.SetStartmode(startmode)
	return o
}

// SetStartmode adds the startmode to the post service details params
func (o *PostServiceDetailsParams) SetStartmode(startmode *string) {
	o.Startmode = startmode
}

// WithState adds the state to the post service details params
func (o *PostServiceDetailsParams) WithState(state *string) *PostServiceDetailsParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the post service details params
func (o *PostServiceDetailsParams) SetState(state *string) {
	o.State = state
}

// WithStatus adds the status to the post service details params
func (o *PostServiceDetailsParams) WithStatus(status *string) *PostServiceDetailsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the post service details params
func (o *PostServiceDetailsParams) SetStatus(status *string) {
	o.Status = status
}

// WithUser adds the user to the post service details params
func (o *PostServiceDetailsParams) WithUser(user *string) *PostServiceDetailsParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the post service details params
func (o *PostServiceDetailsParams) SetUser(user *string) {
	o.User = user
}

// WithVendor adds the vendor to the post service details params
func (o *PostServiceDetailsParams) WithVendor(vendor *string) *PostServiceDetailsParams {
	o.SetVendor(vendor)
	return o
}

// SetVendor adds the vendor to the post service details params
func (o *PostServiceDetailsParams) SetVendor(vendor *string) {
	o.Vendor = vendor
}

// WithWeeks adds the weeks to the post service details params
func (o *PostServiceDetailsParams) WithWeeks(weeks *string) *PostServiceDetailsParams {
	o.SetWeeks(weeks)
	return o
}

// SetWeeks adds the weeks to the post service details params
func (o *PostServiceDetailsParams) SetWeeks(weeks *string) {
	o.Weeks = weeks
}

// WriteToRequest writes these params to a swagger request
func (o *PostServiceDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Appcomp != nil {

		// query param appcomp
		var qrAppcomp string
		if o.Appcomp != nil {
			qrAppcomp = *o.Appcomp
		}
		qAppcomp := qrAppcomp
		if qAppcomp != "" {
			if err := r.SetQueryParam("appcomp", qAppcomp); err != nil {
				return err
			}
		}

	}

	if o.AtLogon != nil {

		// query param at_logon
		var qrAtLogon string
		if o.AtLogon != nil {
			qrAtLogon = *o.AtLogon
		}
		qAtLogon := qrAtLogon
		if qAtLogon != "" {
			if err := r.SetQueryParam("at_logon", qAtLogon); err != nil {
				return err
			}
		}

	}

	if o.AtStartup != nil {

		// query param at_startup
		var qrAtStartup string
		if o.AtStartup != nil {
			qrAtStartup = *o.AtStartup
		}
		qAtStartup := qrAtStartup
		if qAtStartup != "" {
			if err := r.SetQueryParam("at_startup", qAtStartup); err != nil {
				return err
			}
		}

	}

	if o.Category != nil {

		// query param category
		var qrCategory string
		if o.Category != nil {
			qrCategory = *o.Category
		}
		qCategory := qrCategory
		if qCategory != "" {
			if err := r.SetQueryParam("category", qCategory); err != nil {
				return err
			}
		}

	}

	if o.DayOfMonth != nil {

		// query param day_of_month
		var qrDayOfMonth string
		if o.DayOfMonth != nil {
			qrDayOfMonth = *o.DayOfMonth
		}
		qDayOfMonth := qrDayOfMonth
		if qDayOfMonth != "" {
			if err := r.SetQueryParam("day_of_month", qDayOfMonth); err != nil {
				return err
			}
		}

	}

	if o.DayOfWeek != nil {

		// query param day_of_week
		var qrDayOfWeek string
		if o.DayOfWeek != nil {
			qrDayOfWeek = *o.DayOfWeek
		}
		qDayOfWeek := qrDayOfWeek
		if qDayOfWeek != "" {
			if err := r.SetQueryParam("day_of_week", qDayOfWeek); err != nil {
				return err
			}
		}

	}

	if o.Days != nil {

		// query param days
		var qrDays string
		if o.Days != nil {
			qrDays = *o.Days
		}
		qDays := qrDays
		if qDays != "" {
			if err := r.SetQueryParam("days", qDays); err != nil {
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

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.EventBased != nil {

		// query param event_based
		var qrEventBased string
		if o.EventBased != nil {
			qrEventBased = *o.EventBased
		}
		qEventBased := qrEventBased
		if qEventBased != "" {
			if err := r.SetQueryParam("event_based", qEventBased); err != nil {
				return err
			}
		}

	}

	if o.Hours != nil {

		// query param hours
		var qrHours string
		if o.Hours != nil {
			qrHours = *o.Hours
		}
		qHours := qrHours
		if qHours != "" {
			if err := r.SetQueryParam("hours", qHours); err != nil {
				return err
			}
		}

	}

	if o.IdleTime != nil {

		// query param idle_time
		var qrIdleTime string
		if o.IdleTime != nil {
			qrIdleTime = *o.IdleTime
		}
		qIdleTime := qrIdleTime
		if qIdleTime != "" {
			if err := r.SetQueryParam("idle_time", qIdleTime); err != nil {
				return err
			}
		}

	}

	if o.InstallDate != nil {

		// form param install_date
		var frInstallDate string
		if o.InstallDate != nil {
			frInstallDate = *o.InstallDate
		}
		fInstallDate := frInstallDate
		if fInstallDate != "" {
			if err := r.SetFormParam("install_date", fInstallDate); err != nil {
				return err
			}
		}

	}

	if o.Minutes != nil {

		// query param minutes
		var qrMinutes string
		if o.Minutes != nil {
			qrMinutes = *o.Minutes
		}
		qMinutes := qrMinutes
		if qMinutes != "" {
			if err := r.SetQueryParam("minutes", qMinutes); err != nil {
				return err
			}
		}

	}

	if o.MonthOfYear != nil {

		// query param month_of_year
		var qrMonthOfYear string
		if o.MonthOfYear != nil {
			qrMonthOfYear = *o.MonthOfYear
		}
		qMonthOfYear := qrMonthOfYear
		if qMonthOfYear != "" {
			if err := r.SetQueryParam("month_of_year", qMonthOfYear); err != nil {
				return err
			}
		}

	}

	if o.OtherTrigger != nil {

		// query param other_trigger
		var qrOtherTrigger string
		if o.OtherTrigger != nil {
			qrOtherTrigger = *o.OtherTrigger
		}
		qOtherTrigger := qrOtherTrigger
		if qOtherTrigger != "" {
			if err := r.SetQueryParam("other_trigger", qOtherTrigger); err != nil {
				return err
			}
		}

	}

	if o.OtherType != nil {

		// query param other_type
		var qrOtherType string
		if o.OtherType != nil {
			qrOtherType = *o.OtherType
		}
		qOtherType := qrOtherType
		if qOtherType != "" {
			if err := r.SetQueryParam("other_type", qOtherType); err != nil {
				return err
			}
		}

	}

	// query param service_display_name
	qrServiceDisplayName := o.ServiceDisplayName
	qServiceDisplayName := qrServiceDisplayName
	if qServiceDisplayName != "" {
		if err := r.SetQueryParam("service_display_name", qServiceDisplayName); err != nil {
			return err
		}
	}

	// query param service_name
	qrServiceName := o.ServiceName
	qServiceName := qrServiceName
	if qServiceName != "" {
		if err := r.SetQueryParam("service_name", qServiceName); err != nil {
			return err
		}
	}

	if o.ServiceType != nil {

		// query param service_type
		var qrServiceType string
		if o.ServiceType != nil {
			qrServiceType = *o.ServiceType
		}
		qServiceType := qrServiceType
		if qServiceType != "" {
			if err := r.SetQueryParam("service_type", qServiceType); err != nil {
				return err
			}
		}

	}

	if o.Startmode != nil {

		// query param startmode
		var qrStartmode string
		if o.Startmode != nil {
			qrStartmode = *o.Startmode
		}
		qStartmode := qrStartmode
		if qStartmode != "" {
			if err := r.SetQueryParam("startmode", qStartmode); err != nil {
				return err
			}
		}

	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// form param status
		var frStatus string
		if o.Status != nil {
			frStatus = *o.Status
		}
		fStatus := frStatus
		if fStatus != "" {
			if err := r.SetFormParam("status", fStatus); err != nil {
				return err
			}
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

	if o.Vendor != nil {

		// form param vendor
		var frVendor string
		if o.Vendor != nil {
			frVendor = *o.Vendor
		}
		fVendor := frVendor
		if fVendor != "" {
			if err := r.SetFormParam("vendor", fVendor); err != nil {
				return err
			}
		}

	}

	if o.Weeks != nil {

		// query param weeks
		var qrWeeks string
		if o.Weeks != nil {
			qrWeeks = *o.Weeks
		}
		qWeeks := qrWeeks
		if qWeeks != "" {
			if err := r.SetQueryParam("weeks", qWeeks); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
