package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CIMCBootImageNetworkAddress c i m c boot image network address
// swagger:model CIMC.BootImage_NetworkAddress
type CIMCBootImageNetworkAddress struct {

	// gateway
	// Required: true
	Gateway *string `json:"gateway"`

	// ip addr
	// Required: true
	IPAddr *string `json:"ipAddr"`

	// netmask
	// Required: true
	Netmask *string `json:"netmask"`

	// vlan ids
	VlanIds []float64 `json:"vlanIds"`
}

// Validate validates this c i m c boot image network address
func (m *CIMCBootImageNetworkAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPAddr(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetmask(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVlanIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CIMCBootImageNetworkAddress) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *CIMCBootImageNetworkAddress) validateIPAddr(formats strfmt.Registry) error {

	if err := validate.Required("ipAddr", "body", m.IPAddr); err != nil {
		return err
	}

	return nil
}

func (m *CIMCBootImageNetworkAddress) validateNetmask(formats strfmt.Registry) error {

	if err := validate.Required("netmask", "body", m.Netmask); err != nil {
		return err
	}

	return nil
}

func (m *CIMCBootImageNetworkAddress) validateVlanIds(formats strfmt.Registry) error {

	if swag.IsZero(m.VlanIds) { // not required
		return nil
	}

	return nil
}
