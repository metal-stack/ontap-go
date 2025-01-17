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

// GroupPolicyObjectCentralAccessRule group policy object central access rule
//
// swagger:model group_policy_object_central_access_rule
type GroupPolicyObjectCentralAccessRule struct {

	// Policy creation timestamp.
	// Example: 2018-01-01 16:00:00
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"create_time,omitempty" yaml:"create_time,omitempty"`

	// Effective security policy in security descriptor definition language format.
	// Example: O:SYG:SYD:AR(A;;FA;;;WD)
	CurrentPermission *string `json:"current_permission,omitempty" yaml:"current_permission,omitempty"`

	// Description about the policy.
	// Example: rule #1
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`

	// name
	// Example: p1
	// Min Length: 1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// Proposed security policy in security descriptor definition language format.
	// Example: O:SYG:SYD:(A;;FA;;;OW)(A;;FA;;;BA)(A;;FA;;;SY)
	ProposedPermission *string `json:"proposed_permission,omitempty" yaml:"proposed_permission,omitempty"`

	// Criteria to scope resources for which access rules apply.
	// Example: department
	ResourceCriteria *string `json:"resource_criteria,omitempty" yaml:"resource_criteria,omitempty"`

	// svm
	Svm *GroupPolicyObjectCentralAccessRuleInlineSvm `json:"svm,omitempty" yaml:"svm,omitempty"`

	// Last policy modification timestamp.
	// Example: 2018-01-01 16:00:00
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty" yaml:"update_time,omitempty"`
}

// Validate validates this group policy object central access rule
func (m *GroupPolicyObjectCentralAccessRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectCentralAccessRule) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GroupPolicyObjectCentralAccessRule) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *GroupPolicyObjectCentralAccessRule) validateSvm(formats strfmt.Registry) error {
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

func (m *GroupPolicyObjectCentralAccessRule) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this group policy object central access rule based on the context it is used
func (m *GroupPolicyObjectCentralAccessRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectCentralAccessRule) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GroupPolicyObjectCentralAccessRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupPolicyObjectCentralAccessRule) UnmarshalBinary(b []byte) error {
	var res GroupPolicyObjectCentralAccessRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GroupPolicyObjectCentralAccessRuleInlineSvm Will not be populated for objects that are yet to be applied.
//
// swagger:model group_policy_object_central_access_rule_inline_svm
type GroupPolicyObjectCentralAccessRuleInlineSvm struct {

	// links
	Links *GroupPolicyObjectCentralAccessRuleInlineSvmInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this group policy object central access rule inline svm
func (m *GroupPolicyObjectCentralAccessRuleInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectCentralAccessRuleInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this group policy object central access rule inline svm based on the context it is used
func (m *GroupPolicyObjectCentralAccessRuleInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectCentralAccessRuleInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GroupPolicyObjectCentralAccessRuleInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupPolicyObjectCentralAccessRuleInlineSvm) UnmarshalBinary(b []byte) error {
	var res GroupPolicyObjectCentralAccessRuleInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GroupPolicyObjectCentralAccessRuleInlineSvmInlineLinks group policy object central access rule inline svm inline links
//
// swagger:model group_policy_object_central_access_rule_inline_svm_inline__links
type GroupPolicyObjectCentralAccessRuleInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this group policy object central access rule inline svm inline links
func (m *GroupPolicyObjectCentralAccessRuleInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectCentralAccessRuleInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this group policy object central access rule inline svm inline links based on the context it is used
func (m *GroupPolicyObjectCentralAccessRuleInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectCentralAccessRuleInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GroupPolicyObjectCentralAccessRuleInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupPolicyObjectCentralAccessRuleInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res GroupPolicyObjectCentralAccessRuleInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
