// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SnapshotPolicy The Snapshot copy policy object is associated with a read-write volume used to create and delete Snapshot copies at regular intervals.
//
// swagger:model snapshot_policy
type SnapshotPolicy struct {

	// links
	Links *SnapshotPolicyInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// A comment associated with the Snapshot copy policy.
	Comment *string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// Is the Snapshot copy policy enabled?
	// Example: true
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Name of the Snapshot copy policy.
	// Example: default
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set to "svm" when the request is on a data SVM, otherwise set to "cluster".
	// Read Only: true
	// Enum: ["svm","cluster"]
	Scope *string `json:"scope,omitempty" yaml:"scope,omitempty"`

	// snapshot policy inline copies
	SnapshotPolicyInlineCopies []*SnapshotPolicyInlineCopiesInlineArrayItem `json:"copies,omitempty" yaml:"copies,omitempty"`

	// svm
	Svm *SnapshotPolicyInlineSvm `json:"svm,omitempty" yaml:"svm,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this snapshot policy
func (m *SnapshotPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotPolicyInlineCopies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicy) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

var snapshotPolicyTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["svm","cluster"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snapshotPolicyTypeScopePropEnum = append(snapshotPolicyTypeScopePropEnum, v)
	}
}

const (

	// SnapshotPolicyScopeSvm captures enum value "svm"
	SnapshotPolicyScopeSvm string = "svm"

	// SnapshotPolicyScopeCluster captures enum value "cluster"
	SnapshotPolicyScopeCluster string = "cluster"
)

// prop value enum
func (m *SnapshotPolicy) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snapshotPolicyTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnapshotPolicy) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", *m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotPolicy) validateSnapshotPolicyInlineCopies(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotPolicyInlineCopies) { // not required
		return nil
	}

	for i := 0; i < len(m.SnapshotPolicyInlineCopies); i++ {
		if swag.IsZero(m.SnapshotPolicyInlineCopies[i]) { // not required
			continue
		}

		if m.SnapshotPolicyInlineCopies[i] != nil {
			if err := m.SnapshotPolicyInlineCopies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("copies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("copies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapshotPolicy) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot policy based on the context it is used
func (m *SnapshotPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotPolicyInlineCopies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicy) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotPolicy) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotPolicy) contextValidateSnapshotPolicyInlineCopies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SnapshotPolicyInlineCopies); i++ {

		if m.SnapshotPolicyInlineCopies[i] != nil {

			if swag.IsZero(m.SnapshotPolicyInlineCopies[i]) { // not required
				return nil
			}

			if err := m.SnapshotPolicyInlineCopies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("copies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("copies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapshotPolicy) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {

		if swag.IsZero(m.Svm) { // not required
			return nil
		}

		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotPolicy) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPolicy) UnmarshalBinary(b []byte) error {
	var res SnapshotPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotPolicyInlineCopiesInlineArrayItem snapshot policy inline copies inline array item
//
// swagger:model snapshot_policy_inline_copies_inline_array_item
type SnapshotPolicyInlineCopiesInlineArrayItem struct {

	// The number of Snapshot copies to maintain for this schedule.
	Count *int64 `json:"count,omitempty" yaml:"count,omitempty"`

	// The prefix to use while creating Snapshot copies at regular intervals.
	Prefix *string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// The retention period of Snapshot copies for this schedule. The retention period value represents a duration and must be specified in the ISO-8601 duration format. The retention period can be in years, months, days, hours, and minutes. A period specified for years, months, and days is represented in the ISO-8601 format as "P<num>Y", "P<num>M", "P<num>D" respectively, for example "P10Y" represents a duration of 10 years. A duration in hours and minutes is represented by "PT<num>H" and "PT<num>M" respectively. The period string must contain only a single time element that is, either years, months, days, hours, or minutes. A duration which combines different periods is not supported, for example "P1Y10M" is not supported.
	RetentionPeriod *string `json:"retention_period,omitempty" yaml:"retention_period,omitempty"`

	// schedule
	Schedule *SnapshotPolicyInlineCopiesInlineArrayItemInlineSchedule `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// Label for SnapMirror operations
	SnapmirrorLabel *string `json:"snapmirror_label,omitempty" yaml:"snapmirror_label,omitempty"`
}

