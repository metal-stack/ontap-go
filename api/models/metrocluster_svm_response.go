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

// MetroclusterSvmResponse metrocluster svm response
//
// swagger:model metrocluster_svm_response
type MetroclusterSvmResponse struct {

	// links
	Links *CollectionLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// metrocluster svm response inline records
	MetroclusterSvmResponseInlineRecords []*MetroclusterSvm `json:"records,omitempty" yaml:"records,omitempty"`

	// Number of Records
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty" yaml:"num_records,omitempty"`
}

// Validate validates this metrocluster svm response
func (m *MetroclusterSvmResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetroclusterSvmResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterSvmResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *MetroclusterSvmResponse) validateMetroclusterSvmResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.MetroclusterSvmResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.MetroclusterSvmResponseInlineRecords); i++ {
		if swag.IsZero(m.MetroclusterSvmResponseInlineRecords[i]) { // not required
			continue
		}

		if m.MetroclusterSvmResponseInlineRecords[i] != nil {
			if err := m.MetroclusterSvmResponseInlineRecords[i].Validate(formats); err != nil {
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

// ContextValidate validate this metrocluster svm response based on the context it is used
func (m *MetroclusterSvmResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetroclusterSvmResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterSvmResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MetroclusterSvmResponse) contextValidateMetroclusterSvmResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MetroclusterSvmResponseInlineRecords); i++ {

		if m.MetroclusterSvmResponseInlineRecords[i] != nil {

			if swag.IsZero(m.MetroclusterSvmResponseInlineRecords[i]) { // not required
				return nil
			}

			if err := m.MetroclusterSvmResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (m *MetroclusterSvmResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterSvmResponse) UnmarshalBinary(b []byte) error {
	var res MetroclusterSvmResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
