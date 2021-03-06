package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ComputerSystem100ProcessorSummary This object describes the central processors of the system in general detail.
// swagger:model ComputerSystem.1.0.0_ProcessorSummary
type ComputerSystem100ProcessorSummary struct {

	// The number of processors in the system.
	// Read Only: true
	// Minimum: 0
	Count float64 `json:"Count,omitempty"`

	// The processor model for the primary or majority of processors in this system.
	// Read Only: true
	Model string `json:"Model,omitempty"`

	// status
	Status *ResourceStatus `json:"Status,omitempty"`
}

// Validate validates this computer system 1 0 0 processor summary
func (m *ComputerSystem100ProcessorSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
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

func (m *ComputerSystem100ProcessorSummary) validateCount(formats strfmt.Registry) error {

	if swag.IsZero(m.Count) { // not required
		return nil
	}

	if err := validate.Minimum("Count", "body", float64(m.Count), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ComputerSystem100ProcessorSummary) validateStatus(formats strfmt.Registry) error {

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
