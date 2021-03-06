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
	"github.com/go-openapi/swag"
)

// NewPostAutoDiscoveryNetworksParams creates a new PostAutoDiscoveryNetworksParams object
// with the default values initialized.
func NewPostAutoDiscoveryNetworksParams() *PostAutoDiscoveryNetworksParams {
	var (
		snmpVersionDefault = string("v2c")
	)
	return &PostAutoDiscoveryNetworksParams{
		SnmpVersion: &snmpVersionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAutoDiscoveryNetworksParamsWithTimeout creates a new PostAutoDiscoveryNetworksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAutoDiscoveryNetworksParamsWithTimeout(timeout time.Duration) *PostAutoDiscoveryNetworksParams {
	var (
		snmpVersionDefault = string("v2c")
	)
	return &PostAutoDiscoveryNetworksParams{
		SnmpVersion: &snmpVersionDefault,

		timeout: timeout,
	}
}

// NewPostAutoDiscoveryNetworksParamsWithContext creates a new PostAutoDiscoveryNetworksParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAutoDiscoveryNetworksParamsWithContext(ctx context.Context) *PostAutoDiscoveryNetworksParams {
	var (
		snmpVersionDefault = string("v2c")
	)
	return &PostAutoDiscoveryNetworksParams{
		SnmpVersion: &snmpVersionDefault,

		Context: ctx,
	}
}

// NewPostAutoDiscoveryNetworksParamsWithHTTPClient creates a new PostAutoDiscoveryNetworksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAutoDiscoveryNetworksParamsWithHTTPClient(client *http.Client) *PostAutoDiscoveryNetworksParams {
	var (
		snmpVersionDefault = string("v2c")
	)
	return &PostAutoDiscoveryNetworksParams{
		SnmpVersion: &snmpVersionDefault,
		HTTPClient:  client,
	}
}

/*PostAutoDiscoveryNetworksParams contains all the parameters to send to the API endpoint
for the post auto discovery networks operation typically these are written to a http.Request
*/
type PostAutoDiscoveryNetworksParams struct {

	/*AutodiscoverCdpDevices
	  yes to enable CDP/LLDP (added in v8.3.2)

	*/
	AutodiscoverCdpDevices *string
	/*ClearExistingSchedule*/
	ClearExistingSchedule *string
	/*DebugLevel*/
	DebugLevel *string
	/*DeleteOlderMacAssociationAfter
	  number of days (added in v10.4.0)

	*/
	DeleteOlderMacAssociationAfter *string
	/*DeleteSwitchPortNotFound
	  yes or no to delete switch ports not found (added in v10.4.0)

	*/
	DeleteSwitchPortNotFound *string
	/*DiscoverServices*/
	DiscoverServices *string
	/*EndIPAddress
	  End IP address

	*/
	EndIPAddress *string
	/*GetAllSwitchPorts
	  yes or no to get all switch ports (added in v10.4.0)

	*/
	GetAllSwitchPorts *string
	/*Groups
	  name of one or more groups separated by commas

	*/
	Groups *string
	/*HostnamePrecedence*/
	HostnamePrecedence *string
	/*Ipaddress
	  IP address. Required if new

	*/
	Ipaddress *string
	/*MaskBits
	  mask bits (integer only)

	*/
	MaskBits *int64
	/*Name
	  name of the job

	*/
	Name string
	/*ObjectCategory
	  name of subnet category for discovered subnets

	*/
	ObjectCategory *string
	/*PortNamePrefixToIgnoreMacs*/
	PortNamePrefixToIgnoreMacs *string
	/*ScheduleDays
	  Comma separated days of week, where Monday = 0. e.g. 0,1,2 wil set the job for Mon, Tue and Wed. For multiple schedules, separate with a slash (/).

	*/
	ScheduleDays *string
	/*ScheduleTime
	  Time in HH:MM format if you want to schedule the job. Note: Must be formatted as text NOT date. For multiple schedules, separate with a slash (/).

	*/
	ScheduleTime *string
	/*ServiceLevel
	  Must already exist

	*/
	ServiceLevel *string
	/*SkipVlanIndexing
	  yes or no (added in v10.4.0)

	*/
	SkipVlanIndexing *string
	/*SnmpPort
	  snmp port (integer only) (added in v10.4.0)

	*/
	SnmpPort *int64
	/*SnmpString
	  required, if new

	*/
	SnmpString *string
	/*SnmpStringID
	  The id of the password for the community string

	*/
	SnmpStringID *string
	/*SnmpStringIds
	  Can be comma separated list of community string IDs to use multiple community strings

	*/
	SnmpStringIds *string
	/*SnmpStrings
	  Can be comma separated list of community strings to use multiple community strings

	*/
	SnmpStrings *string
	/*SnmpVersion*/
	SnmpVersion *string
	/*Snmpv3AuthMode*/
	Snmpv3AuthMode *string
	/*Snmpv3AuthPassword
	  password (added in v10.4.0)

	*/
	Snmpv3AuthPassword *string
	/*Snmpv3AuthPasswordID
	  The id of the password for the auth password

	*/
	Snmpv3AuthPasswordID *string
	/*Snmpv3AuthProtocol*/
	Snmpv3AuthProtocol *string
	/*Snmpv3Context*/
	Snmpv3Context *string
	/*Snmpv3PrivacyProtocol*/
	Snmpv3PrivacyProtocol *string
	/*Snmpv3PrivacyProtocolPassword
	  password (added in v10.4.0)

	*/
	Snmpv3PrivacyProtocolPassword *string
	/*Snmpv3PrivacyProtocolPasswordID
	  The id of the password for the privacy protocol password

	*/
	Snmpv3PrivacyProtocolPasswordID *string
	/*Snmpv3User
	  name of snmp v3 user (added in v10.4.0)

	*/
	Snmpv3User *string
	/*StripDomainName*/
	StripDomainName *string
	/*UseNameAliasPortDescr
	  yes to use alias for port description during discovery (added in v8.3.2)

	*/
	UseNameAliasPortDescr *string
	/*VlansToIgnore
	  list of vlan ids to ignore separated by commas (added in v10.4.0)

	*/
	VlansToIgnore *string
	/*Vrfgroup
	  name of vrf group for discovered subnets (added in v10.4.0)

	*/
	Vrfgroup *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithTimeout(timeout time.Duration) *PostAutoDiscoveryNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithContext(ctx context.Context) *PostAutoDiscoveryNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithHTTPClient(client *http.Client) *PostAutoDiscoveryNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAutodiscoverCdpDevices adds the autodiscoverCdpDevices to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithAutodiscoverCdpDevices(autodiscoverCdpDevices *string) *PostAutoDiscoveryNetworksParams {
	o.SetAutodiscoverCdpDevices(autodiscoverCdpDevices)
	return o
}

// SetAutodiscoverCdpDevices adds the autodiscoverCdpDevices to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetAutodiscoverCdpDevices(autodiscoverCdpDevices *string) {
	o.AutodiscoverCdpDevices = autodiscoverCdpDevices
}

