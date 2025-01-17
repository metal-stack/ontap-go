// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Nis nis
//
// swagger:model nis
type Nis struct {

	// RFC 2307 nisMapEntry attribute.
	// Example: msSFU30NisMapEntry
	Mapentry *string `json:"mapentry,omitempty" yaml:"mapentry,omitempty"`

	// RFC 2307 nisMapName attribute.
	// Example: msSFU30NisMapName
	Mapname *string `json:"mapname,omitempty" yaml:"mapname,omitempty"`

	// RFC 2307 nisNetgroup object class.
	// Example: msSFU30NisNetGroup
	Netgroup *string `json:"netgroup,omitempty" yaml:"netgroup,omitempty"`

	// RFC 2307 nisNetgroupTriple attribute.
	// Example: msSFU30MemberOfNisNetgroup
	NetgroupTriple *string `json:"netgroup_triple,omitempty" yaml:"netgroup_triple,omitempty"`

	// RFC 2307 nisObject object class.
	// Example: msSFU30NisObject
	Object *string `json:"object,omitempty" yaml:"object,omitempty"`
}

// Validate validates this nis
func (m *Nis) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nis based on context it is used
func (m *Nis) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Nis) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Nis) UnmarshalBinary(b []byte) error {
	var res Nis
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
