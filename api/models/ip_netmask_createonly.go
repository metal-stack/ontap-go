// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// IPNetmaskCreateonly Input as netmask length (16) or IPv4 mask (255.255.0.0).
// Example: 24
//
// swagger:model ip_netmask_createonly
type IPNetmaskCreateonly string

// Validate validates this ip netmask createonly
func (m IPNetmaskCreateonly) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ip netmask createonly based on context it is used
func (m IPNetmaskCreateonly) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
