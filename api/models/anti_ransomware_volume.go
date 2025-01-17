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

// AntiRansomwareVolume Anti-ransomware properties of volumes.
//
// swagger:model anti_ransomware_volume
type AntiRansomwareVolume struct {

	// anti ransomware volume inline attack reports
	// Read Only: true
	AntiRansomwareVolumeInlineAttackReports []*AntiRansomwareAttackReport `json:"attack_reports,omitempty" yaml:"attack_reports,omitempty"`

	// anti ransomware volume inline suspect files
	// Read Only: true
	AntiRansomwareVolumeInlineSuspectFiles []*AntiRansomwareVolumeInlineSuspectFilesInlineArrayItem `json:"suspect_files,omitempty" yaml:"suspect_files,omitempty"`

	// Probability of a ransomware attack.<br>`none` No files are suspected of ransomware activity.<br>`low` A number of files are suspected of ransomware activity.<br>`moderate` A moderate number of files are suspected of ransomware activity.<br>`high` A large number of files are suspected of ransomware activity.
	// Read Only: true
	// Enum: ["none","low","moderate","high"]
	AttackProbability *string `json:"attack_probability,omitempty" yaml:"attack_probability,omitempty"`

	// Time when Anti-ransomware monitoring `state` is set to dry-run value for starting evaluation mode.
	// Read Only: true
	// Format: date-time
	DryRunStartTime *strfmt.DateTime `json:"dry_run_start_time,omitempty" yaml:"dry_run_start_time,omitempty"`

	// event log
	EventLog *AntiRansomwareVolumeInlineEventLog `json:"event_log,omitempty" yaml:"event_log,omitempty"`

	// space
	Space *AntiRansomwareVolumeInlineSpace `json:"space,omitempty" yaml:"space,omitempty"`

	// Anti-ransomware state.<br>`disabled` Anti-ransomware monitoring is disabled on the volume.  This is the default state in a POST operation.<br>`disable_in_progress` Anti-ransomware monitoring is being disabled and a cleanup operation is in effect. Valid in GET operation.<br>`dry_run` Anti-ransomware monitoring is enabled in the evaluation mode.<br>`enabled` Anti-ransomware monitoring is active on the volume.<br>`paused` Anti-ransomware monitoring is paused on the volume.<br>`enable_paused` Anti-ransomware monitoring is paused on the volume from its earlier enabled state. Valid in GET operation. <br>`dry_run_paused` Anti-ransomware monitoring is paused on the volume from its earlier dry_run state. Valid in GET operation. <br>For POST, the valid Anti-ransomware states are only `disabled`, `enabled` and `dry_run`, whereas for PATCH, `paused` is also valid along with the three valid states for POST.
	// Enum: ["disabled","disable_in_progress","dry_run","enabled","paused","enable_paused","dry_run_paused"]
	State *string `json:"state,omitempty" yaml:"state,omitempty"`

	// Indicates whether or not to set the surge values as historical values. This field is no longer supported. Use update-baseline-from-surge instead.
	SurgeAsNormal *bool `json:"surge_as_normal,omitempty" yaml:"surge_as_normal,omitempty"`

	// Sets the observed surge value as the new baseline on a volume.
	UpdateBaselineFromSurge *bool `json:"update_baseline_from_surge,omitempty" yaml:"update_baseline_from_surge,omitempty"`
}

// Validate validates this anti ransomware volume
func (m *AntiRansomwareVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAntiRansomwareVolumeInlineAttackReports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAntiRansomwareVolumeInlineSuspectFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttackProbability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDryRunStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventLog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareVolume) validateAntiRansomwareVolumeInlineAttackReports(formats strfmt.Registry) error {
	if swag.IsZero(m.AntiRansomwareVolumeInlineAttackReports) { // not required
		return nil
	}

	for i := 0; i < len(m.AntiRansomwareVolumeInlineAttackReports); i++ {
		if swag.IsZero(m.AntiRansomwareVolumeInlineAttackReports[i]) { // not required
			continue
		}

		if m.AntiRansomwareVolumeInlineAttackReports[i] != nil {
			if err := m.AntiRansomwareVolumeInlineAttackReports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attack_reports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attack_reports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AntiRansomwareVolume) validateAntiRansomwareVolumeInlineSuspectFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.AntiRansomwareVolumeInlineSuspectFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.AntiRansomwareVolumeInlineSuspectFiles); i++ {
		if swag.IsZero(m.AntiRansomwareVolumeInlineSuspectFiles[i]) { // not required
			continue
		}

		if m.AntiRansomwareVolumeInlineSuspectFiles[i] != nil {
			if err := m.AntiRansomwareVolumeInlineSuspectFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("suspect_files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("suspect_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var antiRansomwareVolumeTypeAttackProbabilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","low","moderate","high"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		antiRansomwareVolumeTypeAttackProbabilityPropEnum = append(antiRansomwareVolumeTypeAttackProbabilityPropEnum, v)
	}
}

const (

	// AntiRansomwareVolumeAttackProbabilityNone captures enum value "none"
	AntiRansomwareVolumeAttackProbabilityNone string = "none"

	// AntiRansomwareVolumeAttackProbabilityLow captures enum value "low"
	AntiRansomwareVolumeAttackProbabilityLow string = "low"

	// AntiRansomwareVolumeAttackProbabilityModerate captures enum value "moderate"
	AntiRansomwareVolumeAttackProbabilityModerate string = "moderate"

	// AntiRansomwareVolumeAttackProbabilityHigh captures enum value "high"
	AntiRansomwareVolumeAttackProbabilityHigh string = "high"
)

// prop value enum
func (m *AntiRansomwareVolume) validateAttackProbabilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, antiRansomwareVolumeTypeAttackProbabilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AntiRansomwareVolume) validateAttackProbability(formats strfmt.Registry) error {
	if swag.IsZero(m.AttackProbability) { // not required
		return nil
	}

	// value enum
	if err := m.validateAttackProbabilityEnum("attack_probability", "body", *m.AttackProbability); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolume) validateDryRunStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.DryRunStartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("dry_run_start_time", "body", "date-time", m.DryRunStartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolume) validateEventLog(formats strfmt.Registry) error {
	if swag.IsZero(m.EventLog) { // not required
		return nil
	}

	if m.EventLog != nil {
		if err := m.EventLog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_log")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_log")
			}
			return err
		}
	}

	return nil
}

