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

// SnaplockLog snaplock log
//
// swagger:model snaplock_log
type SnaplockLog struct {

	// links
	Links *SnaplockLogInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// log archive
	LogArchive *SnaplockLogInlineLogArchive `json:"log_archive,omitempty" yaml:"log_archive,omitempty"`

	// log volume
	LogVolume *SnaplockLogVolume `json:"log_volume,omitempty" yaml:"log_volume,omitempty"`

	// snaplock log inline log files
	// Read Only: true
	SnaplockLogInlineLogFiles []*SnaplockLogFile `json:"log_files,omitempty" yaml:"log_files,omitempty"`

	// svm
	Svm *SnaplockLogInlineSvm `json:"svm,omitempty" yaml:"svm,omitempty"`
}

// Validate validates this snaplock log
func (m *SnaplockLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogArchive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnaplockLogInlineLogFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLog) validateLinks(formats strfmt.Registry) error {
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

func (m *SnaplockLog) validateLogArchive(formats strfmt.Registry) error {
	if swag.IsZero(m.LogArchive) { // not required
		return nil
	}

	if m.LogArchive != nil {
		if err := m.LogArchive.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_archive")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_archive")
			}
			return err
		}
	}

	return nil
}

func (m *SnaplockLog) validateLogVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.LogVolume) { // not required
		return nil
	}

	if m.LogVolume != nil {
		if err := m.LogVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_volume")
			}
			return err
		}
	}

	return nil
}

func (m *SnaplockLog) validateSnaplockLogInlineLogFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.SnaplockLogInlineLogFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.SnaplockLogInlineLogFiles); i++ {
		if swag.IsZero(m.SnaplockLogInlineLogFiles[i]) { // not required
			continue
		}

		if m.SnaplockLogInlineLogFiles[i] != nil {
			if err := m.SnaplockLogInlineLogFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("log_files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("log_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnaplockLog) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snaplock log based on the context it is used
func (m *SnaplockLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogArchive(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnaplockLogInlineLogFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLog) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SnaplockLog) contextValidateLogArchive(ctx context.Context, formats strfmt.Registry) error {

	if m.LogArchive != nil {

		if swag.IsZero(m.LogArchive) { // not required
			return nil
		}

		if err := m.LogArchive.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_archive")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_archive")
			}
			return err
		}
	}

	return nil
}

func (m *SnaplockLog) contextValidateLogVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.LogVolume != nil {

		if swag.IsZero(m.LogVolume) { // not required
			return nil
		}

		if err := m.LogVolume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_volume")
			}
			return err
		}
	}

	return nil
}

func (m *SnaplockLog) contextValidateSnaplockLogInlineLogFiles(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "log_files", "body", []*SnaplockLogFile(m.SnaplockLogInlineLogFiles)); err != nil {
		return err
	}

	for i := 0; i < len(m.SnaplockLogInlineLogFiles); i++ {

		if m.SnaplockLogInlineLogFiles[i] != nil {

			if swag.IsZero(m.SnaplockLogInlineLogFiles[i]) { // not required
				return nil
			}

			if err := m.SnaplockLogInlineLogFiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("log_files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("log_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnaplockLog) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {

		if swag.IsZero(m.Svm) { // not required
			return nil
		}

		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnaplockLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLog) UnmarshalBinary(b []byte) error {
	var res SnaplockLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnaplockLogInlineLinks snaplock log inline links
//
// swagger:model snaplock_log_inline__links
type SnaplockLogInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this snaplock log inline links
func (m *SnaplockLogInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this snaplock log inline links based on the context it is used
func (m *SnaplockLogInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SnaplockLogInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLogInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnaplockLogInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnaplockLogInlineLogArchive snaplock log inline log archive
//
// swagger:model snaplock_log_inline_log_archive
type SnaplockLogInlineLogArchive struct {

	// links
	Links *SnaplockLogInlineLogArchiveInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// Archive the specified SnapLock log file for the given base_name, and create a new log file. If base_name is not mentioned, archive all log files.
	Archive *bool `json:"archive,omitempty" yaml:"archive,omitempty"`

	// Base name of log archive
	// Enum: ["legal_hold","privileged_delete","system"]
	BaseName *string `json:"base_name,omitempty" yaml:"base_name,omitempty"`
}

// Validate validates this snaplock log inline log archive
func (m *SnaplockLogInlineLogArchive) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogInlineLogArchive) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_archive" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_archive" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

var snaplockLogInlineLogArchiveTypeBaseNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["legal_hold","privileged_delete","system"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snaplockLogInlineLogArchiveTypeBaseNamePropEnum = append(snaplockLogInlineLogArchiveTypeBaseNamePropEnum, v)
	}
}

const (

	// SnaplockLogInlineLogArchiveBaseNameLegalHold captures enum value "legal_hold"
	SnaplockLogInlineLogArchiveBaseNameLegalHold string = "legal_hold"

	// SnaplockLogInlineLogArchiveBaseNamePrivilegedDelete captures enum value "privileged_delete"
	SnaplockLogInlineLogArchiveBaseNamePrivilegedDelete string = "privileged_delete"

	// SnaplockLogInlineLogArchiveBaseNameSystem captures enum value "system"
	SnaplockLogInlineLogArchiveBaseNameSystem string = "system"
)

// prop value enum
func (m *SnaplockLogInlineLogArchive) validateBaseNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snaplockLogInlineLogArchiveTypeBaseNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnaplockLogInlineLogArchive) validateBaseName(formats strfmt.Registry) error {
	if swag.IsZero(m.BaseName) { // not required
		return nil
	}

	// value enum
	if err := m.validateBaseNameEnum("log_archive"+"."+"base_name", "body", *m.BaseName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this snaplock log inline log archive based on the context it is used
func (m *SnaplockLogInlineLogArchive) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogInlineLogArchive) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_archive" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_archive" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnaplockLogInlineLogArchive) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLogInlineLogArchive) UnmarshalBinary(b []byte) error {
	var res SnaplockLogInlineLogArchive
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnaplockLogInlineLogArchiveInlineLinks snaplock log inline log archive inline links
//
// swagger:model snaplock_log_inline_log_archive_inline__links
type SnaplockLogInlineLogArchiveInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this snaplock log inline log archive inline links
func (m *SnaplockLogInlineLogArchiveInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogInlineLogArchiveInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_archive" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_archive" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snaplock log inline log archive inline links based on the context it is used
func (m *SnaplockLogInlineLogArchiveInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogInlineLogArchiveInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_archive" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_archive" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnaplockLogInlineLogArchiveInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLogInlineLogArchiveInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnaplockLogInlineLogArchiveInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnaplockLogInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model snaplock_log_inline_svm
type SnaplockLogInlineSvm struct {

	// links
	Links *SnaplockLogInlineSvmInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this snaplock log inline svm
func (m *SnaplockLogInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snaplock log inline svm based on the context it is used
func (m *SnaplockLogInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnaplockLogInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLogInlineSvm) UnmarshalBinary(b []byte) error {
	var res SnaplockLogInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnaplockLogInlineSvmInlineLinks snaplock log inline svm inline links
//
// swagger:model snaplock_log_inline_svm_inline__links
type SnaplockLogInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this snaplock log inline svm inline links
func (m *SnaplockLogInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snaplock log inline svm inline links based on the context it is used
func (m *SnaplockLogInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnaplockLogInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnaplockLogInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnaplockLogInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnaplockLogInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
