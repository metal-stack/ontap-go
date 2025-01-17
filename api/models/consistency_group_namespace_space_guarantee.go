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

// ConsistencyGroupNamespaceSpaceGuarantee Properties that request and report the space guarantee for the NVMe namespace.
//
// swagger:model consistency_group_namespace_space_guarantee
type ConsistencyGroupNamespaceSpaceGuarantee struct {

	// The requested space reservation policy for the NVMe namespace. If _true_, a space reservation is requested for the namespace; if _false_, the namespace is thin provisioned. Guaranteeing a space reservation request for a namespace requires that the volume in which the namespace resides also be space reserved and that the fractional reserve for the volume be 100%.<br/>
	// The space reservation policy for an NVMe namespace is determined by ONTAP.
	//
	Requested *bool `json:"requested,omitempty" yaml:"requested,omitempty"`

	// Reports if the NVMe namespace is space guaranteed.<br/>
	// This property is _true_ if a space guarantee is requested and the containing volume and aggregate support the request. This property is _false_ if a space guarantee is not requested or if a space guarantee is requested and either the containing volume and aggregate do not support the request.
	//
	// Read Only: true
	Reserved *bool `json:"reserved,omitempty" yaml:"reserved,omitempty"`
}

// Validate validates this consistency group namespace space guarantee
func (m *ConsistencyGroupNamespaceSpaceGuarantee) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this consistency group namespace space guarantee based on the context it is used
func (m *ConsistencyGroupNamespaceSpaceGuarantee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReserved(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupNamespaceSpaceGuarantee) contextValidateReserved(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reserved", "body", m.Reserved); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupNamespaceSpaceGuarantee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupNamespaceSpaceGuarantee) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupNamespaceSpaceGuarantee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
