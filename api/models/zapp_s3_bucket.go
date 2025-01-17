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

// ZappS3Bucket A generic S3 bucket application.
//
// swagger:model zapp_s3_bucket
type ZappS3Bucket struct {

	// The list of application components to be created.
	// Required: true
	// Max Items: 10
	// Min Items: 1
	ApplicationComponents []*ZappS3BucketApplicationComponents `json:"application_components" yaml:"application_components"`

	// protection type
	ProtectionType *ZappS3BucketProtectionType `json:"protection_type,omitempty" yaml:"protection_type,omitempty"`
}

// Validate validates this zapp s3 bucket
func (m *ZappS3Bucket) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZappS3Bucket) validateApplicationComponents(formats strfmt.Registry) error {

	if err := validate.Required("application_components", "body", m.ApplicationComponents); err != nil {
		return err
	}

	iApplicationComponentsSize := int64(len(m.ApplicationComponents))

	if err := validate.MinItems("application_components", "body", iApplicationComponentsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("application_components", "body", iApplicationComponentsSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(m.ApplicationComponents); i++ {
		if swag.IsZero(m.ApplicationComponents[i]) { // not required
			continue
		}

		if m.ApplicationComponents[i] != nil {
			if err := m.ApplicationComponents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("application_components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("application_components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ZappS3Bucket) validateProtectionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionType) { // not required
		return nil
	}

	if m.ProtectionType != nil {
		if err := m.ProtectionType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protection_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protection_type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this zapp s3 bucket based on the context it is used
func (m *ZappS3Bucket) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplicationComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtectionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZappS3Bucket) contextValidateApplicationComponents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ApplicationComponents); i++ {

		if m.ApplicationComponents[i] != nil {

			if swag.IsZero(m.ApplicationComponents[i]) { // not required
				return nil
			}

			if err := m.ApplicationComponents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("application_components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("application_components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ZappS3Bucket) contextValidateProtectionType(ctx context.Context, formats strfmt.Registry) error {

	if m.ProtectionType != nil {

		if swag.IsZero(m.ProtectionType) { // not required
			return nil
		}

		if err := m.ProtectionType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protection_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protection_type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ZappS3Bucket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZappS3Bucket) UnmarshalBinary(b []byte) error {
	var res ZappS3Bucket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ZappS3BucketProtectionType zapp s3 bucket protection type
//
// swagger:model ZappS3BucketProtectionType
type ZappS3BucketProtectionType struct {

	// The remote RPO of the application.
	// Enum: ["none","zero"]
	RemoteRpo string `json:"remote_rpo,omitempty" yaml:"remote_rpo,omitempty"`
}

// Validate validates this zapp s3 bucket protection type
func (m *ZappS3BucketProtectionType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRemoteRpo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var zappS3BucketProtectionTypeTypeRemoteRpoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","zero"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		zappS3BucketProtectionTypeTypeRemoteRpoPropEnum = append(zappS3BucketProtectionTypeTypeRemoteRpoPropEnum, v)
	}
}

const (

	// ZappS3BucketProtectionTypeRemoteRpoNone captures enum value "none"
	ZappS3BucketProtectionTypeRemoteRpoNone string = "none"

	// ZappS3BucketProtectionTypeRemoteRpoZero captures enum value "zero"
	ZappS3BucketProtectionTypeRemoteRpoZero string = "zero"
)

// prop value enum
func (m *ZappS3BucketProtectionType) validateRemoteRpoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, zappS3BucketProtectionTypeTypeRemoteRpoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ZappS3BucketProtectionType) validateRemoteRpo(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteRpo) { // not required
		return nil
	}

	// value enum
	if err := m.validateRemoteRpoEnum("protection_type"+"."+"remote_rpo", "body", m.RemoteRpo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this zapp s3 bucket protection type based on context it is used
func (m *ZappS3BucketProtectionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ZappS3BucketProtectionType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZappS3BucketProtectionType) UnmarshalBinary(b []byte) error {
	var res ZappS3BucketProtectionType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