// WithClearExistingSchedule adds the clearExistingSchedule to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithClearExistingSchedule(clearExistingSchedule *string) *PostAutoDiscoveryNetworksParams {
	o.SetClearExistingSchedule(clearExistingSchedule)
	return o
}

// SetClearExistingSchedule adds the clearExistingSchedule to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetClearExistingSchedule(clearExistingSchedule *string) {
	o.ClearExistingSchedule = clearExistingSchedule
}

// WithDebugLevel adds the debugLevel to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithDebugLevel(debugLevel *string) *PostAutoDiscoveryNetworksParams {
	o.SetDebugLevel(debugLevel)
	return o
}

// SetDebugLevel adds the debugLevel to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetDebugLevel(debugLevel *string) {
	o.DebugLevel = debugLevel
}

// WithDeleteOlderMacAssociationAfter adds the deleteOlderMacAssociationAfter to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithDeleteOlderMacAssociationAfter(deleteOlderMacAssociationAfter *string) *PostAutoDiscoveryNetworksParams {
	o.SetDeleteOlderMacAssociationAfter(deleteOlderMacAssociationAfter)
	return o
}

// SetDeleteOlderMacAssociationAfter adds the deleteOlderMacAssociationAfter to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetDeleteOlderMacAssociationAfter(deleteOlderMacAssociationAfter *string) {
	o.DeleteOlderMacAssociationAfter = deleteOlderMacAssociationAfter
}

// WithDeleteSwitchPortNotFound adds the deleteSwitchPortNotFound to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithDeleteSwitchPortNotFound(deleteSwitchPortNotFound *string) *PostAutoDiscoveryNetworksParams {
	o.SetDeleteSwitchPortNotFound(deleteSwitchPortNotFound)
	return o
}

// SetDeleteSwitchPortNotFound adds the deleteSwitchPortNotFound to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetDeleteSwitchPortNotFound(deleteSwitchPortNotFound *string) {
	o.DeleteSwitchPortNotFound = deleteSwitchPortNotFound
}

// WithDiscoverServices adds the discoverServices to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithDiscoverServices(discoverServices *string) *PostAutoDiscoveryNetworksParams {
	o.SetDiscoverServices(discoverServices)
	return o
}

// SetDiscoverServices adds the discoverServices to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetDiscoverServices(discoverServices *string) {
	o.DiscoverServices = discoverServices
}

