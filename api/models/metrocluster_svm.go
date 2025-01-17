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

// MetroclusterSvm Retrieves configuration information for all pairs of SVMs in MetroCluster. REST /api/cluster/metrocluster/svms/?
//
// swagger:model metrocluster_svm
type MetroclusterSvm struct {

	// links
	Links *SelfLink `json:"_links,omitempty" yaml:"_links,omitempty"`

	// cluster
	Cluster *MetroclusterSvmInlineCluster `json:"cluster,omitempty" yaml:"cluster,omitempty"`

	// Configuration state.
	// Read Only: true
	// Enum: ["degraded","healthy","pending_setup","pending_switchback","replication_paused","syncing","unhealthy"]
	ConfigurationState *string `json:"configuration_state,omitempty" yaml:"configuration_state,omitempty"`

	// Reason for SVM object replication failure.
	// Example: Apply failed for Object: volume Method: volume_method. Reason: invalid operation
	// Read Only: true
	FailedReason *Error `json:"failed_reason,omitempty" yaml:"failed_reason,omitempty"`

	// partner svm
	PartnerSvm *MetroclusterSvmInlinePartnerSvm `json:"partner_svm,omitempty" yaml:"partner_svm,omitempty"`

	// svm
	Svm *MetroclusterSvmInlineSvm `json:"svm,omitempty" yaml:"svm,omitempty"`
}

