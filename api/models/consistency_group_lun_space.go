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

// ConsistencyGroupLunSpace The storage space related properties of the LUN.
//
// swagger:model consistency_group_lun_space
type ConsistencyGroupLunSpace struct {

	// guarantee
	Guarantee *ConsistencyGroupLunSpaceInlineGuarantee `json:"guarantee,omitempty" yaml:"guarantee,omitempty"`

	// The total provisioned size of the LUN. The LUN size can be increased but not reduced using the REST interface.
	// The maximum and minimum sizes listed here are the absolute maximum and absolute minimum sizes, in bytes. The actual minimum and maxiumum sizes vary depending on the ONTAP version, ONTAP platform, and the available space in the containing volume and aggregate.
	// For more information, see _Size properties_ in the _docs_ section of the ONTAP REST API documentation.
	//
	// Example: 1073741824
	// Maximum: 1.40737488355328e+14
	// Minimum: 4096
	Size *int64 `json:"size,omitempty" yaml:"size,omitempty"`

	// The amount of space consumed by the main data stream of the LUN.<br/>
	// This value is the total space consumed in the volume by the LUN, including filesystem overhead, but excluding prefix and suffix streams. Due to internal filesystem overhead and the many ways SAN filesystems and applications utilize blocks within a LUN, this value does not necessarily reflect actual consumption/availability from the perspective of the filesystem or application. Without specific knowledge of how the LUN blocks are utilized outside of ONTAP, this property should not be used as an indicator for an out-of-space condition.<br/>
	// For more information, see _Size properties_ in the _docs_ section of the ONTAP REST API documentation.
	//
	// Read Only: true
	Used *int64 `json:"used,omitempty" yaml:"used,omitempty"`
}

// Validate validates this consistency group lun space
func (m *ConsistencyGroupLunSpace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGuarantee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupLunSpace) validateGuarantee(formats strfmt.Registry) error {
	if swag.IsZero(m.Guarantee) { // not required
		return nil
	}

	if m.Guarantee != nil {
		if err := m.Guarantee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guarantee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("guarantee")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupLunSpace) validateSize(formats strfmt.Registry) error {
	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if err := validate.MinimumInt("size", "body", *m.Size, 4096, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("size", "body", *m.Size, 1.40737488355328e+14, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this consistency group lun space based on the context it is used
func (m *ConsistencyGroupLunSpace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGuarantee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupLunSpace) contextValidateGuarantee(ctx context.Context, formats strfmt.Registry) error {

	if m.Guarantee != nil {

		if swag.IsZero(m.Guarantee) { // not required
			return nil
		}

		if err := m.Guarantee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guarantee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("guarantee")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupLunSpace) contextValidateUsed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "used", "body", m.Used); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupLunSpace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupLunSpace) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupLunSpace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsistencyGroupLunSpaceInlineGuarantee Properties that request and report the space guarantee for the LUN.
//
// swagger:model consistency_group_lun_space_inline_guarantee
type ConsistencyGroupLunSpaceInlineGuarantee struct {

	// The requested space reservation policy for the LUN. If _true_, a space reservation is requested for the LUN; if _false_, the LUN is thin provisioned. Guaranteeing a space reservation request for a LUN requires that the volume in which the LUN resides is also space reserved and that the fractional reserve for the volume is 100%. Valid in POST and PATCH.
	//
	Requested *bool `json:"requested,omitempty" yaml:"requested,omitempty"`

	// Reports if the LUN is space guaranteed.<br/>
	// If _true_, a space guarantee is requested and the containing volume and aggregate support the request. If _false_, a space guarantee is not requested or a space guarantee is requested and either the containing volume or aggregate do not support the request.
	//
	// Read Only: true
	Reserved *bool `json:"reserved,omitempty" yaml:"reserved,omitempty"`
}

// Validate validates this consistency group lun space inline guarantee
func (m *ConsistencyGroupLunSpaceInlineGuarantee) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this consistency group lun space inline guarantee based on the context it is used
func (m *ConsistencyGroupLunSpaceInlineGuarantee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReserved(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupLunSpaceInlineGuarantee) contextValidateReserved(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "guarantee"+"."+"reserved", "body", m.Reserved); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupLunSpaceInlineGuarantee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupLunSpaceInlineGuarantee) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupLunSpaceInlineGuarantee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
