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

// TapeDevice tape device
//
// swagger:model tape_device
type TapeDevice struct {

	// alias
	Alias *TapeDeviceInlineAlias `json:"alias,omitempty" yaml:"alias,omitempty"`

	// Block number.
	// Example: 0
	// Read Only: true
	BlockNumber *int64 `json:"block_number,omitempty" yaml:"block_number,omitempty"`

	// Density.
	// Example: low
	// Read Only: true
	// Enum: ["low","medium","high","extended"]
	Density *string `json:"density,omitempty" yaml:"density,omitempty"`

	// description
	// Example: QUANTUM LTO-8 ULTRIUM
	// Read Only: true
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`

	// device id
	// Example: 1a.0
	// Read Only: true
	DeviceID *string `json:"device_id,omitempty" yaml:"device_id,omitempty"`

	// Operational state of the device.
	// Example: read_write_enabled
	// Read Only: true
	// Enum: ["unknown","available","ready_write_enabled","ready_write_protected","offline","in_use","error","reserved_by_another_host","normal","rewinding","erasing"]
	DeviceState *string `json:"device_state,omitempty" yaml:"device_state,omitempty"`

	// File number.
	// Example: 0
	// Read Only: true
	FileNumber *int64 `json:"file_number,omitempty" yaml:"file_number,omitempty"`

	// Device interface type.
	// Example: sas
	// Read Only: true
	// Enum: ["unknown","fibre_channel","sas","pscsi"]
	Interface *string `json:"interface,omitempty" yaml:"interface,omitempty"`

	// node
	Node *TapeDeviceInlineNode `json:"node,omitempty" yaml:"node,omitempty"`

	// online
	Online *bool `json:"online,omitempty" yaml:"online,omitempty"`

	// position
	Position *TapeDeviceInlinePosition `json:"position,omitempty" yaml:"position,omitempty"`

	// reservation type
	// Example: off
	// Read Only: true
	// Enum: ["off","persistent","scsi"]
	ReservationType *string `json:"reservation_type,omitempty" yaml:"reservation_type,omitempty"`

	// Residual count of the last I/O operation.
	// Example: 0
	// Read Only: true
	ResidualCount *int64 `json:"residual_count,omitempty" yaml:"residual_count,omitempty"`

	// serial number
	// Example: 10WT00093
	// Read Only: true
	SerialNumber *string `json:"serial_number,omitempty" yaml:"serial_number,omitempty"`

	// storage port
	StoragePort *TapeDeviceInlineStoragePort `json:"storage_port,omitempty" yaml:"storage_port,omitempty"`

	// tape device inline aliases
	TapeDeviceInlineAliases []*TapeDeviceInlineAliasesInlineArrayItem `json:"aliases,omitempty" yaml:"aliases,omitempty"`

	// tape device inline device names
	// Read Only: true
	TapeDeviceInlineDeviceNames []*TapeDeviceInlineDeviceNamesInlineArrayItem `json:"device_names,omitempty" yaml:"device_names,omitempty"`

	// Tape cartridge format.
	// Example: ["LTO-7 6TB","LTO-7 15TB Compressed","LTO-8 12TB","LTO-8 30TB Compressed"]
	// Read Only: true
	TapeDeviceInlineFormats []*string `json:"formats,omitempty" yaml:"formats,omitempty"`

	// Device type.
	// Example: tape
	// Read Only: true
	// Enum: ["unknown","tape","media_changer"]
	Type *string `json:"type,omitempty" yaml:"type,omitempty"`

	// World Wide Node Name.
	// Example: 500507631295741c
	// Read Only: true
	Wwnn *string `json:"wwnn,omitempty" yaml:"wwnn,omitempty"`

	// World Wide Port Name.
	// Example: 500507631295741c
	// Read Only: true
	Wwpn *string `json:"wwpn,omitempty" yaml:"wwpn,omitempty"`
}

// Validate validates this tape device
func (m *TapeDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlias(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDensity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReservationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoragePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTapeDeviceInlineAliases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTapeDeviceInlineDeviceNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapeDevice) validateAlias(formats strfmt.Registry) error {
	if swag.IsZero(m.Alias) { // not required
		return nil
	}

	if m.Alias != nil {
		if err := m.Alias.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alias")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alias")
			}
			return err
		}
	}

	return nil
}

var tapeDeviceTypeDensityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["low","medium","high","extended"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapeDeviceTypeDensityPropEnum = append(tapeDeviceTypeDensityPropEnum, v)
	}
}

const (

	// TapeDeviceDensityLow captures enum value "low"
	TapeDeviceDensityLow string = "low"

	// TapeDeviceDensityMedium captures enum value "medium"
	TapeDeviceDensityMedium string = "medium"

	// TapeDeviceDensityHigh captures enum value "high"
	TapeDeviceDensityHigh string = "high"

	// TapeDeviceDensityExtended captures enum value "extended"
	TapeDeviceDensityExtended string = "extended"
)

// prop value enum
func (m *TapeDevice) validateDensityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tapeDeviceTypeDensityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TapeDevice) validateDensity(formats strfmt.Registry) error {
	if swag.IsZero(m.Density) { // not required
		return nil
	}

	// value enum
	if err := m.validateDensityEnum("density", "body", *m.Density); err != nil {
		return err
	}

	return nil
}

var tapeDeviceTypeDeviceStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","available","ready_write_enabled","ready_write_protected","offline","in_use","error","reserved_by_another_host","normal","rewinding","erasing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapeDeviceTypeDeviceStatePropEnum = append(tapeDeviceTypeDeviceStatePropEnum, v)
	}
}

const (

	// TapeDeviceDeviceStateUnknown captures enum value "unknown"
	TapeDeviceDeviceStateUnknown string = "unknown"

	// TapeDeviceDeviceStateAvailable captures enum value "available"
	TapeDeviceDeviceStateAvailable string = "available"

	// TapeDeviceDeviceStateReadyWriteEnabled captures enum value "ready_write_enabled"
	TapeDeviceDeviceStateReadyWriteEnabled string = "ready_write_enabled"

	// TapeDeviceDeviceStateReadyWriteProtected captures enum value "ready_write_protected"
	TapeDeviceDeviceStateReadyWriteProtected string = "ready_write_protected"

	// TapeDeviceDeviceStateOffline captures enum value "offline"
	TapeDeviceDeviceStateOffline string = "offline"

	// TapeDeviceDeviceStateInUse captures enum value "in_use"
	TapeDeviceDeviceStateInUse string = "in_use"

	// TapeDeviceDeviceStateError captures enum value "error"
	TapeDeviceDeviceStateError string = "error"

	// TapeDeviceDeviceStateReservedByAnotherHost captures enum value "reserved_by_another_host"
	TapeDeviceDeviceStateReservedByAnotherHost string = "reserved_by_another_host"

	// TapeDeviceDeviceStateNormal captures enum value "normal"
	TapeDeviceDeviceStateNormal string = "normal"

	// TapeDeviceDeviceStateRewinding captures enum value "rewinding"
	TapeDeviceDeviceStateRewinding string = "rewinding"

	// TapeDeviceDeviceStateErasing captures enum value "erasing"
	TapeDeviceDeviceStateErasing string = "erasing"
)

// prop value enum
func (m *TapeDevice) validateDeviceStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tapeDeviceTypeDeviceStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TapeDevice) validateDeviceState(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceState) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeviceStateEnum("device_state", "body", *m.DeviceState); err != nil {
		return err
	}

	return nil
}

var tapeDeviceTypeInterfacePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","fibre_channel","sas","pscsi"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapeDeviceTypeInterfacePropEnum = append(tapeDeviceTypeInterfacePropEnum, v)
	}
}

const (

	// TapeDeviceInterfaceUnknown captures enum value "unknown"
	TapeDeviceInterfaceUnknown string = "unknown"

	// TapeDeviceInterfaceFibreChannel captures enum value "fibre_channel"
	TapeDeviceInterfaceFibreChannel string = "fibre_channel"

	// TapeDeviceInterfaceSas captures enum value "sas"
	TapeDeviceInterfaceSas string = "sas"

	// TapeDeviceInterfacePscsi captures enum value "pscsi"
	TapeDeviceInterfacePscsi string = "pscsi"
)

// prop value enum
func (m *TapeDevice) validateInterfaceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tapeDeviceTypeInterfacePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TapeDevice) validateInterface(formats strfmt.Registry) error {
	if swag.IsZero(m.Interface) { // not required
		return nil
	}

	// value enum
	if err := m.validateInterfaceEnum("interface", "body", *m.Interface); err != nil {
		return err
	}

	return nil
}

func (m *TapeDevice) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *TapeDevice) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if m.Position != nil {
		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

var tapeDeviceTypeReservationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["off","persistent","scsi"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapeDeviceTypeReservationTypePropEnum = append(tapeDeviceTypeReservationTypePropEnum, v)
	}
}

const (

	// TapeDeviceReservationTypeOff captures enum value "off"
	TapeDeviceReservationTypeOff string = "off"

	// TapeDeviceReservationTypePersistent captures enum value "persistent"
	TapeDeviceReservationTypePersistent string = "persistent"

	// TapeDeviceReservationTypeScsi captures enum value "scsi"
	TapeDeviceReservationTypeScsi string = "scsi"
)

// prop value enum
func (m *TapeDevice) validateReservationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tapeDeviceTypeReservationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TapeDevice) validateReservationType(formats strfmt.Registry) error {
	if swag.IsZero(m.ReservationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateReservationTypeEnum("reservation_type", "body", *m.ReservationType); err != nil {
		return err
	}

	return nil
}

func (m *TapeDevice) validateStoragePort(formats strfmt.Registry) error {
	if swag.IsZero(m.StoragePort) { // not required
		return nil
	}

	if m.StoragePort != nil {
		if err := m.StoragePort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_port")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_port")
			}
			return err
		}
	}

	return nil
}

func (m *TapeDevice) validateTapeDeviceInlineAliases(formats strfmt.Registry) error {
	if swag.IsZero(m.TapeDeviceInlineAliases) { // not required
		return nil
	}

	for i := 0; i < len(m.TapeDeviceInlineAliases); i++ {
		if swag.IsZero(m.TapeDeviceInlineAliases[i]) { // not required
			continue
		}

		if m.TapeDeviceInlineAliases[i] != nil {
			if err := m.TapeDeviceInlineAliases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aliases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("aliases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapeDevice) validateTapeDeviceInlineDeviceNames(formats strfmt.Registry) error {
	if swag.IsZero(m.TapeDeviceInlineDeviceNames) { // not required
		return nil
	}

	for i := 0; i < len(m.TapeDeviceInlineDeviceNames); i++ {
		if swag.IsZero(m.TapeDeviceInlineDeviceNames[i]) { // not required
			continue
		}

		if m.TapeDeviceInlineDeviceNames[i] != nil {
			if err := m.TapeDeviceInlineDeviceNames[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("device_names" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("device_names" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var tapeDeviceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","tape","media_changer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapeDeviceTypeTypePropEnum = append(tapeDeviceTypeTypePropEnum, v)
	}
}

const (

	// TapeDeviceTypeUnknown captures enum value "unknown"
	TapeDeviceTypeUnknown string = "unknown"

	// TapeDeviceTypeTape captures enum value "tape"
	TapeDeviceTypeTape string = "tape"

	// TapeDeviceTypeMediaChanger captures enum value "media_changer"
	TapeDeviceTypeMediaChanger string = "media_changer"
)

// prop value enum
func (m *TapeDevice) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tapeDeviceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TapeDevice) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this tape device based on the context it is used
func (m *TapeDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlias(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBlockNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDensity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReservationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResidualCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSerialNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStoragePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTapeDeviceInlineAliases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTapeDeviceInlineDeviceNames(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTapeDeviceInlineFormats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwnn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapeDevice) contextValidateAlias(ctx context.Context, formats strfmt.Registry) error {

	if m.Alias != nil {

		if swag.IsZero(m.Alias) { // not required
			return nil
		}

		if err := m.Alias.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alias")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alias")
			}
			return err
		}
	}

	return nil
}

func (m *TapeDevice) contextValidateBlockNumber(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "block_number", "body", m.BlockNumber); err != nil {
		return err
	}

	return nil
}

func (m *TapeDevice) contextValidateDensity(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "density", "body", m.Density); err != nil {
		return err
	}

	return nil
}

func (m *TapeDevice) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *TapeDevice) contextValidateDeviceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "device_id", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

func (m *TapeDevice) contextValidateDeviceState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "device_state", "body", m.DeviceState); err != nil {
		return err
	}

	return nil
}

func (m *TapeDevice) contextValidateFileNumber(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "file_number", "body", m.FileNumber); err != nil {
		return err
	}

	return nil
}

func (m *TapeDevice) contextValidateInterface(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "interface", "body", m.Interface); err != nil {
		return err
	}

	return nil
}

func (m *TapeDevice) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {

		if swag.IsZero(m.Node) { // not required
			return nil
		}

		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *TapeDevice) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if m.Position != nil {

		if swag.IsZero(m.Position) { // not required
			return nil
		}

		if err := m.Position.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

func (m *TapeDevice) contextValidateReservationType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reservation_type", "body", m.ReservationType); err != nil {
		return err
	}

	return nil
}

func (m *TapeDevice) contextValidateResidualCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "residual_count", "body", m.ResidualCount); err != nil {
		return err
	}

	return nil
}

func (m *TapeDevice) contextValidateSerialNumber(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "serial_number", "body", m.SerialNumber); err != nil {
		return err
	}

	return nil
}

func (m *TapeDevice) contextValidateStoragePort(ctx context.Context, formats strfmt.Registry) error {

	if m.StoragePort != nil {

		if swag.IsZero(m.StoragePort) { // not required
			return nil
		}

		if err := m.StoragePort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_port")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_port")
			}
			return err
		}
	}

	return nil
}

func (m *TapeDevice) contextValidateTapeDeviceInlineAliases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TapeDeviceInlineAliases); i++ {

		if m.TapeDeviceInlineAliases[i] != nil {

			if swag.IsZero(m.TapeDeviceInlineAliases[i]) { // not required
				return nil
			}

			if err := m.TapeDeviceInlineAliases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aliases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("aliases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapeDevice) contextValidateTapeDeviceInlineDeviceNames(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "device_names", "body", []*TapeDeviceInlineDeviceNamesInlineArrayItem(m.TapeDeviceInlineDeviceNames)); err != nil {
		return err
	}

	for i := 0; i < len(m.TapeDeviceInlineDeviceNames); i++ {

		if m.TapeDeviceInlineDeviceNames[i] != nil {

			if swag.IsZero(m.TapeDeviceInlineDeviceNames[i]) { // not required
				return nil
			}

			if err := m.TapeDeviceInlineDeviceNames[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("device_names" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("device_names" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapeDevice) contextValidateTapeDeviceInlineFormats(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "formats", "body", []*string(m.TapeDeviceInlineFormats)); err != nil {
		return err
	}

	return nil
}

func (m *TapeDevice) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *TapeDevice) contextValidateWwnn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "wwnn", "body", m.Wwnn); err != nil {
		return err
	}

	return nil
}

func (m *TapeDevice) contextValidateWwpn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "wwpn", "body", m.Wwpn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapeDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapeDevice) UnmarshalBinary(b []byte) error {
	var res TapeDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapeDeviceInlineAlias tape device inline alias
//
// swagger:model tape_device_inline_alias
type TapeDeviceInlineAlias struct {

	// This field will no longer be supported in a future release. Use aliases.mapping instead.
	// Example: SN[10WT000933]
	Mapping *string `json:"mapping,omitempty" yaml:"mapping,omitempty"`

	// This field will no longer be supported in a future release. Use aliases.name instead.
	// Example: st6
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this tape device inline alias
func (m *TapeDeviceInlineAlias) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this tape device inline alias based on the context it is used
func (m *TapeDeviceInlineAlias) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapeDeviceInlineAlias) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapeDeviceInlineAlias) UnmarshalBinary(b []byte) error {
	var res TapeDeviceInlineAlias
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapeDeviceInlineAliasesInlineArrayItem tape device inline aliases inline array item
//
// swagger:model tape_device_inline_aliases_inline_array_item
type TapeDeviceInlineAliasesInlineArrayItem struct {

	// Alias mapping.
	// Example: SN[10WT000933]
	// Read Only: true
	Mapping *string `json:"mapping,omitempty" yaml:"mapping,omitempty"`

	// Alias name.
	// Example: st6
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this tape device inline aliases inline array item
func (m *TapeDeviceInlineAliasesInlineArrayItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this tape device inline aliases inline array item based on the context it is used
func (m *TapeDeviceInlineAliasesInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMapping(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapeDeviceInlineAliasesInlineArrayItem) contextValidateMapping(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "mapping", "body", m.Mapping); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapeDeviceInlineAliasesInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapeDeviceInlineAliasesInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res TapeDeviceInlineAliasesInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapeDeviceInlineDeviceNamesInlineArrayItem tape device inline device names inline array item
//
// swagger:model tape_device_inline_device_names_inline_array_item
type TapeDeviceInlineDeviceNamesInlineArrayItem struct {

	// Device name for no rewind.
	// Example: nrst6l
	NoRewindDevice *string `json:"no_rewind_device,omitempty" yaml:"no_rewind_device,omitempty"`

	// Device name for rewind.
	// Example: rst6l
	RewindDevice *string `json:"rewind_device,omitempty" yaml:"rewind_device,omitempty"`

	// Device name for unload or reload operations.
	// Example: urst6l
	UnloadReloadDevice *string `json:"unload_reload_device,omitempty" yaml:"unload_reload_device,omitempty"`
}

// Validate validates this tape device inline device names inline array item
func (m *TapeDeviceInlineDeviceNamesInlineArrayItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tape device inline device names inline array item based on context it is used
func (m *TapeDeviceInlineDeviceNamesInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapeDeviceInlineDeviceNamesInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapeDeviceInlineDeviceNamesInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res TapeDeviceInlineDeviceNamesInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapeDeviceInlineNode tape device inline node
//
// swagger:model tape_device_inline_node
type TapeDeviceInlineNode struct {

	// links
	Links *TapeDeviceInlineNodeInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this tape device inline node
func (m *TapeDeviceInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapeDeviceInlineNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tape device inline node based on the context it is used
func (m *TapeDeviceInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapeDeviceInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapeDeviceInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapeDeviceInlineNode) UnmarshalBinary(b []byte) error {
	var res TapeDeviceInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapeDeviceInlineNodeInlineLinks tape device inline node inline links
//
// swagger:model tape_device_inline_node_inline__links
type TapeDeviceInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this tape device inline node inline links
func (m *TapeDeviceInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapeDeviceInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tape device inline node inline links based on the context it is used
func (m *TapeDeviceInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapeDeviceInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapeDeviceInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapeDeviceInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res TapeDeviceInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapeDeviceInlinePosition tape device inline position
//
// swagger:model tape_device_inline_position
type TapeDeviceInlinePosition struct {

	// Number of times to run position operation.
	// Example: 5
	Count *int64 `json:"count,omitempty" yaml:"count,omitempty"`

	// Position operation.
	// Example: rewind
	// Enum: ["weof","fsf","bsf","fsr","bsr","rewind","erase","eom"]
	Operation *string `json:"operation,omitempty" yaml:"operation,omitempty"`
}

// Validate validates this tape device inline position
func (m *TapeDeviceInlinePosition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tapeDeviceInlinePositionTypeOperationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["weof","fsf","bsf","fsr","bsr","rewind","erase","eom"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapeDeviceInlinePositionTypeOperationPropEnum = append(tapeDeviceInlinePositionTypeOperationPropEnum, v)
	}
}

const (

	// TapeDeviceInlinePositionOperationWeof captures enum value "weof"
	TapeDeviceInlinePositionOperationWeof string = "weof"

	// TapeDeviceInlinePositionOperationFsf captures enum value "fsf"
	TapeDeviceInlinePositionOperationFsf string = "fsf"

	// TapeDeviceInlinePositionOperationBsf captures enum value "bsf"
	TapeDeviceInlinePositionOperationBsf string = "bsf"

	// TapeDeviceInlinePositionOperationFsr captures enum value "fsr"
	TapeDeviceInlinePositionOperationFsr string = "fsr"

	// TapeDeviceInlinePositionOperationBsr captures enum value "bsr"
	TapeDeviceInlinePositionOperationBsr string = "bsr"

	// TapeDeviceInlinePositionOperationRewind captures enum value "rewind"
	TapeDeviceInlinePositionOperationRewind string = "rewind"

	// TapeDeviceInlinePositionOperationErase captures enum value "erase"
	TapeDeviceInlinePositionOperationErase string = "erase"

	// TapeDeviceInlinePositionOperationEom captures enum value "eom"
	TapeDeviceInlinePositionOperationEom string = "eom"
)

// prop value enum
func (m *TapeDeviceInlinePosition) validateOperationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tapeDeviceInlinePositionTypeOperationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TapeDeviceInlinePosition) validateOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperationEnum("position"+"."+"operation", "body", *m.Operation); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this tape device inline position based on context it is used
func (m *TapeDeviceInlinePosition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapeDeviceInlinePosition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapeDeviceInlinePosition) UnmarshalBinary(b []byte) error {
	var res TapeDeviceInlinePosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapeDeviceInlineStoragePort tape device inline storage port
//
// swagger:model tape_device_inline_storage_port
type TapeDeviceInlineStoragePort struct {

	// Initiator port.
	// Example: 2b
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this tape device inline storage port
func (m *TapeDeviceInlineStoragePort) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this tape device inline storage port based on the context it is used
func (m *TapeDeviceInlineStoragePort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapeDeviceInlineStoragePort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapeDeviceInlineStoragePort) UnmarshalBinary(b []byte) error {
	var res TapeDeviceInlineStoragePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
