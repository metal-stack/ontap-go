// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NdmpScsi ndmp scsi
//
// swagger:model ndmp_scsi
type NdmpScsi struct {

	// Indicates the NDMP SCSI device ID.
	DeviceID *string `json:"device_id,omitempty" yaml:"device_id,omitempty"`

	// Indicates the NDMP SCSI host adapter.
	HostAdapter *int64 `json:"host_adapter,omitempty" yaml:"host_adapter,omitempty"`

	// Indicates the NDMP SCSI LUN ID.
	LunID *int64 `json:"lun_id,omitempty" yaml:"lun_id,omitempty"`

	// Indicates the NDMP SCSI target ID.
	TargetID *int64 `json:"target_id,omitempty" yaml:"target_id,omitempty"`
}

// Validate validates this ndmp scsi
func (m *NdmpScsi) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ndmp scsi based on context it is used
func (m *NdmpScsi) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NdmpScsi) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NdmpScsi) UnmarshalBinary(b []byte) error {
	var res NdmpScsi
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
