// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// License license
//
// swagger:model license
type License struct {

	// A flag indicating whether the license is currently being enforced.
	// Read Only: true
	Active *bool `json:"active,omitempty" yaml:"active,omitempty"`

	// capacity
	Capacity *LicenseInlineCapacity `json:"capacity,omitempty" yaml:"capacity,omitempty"`

	// compliance
	Compliance *LicenseInlineCompliance `json:"compliance,omitempty" yaml:"compliance,omitempty"`

	// A flag indicating whether the license is in evaluation mode.
	// Read Only: true
	Evaluation *bool `json:"evaluation,omitempty" yaml:"evaluation,omitempty"`

	// Date and time when the license expires.
	// Example: 2019-03-02 19:00:00
	// Read Only: true
	// Format: date-time
	ExpiryTime *strfmt.DateTime `json:"expiry_time,omitempty" yaml:"expiry_time,omitempty"`

	// A string that associates the license with a node or cluster.
	// Example: 456-44-1234
	// Read Only: true
	HostID *string `json:"host_id,omitempty" yaml:"host_id,omitempty"`

	// Name of license that enabled the feature.
	// Example: Core Bundle
	// Read Only: true
	InstalledLicense *string `json:"installed_license,omitempty" yaml:"installed_license,omitempty"`

	// Cluster, node or license manager that owns the license.
	// Example: cluster1
	// Read Only: true
	Owner *string `json:"owner,omitempty" yaml:"owner,omitempty"`

	// Serial number of the license.
	// Example: 123456789
	// Read Only: true
	SerialNumber *string `json:"serial_number,omitempty" yaml:"serial_number,omitempty"`

	// A flag indicating whether the Cloud ONTAP system is going to shutdown as the Cloud platform license has already expired.
	// Read Only: true
	ShutdownImminent *bool `json:"shutdown_imminent,omitempty" yaml:"shutdown_imminent,omitempty"`

	// Date and time when the license starts.
	// Example: 2019-02-02 19:00:00
	// Read Only: true
	// Format: date-time
	StartTime *strfmt.DateTime `json:"start_time,omitempty" yaml:"start_time,omitempty"`
}

// Validate validates this license
func (m *License) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompliance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiryTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *License) validateCapacity(formats strfmt.Registry) error {
	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	if m.Capacity != nil {
		if err := m.Capacity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capacity")
			}
			return err
		}
	}

	return nil
}

func (m *License) validateCompliance(formats strfmt.Registry) error {
	if swag.IsZero(m.Compliance) { // not required
		return nil
	}

	if m.Compliance != nil {
		if err := m.Compliance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("compliance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("compliance")
			}
			return err
		}
	}

	return nil
}

func (m *License) validateExpiryTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiryTime) { // not required
		return nil
	}

	if err := validate.FormatOf("expiry_time", "body", "date-time", m.ExpiryTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *License) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this license based on the context it is used
func (m *License) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActive(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCapacity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCompliance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEvaluation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpiryTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstalledLicense(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSerialNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShutdownImminent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *License) contextValidateActive(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "active", "body", m.Active); err != nil {
		return err
	}

	return nil
}

func (m *License) contextValidateCapacity(ctx context.Context, formats strfmt.Registry) error {

	if m.Capacity != nil {

		if swag.IsZero(m.Capacity) { // not required
			return nil
		}

		if err := m.Capacity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capacity")
			}
			return err
		}
	}

	return nil
}

func (m *License) contextValidateCompliance(ctx context.Context, formats strfmt.Registry) error {

	if m.Compliance != nil {

		if swag.IsZero(m.Compliance) { // not required
			return nil
		}

		if err := m.Compliance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("compliance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("compliance")
			}
			return err
		}
	}

	return nil
}

func (m *License) contextValidateEvaluation(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "evaluation", "body", m.Evaluation); err != nil {
		return err
	}

	return nil
}

func (m *License) contextValidateExpiryTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "expiry_time", "body", m.ExpiryTime); err != nil {
		return err
	}

	return nil
}

func (m *License) contextValidateHostID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "host_id", "body", m.HostID); err != nil {
		return err
	}

	return nil
}

func (m *License) contextValidateInstalledLicense(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "installed_license", "body", m.InstalledLicense); err != nil {
		return err
	}

	return nil
}

func (m *License) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "owner", "body", m.Owner); err != nil {
		return err
	}

	return nil
}

func (m *License) contextValidateSerialNumber(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "serial_number", "body", m.SerialNumber); err != nil {
		return err
	}

	return nil
}

func (m *License) contextValidateShutdownImminent(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "shutdown_imminent", "body", m.ShutdownImminent); err != nil {
		return err
	}

	return nil
}

func (m *License) contextValidateStartTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "start_time", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *License) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *License) UnmarshalBinary(b []byte) error {
	var res License
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LicenseInlineCapacity license inline capacity
//
// swagger:model license_inline_capacity
type LicenseInlineCapacity struct {

	// Licensed capacity size (in bytes) that can be used.
	// Read Only: true
	MaximumSize *int64 `json:"maximum_size,omitempty" yaml:"maximum_size,omitempty"`

	// Capacity that is currently used (in bytes).
	// Read Only: true
	UsedSize *int64 `json:"used_size,omitempty" yaml:"used_size,omitempty"`
}

// Validate validates this license inline capacity
func (m *LicenseInlineCapacity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this license inline capacity based on the context it is used
func (m *LicenseInlineCapacity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMaximumSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsedSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicenseInlineCapacity) contextValidateMaximumSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "capacity"+"."+"maximum_size", "body", m.MaximumSize); err != nil {
		return err
	}

	return nil
}

func (m *LicenseInlineCapacity) contextValidateUsedSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "capacity"+"."+"used_size", "body", m.UsedSize); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LicenseInlineCapacity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseInlineCapacity) UnmarshalBinary(b []byte) error {
	var res LicenseInlineCapacity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LicenseInlineCompliance license inline compliance
//
// swagger:model license_inline_compliance
type LicenseInlineCompliance struct {

	// Compliance state of the license.
	// Example: compliant
	// Read Only: true
	// Enum: ["compliant","noncompliant","unlicensed","unknown"]
	State *string `json:"state,omitempty" yaml:"state,omitempty"`
}

// Validate validates this license inline compliance
func (m *LicenseInlineCompliance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var licenseInlineComplianceTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["compliant","noncompliant","unlicensed","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		licenseInlineComplianceTypeStatePropEnum = append(licenseInlineComplianceTypeStatePropEnum, v)
	}
}

const (

	// LicenseInlineComplianceStateCompliant captures enum value "compliant"
	LicenseInlineComplianceStateCompliant string = "compliant"

	// LicenseInlineComplianceStateNoncompliant captures enum value "noncompliant"
	LicenseInlineComplianceStateNoncompliant string = "noncompliant"

	// LicenseInlineComplianceStateUnlicensed captures enum value "unlicensed"
	LicenseInlineComplianceStateUnlicensed string = "unlicensed"

	// LicenseInlineComplianceStateUnknown captures enum value "unknown"
	LicenseInlineComplianceStateUnknown string = "unknown"
)

// prop value enum
func (m *LicenseInlineCompliance) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, licenseInlineComplianceTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LicenseInlineCompliance) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("compliance"+"."+"state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this license inline compliance based on the context it is used
func (m *LicenseInlineCompliance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicenseInlineCompliance) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "compliance"+"."+"state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LicenseInlineCompliance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseInlineCompliance) UnmarshalBinary(b []byte) error {
	var res LicenseInlineCompliance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
