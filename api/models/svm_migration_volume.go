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

// SvmMigrationVolume Volume transfer information
//
// swagger:model svm_migration_volume
type SvmMigrationVolume struct {

	// links
	Links *SelfLink `json:"_links,omitempty" yaml:"_links,omitempty"`

	// Indicates whether the volume transfer relationship is healthy.
	Healthy *bool `json:"healthy,omitempty" yaml:"healthy,omitempty"`

	// node
	Node *SvmMigrationVolumeInlineNode `json:"node,omitempty" yaml:"node,omitempty"`

	// svm
	Svm *SvmMigrationVolumeInlineSvm `json:"svm,omitempty" yaml:"svm,omitempty"`

	// List of transfer errors
	// Read Only: true
	SvmMigrationVolumeInlineErrors []*SvmMigrationVolumeInlineErrorsInlineArrayItem `json:"errors,omitempty" yaml:"errors,omitempty"`

	// Status of the transfer.
	// Read Only: true
	// Enum: ["Idle","Transferring","Aborting","OutOfSync","InSync","Transitioning","ReadyForCutoverPreCommit","CutoverPreCommitting","CuttingOver"]
	TransferState *string `json:"transfer_state,omitempty" yaml:"transfer_state,omitempty"`

	// volume
	Volume *SvmMigrationVolumeInlineVolume `json:"volume,omitempty" yaml:"volume,omitempty"`
}

// Validate validates this svm migration volume
func (m *SvmMigrationVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvmMigrationVolumeInlineErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransferState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolume) validateLinks(formats strfmt.Registry) error {
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

func (m *SvmMigrationVolume) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *SvmMigrationVolume) validateSvm(formats strfmt.Registry) error {
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

func (m *SvmMigrationVolume) validateSvmMigrationVolumeInlineErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.SvmMigrationVolumeInlineErrors) { // not required
		return nil
	}

	for i := 0; i < len(m.SvmMigrationVolumeInlineErrors); i++ {
		if swag.IsZero(m.SvmMigrationVolumeInlineErrors[i]) { // not required
			continue
		}

		if m.SvmMigrationVolumeInlineErrors[i] != nil {
			if err := m.SvmMigrationVolumeInlineErrors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var svmMigrationVolumeTypeTransferStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Idle","Transferring","Aborting","OutOfSync","InSync","Transitioning","ReadyForCutoverPreCommit","CutoverPreCommitting","CuttingOver"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		svmMigrationVolumeTypeTransferStatePropEnum = append(svmMigrationVolumeTypeTransferStatePropEnum, v)
	}
}

const (

	// SvmMigrationVolumeTransferStateIdle captures enum value "Idle"
	SvmMigrationVolumeTransferStateIdle string = "Idle"

	// SvmMigrationVolumeTransferStateTransferring captures enum value "Transferring"
	SvmMigrationVolumeTransferStateTransferring string = "Transferring"

	// SvmMigrationVolumeTransferStateAborting captures enum value "Aborting"
	SvmMigrationVolumeTransferStateAborting string = "Aborting"

	// SvmMigrationVolumeTransferStateOutOfSync captures enum value "OutOfSync"
	SvmMigrationVolumeTransferStateOutOfSync string = "OutOfSync"

	// SvmMigrationVolumeTransferStateInSync captures enum value "InSync"
	SvmMigrationVolumeTransferStateInSync string = "InSync"

	// SvmMigrationVolumeTransferStateTransitioning captures enum value "Transitioning"
	SvmMigrationVolumeTransferStateTransitioning string = "Transitioning"

	// SvmMigrationVolumeTransferStateReadyForCutoverPreCommit captures enum value "ReadyForCutoverPreCommit"
	SvmMigrationVolumeTransferStateReadyForCutoverPreCommit string = "ReadyForCutoverPreCommit"

	// SvmMigrationVolumeTransferStateCutoverPreCommitting captures enum value "CutoverPreCommitting"
	SvmMigrationVolumeTransferStateCutoverPreCommitting string = "CutoverPreCommitting"

	// SvmMigrationVolumeTransferStateCuttingOver captures enum value "CuttingOver"
	SvmMigrationVolumeTransferStateCuttingOver string = "CuttingOver"
)

// prop value enum
func (m *SvmMigrationVolume) validateTransferStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, svmMigrationVolumeTypeTransferStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SvmMigrationVolume) validateTransferState(formats strfmt.Registry) error {
	if swag.IsZero(m.TransferState) { // not required
		return nil
	}

	// value enum
	if err := m.validateTransferStateEnum("transfer_state", "body", *m.TransferState); err != nil {
		return err
	}

	return nil
}

