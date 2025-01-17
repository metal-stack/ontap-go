// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SecurityCertificateSign security certificate sign
//
// swagger:model security_certificate_sign
type SecurityCertificateSign struct {

	// Certificate expiration time, in ISO 8601 duration format or date and time format. The allowed expiration time range is between 1 day to 10 years.
	// Example: P1DT2H3M4S or '2030-01-25T11:20:13Z'
	ExpiryTime *string `json:"expiry_time,omitempty" yaml:"expiry_time,omitempty"`

	// Hashing function
	// Enum: ["sha256","sha224","sha384","sha512"]
	HashFunction *string `json:"hash_function,omitempty" yaml:"hash_function,omitempty"`

	// Certificate signing request to be signed by the given certificate authority. Request should be in X509 PEM format.
	// Example: '-----BEGIN CERTIFICATE REQUEST----- MIICYDCCAUgCAQAwGzEMMAoGA1UEAxMDQUJDMQswCQYDVQQGEwJVUzCCASIwDQYJ KoZIhvcNAQEBBQADggEPADCCAQoCggEBAPF+82SlqT3Vyu3Jx4IAwHcO5EGwLOxy zQ6KNjz71Fca0n1/A1CbCPyOsSupGVObvdWxX7xLVMJ2SXb7h43GCqYyX6FXJO4F HOpmLvB+jxdeiW7SDbiZyLUlsvA+oRO/uNlcug773QZdKLjJD64erZZMRUNbUJB8 bARxAUi0FPvgTraSQ0UW5sRLiGKeAyKA4wekYe1VgjHRTBizFbD4dI3njfva/2Bl jf+kkulgcLJTuJNtkgeimqMKyraYuleYcYk2K+C//0NuNOuPbDfTXCM7O61vik09 Szi8nLN7OXE9KoAA93U/BCpSfpl8XIb4cGnEr8hgVHOOtZSo+KZBFxMCAwEAAaAA MA0GCSqGSIb3DQEBCwUAA4IBAQC2vFYpvgsFrm5GnPx8tOBD1xsTyYjbWJMD8hAF lFrvF9Sw9QGCtDyacxkwgJhQx8l8JiIS5GOY6WWLBl9FMkLQNAhDL9xF3WF7vfYq RKgrz3bd/Vg96fsRZNYIPLGmoEaqLOh3FOCGc2VbdsR9PwOn3fwthxkIRd6ds6/q jc5cpSmVsCOgu+OKcpRXikYDbkWXfTZ1AhSfn6njBYFdZ9+PNAu/0JRQh5bX60nO 5heniTcAJLwUZP/CQ8nxHY0Wqy+1rAtM33d5cVmhUlBXQSIru/0ZkA/b9fK5Zv8E ZMADYUoEvIG59Vxhyci8lzYf+Mxl8qBSF+ZdC4yWhzDqZtM9 -----END CERTIFICATE REQUEST-----'
	SigningRequest *string `json:"signing_request,omitempty" yaml:"signing_request,omitempty"`
}

// Validate validates this security certificate sign
func (m *SecurityCertificateSign) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHashFunction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var securityCertificateSignTypeHashFunctionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha256","sha224","sha384","sha512"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityCertificateSignTypeHashFunctionPropEnum = append(securityCertificateSignTypeHashFunctionPropEnum, v)
	}
}

const (

	// SecurityCertificateSignHashFunctionSha256 captures enum value "sha256"
	SecurityCertificateSignHashFunctionSha256 string = "sha256"

	// SecurityCertificateSignHashFunctionSha224 captures enum value "sha224"
	SecurityCertificateSignHashFunctionSha224 string = "sha224"

	// SecurityCertificateSignHashFunctionSha384 captures enum value "sha384"
	SecurityCertificateSignHashFunctionSha384 string = "sha384"

	// SecurityCertificateSignHashFunctionSha512 captures enum value "sha512"
	SecurityCertificateSignHashFunctionSha512 string = "sha512"
)

// prop value enum
func (m *SecurityCertificateSign) validateHashFunctionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityCertificateSignTypeHashFunctionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SecurityCertificateSign) validateHashFunction(formats strfmt.Registry) error {
	if swag.IsZero(m.HashFunction) { // not required
		return nil
	}

	// value enum
	if err := m.validateHashFunctionEnum("hash_function", "body", *m.HashFunction); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this security certificate sign based on context it is used
func (m *SecurityCertificateSign) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecurityCertificateSign) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityCertificateSign) UnmarshalBinary(b []byte) error {
	var res SecurityCertificateSign
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
