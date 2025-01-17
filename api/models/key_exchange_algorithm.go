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

// KeyExchangeAlgorithm Key Exchange Algorithm.
// Example: diffie_hellman_group_exchange_sha256
//
// swagger:model key_exchange_algorithm
type KeyExchangeAlgorithm string

func NewKeyExchangeAlgorithm(value KeyExchangeAlgorithm) *KeyExchangeAlgorithm {
	return &value
}

// Pointer returns a pointer to a freshly-allocated KeyExchangeAlgorithm.
func (m KeyExchangeAlgorithm) Pointer() *KeyExchangeAlgorithm {
	return &m
}

const (

	// KeyExchangeAlgorithmDiffieHellmanGroupExchangeSha256 captures enum value "diffie_hellman_group_exchange_sha256"
	KeyExchangeAlgorithmDiffieHellmanGroupExchangeSha256 KeyExchangeAlgorithm = "diffie_hellman_group_exchange_sha256"

	// KeyExchangeAlgorithmCurve25519DashSha256 captures enum value "curve25519-sha256"
	KeyExchangeAlgorithmCurve25519DashSha256 KeyExchangeAlgorithm = "curve25519-sha256"

	// KeyExchangeAlgorithmEcdhSha2Nistp256 captures enum value "ecdh_sha2_nistp256"
	KeyExchangeAlgorithmEcdhSha2Nistp256 KeyExchangeAlgorithm = "ecdh_sha2_nistp256"

	// KeyExchangeAlgorithmEcdhSha2Nistp384 captures enum value "ecdh_sha2_nistp384"
	KeyExchangeAlgorithmEcdhSha2Nistp384 KeyExchangeAlgorithm = "ecdh_sha2_nistp384"

	// KeyExchangeAlgorithmEcdhSha2Nistp521 captures enum value "ecdh_sha2_nistp521"
	KeyExchangeAlgorithmEcdhSha2Nistp521 KeyExchangeAlgorithm = "ecdh_sha2_nistp521"
)

// for schema
var keyExchangeAlgorithmEnum []interface{}

func init() {
	var res []KeyExchangeAlgorithm
	if err := json.Unmarshal([]byte(`["diffie_hellman_group_exchange_sha256","curve25519-sha256","ecdh_sha2_nistp256","ecdh_sha2_nistp384","ecdh_sha2_nistp521"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		keyExchangeAlgorithmEnum = append(keyExchangeAlgorithmEnum, v)
	}
}

func (m KeyExchangeAlgorithm) validateKeyExchangeAlgorithmEnum(path, location string, value KeyExchangeAlgorithm) error {
	if err := validate.EnumCase(path, location, value, keyExchangeAlgorithmEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this key exchange algorithm
func (m KeyExchangeAlgorithm) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateKeyExchangeAlgorithmEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this key exchange algorithm based on context it is used
func (m KeyExchangeAlgorithm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
