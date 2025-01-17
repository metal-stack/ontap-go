// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TopMetricsClientResponse top metrics client response
//
// swagger:model top_metrics_client_response
type TopMetricsClientResponse struct {

	// links
	Links *TopMetricsClientResponseInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// incomplete response reason
	IncompleteResponseReason *TopMetricsClientResponseInlineIncompleteResponseReason `json:"incomplete_response_reason,omitempty" yaml:"incomplete_response_reason,omitempty"`

	// notice
	Notice *TopMetricsClientResponseInlineNotice `json:"notice,omitempty" yaml:"notice,omitempty"`

	// Number of records.
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty" yaml:"num_records,omitempty"`

	// top metrics client response inline records
	TopMetricsClientResponseInlineRecords []*TopMetricsClient `json:"records,omitempty" yaml:"records,omitempty"`
}

// Validate validates this top metrics client response
func (m *TopMetricsClientResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncompleteResponseReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopMetricsClientResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsClientResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *TopMetricsClientResponse) validateIncompleteResponseReason(formats strfmt.Registry) error {
	if swag.IsZero(m.IncompleteResponseReason) { // not required
		return nil
	}

	if m.IncompleteResponseReason != nil {
		if err := m.IncompleteResponseReason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incomplete_response_reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incomplete_response_reason")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsClientResponse) validateNotice(formats strfmt.Registry) error {
	if swag.IsZero(m.Notice) { // not required
		return nil
	}

	if m.Notice != nil {
		if err := m.Notice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notice")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsClientResponse) validateTopMetricsClientResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.TopMetricsClientResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.TopMetricsClientResponseInlineRecords); i++ {
		if swag.IsZero(m.TopMetricsClientResponseInlineRecords[i]) { // not required
			continue
		}

		if m.TopMetricsClientResponseInlineRecords[i] != nil {
			if err := m.TopMetricsClientResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this top metrics client response based on the context it is used
func (m *TopMetricsClientResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncompleteResponseReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTopMetricsClientResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsClientResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TopMetricsClientResponse) contextValidateIncompleteResponseReason(ctx context.Context, formats strfmt.Registry) error {

	if m.IncompleteResponseReason != nil {

		if swag.IsZero(m.IncompleteResponseReason) { // not required
			return nil
		}

		if err := m.IncompleteResponseReason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incomplete_response_reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incomplete_response_reason")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsClientResponse) contextValidateNotice(ctx context.Context, formats strfmt.Registry) error {

	if m.Notice != nil {

		if swag.IsZero(m.Notice) { // not required
			return nil
		}

		if err := m.Notice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notice")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsClientResponse) contextValidateTopMetricsClientResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TopMetricsClientResponseInlineRecords); i++ {

		if m.TopMetricsClientResponseInlineRecords[i] != nil {

			if swag.IsZero(m.TopMetricsClientResponseInlineRecords[i]) { // not required
				return nil
			}

			if err := m.TopMetricsClientResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsClientResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsClientResponse) UnmarshalBinary(b []byte) error {
	var res TopMetricsClientResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsClientResponseInlineIncompleteResponseReason Indicates that the metric report provides incomplete data.
//
// swagger:model top_metrics_client_response_inline_incomplete_response_reason
type TopMetricsClientResponseInlineIncompleteResponseReason struct {

	// Warning code indicating why partial data was reported.
	// Example: 111411207
	// Read Only: true
	Code *string `json:"code,omitempty" yaml:"code,omitempty"`

	// A message describing the reason for partial data.
	// Example: Partial data has been returned for this metric report. Reason: The activity tracking report for this volume is not available because the system is busy collecting tracking data.
	// Read Only: true
	Message *string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Validate validates this top metrics client response inline incomplete response reason
func (m *TopMetricsClientResponseInlineIncompleteResponseReason) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this top metrics client response inline incomplete response reason based on the context it is used
func (m *TopMetricsClientResponseInlineIncompleteResponseReason) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *TopMetricsClientResponseInlineIncompleteResponseReason) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "incomplete_response_reason"+"."+"code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsClientResponseInlineIncompleteResponseReason) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "incomplete_response_reason"+"."+"message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsClientResponseInlineIncompleteResponseReason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsClientResponseInlineIncompleteResponseReason) UnmarshalBinary(b []byte) error {
	var res TopMetricsClientResponseInlineIncompleteResponseReason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsClientResponseInlineLinks top metrics client response inline links
//
// swagger:model top_metrics_client_response_inline__links
type TopMetricsClientResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty" yaml:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this top metrics client response inline links
func (m *TopMetricsClientResponseInlineLinks) Validate(formats strfmt.Registry) error {
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

func (m *TopMetricsClientResponseInlineLinks) validateNext(formats strfmt.Registry) error {
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

func (m *TopMetricsClientResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics client response inline links based on the context it is used
func (m *TopMetricsClientResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *TopMetricsClientResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TopMetricsClientResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsClientResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsClientResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res TopMetricsClientResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsClientResponseInlineNotice Optional field that indicates why no records are returned by the volume activity tracking REST API.
//
// swagger:model top_metrics_client_response_inline_notice
type TopMetricsClientResponseInlineNotice struct {

	// Warning code indicating why no records are returned.
	// Example: 111411207
	// Read Only: true
	Code *string `json:"code,omitempty" yaml:"code,omitempty"`

	// Details why no records are returned.
	// Example: No read/write traffic on volume
	// Read Only: true
	Message *string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Validate validates this top metrics client response inline notice
func (m *TopMetricsClientResponseInlineNotice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this top metrics client response inline notice based on the context it is used
func (m *TopMetricsClientResponseInlineNotice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *TopMetricsClientResponseInlineNotice) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "notice"+"."+"code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsClientResponseInlineNotice) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "notice"+"."+"message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsClientResponseInlineNotice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsClientResponseInlineNotice) UnmarshalBinary(b []byte) error {
	var res TopMetricsClientResponseInlineNotice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}