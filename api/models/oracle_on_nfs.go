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

// OracleOnNfs Oracle using NFS.
//
// swagger:model oracle_on_nfs
type OracleOnNfs struct {

	// archive log
	ArchiveLog *OracleOnNfsArchiveLog `json:"archive_log,omitempty" yaml:"archive_log,omitempty"`

	// db
	// Required: true
	Db *OracleOnNfsDb `json:"db" yaml:"db"`

	// The list of NFS access controls. You must provide either 'host' or 'access' to enable NFS access.
	NfsAccess []*AppNfsAccess `json:"nfs_access" yaml:"nfs_access"`

	// ora home
	OraHome *OracleOnNfsOraHome `json:"ora_home,omitempty" yaml:"ora_home,omitempty"`

	// protection type
	ProtectionType *OracleOnNfsProtectionType `json:"protection_type,omitempty" yaml:"protection_type,omitempty"`

	// redo log
	// Required: true
	RedoLog *OracleOnNfsRedoLog `json:"redo_log" yaml:"redo_log"`
}

// Validate validates this oracle on nfs
func (m *OracleOnNfs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchiveLog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfsAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOraHome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedoLog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleOnNfs) validateArchiveLog(formats strfmt.Registry) error {
	if swag.IsZero(m.ArchiveLog) { // not required
		return nil
	}

	if m.ArchiveLog != nil {
		if err := m.ArchiveLog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archive_log")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("archive_log")
			}
			return err
		}
	}

	return nil
}

func (m *OracleOnNfs) validateDb(formats strfmt.Registry) error {

	if err := validate.Required("db", "body", m.Db); err != nil {
		return err
	}

	if m.Db != nil {
		if err := m.Db.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("db")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("db")
			}
			return err
		}
	}

	return nil
}

func (m *OracleOnNfs) validateNfsAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.NfsAccess) { // not required
		return nil
	}

	for i := 0; i < len(m.NfsAccess); i++ {
		if swag.IsZero(m.NfsAccess[i]) { // not required
			continue
		}

		if m.NfsAccess[i] != nil {
			if err := m.NfsAccess[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nfs_access" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nfs_access" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OracleOnNfs) validateOraHome(formats strfmt.Registry) error {
	if swag.IsZero(m.OraHome) { // not required
		return nil
	}

	if m.OraHome != nil {
		if err := m.OraHome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ora_home")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ora_home")
			}
			return err
		}
	}

	return nil
}

func (m *OracleOnNfs) validateProtectionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionType) { // not required
		return nil
	}

	if m.ProtectionType != nil {
		if err := m.ProtectionType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protection_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protection_type")
			}
			return err
		}
	}

	return nil
}

func (m *OracleOnNfs) validateRedoLog(formats strfmt.Registry) error {

	if err := validate.Required("redo_log", "body", m.RedoLog); err != nil {
		return err
	}

	if m.RedoLog != nil {
		if err := m.RedoLog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redo_log")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redo_log")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this oracle on nfs based on the context it is used
func (m *OracleOnNfs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArchiveLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNfsAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOraHome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtectionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRedoLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleOnNfs) contextValidateArchiveLog(ctx context.Context, formats strfmt.Registry) error {

	if m.ArchiveLog != nil {

		if swag.IsZero(m.ArchiveLog) { // not required
			return nil
		}

		if err := m.ArchiveLog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archive_log")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("archive_log")
			}
			return err
		}
	}

	return nil
}

func (m *OracleOnNfs) contextValidateDb(ctx context.Context, formats strfmt.Registry) error {

	if m.Db != nil {

		if err := m.Db.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("db")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("db")
			}
			return err
		}
	}

	return nil
}

func (m *OracleOnNfs) contextValidateNfsAccess(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NfsAccess); i++ {

		if m.NfsAccess[i] != nil {

			if swag.IsZero(m.NfsAccess[i]) { // not required
				return nil
			}

			if err := m.NfsAccess[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nfs_access" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nfs_access" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OracleOnNfs) contextValidateOraHome(ctx context.Context, formats strfmt.Registry) error {

	if m.OraHome != nil {

		if swag.IsZero(m.OraHome) { // not required
			return nil
		}

		if err := m.OraHome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ora_home")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ora_home")
			}
			return err
		}
	}

	return nil
}

