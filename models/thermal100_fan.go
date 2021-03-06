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

// Thermal100Fan This is the base type for addressable members of an array.
// swagger:model Thermal.1.0.0_Fan
type Thermal100Fan struct {

	// Name of the fan
	// Read Only: true
	FanName string `json:"FanName,omitempty"`

	// Below normal range but not yet fatal
	// Read Only: true
	LowerThresholdCritical float64 `json:"LowerThresholdCritical,omitempty"`

	// Below normal range and is fatal
	// Read Only: true
	LowerThresholdFatal float64 `json:"LowerThresholdFatal,omitempty"`

	// Below normal range
	// Read Only: true
	LowerThresholdNonCritical float64 `json:"LowerThresholdNonCritical,omitempty"`

	// Maximum value for ReadingRPM
	// Read Only: true
	MaxReadingRange float64 `json:"MaxReadingRange,omitempty"`

	// This is the identifier for the member within the collection.
	MemberID string `json:"MemberId,omitempty"`

	// Minimum value for ReadingRPM
	// Read Only: true
	MinReadingRange float64 `json:"MinReadingRange,omitempty"`

	// This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	Oem ResourceOem `json:"Oem,omitempty"`

	// Describes the area or device associated with this fan.
	// Read Only: true
	PhysicalContext string `json:"PhysicalContext,omitempty"`

	// Current fan speed in RPM
	// Read Only: true
	ReadingRPM float64 `json:"ReadingRPM,omitempty"`

	// This structure is used to show redundancy for fans.  The Component ids will reference the members of the redundancy groups.
	// Read Only: true
	Redundancy []*Odata400IDRef `json:"Redundancy"`

	// redundancy at odata count
	// Read Only: true
	RedundancyAtOdataCount float64 `json:"Redundancy@odata.count,omitempty"`

	// redundancy at odata navigation link
	RedundancyAtOdataNavigationLink *Odata400IDRef `json:"Redundancy@odata.navigationLink,omitempty"`

	// The ID(s) of the resources serviced with this fan
	RelatedItem []*Odata400IDRef `json:"RelatedItem"`

	// related item at odata count
	// Read Only: true
	RelatedItemAtOdataCount float64 `json:"RelatedItem@odata.count,omitempty"`

	// related item at odata navigation link
	RelatedItemAtOdataNavigationLink *Odata400IDRef `json:"RelatedItem@odata.navigationLink,omitempty"`

	// status
	Status *ResourceStatus `json:"Status,omitempty"`

	// Above normal range but not yet fatal
	// Read Only: true
	UpperThresholdCritical float64 `json:"UpperThresholdCritical,omitempty"`

	// Above normal range and is fatal
	// Read Only: true
	UpperThresholdFatal float64 `json:"UpperThresholdFatal,omitempty"`

	// Above normal range
	// Read Only: true
	UpperThresholdNonCritical float64 `json:"UpperThresholdNonCritical,omitempty"`
}

