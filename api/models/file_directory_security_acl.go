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

// FileDirectorySecurityACL Manages the DACLS or SACLS.
//
// swagger:model file_directory_security_acl
type FileDirectorySecurityACL struct {

	// Specifies whether the ACL is for DACL or SACL. It is a required field.
	// The available values are:
	// * access_allow                     - DACL for allow access
	// * access_deny                      - DACL for deny access
	// * audit_success                    - SACL for success access
	// * audit_failure                    - SACL for failure access
	//
	// Example: access_allow
	// Enum: ["access_allow","access_deny","audit_failure","audit_success"]
	Access *string `json:"access,omitempty" yaml:"access,omitempty"`

	// Access Control Level specifies the access control of the task to be applied. Valid values
	// are "file-directory" or "Storage-Level Access Guard (SLAG)". SLAG is used to apply the
	// specified security descriptors with the task for the volume or qtree. Otherwise, the
	// security descriptors are applied on files and directories at the specified path. The
	// value SLAG is not supported on FlexGroups volumes. The default value is "file-directory"
	// ('-' and '_' are interchangeable).
	//
	// Example: file_directory
	// Enum: ["file_directory","slag"]
	AccessControl *string `json:"access_control,omitempty" yaml:"access_control,omitempty"`

	// advanced rights
	AdvancedRights *AdvancedRights `json:"advanced_rights,omitempty" yaml:"advanced_rights,omitempty"`

	// apply to
	ApplyTo *ApplyTo `json:"apply_to,omitempty" yaml:"apply_to,omitempty"`

	// Specifies that permissions on this file or directory cannot be replaced.
	//
	// Example: ["/dir1/dir2/","/parent/dir3"]
	FileDirectorySecurityACLInlineIgnorePaths []*string `json:"ignore_paths,omitempty" yaml:"ignore_paths,omitempty"`

	// Specifies how to propagate security settings to child subfolders and files.
	// This setting determines how child files/folders contained within a parent
	// folder inherit access control and audit information from the parent folder.
	// The available values are:
	// * propagate    - propagate inheritable permissions to all subfolders and files
	// * ignore       - ignore inheritable permissions
	// * replace      - replace existing permissions on all subfolders and files with inheritable permissions
	//
	// Enum: ["propagate","ignore","replace"]
	PropagationMode *string `json:"propagation_mode,omitempty" yaml:"propagation_mode,omitempty"`

	// rights
	Rights *Rights `json:"rights,omitempty" yaml:"rights,omitempty"`

	// Specifies the account to which the ACE applies.
	// You can specify either name or SID.
	//
	// Example: S-1-5-21-2233347455-2266964949-1780268902-69304
	User *string `json:"user,omitempty" yaml:"user,omitempty"`
}

// Validate validates this file directory security acl
func (m *FileDirectorySecurityACL) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccessControl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdvancedRights(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplyTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropagationMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRights(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var fileDirectorySecurityAclTypeAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["access_allow","access_deny","audit_failure","audit_success"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fileDirectorySecurityAclTypeAccessPropEnum = append(fileDirectorySecurityAclTypeAccessPropEnum, v)
	}
}

const (

	// FileDirectorySecurityACLAccessAccessAllow captures enum value "access_allow"
	FileDirectorySecurityACLAccessAccessAllow string = "access_allow"

	// FileDirectorySecurityACLAccessAccessDeny captures enum value "access_deny"
	FileDirectorySecurityACLAccessAccessDeny string = "access_deny"

	// FileDirectorySecurityACLAccessAuditFailure captures enum value "audit_failure"
	FileDirectorySecurityACLAccessAuditFailure string = "audit_failure"

	// FileDirectorySecurityACLAccessAuditSuccess captures enum value "audit_success"
	FileDirectorySecurityACLAccessAuditSuccess string = "audit_success"
)

// prop value enum
func (m *FileDirectorySecurityACL) validateAccessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fileDirectorySecurityAclTypeAccessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FileDirectorySecurityACL) validateAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.Access) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessEnum("access", "body", *m.Access); err != nil {
		return err
	}

	return nil
}

