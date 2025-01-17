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

// ApplicationLunObject LUN object
//
// swagger:model application_lun_object
type ApplicationLunObject struct {

	// LUN creation time
	// Read Only: true
	// Format: date-time
	CreationTimestamp *strfmt.DateTime `json:"creation_timestamp,omitempty" yaml:"creation_timestamp,omitempty"`

	// LUN path
	// Read Only: true
	Path *string `json:"path,omitempty" yaml:"path,omitempty"`

	// LUN size
	// Read Only: true
	Size *int64 `json:"size,omitempty" yaml:"size,omitempty"`

	// LUN UUID
	// Read Only: true
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this application lun object
func (m *ApplicationLunObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationLunObject) validateCreationTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_timestamp", "body", "date-time", m.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this application lun object based on the context it is used
func (m *ApplicationLunObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreationTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationLunObject) contextValidateCreationTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "creation_timestamp", "body", m.CreationTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationLunObject) contextValidatePath(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationLunObject) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationLunObject) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationLunObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationLunObject) UnmarshalBinary(b []byte) error {
	var res ApplicationLunObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
