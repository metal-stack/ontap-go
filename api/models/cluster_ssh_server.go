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

// ClusterSSHServer cluster ssh server
//
// swagger:model cluster_ssh_server
type ClusterSSHServer struct {

	// links
	Links *ClusterSSHServerInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// Ciphers for encrypting the data.
	// Example: ["aes256_ctr","aes192_ctr","aes128_ctr"]
	ClusterSSHServerInlineCiphers []*Cipher `json:"ciphers,omitempty" yaml:"ciphers,omitempty"`

	// Key exchange algorithms.
	// Example: ["diffie_hellman_group_exchange_sha256","ecdh_sha2_nistp256"]
	ClusterSSHServerInlineKeyExchangeAlgorithms []*KeyExchangeAlgorithm `json:"key_exchange_algorithms,omitempty" yaml:"key_exchange_algorithms,omitempty"`

	// MAC algorithms.
	// Example: ["hmac_sha2_512","hmac_sha2_512_etm"]
	ClusterSSHServerInlineMacAlgorithms []*MacAlgorithm `json:"mac_algorithms,omitempty" yaml:"mac_algorithms,omitempty"`

	// Maximum connections allowed per second.
	// Maximum: 70
	// Minimum: 1
	ConnectionsPerSecond *int64 `json:"connections_per_second,omitempty" yaml:"connections_per_second,omitempty"`

	// Maximum authentication retries allowed before closing the connection.
	// Maximum: 6
	// Minimum: 2
	MaxAuthenticationRetryCount *int64 `json:"max_authentication_retry_count,omitempty" yaml:"max_authentication_retry_count,omitempty"`

	// Maximum possible simultaneous connections.
	// Maximum: 128
	// Minimum: 1
	MaxInstances *int64 `json:"max_instances,omitempty" yaml:"max_instances,omitempty"`

	// Maximum connections from the same client host.
	// Maximum: 64
	// Minimum: 1
	PerSourceLimit *int64 `json:"per_source_limit,omitempty" yaml:"per_source_limit,omitempty"`
}

// Validate validates this cluster ssh server
func (m *ClusterSSHServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterSSHServerInlineCiphers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterSSHServerInlineKeyExchangeAlgorithms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterSSHServerInlineMacAlgorithms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionsPerSecond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxAuthenticationRetryCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxInstances(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerSourceLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSSHServer) validateLinks(formats strfmt.Registry) error {
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

func (m *ClusterSSHServer) validateClusterSSHServerInlineCiphers(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterSSHServerInlineCiphers) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterSSHServerInlineCiphers); i++ {
		if swag.IsZero(m.ClusterSSHServerInlineCiphers[i]) { // not required
			continue
		}

		if m.ClusterSSHServerInlineCiphers[i] != nil {
			if err := m.ClusterSSHServerInlineCiphers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ciphers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ciphers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterSSHServer) validateClusterSSHServerInlineKeyExchangeAlgorithms(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterSSHServerInlineKeyExchangeAlgorithms) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterSSHServerInlineKeyExchangeAlgorithms); i++ {
		if swag.IsZero(m.ClusterSSHServerInlineKeyExchangeAlgorithms[i]) { // not required
			continue
		}

		if m.ClusterSSHServerInlineKeyExchangeAlgorithms[i] != nil {
			if err := m.ClusterSSHServerInlineKeyExchangeAlgorithms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("key_exchange_algorithms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("key_exchange_algorithms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterSSHServer) validateClusterSSHServerInlineMacAlgorithms(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterSSHServerInlineMacAlgorithms) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterSSHServerInlineMacAlgorithms); i++ {
		if swag.IsZero(m.ClusterSSHServerInlineMacAlgorithms[i]) { // not required
			continue
		}

		if m.ClusterSSHServerInlineMacAlgorithms[i] != nil {
			if err := m.ClusterSSHServerInlineMacAlgorithms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mac_algorithms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mac_algorithms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterSSHServer) validateConnectionsPerSecond(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectionsPerSecond) { // not required
		return nil
	}

	if err := validate.MinimumInt("connections_per_second", "body", *m.ConnectionsPerSecond, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("connections_per_second", "body", *m.ConnectionsPerSecond, 70, false); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSSHServer) validateMaxAuthenticationRetryCount(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxAuthenticationRetryCount) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_authentication_retry_count", "body", *m.MaxAuthenticationRetryCount, 2, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max_authentication_retry_count", "body", *m.MaxAuthenticationRetryCount, 6, false); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSSHServer) validateMaxInstances(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxInstances) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_instances", "body", *m.MaxInstances, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max_instances", "body", *m.MaxInstances, 128, false); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSSHServer) validatePerSourceLimit(formats strfmt.Registry) error {
	if swag.IsZero(m.PerSourceLimit) { // not required
		return nil
	}

	if err := validate.MinimumInt("per_source_limit", "body", *m.PerSourceLimit, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("per_source_limit", "body", *m.PerSourceLimit, 64, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cluster ssh server based on the context it is used
func (m *ClusterSSHServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterSSHServerInlineCiphers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterSSHServerInlineKeyExchangeAlgorithms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterSSHServerInlineMacAlgorithms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSSHServer) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterSSHServer) contextValidateClusterSSHServerInlineCiphers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterSSHServerInlineCiphers); i++ {

		if m.ClusterSSHServerInlineCiphers[i] != nil {

			if swag.IsZero(m.ClusterSSHServerInlineCiphers[i]) { // not required
				return nil
			}

			if err := m.ClusterSSHServerInlineCiphers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ciphers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ciphers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterSSHServer) contextValidateClusterSSHServerInlineKeyExchangeAlgorithms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterSSHServerInlineKeyExchangeAlgorithms); i++ {

		if m.ClusterSSHServerInlineKeyExchangeAlgorithms[i] != nil {

			if swag.IsZero(m.ClusterSSHServerInlineKeyExchangeAlgorithms[i]) { // not required
				return nil
			}

			if err := m.ClusterSSHServerInlineKeyExchangeAlgorithms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("key_exchange_algorithms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("key_exchange_algorithms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterSSHServer) contextValidateClusterSSHServerInlineMacAlgorithms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterSSHServerInlineMacAlgorithms); i++ {

		if m.ClusterSSHServerInlineMacAlgorithms[i] != nil {

			if swag.IsZero(m.ClusterSSHServerInlineMacAlgorithms[i]) { // not required
				return nil
			}

			if err := m.ClusterSSHServerInlineMacAlgorithms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mac_algorithms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mac_algorithms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSSHServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSSHServer) UnmarshalBinary(b []byte) error {
	var res ClusterSSHServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterSSHServerInlineLinks cluster ssh server inline links
//
// swagger:model cluster_ssh_server_inline__links
type ClusterSSHServerInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this cluster ssh server inline links
func (m *ClusterSSHServerInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSSHServerInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this cluster ssh server inline links based on the context it is used
func (m *ClusterSSHServerInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSSHServerInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClusterSSHServerInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSSHServerInlineLinks) UnmarshalBinary(b []byte) error {
	var res ClusterSSHServerInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}