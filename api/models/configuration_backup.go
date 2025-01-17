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

// ConfigurationBackup configuration backup
//
// swagger:model configuration_backup
type ConfigurationBackup struct {

	// password
	// Example: yourpassword
	// Format: password
	Password *strfmt.Password `json:"password,omitempty" yaml:"password,omitempty"`

	// An external backup location for the cluster configuration. This is mostly required for single node clusters where node and cluster configuration backups cannot be copied to other nodes in the cluster.
	// Example: http://10.224.65.198/backups
	URL *string `json:"url,omitempty" yaml:"url,omitempty"`

	// username
	// Example: me
	Username *string `json:"username,omitempty" yaml:"username,omitempty"`

	// Use this parameter with the value "true" to validate the digital certificate of the remote server. Digital certificate validation is available only when the HTTPS protocol is used in the URL; it is disabled by default.
	ValidateCertificate *bool `json:"validate_certificate,omitempty" yaml:"validate_certificate,omitempty"`
}

// Validate validates this configuration backup
func (m *ConfigurationBackup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigurationBackup) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.FormatOf("password", "body", "password", m.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this configuration backup based on context it is used
func (m *ConfigurationBackup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigurationBackup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigurationBackup) UnmarshalBinary(b []byte) error {
	var res ConfigurationBackup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
