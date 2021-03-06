package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Manager100Manager This is the schema definition for a Manager.  Examples of managers are BMCs, Enclosure Managers, Management Controllers and other subsystems assigned managability functions.
// swagger:model Manager.1.0.0_Manager
type Manager100Manager struct {

	// at odata context
	// Read Only: true
	AtOdataContext strfmt.URI `json:"@odata.context,omitempty"`

	// at odata id
	// Read Only: true
	AtOdataID strfmt.URI `json:"@odata.id,omitempty"`

	// at odata type
	// Read Only: true
	AtOdataType string `json:"@odata.type,omitempty"`

	// actions
	Actions *Manager100ManagerActions `json:"Actions,omitempty"`

	// Information about the Command Shell service provided by this manager.
	CommandShell *Manager100CommandShell `json:"CommandShell,omitempty"`

	// The current DateTime (with offset) for the manager, used to set or read time.
	DateTime strfmt.DateTime `json:"DateTime,omitempty"`

	// The time offset from UTC that the DateTime property is set to in format: +06:00 .
	// Pattern: ([-+][0-1][0-9]:[0-5][0-9])
	DateTimeLocalOffset string `json:"DateTimeLocalOffset,omitempty"`

	// Provides a description of this resource and is used for commonality  in the schema definitions.
	// Read Only: true
	Description string `json:"Description,omitempty"`

	// This is a reference to a collection of NICs that this manager uses for network communication.  It is here that clients will find NIC configuration options and settings.
	// Read Only: true
	EthernetInterfaces *EthernetInterfaceCollectionEthernetInterfaceCollection `json:"EthernetInterfaces,omitempty"`

	// The firmware version of this Manager
	// Read Only: true
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The value of this property shall contain the information about the Graphical Console (KVM-IP) service of this manager.
	GraphicalConsole *Manager100GraphicalConsole `json:"GraphicalConsole,omitempty"`

	// Uniquely identifies the resource within the collection of like resources.
	// Read Only: true
	ID string `json:"Id,omitempty"`

	// links
	Links *Manager100ManagerLinks `json:"Links,omitempty"`

	// This is a reference to a collection of Logs used by the manager.
	// Read Only: true
	LogServices *LogServiceCollectionLogServiceCollection `json:"LogServices,omitempty"`

	// This property represents the type of manager that this resource represents.
	// Read Only: true
	ManagerType string `json:"ManagerType,omitempty"`

	// The model information of this Manager as defined by the manufacturer
	// Read Only: true
	Model string `json:"Model,omitempty"`

	// The name of the resource or array element.
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// This is a reference to the network services and their settings that the manager controls.  It is here that clients will find network configuration options as well as network services.
	// Read Only: true
	NetworkProtocol *ManagerNetworkProtocol100ManagerNetworkProtocol `json:"NetworkProtocol,omitempty"`

	// This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	Oem ResourceOem `json:"Oem,omitempty"`

	// Redundancy information for the managers of this system
	// Read Only: true
	Redundancy []*Odata400IDRef `json:"Redundancy"`

	// redundancy at odata count
	// Read Only: true
	RedundancyAtOdataCount float64 `json:"Redundancy@odata.count,omitempty"`

	// redundancy at odata navigation link
	RedundancyAtOdataNavigationLink *Odata400IDRef `json:"Redundancy@odata.navigationLink,omitempty"`

	// Information about the Serial Console service provided by this manager.
	SerialConsole *Manager100SerialConsole `json:"SerialConsole,omitempty"`

	// This is a reference to a collection of serial interfaces that this manager uses for serial and console communication.  It is here that clients will find serial configuration options and settings.
	// Read Only: true
	SerialInterfaces *SerialInterfaceCollectionSerialInterfaceCollection `json:"SerialInterfaces,omitempty"`

	// The UUID of the Redfish Service Entry Point provided by this manager
	// Read Only: true
	// Pattern: ([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})
	ServiceEntryPointUUID string `json:"ServiceEntryPointUUID,omitempty"`

	// status
	Status *ResourceStatus `json:"Status,omitempty"`

	// The Universal Unique Identifier (UUID) for this Manager
	// Read Only: true
	// Pattern: ([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})
	UUID string `json:"UUID,omitempty"`

	// This is a reference to the Virtual Media services for this particular manager.
	// Read Only: true
	VirtualMedia *VirtualMediaCollectionVirtualMediaCollection `json:"VirtualMedia,omitempty"`
}

