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

// Duo Duo profile for the SVM or cluster-management server (Cserver).
//
// swagger:model duo
type Duo struct {

	// links
	Links *DuoInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The URL at which the Duo API is hosted.
	// Example: api-****.duo.com
	APIHost *string `json:"api_host,omitempty" yaml:"api_host,omitempty"`

	// Automatically sends a push notification for authentication when using Duo.
	// Example: true
	AutoPush *bool `json:"auto_push,omitempty" yaml:"auto_push,omitempty"`

	// Comment for the Duo profile.
	Comment *string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// Determines the behavior of the system when it cannot communicate with the Duo service.
	// Example: safe
	// Enum: ["safe","secure"]
	FailMode *string `json:"fail_mode,omitempty" yaml:"fail_mode,omitempty"`

	// The SHA fingerprint corresponding to the Duo secret key.
	// Read Only: true
	Fingerprint *string `json:"fingerprint,omitempty" yaml:"fingerprint,omitempty"`

	// Specifies the HTTP proxy server to be used when connecting to the Duo service.
	// Example: IPaddress:port
	HTTPProxy *string `json:"http_proxy,omitempty" yaml:"http_proxy,omitempty"`

	// The Integration Key associated with the Duo profile.
	IntegrationKey *string `json:"integration_key,omitempty" yaml:"integration_key,omitempty"`

	// Indicates whether the Duo authentication feature is active or inactive.
	// Example: true
	IsEnabled *bool `json:"is_enabled,omitempty" yaml:"is_enabled,omitempty"`

	// The maximum number of authentication attempts allowed for a user before the process is terminated.
	// Example: 1
	// Maximum: 3
	// Minimum: 1
	MaxPrompts *int64 `json:"max_prompts,omitempty" yaml:"max_prompts,omitempty"`

	// owner
	Owner *DuoInlineOwner `json:"owner,omitempty" yaml:"owner,omitempty"`

	// Additional information sent along with the push notification for Duo authentication.
	// Example: true
	PushInfo *bool `json:"push_info,omitempty" yaml:"push_info,omitempty"`

	// The Secret Key associated with the Duo profile.
	SecretKey *string `json:"secret_key,omitempty" yaml:"secret_key,omitempty"`

	// Information on the reachability status of Duo.
	// Example: OK
	// Read Only: true
	Status *string `json:"status,omitempty" yaml:"status,omitempty"`
}

// Validate validates this duo
func (m *Duo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxPrompts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Duo) validateLinks(formats strfmt.Registry) error {
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

var duoTypeFailModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["safe","secure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		duoTypeFailModePropEnum = append(duoTypeFailModePropEnum, v)
	}
}

const (

	// DuoFailModeSafe captures enum value "safe"
	DuoFailModeSafe string = "safe"

	// DuoFailModeSecure captures enum value "secure"
	DuoFailModeSecure string = "secure"
)

// prop value enum
func (m *Duo) validateFailModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, duoTypeFailModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Duo) validateFailMode(formats strfmt.Registry) error {
	if swag.IsZero(m.FailMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateFailModeEnum("fail_mode", "body", *m.FailMode); err != nil {
		return err
	}

	return nil
}

func (m *Duo) validateMaxPrompts(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxPrompts) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_prompts", "body", *m.MaxPrompts, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max_prompts", "body", *m.MaxPrompts, 3, false); err != nil {
		return err
	}

	return nil
}

func (m *Duo) validateOwner(formats strfmt.Registry) error {
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

// ContextValidate validate this duo based on the context it is used
func (m *Duo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFingerprint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Duo) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Duo) contextValidateFingerprint(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "fingerprint", "body", m.Fingerprint); err != nil {
		return err
	}

	return nil
}

func (m *Duo) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Duo) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Duo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Duo) UnmarshalBinary(b []byte) error {
	var res Duo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DuoInlineLinks duo inline links
//
// swagger:model duo_inline__links
type DuoInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty" yaml:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this duo inline links
func (m *DuoInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DuoInlineLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *DuoInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this duo inline links based on the context it is used
func (m *DuoInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DuoInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {

		if swag.IsZero(m.Next) { // not required
			return nil
		}

		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *DuoInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DuoInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DuoInlineLinks) UnmarshalBinary(b []byte) error {
	var res DuoInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DuoInlineOwner SVM name and UUID for which the Duo profile is configured.
//
// swagger:model duo_inline_owner
type DuoInlineOwner struct {

	// links
	Links *DuoInlineOwnerInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this duo inline owner
func (m *DuoInlineOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DuoInlineOwner) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this duo inline owner based on the context it is used
func (m *DuoInlineOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DuoInlineOwner) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DuoInlineOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DuoInlineOwner) UnmarshalBinary(b []byte) error {
	var res DuoInlineOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DuoInlineOwnerInlineLinks duo inline owner inline links
//
// swagger:model duo_inline_owner_inline__links
type DuoInlineOwnerInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this duo inline owner inline links
func (m *DuoInlineOwnerInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DuoInlineOwnerInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this duo inline owner inline links based on the context it is used
func (m *DuoInlineOwnerInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DuoInlineOwnerInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DuoInlineOwnerInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DuoInlineOwnerInlineLinks) UnmarshalBinary(b []byte) error {
	var res DuoInlineOwnerInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}