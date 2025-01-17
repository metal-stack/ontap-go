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

// ClusterSpace Provides information on cluster-wide storage details across the different tiers. Storage details include storage efficiency, block storage and cloud storage information.
//
// swagger:model cluster_space
type ClusterSpace struct {

	// block storage
	BlockStorage *ClusterSpaceInlineBlockStorage `json:"block_storage,omitempty" yaml:"block_storage,omitempty"`

	// cloud storage
	CloudStorage *ClusterSpaceInlineCloudStorage `json:"cloud_storage,omitempty" yaml:"cloud_storage,omitempty"`

	// efficiency
	Efficiency *ClusterSpaceInlineEfficiency `json:"efficiency,omitempty" yaml:"efficiency,omitempty"`

	// efficiency without snapshots
	EfficiencyWithoutSnapshots *ClusterSpaceInlineEfficiencyWithoutSnapshots `json:"efficiency_without_snapshots,omitempty" yaml:"efficiency_without_snapshots,omitempty"`

	// efficiency without snapshots flexclones
	EfficiencyWithoutSnapshotsFlexclones *ClusterSpaceInlineEfficiencyWithoutSnapshotsFlexclones `json:"efficiency_without_snapshots_flexclones,omitempty" yaml:"efficiency_without_snapshots_flexclones,omitempty"`
}

