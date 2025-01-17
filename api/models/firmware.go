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
)

// Firmware firmware
//
// swagger:model firmware
type Firmware struct {

	// disk
	// Read Only: true
	Disk *FirmwareDisk `json:"disk,omitempty" yaml:"disk,omitempty"`

	// dqp
	// Read Only: true
	Dqp *FirmwareDqp `json:"dqp,omitempty" yaml:"dqp,omitempty"`

	// firmware inline cluster fw progress
	FirmwareInlineClusterFwProgress []*FirmwareUpdateProgress `json:"cluster_fw_progress,omitempty" yaml:"cluster_fw_progress,omitempty"`

	// shelf
	// Read Only: true
	Shelf *FirmwareShelf `json:"shelf,omitempty" yaml:"shelf,omitempty"`

	// sp bmc
	// Read Only: true
	SpBmc *FirmwareSpBmc `json:"sp_bmc,omitempty" yaml:"sp_bmc,omitempty"`
}

// Validate validates this firmware
func (m *Firmware) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDqp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirmwareInlineClusterFwProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShelf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpBmc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Firmware) validateDisk(formats strfmt.Registry) error {
	if swag.IsZero(m.Disk) { // not required
		return nil
	}

	if m.Disk != nil {
		if err := m.Disk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk")
			}
			return err
		}
	}

	return nil
}

func (m *Firmware) validateDqp(formats strfmt.Registry) error {
	if swag.IsZero(m.Dqp) { // not required
		return nil
	}

	if m.Dqp != nil {
		if err := m.Dqp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dqp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dqp")
			}
			return err
		}
	}

	return nil
}

func (m *Firmware) validateFirmwareInlineClusterFwProgress(formats strfmt.Registry) error {
	if swag.IsZero(m.FirmwareInlineClusterFwProgress) { // not required
		return nil
	}

	for i := 0; i < len(m.FirmwareInlineClusterFwProgress); i++ {
		if swag.IsZero(m.FirmwareInlineClusterFwProgress[i]) { // not required
			continue
		}

		if m.FirmwareInlineClusterFwProgress[i] != nil {
			if err := m.FirmwareInlineClusterFwProgress[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_fw_progress" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster_fw_progress" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Firmware) validateShelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Shelf) { // not required
		return nil
	}

	if m.Shelf != nil {
		if err := m.Shelf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shelf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shelf")
			}
			return err
		}
	}

	return nil
}

func (m *Firmware) validateSpBmc(formats strfmt.Registry) error {
	if swag.IsZero(m.SpBmc) { // not required
		return nil
	}

	if m.SpBmc != nil {
		if err := m.SpBmc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sp_bmc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sp_bmc")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this firmware based on the context it is used
func (m *Firmware) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDqp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirmwareInlineClusterFwProgress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpBmc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Firmware) contextValidateDisk(ctx context.Context, formats strfmt.Registry) error {

	if m.Disk != nil {

		if swag.IsZero(m.Disk) { // not required
			return nil
		}

		if err := m.Disk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk")
			}
			return err
		}
	}

	return nil
}

func (m *Firmware) contextValidateDqp(ctx context.Context, formats strfmt.Registry) error {

	if m.Dqp != nil {

		if swag.IsZero(m.Dqp) { // not required
			return nil
		}

		if err := m.Dqp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dqp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dqp")
			}
			return err
		}
	}

	return nil
}

func (m *Firmware) contextValidateFirmwareInlineClusterFwProgress(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FirmwareInlineClusterFwProgress); i++ {

		if m.FirmwareInlineClusterFwProgress[i] != nil {

			if swag.IsZero(m.FirmwareInlineClusterFwProgress[i]) { // not required
				return nil
			}

			if err := m.FirmwareInlineClusterFwProgress[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_fw_progress" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster_fw_progress" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Firmware) contextValidateShelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Shelf != nil {

		if swag.IsZero(m.Shelf) { // not required
			return nil
		}

		if err := m.Shelf.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shelf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shelf")
			}
			return err
		}
	}

	return nil
}

func (m *Firmware) contextValidateSpBmc(ctx context.Context, formats strfmt.Registry) error {

	if m.SpBmc != nil {

		if swag.IsZero(m.SpBmc) { // not required
			return nil
		}

		if err := m.SpBmc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sp_bmc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sp_bmc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Firmware) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Firmware) UnmarshalBinary(b []byte) error {
	var res Firmware
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
