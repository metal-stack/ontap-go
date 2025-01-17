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

// FileDirectorySecurity Manages New Technology File System (NTFS) security and NTFS audit policies.
//
// swagger:model file_directory_security
type FileDirectorySecurity struct {

	// An Access Control Level specifies the access control of the task to be applied. Valid values
	// are "file-directory" or "Storage-Level Access Guard (SLAG)". SLAG is used to apply the
	// specified security descriptors with the task for the volume or qtree. Otherwise, the
	// security descriptors are applied on files and directories at the specified path.
	// The value SLAG is not supported on FlexGroups volumes. The default value is "file-directory"
	// ('-' and '_' are interchangeable).
	//
	// Example: file_directory
	// Enum: ["file_directory","slag"]
	AccessControl *string `json:"access_control,omitempty" yaml:"access_control,omitempty"`

	// Specifies the control flags in the SD. It is a Hexadecimal Value.
	//
	// Example: 8014
	ControlFlags *string `json:"control_flags,omitempty" yaml:"control_flags,omitempty"`

	// Specifies the file attributes on this file or directory.
	//
	// Example: 10
	// Read Only: true
	DosAttributes *string `json:"dos_attributes,omitempty" yaml:"dos_attributes,omitempty"`

	// Specifies the effective style of the SD. The following values are supported:
	// * unix - UNIX style
	// * ntfs - NTFS style
	// * mixed - Mixed style
	// * unified - Unified style
	//
	// Example: mixed
	// Read Only: true
	// Enum: ["unix","ntfs","mixed","unified"]
	EffectiveStyle *string `json:"effective_style,omitempty" yaml:"effective_style,omitempty"`

	// A discretionary access security list (DACL) identifies the trustees that are allowed or denied access
	// to a securable object. When a process tries to access a securable
	// object, the system checks the access control entries (ACEs) in the
	// object's DACL to determine whether to grant access to it.
	//
	FileDirectorySecurityInlineAcls []*ACL `json:"acls,omitempty" yaml:"acls,omitempty"`

	// Specifies that permissions on this file or directory cannot be replaced.
	//
	// Example: ["/dir1/dir2/","/parent/dir3"]
	FileDirectorySecurityInlineIgnorePaths []*string `json:"ignore_paths,omitempty" yaml:"ignore_paths,omitempty"`

	// Specifies the owner's primary group.
	// You can specify the owner group using either a group name or SID.
	//
	// Example: S-1-5-21-2233347455-2266964949-1780268902-69700
	Group *string `json:"group,omitempty" yaml:"group,omitempty"`

	// Specifies group ID on this file or directory.
	//
	// Example: 2
	// Read Only: true
	GroupID *string `json:"group_id,omitempty" yaml:"group_id,omitempty"`

	// Specifies the File Inode number.
	//
	// Example: 64
	// Read Only: true
	Inode *int64 `json:"inode,omitempty" yaml:"inode,omitempty"`

	// Specifies the mode bits on this file or directory.
	//
	// Example: 777
	// Read Only: true
	ModeBits *int64 `json:"mode_bits,omitempty" yaml:"mode_bits,omitempty"`

	// Specifies the owner of the SD.
	// You can specify the owner using either a user name or security identifier (SID).
	// The owner of the SD can modify the permissions on the
	// file (or folder) or files (or folders) to which the SD
	// is applied and can give other users the right to take ownership
	// of the object or objects to which the SD is applied.
	//
	// Example: S-1-5-21-2233347455-2266964949-1780268902-69304
	Owner *string `json:"owner,omitempty" yaml:"owner,omitempty"`

	// Specifies how to propagate security settings to child subfolders and files.
	// This setting determines how child files/folders contained within a parent
	// folder inherit access control and audit information from the parent folder.
	// The available values are:
	// * propagate    - propagate inheritable permissions to all subfolders and files
	// * ignore       - ignore inheritable permissions
	// * replace      - replace existing permissions on all subfolders and files with inheritable permissions
	//
	// Example: propagate
	// Enum: ["propagate","ignore","replace"]
	PropagationMode *string `json:"propagation_mode,omitempty" yaml:"propagation_mode,omitempty"`

	// Specifies the security style of the SD. The following values are supported:
	// * unix - UNIX style
	// * ntfs - NTFS style
	// * mixed - Mixed style
	// * unified - Unified style
	//
	// Example: ntfs
	// Read Only: true
	// Enum: ["unix","ntfs","mixed","unified"]
	SecurityStyle *string `json:"security_style,omitempty" yaml:"security_style,omitempty"`

	// Specifies the textual format of file attributes on this file or directory.
	//
	// Example: ---A----
	// Read Only: true
	TextDosAttr *string `json:"text_dos_attr,omitempty" yaml:"text_dos_attr,omitempty"`

	// Specifies the textual format of mode bits on this file or directory.
	//
	// Example: rwxrwxrwx
	// Read Only: true
	TextModeBits *string `json:"text_mode_bits,omitempty" yaml:"text_mode_bits,omitempty"`

	// Specifies user ID of this file or directory.
	//
	// Example: 10
	// Read Only: true
	UserID *string `json:"user_id,omitempty" yaml:"user_id,omitempty"`
}

