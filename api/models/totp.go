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

// Totp TOTP profile for the user account used to access SSH.
//
// swagger:model totp
type Totp struct {

	// links
	Links *TotpInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// account
	Account *AccountReference `json:"account,omitempty" yaml:"account,omitempty"`

	// Optional comment for the TOTP profile.
	Comment *string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// Status of the TOTP profile.
	// Example: false
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// owner
	Owner *TotpInlineOwner `json:"owner,omitempty" yaml:"owner,omitempty"`

	// Scope of the entity. Set to "cluster" for cluster owned objects and to "svm" for SVM owned objects.
	// Read Only: true
	// Enum: ["cluster","svm"]
	Scope *string `json:"scope,omitempty" yaml:"scope,omitempty"`

	// SHA fingerprint for the TOTP secret key.
	// Read Only: true
	ShaFingerprint *string `json:"sha_fingerprint,omitempty" yaml:"sha_fingerprint,omitempty"`
}

// Validate validates this totp
func (m *Totp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Totp) validateLinks(formats strfmt.Registry) error {
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

func (m *Totp) validateAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (m *Totp) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

var totpTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cluster","svm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		totpTypeScopePropEnum = append(totpTypeScopePropEnum, v)
	}
}

const (

	// TotpScopeCluster captures enum value "cluster"
	TotpScopeCluster string = "cluster"

	// TotpScopeSvm captures enum value "svm"
	TotpScopeSvm string = "svm"
)

// prop value enum
func (m *Totp) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, totpTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Totp) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", *m.Scope); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this totp based on the context it is used
func (m *Totp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShaFingerprint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Totp) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Totp) contextValidateAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.Account != nil {

		if swag.IsZero(m.Account) { // not required
			return nil
		}

		if err := m.Account.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (m *Totp) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {

		if swag.IsZero(m.Owner) { // not required
			return nil
		}

		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *Totp) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *Totp) contextValidateShaFingerprint(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sha_fingerprint", "body", m.ShaFingerprint); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Totp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Totp) UnmarshalBinary(b []byte) error {
	var res Totp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TotpInlineLinks totp inline links
//
// swagger:model totp_inline__links
type TotpInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this totp inline links
func (m *TotpInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TotpInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this totp inline links based on the context it is used
func (m *TotpInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TotpInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TotpInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TotpInlineLinks) UnmarshalBinary(b []byte) error {
	var res TotpInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TotpInlineOwner Owner name and UUID that uniquely identifies the TOTP profile.
//
// swagger:model totp_inline_owner
type TotpInlineOwner struct {

	// links
	Links *TotpInlineOwnerInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this totp inline owner
func (m *TotpInlineOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TotpInlineOwner) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this totp inline owner based on the context it is used
func (m *TotpInlineOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TotpInlineOwner) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TotpInlineOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TotpInlineOwner) UnmarshalBinary(b []byte) error {
	var res TotpInlineOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TotpInlineOwnerInlineLinks totp inline owner inline links
//
// swagger:model totp_inline_owner_inline__links
type TotpInlineOwnerInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this totp inline owner inline links
func (m *TotpInlineOwnerInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TotpInlineOwnerInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this totp inline owner inline links based on the context it is used
func (m *TotpInlineOwnerInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TotpInlineOwnerInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TotpInlineOwnerInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TotpInlineOwnerInlineLinks) UnmarshalBinary(b []byte) error {
	var res TotpInlineOwnerInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}