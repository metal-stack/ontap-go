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

// ExportRules export rules
//
// swagger:model export_rules
type ExportRules struct {

	// links
	Links *ExportRulesInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

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
	ExportRulesInlineClients []*ExportClients `json:"clients,omitempty" yaml:"clients,omitempty"`

	// Authentication flavors that the read-only access rule governs
	//
	ExportRulesInlineRoRule []*ExportAuthenticationFlavor `json:"ro_rule,omitempty" yaml:"ro_rule,omitempty"`

	// Authentication flavors that the read/write access rule governs
	//
	ExportRulesInlineRwRule []*ExportAuthenticationFlavor `json:"rw_rule,omitempty" yaml:"rw_rule,omitempty"`

	// Authentication flavors that the superuser security type governs
	//
	ExportRulesInlineSuperuser []*ExportAuthenticationFlavor `json:"superuser,omitempty" yaml:"superuser,omitempty"`

	// Index of the rule within the export policy.
	//
	Index *int64 `json:"index,omitempty" yaml:"index,omitempty"`

	// NTFS export UNIX security options.
	// Enum: ["fail","ignore"]
	NtfsUnixSecurity *string `json:"ntfs_unix_security,omitempty" yaml:"ntfs_unix_security,omitempty"`

	// protocols
	Protocols []*string `json:"protocols,omitempty" yaml:"protocols,omitempty"`
}

// Validate validates this export rules
func (m *ExportRules) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChownMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportRulesInlineClients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportRulesInlineRoRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportRulesInlineRwRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportRulesInlineSuperuser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNtfsUnixSecurity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocols(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportRules) validateLinks(formats strfmt.Registry) error {
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

var exportRulesTypeChownModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["restricted","unrestricted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		exportRulesTypeChownModePropEnum = append(exportRulesTypeChownModePropEnum, v)
	}
}

const (

	// ExportRulesChownModeRestricted captures enum value "restricted"
	ExportRulesChownModeRestricted string = "restricted"

	// ExportRulesChownModeUnrestricted captures enum value "unrestricted"
	ExportRulesChownModeUnrestricted string = "unrestricted"
)

