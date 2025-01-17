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

// ConsistencyGroupSnapshot consistency group snapshot
//
// swagger:model consistency_group_snapshot
type ConsistencyGroupSnapshot struct {

	// links
	Links *SelfLink `json:"_links,omitempty" yaml:"_links,omitempty"`

	// Comment for the Snapshot copy.
	//
	// Example: My Snapshot copy comment
	Comment *string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// consistency group
	ConsistencyGroup *ConsistencyGroupSnapshotInlineConsistencyGroup `json:"consistency_group,omitempty" yaml:"consistency_group,omitempty"`

	// List of volumes which are not in the Snapshot copy.
	//
	// Read Only: true
	ConsistencyGroupSnapshotInlineMissingVolumes []*VolumeReference `json:"missing_volumes,omitempty" yaml:"missing_volumes,omitempty"`

	// List of volume and snapshot identifiers for each volume in the Snapshot copy.
	//
	// Read Only: true
	ConsistencyGroupSnapshotInlineSnapshotVolumes []*ConsistencyGroupVolumeSnapshot `json:"snapshot_volumes,omitempty" yaml:"snapshot_volumes,omitempty"`

	// Consistency type. This is for categorization purposes only. A Snapshot copy should not be set to 'application consistent' unless the host application is quiesced for the Snapshot copy. Valid in POST.
	//
	// Example: crash
	// Enum: ["crash","application"]
	ConsistencyType *string `json:"consistency_type,omitempty" yaml:"consistency_type,omitempty"`

	// Time the snapshot copy was created
	//
	// Example: 2020-10-25 11:20:00
	// Read Only: true
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"create_time,omitempty" yaml:"create_time,omitempty"`

	// Indicates whether the Snapshot copy taken is partial or not.
	//
	// Example: false
	// Read Only: true
	IsPartial *bool `json:"is_partial,omitempty" yaml:"is_partial,omitempty"`

	// Name of the Snapshot copy.
	//
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// Snapmirror Label for the Snapshot copy.
	//
	// Example: sm_label
	SnapmirrorLabel *string `json:"snapmirror_label,omitempty" yaml:"snapmirror_label,omitempty"`

	// The SVM in which the consistency group is located.
	//
	Svm *SvmReference `json:"svm,omitempty" yaml:"svm,omitempty"`

	// The unique identifier of the Snapshot copy. The UUID is generated
	// by ONTAP when the Snapshot copy is created.
	//
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`

	// Specifies whether a write fence will be taken when creating the Snapshot copy. The default is false if there is only one volume in the consistency group, otherwise the default is true.
	//
	WriteFence *bool `json:"write_fence,omitempty" yaml:"write_fence,omitempty"`
}

