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

// IPServicePolicy ip service policy
//
// swagger:model ip_service_policy
type IPServicePolicy struct {

	// links
	Links *IPServicePolicyInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// ip service policy inline services
	IPServicePolicyInlineServices []*IPService `json:"services,omitempty" yaml:"services,omitempty"`

	// ipspace
	Ipspace *IPServicePolicyInlineIpspace `json:"ipspace,omitempty" yaml:"ipspace,omitempty"`

	// is built in
	// Read Only: true
	IsBuiltIn *bool `json:"is_built_in,omitempty" yaml:"is_built_in,omitempty"`

	// name
	// Example: default-intercluster
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set to "svm" for interfaces owned by an SVM. Otherwise, set to "cluster".
	// Enum: ["svm","cluster"]
	Scope *string `json:"scope,omitempty" yaml:"scope,omitempty"`

	// svm
	Svm *IPServicePolicyInlineSvm `json:"svm,omitempty" yaml:"svm,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this ip service policy
func (m *IPServicePolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPServicePolicyInlineServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIpspace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
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

func (m *IPServicePolicy) validateLinks(formats strfmt.Registry) error {
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

func (m *IPServicePolicy) validateIPServicePolicyInlineServices(formats strfmt.Registry) error {
	if swag.IsZero(m.IPServicePolicyInlineServices) { // not required
		return nil
	}

	for i := 0; i < len(m.IPServicePolicyInlineServices); i++ {
		if swag.IsZero(m.IPServicePolicyInlineServices[i]) { // not required
			continue
		}

		if m.IPServicePolicyInlineServices[i] != nil {
			if err := m.IPServicePolicyInlineServices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IPServicePolicy) validateIpspace(formats strfmt.Registry) error {
	if swag.IsZero(m.Ipspace) { // not required
		return nil
	}

	if m.Ipspace != nil {
		if err := m.Ipspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipspace")
			}
			return err
		}
	}

	return nil
}

var ipServicePolicyTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["svm","cluster"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipServicePolicyTypeScopePropEnum = append(ipServicePolicyTypeScopePropEnum, v)
	}
}

const (

	// IPServicePolicyScopeSvm captures enum value "svm"
	IPServicePolicyScopeSvm string = "svm"

	// IPServicePolicyScopeCluster captures enum value "cluster"
	IPServicePolicyScopeCluster string = "cluster"
)

// prop value enum
func (m *IPServicePolicy) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ipServicePolicyTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IPServicePolicy) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", *m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *IPServicePolicy) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this ip service policy based on the context it is used
func (m *IPServicePolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPServicePolicyInlineServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIpspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsBuiltIn(ctx, formats); err != nil {
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

func (m *IPServicePolicy) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IPServicePolicy) contextValidateIPServicePolicyInlineServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPServicePolicyInlineServices); i++ {

		if m.IPServicePolicyInlineServices[i] != nil {

			if swag.IsZero(m.IPServicePolicyInlineServices[i]) { // not required
				return nil
			}

			if err := m.IPServicePolicyInlineServices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IPServicePolicy) contextValidateIpspace(ctx context.Context, formats strfmt.Registry) error {

	if m.Ipspace != nil {

		if swag.IsZero(m.Ipspace) { // not required
			return nil
		}

		if err := m.Ipspace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipspace")
			}
			return err
		}
	}

	return nil
}

func (m *IPServicePolicy) contextValidateIsBuiltIn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_built_in", "body", m.IsBuiltIn); err != nil {
		return err
	}

	return nil
}

func (m *IPServicePolicy) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IPServicePolicy) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPServicePolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPServicePolicy) UnmarshalBinary(b []byte) error {
	var res IPServicePolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPServicePolicyInlineIpspace ip service policy inline ipspace
//
// swagger:model ip_service_policy_inline_ipspace
type IPServicePolicyInlineIpspace struct {

	// links
	Links *IPServicePolicyInlineIpspaceInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// IPspace name
	// Example: Default
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// IPspace UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this ip service policy inline ipspace
func (m *IPServicePolicyInlineIpspace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPServicePolicyInlineIpspace) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipspace" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipspace" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ip service policy inline ipspace based on the context it is used
func (m *IPServicePolicyInlineIpspace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPServicePolicyInlineIpspace) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipspace" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipspace" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPServicePolicyInlineIpspace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPServicePolicyInlineIpspace) UnmarshalBinary(b []byte) error {
	var res IPServicePolicyInlineIpspace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPServicePolicyInlineIpspaceInlineLinks ip service policy inline ipspace inline links
//
// swagger:model ip_service_policy_inline_ipspace_inline__links
type IPServicePolicyInlineIpspaceInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this ip service policy inline ipspace inline links
func (m *IPServicePolicyInlineIpspaceInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPServicePolicyInlineIpspaceInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipspace" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipspace" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ip service policy inline ipspace inline links based on the context it is used
func (m *IPServicePolicyInlineIpspaceInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPServicePolicyInlineIpspaceInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipspace" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipspace" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPServicePolicyInlineIpspaceInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPServicePolicyInlineIpspaceInlineLinks) UnmarshalBinary(b []byte) error {
	var res IPServicePolicyInlineIpspaceInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPServicePolicyInlineLinks ip service policy inline links
//
// swagger:model ip_service_policy_inline__links
type IPServicePolicyInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this ip service policy inline links
func (m *IPServicePolicyInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPServicePolicyInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ip service policy inline links based on the context it is used
func (m *IPServicePolicyInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPServicePolicyInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IPServicePolicyInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPServicePolicyInlineLinks) UnmarshalBinary(b []byte) error {
	var res IPServicePolicyInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPServicePolicyInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model ip_service_policy_inline_svm
type IPServicePolicyInlineSvm struct {

	// links
	Links *IPServicePolicyInlineSvmInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this ip service policy inline svm
func (m *IPServicePolicyInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPServicePolicyInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this ip service policy inline svm based on the context it is used
func (m *IPServicePolicyInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPServicePolicyInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IPServicePolicyInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPServicePolicyInlineSvm) UnmarshalBinary(b []byte) error {
	var res IPServicePolicyInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPServicePolicyInlineSvmInlineLinks ip service policy inline svm inline links
//
// swagger:model ip_service_policy_inline_svm_inline__links
type IPServicePolicyInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this ip service policy inline svm inline links
func (m *IPServicePolicyInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPServicePolicyInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ip service policy inline svm inline links based on the context it is used
func (m *IPServicePolicyInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPServicePolicyInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IPServicePolicyInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPServicePolicyInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res IPServicePolicyInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
