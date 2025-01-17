// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Cn cn
//
// swagger:model cn
type Cn struct {

	// RFC 2256 cn attribute used by RFC 2307 when working with groups.
	// Example: cn
	Group *string `json:"group,omitempty" yaml:"group,omitempty"`

	// RFC 2256 cn attribute used by RFC 2307 when working with netgroups.
	// Example: name
	Netgroup *string `json:"netgroup,omitempty" yaml:"netgroup,omitempty"`
}

// Validate validates this cn
func (m *Cn) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cn based on context it is used
func (m *Cn) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Cn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cn) UnmarshalBinary(b []byte) error {
	var res Cn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}