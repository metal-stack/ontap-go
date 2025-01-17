// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ShadowcopyAddFiles shadowcopy add files
//
// swagger:model shadowcopy_add_files
type ShadowcopyAddFiles struct {

	// The universally-unique identifier of the storage's shadow copy set.
	// Example: f8328660-00e6-11e6-80d9-005056bd65a9
	StorageShadowcopySetUUID *string `json:"storage_shadowcopy_set_uuid,omitempty" yaml:"storage_shadowcopy_set_uuid,omitempty"`

	// The universally-unique identifier of the storage's shadow copy.
	// Example: fef32805-1f19-40ba-9b82-ebf277517e7e
	StorageShadowcopyUUID *string `json:"storage_shadowcopy_uuid,omitempty" yaml:"storage_shadowcopy_uuid,omitempty"`
}

// Validate validates this shadowcopy add files
func (m *ShadowcopyAddFiles) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this shadowcopy add files based on context it is used
func (m *ShadowcopyAddFiles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ShadowcopyAddFiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShadowcopyAddFiles) UnmarshalBinary(b []byte) error {
	var res ShadowcopyAddFiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