// Validate validates this snapshot policy inline copies inline array item
func (m *SnapshotPolicyInlineCopiesInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyInlineCopiesInlineArrayItem) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot policy inline copies inline array item based on the context it is used
func (m *SnapshotPolicyInlineCopiesInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyInlineCopiesInlineArrayItem) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {

		if swag.IsZero(m.Schedule) { // not required
			return nil
		}

		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotPolicyInlineCopiesInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPolicyInlineCopiesInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res SnapshotPolicyInlineCopiesInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotPolicyInlineCopiesInlineArrayItemInlineSchedule snapshot policy inline copies inline array item inline schedule
//
// swagger:model snapshot_policy_inline_copies_inline_array_item_inline_schedule
type SnapshotPolicyInlineCopiesInlineArrayItemInlineSchedule struct {

	// links
	Links *SnapshotPolicyInlineCopiesInlineArrayItemInlineScheduleInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// Job schedule name
	// Example: weekly
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// Job schedule UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this snapshot policy inline copies inline array item inline schedule
func (m *SnapshotPolicyInlineCopiesInlineArrayItemInlineSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyInlineCopiesInlineArrayItemInlineSchedule) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot policy inline copies inline array item inline schedule based on the context it is used
func (m *SnapshotPolicyInlineCopiesInlineArrayItemInlineSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyInlineCopiesInlineArrayItemInlineSchedule) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotPolicyInlineCopiesInlineArrayItemInlineSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPolicyInlineCopiesInlineArrayItemInlineSchedule) UnmarshalBinary(b []byte) error {
	var res SnapshotPolicyInlineCopiesInlineArrayItemInlineSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotPolicyInlineCopiesInlineArrayItemInlineScheduleInlineLinks snapshot policy inline copies inline array item inline schedule inline links
//
// swagger:model snapshot_policy_inline_copies_inline_array_item_inline_schedule_inline__links
type SnapshotPolicyInlineCopiesInlineArrayItemInlineScheduleInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this snapshot policy inline copies inline array item inline schedule inline links
func (m *SnapshotPolicyInlineCopiesInlineArrayItemInlineScheduleInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyInlineCopiesInlineArrayItemInlineScheduleInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot policy inline copies inline array item inline schedule inline links based on the context it is used
func (m *SnapshotPolicyInlineCopiesInlineArrayItemInlineScheduleInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyInlineCopiesInlineArrayItemInlineScheduleInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotPolicyInlineCopiesInlineArrayItemInlineScheduleInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPolicyInlineCopiesInlineArrayItemInlineScheduleInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnapshotPolicyInlineCopiesInlineArrayItemInlineScheduleInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotPolicyInlineLinks snapshot policy inline links
//
// swagger:model snapshot_policy_inline__links
type SnapshotPolicyInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this snapshot policy inline links
func (m *SnapshotPolicyInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot policy inline links based on the context it is used
func (m *SnapshotPolicyInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotPolicyInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPolicyInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnapshotPolicyInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotPolicyInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model snapshot_policy_inline_svm
type SnapshotPolicyInlineSvm struct {

	// links
	Links *SnapshotPolicyInlineSvmInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this snapshot policy inline svm
func (m *SnapshotPolicyInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot policy inline svm based on the context it is used
func (m *SnapshotPolicyInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotPolicyInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPolicyInlineSvm) UnmarshalBinary(b []byte) error {
	var res SnapshotPolicyInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotPolicyInlineSvmInlineLinks snapshot policy inline svm inline links
//
// swagger:model snapshot_policy_inline_svm_inline__links
type SnapshotPolicyInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this snapshot policy inline svm inline links
func (m *SnapshotPolicyInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot policy inline svm inline links based on the context it is used
func (m *SnapshotPolicyInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPolicyInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotPolicyInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPolicyInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnapshotPolicyInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}