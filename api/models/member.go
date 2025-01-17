// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Member member
//
// swagger:model member
type Member struct {

	// RFC 2307 memberNisNetgroup attribute.
	// Example: msSFU30MemberNisNetgroup
	NisNetgroup *string `json:"nis_netgroup,omitempty" yaml:"nis_netgroup,omitempty"`

	// RFC 2307 memberUid attribute.
	// Example: msSFU30MemberUid
	UID *string `json:"uid,omitempty" yaml:"uid,omitempty"`
}

// Validate validates this member
func (m *Member) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this member based on context it is used
func (m *Member) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Member) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Member) UnmarshalBinary(b []byte) error {
	var res Member
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
