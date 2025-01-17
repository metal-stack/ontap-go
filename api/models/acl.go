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

// ACL An ACE is an element in an access control list (ACL).
// An ACL can have zero or more ACEs. Each ACE controls or monitors access to
// an object by a specified trustee.
//
// swagger:model acl
type ACL struct {

	// Specifies whether the ACL is for DACL or SACL.
	// The available values are:
	// * access_allow                     - DACL for allow access
	// * access_deny                      - DACL for deny access
	// * access_allowed_callback          - CALLBACK for allowed access
	// * access_denied_callback           - CALLBACK for denied access
	// * access_allowed_callback_object   - CALLBACK OBJECT for allowed access
	// * access_denied_callback_object    - CALLBACK OBJECT for denied access
	// * system_audit_callback            - SYSTEM Audit Callback ace
	// * system_audit_callback_object     - SYSTEM Audit Callback Object ace
	// * system_resource_attribute        - SYSTEM Resource Attribute
	// * system_scoped_policy_id          - SYSTEM Scope Policy ID
	// * audit_success                    - SACL for success access
	// * audit_failure                    - SACL for failure access
	// * audit_success_and_failure        - SACL for both success and failure access
	//
	// Example: access_allow
	// Enum: ["access_allow","access_deny","access_allowed_callback","access_denied_callback","access_allowed_callback_object","access_denied_callback_object","system_audit_callback","system_audit_callback_object","system_resource_attribute","system_scoped_policy_id","audit_failure","audit_success","audit_success_and_failure"]
	Access *string `json:"access,omitempty" yaml:"access,omitempty"`

	// An Access Control Level specifies the access control of the task to be applied. Valid values
	// are "file-directory" or "Storage-Level Access Guard (SLAG)". SLAG is used to apply the
	// specified security descriptors with the task for the volume or qtree. Otherwise, the security
	// descriptors are applied on files and directories at the specified path. The value SLAG is not
	// supported on FlexGroups volumes. The default value is "file-directory"
	// ('-' and '_' are interchangeable).
	//
	// Example: file_directory
	// Read Only: true
	// Enum: ["file_directory","slag"]
	AccessControl *string `json:"access_control,omitempty" yaml:"access_control,omitempty"`

	// advanced rights
	AdvancedRights *AdvancedRights `json:"advanced_rights,omitempty" yaml:"advanced_rights,omitempty"`

	// apply to
	ApplyTo *ApplyTo `json:"apply_to,omitempty" yaml:"apply_to,omitempty"`

	// Indicates whether or not the ACE flag is inherited.
	//
	// Example: true
	// Read Only: true
	Inherited *bool `json:"inherited,omitempty" yaml:"inherited,omitempty"`

	// rights
	Rights *Rights `json:"rights,omitempty" yaml:"rights,omitempty"`

	// Specifies the account to which the ACE applies.
	// You can specify either name or SID.
	//
	// Example: S-1-5-21-2233347455-2266964949-1780268902-69304
	User *string `json:"user,omitempty" yaml:"user,omitempty"`
}

// Validate validates this acl
func (m *ACL) Validate(formats strfmt.Registry) error {
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

	if err := m.validateRights(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var aclTypeAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["access_allow","access_deny","access_allowed_callback","access_denied_callback","access_allowed_callback_object","access_denied_callback_object","system_audit_callback","system_audit_callback_object","system_resource_attribute","system_scoped_policy_id","audit_failure","audit_success","audit_success_and_failure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aclTypeAccessPropEnum = append(aclTypeAccessPropEnum, v)
	}
}

const (

	// ACLAccessAccessAllow captures enum value "access_allow"
	ACLAccessAccessAllow string = "access_allow"

	// ACLAccessAccessDeny captures enum value "access_deny"
	ACLAccessAccessDeny string = "access_deny"

	// ACLAccessAccessAllowedCallback captures enum value "access_allowed_callback"
	ACLAccessAccessAllowedCallback string = "access_allowed_callback"

	// ACLAccessAccessDeniedCallback captures enum value "access_denied_callback"
	ACLAccessAccessDeniedCallback string = "access_denied_callback"

	// ACLAccessAccessAllowedCallbackObject captures enum value "access_allowed_callback_object"
	ACLAccessAccessAllowedCallbackObject string = "access_allowed_callback_object"

	// ACLAccessAccessDeniedCallbackObject captures enum value "access_denied_callback_object"
	ACLAccessAccessDeniedCallbackObject string = "access_denied_callback_object"

	// ACLAccessSystemAuditCallback captures enum value "system_audit_callback"
	ACLAccessSystemAuditCallback string = "system_audit_callback"

	// ACLAccessSystemAuditCallbackObject captures enum value "system_audit_callback_object"
	ACLAccessSystemAuditCallbackObject string = "system_audit_callback_object"

	// ACLAccessSystemResourceAttribute captures enum value "system_resource_attribute"
	ACLAccessSystemResourceAttribute string = "system_resource_attribute"

	// ACLAccessSystemScopedPolicyID captures enum value "system_scoped_policy_id"
	ACLAccessSystemScopedPolicyID string = "system_scoped_policy_id"

	// ACLAccessAuditFailure captures enum value "audit_failure"
	ACLAccessAuditFailure string = "audit_failure"

	// ACLAccessAuditSuccess captures enum value "audit_success"
	ACLAccessAuditSuccess string = "audit_success"

	// ACLAccessAuditSuccessAndFailure captures enum value "audit_success_and_failure"
	ACLAccessAuditSuccessAndFailure string = "audit_success_and_failure"
)

// prop value enum
func (m *ACL) validateAccessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, aclTypeAccessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ACL) validateAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.Access) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessEnum("access", "body", *m.Access); err != nil {
		return err
	}

	return nil
}

var aclTypeAccessControlPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["file_directory","slag"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aclTypeAccessControlPropEnum = append(aclTypeAccessControlPropEnum, v)
	}
}

const (

	// ACLAccessControlFileDirectory captures enum value "file_directory"
	ACLAccessControlFileDirectory string = "file_directory"

	// ACLAccessControlSlag captures enum value "slag"
	ACLAccessControlSlag string = "slag"
)

// prop value enum
func (m *ACL) validateAccessControlEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, aclTypeAccessControlPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ACL) validateAccessControl(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessControl) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessControlEnum("access_control", "body", *m.AccessControl); err != nil {
		return err
	}

	return nil
}

func (m *ACL) validateAdvancedRights(formats strfmt.Registry) error {
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

func (m *ACL) validateApplyTo(formats strfmt.Registry) error {
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

func (m *ACL) validateRights(formats strfmt.Registry) error {
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

// ContextValidate validate this acl based on the context it is used
func (m *ACL) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessControl(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdvancedRights(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApplyTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInherited(ctx, formats); err != nil {
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

func (m *ACL) contextValidateAccessControl(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "access_control", "body", m.AccessControl); err != nil {
		return err
	}

	return nil
}

func (m *ACL) contextValidateAdvancedRights(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ACL) contextValidateApplyTo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ACL) contextValidateInherited(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "inherited", "body", m.Inherited); err != nil {
		return err
	}

	return nil
}

func (m *ACL) contextValidateRights(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ACL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ACL) UnmarshalBinary(b []byte) error {
	var res ACL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}