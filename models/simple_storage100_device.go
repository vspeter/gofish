package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// SimpleStorage100Device simple storage 1 0 0 device
// swagger:model SimpleStorage.1.0.0_Device
type SimpleStorage100Device struct {

	// The name of the manufacturer of this device
	// Read Only: true
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The product model number of this device
	// Read Only: true
	Model string `json:"Model,omitempty"`

	// The name of the resource or array element.
	// Required: true
	// Read Only: true
	Name string `json:"Name"`

	// oem
	Oem ResourceOem `json:"Oem,omitempty"`

	// status
	Status *ResourceStatus `json:"Status,omitempty"`
}

// Validate validates this simple storage 1 0 0 device
func (m *SimpleStorage100Device) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
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

func (m *SimpleStorage100Device) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("Name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *SimpleStorage100Device) validateStatus(formats strfmt.Registry) error {

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