// Validate validates this manager 1 0 0 manager
func (m *Manager100Manager) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCommandShell(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDateTimeLocalOffset(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEthernetInterfaces(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGraphicalConsole(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLogServices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateManagerType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetworkProtocol(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRedundancy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRedundancyAtOdataNavigationLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSerialConsole(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSerialInterfaces(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServiceEntryPointUUID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVirtualMedia(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Manager100Manager) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	if m.Actions != nil {

		if err := m.Actions.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager100Manager) validateCommandShell(formats strfmt.Registry) error {

	if swag.IsZero(m.CommandShell) { // not required
		return nil
	}

	if m.CommandShell != nil {

		if err := m.CommandShell.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager100Manager) validateDateTimeLocalOffset(formats strfmt.Registry) error {

	if swag.IsZero(m.DateTimeLocalOffset) { // not required
		return nil
	}

	if err := validate.Pattern("DateTimeLocalOffset", "body", string(m.DateTimeLocalOffset), `([-+][0-1][0-9]:[0-5][0-9])`); err != nil {
		return err
	}

	return nil
}

func (m *Manager100Manager) validateEthernetInterfaces(formats strfmt.Registry) error {

	if swag.IsZero(m.EthernetInterfaces) { // not required
		return nil
	}

	if m.EthernetInterfaces != nil {

		if err := m.EthernetInterfaces.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager100Manager) validateGraphicalConsole(formats strfmt.Registry) error {

	if swag.IsZero(m.GraphicalConsole) { // not required
		return nil
	}

	if m.GraphicalConsole != nil {

		if err := m.GraphicalConsole.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager100Manager) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {

		if err := m.Links.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager100Manager) validateLogServices(formats strfmt.Registry) error {

	if swag.IsZero(m.LogServices) { // not required
		return nil
	}

	if m.LogServices != nil {

		if err := m.LogServices.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var manager100ManagerTypeManagerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ManagementController","EnclosureManager","BMC","RackManager","AuxiliaryController"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		manager100ManagerTypeManagerTypePropEnum = append(manager100ManagerTypeManagerTypePropEnum, v)
	}
}

const (
	// Manager100ManagerManagerTypeManagementController captures enum value "ManagementController"
	Manager100ManagerManagerTypeManagementController string = "ManagementController"
	// Manager100ManagerManagerTypeEnclosureManager captures enum value "EnclosureManager"
	Manager100ManagerManagerTypeEnclosureManager string = "EnclosureManager"
	// Manager100ManagerManagerTypeBMC captures enum value "BMC"
	Manager100ManagerManagerTypeBMC string = "BMC"
	// Manager100ManagerManagerTypeRackManager captures enum value "RackManager"
	Manager100ManagerManagerTypeRackManager string = "RackManager"
	// Manager100ManagerManagerTypeAuxiliaryController captures enum value "AuxiliaryController"
	Manager100ManagerManagerTypeAuxiliaryController string = "AuxiliaryController"
)

// prop value enum
func (m *Manager100Manager) validateManagerTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, manager100ManagerTypeManagerTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Manager100Manager) validateManagerType(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagerType) { // not required
		return nil
	}

	// value enum
	if err := m.validateManagerTypeEnum("ManagerType", "body", m.ManagerType); err != nil {
		return err
	}

	return nil
}

func (m *Manager100Manager) validateNetworkProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkProtocol) { // not required
		return nil
	}

	if m.NetworkProtocol != nil {

		if err := m.NetworkProtocol.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager100Manager) validateRedundancy(formats strfmt.Registry) error {

	if swag.IsZero(m.Redundancy) { // not required
		return nil
	}

	for i := 0; i < len(m.Redundancy); i++ {

		if swag.IsZero(m.Redundancy[i]) { // not required
			continue
		}

		if m.Redundancy[i] != nil {

			if err := m.Redundancy[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Manager100Manager) validateRedundancyAtOdataNavigationLink(formats strfmt.Registry) error {

	if swag.IsZero(m.RedundancyAtOdataNavigationLink) { // not required
		return nil
	}

	if m.RedundancyAtOdataNavigationLink != nil {

		if err := m.RedundancyAtOdataNavigationLink.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager100Manager) validateSerialConsole(formats strfmt.Registry) error {

	if swag.IsZero(m.SerialConsole) { // not required
		return nil
	}

	if m.SerialConsole != nil {

		if err := m.SerialConsole.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager100Manager) validateSerialInterfaces(formats strfmt.Registry) error {

	if swag.IsZero(m.SerialInterfaces) { // not required
		return nil
	}

	if m.SerialInterfaces != nil {

		if err := m.SerialInterfaces.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager100Manager) validateServiceEntryPointUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceEntryPointUUID) { // not required
		return nil
	}

	if err := validate.Pattern("ServiceEntryPointUUID", "body", string(m.ServiceEntryPointUUID), `([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})`); err != nil {
		return err
	}

	return nil
}

func (m *Manager100Manager) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {

		if err := m.Status.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager100Manager) validateUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.Pattern("UUID", "body", string(m.UUID), `([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})`); err != nil {
		return err
	}

	return nil
}

func (m *Manager100Manager) validateVirtualMedia(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualMedia) { // not required
		return nil
	}

	if m.VirtualMedia != nil {

		if err := m.VirtualMedia.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// Manager100ManagerActions The Actions object contains the available custom actions on this resource.
// swagger:model Manager100ManagerActions
type Manager100ManagerActions struct {

	// manager force failover
	NrManagerForceFailover *Manager100ForceFailover `json:"#Manager.ForceFailover,omitempty"`

	// manager modify redundancy set
	NrManagerModifyRedundancySet *Manager100ModifyRedundancySet `json:"#Manager.ModifyRedundancySet,omitempty"`

	// manager reset
	NrManagerReset *Manager100Reset `json:"#Manager.Reset,omitempty"`

	// oem
	Oem interface{} `json:"Oem,omitempty"`
}

// Validate validates this manager100 manager actions
func (m *Manager100ManagerActions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNrManagerForceFailover(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNrManagerModifyRedundancySet(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNrManagerReset(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Manager100ManagerActions) validateNrManagerForceFailover(formats strfmt.Registry) error {

	if swag.IsZero(m.NrManagerForceFailover) { // not required
		return nil
	}

	if m.NrManagerForceFailover != nil {

		if err := m.NrManagerForceFailover.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager100ManagerActions) validateNrManagerModifyRedundancySet(formats strfmt.Registry) error {

	if swag.IsZero(m.NrManagerModifyRedundancySet) { // not required
		return nil
	}

	if m.NrManagerModifyRedundancySet != nil {

		if err := m.NrManagerModifyRedundancySet.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager100ManagerActions) validateNrManagerReset(formats strfmt.Registry) error {

	if swag.IsZero(m.NrManagerReset) { // not required
		return nil
	}

	if m.NrManagerReset != nil {

		if err := m.NrManagerReset.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// Manager100ManagerLinks This object contains the links to other resources that are related to this resource.
// swagger:model Manager100ManagerLinks
type Manager100ManagerLinks struct {

	// This property is an array of references to the chasis that this manager has control over.
	// Read Only: true
	ManagerForChassis []*Odata400IDRef `json:"ManagerForChassis"`

	// manager for chassis at odata count
	// Read Only: true
	ManagerForChassisAtOdataCount float64 `json:"ManagerForChassis@odata.count,omitempty"`

	// manager for chassis at odata navigation link
	ManagerForChassisAtOdataNavigationLink *Odata400IDRef `json:"ManagerForChassis@odata.navigationLink,omitempty"`

	// This property is an array of references to the systems that this manager has control over.
	// Read Only: true
	ManagerForServers []*Odata400IDRef `json:"ManagerForServers"`

	// manager for servers at odata count
	// Read Only: true
	ManagerForServersAtOdataCount float64 `json:"ManagerForServers@odata.count,omitempty"`

	// manager for servers at odata navigation link
	ManagerForServersAtOdataNavigationLink *Odata400IDRef `json:"ManagerForServers@odata.navigationLink,omitempty"`

	// Oem extension object.
	Oem ResourceOem `json:"Oem,omitempty"`
}

// Validate validates this manager100 manager links
func (m *Manager100ManagerLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManagerForChassis(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateManagerForChassisAtOdataNavigationLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateManagerForServers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateManagerForServersAtOdataNavigationLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Manager100ManagerLinks) validateManagerForChassis(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagerForChassis) { // not required
		return nil
	}

	for i := 0; i < len(m.ManagerForChassis); i++ {

		if swag.IsZero(m.ManagerForChassis[i]) { // not required
			continue
		}

		if m.ManagerForChassis[i] != nil {

			if err := m.ManagerForChassis[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Manager100ManagerLinks) validateManagerForChassisAtOdataNavigationLink(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagerForChassisAtOdataNavigationLink) { // not required
		return nil
	}

	if m.ManagerForChassisAtOdataNavigationLink != nil {

		if err := m.ManagerForChassisAtOdataNavigationLink.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager100ManagerLinks) validateManagerForServers(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagerForServers) { // not required
		return nil
	}

	for i := 0; i < len(m.ManagerForServers); i++ {

		if swag.IsZero(m.ManagerForServers[i]) { // not required
			continue
		}

		if m.ManagerForServers[i] != nil {

			if err := m.ManagerForServers[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Manager100ManagerLinks) validateManagerForServersAtOdataNavigationLink(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagerForServersAtOdataNavigationLink) { // not required
		return nil
	}

	if m.ManagerForServersAtOdataNavigationLink != nil {

		if err := m.ManagerForServersAtOdataNavigationLink.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