// Validate validates this file directory security
func (m *FileDirectorySecurity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessControl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveStyle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileDirectorySecurityInlineAcls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropagationMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityStyle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var fileDirectorySecurityTypeAccessControlPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["file_directory","slag"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fileDirectorySecurityTypeAccessControlPropEnum = append(fileDirectorySecurityTypeAccessControlPropEnum, v)
	}
}

const (

	// FileDirectorySecurityAccessControlFileDirectory captures enum value "file_directory"
	FileDirectorySecurityAccessControlFileDirectory string = "file_directory"

	// FileDirectorySecurityAccessControlSlag captures enum value "slag"
	FileDirectorySecurityAccessControlSlag string = "slag"
)

// prop value enum
func (m *FileDirectorySecurity) validateAccessControlEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fileDirectorySecurityTypeAccessControlPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FileDirectorySecurity) validateAccessControl(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessControl) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessControlEnum("access_control", "body", *m.AccessControl); err != nil {
		return err
	}

	return nil
}

var fileDirectorySecurityTypeEffectiveStylePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unix","ntfs","mixed","unified"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fileDirectorySecurityTypeEffectiveStylePropEnum = append(fileDirectorySecurityTypeEffectiveStylePropEnum, v)
	}
}

const (

	// FileDirectorySecurityEffectiveStyleUnix captures enum value "unix"
	FileDirectorySecurityEffectiveStyleUnix string = "unix"

	// FileDirectorySecurityEffectiveStyleNtfs captures enum value "ntfs"
	FileDirectorySecurityEffectiveStyleNtfs string = "ntfs"

	// FileDirectorySecurityEffectiveStyleMixed captures enum value "mixed"
	FileDirectorySecurityEffectiveStyleMixed string = "mixed"

	// FileDirectorySecurityEffectiveStyleUnified captures enum value "unified"
	FileDirectorySecurityEffectiveStyleUnified string = "unified"
)

// prop value enum
func (m *FileDirectorySecurity) validateEffectiveStyleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fileDirectorySecurityTypeEffectiveStylePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FileDirectorySecurity) validateEffectiveStyle(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectiveStyle) { // not required
		return nil
	}

	// value enum
	if err := m.validateEffectiveStyleEnum("effective_style", "body", *m.EffectiveStyle); err != nil {
		return err
	}

	return nil
}