func (m *AntiRansomwareVolume) validateSpace(formats strfmt.Registry) error {
	if swag.IsZero(m.Space) { // not required
		return nil
	}

	if m.Space != nil {
		if err := m.Space.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("space")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("space")
			}
			return err
		}
	}

	return nil
}

var antiRansomwareVolumeTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["disabled","disable_in_progress","dry_run","enabled","paused","enable_paused","dry_run_paused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		antiRansomwareVolumeTypeStatePropEnum = append(antiRansomwareVolumeTypeStatePropEnum, v)
	}
}

const (

	// AntiRansomwareVolumeStateDisabled captures enum value "disabled"
	AntiRansomwareVolumeStateDisabled string = "disabled"

	// AntiRansomwareVolumeStateDisableInProgress captures enum value "disable_in_progress"
	AntiRansomwareVolumeStateDisableInProgress string = "disable_in_progress"

	// AntiRansomwareVolumeStateDryRun captures enum value "dry_run"
	AntiRansomwareVolumeStateDryRun string = "dry_run"

	// AntiRansomwareVolumeStateEnabled captures enum value "enabled"
	AntiRansomwareVolumeStateEnabled string = "enabled"

	// AntiRansomwareVolumeStatePaused captures enum value "paused"
	AntiRansomwareVolumeStatePaused string = "paused"

	// AntiRansomwareVolumeStateEnablePaused captures enum value "enable_paused"
	AntiRansomwareVolumeStateEnablePaused string = "enable_paused"

	// AntiRansomwareVolumeStateDryRunPaused captures enum value "dry_run_paused"
	AntiRansomwareVolumeStateDryRunPaused string = "dry_run_paused"
)

// prop value enum
func (m *AntiRansomwareVolume) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, antiRansomwareVolumeTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AntiRansomwareVolume) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this anti ransomware volume based on the context it is used
func (m *AntiRansomwareVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAntiRansomwareVolumeInlineAttackReports(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAntiRansomwareVolumeInlineSuspectFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAttackProbability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDryRunStartTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareVolume) contextValidateAntiRansomwareVolumeInlineAttackReports(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "attack_reports", "body", []*AntiRansomwareAttackReport(m.AntiRansomwareVolumeInlineAttackReports)); err != nil {
		return err
	}

	for i := 0; i < len(m.AntiRansomwareVolumeInlineAttackReports); i++ {

		if m.AntiRansomwareVolumeInlineAttackReports[i] != nil {

			if swag.IsZero(m.AntiRansomwareVolumeInlineAttackReports[i]) { // not required
				return nil
			}

			if err := m.AntiRansomwareVolumeInlineAttackReports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attack_reports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attack_reports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AntiRansomwareVolume) contextValidateAntiRansomwareVolumeInlineSuspectFiles(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "suspect_files", "body", []*AntiRansomwareVolumeInlineSuspectFilesInlineArrayItem(m.AntiRansomwareVolumeInlineSuspectFiles)); err != nil {
		return err
	}

	for i := 0; i < len(m.AntiRansomwareVolumeInlineSuspectFiles); i++ {

		if m.AntiRansomwareVolumeInlineSuspectFiles[i] != nil {

			if swag.IsZero(m.AntiRansomwareVolumeInlineSuspectFiles[i]) { // not required
				return nil
			}

			if err := m.AntiRansomwareVolumeInlineSuspectFiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("suspect_files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("suspect_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AntiRansomwareVolume) contextValidateAttackProbability(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "attack_probability", "body", m.AttackProbability); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolume) contextValidateDryRunStartTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dry_run_start_time", "body", m.DryRunStartTime); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolume) contextValidateEventLog(ctx context.Context, formats strfmt.Registry) error {

	if m.EventLog != nil {

		if swag.IsZero(m.EventLog) { // not required
			return nil
		}

		if err := m.EventLog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_log")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_log")
			}
			return err
		}
	}

	return nil
}