// WithEndIPAddress adds the endIPAddress to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithEndIPAddress(endIPAddress *string) *PostAutoDiscoveryNetworksParams {
	o.SetEndIPAddress(endIPAddress)
	return o
}

// SetEndIPAddress adds the endIpAddress to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetEndIPAddress(endIPAddress *string) {
	o.EndIPAddress = endIPAddress
}

// WithGetAllSwitchPorts adds the getAllSwitchPorts to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithGetAllSwitchPorts(getAllSwitchPorts *string) *PostAutoDiscoveryNetworksParams {
	o.SetGetAllSwitchPorts(getAllSwitchPorts)
	return o
}

// SetGetAllSwitchPorts adds the getAllSwitchPorts to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetGetAllSwitchPorts(getAllSwitchPorts *string) {
	o.GetAllSwitchPorts = getAllSwitchPorts
}

// WithGroups adds the groups to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithGroups(groups *string) *PostAutoDiscoveryNetworksParams {
	o.SetGroups(groups)
	return o
}

// SetGroups adds the groups to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetGroups(groups *string) {
	o.Groups = groups
}

// WithHostnamePrecedence adds the hostnamePrecedence to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithHostnamePrecedence(hostnamePrecedence *string) *PostAutoDiscoveryNetworksParams {
	o.SetHostnamePrecedence(hostnamePrecedence)
	return o
}

// SetHostnamePrecedence adds the hostnamePrecedence to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetHostnamePrecedence(hostnamePrecedence *string) {
	o.HostnamePrecedence = hostnamePrecedence
}

// WithIpaddress adds the ipaddress to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithIpaddress(ipaddress *string) *PostAutoDiscoveryNetworksParams {
	o.SetIpaddress(ipaddress)
	return o
}

// SetIpaddress adds the ipaddress to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetIpaddress(ipaddress *string) {
	o.Ipaddress = ipaddress
}

// WithMaskBits adds the maskBits to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithMaskBits(maskBits *int64) *PostAutoDiscoveryNetworksParams {
	o.SetMaskBits(maskBits)
	return o
}

// SetMaskBits adds the maskBits to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetMaskBits(maskBits *int64) {
	o.MaskBits = maskBits
}

// WithName adds the name to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithName(name string) *PostAutoDiscoveryNetworksParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetName(name string) {
	o.Name = name
}

// WithObjectCategory adds the objectCategory to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithObjectCategory(objectCategory *string) *PostAutoDiscoveryNetworksParams {
	o.SetObjectCategory(objectCategory)
	return o
}

// SetObjectCategory adds the objectCategory to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetObjectCategory(objectCategory *string) {
	o.ObjectCategory = objectCategory
}

// WithPortNamePrefixToIgnoreMacs adds the portNamePrefixToIgnoreMacs to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithPortNamePrefixToIgnoreMacs(portNamePrefixToIgnoreMacs *string) *PostAutoDiscoveryNetworksParams {
	o.SetPortNamePrefixToIgnoreMacs(portNamePrefixToIgnoreMacs)
	return o
}

// SetPortNamePrefixToIgnoreMacs adds the portNamePrefixToIgnoreMacs to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetPortNamePrefixToIgnoreMacs(portNamePrefixToIgnoreMacs *string) {
	o.PortNamePrefixToIgnoreMacs = portNamePrefixToIgnoreMacs
}

// WithScheduleDays adds the scheduleDays to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithScheduleDays(scheduleDays *string) *PostAutoDiscoveryNetworksParams {
	o.SetScheduleDays(scheduleDays)
	return o
}

// SetScheduleDays adds the scheduleDays to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetScheduleDays(scheduleDays *string) {
	o.ScheduleDays = scheduleDays
}

// WithScheduleTime adds the scheduleTime to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithScheduleTime(scheduleTime *string) *PostAutoDiscoveryNetworksParams {
	o.SetScheduleTime(scheduleTime)
	return o
}

// SetScheduleTime adds the scheduleTime to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetScheduleTime(scheduleTime *string) {
	o.ScheduleTime = scheduleTime
}

// WithServiceLevel adds the serviceLevel to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithServiceLevel(serviceLevel *string) *PostAutoDiscoveryNetworksParams {
	o.SetServiceLevel(serviceLevel)
	return o
}

// SetServiceLevel adds the serviceLevel to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetServiceLevel(serviceLevel *string) {
	o.ServiceLevel = serviceLevel
}

// WithSkipVlanIndexing adds the skipVlanIndexing to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSkipVlanIndexing(skipVlanIndexing *string) *PostAutoDiscoveryNetworksParams {
	o.SetSkipVlanIndexing(skipVlanIndexing)
	return o
}

