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

// CifsDomainPreferredDc cifs domain preferred dc
//
// swagger:model cifs_domain_preferred_dc
type CifsDomainPreferredDc struct {

	// Fully Qualified Domain Name.
	//
	// Example: test.com
	Fqdn *string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`

	// IP address of the preferred domain controller (DC). The address can be either an IPv4 or an IPv6 address.
	//
	// Example: 4.4.4.4
	ServerIP *string `json:"server_ip,omitempty" yaml:"server_ip,omitempty"`

	// status
	Status *CifsDomainPreferredDcInlineStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

// Validate validates this cifs domain preferred dc
func (m *CifsDomainPreferredDc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainPreferredDc) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cifs domain preferred dc based on the context it is used
func (m *CifsDomainPreferredDc) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainPreferredDc) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsDomainPreferredDc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsDomainPreferredDc) UnmarshalBinary(b []byte) error {
	var res CifsDomainPreferredDc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsDomainPreferredDcInlineStatus Status of CIFS preferred domain controller.
//
// swagger:model cifs_domain_preferred_dc_inline_status
type CifsDomainPreferredDcInlineStatus struct {

	// Provides a detailed description of the state if the state is 'down' or
	// the response time of the DNS server if the state is 'up'.
	//
	// Example: Response time (msec): 111
	// Read Only: true
	Details *string `json:"details,omitempty" yaml:"details,omitempty"`

	// Indicates whether or not the domain controller is reachable.
	// Example: true
	// Read Only: true
	Reachable *bool `json:"reachable,omitempty" yaml:"reachable,omitempty"`
}

// Validate validates this cifs domain preferred dc inline status
func (m *CifsDomainPreferredDcInlineStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this cifs domain preferred dc inline status based on the context it is used
func (m *CifsDomainPreferredDcInlineStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReachable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsDomainPreferredDcInlineStatus) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status"+"."+"details", "body", m.Details); err != nil {
		return err
	}

	return nil
}

func (m *CifsDomainPreferredDcInlineStatus) contextValidateReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status"+"."+"reachable", "body", m.Reachable); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsDomainPreferredDcInlineStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsDomainPreferredDcInlineStatus) UnmarshalBinary(b []byte) error {
	var res CifsDomainPreferredDcInlineStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
