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

// S3BucketPolicyStatement Specifies information about a single access permission.
//
// swagger:model s3_bucket_policy_statement
type S3BucketPolicyStatement struct {

	// Specifies whether access is allowed or denied when a user requests the specific action. If access (to allow) is not granted explicitly to a resource, access is implicitly denied. Access can also be denied explicitly to a resource, in order to make sure that a user cannot access it, even if a different policy grants access.
	// Example: allow
	// Enum: ["allow","deny"]
	Effect *string `json:"effect,omitempty" yaml:"effect,omitempty"`

	// s3 bucket policy statement inline actions
	// Example: ["GetObject","PutObject","DeleteObject","ListBucket"]
	S3BucketPolicyStatementInlineActions []*string `json:"actions,omitempty" yaml:"actions,omitempty"`

	// Specifies bucket policy conditions.
	S3BucketPolicyStatementInlineConditions []*S3BucketPolicyCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`

	// s3 bucket policy statement inline principals
	// Example: ["user1","group/grp1","nasgroup/group1"]
	S3BucketPolicyStatementInlinePrincipals []*string `json:"principals,omitempty" yaml:"principals,omitempty"`

	// s3 bucket policy statement inline resources
	// Example: ["bucket1","bucket1/*"]
	S3BucketPolicyStatementInlineResources []*string `json:"resources,omitempty" yaml:"resources,omitempty"`

	// Specifies the statement identifier used to differentiate between statements. The sid length can range from 1 to 256 characters and can only contain the following combination of characters 0-9, A-Z, and a-z. Special characters are not valid.
	// Example: FullAccessToUser1
	// Max Length: 256
	// Min Length: 0
	// Pattern: ^[0-9A-Za-z]{0,256}$
	Sid *string `json:"sid,omitempty" yaml:"sid,omitempty"`
}

// Validate validates this s3 bucket policy statement
func (m *S3BucketPolicyStatement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3BucketPolicyStatementInlineConditions(formats); err != nil {
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

var s3BucketPolicyStatementTypeEffectPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow","deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		s3BucketPolicyStatementTypeEffectPropEnum = append(s3BucketPolicyStatementTypeEffectPropEnum, v)
	}
}

const (

	// S3BucketPolicyStatementEffectAllow captures enum value "allow"
	S3BucketPolicyStatementEffectAllow string = "allow"

	// S3BucketPolicyStatementEffectDeny captures enum value "deny"
	S3BucketPolicyStatementEffectDeny string = "deny"
)

// prop value enum
func (m *S3BucketPolicyStatement) validateEffectEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, s3BucketPolicyStatementTypeEffectPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *S3BucketPolicyStatement) validateEffect(formats strfmt.Registry) error {
	if swag.IsZero(m.Effect) { // not required
		return nil
	}

	// value enum
	if err := m.validateEffectEnum("effect", "body", *m.Effect); err != nil {
		return err
	}

	return nil
}

func (m *S3BucketPolicyStatement) validateS3BucketPolicyStatementInlineConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.S3BucketPolicyStatementInlineConditions) { // not required
		return nil
	}

	for i := 0; i < len(m.S3BucketPolicyStatementInlineConditions); i++ {
		if swag.IsZero(m.S3BucketPolicyStatementInlineConditions[i]) { // not required
			continue
		}

		if m.S3BucketPolicyStatementInlineConditions[i] != nil {
			if err := m.S3BucketPolicyStatementInlineConditions[i].Validate(formats); err != nil {
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

func (m *S3BucketPolicyStatement) validateSid(formats strfmt.Registry) error {
	if swag.IsZero(m.Sid) { // not required
		return nil
	}

	if err := validate.MinLength("sid", "body", *m.Sid, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("sid", "body", *m.Sid, 256); err != nil {
		return err
	}

	if err := validate.Pattern("sid", "body", *m.Sid, `^[0-9A-Za-z]{0,256}$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this s3 bucket policy statement based on the context it is used
func (m *S3BucketPolicyStatement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateS3BucketPolicyStatementInlineConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketPolicyStatement) contextValidateS3BucketPolicyStatementInlineConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.S3BucketPolicyStatementInlineConditions); i++ {

		if m.S3BucketPolicyStatementInlineConditions[i] != nil {

			if swag.IsZero(m.S3BucketPolicyStatementInlineConditions[i]) { // not required
				return nil
			}

			if err := m.S3BucketPolicyStatementInlineConditions[i].ContextValidate(ctx, formats); err != nil {
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
func (m *S3BucketPolicyStatement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketPolicyStatement) UnmarshalBinary(b []byte) error {
	var res S3BucketPolicyStatement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}