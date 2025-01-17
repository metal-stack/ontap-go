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

// TopMetricsSvmClientExcludedVolume List of volumes and their details as to why they are not included in the SVM activity tracking REST API.
//
// swagger:model top_metrics_svm_client_excluded_volume
type TopMetricsSvmClientExcludedVolume struct {

	// reason
	Reason *TopMetricsSvmClientExcludedVolumeInlineReason `json:"reason,omitempty" yaml:"reason,omitempty"`

	// volume
	Volume *TopMetricsSvmClientExcludedVolumeInlineVolume `json:"volume,omitempty" yaml:"volume,omitempty"`
}

// Validate validates this top metrics svm client excluded volume
func (m *TopMetricsSvmClientExcludedVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReason(formats); err != nil {
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

func (m *TopMetricsSvmClientExcludedVolume) validateReason(formats strfmt.Registry) error {
	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if m.Reason != nil {
		if err := m.Reason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reason")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmClientExcludedVolume) validateVolume(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics svm client excluded volume based on the context it is used
func (m *TopMetricsSvmClientExcludedVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReason(ctx, formats); err != nil {
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

func (m *TopMetricsSvmClientExcludedVolume) contextValidateReason(ctx context.Context, formats strfmt.Registry) error {

	if m.Reason != nil {

		if swag.IsZero(m.Reason) { // not required
			return nil
		}

		if err := m.Reason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reason")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmClientExcludedVolume) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsSvmClientExcludedVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmClientExcludedVolume) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmClientExcludedVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmClientExcludedVolumeInlineReason top metrics svm client excluded volume inline reason
//
// swagger:model top_metrics_svm_client_excluded_volume_inline_reason
type TopMetricsSvmClientExcludedVolumeInlineReason struct {

	// Warning code indicating why the volume is not included in the SVM activity tracking REST API.
	// Example: 111411207
	// Read Only: true
	Code *string `json:"code,omitempty" yaml:"code,omitempty"`

	// Details why the volume is not included in the SVM activity tracking REST API.
	// Example: The volume is offline.
	// Read Only: true
	Message *string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Validate validates this top metrics svm client excluded volume inline reason
func (m *TopMetricsSvmClientExcludedVolumeInlineReason) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this top metrics svm client excluded volume inline reason based on the context it is used
func (m *TopMetricsSvmClientExcludedVolumeInlineReason) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *TopMetricsSvmClientExcludedVolumeInlineReason) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reason"+"."+"code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsSvmClientExcludedVolumeInlineReason) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reason"+"."+"message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsSvmClientExcludedVolumeInlineReason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmClientExcludedVolumeInlineReason) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmClientExcludedVolumeInlineReason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmClientExcludedVolumeInlineVolume top metrics svm client excluded volume inline volume
//
// swagger:model top_metrics_svm_client_excluded_volume_inline_volume
type TopMetricsSvmClientExcludedVolumeInlineVolume struct {

	// links
	Links *TopMetricsSvmClientExcludedVolumeInlineVolumeInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the volume. This field cannot be specified in a POST or PATCH method.
	// Example: volume1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	// Example: 028baa66-41bd-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this top metrics svm client excluded volume inline volume
func (m *TopMetricsSvmClientExcludedVolumeInlineVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmClientExcludedVolumeInlineVolume) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics svm client excluded volume inline volume based on the context it is used
func (m *TopMetricsSvmClientExcludedVolumeInlineVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmClientExcludedVolumeInlineVolume) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsSvmClientExcludedVolumeInlineVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmClientExcludedVolumeInlineVolume) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmClientExcludedVolumeInlineVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmClientExcludedVolumeInlineVolumeInlineLinks top metrics svm client excluded volume inline volume inline links
//
// swagger:model top_metrics_svm_client_excluded_volume_inline_volume_inline__links
type TopMetricsSvmClientExcludedVolumeInlineVolumeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this top metrics svm client excluded volume inline volume inline links
func (m *TopMetricsSvmClientExcludedVolumeInlineVolumeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmClientExcludedVolumeInlineVolumeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics svm client excluded volume inline volume inline links based on the context it is used
func (m *TopMetricsSvmClientExcludedVolumeInlineVolumeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmClientExcludedVolumeInlineVolumeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsSvmClientExcludedVolumeInlineVolumeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmClientExcludedVolumeInlineVolumeInlineLinks) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmClientExcludedVolumeInlineVolumeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
