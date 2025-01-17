// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmsParameterCriterion Criterion used for parameter based filtering
//
// swagger:model ems_parameter_criterion
type EmsParameterCriterion struct {

	// Parameter name pattern. Wildcard character '*' is supported.
	// Example: vol
	NamePattern *string `json:"name_pattern,omitempty" yaml:"name_pattern,omitempty"`

	// Parameter value pattern. Wildcard character '*' is supported.
	// Example: cloud*
	ValuePattern *string `json:"value_pattern,omitempty" yaml:"value_pattern,omitempty"`
}

// Validate validates this ems parameter criterion
func (m *EmsParameterCriterion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ems parameter criterion based on context it is used
func (m *EmsParameterCriterion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmsParameterCriterion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsParameterCriterion) UnmarshalBinary(b []byte) error {
	var res EmsParameterCriterion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