func (m *OracleOnNfs) contextValidateProtectionType(ctx context.Context, formats strfmt.Registry) error {

	if m.ProtectionType != nil {

		if swag.IsZero(m.ProtectionType) { // not required
			return nil
		}

		if err := m.ProtectionType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protection_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protection_type")
			}
			return err
		}
	}

	return nil
}

func (m *OracleOnNfs) contextValidateRedoLog(ctx context.Context, formats strfmt.Registry) error {

	if m.RedoLog != nil {

		if err := m.RedoLog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redo_log")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redo_log")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OracleOnNfs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleOnNfs) UnmarshalBinary(b []byte) error {
	var res OracleOnNfs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OracleOnNfsArchiveLog oracle on nfs archive log
//
// swagger:model OracleOnNfsArchiveLog
type OracleOnNfsArchiveLog struct {

	// The size of the archive log. Usage: {&lt;integer&gt;[KB|MB|GB|TB|PB]}
	Size int64 `json:"size,omitempty" yaml:"size,omitempty"`

	// storage service
	StorageService *OracleOnNfsArchiveLogStorageService `json:"storage_service,omitempty" yaml:"storage_service,omitempty"`
}

// Validate validates this oracle on nfs archive log
func (m *OracleOnNfsArchiveLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStorageService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleOnNfsArchiveLog) validateStorageService(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageService) { // not required
		return nil
	}

	if m.StorageService != nil {
		if err := m.StorageService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archive_log" + "." + "storage_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("archive_log" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this oracle on nfs archive log based on the context it is used
func (m *OracleOnNfsArchiveLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleOnNfsArchiveLog) contextValidateStorageService(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageService != nil {

		if swag.IsZero(m.StorageService) { // not required
			return nil
		}

		if err := m.StorageService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archive_log" + "." + "storage_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("archive_log" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OracleOnNfsArchiveLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleOnNfsArchiveLog) UnmarshalBinary(b []byte) error {
	var res OracleOnNfsArchiveLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OracleOnNfsArchiveLogStorageService oracle on nfs archive log storage service
//
// swagger:model OracleOnNfsArchiveLogStorageService
type OracleOnNfsArchiveLogStorageService struct {

	// The storage service of the archive log.
	// Enum: ["extreme","performance","value"]
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this oracle on nfs archive log storage service
func (m *OracleOnNfsArchiveLogStorageService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var oracleOnNfsArchiveLogStorageServiceTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extreme","performance","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oracleOnNfsArchiveLogStorageServiceTypeNamePropEnum = append(oracleOnNfsArchiveLogStorageServiceTypeNamePropEnum, v)
	}
}

const (

	// OracleOnNfsArchiveLogStorageServiceNameExtreme captures enum value "extreme"
	OracleOnNfsArchiveLogStorageServiceNameExtreme string = "extreme"

	// OracleOnNfsArchiveLogStorageServiceNamePerformance captures enum value "performance"
	OracleOnNfsArchiveLogStorageServiceNamePerformance string = "performance"

	// OracleOnNfsArchiveLogStorageServiceNameValue captures enum value "value"
	OracleOnNfsArchiveLogStorageServiceNameValue string = "value"
)

// prop value enum
func (m *OracleOnNfsArchiveLogStorageService) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oracleOnNfsArchiveLogStorageServiceTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OracleOnNfsArchiveLogStorageService) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("archive_log"+"."+"storage_service"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this oracle on nfs archive log storage service based on context it is used
func (m *OracleOnNfsArchiveLogStorageService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OracleOnNfsArchiveLogStorageService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleOnNfsArchiveLogStorageService) UnmarshalBinary(b []byte) error {
	var res OracleOnNfsArchiveLogStorageService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OracleOnNfsDb oracle on nfs db
//
// swagger:model OracleOnNfsDb
type OracleOnNfsDb struct {

	// The size of the database. Usage: {&lt;integer&gt;[KB|MB|GB|TB|PB]}
	// Required: true
	Size *int64 `json:"size" yaml:"size"`

	// storage service
	StorageService *OracleOnNfsDbStorageService `json:"storage_service,omitempty" yaml:"storage_service,omitempty"`
}

// Validate validates this oracle on nfs db
func (m *OracleOnNfsDb) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleOnNfsDb) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("db"+"."+"size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *OracleOnNfsDb) validateStorageService(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageService) { // not required
		return nil
	}

	if m.StorageService != nil {
		if err := m.StorageService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("db" + "." + "storage_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("db" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this oracle on nfs db based on the context it is used
func (m *OracleOnNfsDb) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleOnNfsDb) contextValidateStorageService(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageService != nil {

		if swag.IsZero(m.StorageService) { // not required
			return nil
		}

		if err := m.StorageService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("db" + "." + "storage_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("db" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OracleOnNfsDb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleOnNfsDb) UnmarshalBinary(b []byte) error {
	var res OracleOnNfsDb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OracleOnNfsDbStorageService oracle on nfs db storage service
//
// swagger:model OracleOnNfsDbStorageService
type OracleOnNfsDbStorageService struct {

	// The storage service of the database.
	// Enum: ["extreme","performance","value"]
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this oracle on nfs db storage service
func (m *OracleOnNfsDbStorageService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var oracleOnNfsDbStorageServiceTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extreme","performance","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oracleOnNfsDbStorageServiceTypeNamePropEnum = append(oracleOnNfsDbStorageServiceTypeNamePropEnum, v)
	}
}

const (

	// OracleOnNfsDbStorageServiceNameExtreme captures enum value "extreme"
	OracleOnNfsDbStorageServiceNameExtreme string = "extreme"

	// OracleOnNfsDbStorageServiceNamePerformance captures enum value "performance"
	OracleOnNfsDbStorageServiceNamePerformance string = "performance"

	// OracleOnNfsDbStorageServiceNameValue captures enum value "value"
	OracleOnNfsDbStorageServiceNameValue string = "value"
)

// prop value enum
func (m *OracleOnNfsDbStorageService) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oracleOnNfsDbStorageServiceTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OracleOnNfsDbStorageService) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("db"+"."+"storage_service"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this oracle on nfs db storage service based on context it is used
func (m *OracleOnNfsDbStorageService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OracleOnNfsDbStorageService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleOnNfsDbStorageService) UnmarshalBinary(b []byte) error {
	var res OracleOnNfsDbStorageService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OracleOnNfsOraHome oracle on nfs ora home
//
// swagger:model OracleOnNfsOraHome
type OracleOnNfsOraHome struct {

	// The size of the ORACLE_HOME storage volume. Usage: {&lt;integer&gt;[KB|MB|GB|TB|PB]}
	Size int64 `json:"size,omitempty" yaml:"size,omitempty"`

	// storage service
	StorageService *OracleOnNfsOraHomeStorageService `json:"storage_service,omitempty" yaml:"storage_service,omitempty"`
}

// Validate validates this oracle on nfs ora home
func (m *OracleOnNfsOraHome) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStorageService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleOnNfsOraHome) validateStorageService(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageService) { // not required
		return nil
	}

	if m.StorageService != nil {
		if err := m.StorageService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ora_home" + "." + "storage_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ora_home" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this oracle on nfs ora home based on the context it is used
func (m *OracleOnNfsOraHome) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleOnNfsOraHome) contextValidateStorageService(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageService != nil {

		if swag.IsZero(m.StorageService) { // not required
			return nil
		}

		if err := m.StorageService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ora_home" + "." + "storage_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ora_home" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OracleOnNfsOraHome) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleOnNfsOraHome) UnmarshalBinary(b []byte) error {
	var res OracleOnNfsOraHome
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OracleOnNfsOraHomeStorageService oracle on nfs ora home storage service
//
// swagger:model OracleOnNfsOraHomeStorageService
type OracleOnNfsOraHomeStorageService struct {

	// The storage service of the ORACLE_HOME storage volume.
	// Enum: ["extreme","performance","value"]
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this oracle on nfs ora home storage service
func (m *OracleOnNfsOraHomeStorageService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var oracleOnNfsOraHomeStorageServiceTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extreme","performance","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oracleOnNfsOraHomeStorageServiceTypeNamePropEnum = append(oracleOnNfsOraHomeStorageServiceTypeNamePropEnum, v)
	}
}

const (

	// OracleOnNfsOraHomeStorageServiceNameExtreme captures enum value "extreme"
	OracleOnNfsOraHomeStorageServiceNameExtreme string = "extreme"

	// OracleOnNfsOraHomeStorageServiceNamePerformance captures enum value "performance"
	OracleOnNfsOraHomeStorageServiceNamePerformance string = "performance"

	// OracleOnNfsOraHomeStorageServiceNameValue captures enum value "value"
	OracleOnNfsOraHomeStorageServiceNameValue string = "value"
)

// prop value enum
func (m *OracleOnNfsOraHomeStorageService) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oracleOnNfsOraHomeStorageServiceTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OracleOnNfsOraHomeStorageService) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("ora_home"+"."+"storage_service"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this oracle on nfs ora home storage service based on context it is used
func (m *OracleOnNfsOraHomeStorageService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OracleOnNfsOraHomeStorageService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleOnNfsOraHomeStorageService) UnmarshalBinary(b []byte) error {
	var res OracleOnNfsOraHomeStorageService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OracleOnNfsProtectionType oracle on nfs protection type
//
// swagger:model OracleOnNfsProtectionType
type OracleOnNfsProtectionType struct {

	// The local RPO of the application.
	// Enum: ["hourly","none"]
	LocalRpo string `json:"local_rpo,omitempty" yaml:"local_rpo,omitempty"`

	// The remote RPO of the application.
	// Enum: ["none","zero"]
	RemoteRpo string `json:"remote_rpo,omitempty" yaml:"remote_rpo,omitempty"`
}

// Validate validates this oracle on nfs protection type
func (m *OracleOnNfsProtectionType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocalRpo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteRpo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var oracleOnNfsProtectionTypeTypeLocalRpoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hourly","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oracleOnNfsProtectionTypeTypeLocalRpoPropEnum = append(oracleOnNfsProtectionTypeTypeLocalRpoPropEnum, v)
	}
}

const (

	// OracleOnNfsProtectionTypeLocalRpoHourly captures enum value "hourly"
	OracleOnNfsProtectionTypeLocalRpoHourly string = "hourly"

	// OracleOnNfsProtectionTypeLocalRpoNone captures enum value "none"
	OracleOnNfsProtectionTypeLocalRpoNone string = "none"
)

// prop value enum
func (m *OracleOnNfsProtectionType) validateLocalRpoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oracleOnNfsProtectionTypeTypeLocalRpoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OracleOnNfsProtectionType) validateLocalRpo(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalRpo) { // not required
		return nil
	}

	// value enum
	if err := m.validateLocalRpoEnum("protection_type"+"."+"local_rpo", "body", m.LocalRpo); err != nil {
		return err
	}

	return nil
}

var oracleOnNfsProtectionTypeTypeRemoteRpoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","zero"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oracleOnNfsProtectionTypeTypeRemoteRpoPropEnum = append(oracleOnNfsProtectionTypeTypeRemoteRpoPropEnum, v)
	}
}

const (

	// OracleOnNfsProtectionTypeRemoteRpoNone captures enum value "none"
	OracleOnNfsProtectionTypeRemoteRpoNone string = "none"

	// OracleOnNfsProtectionTypeRemoteRpoZero captures enum value "zero"
	OracleOnNfsProtectionTypeRemoteRpoZero string = "zero"
)

// prop value enum
func (m *OracleOnNfsProtectionType) validateRemoteRpoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oracleOnNfsProtectionTypeTypeRemoteRpoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OracleOnNfsProtectionType) validateRemoteRpo(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteRpo) { // not required
		return nil
	}

	// value enum
	if err := m.validateRemoteRpoEnum("protection_type"+"."+"remote_rpo", "body", m.RemoteRpo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this oracle on nfs protection type based on context it is used
func (m *OracleOnNfsProtectionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OracleOnNfsProtectionType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleOnNfsProtectionType) UnmarshalBinary(b []byte) error {
	var res OracleOnNfsProtectionType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OracleOnNfsRedoLog oracle on nfs redo log
//
// swagger:model OracleOnNfsRedoLog
type OracleOnNfsRedoLog struct {

	// Specifies whether the redo log group should be mirrored.
	// Enum: [false,true]
	Mirrored *bool `json:"mirrored,omitempty" yaml:"mirrored,omitempty"`

	// The size of the redo log group. Usage: {&lt;integer&gt;[KB|MB|GB|TB|PB]}
	// Required: true
	Size *int64 `json:"size" yaml:"size"`

	// storage service
	StorageService *OracleOnNfsRedoLogStorageService `json:"storage_service,omitempty" yaml:"storage_service,omitempty"`
}

// Validate validates this oracle on nfs redo log
func (m *OracleOnNfsRedoLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMirrored(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var oracleOnNfsRedoLogTypeMirroredPropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[false,true]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oracleOnNfsRedoLogTypeMirroredPropEnum = append(oracleOnNfsRedoLogTypeMirroredPropEnum, v)
	}
}

// prop value enum
func (m *OracleOnNfsRedoLog) validateMirroredEnum(path, location string, value bool) error {
	if err := validate.EnumCase(path, location, value, oracleOnNfsRedoLogTypeMirroredPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OracleOnNfsRedoLog) validateMirrored(formats strfmt.Registry) error {
	if swag.IsZero(m.Mirrored) { // not required
		return nil
	}

	// value enum
	if err := m.validateMirroredEnum("redo_log"+"."+"mirrored", "body", *m.Mirrored); err != nil {
		return err
	}

	return nil
}

func (m *OracleOnNfsRedoLog) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("redo_log"+"."+"size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *OracleOnNfsRedoLog) validateStorageService(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageService) { // not required
		return nil
	}

	if m.StorageService != nil {
		if err := m.StorageService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redo_log" + "." + "storage_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redo_log" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this oracle on nfs redo log based on the context it is used
func (m *OracleOnNfsRedoLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleOnNfsRedoLog) contextValidateStorageService(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageService != nil {

		if swag.IsZero(m.StorageService) { // not required
			return nil
		}

		if err := m.StorageService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redo_log" + "." + "storage_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redo_log" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OracleOnNfsRedoLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleOnNfsRedoLog) UnmarshalBinary(b []byte) error {
	var res OracleOnNfsRedoLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OracleOnNfsRedoLogStorageService oracle on nfs redo log storage service
//
// swagger:model OracleOnNfsRedoLogStorageService
type OracleOnNfsRedoLogStorageService struct {

	// The storage service of the redo log group.
	// Enum: ["extreme","performance","value"]
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this oracle on nfs redo log storage service
func (m *OracleOnNfsRedoLogStorageService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var oracleOnNfsRedoLogStorageServiceTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extreme","performance","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oracleOnNfsRedoLogStorageServiceTypeNamePropEnum = append(oracleOnNfsRedoLogStorageServiceTypeNamePropEnum, v)
	}
}

const (

	// OracleOnNfsRedoLogStorageServiceNameExtreme captures enum value "extreme"
	OracleOnNfsRedoLogStorageServiceNameExtreme string = "extreme"

	// OracleOnNfsRedoLogStorageServiceNamePerformance captures enum value "performance"
	OracleOnNfsRedoLogStorageServiceNamePerformance string = "performance"

	// OracleOnNfsRedoLogStorageServiceNameValue captures enum value "value"
	OracleOnNfsRedoLogStorageServiceNameValue string = "value"
)

// prop value enum
func (m *OracleOnNfsRedoLogStorageService) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oracleOnNfsRedoLogStorageServiceTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OracleOnNfsRedoLogStorageService) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("redo_log"+"."+"storage_service"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this oracle on nfs redo log storage service based on context it is used
func (m *OracleOnNfsRedoLogStorageService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OracleOnNfsRedoLogStorageService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleOnNfsRedoLogStorageService) UnmarshalBinary(b []byte) error {
	var res OracleOnNfsRedoLogStorageService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