var fileDirectorySecurityAclTypeAccessControlPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["file_directory","slag"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fileDirectorySecurityAclTypeAccessControlPropEnum = append(fileDirectorySecurityAclTypeAccessControlPropEnum, v)
	}
}

const (

	// FileDirectorySecurityACLAccessControlFileDirectory captures enum value "file_directory"
	FileDirectorySecurityACLAccessControlFileDirectory string = "file_directory"

	// FileDirectorySecurityACLAccessControlSlag captures enum value "slag"
	FileDirectorySecurityACLAccessControlSlag string = "slag"
)

// prop value enum
func (m *FileDirectorySecurityACL) validateAccessControlEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fileDirectorySecurityAclTypeAccessControlPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FileDirectorySecurityACL) validateAccessControl(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessControl) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessControlEnum("access_control", "body", *m.AccessControl); err != nil {
		return err
	}

	return nil
}

func (m *FileDirectorySecurityACL) validateAdvancedRights(formats strfmt.Registry) error {
	if swag.IsZero(m.AdvancedRights) { // not required
		return nil
	}

	if m.AdvancedRights != nil {
		if err := m.AdvancedRights.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advanced_rights")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("advanced_rights")
			}
			return err
		}
	}

	return nil
}

func (m *FileDirectorySecurityACL) validateApplyTo(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplyTo) { // not required
		return nil
	}

	if m.ApplyTo != nil {
		if err := m.ApplyTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apply_to")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apply_to")
			}
			return err
		}
	}

	return nil
}

var fileDirectorySecurityAclTypePropagationModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["propagate","ignore","replace"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fileDirectorySecurityAclTypePropagationModePropEnum = append(fileDirectorySecurityAclTypePropagationModePropEnum, v)
	}
}

const (

	// FileDirectorySecurityACLPropagationModePropagate captures enum value "propagate"
	FileDirectorySecurityACLPropagationModePropagate string = "propagate"

	// FileDirectorySecurityACLPropagationModeIgnore captures enum value "ignore"
	FileDirectorySecurityACLPropagationModeIgnore string = "ignore"

	// FileDirectorySecurityACLPropagationModeReplace captures enum value "replace"
	FileDirectorySecurityACLPropagationModeReplace string = "replace"
)

// prop value enum
func (m *FileDirectorySecurityACL) validatePropagationModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fileDirectorySecurityAclTypePropagationModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FileDirectorySecurityACL) validatePropagationMode(formats strfmt.Registry) error {
	if swag.IsZero(m.PropagationMode) { // not required
		return nil
	}

	// value enum
	if err := m.validatePropagationModeEnum("propagation_mode", "body", *m.PropagationMode); err != nil {
		return err
	}

	return nil
}

func (m *FileDirectorySecurityACL) validateRights(formats strfmt.Registry) error {
	if swag.IsZero(m.Rights) { // not required
		return nil
	}

	if m.Rights != nil {
		if err := m.Rights.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rights")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rights")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this file directory security acl based on the context it is used
func (m *FileDirectorySecurityACL) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdvancedRights(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApplyTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRights(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileDirectorySecurityACL) contextValidateAdvancedRights(ctx context.Context, formats strfmt.Registry) error {

	if m.AdvancedRights != nil {

		if swag.IsZero(m.AdvancedRights) { // not required
			return nil
		}

		if err := m.AdvancedRights.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advanced_rights")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("advanced_rights")
			}
			return err
		}
	}

	return nil
}

func (m *FileDirectorySecurityACL) contextValidateApplyTo(ctx context.Context, formats strfmt.Registry) error {

	if m.ApplyTo != nil {

		if swag.IsZero(m.ApplyTo) { // not required
			return nil
		}

		if err := m.ApplyTo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apply_to")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apply_to")
			}
			return err
		}
	}

	return nil
}

func (m *FileDirectorySecurityACL) contextValidateRights(ctx context.Context, formats strfmt.Registry) error {

	if m.Rights != nil {

		if swag.IsZero(m.Rights) { // not required
			return nil
		}

		if err := m.Rights.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rights")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rights")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileDirectorySecurityACL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileDirectorySecurityACL) UnmarshalBinary(b []byte) error {
	var res FileDirectorySecurityACL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