// Validate validates this cluster space
func (m *ClusterSpace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEfficiency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEfficiencyWithoutSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEfficiencyWithoutSnapshotsFlexclones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpace) validateBlockStorage(formats strfmt.Registry) error {
	if swag.IsZero(m.BlockStorage) { // not required
		return nil
	}

	if m.BlockStorage != nil {
		if err := m.BlockStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block_storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block_storage")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpace) validateCloudStorage(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudStorage) { // not required
		return nil
	}

	if m.CloudStorage != nil {
		if err := m.CloudStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud_storage")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpace) validateEfficiency(formats strfmt.Registry) error {
	if swag.IsZero(m.Efficiency) { // not required
		return nil
	}

	if m.Efficiency != nil {
		if err := m.Efficiency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("efficiency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("efficiency")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpace) validateEfficiencyWithoutSnapshots(formats strfmt.Registry) error {
	if swag.IsZero(m.EfficiencyWithoutSnapshots) { // not required
		return nil
	}

	if m.EfficiencyWithoutSnapshots != nil {
		if err := m.EfficiencyWithoutSnapshots.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("efficiency_without_snapshots")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("efficiency_without_snapshots")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpace) validateEfficiencyWithoutSnapshotsFlexclones(formats strfmt.Registry) error {
	if swag.IsZero(m.EfficiencyWithoutSnapshotsFlexclones) { // not required
		return nil
	}

	if m.EfficiencyWithoutSnapshotsFlexclones != nil {
		if err := m.EfficiencyWithoutSnapshotsFlexclones.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("efficiency_without_snapshots_flexclones")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("efficiency_without_snapshots_flexclones")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster space based on the context it is used
func (m *ClusterSpace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlockStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCloudStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEfficiency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEfficiencyWithoutSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEfficiencyWithoutSnapshotsFlexclones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpace) contextValidateBlockStorage(ctx context.Context, formats strfmt.Registry) error {

	if m.BlockStorage != nil {

		if swag.IsZero(m.BlockStorage) { // not required
			return nil
		}

		if err := m.BlockStorage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block_storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block_storage")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpace) contextValidateCloudStorage(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudStorage != nil {

		if swag.IsZero(m.CloudStorage) { // not required
			return nil
		}

		if err := m.CloudStorage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_storage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud_storage")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpace) contextValidateEfficiency(ctx context.Context, formats strfmt.Registry) error {

	if m.Efficiency != nil {

		if swag.IsZero(m.Efficiency) { // not required
			return nil
		}

		if err := m.Efficiency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("efficiency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("efficiency")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpace) contextValidateEfficiencyWithoutSnapshots(ctx context.Context, formats strfmt.Registry) error {

	if m.EfficiencyWithoutSnapshots != nil {

		if swag.IsZero(m.EfficiencyWithoutSnapshots) { // not required
			return nil
		}

		if err := m.EfficiencyWithoutSnapshots.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("efficiency_without_snapshots")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("efficiency_without_snapshots")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpace) contextValidateEfficiencyWithoutSnapshotsFlexclones(ctx context.Context, formats strfmt.Registry) error {

	if m.EfficiencyWithoutSnapshotsFlexclones != nil {

		if swag.IsZero(m.EfficiencyWithoutSnapshotsFlexclones) { // not required
			return nil
		}

		if err := m.EfficiencyWithoutSnapshotsFlexclones.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("efficiency_without_snapshots_flexclones")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("efficiency_without_snapshots_flexclones")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSpace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSpace) UnmarshalBinary(b []byte) error {
	var res ClusterSpace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterSpaceInlineBlockStorage Configuration information for the locally attached portion of the storage across the cluster. When a cloud store is also used by the storage, this is referred to as the performance tier.
//
// swagger:model cluster_space_inline_block_storage
type ClusterSpaceInlineBlockStorage struct {

	// Available space across the cluster.
	Available *int64 `json:"available,omitempty" yaml:"available,omitempty"`

	// Inactive data across the cluster.
	//
	//
	InactiveData *int64 `json:"inactive_data,omitempty" yaml:"inactive_data,omitempty"`

	// Configuration information based on type of media. For example, SSD media type information includes the sum of all the SSD storage across the cluster.
	Medias []*ClusterSpaceBlockStorageMediasItems0 `json:"medias,omitempty" yaml:"medias,omitempty"`

	// Total physical used space across the cluster.
	PhysicalUsed *int64 `json:"physical_used,omitempty" yaml:"physical_used,omitempty"`

	// Total space across the cluster.
	Size *int64 `json:"size,omitempty" yaml:"size,omitempty"`

	// Used space (includes volume reserves) across the cluster.
	//
	//
	Used *int64 `json:"used,omitempty" yaml:"used,omitempty"`
}

// Validate validates this cluster space inline block storage
func (m *ClusterSpaceInlineBlockStorage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMedias(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpaceInlineBlockStorage) validateMedias(formats strfmt.Registry) error {
	if swag.IsZero(m.Medias) { // not required
		return nil
	}

	for i := 0; i < len(m.Medias); i++ {
		if swag.IsZero(m.Medias[i]) { // not required
			continue
		}

		if m.Medias[i] != nil {
			if err := m.Medias[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("block_storage" + "." + "medias" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("block_storage" + "." + "medias" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cluster space inline block storage based on the context it is used
func (m *ClusterSpaceInlineBlockStorage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMedias(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpaceInlineBlockStorage) contextValidateMedias(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Medias); i++ {

		if m.Medias[i] != nil {

			if swag.IsZero(m.Medias[i]) { // not required
				return nil
			}

			if err := m.Medias[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("block_storage" + "." + "medias" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("block_storage" + "." + "medias" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSpaceInlineBlockStorage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSpaceInlineBlockStorage) UnmarshalBinary(b []byte) error {
	var res ClusterSpaceInlineBlockStorage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterSpaceBlockStorageMediasItems0 cluster space block storage medias items0
//
// swagger:model ClusterSpaceBlockStorageMediasItems0
type ClusterSpaceBlockStorageMediasItems0 struct {

	// Available space across the cluster based on media type.
	Available *int64 `json:"available,omitempty" yaml:"available,omitempty"`

	// efficiency
	Efficiency *ClusterSpaceBlockStorageMediasItems0Efficiency `json:"efficiency,omitempty" yaml:"efficiency,omitempty"`

	// efficiency without snapshots
	EfficiencyWithoutSnapshots *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshots `json:"efficiency_without_snapshots,omitempty" yaml:"efficiency_without_snapshots,omitempty"`

	// efficiency without snapshots flexclones
	EfficiencyWithoutSnapshotsFlexclones *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshotsFlexclones `json:"efficiency_without_snapshots_flexclones,omitempty" yaml:"efficiency_without_snapshots_flexclones,omitempty"`

	// Total physical used space across the cluster based on media type.
	PhysicalUsed *int64 `json:"physical_used,omitempty" yaml:"physical_used,omitempty"`

	// Total space across the cluster based on media type.
	Size *int64 `json:"size,omitempty" yaml:"size,omitempty"`

	// The type of media being used.
	// Enum: ["hdd","hybrid","lun","ssd","vmdisk"]
	Type *string `json:"type,omitempty" yaml:"type,omitempty"`

	// Used space across the cluster based on media type.
	//
	//
	Used *int64 `json:"used,omitempty" yaml:"used,omitempty"`
}

// Validate validates this cluster space block storage medias items0
func (m *ClusterSpaceBlockStorageMediasItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEfficiency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEfficiencyWithoutSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEfficiencyWithoutSnapshotsFlexclones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0) validateEfficiency(formats strfmt.Registry) error {
	if swag.IsZero(m.Efficiency) { // not required
		return nil
	}

	if m.Efficiency != nil {
		if err := m.Efficiency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("efficiency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("efficiency")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0) validateEfficiencyWithoutSnapshots(formats strfmt.Registry) error {
	if swag.IsZero(m.EfficiencyWithoutSnapshots) { // not required
		return nil
	}

	if m.EfficiencyWithoutSnapshots != nil {
		if err := m.EfficiencyWithoutSnapshots.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("efficiency_without_snapshots")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("efficiency_without_snapshots")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0) validateEfficiencyWithoutSnapshotsFlexclones(formats strfmt.Registry) error {
	if swag.IsZero(m.EfficiencyWithoutSnapshotsFlexclones) { // not required
		return nil
	}

	if m.EfficiencyWithoutSnapshotsFlexclones != nil {
		if err := m.EfficiencyWithoutSnapshotsFlexclones.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("efficiency_without_snapshots_flexclones")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("efficiency_without_snapshots_flexclones")
			}
			return err
		}
	}

	return nil
}

var clusterSpaceBlockStorageMediasItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hdd","hybrid","lun","ssd","vmdisk"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterSpaceBlockStorageMediasItems0TypeTypePropEnum = append(clusterSpaceBlockStorageMediasItems0TypeTypePropEnum, v)
	}
}

const (

	// ClusterSpaceBlockStorageMediasItems0TypeHdd captures enum value "hdd"
	ClusterSpaceBlockStorageMediasItems0TypeHdd string = "hdd"

	// ClusterSpaceBlockStorageMediasItems0TypeHybrid captures enum value "hybrid"
	ClusterSpaceBlockStorageMediasItems0TypeHybrid string = "hybrid"

	// ClusterSpaceBlockStorageMediasItems0TypeLun captures enum value "lun"
	ClusterSpaceBlockStorageMediasItems0TypeLun string = "lun"

	// ClusterSpaceBlockStorageMediasItems0TypeSsd captures enum value "ssd"
	ClusterSpaceBlockStorageMediasItems0TypeSsd string = "ssd"

	// ClusterSpaceBlockStorageMediasItems0TypeVmdisk captures enum value "vmdisk"
	ClusterSpaceBlockStorageMediasItems0TypeVmdisk string = "vmdisk"
)

// prop value enum
func (m *ClusterSpaceBlockStorageMediasItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterSpaceBlockStorageMediasItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cluster space block storage medias items0 based on the context it is used
func (m *ClusterSpaceBlockStorageMediasItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEfficiency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEfficiencyWithoutSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEfficiencyWithoutSnapshotsFlexclones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0) contextValidateEfficiency(ctx context.Context, formats strfmt.Registry) error {

	if m.Efficiency != nil {

		if swag.IsZero(m.Efficiency) { // not required
			return nil
		}

		if err := m.Efficiency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("efficiency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("efficiency")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0) contextValidateEfficiencyWithoutSnapshots(ctx context.Context, formats strfmt.Registry) error {

	if m.EfficiencyWithoutSnapshots != nil {

		if swag.IsZero(m.EfficiencyWithoutSnapshots) { // not required
			return nil
		}

		if err := m.EfficiencyWithoutSnapshots.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("efficiency_without_snapshots")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("efficiency_without_snapshots")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0) contextValidateEfficiencyWithoutSnapshotsFlexclones(ctx context.Context, formats strfmt.Registry) error {

	if m.EfficiencyWithoutSnapshotsFlexclones != nil {

		if swag.IsZero(m.EfficiencyWithoutSnapshotsFlexclones) { // not required
			return nil
		}

		if err := m.EfficiencyWithoutSnapshotsFlexclones.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("efficiency_without_snapshots_flexclones")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("efficiency_without_snapshots_flexclones")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSpaceBlockStorageMediasItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSpaceBlockStorageMediasItems0) UnmarshalBinary(b []byte) error {
	var res ClusterSpaceBlockStorageMediasItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterSpaceBlockStorageMediasItems0Efficiency Storage efficiency.
//
// swagger:model ClusterSpaceBlockStorageMediasItems0Efficiency
type ClusterSpaceBlockStorageMediasItems0Efficiency struct {

	// Logical used
	// Read Only: true
	LogicalUsed *int64 `json:"logical_used,omitempty" yaml:"logical_used,omitempty"`

	// Data reduction ratio (logical_used / used)
	// Read Only: true
	Ratio *float64 `json:"ratio,omitempty" yaml:"ratio,omitempty"`

	// Space saved by storage efficiencies (logical_used - used)
	// Read Only: true
	Savings *int64 `json:"savings,omitempty" yaml:"savings,omitempty"`
}

// Validate validates this cluster space block storage medias items0 efficiency
func (m *ClusterSpaceBlockStorageMediasItems0Efficiency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this cluster space block storage medias items0 efficiency based on the context it is used
func (m *ClusterSpaceBlockStorageMediasItems0Efficiency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogicalUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRatio(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSavings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0Efficiency) contextValidateLogicalUsed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency"+"."+"logical_used", "body", m.LogicalUsed); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0Efficiency) contextValidateRatio(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency"+"."+"ratio", "body", m.Ratio); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0Efficiency) contextValidateSavings(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency"+"."+"savings", "body", m.Savings); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSpaceBlockStorageMediasItems0Efficiency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSpaceBlockStorageMediasItems0Efficiency) UnmarshalBinary(b []byte) error {
	var res ClusterSpaceBlockStorageMediasItems0Efficiency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshots Storage efficiency that does not include the savings provided by Snapshot copies.
//
// swagger:model ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshots
type ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshots struct {

	// Logical used
	// Read Only: true
	LogicalUsed *int64 `json:"logical_used,omitempty" yaml:"logical_used,omitempty"`

	// Data reduction ratio (logical_used / used)
	// Read Only: true
	Ratio *float64 `json:"ratio,omitempty" yaml:"ratio,omitempty"`

	// Space saved by storage efficiencies (logical_used - used)
	// Read Only: true
	Savings *int64 `json:"savings,omitempty" yaml:"savings,omitempty"`
}

// Validate validates this cluster space block storage medias items0 efficiency without snapshots
func (m *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshots) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this cluster space block storage medias items0 efficiency without snapshots based on the context it is used
func (m *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshots) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogicalUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRatio(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSavings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshots) contextValidateLogicalUsed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency_without_snapshots"+"."+"logical_used", "body", m.LogicalUsed); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshots) contextValidateRatio(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency_without_snapshots"+"."+"ratio", "body", m.Ratio); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshots) contextValidateSavings(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency_without_snapshots"+"."+"savings", "body", m.Savings); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshots) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshots) UnmarshalBinary(b []byte) error {
	var res ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshots
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshotsFlexclones Storage efficiency that does not include the savings provided by Snapshot copies and FlexClone volumes.
//
// swagger:model ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshotsFlexclones
type ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshotsFlexclones struct {

	// Logical used
	// Read Only: true
	LogicalUsed *int64 `json:"logical_used,omitempty" yaml:"logical_used,omitempty"`

	// Data reduction ratio (logical_used / used)
	// Read Only: true
	Ratio *float64 `json:"ratio,omitempty" yaml:"ratio,omitempty"`

	// Space saved by storage efficiencies (logical_used - used)
	// Read Only: true
	Savings *int64 `json:"savings,omitempty" yaml:"savings,omitempty"`
}

// Validate validates this cluster space block storage medias items0 efficiency without snapshots flexclones
func (m *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshotsFlexclones) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this cluster space block storage medias items0 efficiency without snapshots flexclones based on the context it is used
func (m *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshotsFlexclones) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogicalUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRatio(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSavings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshotsFlexclones) contextValidateLogicalUsed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency_without_snapshots_flexclones"+"."+"logical_used", "body", m.LogicalUsed); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshotsFlexclones) contextValidateRatio(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency_without_snapshots_flexclones"+"."+"ratio", "body", m.Ratio); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshotsFlexclones) contextValidateSavings(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency_without_snapshots_flexclones"+"."+"savings", "body", m.Savings); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshotsFlexclones) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshotsFlexclones) UnmarshalBinary(b []byte) error {
	var res ClusterSpaceBlockStorageMediasItems0EfficiencyWithoutSnapshotsFlexclones
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterSpaceInlineCloudStorage Configuration information for the cloud storage portion of all the aggregates across the cluster. This is referred to as the capacity tier.
//
// swagger:model cluster_space_inline_cloud_storage
type ClusterSpaceInlineCloudStorage struct {

	// Total space used in cloud.
	// Read Only: true
	Used *int64 `json:"used,omitempty" yaml:"used,omitempty"`
}

// Validate validates this cluster space inline cloud storage
func (m *ClusterSpaceInlineCloudStorage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this cluster space inline cloud storage based on the context it is used
func (m *ClusterSpaceInlineCloudStorage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpaceInlineCloudStorage) contextValidateUsed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cloud_storage"+"."+"used", "body", m.Used); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSpaceInlineCloudStorage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSpaceInlineCloudStorage) UnmarshalBinary(b []byte) error {
	var res ClusterSpaceInlineCloudStorage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterSpaceInlineEfficiency Storage efficiency.
//
// swagger:model cluster_space_inline_efficiency
type ClusterSpaceInlineEfficiency struct {

	// Logical used
	// Read Only: true
	LogicalUsed *int64 `json:"logical_used,omitempty" yaml:"logical_used,omitempty"`

	// Data reduction ratio (logical_used / used)
	// Read Only: true
	Ratio *float64 `json:"ratio,omitempty" yaml:"ratio,omitempty"`

	// Space saved by storage efficiencies (logical_used - used)
	// Read Only: true
	Savings *int64 `json:"savings,omitempty" yaml:"savings,omitempty"`
}

// Validate validates this cluster space inline efficiency
func (m *ClusterSpaceInlineEfficiency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this cluster space inline efficiency based on the context it is used
func (m *ClusterSpaceInlineEfficiency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogicalUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRatio(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSavings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpaceInlineEfficiency) contextValidateLogicalUsed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency"+"."+"logical_used", "body", m.LogicalUsed); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSpaceInlineEfficiency) contextValidateRatio(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency"+"."+"ratio", "body", m.Ratio); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSpaceInlineEfficiency) contextValidateSavings(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency"+"."+"savings", "body", m.Savings); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSpaceInlineEfficiency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSpaceInlineEfficiency) UnmarshalBinary(b []byte) error {
	var res ClusterSpaceInlineEfficiency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterSpaceInlineEfficiencyWithoutSnapshots Storage efficiency that does not include the savings provided by Snapshot copies.
//
// swagger:model cluster_space_inline_efficiency_without_snapshots
type ClusterSpaceInlineEfficiencyWithoutSnapshots struct {

	// Logical used
	// Read Only: true
	LogicalUsed *int64 `json:"logical_used,omitempty" yaml:"logical_used,omitempty"`

	// Data reduction ratio (logical_used / used)
	// Read Only: true
	Ratio *float64 `json:"ratio,omitempty" yaml:"ratio,omitempty"`

	// Space saved by storage efficiencies (logical_used - used)
	// Read Only: true
	Savings *int64 `json:"savings,omitempty" yaml:"savings,omitempty"`
}

// Validate validates this cluster space inline efficiency without snapshots
func (m *ClusterSpaceInlineEfficiencyWithoutSnapshots) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this cluster space inline efficiency without snapshots based on the context it is used
func (m *ClusterSpaceInlineEfficiencyWithoutSnapshots) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogicalUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRatio(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSavings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpaceInlineEfficiencyWithoutSnapshots) contextValidateLogicalUsed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency_without_snapshots"+"."+"logical_used", "body", m.LogicalUsed); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSpaceInlineEfficiencyWithoutSnapshots) contextValidateRatio(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency_without_snapshots"+"."+"ratio", "body", m.Ratio); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSpaceInlineEfficiencyWithoutSnapshots) contextValidateSavings(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency_without_snapshots"+"."+"savings", "body", m.Savings); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSpaceInlineEfficiencyWithoutSnapshots) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSpaceInlineEfficiencyWithoutSnapshots) UnmarshalBinary(b []byte) error {
	var res ClusterSpaceInlineEfficiencyWithoutSnapshots
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterSpaceInlineEfficiencyWithoutSnapshotsFlexclones Storage efficiency that does not include the savings provided by Snapshot copies and FlexClone volumes.
//
// swagger:model cluster_space_inline_efficiency_without_snapshots_flexclones
type ClusterSpaceInlineEfficiencyWithoutSnapshotsFlexclones struct {

	// Logical used
	// Read Only: true
	LogicalUsed *int64 `json:"logical_used,omitempty" yaml:"logical_used,omitempty"`

	// Data reduction ratio (logical_used / used)
	// Read Only: true
	Ratio *float64 `json:"ratio,omitempty" yaml:"ratio,omitempty"`

	// Space saved by storage efficiencies (logical_used - used)
	// Read Only: true
	Savings *int64 `json:"savings,omitempty" yaml:"savings,omitempty"`
}

// Validate validates this cluster space inline efficiency without snapshots flexclones
func (m *ClusterSpaceInlineEfficiencyWithoutSnapshotsFlexclones) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this cluster space inline efficiency without snapshots flexclones based on the context it is used
func (m *ClusterSpaceInlineEfficiencyWithoutSnapshotsFlexclones) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogicalUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRatio(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSavings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpaceInlineEfficiencyWithoutSnapshotsFlexclones) contextValidateLogicalUsed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency_without_snapshots_flexclones"+"."+"logical_used", "body", m.LogicalUsed); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSpaceInlineEfficiencyWithoutSnapshotsFlexclones) contextValidateRatio(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency_without_snapshots_flexclones"+"."+"ratio", "body", m.Ratio); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSpaceInlineEfficiencyWithoutSnapshotsFlexclones) contextValidateSavings(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "efficiency_without_snapshots_flexclones"+"."+"savings", "body", m.Savings); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSpaceInlineEfficiencyWithoutSnapshotsFlexclones) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSpaceInlineEfficiencyWithoutSnapshotsFlexclones) UnmarshalBinary(b []byte) error {
	var res ClusterSpaceInlineEfficiencyWithoutSnapshotsFlexclones
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
