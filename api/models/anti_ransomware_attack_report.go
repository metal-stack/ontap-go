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

// AntiRansomwareAttackReport anti ransomware attack report
//
// swagger:model anti_ransomware_attack_report
type AntiRansomwareAttackReport struct {

	// links
	Links *AntiRansomwareAttackReportInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// Timestamp at which ransomware attack is observed.
	// Example: 2021-06-01 15:06:41
	// Format: date-time
	Time *strfmt.DateTime `json:"time,omitempty" yaml:"time,omitempty"`
}

// Validate validates this anti ransomware attack report
func (m *AntiRansomwareAttackReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareAttackReport) validateLinks(formats strfmt.Registry) error {
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

func (m *AntiRansomwareAttackReport) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this anti ransomware attack report based on the context it is used
func (m *AntiRansomwareAttackReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareAttackReport) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AntiRansomwareAttackReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntiRansomwareAttackReport) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareAttackReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AntiRansomwareAttackReportInlineLinks anti ransomware attack report inline links
//
// swagger:model anti_ransomware_attack_report_inline__links
type AntiRansomwareAttackReportInlineLinks struct {

	// suspects
	Suspects *Href `json:"suspects,omitempty" yaml:"suspects,omitempty"`
}

// Validate validates this anti ransomware attack report inline links
func (m *AntiRansomwareAttackReportInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSuspects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareAttackReportInlineLinks) validateSuspects(formats strfmt.Registry) error {
	if swag.IsZero(m.Suspects) { // not required
		return nil
	}

	if m.Suspects != nil {
		if err := m.Suspects.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "suspects")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "suspects")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this anti ransomware attack report inline links based on the context it is used
func (m *AntiRansomwareAttackReportInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSuspects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareAttackReportInlineLinks) contextValidateSuspects(ctx context.Context, formats strfmt.Registry) error {

	if m.Suspects != nil {

		if swag.IsZero(m.Suspects) { // not required
			return nil
		}

		if err := m.Suspects.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "suspects")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "suspects")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntiRansomwareAttackReportInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntiRansomwareAttackReportInlineLinks) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareAttackReportInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}