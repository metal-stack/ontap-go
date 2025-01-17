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

// S3BucketLifecycleExpiration Information about the expiration lifecycle management action.
//
// swagger:model s3_bucket_lifecycle_expiration
type S3BucketLifecycleExpiration struct {

	// links
	Links *S3BucketLifecycleExpirationInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// Cleanup object delete markers.
	ExpiredObjectDeleteMarker *bool `json:"expired_object_delete_marker,omitempty" yaml:"expired_object_delete_marker,omitempty"`

	// Number of days since creation after which objects can be deleted. This cannot be used along with object_expiry_date.
	// Example: 100
	ObjectAgeDays *int64 `json:"object_age_days,omitempty" yaml:"object_age_days,omitempty"`

	// Specific date from when objects can expire. This cannot be used with object_age_days.
	// Example: 2039-09-23 00:00:00
	// Format: date-time
	ObjectExpiryDate *strfmt.DateTime `json:"object_expiry_date,omitempty" yaml:"object_expiry_date,omitempty"`
}

// Validate validates this s3 bucket lifecycle expiration
func (m *S3BucketLifecycleExpiration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectExpiryDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleExpiration) validateLinks(formats strfmt.Registry) error {
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

func (m *S3BucketLifecycleExpiration) validateObjectExpiryDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectExpiryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("object_expiry_date", "body", "date-time", m.ObjectExpiryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this s3 bucket lifecycle expiration based on the context it is used
func (m *S3BucketLifecycleExpiration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleExpiration) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *S3BucketLifecycleExpiration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleExpiration) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleExpiration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketLifecycleExpirationInlineLinks s3 bucket lifecycle expiration inline links
//
// swagger:model s3_bucket_lifecycle_expiration_inline__links
type S3BucketLifecycleExpirationInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this s3 bucket lifecycle expiration inline links
func (m *S3BucketLifecycleExpirationInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleExpirationInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this s3 bucket lifecycle expiration inline links based on the context it is used
func (m *S3BucketLifecycleExpirationInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleExpirationInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *S3BucketLifecycleExpirationInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleExpirationInlineLinks) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleExpirationInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}