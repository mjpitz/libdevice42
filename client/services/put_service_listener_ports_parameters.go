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
	"github.com/go-openapi/swag"
)

// NewPutServiceListenerPortsParams creates a new PutServiceListenerPortsParams object
// with the default values initialized.
func NewPutServiceListenerPortsParams() *PutServiceListenerPortsParams {
	var ()
	return &PutServiceListenerPortsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutServiceListenerPortsParamsWithTimeout creates a new PutServiceListenerPortsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutServiceListenerPortsParamsWithTimeout(timeout time.Duration) *PutServiceListenerPortsParams {
	var ()
	return &PutServiceListenerPortsParams{

		timeout: timeout,
	}
}

// NewPutServiceListenerPortsParamsWithContext creates a new PutServiceListenerPortsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutServiceListenerPortsParamsWithContext(ctx context.Context) *PutServiceListenerPortsParams {
	var ()
	return &PutServiceListenerPortsParams{

		Context: ctx,
	}
}

// NewPutServiceListenerPortsParamsWithHTTPClient creates a new PutServiceListenerPortsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutServiceListenerPortsParamsWithHTTPClient(client *http.Client) *PutServiceListenerPortsParams {
	var ()
	return &PutServiceListenerPortsParams{
		HTTPClient: client,
	}
}

/*PutServiceListenerPortsParams contains all the parameters to send to the API endpoint
for the put service listener ports operation typically these are written to a http.Request
*/
type PutServiceListenerPortsParams struct {

	/*DeviceID
	  device id

	*/
	DeviceID *int64
	/*DeviceName
	  the name of the device

	*/
	DeviceName *string
	/*DiscoveredProcess
	  the process name that has a handle to the port

	*/
	DiscoveredProcess *string
	/*DiscoveredProcessDisplayName
	  discovered process display name

	*/
	DiscoveredProcessDisplayName *string
	/*DiscoveredService
	  the name of the discovered service listening on this IP/port

	*/
	DiscoveredService *string
	/*DiscoveredServiceDisplayName
	  discovered service display name

	*/
	DiscoveredServiceDisplayName *string
	/*DiscoveredServiceID
	  device service id

	*/
	DiscoveredServiceID *int64
	/*ID
	  service port id

	*/
	ID int64
	/*ListeningIP
	  the name of the device

	*/
	ListeningIP *string
	/*MappedService
	  the name of the mapped service listening on this IP/port

	*/
	MappedService *string
	/*MappedServiceDisplayName
	  mapped service display name

	*/
	MappedServiceDisplayName *string
	/*Port
	  the listening port on this device

	*/
	Port *int64
	/*Protocol
	  the transport protocol, ie TCP

	*/
	Protocol *string
	/*RemoteIps
	  the comma separated list of remote IPs that are connected to this listening IP/port

	*/
	RemoteIps *string
	/*RemotePortTime
	  remote port time

	*/
	RemotePortTime *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithTimeout(timeout time.Duration) *PutServiceListenerPortsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithContext(ctx context.Context) *PutServiceListenerPortsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithHTTPClient(client *http.Client) *PutServiceListenerPortsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithDeviceID(deviceID *int64) *PutServiceListenerPortsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetDeviceID(deviceID *int64) {
	o.DeviceID = deviceID
}

// WithDeviceName adds the deviceName to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithDeviceName(deviceName *string) *PutServiceListenerPortsParams {
	o.SetDeviceName(deviceName)
	return o
}

// SetDeviceName adds the deviceName to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetDeviceName(deviceName *string) {
	o.DeviceName = deviceName
}

// WithDiscoveredProcess adds the discoveredProcess to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithDiscoveredProcess(discoveredProcess *string) *PutServiceListenerPortsParams {
	o.SetDiscoveredProcess(discoveredProcess)
	return o
}

// SetDiscoveredProcess adds the discoveredProcess to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetDiscoveredProcess(discoveredProcess *string) {
	o.DiscoveredProcess = discoveredProcess
}

// WithDiscoveredProcessDisplayName adds the discoveredProcessDisplayName to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithDiscoveredProcessDisplayName(discoveredProcessDisplayName *string) *PutServiceListenerPortsParams {
	o.SetDiscoveredProcessDisplayName(discoveredProcessDisplayName)
	return o
}

// SetDiscoveredProcessDisplayName adds the discoveredProcessDisplayName to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetDiscoveredProcessDisplayName(discoveredProcessDisplayName *string) {
	o.DiscoveredProcessDisplayName = discoveredProcessDisplayName
}

// WithDiscoveredService adds the discoveredService to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithDiscoveredService(discoveredService *string) *PutServiceListenerPortsParams {
	o.SetDiscoveredService(discoveredService)
	return o
}

