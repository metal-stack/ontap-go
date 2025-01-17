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

// FpolicyEvent The information that a FPolicy process needs to determine what file access operations to monitor and for which of the monitored events notifications should be sent to the external FPolicy server.
//
// swagger:model fpolicy_event
type FpolicyEvent struct {

	// file operations
	FileOperations *FpolicyEventInlineFileOperations `json:"file_operations,omitempty" yaml:"file_operations,omitempty"`

	// filters
	Filters *FpolicyEventInlineFilters `json:"filters,omitempty" yaml:"filters,omitempty"`

	// Specifies whether failed file operations monitoring is required.
	MonitorFileopFailure *bool `json:"monitor_fileop_failure,omitempty" yaml:"monitor_fileop_failure,omitempty"`

	// Specifies the name of the FPolicy event.
	// Example: event_cifs
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// Protocol for which event is created. If you specify protocol, then you
	// must also specify a valid value for the file operation parameters.
	//   The value of this parameter must be one of the following:
	//     * cifs  - for the CIFS protocol.
	//     * nfsv3 - for the NFSv3 protocol.
	//     * nfsv4 - for the NFSv4 protocol.
	//
	// Enum: ["cifs","nfsv3","nfsv4"]
	Protocol *string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// svm
	Svm *FpolicyEventInlineSvm `json:"svm,omitempty" yaml:"svm,omitempty"`

	// Specifies whether volume operation monitoring is required.
	VolumeMonitoring *bool `json:"volume_monitoring,omitempty" yaml:"volume_monitoring,omitempty"`
}

// Validate validates this fpolicy event
func (m *FpolicyEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileOperations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
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

func (m *FpolicyEvent) validateFileOperations(formats strfmt.Registry) error {
	if swag.IsZero(m.FileOperations) { // not required
		return nil
	}

	if m.FileOperations != nil {
		if err := m.FileOperations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file_operations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("file_operations")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyEvent) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	if m.Filters != nil {
		if err := m.Filters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filters")
			}
			return err
		}
	}

	return nil
}

var fpolicyEventTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cifs","nfsv3","nfsv4"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fpolicyEventTypeProtocolPropEnum = append(fpolicyEventTypeProtocolPropEnum, v)
	}
}

const (

	// FpolicyEventProtocolCifs captures enum value "cifs"
	FpolicyEventProtocolCifs string = "cifs"

	// FpolicyEventProtocolNfsv3 captures enum value "nfsv3"
	FpolicyEventProtocolNfsv3 string = "nfsv3"

	// FpolicyEventProtocolNfsv4 captures enum value "nfsv4"
	FpolicyEventProtocolNfsv4 string = "nfsv4"
)