// prop value enum
func (m *ExportRules) validateChownModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, exportRulesTypeChownModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExportRules) validateChownMode(formats strfmt.Registry) error {
	if swag.IsZero(m.ChownMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateChownModeEnum("chown_mode", "body", *m.ChownMode); err != nil {
		return err
	}

	return nil
}

func (m *ExportRules) validateExportRulesInlineClients(formats strfmt.Registry) error {
	if swag.IsZero(m.ExportRulesInlineClients) { // not required
		return nil
	}

	for i := 0; i < len(m.ExportRulesInlineClients); i++ {
		if swag.IsZero(m.ExportRulesInlineClients[i]) { // not required
			continue
		}

		if m.ExportRulesInlineClients[i] != nil {
			if err := m.ExportRulesInlineClients[i].Validate(formats); err != nil {
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

func (m *ExportRules) validateExportRulesInlineRoRule(formats strfmt.Registry) error {
	if swag.IsZero(m.ExportRulesInlineRoRule) { // not required
		return nil
	}

	for i := 0; i < len(m.ExportRulesInlineRoRule); i++ {
		if swag.IsZero(m.ExportRulesInlineRoRule[i]) { // not required
			continue
		}

		if m.ExportRulesInlineRoRule[i] != nil {
			if err := m.ExportRulesInlineRoRule[i].Validate(formats); err != nil {
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

func (m *ExportRules) validateExportRulesInlineRwRule(formats strfmt.Registry) error {
	if swag.IsZero(m.ExportRulesInlineRwRule) { // not required
		return nil
	}

	for i := 0; i < len(m.ExportRulesInlineRwRule); i++ {
		if swag.IsZero(m.ExportRulesInlineRwRule[i]) { // not required
			continue
		}

		if m.ExportRulesInlineRwRule[i] != nil {
			if err := m.ExportRulesInlineRwRule[i].Validate(formats); err != nil {
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

func (m *ExportRules) validateExportRulesInlineSuperuser(formats strfmt.Registry) error {
	if swag.IsZero(m.ExportRulesInlineSuperuser) { // not required
		return nil
	}

	for i := 0; i < len(m.ExportRulesInlineSuperuser); i++ {
		if swag.IsZero(m.ExportRulesInlineSuperuser[i]) { // not required
			continue
		}

		if m.ExportRulesInlineSuperuser[i] != nil {
			if err := m.ExportRulesInlineSuperuser[i].Validate(formats); err != nil {
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

var exportRulesTypeNtfsUnixSecurityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fail","ignore"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		exportRulesTypeNtfsUnixSecurityPropEnum = append(exportRulesTypeNtfsUnixSecurityPropEnum, v)
	}
}

const (

	// ExportRulesNtfsUnixSecurityFail captures enum value "fail"
	ExportRulesNtfsUnixSecurityFail string = "fail"

	// ExportRulesNtfsUnixSecurityIgnore captures enum value "ignore"
	ExportRulesNtfsUnixSecurityIgnore string = "ignore"
)

// prop value enum
func (m *ExportRules) validateNtfsUnixSecurityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, exportRulesTypeNtfsUnixSecurityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExportRules) validateNtfsUnixSecurity(formats strfmt.Registry) error {
	if swag.IsZero(m.NtfsUnixSecurity) { // not required
		return nil
	}

	// value enum
	if err := m.validateNtfsUnixSecurityEnum("ntfs_unix_security", "body", *m.NtfsUnixSecurity); err != nil {
		return err
	}

	return nil
}

var exportRulesProtocolsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["any","nfs","nfs3","nfs4","cifs","flexcache"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		exportRulesProtocolsItemsEnum = append(exportRulesProtocolsItemsEnum, v)
	}
}

func (m *ExportRules) validateProtocolsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, exportRulesProtocolsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExportRules) validateProtocols(formats strfmt.Registry) error {
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

// ContextValidate validate this export rules based on the context it is used
func (m *ExportRules) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExportRulesInlineClients(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExportRulesInlineRoRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExportRulesInlineRwRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExportRulesInlineSuperuser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportRules) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ExportRules) contextValidateExportRulesInlineClients(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExportRulesInlineClients); i++ {

		if m.ExportRulesInlineClients[i] != nil {

			if swag.IsZero(m.ExportRulesInlineClients[i]) { // not required
				return nil
			}

			if err := m.ExportRulesInlineClients[i].ContextValidate(ctx, formats); err != nil {
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

func (m *ExportRules) contextValidateExportRulesInlineRoRule(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExportRulesInlineRoRule); i++ {

		if m.ExportRulesInlineRoRule[i] != nil {

			if swag.IsZero(m.ExportRulesInlineRoRule[i]) { // not required
				return nil
			}

			if err := m.ExportRulesInlineRoRule[i].ContextValidate(ctx, formats); err != nil {
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

func (m *ExportRules) contextValidateExportRulesInlineRwRule(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExportRulesInlineRwRule); i++ {

		if m.ExportRulesInlineRwRule[i] != nil {

			if swag.IsZero(m.ExportRulesInlineRwRule[i]) { // not required
				return nil
			}

			if err := m.ExportRulesInlineRwRule[i].ContextValidate(ctx, formats); err != nil {
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

func (m *ExportRules) contextValidateExportRulesInlineSuperuser(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExportRulesInlineSuperuser); i++ {

		if m.ExportRulesInlineSuperuser[i] != nil {

			if swag.IsZero(m.ExportRulesInlineSuperuser[i]) { // not required
				return nil
			}

			if err := m.ExportRulesInlineSuperuser[i].ContextValidate(ctx, formats); err != nil {
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

// MarshalBinary interface implementation
func (m *ExportRules) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportRules) UnmarshalBinary(b []byte) error {
	var res ExportRules
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ExportRulesInlineLinks export rules inline links
//
// swagger:model export_rules_inline__links
type ExportRulesInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this export rules inline links
func (m *ExportRulesInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportRulesInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this export rules inline links based on the context it is used
func (m *ExportRulesInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportRulesInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ExportRulesInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportRulesInlineLinks) UnmarshalBinary(b []byte) error {
	var res ExportRulesInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
