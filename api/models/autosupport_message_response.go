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

// AutosupportMessageResponse autosupport message response
//
// swagger:model autosupport_message_response
type AutosupportMessageResponse struct {

	// links
	Links *AutosupportMessageResponseInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// List of messages invoked on the node
	// Read Only: true
	AutosupportMessageResponseInlineRecords []*AutosupportMessage `json:"records,omitempty" yaml:"records,omitempty"`

	// Number of records
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty" yaml:"num_records,omitempty"`
}

// Validate validates this autosupport message response
func (m *AutosupportMessageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutosupportMessageResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutosupportMessageResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *AutosupportMessageResponse) validateAutosupportMessageResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.AutosupportMessageResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.AutosupportMessageResponseInlineRecords); i++ {
		if swag.IsZero(m.AutosupportMessageResponseInlineRecords[i]) { // not required
			continue
		}

		if m.AutosupportMessageResponseInlineRecords[i] != nil {
			if err := m.AutosupportMessageResponseInlineRecords[i].Validate(formats); err != nil {
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

// ContextValidate validate this autosupport message response based on the context it is used
func (m *AutosupportMessageResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAutosupportMessageResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutosupportMessageResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AutosupportMessageResponse) contextValidateAutosupportMessageResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "records", "body", []*AutosupportMessage(m.AutosupportMessageResponseInlineRecords)); err != nil {
		return err
	}

	for i := 0; i < len(m.AutosupportMessageResponseInlineRecords); i++ {

		if m.AutosupportMessageResponseInlineRecords[i] != nil {

			if swag.IsZero(m.AutosupportMessageResponseInlineRecords[i]) { // not required
				return nil
			}

			if err := m.AutosupportMessageResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (m *AutosupportMessageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutosupportMessageResponse) UnmarshalBinary(b []byte) error {
	var res AutosupportMessageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AutosupportMessageResponseInlineLinks autosupport message response inline links
//
// swagger:model autosupport_message_response_inline__links
type AutosupportMessageResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty" yaml:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this autosupport message response inline links
func (m *AutosupportMessageResponseInlineLinks) Validate(formats strfmt.Registry) error {
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

func (m *AutosupportMessageResponseInlineLinks) validateNext(formats strfmt.Registry) error {
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

func (m *AutosupportMessageResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this autosupport message response inline links based on the context it is used
func (m *AutosupportMessageResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *AutosupportMessageResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AutosupportMessageResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AutosupportMessageResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutosupportMessageResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res AutosupportMessageResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