// SetSkipVlanIndexing adds the skipVlanIndexing to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSkipVlanIndexing(skipVlanIndexing *string) {
	o.SkipVlanIndexing = skipVlanIndexing
}

// WithSnmpPort adds the snmpPort to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSnmpPort(snmpPort *int64) *PostAutoDiscoveryNetworksParams {
	o.SetSnmpPort(snmpPort)
	return o
}

// SetSnmpPort adds the snmpPort to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSnmpPort(snmpPort *int64) {
	o.SnmpPort = snmpPort
}

// WithSnmpString adds the snmpString to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSnmpString(snmpString *string) *PostAutoDiscoveryNetworksParams {
	o.SetSnmpString(snmpString)
	return o
}

// SetSnmpString adds the snmpString to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSnmpString(snmpString *string) {
	o.SnmpString = snmpString
}

// WithSnmpStringID adds the snmpStringID to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSnmpStringID(snmpStringID *string) *PostAutoDiscoveryNetworksParams {
	o.SetSnmpStringID(snmpStringID)
	return o
}

// SetSnmpStringID adds the snmpStringId to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSnmpStringID(snmpStringID *string) {
	o.SnmpStringID = snmpStringID
}

// WithSnmpStringIds adds the snmpStringIds to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSnmpStringIds(snmpStringIds *string) *PostAutoDiscoveryNetworksParams {
	o.SetSnmpStringIds(snmpStringIds)
	return o
}

// SetSnmpStringIds adds the snmpStringIds to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSnmpStringIds(snmpStringIds *string) {
	o.SnmpStringIds = snmpStringIds
}

// WithSnmpStrings adds the snmpStrings to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSnmpStrings(snmpStrings *string) *PostAutoDiscoveryNetworksParams {
	o.SetSnmpStrings(snmpStrings)
	return o
}

// SetSnmpStrings adds the snmpStrings to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSnmpStrings(snmpStrings *string) {
	o.SnmpStrings = snmpStrings
}

// WithSnmpVersion adds the snmpVersion to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSnmpVersion(snmpVersion *string) *PostAutoDiscoveryNetworksParams {
	o.SetSnmpVersion(snmpVersion)
	return o
}

// SetSnmpVersion adds the snmpVersion to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSnmpVersion(snmpVersion *string) {
	o.SnmpVersion = snmpVersion
}

// WithSnmpv3AuthMode adds the snmpv3AuthMode to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSnmpv3AuthMode(snmpv3AuthMode *string) *PostAutoDiscoveryNetworksParams {
	o.SetSnmpv3AuthMode(snmpv3AuthMode)
	return o
}

// SetSnmpv3AuthMode adds the snmpv3AuthMode to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSnmpv3AuthMode(snmpv3AuthMode *string) {
	o.Snmpv3AuthMode = snmpv3AuthMode
}

// WithSnmpv3AuthPassword adds the snmpv3AuthPassword to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSnmpv3AuthPassword(snmpv3AuthPassword *string) *PostAutoDiscoveryNetworksParams {
	o.SetSnmpv3AuthPassword(snmpv3AuthPassword)
	return o
}

// SetSnmpv3AuthPassword adds the snmpv3AuthPassword to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSnmpv3AuthPassword(snmpv3AuthPassword *string) {
	o.Snmpv3AuthPassword = snmpv3AuthPassword
}

// WithSnmpv3AuthPasswordID adds the snmpv3AuthPasswordID to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSnmpv3AuthPasswordID(snmpv3AuthPasswordID *string) *PostAutoDiscoveryNetworksParams {
	o.SetSnmpv3AuthPasswordID(snmpv3AuthPasswordID)
	return o
}

// SetSnmpv3AuthPasswordID adds the snmpv3AuthPasswordId to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSnmpv3AuthPasswordID(snmpv3AuthPasswordID *string) {
	o.Snmpv3AuthPasswordID = snmpv3AuthPasswordID
}

// WithSnmpv3AuthProtocol adds the snmpv3AuthProtocol to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSnmpv3AuthProtocol(snmpv3AuthProtocol *string) *PostAutoDiscoveryNetworksParams {
	o.SetSnmpv3AuthProtocol(snmpv3AuthProtocol)
	return o
}

// SetSnmpv3AuthProtocol adds the snmpv3AuthProtocol to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSnmpv3AuthProtocol(snmpv3AuthProtocol *string) {
	o.Snmpv3AuthProtocol = snmpv3AuthProtocol
}

// WithSnmpv3Context adds the snmpv3Context to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSnmpv3Context(snmpv3Context *string) *PostAutoDiscoveryNetworksParams {
	o.SetSnmpv3Context(snmpv3Context)
	return o
}

