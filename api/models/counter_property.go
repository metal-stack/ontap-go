// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CounterProperty Single string counter entry.
//
// swagger:model counter_property
type CounterProperty struct {

	// Property name.
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// Property value.
	Value *string `json:"value,omitempty" yaml:"value,omitempty"`
}

// Validate validates this counter property
func (m *CounterProperty) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this counter property based on context it is used
func (m *CounterProperty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CounterProperty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CounterProperty) UnmarshalBinary(b []byte) error {
	var res CounterProperty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
