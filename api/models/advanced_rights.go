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

// AdvancedRights Specifies the advanced access right controlled by the ACE for the account specified.
//
//	You can specify more than one "advanced-rights" value by using a comma-delimited list.
//
// swagger:model advanced_rights
type AdvancedRights struct {

	// Append DAta
	AppendData *bool `json:"append_data,omitempty" yaml:"append_data,omitempty"`

	// Delete
	Delete *bool `json:"delete,omitempty" yaml:"delete,omitempty"`

	// Delete Child
	DeleteChild *bool `json:"delete_child,omitempty" yaml:"delete_child,omitempty"`

	// Execute File
	ExecuteFile *bool `json:"execute_file,omitempty" yaml:"execute_file,omitempty"`

	// Full Control
	FullControl *bool `json:"full_control,omitempty" yaml:"full_control,omitempty"`

	// Read Attributes
	ReadAttr *bool `json:"read_attr,omitempty" yaml:"read_attr,omitempty"`

	// Read Data
	ReadData *bool `json:"read_data,omitempty" yaml:"read_data,omitempty"`

	// Read Extended Attributes
	ReadEa *bool `json:"read_ea,omitempty" yaml:"read_ea,omitempty"`

	// Read Permissions
	ReadPerm *bool `json:"read_perm,omitempty" yaml:"read_perm,omitempty"`

	// Synchronize
	// Read Only: true
	Synchronize *bool `json:"synchronize,omitempty" yaml:"synchronize,omitempty"`

	// Write Attributes
	WriteAttr *bool `json:"write_attr,omitempty" yaml:"write_attr,omitempty"`

	// Write Data
	WriteData *bool `json:"write_data,omitempty" yaml:"write_data,omitempty"`

	// Write Extended Attributes
	WriteEa *bool `json:"write_ea,omitempty" yaml:"write_ea,omitempty"`

	// Write Owner
	WriteOwner *bool `json:"write_owner,omitempty" yaml:"write_owner,omitempty"`

	// Write Permission
	WritePerm *bool `json:"write_perm,omitempty" yaml:"write_perm,omitempty"`
}

// Validate validates this advanced rights
func (m *AdvancedRights) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this advanced rights based on the context it is used
func (m *AdvancedRights) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSynchronize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdvancedRights) contextValidateSynchronize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "synchronize", "body", m.Synchronize); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdvancedRights) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdvancedRights) UnmarshalBinary(b []byte) error {
	var res AdvancedRights
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}