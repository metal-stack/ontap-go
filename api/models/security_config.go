// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SecurityConfig security config
//
// swagger:model security_config
type SecurityConfig struct {

	// links
	Links *SecurityConfigInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// fips
	Fips *SecurityConfigInlineFips `json:"fips,omitempty" yaml:"fips,omitempty"`

	// management protocols
	ManagementProtocols *SecurityConfigInlineManagementProtocols `json:"management_protocols,omitempty" yaml:"management_protocols,omitempty"`

	// onboard key manager configurable status
	OnboardKeyManagerConfigurableStatus *SecurityConfigInlineOnboardKeyManagerConfigurableStatus `json:"onboard_key_manager_configurable_status,omitempty" yaml:"onboard_key_manager_configurable_status,omitempty"`

	// software data encryption
	SoftwareDataEncryption *SecurityConfigInlineSoftwareDataEncryption `json:"software_data_encryption,omitempty" yaml:"software_data_encryption,omitempty"`

	// tls
	TLS *SecurityConfigInlineTLS `json:"tls,omitempty" yaml:"tls,omitempty"`
}

// Validate validates this security config
func (m *SecurityConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFips(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagementProtocols(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnboardKeyManagerConfigurableStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareDataEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityConfig) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityConfig) validateFips(formats strfmt.Registry) error {
	if swag.IsZero(m.Fips) { // not required
		return nil
	}

	if m.Fips != nil {
		if err := m.Fips.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fips")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fips")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityConfig) validateManagementProtocols(formats strfmt.Registry) error {
	if swag.IsZero(m.ManagementProtocols) { // not required
		return nil
	}

	if m.ManagementProtocols != nil {
		if err := m.ManagementProtocols.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("management_protocols")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("management_protocols")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityConfig) validateOnboardKeyManagerConfigurableStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.OnboardKeyManagerConfigurableStatus) { // not required
		return nil
	}

	if m.OnboardKeyManagerConfigurableStatus != nil {
		if err := m.OnboardKeyManagerConfigurableStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("onboard_key_manager_configurable_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("onboard_key_manager_configurable_status")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityConfig) validateSoftwareDataEncryption(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareDataEncryption) { // not required
		return nil
	}

	if m.SoftwareDataEncryption != nil {
		if err := m.SoftwareDataEncryption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_data_encryption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_data_encryption")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityConfig) validateTLS(formats strfmt.Registry) error {
	if swag.IsZero(m.TLS) { // not required
		return nil
	}

	if m.TLS != nil {
		if err := m.TLS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security config based on the context it is used
func (m *SecurityConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFips(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateManagementProtocols(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOnboardKeyManagerConfigurableStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareDataEncryption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTLS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityConfig) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityConfig) contextValidateFips(ctx context.Context, formats strfmt.Registry) error {

	if m.Fips != nil {

		if swag.IsZero(m.Fips) { // not required
			return nil
		}

		if err := m.Fips.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fips")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fips")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityConfig) contextValidateManagementProtocols(ctx context.Context, formats strfmt.Registry) error {

	if m.ManagementProtocols != nil {

		if swag.IsZero(m.ManagementProtocols) { // not required
			return nil
		}

		if err := m.ManagementProtocols.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("management_protocols")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("management_protocols")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityConfig) contextValidateOnboardKeyManagerConfigurableStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.OnboardKeyManagerConfigurableStatus != nil {

		if swag.IsZero(m.OnboardKeyManagerConfigurableStatus) { // not required
			return nil
		}

		if err := m.OnboardKeyManagerConfigurableStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("onboard_key_manager_configurable_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("onboard_key_manager_configurable_status")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityConfig) contextValidateSoftwareDataEncryption(ctx context.Context, formats strfmt.Registry) error {

	if m.SoftwareDataEncryption != nil {

		if swag.IsZero(m.SoftwareDataEncryption) { // not required
			return nil
		}

		if err := m.SoftwareDataEncryption.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_data_encryption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_data_encryption")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityConfig) contextValidateTLS(ctx context.Context, formats strfmt.Registry) error {

	if m.TLS != nil {

		if swag.IsZero(m.TLS) { // not required
			return nil
		}

		if err := m.TLS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityConfig) UnmarshalBinary(b []byte) error {
	var res SecurityConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityConfigInlineFips Cluster-wide Federal Information Processing Standards (FIPS) mode information.
//
// swagger:model security_config_inline_fips
type SecurityConfigInlineFips struct {

	// Indicates whether or not the software FIPS mode is enabled on the cluster. Our FIPS compliance involves configuring the use of only approved algorithms in applicable contexts (for example TLS), as well as the use of formally validated cryptographic module software implementations, where applicable. The US government documents concerning FIPS 140-2 outline the relevant security policies in detail.
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

// Validate validates this security config inline fips
func (m *SecurityConfigInlineFips) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this security config inline fips based on context it is used
func (m *SecurityConfigInlineFips) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecurityConfigInlineFips) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityConfigInlineFips) UnmarshalBinary(b []byte) error {
	var res SecurityConfigInlineFips
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityConfigInlineLinks security config inline links
//
// swagger:model security_config_inline__links
type SecurityConfigInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this security config inline links
func (m *SecurityConfigInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityConfigInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security config inline links based on the context it is used
func (m *SecurityConfigInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityConfigInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityConfigInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityConfigInlineLinks) UnmarshalBinary(b []byte) error {
	var res SecurityConfigInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityConfigInlineManagementProtocols Cluster-wide security protocols related information.
//
// swagger:model security_config_inline_management_protocols
type SecurityConfigInlineManagementProtocols struct {

	// Indicates whether or not security protocol rsh is enabled on the cluster.
	RshEnabled *bool `json:"rsh_enabled,omitempty" yaml:"rsh_enabled,omitempty"`

	// Indicates whether or not security protocol telnet is enabled on the cluster.
	TelnetEnabled *bool `json:"telnet_enabled,omitempty" yaml:"telnet_enabled,omitempty"`
}

// Validate validates this security config inline management protocols
func (m *SecurityConfigInlineManagementProtocols) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this security config inline management protocols based on context it is used
func (m *SecurityConfigInlineManagementProtocols) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecurityConfigInlineManagementProtocols) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityConfigInlineManagementProtocols) UnmarshalBinary(b []byte) error {
	var res SecurityConfigInlineManagementProtocols
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityConfigInlineOnboardKeyManagerConfigurableStatus Indicates whether the Onboard Key Manager can be configured in the cluster.
//
// swagger:model security_config_inline_onboard_key_manager_configurable_status
type SecurityConfigInlineOnboardKeyManagerConfigurableStatus struct {

	// Code corresponding to the status message. Returns a 0 if the Onboard Key Manager can be configured in the cluster.
	// Example: 65537300
	Code *int64 `json:"code,omitempty" yaml:"code,omitempty"`

	// Reason that Onboard Key Manager cannot be configured in the cluster.
	// Example: No platform support for volume encryption in following nodes - node1, node2.
	Message *string `json:"message,omitempty" yaml:"message,omitempty"`

	// Set to true if the Onboard Key Manager can be configured in the cluster.
	Supported *bool `json:"supported,omitempty" yaml:"supported,omitempty"`
}

// Validate validates this security config inline onboard key manager configurable status
func (m *SecurityConfigInlineOnboardKeyManagerConfigurableStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this security config inline onboard key manager configurable status based on context it is used
func (m *SecurityConfigInlineOnboardKeyManagerConfigurableStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecurityConfigInlineOnboardKeyManagerConfigurableStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityConfigInlineOnboardKeyManagerConfigurableStatus) UnmarshalBinary(b []byte) error {
	var res SecurityConfigInlineOnboardKeyManagerConfigurableStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityConfigInlineSoftwareDataEncryption Cluster-wide software data encryption related information.
//
// swagger:model security_config_inline_software_data_encryption
type SecurityConfigInlineSoftwareDataEncryption struct {

	// Indicates whether or not software encryption conversion is enabled on the cluster. A PATCH request initiates the conversion of all non-encrypted metadata volumes in the cluster to encrypted metadata volumes and all non-NAE aggregates to NAE aggregates. For the PATCH request to start, the cluster must have either an Onboard or an external key manager set up and the aggregates should either be empty or have only metadata volumes. No data volumes should be present in any of the aggregates in the cluster. For MetroCluster configurations, a PATCH request enables conversion on all the aggregates and metadata volumes of both local and remote clusters and is not allowed when the MetroCluster is in switchover state.
	ConversionEnabled *bool `json:"conversion_enabled,omitempty" yaml:"conversion_enabled,omitempty"`

	// Indicates whether or not default software data at rest encryption is disabled on the cluster.
	DisabledByDefault *bool `json:"disabled_by_default,omitempty" yaml:"disabled_by_default,omitempty"`
}

// Validate validates this security config inline software data encryption
func (m *SecurityConfigInlineSoftwareDataEncryption) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this security config inline software data encryption based on context it is used
func (m *SecurityConfigInlineSoftwareDataEncryption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecurityConfigInlineSoftwareDataEncryption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityConfigInlineSoftwareDataEncryption) UnmarshalBinary(b []byte) error {
	var res SecurityConfigInlineSoftwareDataEncryption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityConfigInlineTLS Cluster-wide Transport Layer Security (TLS) configuration information
//
// swagger:model security_config_inline_tls
type SecurityConfigInlineTLS struct {

	// Names a cipher suite that the system can select during TLS handshakes. A list of available options can be found on the Internet Assigned Number Authority (IANA) website.
	CipherSuites []*string `json:"cipher_suites,omitempty" yaml:"cipher_suites,omitempty"`

	// Names a TLS protocol version that the system can select during TLS handshakes. The use of SSLv3 or TLSv1 is discouraged.
	ProtocolVersions []*string `json:"protocol_versions,omitempty" yaml:"protocol_versions,omitempty"`
}

// Validate validates this security config inline tls
func (m *SecurityConfigInlineTLS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocolVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var securityConfigInlineTlsProtocolVersionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SSLv3","TLSv1","TLSv1.1","TLSv1.2","TLSv1.3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityConfigInlineTlsProtocolVersionsItemsEnum = append(securityConfigInlineTlsProtocolVersionsItemsEnum, v)
	}
}

func (m *SecurityConfigInlineTLS) validateProtocolVersionsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityConfigInlineTlsProtocolVersionsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SecurityConfigInlineTLS) validateProtocolVersions(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtocolVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.ProtocolVersions); i++ {
		if swag.IsZero(m.ProtocolVersions[i]) { // not required
			continue
		}

		// value enum
		if err := m.validateProtocolVersionsItemsEnum("tls"+"."+"protocol_versions"+"."+strconv.Itoa(i), "body", *m.ProtocolVersions[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this security config inline tls based on context it is used
func (m *SecurityConfigInlineTLS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecurityConfigInlineTLS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityConfigInlineTLS) UnmarshalBinary(b []byte) error {
	var res SecurityConfigInlineTLS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