// Validate validates this consistency group snapshot
func (m *ConsistencyGroupSnapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistencyGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistencyGroupSnapshotInlineMissingVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistencyGroupSnapshotInlineSnapshotVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistencyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateTime(formats); err != nil {
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

func (m *ConsistencyGroupSnapshot) validateLinks(formats strfmt.Registry) error {
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

func (m *ConsistencyGroupSnapshot) validateConsistencyGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistencyGroup) { // not required
		return nil
	}

	if m.ConsistencyGroup != nil {
		if err := m.ConsistencyGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistency_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consistency_group")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupSnapshot) validateConsistencyGroupSnapshotInlineMissingVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistencyGroupSnapshotInlineMissingVolumes) { // not required
		return nil
	}

	for i := 0; i < len(m.ConsistencyGroupSnapshotInlineMissingVolumes); i++ {
		if swag.IsZero(m.ConsistencyGroupSnapshotInlineMissingVolumes[i]) { // not required
			continue
		}

		if m.ConsistencyGroupSnapshotInlineMissingVolumes[i] != nil {
			if err := m.ConsistencyGroupSnapshotInlineMissingVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("missing_volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("missing_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsistencyGroupSnapshot) validateConsistencyGroupSnapshotInlineSnapshotVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistencyGroupSnapshotInlineSnapshotVolumes) { // not required
		return nil
	}

	for i := 0; i < len(m.ConsistencyGroupSnapshotInlineSnapshotVolumes); i++ {
		if swag.IsZero(m.ConsistencyGroupSnapshotInlineSnapshotVolumes[i]) { // not required
			continue
		}

		if m.ConsistencyGroupSnapshotInlineSnapshotVolumes[i] != nil {
			if err := m.ConsistencyGroupSnapshotInlineSnapshotVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshot_volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshot_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var consistencyGroupSnapshotTypeConsistencyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["crash","application"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consistencyGroupSnapshotTypeConsistencyTypePropEnum = append(consistencyGroupSnapshotTypeConsistencyTypePropEnum, v)
	}
}

const (

	// ConsistencyGroupSnapshotConsistencyTypeCrash captures enum value "crash"
	ConsistencyGroupSnapshotConsistencyTypeCrash string = "crash"

	// ConsistencyGroupSnapshotConsistencyTypeApplication captures enum value "application"
	ConsistencyGroupSnapshotConsistencyTypeApplication string = "application"
)

// prop value enum
func (m *ConsistencyGroupSnapshot) validateConsistencyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consistencyGroupSnapshotTypeConsistencyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsistencyGroupSnapshot) validateConsistencyType(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistencyType) { // not required
		return nil
	}

	// value enum
	if err := m.validateConsistencyTypeEnum("consistency_type", "body", *m.ConsistencyType); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupSnapshot) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupSnapshot) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this consistency group snapshot based on the context it is used
func (m *ConsistencyGroupSnapshot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsistencyGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsistencyGroupSnapshotInlineMissingVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsistencyGroupSnapshotInlineSnapshotVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsPartial(ctx, formats); err != nil {
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

func (m *ConsistencyGroupSnapshot) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupSnapshot) contextValidateConsistencyGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.ConsistencyGroup != nil {

		if swag.IsZero(m.ConsistencyGroup) { // not required
			return nil
		}

		if err := m.ConsistencyGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistency_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consistency_group")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupSnapshot) contextValidateConsistencyGroupSnapshotInlineMissingVolumes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "missing_volumes", "body", []*VolumeReference(m.ConsistencyGroupSnapshotInlineMissingVolumes)); err != nil {
		return err
	}

	for i := 0; i < len(m.ConsistencyGroupSnapshotInlineMissingVolumes); i++ {

		if m.ConsistencyGroupSnapshotInlineMissingVolumes[i] != nil {

			if swag.IsZero(m.ConsistencyGroupSnapshotInlineMissingVolumes[i]) { // not required
				return nil
			}

			if err := m.ConsistencyGroupSnapshotInlineMissingVolumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("missing_volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("missing_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsistencyGroupSnapshot) contextValidateConsistencyGroupSnapshotInlineSnapshotVolumes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "snapshot_volumes", "body", []*ConsistencyGroupVolumeSnapshot(m.ConsistencyGroupSnapshotInlineSnapshotVolumes)); err != nil {
		return err
	}

	for i := 0; i < len(m.ConsistencyGroupSnapshotInlineSnapshotVolumes); i++ {

		if m.ConsistencyGroupSnapshotInlineSnapshotVolumes[i] != nil {

			if swag.IsZero(m.ConsistencyGroupSnapshotInlineSnapshotVolumes[i]) { // not required
				return nil
			}

			if err := m.ConsistencyGroupSnapshotInlineSnapshotVolumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshot_volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshot_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsistencyGroupSnapshot) contextValidateCreateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "create_time", "body", m.CreateTime); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupSnapshot) contextValidateIsPartial(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_partial", "body", m.IsPartial); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupSnapshot) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupSnapshot) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupSnapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupSnapshot) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupSnapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsistencyGroupSnapshotInlineConsistencyGroup The consistency group of the Snapshot copy.
//
// swagger:model consistency_group_snapshot_inline_consistency_group
type ConsistencyGroupSnapshotInlineConsistencyGroup struct {

	// links
	Links *SelfLink `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the consistency group.
	// Example: my_consistency_group
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// The unique identifier of the consistency group.
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this consistency group snapshot inline consistency group
func (m *ConsistencyGroupSnapshotInlineConsistencyGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupSnapshotInlineConsistencyGroup) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistency_group" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consistency_group" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this consistency group snapshot inline consistency group based on the context it is used
func (m *ConsistencyGroupSnapshotInlineConsistencyGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupSnapshotInlineConsistencyGroup) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistency_group" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consistency_group" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupSnapshotInlineConsistencyGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupSnapshotInlineConsistencyGroup) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupSnapshotInlineConsistencyGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
