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

// MacAlgorithm MAC Algorithm.
// Example: hmac_sha2_256
//
// swagger:model mac_algorithm
type MacAlgorithm string

func NewMacAlgorithm(value MacAlgorithm) *MacAlgorithm {
	return &value
}

// Pointer returns a pointer to a freshly-allocated MacAlgorithm.
func (m MacAlgorithm) Pointer() *MacAlgorithm {
	return &m
}

const (

	// MacAlgorithmHmacSha2256 captures enum value "hmac_sha2_256"
	MacAlgorithmHmacSha2256 MacAlgorithm = "hmac_sha2_256"

	// MacAlgorithmHmacSha2512 captures enum value "hmac_sha2_512"
	MacAlgorithmHmacSha2512 MacAlgorithm = "hmac_sha2_512"

	// MacAlgorithmHmacSha2256Etm captures enum value "hmac_sha2_256_etm"
	MacAlgorithmHmacSha2256Etm MacAlgorithm = "hmac_sha2_256_etm"

	// MacAlgorithmHmacSha2512Etm captures enum value "hmac_sha2_512_etm"
	MacAlgorithmHmacSha2512Etm MacAlgorithm = "hmac_sha2_512_etm"

	// MacAlgorithmUmac64 captures enum value "umac_64"
	MacAlgorithmUmac64 MacAlgorithm = "umac_64"

	// MacAlgorithmUmac128 captures enum value "umac_128"
	MacAlgorithmUmac128 MacAlgorithm = "umac_128"

	// MacAlgorithmUmac64Etm captures enum value "umac_64_etm"
	MacAlgorithmUmac64Etm MacAlgorithm = "umac_64_etm"

	// MacAlgorithmUmac128Etm captures enum value "umac_128_etm"
	MacAlgorithmUmac128Etm MacAlgorithm = "umac_128_etm"
)

// for schema
var macAlgorithmEnum []interface{}

func init() {
	var res []MacAlgorithm
	if err := json.Unmarshal([]byte(`["hmac_sha2_256","hmac_sha2_512","hmac_sha2_256_etm","hmac_sha2_512_etm","umac_64","umac_128","umac_64_etm","umac_128_etm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		macAlgorithmEnum = append(macAlgorithmEnum, v)
	}
}

func (m MacAlgorithm) validateMacAlgorithmEnum(path, location string, value MacAlgorithm) error {
	if err := validate.EnumCase(path, location, value, macAlgorithmEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this mac algorithm
func (m MacAlgorithm) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMacAlgorithmEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this mac algorithm based on context it is used
func (m MacAlgorithm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