func (m *SvmMigrationVolume) validateVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration volume based on the context it is used
func (m *SvmMigrationVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvmMigrationVolumeInlineErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransferState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolume) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SvmMigrationVolume) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {

		if swag.IsZero(m.Node) { // not required
			return nil
		}

		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *SvmMigrationVolume) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SvmMigrationVolume) contextValidateSvmMigrationVolumeInlineErrors(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "errors", "body", []*SvmMigrationVolumeInlineErrorsInlineArrayItem(m.SvmMigrationVolumeInlineErrors)); err != nil {
		return err
	}

	for i := 0; i < len(m.SvmMigrationVolumeInlineErrors); i++ {

		if m.SvmMigrationVolumeInlineErrors[i] != nil {

			if swag.IsZero(m.SvmMigrationVolumeInlineErrors[i]) { // not required
				return nil
			}

			if err := m.SvmMigrationVolumeInlineErrors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SvmMigrationVolume) contextValidateTransferState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "transfer_state", "body", m.TransferState); err != nil {
		return err
	}

	return nil
}

func (m *SvmMigrationVolume) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.Volume != nil {

		if swag.IsZero(m.Volume) { // not required
			return nil
		}

		if err := m.Volume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolume) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumeInlineErrorsInlineArrayItem svm migration volume inline errors inline array item
//
// swagger:model svm_migration_volume_inline_errors_inline_array_item
type SvmMigrationVolumeInlineErrorsInlineArrayItem struct {

	// Argument code
	// Read Only: true
	Code *string `json:"code,omitempty" yaml:"code,omitempty"`

	// Message argument
	// Read Only: true
	Message *string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Validate validates this svm migration volume inline errors inline array item
func (m *SvmMigrationVolumeInlineErrorsInlineArrayItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this svm migration volume inline errors inline array item based on the context it is used
func (m *SvmMigrationVolumeInlineErrorsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumeInlineErrorsInlineArrayItem) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *SvmMigrationVolumeInlineErrorsInlineArrayItem) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationVolumeInlineErrorsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumeInlineErrorsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumeInlineErrorsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumeInlineNode Node in the destination cluster where the volume is hosted
//
// swagger:model svm_migration_volume_inline_node
type SvmMigrationVolumeInlineNode struct {

	// links
	Links *SvmMigrationVolumeInlineNodeInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this svm migration volume inline node
func (m *SvmMigrationVolumeInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumeInlineNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration volume inline node based on the context it is used
func (m *SvmMigrationVolumeInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumeInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationVolumeInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumeInlineNode) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumeInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumeInlineNodeInlineLinks svm migration volume inline node inline links
//
// swagger:model svm_migration_volume_inline_node_inline__links
type SvmMigrationVolumeInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this svm migration volume inline node inline links
func (m *SvmMigrationVolumeInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumeInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration volume inline node inline links based on the context it is used
func (m *SvmMigrationVolumeInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumeInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationVolumeInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumeInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumeInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumeInlineSvm SVM information
//
// swagger:model svm_migration_volume_inline_svm
type SvmMigrationVolumeInlineSvm struct {

	// links
	Links *SvmMigrationVolumeInlineSvmInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this svm migration volume inline svm
func (m *SvmMigrationVolumeInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumeInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this svm migration volume inline svm based on the context it is used
func (m *SvmMigrationVolumeInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumeInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SvmMigrationVolumeInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumeInlineSvm) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumeInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumeInlineSvmInlineLinks svm migration volume inline svm inline links
//
// swagger:model svm_migration_volume_inline_svm_inline__links
type SvmMigrationVolumeInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this svm migration volume inline svm inline links
func (m *SvmMigrationVolumeInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumeInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this svm migration volume inline svm inline links based on the context it is used
func (m *SvmMigrationVolumeInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumeInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SvmMigrationVolumeInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumeInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumeInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumeInlineVolume Volume information in the destination cluster
//
// swagger:model svm_migration_volume_inline_volume
type SvmMigrationVolumeInlineVolume struct {

	// links
	Links *SvmMigrationVolumeInlineVolumeInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the volume. This field cannot be specified in a POST or PATCH method.
	// Example: volume1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	// Example: 028baa66-41bd-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this svm migration volume inline volume
func (m *SvmMigrationVolumeInlineVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumeInlineVolume) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration volume inline volume based on the context it is used
func (m *SvmMigrationVolumeInlineVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumeInlineVolume) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationVolumeInlineVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumeInlineVolume) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumeInlineVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumeInlineVolumeInlineLinks svm migration volume inline volume inline links
//
// swagger:model svm_migration_volume_inline_volume_inline__links
type SvmMigrationVolumeInlineVolumeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this svm migration volume inline volume inline links
func (m *SvmMigrationVolumeInlineVolumeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumeInlineVolumeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration volume inline volume inline links based on the context it is used
func (m *SvmMigrationVolumeInlineVolumeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumeInlineVolumeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationVolumeInlineVolumeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumeInlineVolumeInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumeInlineVolumeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}