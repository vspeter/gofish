package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CIMCBootImageNetworkDevice c i m c boot image network device
// swagger:model CIMC.BootImage_NetworkDevice
type CIMCBootImageNetworkDevice struct {

	// device
	// Required: true
	Device *string `json:"device"`

	// ipv4
	IPV4 *CIMCBootImageNetworkAddress `json:"ipv4,omitempty"`

	// ipv6
	IPV6 *CIMCBootImageNetworkAddress `json:"ipv6,omitempty"`
}

// Validate validates this c i m c boot image network device
func (m *CIMCBootImageNetworkDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPV4(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPV6(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CIMCBootImageNetworkDevice) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

func (m *CIMCBootImageNetworkDevice) validateIPV4(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV4) { // not required
		return nil
	}

	if m.IPV4 != nil {

		if err := m.IPV4.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CIMCBootImageNetworkDevice) validateIPV6(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV6) { // not required
		return nil
	}

	if m.IPV6 != nil {

		if err := m.IPV6.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}