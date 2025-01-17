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

// ExportRule export rule
//
// swagger:model export_rule
type ExportRule struct {

	// links
	Links *ExportRuleInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// Specifies whether or not device creation is allowed.
	AllowDeviceCreation *bool `json:"allow_device_creation,omitempty" yaml:"allow_device_creation,omitempty"`

	// Specifies whether or not SetUID bits in SETATTR Op is to be honored.
	AllowSuid *bool `json:"allow_suid,omitempty" yaml:"allow_suid,omitempty"`

	// User ID To Which Anonymous Users Are Mapped.
	AnonymousUser *string `json:"anonymous_user,omitempty" yaml:"anonymous_user,omitempty"`

	// Specifies who is authorized to change the ownership mode of a file.
	// Enum: ["restricted","unrestricted"]
	ChownMode *string `json:"chown_mode,omitempty" yaml:"chown_mode,omitempty"`

	// Array of client matches
	ExportRuleInlineClients []*ExportClients `json:"clients,omitempty" yaml:"clients,omitempty"`

	// Authentication flavors that the read-only access rule governs
	//
	ExportRuleInlineRoRule []*ExportAuthenticationFlavor `json:"ro_rule,omitempty" yaml:"ro_rule,omitempty"`

	// Authentication flavors that the read/write access rule governs
	//
	ExportRuleInlineRwRule []*ExportAuthenticationFlavor `json:"rw_rule,omitempty" yaml:"rw_rule,omitempty"`

	// Authentication flavors that the superuser security type governs
	//
	ExportRuleInlineSuperuser []*ExportAuthenticationFlavor `json:"superuser,omitempty" yaml:"superuser,omitempty"`

	// Index of the rule within the export policy.
	//
	Index *int64 `json:"index,omitempty" yaml:"index,omitempty"`

	// NTFS export UNIX security options.
	// Enum: ["fail","ignore"]
	NtfsUnixSecurity *string `json:"ntfs_unix_security,omitempty" yaml:"ntfs_unix_security,omitempty"`

	// policy
	Policy *ExportRuleInlinePolicy `json:"policy,omitempty" yaml:"policy,omitempty"`

	// protocols
	Protocols []*string `json:"protocols,omitempty" yaml:"protocols,omitempty"`

	// svm
	Svm *ExportRuleInlineSvm `json:"svm,omitempty" yaml:"svm,omitempty"`
}