func (m *FileDirectorySecurity) validateFileDirectorySecurityInlineAcls(formats strfmt.Registry) error {
	if swag.IsZero(m.FileDirectorySecurityInlineAcls) { // not required
		return nil
	}

	for i := 0; i < len(m.FileDirectorySecurityInlineAcls); i++ {
		if swag.IsZero(m.FileDirectorySecurityInlineAcls[i]) { // not required
			continue
		}

		if m.FileDirectorySecurityInlineAcls[i] != nil {
			if err := m.FileDirectorySecurityInlineAcls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("acls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var fileDirectorySecurityTypePropagationModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["propagate","ignore","replace"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fileDirectorySecurityTypePropagationModePropEnum = append(fileDirectorySecurityTypePropagationModePropEnum, v)
	}
}

const (

	// FileDirectorySecurityPropagationModePropagate captures enum value "propagate"
	FileDirectorySecurityPropagationModePropagate string = "propagate"

	// FileDirectorySecurityPropagationModeIgnore captures enum value "ignore"
	FileDirectorySecurityPropagationModeIgnore string = "ignore"

	// FileDirectorySecurityPropagationModeReplace captures enum value "replace"
	FileDirectorySecurityPropagationModeReplace string = "replace"
)

// prop value enum
func (m *FileDirectorySecurity) validatePropagationModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fileDirectorySecurityTypePropagationModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FileDirectorySecurity) validatePropagationMode(formats strfmt.Registry) error {
	if swag.IsZero(m.PropagationMode) { // not required
		return nil
	}

	// value enum
	if err := m.validatePropagationModeEnum("propagation_mode", "body", *m.PropagationMode); err != nil {
		return err
	}

	return nil
}

var fileDirectorySecurityTypeSecurityStylePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unix","ntfs","mixed","unified"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fileDirectorySecurityTypeSecurityStylePropEnum = append(fileDirectorySecurityTypeSecurityStylePropEnum, v)
	}
}

const (

	// FileDirectorySecuritySecurityStyleUnix captures enum value "unix"
	FileDirectorySecuritySecurityStyleUnix string = "unix"

	// FileDirectorySecuritySecurityStyleNtfs captures enum value "ntfs"
	FileDirectorySecuritySecurityStyleNtfs string = "ntfs"

	// FileDirectorySecuritySecurityStyleMixed captures enum value "mixed"
	FileDirectorySecuritySecurityStyleMixed string = "mixed"

	// FileDirectorySecuritySecurityStyleUnified captures enum value "unified"
	FileDirectorySecuritySecurityStyleUnified string = "unified"
)

// prop value enum
func (m *FileDirectorySecurity) validateSecurityStyleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fileDirectorySecurityTypeSecurityStylePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FileDirectorySecurity) validateSecurityStyle(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityStyle) { // not required
		return nil
	}

	// value enum
	if err := m.validateSecurityStyleEnum("security_style", "body", *m.SecurityStyle); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this file directory security based on the context it is used
func (m *FileDirectorySecurity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDosAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEffectiveStyle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileDirectorySecurityInlineAcls(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroupID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModeBits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityStyle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTextDosAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTextModeBits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileDirectorySecurity) contextValidateDosAttributes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dos_attributes", "body", m.DosAttributes); err != nil {
		return err
	}

	return nil
}

func (m *FileDirectorySecurity) contextValidateEffectiveStyle(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "effective_style", "body", m.EffectiveStyle); err != nil {
		return err
	}

	return nil
}

func (m *FileDirectorySecurity) contextValidateFileDirectorySecurityInlineAcls(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FileDirectorySecurityInlineAcls); i++ {

		if m.FileDirectorySecurityInlineAcls[i] != nil {

			if swag.IsZero(m.FileDirectorySecurityInlineAcls[i]) { // not required
				return nil
			}

			if err := m.FileDirectorySecurityInlineAcls[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("acls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FileDirectorySecurity) contextValidateGroupID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "group_id", "body", m.GroupID); err != nil {
		return err
	}

	return nil
}

func (m *FileDirectorySecurity) contextValidateInode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "inode", "body", m.Inode); err != nil {
		return err
	}

	return nil
}

func (m *FileDirectorySecurity) contextValidateModeBits(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "mode_bits", "body", m.ModeBits); err != nil {
		return err
	}

	return nil
}

func (m *FileDirectorySecurity) contextValidateSecurityStyle(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "security_style", "body", m.SecurityStyle); err != nil {
		return err
	}

	return nil
}

func (m *FileDirectorySecurity) contextValidateTextDosAttr(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "text_dos_attr", "body", m.TextDosAttr); err != nil {
		return err
	}

	return nil
}

func (m *FileDirectorySecurity) contextValidateTextModeBits(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "text_mode_bits", "body", m.TextModeBits); err != nil {
		return err
	}

	return nil
}

func (m *FileDirectorySecurity) contextValidateUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileDirectorySecurity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileDirectorySecurity) UnmarshalBinary(b []byte) error {
	var res FileDirectorySecurity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}