// SetSnmpv3Context adds the snmpv3Context to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSnmpv3Context(snmpv3Context *string) {
	o.Snmpv3Context = snmpv3Context
}

// WithSnmpv3PrivacyProtocol adds the snmpv3PrivacyProtocol to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSnmpv3PrivacyProtocol(snmpv3PrivacyProtocol *string) *PostAutoDiscoveryNetworksParams {
	o.SetSnmpv3PrivacyProtocol(snmpv3PrivacyProtocol)
	return o
}

// SetSnmpv3PrivacyProtocol adds the snmpv3PrivacyProtocol to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSnmpv3PrivacyProtocol(snmpv3PrivacyProtocol *string) {
	o.Snmpv3PrivacyProtocol = snmpv3PrivacyProtocol
}

// WithSnmpv3PrivacyProtocolPassword adds the snmpv3PrivacyProtocolPassword to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSnmpv3PrivacyProtocolPassword(snmpv3PrivacyProtocolPassword *string) *PostAutoDiscoveryNetworksParams {
	o.SetSnmpv3PrivacyProtocolPassword(snmpv3PrivacyProtocolPassword)
	return o
}

// SetSnmpv3PrivacyProtocolPassword adds the snmpv3PrivacyProtocolPassword to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSnmpv3PrivacyProtocolPassword(snmpv3PrivacyProtocolPassword *string) {
	o.Snmpv3PrivacyProtocolPassword = snmpv3PrivacyProtocolPassword
}

// WithSnmpv3PrivacyProtocolPasswordID adds the snmpv3PrivacyProtocolPasswordID to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSnmpv3PrivacyProtocolPasswordID(snmpv3PrivacyProtocolPasswordID *string) *PostAutoDiscoveryNetworksParams {
	o.SetSnmpv3PrivacyProtocolPasswordID(snmpv3PrivacyProtocolPasswordID)
	return o
}

// SetSnmpv3PrivacyProtocolPasswordID adds the snmpv3PrivacyProtocolPasswordId to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSnmpv3PrivacyProtocolPasswordID(snmpv3PrivacyProtocolPasswordID *string) {
	o.Snmpv3PrivacyProtocolPasswordID = snmpv3PrivacyProtocolPasswordID
}

// WithSnmpv3User adds the snmpv3User to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithSnmpv3User(snmpv3User *string) *PostAutoDiscoveryNetworksParams {
	o.SetSnmpv3User(snmpv3User)
	return o
}

// SetSnmpv3User adds the snmpv3User to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetSnmpv3User(snmpv3User *string) {
	o.Snmpv3User = snmpv3User
}

// WithStripDomainName adds the stripDomainName to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithStripDomainName(stripDomainName *string) *PostAutoDiscoveryNetworksParams {
	o.SetStripDomainName(stripDomainName)
	return o
}

// SetStripDomainName adds the stripDomainName to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetStripDomainName(stripDomainName *string) {
	o.StripDomainName = stripDomainName
}

// WithUseNameAliasPortDescr adds the useNameAliasPortDescr to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithUseNameAliasPortDescr(useNameAliasPortDescr *string) *PostAutoDiscoveryNetworksParams {
	o.SetUseNameAliasPortDescr(useNameAliasPortDescr)
	return o
}

// SetUseNameAliasPortDescr adds the useNameAliasPortDescr to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetUseNameAliasPortDescr(useNameAliasPortDescr *string) {
	o.UseNameAliasPortDescr = useNameAliasPortDescr
}

// WithVlansToIgnore adds the vlansToIgnore to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithVlansToIgnore(vlansToIgnore *string) *PostAutoDiscoveryNetworksParams {
	o.SetVlansToIgnore(vlansToIgnore)
	return o
}

// SetVlansToIgnore adds the vlansToIgnore to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetVlansToIgnore(vlansToIgnore *string) {
	o.VlansToIgnore = vlansToIgnore
}

// WithVrfgroup adds the vrfgroup to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) WithVrfgroup(vrfgroup *string) *PostAutoDiscoveryNetworksParams {
	o.SetVrfgroup(vrfgroup)
	return o
}

// SetVrfgroup adds the vrfgroup to the post auto discovery networks params
func (o *PostAutoDiscoveryNetworksParams) SetVrfgroup(vrfgroup *string) {
	o.Vrfgroup = vrfgroup
}

