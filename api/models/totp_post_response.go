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
)

// TotpPostResponse totp post response
//
// swagger:model totp_post_response
type TotpPostResponse struct {

	// Number of records.
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty" yaml:"num_records,omitempty"`

	// totp post response inline records
	TotpPostResponseInlineRecords []*TotpPost `json:"records,omitempty" yaml:"records,omitempty"`
}

// Validate validates this totp post response
func (m *TotpPostResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTotpPostResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TotpPostResponse) validateTotpPostResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.TotpPostResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.TotpPostResponseInlineRecords); i++ {
		if swag.IsZero(m.TotpPostResponseInlineRecords[i]) { // not required
			continue
		}

		if m.TotpPostResponseInlineRecords[i] != nil {
			if err := m.TotpPostResponseInlineRecords[i].Validate(formats); err != nil {
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

// ContextValidate validate this totp post response based on the context it is used
func (m *TotpPostResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTotpPostResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TotpPostResponse) contextValidateTotpPostResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TotpPostResponseInlineRecords); i++ {

		if m.TotpPostResponseInlineRecords[i] != nil {

			if swag.IsZero(m.TotpPostResponseInlineRecords[i]) { // not required
				return nil
			}

			if err := m.TotpPostResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (m *TotpPostResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TotpPostResponse) UnmarshalBinary(b []byte) error {
	var res TotpPostResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}