// SetDiscoveredService adds the discoveredService to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetDiscoveredService(discoveredService *string) {
	o.DiscoveredService = discoveredService
}

// WithDiscoveredServiceDisplayName adds the discoveredServiceDisplayName to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithDiscoveredServiceDisplayName(discoveredServiceDisplayName *string) *PutServiceListenerPortsParams {
	o.SetDiscoveredServiceDisplayName(discoveredServiceDisplayName)
	return o
}

// SetDiscoveredServiceDisplayName adds the discoveredServiceDisplayName to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetDiscoveredServiceDisplayName(discoveredServiceDisplayName *string) {
	o.DiscoveredServiceDisplayName = discoveredServiceDisplayName
}

// WithDiscoveredServiceID adds the discoveredServiceID to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithDiscoveredServiceID(discoveredServiceID *int64) *PutServiceListenerPortsParams {
	o.SetDiscoveredServiceID(discoveredServiceID)
	return o
}

// SetDiscoveredServiceID adds the discoveredServiceId to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetDiscoveredServiceID(discoveredServiceID *int64) {
	o.DiscoveredServiceID = discoveredServiceID
}

// WithID adds the id to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithID(id int64) *PutServiceListenerPortsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetID(id int64) {
	o.ID = id
}

// WithListeningIP adds the listeningIP to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithListeningIP(listeningIP *string) *PutServiceListenerPortsParams {
	o.SetListeningIP(listeningIP)
	return o
}

// SetListeningIP adds the listeningIp to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetListeningIP(listeningIP *string) {
	o.ListeningIP = listeningIP
}

// WithMappedService adds the mappedService to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithMappedService(mappedService *string) *PutServiceListenerPortsParams {
	o.SetMappedService(mappedService)
	return o
}

// SetMappedService adds the mappedService to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetMappedService(mappedService *string) {
	o.MappedService = mappedService
}

// WithMappedServiceDisplayName adds the mappedServiceDisplayName to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithMappedServiceDisplayName(mappedServiceDisplayName *string) *PutServiceListenerPortsParams {
	o.SetMappedServiceDisplayName(mappedServiceDisplayName)
	return o
}

// SetMappedServiceDisplayName adds the mappedServiceDisplayName to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetMappedServiceDisplayName(mappedServiceDisplayName *string) {
	o.MappedServiceDisplayName = mappedServiceDisplayName
}

// WithPort adds the port to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithPort(port *int64) *PutServiceListenerPortsParams {
	o.SetPort(port)
	return o
}

// SetPort adds the port to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetPort(port *int64) {
	o.Port = port
}

// WithProtocol adds the protocol to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithProtocol(protocol *string) *PutServiceListenerPortsParams {
	o.SetProtocol(protocol)
	return o
}

// SetProtocol adds the protocol to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetProtocol(protocol *string) {
	o.Protocol = protocol
}

// WithRemoteIps adds the remoteIps to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithRemoteIps(remoteIps *string) *PutServiceListenerPortsParams {
	o.SetRemoteIps(remoteIps)
	return o
}

// SetRemoteIps adds the remoteIps to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetRemoteIps(remoteIps *string) {
	o.RemoteIps = remoteIps
}

// WithRemotePortTime adds the remotePortTime to the put service listener ports params
func (o *PutServiceListenerPortsParams) WithRemotePortTime(remotePortTime *string) *PutServiceListenerPortsParams {
	o.SetRemotePortTime(remotePortTime)
	return o
}

// SetRemotePortTime adds the remotePortTime to the put service listener ports params
func (o *PutServiceListenerPortsParams) SetRemotePortTime(remotePortTime *string) {
	o.RemotePortTime = remotePortTime
}