// Validate validates this metrocluster svm
func (m *MetroclusterSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigurationState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailedReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartnerSvm(formats); err != nil {
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

func (m *MetroclusterSvm) validateLinks(formats strfmt.Registry) error {
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

func (m *MetroclusterSvm) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

var metroclusterSvmTypeConfigurationStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["degraded","healthy","pending_setup","pending_switchback","replication_paused","syncing","unhealthy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterSvmTypeConfigurationStatePropEnum = append(metroclusterSvmTypeConfigurationStatePropEnum, v)
	}
}

const (

	// MetroclusterSvmConfigurationStateDegraded captures enum value "degraded"
	MetroclusterSvmConfigurationStateDegraded string = "degraded"

	// MetroclusterSvmConfigurationStateHealthy captures enum value "healthy"
	MetroclusterSvmConfigurationStateHealthy string = "healthy"

	// MetroclusterSvmConfigurationStatePendingSetup captures enum value "pending_setup"
	MetroclusterSvmConfigurationStatePendingSetup string = "pending_setup"

	// MetroclusterSvmConfigurationStatePendingSwitchback captures enum value "pending_switchback"
	MetroclusterSvmConfigurationStatePendingSwitchback string = "pending_switchback"

	// MetroclusterSvmConfigurationStateReplicationPaused captures enum value "replication_paused"
	MetroclusterSvmConfigurationStateReplicationPaused string = "replication_paused"

	// MetroclusterSvmConfigurationStateSyncing captures enum value "syncing"
	MetroclusterSvmConfigurationStateSyncing string = "syncing"

	// MetroclusterSvmConfigurationStateUnhealthy captures enum value "unhealthy"
	MetroclusterSvmConfigurationStateUnhealthy string = "unhealthy"
)

// prop value enum
func (m *MetroclusterSvm) validateConfigurationStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metroclusterSvmTypeConfigurationStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetroclusterSvm) validateConfigurationState(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigurationState) { // not required
		return nil
	}

	// value enum
	if err := m.validateConfigurationStateEnum("configuration_state", "body", *m.ConfigurationState); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterSvm) validateFailedReason(formats strfmt.Registry) error {
	if swag.IsZero(m.FailedReason) { // not required
		return nil
	}

	if m.FailedReason != nil {
		if err := m.FailedReason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("failed_reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("failed_reason")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterSvm) validatePartnerSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.PartnerSvm) { // not required
		return nil
	}

	if m.PartnerSvm != nil {
		if err := m.PartnerSvm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partner_svm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partner_svm")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterSvm) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this metrocluster svm based on the context it is used
func (m *MetroclusterSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfigurationState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFailedReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePartnerSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MetroclusterSvm) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {

		if swag.IsZero(m.Cluster) { // not required
			return nil
		}

		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterSvm) contextValidateConfigurationState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "configuration_state", "body", m.ConfigurationState); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterSvm) contextValidateFailedReason(ctx context.Context, formats strfmt.Registry) error {

	if m.FailedReason != nil {

		if swag.IsZero(m.FailedReason) { // not required
			return nil
		}

		if err := m.FailedReason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("failed_reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("failed_reason")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterSvm) contextValidatePartnerSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.PartnerSvm != nil {

		if swag.IsZero(m.PartnerSvm) { // not required
			return nil
		}

		if err := m.PartnerSvm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partner_svm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partner_svm")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterSvm) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *MetroclusterSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterSvm) UnmarshalBinary(b []byte) error {
	var res MetroclusterSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterSvmInlineCluster metrocluster svm inline cluster
//
// swagger:model metrocluster_svm_inline_cluster
type MetroclusterSvmInlineCluster struct {

	// links
	Links *MetroclusterSvmInlineClusterInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// name
	// Example: cluster1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Required: true
	// Format: uuid
	UUID *strfmt.UUID `json:"uuid" yaml:"uuid"`
}

// Validate validates this metrocluster svm inline cluster
func (m *MetroclusterSvmInlineCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterSvmInlineCluster) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterSvmInlineCluster) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("cluster"+"."+"uuid", "body", m.UUID); err != nil {
		return err
	}

	if err := validate.FormatOf("cluster"+"."+"uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this metrocluster svm inline cluster based on the context it is used
func (m *MetroclusterSvmInlineCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterSvmInlineCluster) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterSvmInlineCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterSvmInlineCluster) UnmarshalBinary(b []byte) error {
	var res MetroclusterSvmInlineCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterSvmInlineClusterInlineLinks metrocluster svm inline cluster inline links
//
// swagger:model metrocluster_svm_inline_cluster_inline__links
type MetroclusterSvmInlineClusterInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this metrocluster svm inline cluster inline links
func (m *MetroclusterSvmInlineClusterInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterSvmInlineClusterInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metrocluster svm inline cluster inline links based on the context it is used
func (m *MetroclusterSvmInlineClusterInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterSvmInlineClusterInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterSvmInlineClusterInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterSvmInlineClusterInlineLinks) UnmarshalBinary(b []byte) error {
	var res MetroclusterSvmInlineClusterInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterSvmInlinePartnerSvm metrocluster svm inline partner svm
//
// swagger:model metrocluster_svm_inline_partner_svm
type MetroclusterSvmInlinePartnerSvm struct {

	// MetroCluster partner SVM name.
	// Read Only: true
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// MetroCluster partner SVM UUID.
	// Read Only: true
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this metrocluster svm inline partner svm
func (m *MetroclusterSvmInlinePartnerSvm) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this metrocluster svm inline partner svm based on the context it is used
func (m *MetroclusterSvmInlinePartnerSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
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

func (m *MetroclusterSvmInlinePartnerSvm) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "partner_svm"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterSvmInlinePartnerSvm) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "partner_svm"+"."+"uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterSvmInlinePartnerSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterSvmInlinePartnerSvm) UnmarshalBinary(b []byte) error {
	var res MetroclusterSvmInlinePartnerSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterSvmInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model metrocluster_svm_inline_svm
type MetroclusterSvmInlineSvm struct {

	// links
	Links *MetroclusterSvmInlineSvmInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	// Required: true
	UUID *string `json:"uuid" yaml:"uuid"`
}

// Validate validates this metrocluster svm inline svm
func (m *MetroclusterSvmInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterSvmInlineSvm) validateLinks(formats strfmt.Registry) error {
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

func (m *MetroclusterSvmInlineSvm) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("svm"+"."+"uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this metrocluster svm inline svm based on the context it is used
func (m *MetroclusterSvmInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterSvmInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MetroclusterSvmInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterSvmInlineSvm) UnmarshalBinary(b []byte) error {
	var res MetroclusterSvmInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterSvmInlineSvmInlineLinks metrocluster svm inline svm inline links
//
// swagger:model metrocluster_svm_inline_svm_inline__links
type MetroclusterSvmInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this metrocluster svm inline svm inline links
func (m *MetroclusterSvmInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterSvmInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this metrocluster svm inline svm inline links based on the context it is used
func (m *MetroclusterSvmInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterSvmInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MetroclusterSvmInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterSvmInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res MetroclusterSvmInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
