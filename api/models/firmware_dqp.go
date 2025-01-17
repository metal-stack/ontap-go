// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FirmwareDqp firmware dqp
//
// swagger:model firmware_dqp
type FirmwareDqp struct {

	// Firmware file name
	// Example: qual_devices_v3
	// Read Only: true
	FileName *string `json:"file_name,omitempty" yaml:"file_name,omitempty"`

	// record count
	// Read Only: true
	RecordCount *FirmwareDqpRecordCount `json:"record_count,omitempty" yaml:"record_count,omitempty"`

	// Firmware revision
	// Example: 20200117
	// Read Only: true
	Revision *string `json:"revision,omitempty" yaml:"revision,omitempty"`

	// Firmware version
	// Example: 3.18
	// Read Only: true
	Version *string `json:"version,omitempty" yaml:"version,omitempty"`
}

// Validate validates this firmware dqp
func (m *FirmwareDqp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecordCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareDqp) validateRecordCount(formats strfmt.Registry) error {
	if swag.IsZero(m.RecordCount) { // not required
		return nil
	}

	if m.RecordCount != nil {
		if err := m.RecordCount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("record_count")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("record_count")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this firmware dqp based on the context it is used
func (m *FirmwareDqp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFileName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecordCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareDqp) contextValidateFileName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "file_name", "body", m.FileName); err != nil {
		return err
	}

	return nil
}

func (m *FirmwareDqp) contextValidateRecordCount(ctx context.Context, formats strfmt.Registry) error {

	if m.RecordCount != nil {

		if swag.IsZero(m.RecordCount) { // not required
			return nil
		}

		if err := m.RecordCount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("record_count")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("record_count")
			}
			return err
		}
	}

	return nil
}

func (m *FirmwareDqp) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "revision", "body", m.Revision); err != nil {
		return err
	}

	return nil
}

func (m *FirmwareDqp) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FirmwareDqp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareDqp) UnmarshalBinary(b []byte) error {
	var res FirmwareDqp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