// WriteToRequest writes these params to a swagger request
func (o *PutServiceListenerPortsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeviceID != nil {

		// form param device_id
		var frDeviceID int64
		if o.DeviceID != nil {
			frDeviceID = *o.DeviceID
		}
		fDeviceID := swag.FormatInt64(frDeviceID)
		if fDeviceID != "" {
			if err := r.SetFormParam("device_id", fDeviceID); err != nil {
				return err
			}
		}

	}

	if o.DeviceName != nil {

		// form param device_name
		var frDeviceName string
		if o.DeviceName != nil {
			frDeviceName = *o.DeviceName
		}
		fDeviceName := frDeviceName
		if fDeviceName != "" {
			if err := r.SetFormParam("device_name", fDeviceName); err != nil {
				return err
			}
		}

	}

	if o.DiscoveredProcess != nil {

		// form param discovered_process
		var frDiscoveredProcess string
		if o.DiscoveredProcess != nil {
			frDiscoveredProcess = *o.DiscoveredProcess
		}
		fDiscoveredProcess := frDiscoveredProcess
		if fDiscoveredProcess != "" {
			if err := r.SetFormParam("discovered_process", fDiscoveredProcess); err != nil {
				return err
			}
		}

	}

	if o.DiscoveredProcessDisplayName != nil {

		// form param discovered_process_display_name
		var frDiscoveredProcessDisplayName string
		if o.DiscoveredProcessDisplayName != nil {
			frDiscoveredProcessDisplayName = *o.DiscoveredProcessDisplayName
		}
		fDiscoveredProcessDisplayName := frDiscoveredProcessDisplayName
		if fDiscoveredProcessDisplayName != "" {
			if err := r.SetFormParam("discovered_process_display_name", fDiscoveredProcessDisplayName); err != nil {
				return err
			}
		}

	}

	if o.DiscoveredService != nil {

		// form param discovered_service
		var frDiscoveredService string
		if o.DiscoveredService != nil {
			frDiscoveredService = *o.DiscoveredService
		}
		fDiscoveredService := frDiscoveredService
		if fDiscoveredService != "" {
			if err := r.SetFormParam("discovered_service", fDiscoveredService); err != nil {
				return err
			}
		}

	}

	if o.DiscoveredServiceDisplayName != nil {

		// form param discovered_service_display_name
		var frDiscoveredServiceDisplayName string
		if o.DiscoveredServiceDisplayName != nil {
			frDiscoveredServiceDisplayName = *o.DiscoveredServiceDisplayName
		}
		fDiscoveredServiceDisplayName := frDiscoveredServiceDisplayName
		if fDiscoveredServiceDisplayName != "" {
			if err := r.SetFormParam("discovered_service_display_name", fDiscoveredServiceDisplayName); err != nil {
				return err
			}
		}

	}

	if o.DiscoveredServiceID != nil {

		// form param discovered_service_id
		var frDiscoveredServiceID int64
		if o.DiscoveredServiceID != nil {
			frDiscoveredServiceID = *o.DiscoveredServiceID
		}
		fDiscoveredServiceID := swag.FormatInt64(frDiscoveredServiceID)
		if fDiscoveredServiceID != "" {
			if err := r.SetFormParam("discovered_service_id", fDiscoveredServiceID); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.ListeningIP != nil {

		// form param listening_ip
		var frListeningIP string
		if o.ListeningIP != nil {
			frListeningIP = *o.ListeningIP
		}
		fListeningIP := frListeningIP
		if fListeningIP != "" {
			if err := r.SetFormParam("listening_ip", fListeningIP); err != nil {
				return err
			}
		}

	}

	if o.MappedService != nil {

		// form param mapped_service
		var frMappedService string
		if o.MappedService != nil {
			frMappedService = *o.MappedService
		}
		fMappedService := frMappedService
		if fMappedService != "" {
			if err := r.SetFormParam("mapped_service", fMappedService); err != nil {
				return err
			}
		}

	}

	if o.MappedServiceDisplayName != nil {

		// form param mapped_service_display_name
		var frMappedServiceDisplayName string
		if o.MappedServiceDisplayName != nil {
			frMappedServiceDisplayName = *o.MappedServiceDisplayName
		}
		fMappedServiceDisplayName := frMappedServiceDisplayName
		if fMappedServiceDisplayName != "" {
			if err := r.SetFormParam("mapped_service_display_name", fMappedServiceDisplayName); err != nil {
				return err
			}
		}

	}

	if o.Port != nil {

		// form param port
		var frPort int64
		if o.Port != nil {
			frPort = *o.Port
		}
		fPort := swag.FormatInt64(frPort)
		if fPort != "" {
			if err := r.SetFormParam("port", fPort); err != nil {
				return err
			}
		}

	}

	if o.Protocol != nil {

		// form param protocol
		var frProtocol string
		if o.Protocol != nil {
			frProtocol = *o.Protocol
		}
		fProtocol := frProtocol
		if fProtocol != "" {
			if err := r.SetFormParam("protocol", fProtocol); err != nil {
				return err
			}
		}

	}

	if o.RemoteIps != nil {

		// form param remote_ips
		var frRemoteIps string
		if o.RemoteIps != nil {
			frRemoteIps = *o.RemoteIps
		}
		fRemoteIps := frRemoteIps
		if fRemoteIps != "" {
			if err := r.SetFormParam("remote_ips", fRemoteIps); err != nil {
				return err
			}
		}

	}

	if o.RemotePortTime != nil {

		// form param remote_port_time
		var frRemotePortTime string
		if o.RemotePortTime != nil {
			frRemotePortTime = *o.RemotePortTime
		}
		fRemotePortTime := frRemotePortTime
		if fRemotePortTime != "" {
			if err := r.SetFormParam("remote_port_time", fRemotePortTime); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
