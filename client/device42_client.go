// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/libdevice42/client/admin_users_groups"
	"github.com/mjpitz/libdevice42/client/application_components"
	"github.com/mjpitz/libdevice42/client/asset_device_life_cycle"
	"github.com/mjpitz/libdevice42/client/assets"
	"github.com/mjpitz/libdevice42/client/audit_log"
	"github.com/mjpitz/libdevice42/client/auto_discovery"
	"github.com/mjpitz/libdevice42/client/buildings"
	"github.com/mjpitz/libdevice42/client/cables"
	"github.com/mjpitz/libdevice42/client/certificates"
	"github.com/mjpitz/libdevice42/client/customers"
	"github.com/mjpitz/libdevice42/client/devices"
	"github.com/mjpitz/libdevice42/client/end_users"
	"github.com/mjpitz/libdevice42/client/get_all_tags"
	"github.com/mjpitz/libdevice42/client/hardware_models"
	"github.com/mjpitz/libdevice42/client/history"
	"github.com/mjpitz/libdevice42/client/ip_a_m"
	"github.com/mjpitz/libdevice42/client/object_categories"
	"github.com/mjpitz/libdevice42/client/operating_systems"
	"github.com/mjpitz/libdevice42/client/p_d_u"
	"github.com/mjpitz/libdevice42/client/parts_management"
	"github.com/mjpitz/libdevice42/client/passwords"
	"github.com/mjpitz/libdevice42/client/patch_panels"
	"github.com/mjpitz/libdevice42/client/power_circuits"
	"github.com/mjpitz/libdevice42/client/purchasing"
	"github.com/mjpitz/libdevice42/client/racks"
	"github.com/mjpitz/libdevice42/client/reports"
	"github.com/mjpitz/libdevice42/client/rooms"
	"github.com/mjpitz/libdevice42/client/service_levels"
	"github.com/mjpitz/libdevice42/client/services"
	"github.com/mjpitz/libdevice42/client/software"
	"github.com/mjpitz/libdevice42/client/telco_circuits"
	"github.com/mjpitz/libdevice42/client/vendors"
)

// Default device42 HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "<host>"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new device42 HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Device42 {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new device42 HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Device42 {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new device42 client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Device42 {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Device42)
	cli.Transport = transport
	cli.AdminUsersGroups = admin_users_groups.New(transport, formats)
	cli.ApplicationComponents = application_components.New(transport, formats)
	cli.AssetDeviceLifeCycle = asset_device_life_cycle.New(transport, formats)
	cli.Assets = assets.New(transport, formats)
	cli.AuditLog = audit_log.New(transport, formats)
	cli.AutoDiscovery = auto_discovery.New(transport, formats)
	cli.Buildings = buildings.New(transport, formats)
	cli.Cables = cables.New(transport, formats)
	cli.Certificates = certificates.New(transport, formats)
	cli.Customers = customers.New(transport, formats)
	cli.Devices = devices.New(transport, formats)
	cli.EndUsers = end_users.New(transport, formats)
	cli.GetAllTags = get_all_tags.New(transport, formats)
	cli.HardwareModels = hardware_models.New(transport, formats)
	cli.History = history.New(transport, formats)
	cli.IPam = ip_a_m.New(transport, formats)
	cli.ObjectCategories = object_categories.New(transport, formats)
	cli.OperatingSystems = operating_systems.New(transport, formats)
	cli.Pdu = p_d_u.New(transport, formats)
	cli.PartsManagement = parts_management.New(transport, formats)
	cli.Passwords = passwords.New(transport, formats)
	cli.PatchPanels = patch_panels.New(transport, formats)
	cli.PowerCircuits = power_circuits.New(transport, formats)
	cli.Purchasing = purchasing.New(transport, formats)
	cli.Racks = racks.New(transport, formats)
	cli.Reports = reports.New(transport, formats)
	cli.Rooms = rooms.New(transport, formats)
	cli.ServiceLevels = service_levels.New(transport, formats)
	cli.Services = services.New(transport, formats)
	cli.Software = software.New(transport, formats)
	cli.TelcoCircuits = telco_circuits.New(transport, formats)
	cli.Vendors = vendors.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Device42 is a client for device42
type Device42 struct {
	AdminUsersGroups admin_users_groups.ClientService

	ApplicationComponents application_components.ClientService

	AssetDeviceLifeCycle asset_device_life_cycle.ClientService

	Assets assets.ClientService

	AuditLog audit_log.ClientService

	AutoDiscovery auto_discovery.ClientService

	Buildings buildings.ClientService

	Cables cables.ClientService

	Certificates certificates.ClientService

	Customers customers.ClientService

	Devices devices.ClientService

	EndUsers end_users.ClientService

	GetAllTags get_all_tags.ClientService

	HardwareModels hardware_models.ClientService

	History history.ClientService

	IPam ip_a_m.ClientService

	ObjectCategories object_categories.ClientService

	OperatingSystems operating_systems.ClientService

	Pdu p_d_u.ClientService

	PartsManagement parts_management.ClientService

	Passwords passwords.ClientService

	PatchPanels patch_panels.ClientService

	PowerCircuits power_circuits.ClientService

	Purchasing purchasing.ClientService

	Racks racks.ClientService

	Reports reports.ClientService

	Rooms rooms.ClientService

	ServiceLevels service_levels.ClientService

	Services services.ClientService

	Software software.ClientService

	TelcoCircuits telco_circuits.ClientService

	Vendors vendors.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Device42) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AdminUsersGroups.SetTransport(transport)
	c.ApplicationComponents.SetTransport(transport)
	c.AssetDeviceLifeCycle.SetTransport(transport)
	c.Assets.SetTransport(transport)
	c.AuditLog.SetTransport(transport)
	c.AutoDiscovery.SetTransport(transport)
	c.Buildings.SetTransport(transport)
	c.Cables.SetTransport(transport)
	c.Certificates.SetTransport(transport)
	c.Customers.SetTransport(transport)
	c.Devices.SetTransport(transport)
	c.EndUsers.SetTransport(transport)
	c.GetAllTags.SetTransport(transport)
	c.HardwareModels.SetTransport(transport)
	c.History.SetTransport(transport)
	c.IPam.SetTransport(transport)
	c.ObjectCategories.SetTransport(transport)
	c.OperatingSystems.SetTransport(transport)
	c.Pdu.SetTransport(transport)
	c.PartsManagement.SetTransport(transport)
	c.Passwords.SetTransport(transport)
	c.PatchPanels.SetTransport(transport)
	c.PowerCircuits.SetTransport(transport)
	c.Purchasing.SetTransport(transport)
	c.Racks.SetTransport(transport)
	c.Reports.SetTransport(transport)
	c.Rooms.SetTransport(transport)
	c.ServiceLevels.SetTransport(transport)
	c.Services.SetTransport(transport)
	c.Software.SetTransport(transport)
	c.TelcoCircuits.SetTransport(transport)
	c.Vendors.SetTransport(transport)
}