// WriteToRequest writes these params to a swagger request
func (o *PostAutoDiscoveryNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AutodiscoverCdpDevices != nil {

		// form param autodiscover_cdp_devices
		var frAutodiscoverCdpDevices string
		if o.AutodiscoverCdpDevices != nil {
			frAutodiscoverCdpDevices = *o.AutodiscoverCdpDevices
		}
		fAutodiscoverCdpDevices := frAutodiscoverCdpDevices
		if fAutodiscoverCdpDevices != "" {
			if err := r.SetFormParam("autodiscover_cdp_devices", fAutodiscoverCdpDevices); err != nil {
				return err
			}
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

	if o.DeleteOlderMacAssociationAfter != nil {

		// form param delete_older_mac_association_after
		var frDeleteOlderMacAssociationAfter string
		if o.DeleteOlderMacAssociationAfter != nil {
			frDeleteOlderMacAssociationAfter = *o.DeleteOlderMacAssociationAfter
		}
		fDeleteOlderMacAssociationAfter := frDeleteOlderMacAssociationAfter
		if fDeleteOlderMacAssociationAfter != "" {
			if err := r.SetFormParam("delete_older_mac_association_after", fDeleteOlderMacAssociationAfter); err != nil {
				return err
			}
		}

	}

	if o.DeleteSwitchPortNotFound != nil {

		// form param delete_switch_port_not_found
		var frDeleteSwitchPortNotFound string
		if o.DeleteSwitchPortNotFound != nil {
			frDeleteSwitchPortNotFound = *o.DeleteSwitchPortNotFound
		}
		fDeleteSwitchPortNotFound := frDeleteSwitchPortNotFound
		if fDeleteSwitchPortNotFound != "" {
			if err := r.SetFormParam("delete_switch_port_not_found", fDeleteSwitchPortNotFound); err != nil {
				return err
			}
		}

	}

	if o.DiscoverServices != nil {

		// form param discover_services
		var frDiscoverServices string
		if o.DiscoverServices != nil {
			frDiscoverServices = *o.DiscoverServices
		}
		fDiscoverServices := frDiscoverServices
		if fDiscoverServices != "" {
			if err := r.SetFormParam("discover_services", fDiscoverServices); err != nil {
				return err
			}
		}

	}

	if o.EndIPAddress != nil {

		// form param end_ip_address
		var frEndIPAddress string
		if o.EndIPAddress != nil {
			frEndIPAddress = *o.EndIPAddress
		}
		fEndIPAddress := frEndIPAddress
		if fEndIPAddress != "" {
			if err := r.SetFormParam("end_ip_address", fEndIPAddress); err != nil {
				return err
			}
		}

	}

	if o.GetAllSwitchPorts != nil {

		// form param get_all_switch_ports
		var frGetAllSwitchPorts string
		if o.GetAllSwitchPorts != nil {
			frGetAllSwitchPorts = *o.GetAllSwitchPorts
		}
		fGetAllSwitchPorts := frGetAllSwitchPorts
		if fGetAllSwitchPorts != "" {
			if err := r.SetFormParam("get_all_switch_ports", fGetAllSwitchPorts); err != nil {
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

	if o.HostnamePrecedence != nil {

		// form param hostname_precedence
		var frHostnamePrecedence string
		if o.HostnamePrecedence != nil {
			frHostnamePrecedence = *o.HostnamePrecedence
		}
		fHostnamePrecedence := frHostnamePrecedence
		if fHostnamePrecedence != "" {
			if err := r.SetFormParam("hostname_precedence", fHostnamePrecedence); err != nil {
				return err
			}
		}

	}

	if o.Ipaddress != nil {

		// form param ipaddress
		var frIpaddress string
		if o.Ipaddress != nil {
			frIpaddress = *o.Ipaddress
		}
		fIpaddress := frIpaddress
		if fIpaddress != "" {
			if err := r.SetFormParam("ipaddress", fIpaddress); err != nil {
				return err
			}
		}

	}

	if o.MaskBits != nil {

		// form param mask_bits
		var frMaskBits int64
		if o.MaskBits != nil {
			frMaskBits = *o.MaskBits
		}
		fMaskBits := swag.FormatInt64(frMaskBits)
		if fMaskBits != "" {
			if err := r.SetFormParam("mask_bits", fMaskBits); err != nil {
				return err
			}
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

	if o.PortNamePrefixToIgnoreMacs != nil {

		// form param port_name_prefix_to_ignore_macs
		var frPortNamePrefixToIgnoreMacs string
		if o.PortNamePrefixToIgnoreMacs != nil {
			frPortNamePrefixToIgnoreMacs = *o.PortNamePrefixToIgnoreMacs
		}
		fPortNamePrefixToIgnoreMacs := frPortNamePrefixToIgnoreMacs
		if fPortNamePrefixToIgnoreMacs != "" {
			if err := r.SetFormParam("port_name_prefix_to_ignore_macs", fPortNamePrefixToIgnoreMacs); err != nil {
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

	if o.ServiceLevel != nil {

		// form param service_level
		var frServiceLevel string
		if o.ServiceLevel != nil {
			frServiceLevel = *o.ServiceLevel
		}
		fServiceLevel := frServiceLevel
		if fServiceLevel != "" {
			if err := r.SetFormParam("service_level", fServiceLevel); err != nil {
				return err
			}
		}

	}

	if o.SkipVlanIndexing != nil {

		// form param skip_vlan_indexing
		var frSkipVlanIndexing string
		if o.SkipVlanIndexing != nil {
			frSkipVlanIndexing = *o.SkipVlanIndexing
		}
		fSkipVlanIndexing := frSkipVlanIndexing
		if fSkipVlanIndexing != "" {
			if err := r.SetFormParam("skip_vlan_indexing", fSkipVlanIndexing); err != nil {
				return err
			}
		}

	}

	if o.SnmpPort != nil {

		// form param snmp_port
		var frSnmpPort int64
		if o.SnmpPort != nil {
			frSnmpPort = *o.SnmpPort
		}
		fSnmpPort := swag.FormatInt64(frSnmpPort)
		if fSnmpPort != "" {
			if err := r.SetFormParam("snmp_port", fSnmpPort); err != nil {
				return err
			}
		}

	}

	if o.SnmpString != nil {

		// form param snmp_string
		var frSnmpString string
		if o.SnmpString != nil {
			frSnmpString = *o.SnmpString
		}
		fSnmpString := frSnmpString
		if fSnmpString != "" {
			if err := r.SetFormParam("snmp_string", fSnmpString); err != nil {
				return err
			}
		}

	}

	if o.SnmpStringID != nil {

		// form param snmp_string_id
		var frSnmpStringID string
		if o.SnmpStringID != nil {
			frSnmpStringID = *o.SnmpStringID
		}
		fSnmpStringID := frSnmpStringID
		if fSnmpStringID != "" {
			if err := r.SetFormParam("snmp_string_id", fSnmpStringID); err != nil {
				return err
			}
		}

	}

	if o.SnmpStringIds != nil {

		// form param snmp_string_ids
		var frSnmpStringIds string
		if o.SnmpStringIds != nil {
			frSnmpStringIds = *o.SnmpStringIds
		}
		fSnmpStringIds := frSnmpStringIds
		if fSnmpStringIds != "" {
			if err := r.SetFormParam("snmp_string_ids", fSnmpStringIds); err != nil {
				return err
			}
		}

	}

	if o.SnmpStrings != nil {

		// form param snmp_strings
		var frSnmpStrings string
		if o.SnmpStrings != nil {
			frSnmpStrings = *o.SnmpStrings
		}
		fSnmpStrings := frSnmpStrings
		if fSnmpStrings != "" {
			if err := r.SetFormParam("snmp_strings", fSnmpStrings); err != nil {
				return err
			}
		}

	}

	if o.SnmpVersion != nil {

		// form param snmp_version
		var frSnmpVersion string
		if o.SnmpVersion != nil {
			frSnmpVersion = *o.SnmpVersion
		}
		fSnmpVersion := frSnmpVersion
		if fSnmpVersion != "" {
			if err := r.SetFormParam("snmp_version", fSnmpVersion); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3AuthMode != nil {

		// form param snmpv3_auth_mode
		var frSnmpv3AuthMode string
		if o.Snmpv3AuthMode != nil {
			frSnmpv3AuthMode = *o.Snmpv3AuthMode
		}
		fSnmpv3AuthMode := frSnmpv3AuthMode
		if fSnmpv3AuthMode != "" {
			if err := r.SetFormParam("snmpv3_auth_mode", fSnmpv3AuthMode); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3AuthPassword != nil {

		// form param snmpv3_auth_password
		var frSnmpv3AuthPassword string
		if o.Snmpv3AuthPassword != nil {
			frSnmpv3AuthPassword = *o.Snmpv3AuthPassword
		}
		fSnmpv3AuthPassword := frSnmpv3AuthPassword
		if fSnmpv3AuthPassword != "" {
			if err := r.SetFormParam("snmpv3_auth_password", fSnmpv3AuthPassword); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3AuthPasswordID != nil {

		// form param snmpv3_auth_password_id
		var frSnmpv3AuthPasswordID string
		if o.Snmpv3AuthPasswordID != nil {
			frSnmpv3AuthPasswordID = *o.Snmpv3AuthPasswordID
		}
		fSnmpv3AuthPasswordID := frSnmpv3AuthPasswordID
		if fSnmpv3AuthPasswordID != "" {
			if err := r.SetFormParam("snmpv3_auth_password_id", fSnmpv3AuthPasswordID); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3AuthProtocol != nil {

		// form param snmpv3_auth_protocol
		var frSnmpv3AuthProtocol string
		if o.Snmpv3AuthProtocol != nil {
			frSnmpv3AuthProtocol = *o.Snmpv3AuthProtocol
		}
		fSnmpv3AuthProtocol := frSnmpv3AuthProtocol
		if fSnmpv3AuthProtocol != "" {
			if err := r.SetFormParam("snmpv3_auth_protocol", fSnmpv3AuthProtocol); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3Context != nil {

		// form param snmpv3_context
		var frSnmpv3Context string
		if o.Snmpv3Context != nil {
			frSnmpv3Context = *o.Snmpv3Context
		}
		fSnmpv3Context := frSnmpv3Context
		if fSnmpv3Context != "" {
			if err := r.SetFormParam("snmpv3_context", fSnmpv3Context); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3PrivacyProtocol != nil {

		// form param snmpv3_privacy_protocol
		var frSnmpv3PrivacyProtocol string
		if o.Snmpv3PrivacyProtocol != nil {
			frSnmpv3PrivacyProtocol = *o.Snmpv3PrivacyProtocol
		}
		fSnmpv3PrivacyProtocol := frSnmpv3PrivacyProtocol
		if fSnmpv3PrivacyProtocol != "" {
			if err := r.SetFormParam("snmpv3_privacy_protocol", fSnmpv3PrivacyProtocol); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3PrivacyProtocolPassword != nil {

		// form param snmpv3_privacy_protocol_password
		var frSnmpv3PrivacyProtocolPassword string
		if o.Snmpv3PrivacyProtocolPassword != nil {
			frSnmpv3PrivacyProtocolPassword = *o.Snmpv3PrivacyProtocolPassword
		}
		fSnmpv3PrivacyProtocolPassword := frSnmpv3PrivacyProtocolPassword
		if fSnmpv3PrivacyProtocolPassword != "" {
			if err := r.SetFormParam("snmpv3_privacy_protocol_password", fSnmpv3PrivacyProtocolPassword); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3PrivacyProtocolPasswordID != nil {

		// form param snmpv3_privacy_protocol_password_id
		var frSnmpv3PrivacyProtocolPasswordID string
		if o.Snmpv3PrivacyProtocolPasswordID != nil {
			frSnmpv3PrivacyProtocolPasswordID = *o.Snmpv3PrivacyProtocolPasswordID
		}
		fSnmpv3PrivacyProtocolPasswordID := frSnmpv3PrivacyProtocolPasswordID
		if fSnmpv3PrivacyProtocolPasswordID != "" {
			if err := r.SetFormParam("snmpv3_privacy_protocol_password_id", fSnmpv3PrivacyProtocolPasswordID); err != nil {
				return err
			}
		}

	}

	if o.Snmpv3User != nil {

		// form param snmpv3_user
		var frSnmpv3User string
		if o.Snmpv3User != nil {
			frSnmpv3User = *o.Snmpv3User
		}
		fSnmpv3User := frSnmpv3User
		if fSnmpv3User != "" {
			if err := r.SetFormParam("snmpv3_user", fSnmpv3User); err != nil {
				return err
			}
		}

	}

	if o.StripDomainName != nil {

		// form param strip_domain_name
		var frStripDomainName string
		if o.StripDomainName != nil {
			frStripDomainName = *o.StripDomainName
		}
		fStripDomainName := frStripDomainName
		if fStripDomainName != "" {
			if err := r.SetFormParam("strip_domain_name", fStripDomainName); err != nil {
				return err
			}
		}

	}

	if o.UseNameAliasPortDescr != nil {

		// form param use_name_alias_port_descr
		var frUseNameAliasPortDescr string
		if o.UseNameAliasPortDescr != nil {
			frUseNameAliasPortDescr = *o.UseNameAliasPortDescr
		}
		fUseNameAliasPortDescr := frUseNameAliasPortDescr
		if fUseNameAliasPortDescr != "" {
			if err := r.SetFormParam("use_name_alias_port_descr", fUseNameAliasPortDescr); err != nil {
				return err
			}
		}

	}

	if o.VlansToIgnore != nil {

		// form param vlans_to_ignore
		var frVlansToIgnore string
		if o.VlansToIgnore != nil {
			frVlansToIgnore = *o.VlansToIgnore
		}
		fVlansToIgnore := frVlansToIgnore
		if fVlansToIgnore != "" {
			if err := r.SetFormParam("vlans_to_ignore", fVlansToIgnore); err != nil {
				return err
			}
		}

	}

	if o.Vrfgroup != nil {

		// form param vrfgroup
		var frVrfgroup string
		if o.Vrfgroup != nil {
			frVrfgroup = *o.Vrfgroup
		}
		fVrfgroup := frVrfgroup
		if fVrfgroup != "" {
			if err := r.SetFormParam("vrfgroup", fVrfgroup); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