// Validate validates this export rule
func (m *ExportRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChownMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportRuleInlineClients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportRuleInlineRoRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportRuleInlineRwRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportRuleInlineSuperuser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNtfsUnixSecurity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocols(formats); err != nil {
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

func (m *ExportRule) validateLinks(formats strfmt.Registry) error {
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

var exportRuleTypeChownModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["restricted","unrestricted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		exportRuleTypeChownModePropEnum = append(exportRuleTypeChownModePropEnum, v)
	}
}

const (

	// ExportRuleChownModeRestricted captures enum value "restricted"
	ExportRuleChownModeRestricted string = "restricted"

	// ExportRuleChownModeUnrestricted captures enum value "unrestricted"
	ExportRuleChownModeUnrestricted string = "unrestricted"
)

// prop value enum
func (m *ExportRule) validateChownModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, exportRuleTypeChownModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExportRule) validateChownMode(formats strfmt.Registry) error {
	if swag.IsZero(m.ChownMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateChownModeEnum("chown_mode", "body", *m.ChownMode); err != nil {
		return err
	}

	return nil
}

func (m *ExportRule) validateExportRuleInlineClients(formats strfmt.Registry) error {
	if swag.IsZero(m.ExportRuleInlineClients) { // not required
		return nil
	}

	for i := 0; i < len(m.ExportRuleInlineClients); i++ {
		if swag.IsZero(m.ExportRuleInlineClients[i]) { // not required
			continue
		}

		if m.ExportRuleInlineClients[i] != nil {
			if err := m.ExportRuleInlineClients[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clients" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExportRule) validateExportRuleInlineRoRule(formats strfmt.Registry) error {
	if swag.IsZero(m.ExportRuleInlineRoRule) { // not required
		return nil
	}

	for i := 0; i < len(m.ExportRuleInlineRoRule); i++ {
		if swag.IsZero(m.ExportRuleInlineRoRule[i]) { // not required
			continue
		}

		if m.ExportRuleInlineRoRule[i] != nil {
			if err := m.ExportRuleInlineRoRule[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ro_rule" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ro_rule" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExportRule) validateExportRuleInlineRwRule(formats strfmt.Registry) error {
	if swag.IsZero(m.ExportRuleInlineRwRule) { // not required
		return nil
	}

	for i := 0; i < len(m.ExportRuleInlineRwRule); i++ {
		if swag.IsZero(m.ExportRuleInlineRwRule[i]) { // not required
			continue
		}

		if m.ExportRuleInlineRwRule[i] != nil {
			if err := m.ExportRuleInlineRwRule[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rw_rule" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rw_rule" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExportRule) validateExportRuleInlineSuperuser(formats strfmt.Registry) error {
	if swag.IsZero(m.ExportRuleInlineSuperuser) { // not required
		return nil
	}

	for i := 0; i < len(m.ExportRuleInlineSuperuser); i++ {
		if swag.IsZero(m.ExportRuleInlineSuperuser[i]) { // not required
			continue
		}

		if m.ExportRuleInlineSuperuser[i] != nil {
			if err := m.ExportRuleInlineSuperuser[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("superuser" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("superuser" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var exportRuleTypeNtfsUnixSecurityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fail","ignore"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		exportRuleTypeNtfsUnixSecurityPropEnum = append(exportRuleTypeNtfsUnixSecurityPropEnum, v)
	}
}

const (

	// ExportRuleNtfsUnixSecurityFail captures enum value "fail"
	ExportRuleNtfsUnixSecurityFail string = "fail"

	// ExportRuleNtfsUnixSecurityIgnore captures enum value "ignore"
	ExportRuleNtfsUnixSecurityIgnore string = "ignore"
)

// prop value enum
func (m *ExportRule) validateNtfsUnixSecurityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, exportRuleTypeNtfsUnixSecurityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExportRule) validateNtfsUnixSecurity(formats strfmt.Registry) error {
	if swag.IsZero(m.NtfsUnixSecurity) { // not required
		return nil
	}

	// value enum
	if err := m.validateNtfsUnixSecurityEnum("ntfs_unix_security", "body", *m.NtfsUnixSecurity); err != nil {
		return err
	}

	return nil
}

func (m *ExportRule) validatePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if m.Policy != nil {
		if err := m.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

var exportRuleProtocolsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["any","nfs","nfs3","nfs4","cifs","flexcache"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		exportRuleProtocolsItemsEnum = append(exportRuleProtocolsItemsEnum, v)
	}
}

func (m *ExportRule) validateProtocolsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, exportRuleProtocolsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExportRule) validateProtocols(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocols) { // not required
		return nil
	}

	for i := 0; i < len(m.Protocols); i++ {
		if swag.IsZero(m.Protocols[i]) { // not required
			continue
		}

		// value enum
		if err := m.validateProtocolsItemsEnum("protocols"+"."+strconv.Itoa(i), "body", *m.Protocols[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExportRule) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this export rule based on the context it is used
func (m *ExportRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExportRuleInlineClients(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExportRuleInlineRoRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExportRuleInlineRwRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExportRuleInlineSuperuser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicy(ctx, formats); err != nil {
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

func (m *ExportRule) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ExportRule) contextValidateExportRuleInlineClients(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExportRuleInlineClients); i++ {

		if m.ExportRuleInlineClients[i] != nil {

			if swag.IsZero(m.ExportRuleInlineClients[i]) { // not required
				return nil
			}

			if err := m.ExportRuleInlineClients[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clients" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExportRule) contextValidateExportRuleInlineRoRule(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExportRuleInlineRoRule); i++ {

		if m.ExportRuleInlineRoRule[i] != nil {

			if swag.IsZero(m.ExportRuleInlineRoRule[i]) { // not required
				return nil
			}

			if err := m.ExportRuleInlineRoRule[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ro_rule" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ro_rule" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExportRule) contextValidateExportRuleInlineRwRule(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExportRuleInlineRwRule); i++ {

		if m.ExportRuleInlineRwRule[i] != nil {

			if swag.IsZero(m.ExportRuleInlineRwRule[i]) { // not required
				return nil
			}

			if err := m.ExportRuleInlineRwRule[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rw_rule" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rw_rule" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExportRule) contextValidateExportRuleInlineSuperuser(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExportRuleInlineSuperuser); i++ {

		if m.ExportRuleInlineSuperuser[i] != nil {

			if swag.IsZero(m.ExportRuleInlineSuperuser[i]) { // not required
				return nil
			}

			if err := m.ExportRuleInlineSuperuser[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("superuser" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("superuser" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExportRule) contextValidatePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.Policy != nil {

		if swag.IsZero(m.Policy) { // not required
			return nil
		}

		if err := m.Policy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

func (m *ExportRule) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ExportRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportRule) UnmarshalBinary(b []byte) error {
	var res ExportRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ExportRuleInlineLinks export rule inline links
//
// swagger:model export_rule_inline__links
type ExportRuleInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this export rule inline links
func (m *ExportRuleInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportRuleInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this export rule inline links based on the context it is used
func (m *ExportRuleInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportRuleInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ExportRuleInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportRuleInlineLinks) UnmarshalBinary(b []byte) error {
	var res ExportRuleInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ExportRuleInlinePolicy export rule inline policy
//
// swagger:model export_rule_inline_policy
type ExportRuleInlinePolicy struct {

	// Export policy ID
	ID *int64 `json:"id,omitempty" yaml:"id,omitempty"`

	// Export policy name
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this export rule inline policy
func (m *ExportRuleInlinePolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this export rule inline policy based on the context it is used
func (m *ExportRuleInlinePolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ExportRuleInlinePolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportRuleInlinePolicy) UnmarshalBinary(b []byte) error {
	var res ExportRuleInlinePolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ExportRuleInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model export_rule_inline_svm
type ExportRuleInlineSvm struct {

	// links
	Links *ExportRuleInlineSvmInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this export rule inline svm
func (m *ExportRuleInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportRuleInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this export rule inline svm based on the context it is used
func (m *ExportRuleInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportRuleInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ExportRuleInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportRuleInlineSvm) UnmarshalBinary(b []byte) error {
	var res ExportRuleInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ExportRuleInlineSvmInlineLinks export rule inline svm inline links
//
// swagger:model export_rule_inline_svm_inline__links
type ExportRuleInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this export rule inline svm inline links
func (m *ExportRuleInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportRuleInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this export rule inline svm inline links based on the context it is used
func (m *ExportRuleInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportRuleInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ExportRuleInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportRuleInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res ExportRuleInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
