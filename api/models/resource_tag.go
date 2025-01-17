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

// ResourceTag A resource tag is a way to group resources in the API together for identification or tracking purposes.
//
// swagger:model resource_tag
type ResourceTag struct {

	// The number of resources that are currently using this tag.
	// Minimum: 1
	NumResources *int64 `json:"num_resources,omitempty" yaml:"num_resources,omitempty"`

	// A key:value formatted string representing the tag's name.
	// Example: team:accounting
	// Max Length: 200
	Value *string `json:"value,omitempty" yaml:"value,omitempty"`
}

// Validate validates this resource tag
func (m *ResourceTag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceTag) validateNumResources(formats strfmt.Registry) error {
	if swag.IsZero(m.NumResources) { // not required
		return nil
	}

	if err := validate.MinimumInt("num_resources", "body", *m.NumResources, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *ResourceTag) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if err := validate.MaxLength("value", "body", *m.Value, 200); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this resource tag based on context it is used
func (m *ResourceTag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceTag) UnmarshalBinary(b []byte) error {
	var res ResourceTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