// Validate validates this thermal 1 0 0 fan
func (m *Thermal100Fan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhysicalContext(formats); err != nil {
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

	if err := m.validateRelatedItem(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRelatedItemAtOdataNavigationLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var thermal100FanTypePhysicalContextPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Room","Intake","Exhaust","Front","Back","Upper","Lower","CPU","GPU","Backplane","SystemBoard","PowerSupply","VoltageRegulator","StorageDevice","NetworkingDevice","ComputeBay","StorageBay","NetworkBay","ExpansionBay","PowerSupplyBay"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		thermal100FanTypePhysicalContextPropEnum = append(thermal100FanTypePhysicalContextPropEnum, v)
	}
}

const (
	// Thermal100FanPhysicalContextRoom captures enum value "Room"
	Thermal100FanPhysicalContextRoom string = "Room"
	// Thermal100FanPhysicalContextIntake captures enum value "Intake"
	Thermal100FanPhysicalContextIntake string = "Intake"
	// Thermal100FanPhysicalContextExhaust captures enum value "Exhaust"
	Thermal100FanPhysicalContextExhaust string = "Exhaust"
	// Thermal100FanPhysicalContextFront captures enum value "Front"
	Thermal100FanPhysicalContextFront string = "Front"
	// Thermal100FanPhysicalContextBack captures enum value "Back"
	Thermal100FanPhysicalContextBack string = "Back"
	// Thermal100FanPhysicalContextUpper captures enum value "Upper"
	Thermal100FanPhysicalContextUpper string = "Upper"
	// Thermal100FanPhysicalContextLower captures enum value "Lower"
	Thermal100FanPhysicalContextLower string = "Lower"
	// Thermal100FanPhysicalContextCPU captures enum value "CPU"
	Thermal100FanPhysicalContextCPU string = "CPU"
	// Thermal100FanPhysicalContextGPU captures enum value "GPU"
	Thermal100FanPhysicalContextGPU string = "GPU"
	// Thermal100FanPhysicalContextBackplane captures enum value "Backplane"
	Thermal100FanPhysicalContextBackplane string = "Backplane"
	// Thermal100FanPhysicalContextSystemBoard captures enum value "SystemBoard"
	Thermal100FanPhysicalContextSystemBoard string = "SystemBoard"
	// Thermal100FanPhysicalContextPowerSupply captures enum value "PowerSupply"
	Thermal100FanPhysicalContextPowerSupply string = "PowerSupply"
	// Thermal100FanPhysicalContextVoltageRegulator captures enum value "VoltageRegulator"
	Thermal100FanPhysicalContextVoltageRegulator string = "VoltageRegulator"
	// Thermal100FanPhysicalContextStorageDevice captures enum value "StorageDevice"
	Thermal100FanPhysicalContextStorageDevice string = "StorageDevice"
	// Thermal100FanPhysicalContextNetworkingDevice captures enum value "NetworkingDevice"
	Thermal100FanPhysicalContextNetworkingDevice string = "NetworkingDevice"
	// Thermal100FanPhysicalContextComputeBay captures enum value "ComputeBay"
	Thermal100FanPhysicalContextComputeBay string = "ComputeBay"
	// Thermal100FanPhysicalContextStorageBay captures enum value "StorageBay"
	Thermal100FanPhysicalContextStorageBay string = "StorageBay"
	// Thermal100FanPhysicalContextNetworkBay captures enum value "NetworkBay"
	Thermal100FanPhysicalContextNetworkBay string = "NetworkBay"
	// Thermal100FanPhysicalContextExpansionBay captures enum value "ExpansionBay"
	Thermal100FanPhysicalContextExpansionBay string = "ExpansionBay"
	// Thermal100FanPhysicalContextPowerSupplyBay captures enum value "PowerSupplyBay"
	Thermal100FanPhysicalContextPowerSupplyBay string = "PowerSupplyBay"
)

// prop value enum
func (m *Thermal100Fan) validatePhysicalContextEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, thermal100FanTypePhysicalContextPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Thermal100Fan) validatePhysicalContext(formats strfmt.Registry) error {

	if swag.IsZero(m.PhysicalContext) { // not required
		return nil
	}

	// value enum
	if err := m.validatePhysicalContextEnum("PhysicalContext", "body", m.PhysicalContext); err != nil {
		return err
	}

	return nil
}

func (m *Thermal100Fan) validateRedundancy(formats strfmt.Registry) error {

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

func (m *Thermal100Fan) validateRedundancyAtOdataNavigationLink(formats strfmt.Registry) error {

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

func (m *Thermal100Fan) validateRelatedItem(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedItem) { // not required
		return nil
	}

	for i := 0; i < len(m.RelatedItem); i++ {

		if swag.IsZero(m.RelatedItem[i]) { // not required
			continue
		}

		if m.RelatedItem[i] != nil {

			if err := m.RelatedItem[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Thermal100Fan) validateRelatedItemAtOdataNavigationLink(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedItemAtOdataNavigationLink) { // not required
		return nil
	}

	if m.RelatedItemAtOdataNavigationLink != nil {

		if err := m.RelatedItemAtOdataNavigationLink.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Thermal100Fan) validateStatus(formats strfmt.Registry) error {

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