// prop value enum
func (m *FpolicyEvent) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fpolicyEventTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FpolicyEvent) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *FpolicyEvent) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this fpolicy event based on the context it is used
func (m *FpolicyEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFileOperations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilters(ctx, formats); err != nil {
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

func (m *FpolicyEvent) contextValidateFileOperations(ctx context.Context, formats strfmt.Registry) error {

	if m.FileOperations != nil {

		if swag.IsZero(m.FileOperations) { // not required
			return nil
		}

		if err := m.FileOperations.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file_operations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("file_operations")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyEvent) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	if m.Filters != nil {

		if swag.IsZero(m.Filters) { // not required
			return nil
		}

		if err := m.Filters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filters")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyEvent) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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
func (m *FpolicyEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyEvent) UnmarshalBinary(b []byte) error {
	var res FpolicyEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyEventInlineFileOperations Specifies the file operations for the FPolicy event. You must specify a valid protocol in the protocol parameter.
// The event will check the operations specified from all client requests using the protocol.
//
// swagger:model fpolicy_event_inline_file_operations
type FpolicyEventInlineFileOperations struct {

	// Access operations
	Access *bool `json:"access,omitempty" yaml:"access,omitempty"`

	// File close operations
	Close *bool `json:"close,omitempty" yaml:"close,omitempty"`

	// File create operations
	Create *bool `json:"create,omitempty" yaml:"create,omitempty"`

	// Directory create operations
	CreateDir *bool `json:"create_dir,omitempty" yaml:"create_dir,omitempty"`

	// File delete operations
	Delete *bool `json:"delete,omitempty" yaml:"delete,omitempty"`

	// Directory delete operations
	DeleteDir *bool `json:"delete_dir,omitempty" yaml:"delete_dir,omitempty"`

	// Get attribute operations
	Getattr *bool `json:"getattr,omitempty" yaml:"getattr,omitempty"`

	// Link operations
	Link *bool `json:"link,omitempty" yaml:"link,omitempty"`

	// Lookup operations
	Lookup *bool `json:"lookup,omitempty" yaml:"lookup,omitempty"`

	// File open operations
	Open *bool `json:"open,omitempty" yaml:"open,omitempty"`

	// File read operations
	Read *bool `json:"read,omitempty" yaml:"read,omitempty"`

	// File rename operations
	Rename *bool `json:"rename,omitempty" yaml:"rename,omitempty"`

	// Directory rename operations
	RenameDir *bool `json:"rename_dir,omitempty" yaml:"rename_dir,omitempty"`

	// Set attribute operations
	Setattr *bool `json:"setattr,omitempty" yaml:"setattr,omitempty"`

	// Symbolic link operations
	Symlink *bool `json:"symlink,omitempty" yaml:"symlink,omitempty"`

	// File write operations
	Write *bool `json:"write,omitempty" yaml:"write,omitempty"`
}

// Validate validates this fpolicy event inline file operations
func (m *FpolicyEventInlineFileOperations) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fpolicy event inline file operations based on context it is used
func (m *FpolicyEventInlineFileOperations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyEventInlineFileOperations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyEventInlineFileOperations) UnmarshalBinary(b []byte) error {
	var res FpolicyEventInlineFileOperations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyEventInlineFilters Specifies the list of filters for a given file operation for the specified protocol.
// When you specify the filters, you must specify the valid protocols and a valid file operations.
//
// swagger:model fpolicy_event_inline_filters
type FpolicyEventInlineFilters struct {

	// Filter the client request for close with modification.
	CloseWithModification *bool `json:"close_with_modification,omitempty" yaml:"close_with_modification,omitempty"`

	// Filter the client request for close with read.
	CloseWithRead *bool `json:"close_with_read,omitempty" yaml:"close_with_read,omitempty"`

	// Filter the client request for close without modification.
	CloseWithoutModification *bool `json:"close_without_modification,omitempty" yaml:"close_without_modification,omitempty"`

	// Filter the client requests for directory operations. When this filter is specified directory operations are not monitored.
	ExcludeDirectory *bool `json:"exclude_directory,omitempty" yaml:"exclude_directory,omitempty"`

	// Filter the client requests for the first-read.
	FirstRead *bool `json:"first_read,omitempty" yaml:"first_read,omitempty"`

	// Filter the client requests for the first-write.
	FirstWrite *bool `json:"first_write,omitempty" yaml:"first_write,omitempty"`

	// Filter the client request for alternate data stream.
	MonitorAds *bool `json:"monitor_ads,omitempty" yaml:"monitor_ads,omitempty"`

	// Filter the client request for offline bit set. FPolicy server receives notification only when offline files are accessed.
	OfflineBit *bool `json:"offline_bit,omitempty" yaml:"offline_bit,omitempty"`

	// Filter the client request for open with delete intent.
	OpenWithDeleteIntent *bool `json:"open_with_delete_intent,omitempty" yaml:"open_with_delete_intent,omitempty"`

	// Filter the client request for open with write intent.
	OpenWithWriteIntent *bool `json:"open_with_write_intent,omitempty" yaml:"open_with_write_intent,omitempty"`

	// Filter the client setattr requests for changing the access time of a file or directory.
	SetattrWithAccessTimeChange *bool `json:"setattr_with_access_time_change,omitempty" yaml:"setattr_with_access_time_change,omitempty"`

	// Filter the client setattr requests for changing the allocation size of a file.
	SetattrWithAllocationSizeChange *bool `json:"setattr_with_allocation_size_change,omitempty" yaml:"setattr_with_allocation_size_change,omitempty"`

	// Filter the client setattr requests for changing the creation time of a file or directory.
	SetattrWithCreationTimeChange *bool `json:"setattr_with_creation_time_change,omitempty" yaml:"setattr_with_creation_time_change,omitempty"`

	// Filter the client setattr requests for changing dacl on a file or directory.
	SetattrWithDaclChange *bool `json:"setattr_with_dacl_change,omitempty" yaml:"setattr_with_dacl_change,omitempty"`

	// Filter the client setattr requests for changing group of a file or directory.
	SetattrWithGroupChange *bool `json:"setattr_with_group_change,omitempty" yaml:"setattr_with_group_change,omitempty"`

	// Filter the client setattr requests for changing the mode bits on a file or directory.
	SetattrWithModeChange *bool `json:"setattr_with_mode_change,omitempty" yaml:"setattr_with_mode_change,omitempty"`

	// Filter the client setattr requests for changing the modification time of a file or directory.
	SetattrWithModifyTimeChange *bool `json:"setattr_with_modify_time_change,omitempty" yaml:"setattr_with_modify_time_change,omitempty"`

	// Filter the client setattr requests for changing owner of a file or directory.
	SetattrWithOwnerChange *bool `json:"setattr_with_owner_change,omitempty" yaml:"setattr_with_owner_change,omitempty"`

	// Filter the client setattr requests for changing sacl on a file or directory.
	SetattrWithSaclChange *bool `json:"setattr_with_sacl_change,omitempty" yaml:"setattr_with_sacl_change,omitempty"`

	// Filter the client setattr requests for changing the size of a file.
	SetattrWithSizeChange *bool `json:"setattr_with_size_change,omitempty" yaml:"setattr_with_size_change,omitempty"`

	// Filter the client request for write with size change.
	WriteWithSizeChange *bool `json:"write_with_size_change,omitempty" yaml:"write_with_size_change,omitempty"`
}

// Validate validates this fpolicy event inline filters
func (m *FpolicyEventInlineFilters) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fpolicy event inline filters based on context it is used
func (m *FpolicyEventInlineFilters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyEventInlineFilters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyEventInlineFilters) UnmarshalBinary(b []byte) error {
	var res FpolicyEventInlineFilters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyEventInlineSvm fpolicy event inline svm
//
// swagger:model fpolicy_event_inline_svm
type FpolicyEventInlineSvm struct {

	// SVM UUID
	// Read Only: true
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this fpolicy event inline svm
func (m *FpolicyEventInlineSvm) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this fpolicy event inline svm based on the context it is used
func (m *FpolicyEventInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyEventInlineSvm) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "svm"+"."+"uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyEventInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyEventInlineSvm) UnmarshalBinary(b []byte) error {
	var res FpolicyEventInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}