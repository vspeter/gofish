package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ComputerSystem100MemorySummary This object describes the memory of the system in general detail.
// swagger:model ComputerSystem.1.0.0_MemorySummary
type ComputerSystem100MemorySummary struct {

	// status
	Status *ResourceStatus `json:"Status,omitempty"`

	// The total installed, operating system-accessible memory (RAM), measured in GiB.
	// Read Only: true
	// Minimum: 0
	TotalSystemMemoryGiB float64 `json:"TotalSystemMemoryGiB,omitempty"`
}

// Validate validates this computer system 1 0 0 memory summary
func (m *ComputerSystem100MemorySummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTotalSystemMemoryGiB(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerSystem100MemorySummary) validateStatus(formats strfmt.Registry) error {

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

func (m *ComputerSystem100MemorySummary) validateTotalSystemMemoryGiB(formats strfmt.Registry) error {

	if swag.IsZero(m.TotalSystemMemoryGiB) { // not required
		return nil
	}

	if err := validate.Minimum("TotalSystemMemoryGiB", "body", float64(m.TotalSystemMemoryGiB), 0, false); err != nil {
		return err
	}

	return nil
}
