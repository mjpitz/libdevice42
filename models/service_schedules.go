// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceSchedules service schedules
//
// swagger:model Service_schedules
type ServiceSchedules []*ServiceSchedulesItems0

// Validate validates this service schedules
func (m ServiceSchedules) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ServiceSchedulesItems0 service schedules items0
//
// swagger:model ServiceSchedulesItems0
type ServiceSchedulesItems0 struct {

	// arguments
	Arguments string `json:"arguments,omitempty"`

	// at logon
	AtLogon bool `json:"at_logon,omitempty"`

	// at startup
	AtStartup bool `json:"at_startup,omitempty"`

	// caption
	Caption string `json:"caption,omitempty"`

	// day of month
	DayOfMonth string `json:"day_of_month,omitempty"`

	// day of week
	DayOfWeek string `json:"day_of_week,omitempty"`

	// days
	Days string `json:"days,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// device
	Device string `json:"device,omitempty"`

	// device id
	DeviceID int64 `json:"device_id,omitempty"`

	// event based
	EventBased bool `json:"event_based,omitempty"`

	// hours
	Hours string `json:"hours,omitempty"`

	// idle time
	IdleTime bool `json:"idle_time,omitempty"`

	// install date
	InstallDate string `json:"install_Date,omitempty"`

	// minutes
	Minutes string `json:"minutes,omitempty"`

	// month of year
	MonthOfYear string `json:"month_of_year,omitempty"`

	// other trigger
	OtherTrigger bool `json:"other_trigger,omitempty"`

	// other type
	OtherType string `json:"other_type,omitempty"`

	// service name
	ServiceName string `json:"service_name,omitempty"`

	// service schedule id
	ServiceScheduleID int64 `json:"service_schedule_id,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// user
	User string `json:"user,omitempty"`

	// user id
	UserID int64 `json:"user_id,omitempty"`

	// weeks
	Weeks string `json:"weeks,omitempty"`
}

// Validate validates this service schedules items0
func (m *ServiceSchedulesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceSchedulesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceSchedulesItems0) UnmarshalBinary(b []byte) error {
	var res ServiceSchedulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
