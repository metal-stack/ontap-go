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

// LayoutRequirement layout requirement
//
// swagger:model layout_requirement
type LayoutRequirement struct {

	// Minimum number of disks to create an aggregate.
	// Example: 6
	// Read Only: true
	AggregateMinDisks *int64 `json:"aggregate_min_disks,omitempty" yaml:"aggregate_min_disks,omitempty"`

	// Indicates if this RAID type is the default.
	// Read Only: true
	Default *bool `json:"default,omitempty" yaml:"default,omitempty"`

	// raid group
	RaidGroup *LayoutRequirementInlineRaidGroup `json:"raid_group,omitempty" yaml:"raid_group,omitempty"`

	// RAID type.
	// Read Only: true
	// Enum: ["raid_dp","raid_tec","raid4","raid0"]
	RaidType *string `json:"raid_type,omitempty" yaml:"raid_type,omitempty"`
}

// Validate validates this layout requirement
func (m *LayoutRequirement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRaidGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRaidType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LayoutRequirement) validateRaidGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.RaidGroup) { // not required
		return nil
	}

	if m.RaidGroup != nil {
		if err := m.RaidGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("raid_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("raid_group")
			}
			return err
		}
	}

	return nil
}

var layoutRequirementTypeRaidTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["raid_dp","raid_tec","raid4","raid0"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		layoutRequirementTypeRaidTypePropEnum = append(layoutRequirementTypeRaidTypePropEnum, v)
	}
}

const (

	// LayoutRequirementRaidTypeRaidDp captures enum value "raid_dp"
	LayoutRequirementRaidTypeRaidDp string = "raid_dp"

	// LayoutRequirementRaidTypeRaidTec captures enum value "raid_tec"
	LayoutRequirementRaidTypeRaidTec string = "raid_tec"

	// LayoutRequirementRaidTypeRaid4 captures enum value "raid4"
	LayoutRequirementRaidTypeRaid4 string = "raid4"

	// LayoutRequirementRaidTypeRaid0 captures enum value "raid0"
	LayoutRequirementRaidTypeRaid0 string = "raid0"
)

// prop value enum
func (m *LayoutRequirement) validateRaidTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, layoutRequirementTypeRaidTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LayoutRequirement) validateRaidType(formats strfmt.Registry) error {
	if swag.IsZero(m.RaidType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRaidTypeEnum("raid_type", "body", *m.RaidType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this layout requirement based on the context it is used
func (m *LayoutRequirement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregateMinDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefault(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRaidGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRaidType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LayoutRequirement) contextValidateAggregateMinDisks(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "aggregate_min_disks", "body", m.AggregateMinDisks); err != nil {
		return err
	}

	return nil
}

func (m *LayoutRequirement) contextValidateDefault(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "default", "body", m.Default); err != nil {
		return err
	}

	return nil
}

func (m *LayoutRequirement) contextValidateRaidGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.RaidGroup != nil {

		if swag.IsZero(m.RaidGroup) { // not required
			return nil
		}

		if err := m.RaidGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("raid_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("raid_group")
			}
			return err
		}
	}

	return nil
}

func (m *LayoutRequirement) contextValidateRaidType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "raid_type", "body", m.RaidType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LayoutRequirement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LayoutRequirement) UnmarshalBinary(b []byte) error {
	var res LayoutRequirement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LayoutRequirementInlineRaidGroup layout requirement inline raid group
//
// swagger:model layout_requirement_inline_raid_group
type LayoutRequirementInlineRaidGroup struct {

	// Default number of disks in a RAID group.
	// Example: 16
	// Read Only: true
	Default *int64 `json:"default,omitempty" yaml:"default,omitempty"`

	// Maximum number of disks allowed in a RAID group.
	// Example: 28
	// Read Only: true
	Max *int64 `json:"max,omitempty" yaml:"max,omitempty"`

	// Minimum number of disks allowed in a RAID group.
	// Example: 5
	// Read Only: true
	Min *int64 `json:"min,omitempty" yaml:"min,omitempty"`
}

// Validate validates this layout requirement inline raid group
func (m *LayoutRequirementInlineRaidGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this layout requirement inline raid group based on the context it is used
func (m *LayoutRequirementInlineRaidGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefault(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMax(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LayoutRequirementInlineRaidGroup) contextValidateDefault(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "raid_group"+"."+"default", "body", m.Default); err != nil {
		return err
	}

	return nil
}

func (m *LayoutRequirementInlineRaidGroup) contextValidateMax(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "raid_group"+"."+"max", "body", m.Max); err != nil {
		return err
	}

	return nil
}

func (m *LayoutRequirementInlineRaidGroup) contextValidateMin(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "raid_group"+"."+"min", "body", m.Min); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LayoutRequirementInlineRaidGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LayoutRequirementInlineRaidGroup) UnmarshalBinary(b []byte) error {
	var res LayoutRequirementInlineRaidGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