func (m *AntiRansomwareVolume) contextValidateSpace(ctx context.Context, formats strfmt.Registry) error {

	if m.Space != nil {

		if swag.IsZero(m.Space) { // not required
			return nil
		}

		if err := m.Space.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("space")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("space")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntiRansomwareVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntiRansomwareVolume) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AntiRansomwareVolumeInlineEventLog anti ransomware volume inline event log
//
// swagger:model anti_ransomware_volume_inline_event_log
type AntiRansomwareVolumeInlineEventLog struct {

	// Specifies whether to send an EMS when a new file extension is discovered.
	IsEnabledOnNewFileExtensionSeen *bool `json:"is_enabled_on_new_file_extension_seen,omitempty" yaml:"is_enabled_on_new_file_extension_seen,omitempty"`

	// Specifies whether to send an EMS when a Snapshot copy is created.
	IsEnabledOnSnapshotCopyCreation *bool `json:"is_enabled_on_snapshot_copy_creation,omitempty" yaml:"is_enabled_on_snapshot_copy_creation,omitempty"`
}

// Validate validates this anti ransomware volume inline event log
func (m *AntiRansomwareVolumeInlineEventLog) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this anti ransomware volume inline event log based on context it is used
func (m *AntiRansomwareVolumeInlineEventLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AntiRansomwareVolumeInlineEventLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntiRansomwareVolumeInlineEventLog) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareVolumeInlineEventLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AntiRansomwareVolumeInlineSpace anti ransomware volume inline space
//
// swagger:model anti_ransomware_volume_inline_space
type AntiRansomwareVolumeInlineSpace struct {

	// Total number of Anti-ransomware backup Snapshot copies.
	// Read Only: true
	SnapshotCount *int64 `json:"snapshot_count,omitempty" yaml:"snapshot_count,omitempty"`

	// Total space in bytes used by the Anti-ransomware feature.
	// Read Only: true
	Used *int64 `json:"used,omitempty" yaml:"used,omitempty"`

	// Space in bytes used by the Anti-ransomware analytics logs.
	// Read Only: true
	UsedByLogs *int64 `json:"used_by_logs,omitempty" yaml:"used_by_logs,omitempty"`

	// Space in bytes used by the Anti-ransomware backup Snapshot copies.
	// Read Only: true
	UsedBySnapshots *int64 `json:"used_by_snapshots,omitempty" yaml:"used_by_snapshots,omitempty"`
}

// Validate validates this anti ransomware volume inline space
func (m *AntiRansomwareVolumeInlineSpace) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this anti ransomware volume inline space based on the context it is used
func (m *AntiRansomwareVolumeInlineSpace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSnapshotCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsedByLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsedBySnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareVolumeInlineSpace) contextValidateSnapshotCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "space"+"."+"snapshot_count", "body", m.SnapshotCount); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeInlineSpace) contextValidateUsed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "space"+"."+"used", "body", m.Used); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeInlineSpace) contextValidateUsedByLogs(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "space"+"."+"used_by_logs", "body", m.UsedByLogs); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeInlineSpace) contextValidateUsedBySnapshots(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "space"+"."+"used_by_snapshots", "body", m.UsedBySnapshots); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntiRansomwareVolumeInlineSpace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntiRansomwareVolumeInlineSpace) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareVolumeInlineSpace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AntiRansomwareVolumeInlineSuspectFilesInlineArrayItem anti ransomware volume inline suspect files inline array item
//
// swagger:model anti_ransomware_volume_inline_suspect_files_inline_array_item
type AntiRansomwareVolumeInlineSuspectFilesInlineArrayItem struct {

	// Total number of `suspect_files.format` files observed by the Anti-ransomware analytics engine on the volume.
	// Read Only: true
	Count *int64 `json:"count,omitempty" yaml:"count,omitempty"`

	// Indicates the entropy level of this file type.
	// Read Only: true
	Entropy *string `json:"entropy,omitempty" yaml:"entropy,omitempty"`

	// File formats observed by the Anti-ransomware analytics engine on the volume.
	// Read Only: true
	Format *string `json:"format,omitempty" yaml:"format,omitempty"`
}

// Validate validates this anti ransomware volume inline suspect files inline array item
func (m *AntiRansomwareVolumeInlineSuspectFilesInlineArrayItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this anti ransomware volume inline suspect files inline array item based on the context it is used
func (m *AntiRansomwareVolumeInlineSuspectFilesInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntropy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareVolumeInlineSuspectFilesInlineArrayItem) contextValidateCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeInlineSuspectFilesInlineArrayItem) contextValidateEntropy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "entropy", "body", m.Entropy); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeInlineSuspectFilesInlineArrayItem) contextValidateFormat(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntiRansomwareVolumeInlineSuspectFilesInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntiRansomwareVolumeInlineSuspectFilesInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareVolumeInlineSuspectFilesInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
