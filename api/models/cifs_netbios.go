// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CifsNetbios cifs netbios
//
// swagger:model cifs_netbios
type CifsNetbios struct {

	// cifs netbios inline aliases
	// Example: ["ALIAS_1","ALIAS_2","ALIAS_3"]
	// Max Items: 200
	CifsNetbiosInlineAliases []*string `json:"aliases,omitempty" yaml:"aliases,omitempty"`

	// cifs netbios inline wins servers
	// Example: ["10.224.65.20","10.224.65.21"]
	// Max Items: 4
	CifsNetbiosInlineWinsServers []*string `json:"wins_servers,omitempty" yaml:"wins_servers,omitempty"`

	// Specifies whether NetBios name service (NBNS) is enabled for the CIFS. If this service is enabled, the CIFS server will start sending the broadcast for name registration.
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

// Validate validates this cifs netbios
func (m *CifsNetbios) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCifsNetbiosInlineAliases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCifsNetbiosInlineWinsServers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsNetbios) validateCifsNetbiosInlineAliases(formats strfmt.Registry) error {
	if swag.IsZero(m.CifsNetbiosInlineAliases) { // not required
		return nil
	}

	iCifsNetbiosInlineAliasesSize := int64(len(m.CifsNetbiosInlineAliases))

	if err := validate.MaxItems("aliases", "body", iCifsNetbiosInlineAliasesSize, 200); err != nil {
		return err
	}

	for i := 0; i < len(m.CifsNetbiosInlineAliases); i++ {
		if swag.IsZero(m.CifsNetbiosInlineAliases[i]) { // not required
			continue
		}

		if err := validate.MinLength("aliases"+"."+strconv.Itoa(i), "body", *m.CifsNetbiosInlineAliases[i], 1); err != nil {
			return err
		}

		if err := validate.MaxLength("aliases"+"."+strconv.Itoa(i), "body", *m.CifsNetbiosInlineAliases[i], 15); err != nil {
			return err
		}

	}

	return nil
}

func (m *CifsNetbios) validateCifsNetbiosInlineWinsServers(formats strfmt.Registry) error {
	if swag.IsZero(m.CifsNetbiosInlineWinsServers) { // not required
		return nil
	}

	iCifsNetbiosInlineWinsServersSize := int64(len(m.CifsNetbiosInlineWinsServers))

	if err := validate.MaxItems("wins_servers", "body", iCifsNetbiosInlineWinsServersSize, 4); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cifs netbios based on context it is used
func (m *CifsNetbios) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CifsNetbios) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsNetbios) UnmarshalBinary(b []byte) error {
	var res CifsNetbios
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}