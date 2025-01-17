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

// ZappS3BucketApplicationComponentsAccessPolicies The list of S3 objectstore policies to be created.
//
// swagger:model zapp_s3_bucket_application_components_access_policies
type ZappS3BucketApplicationComponentsAccessPolicies struct {

	// actions
	Actions []string `json:"actions" yaml:"actions"`

	// conditions.
	Conditions []*ZappS3BucketApplicationComponentsAccessPoliciesConditions `json:"conditions" yaml:"conditions"`

	// Allow or Deny Access.
	// Required: true
	// Enum: ["allow","deny"]
	Effect *string `json:"effect" yaml:"effect"`

	// principals
	Principals []string `json:"principals" yaml:"principals"`

	// resources
	Resources []string `json:"resources" yaml:"resources"`

	// Statement Identifier Usage: &lt;(size 1..256)&gt;
	// Max Length: 256
	// Min Length: 1
	Sid string `json:"sid,omitempty" yaml:"sid,omitempty"`
}

// Validate validates this zapp s3 bucket application components access policies
func (m *ZappS3BucketApplicationComponentsAccessPolicies) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZappS3BucketApplicationComponentsAccessPolicies) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var zappS3BucketApplicationComponentsAccessPoliciesTypeEffectPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow","deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		zappS3BucketApplicationComponentsAccessPoliciesTypeEffectPropEnum = append(zappS3BucketApplicationComponentsAccessPoliciesTypeEffectPropEnum, v)
	}
}

const (

	// ZappS3BucketApplicationComponentsAccessPoliciesEffectAllow captures enum value "allow"
	ZappS3BucketApplicationComponentsAccessPoliciesEffectAllow string = "allow"

	// ZappS3BucketApplicationComponentsAccessPoliciesEffectDeny captures enum value "deny"
	ZappS3BucketApplicationComponentsAccessPoliciesEffectDeny string = "deny"
)

// prop value enum
func (m *ZappS3BucketApplicationComponentsAccessPolicies) validateEffectEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, zappS3BucketApplicationComponentsAccessPoliciesTypeEffectPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ZappS3BucketApplicationComponentsAccessPolicies) validateEffect(formats strfmt.Registry) error {

	if err := validate.Required("effect", "body", m.Effect); err != nil {
		return err
	}

	// value enum
	if err := m.validateEffectEnum("effect", "body", *m.Effect); err != nil {
		return err
	}

	return nil
}

func (m *ZappS3BucketApplicationComponentsAccessPolicies) validateSid(formats strfmt.Registry) error {
	if swag.IsZero(m.Sid) { // not required
		return nil
	}

	if err := validate.MinLength("sid", "body", m.Sid, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("sid", "body", m.Sid, 256); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this zapp s3 bucket application components access policies based on the context it is used
func (m *ZappS3BucketApplicationComponentsAccessPolicies) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZappS3BucketApplicationComponentsAccessPolicies) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Conditions); i++ {

		if m.Conditions[i] != nil {

			if swag.IsZero(m.Conditions[i]) { // not required
				return nil
			}

			if err := m.Conditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ZappS3BucketApplicationComponentsAccessPolicies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZappS3BucketApplicationComponentsAccessPolicies) UnmarshalBinary(b []byte) error {
	var res ZappS3BucketApplicationComponentsAccessPolicies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
