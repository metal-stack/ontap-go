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

// ConsistencyGroupNvmeHostDhHmacChap A container for properties of NVMe in-band authentication with the DH-HMAC-CHAP protocol.
//
// swagger:model consistency_group_nvme_host_dh_hmac_chap
type ConsistencyGroupNvmeHostDhHmacChap struct {

	// The controller secret for NVMe in-band authentication. The value of this property is used by the NVMe host to authenticate the NVMe controller while establishing a connection. If unset, the controller is not authenticated. When supplied, the property `host_secret_key` must also be supplied. Optional in POST.<br/>
	// This property is write-only. The `mode` property can be used to identify if a controller secret has been set for the host, but the controller secret value cannot be read. To change the value, the host must be deleted from the subsystem and re-added.
	//
	// Example: DHHC-1:00:ia6zGodOr4SEG0Zzaw398rpY0wqipUWj4jWjUh4HWUz6aQ2n:
	ControllerSecretKey *string `json:"controller_secret_key,omitempty" yaml:"controller_secret_key,omitempty"`

	// The Diffie-Hellman group size for NVMe in-band authentication. When property `host_secret_key` is provided, this property defaults to `2048_bit`. When supplied, the property `host_secret_key` must also be supplied. Optional in POST.
	//
	// Enum: ["none","2048_bit","3072_bit","4096_bit","6144_bit","8192_bit"]
	GroupSize *string `json:"group_size,omitempty" yaml:"group_size,omitempty"`

	// The hash function for NVMe in-band authentication. When property `host_secret_key` is provided, this property defaults to `sha_256`. When supplied, the property `host_secret_key` must also be supplied. Optional in POST.
	//
	// Enum: ["sha_256","sha_512"]
	HashFunction *string `json:"hash_function,omitempty" yaml:"hash_function,omitempty"`

	// The host secret for NVMe in-band authentication. The value of this property is used by the NVMe controller to authenticate the NVMe host while establishing a connection. If unset, no authentication is performed by the host or controller. This property must be supplied if any other NVMe in-band authentication properties are supplied. Optional in POST.<br/>
	// This property is write-only. The `mode` property can be used to identify if a host secret has been set for the host, but the host secret value cannot be read. To change the value, the host must be deleted from the subsystem and re-added.
	//
	// Example: DHHC-1:00:ia6zGodOr4SEG0Zzaw398rpY0wqipUWj4jWjUh4HWUz6aQ2n:
	HostSecretKey *string `json:"host_secret_key,omitempty" yaml:"host_secret_key,omitempty"`
}

// Validate validates this consistency group nvme host dh hmac chap
func (m *ConsistencyGroupNvmeHostDhHmacChap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHashFunction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var consistencyGroupNvmeHostDhHmacChapTypeGroupSizePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","2048_bit","3072_bit","4096_bit","6144_bit","8192_bit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consistencyGroupNvmeHostDhHmacChapTypeGroupSizePropEnum = append(consistencyGroupNvmeHostDhHmacChapTypeGroupSizePropEnum, v)
	}
}

const (

	// ConsistencyGroupNvmeHostDhHmacChapGroupSizeNone captures enum value "none"
	ConsistencyGroupNvmeHostDhHmacChapGroupSizeNone string = "none"

	// ConsistencyGroupNvmeHostDhHmacChapGroupSizeNr2048Bit captures enum value "2048_bit"
	ConsistencyGroupNvmeHostDhHmacChapGroupSizeNr2048Bit string = "2048_bit"

	// ConsistencyGroupNvmeHostDhHmacChapGroupSizeNr3072Bit captures enum value "3072_bit"
	ConsistencyGroupNvmeHostDhHmacChapGroupSizeNr3072Bit string = "3072_bit"

	// ConsistencyGroupNvmeHostDhHmacChapGroupSizeNr4096Bit captures enum value "4096_bit"
	ConsistencyGroupNvmeHostDhHmacChapGroupSizeNr4096Bit string = "4096_bit"

	// ConsistencyGroupNvmeHostDhHmacChapGroupSizeNr6144Bit captures enum value "6144_bit"
	ConsistencyGroupNvmeHostDhHmacChapGroupSizeNr6144Bit string = "6144_bit"

	// ConsistencyGroupNvmeHostDhHmacChapGroupSizeNr8192Bit captures enum value "8192_bit"
	ConsistencyGroupNvmeHostDhHmacChapGroupSizeNr8192Bit string = "8192_bit"
)

// prop value enum
func (m *ConsistencyGroupNvmeHostDhHmacChap) validateGroupSizeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consistencyGroupNvmeHostDhHmacChapTypeGroupSizePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsistencyGroupNvmeHostDhHmacChap) validateGroupSize(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupSize) { // not required
		return nil
	}

	// value enum
	if err := m.validateGroupSizeEnum("group_size", "body", *m.GroupSize); err != nil {
		return err
	}

	return nil
}

var consistencyGroupNvmeHostDhHmacChapTypeHashFunctionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha_256","sha_512"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consistencyGroupNvmeHostDhHmacChapTypeHashFunctionPropEnum = append(consistencyGroupNvmeHostDhHmacChapTypeHashFunctionPropEnum, v)
	}
}

const (

	// ConsistencyGroupNvmeHostDhHmacChapHashFunctionSha256 captures enum value "sha_256"
	ConsistencyGroupNvmeHostDhHmacChapHashFunctionSha256 string = "sha_256"

	// ConsistencyGroupNvmeHostDhHmacChapHashFunctionSha512 captures enum value "sha_512"
	ConsistencyGroupNvmeHostDhHmacChapHashFunctionSha512 string = "sha_512"
)

// prop value enum
func (m *ConsistencyGroupNvmeHostDhHmacChap) validateHashFunctionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consistencyGroupNvmeHostDhHmacChapTypeHashFunctionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsistencyGroupNvmeHostDhHmacChap) validateHashFunction(formats strfmt.Registry) error {
	if swag.IsZero(m.HashFunction) { // not required
		return nil
	}

	// value enum
	if err := m.validateHashFunctionEnum("hash_function", "body", *m.HashFunction); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this consistency group nvme host dh hmac chap based on context it is used
func (m *ConsistencyGroupNvmeHostDhHmacChap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupNvmeHostDhHmacChap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupNvmeHostDhHmacChap) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupNvmeHostDhHmacChap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
