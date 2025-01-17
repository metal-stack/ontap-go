// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NsswitchSource Source location for a name service
//
// swagger:model nsswitch_source
type NsswitchSource string

func NewNsswitchSource(value NsswitchSource) *NsswitchSource {
	return &value
}

// Pointer returns a pointer to a freshly-allocated NsswitchSource.
func (m NsswitchSource) Pointer() *NsswitchSource {
	return &m
}

const (

	// NsswitchSourceFiles captures enum value "files"
	NsswitchSourceFiles NsswitchSource = "files"

	// NsswitchSourceDNS captures enum value "dns"
	NsswitchSourceDNS NsswitchSource = "dns"

	// NsswitchSourceLdap captures enum value "ldap"
	NsswitchSourceLdap NsswitchSource = "ldap"

	// NsswitchSourceNis captures enum value "nis"
	NsswitchSourceNis NsswitchSource = "nis"
)

// for schema
var nsswitchSourceEnum []interface{}

func init() {
	var res []NsswitchSource
	if err := json.Unmarshal([]byte(`["files","dns","ldap","nis"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nsswitchSourceEnum = append(nsswitchSourceEnum, v)
	}
}

func (m NsswitchSource) validateNsswitchSourceEnum(path, location string, value NsswitchSource) error {
	if err := validate.EnumCase(path, location, value, nsswitchSourceEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this nsswitch source
func (m NsswitchSource) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateNsswitchSourceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this nsswitch source based on context it is used
func (m NsswitchSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
