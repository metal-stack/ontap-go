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

// ActiveDirectory active directory
//
// swagger:model active_directory
type ActiveDirectory struct {

	// Specifies the discovered servers records.
	// Read Only: true
	ActiveDirectoryInlineDiscoveredServers []*ActiveDirectoryInlineDiscoveredServersInlineArrayItem `json:"discovered_servers,omitempty" yaml:"discovered_servers,omitempty"`

	// Specifies the preferred domain controller (DC) records.
	// Read Only: true
	ActiveDirectoryInlinePreferredDcs []*ActiveDirectoryInlinePreferredDcsInlineArrayItem `json:"preferred_dcs,omitempty" yaml:"preferred_dcs,omitempty"`

	// If set to true and a machine account exists with the same name as specified in "name" in Active Directory, it will be overwritten and reused.
	// Example: false
	ForceAccountOverwrite *bool `json:"force_account_overwrite,omitempty" yaml:"force_account_overwrite,omitempty"`

	// Fully qualified domain name.
	// Example: server1.com
	// Max Length: 254
	Fqdn *string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`

	// Active Directory (AD) account NetBIOS name.
	// Example: account1
	// Max Length: 15
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// Organizational unit under which the Active Directory account will be created.
	// Example: CN=Test
	OrganizationalUnit *string `json:"organizational_unit,omitempty" yaml:"organizational_unit,omitempty"`

	// Administrator password required for Active Directory account creation, modification and deletion.
	// Example: testpwd
	// Min Length: 1
	Password *string `json:"password,omitempty" yaml:"password,omitempty"`

	// svm
	Svm *ActiveDirectoryInlineSvm `json:"svm,omitempty" yaml:"svm,omitempty"`

	// Administrator username required for Active Directory account creation, modification and deletion.
	// Example: admin
	// Min Length: 1
	Username *string `json:"username,omitempty" yaml:"username,omitempty"`
}

// Validate validates this active directory
func (m *ActiveDirectory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveDirectoryInlineDiscoveredServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActiveDirectoryInlinePreferredDcs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFqdn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectory) validateActiveDirectoryInlineDiscoveredServers(formats strfmt.Registry) error {
	if swag.IsZero(m.ActiveDirectoryInlineDiscoveredServers) { // not required
		return nil
	}

	for i := 0; i < len(m.ActiveDirectoryInlineDiscoveredServers); i++ {
		if swag.IsZero(m.ActiveDirectoryInlineDiscoveredServers[i]) { // not required
			continue
		}

		if m.ActiveDirectoryInlineDiscoveredServers[i] != nil {
			if err := m.ActiveDirectoryInlineDiscoveredServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("discovered_servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("discovered_servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ActiveDirectory) validateActiveDirectoryInlinePreferredDcs(formats strfmt.Registry) error {
	if swag.IsZero(m.ActiveDirectoryInlinePreferredDcs) { // not required
		return nil
	}

	for i := 0; i < len(m.ActiveDirectoryInlinePreferredDcs); i++ {
		if swag.IsZero(m.ActiveDirectoryInlinePreferredDcs[i]) { // not required
			continue
		}

		if m.ActiveDirectoryInlinePreferredDcs[i] != nil {
			if err := m.ActiveDirectoryInlinePreferredDcs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("preferred_dcs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("preferred_dcs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ActiveDirectory) validateFqdn(formats strfmt.Registry) error {
	if swag.IsZero(m.Fqdn) { // not required
		return nil
	}

	if err := validate.MaxLength("fqdn", "body", *m.Fqdn, 254); err != nil {
		return err
	}

	return nil
}

func (m *ActiveDirectory) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", *m.Name, 15); err != nil {
		return err
	}

	return nil
}

func (m *ActiveDirectory) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.MinLength("password", "body", *m.Password, 1); err != nil {
		return err
	}

	return nil
}

func (m *ActiveDirectory) validateSvm(formats strfmt.Registry) error {
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

func (m *ActiveDirectory) validateUsername(formats strfmt.Registry) error {
	if swag.IsZero(m.Username) { // not required
		return nil
	}

	if err := validate.MinLength("username", "body", *m.Username, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this active directory based on the context it is used
func (m *ActiveDirectory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActiveDirectoryInlineDiscoveredServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateActiveDirectoryInlinePreferredDcs(ctx, formats); err != nil {
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

func (m *ActiveDirectory) contextValidateActiveDirectoryInlineDiscoveredServers(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "discovered_servers", "body", []*ActiveDirectoryInlineDiscoveredServersInlineArrayItem(m.ActiveDirectoryInlineDiscoveredServers)); err != nil {
		return err
	}

	for i := 0; i < len(m.ActiveDirectoryInlineDiscoveredServers); i++ {

		if m.ActiveDirectoryInlineDiscoveredServers[i] != nil {

			if swag.IsZero(m.ActiveDirectoryInlineDiscoveredServers[i]) { // not required
				return nil
			}

			if err := m.ActiveDirectoryInlineDiscoveredServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("discovered_servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("discovered_servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ActiveDirectory) contextValidateActiveDirectoryInlinePreferredDcs(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "preferred_dcs", "body", []*ActiveDirectoryInlinePreferredDcsInlineArrayItem(m.ActiveDirectoryInlinePreferredDcs)); err != nil {
		return err
	}

	for i := 0; i < len(m.ActiveDirectoryInlinePreferredDcs); i++ {

		if m.ActiveDirectoryInlinePreferredDcs[i] != nil {

			if swag.IsZero(m.ActiveDirectoryInlinePreferredDcs[i]) { // not required
				return nil
			}

			if err := m.ActiveDirectoryInlinePreferredDcs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("preferred_dcs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("preferred_dcs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ActiveDirectory) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ActiveDirectory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveDirectory) UnmarshalBinary(b []byte) error {
	var res ActiveDirectory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ActiveDirectoryInlineDiscoveredServersInlineArrayItem active directory inline discovered servers inline array item
//
// swagger:model active_directory_inline_discovered_servers_inline_array_item
type ActiveDirectoryInlineDiscoveredServersInlineArrayItem struct {

	// The Active Directory domain that the discovered server is a member of.
	// Example: server1.com
	Domain *string `json:"domain,omitempty" yaml:"domain,omitempty"`

	// node
	Node *ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNode `json:"node,omitempty" yaml:"node,omitempty"`

	// The preference level of the server that was discovered.
	// Example: preferred
	// Enum: ["unknown","preferred","favored","adequate"]
	Preference *string `json:"preference,omitempty" yaml:"preference,omitempty"`

	// server
	Server *ActiveDirectoryDiscoveredServerReference `json:"server,omitempty" yaml:"server,omitempty"`

	// The status of the connection to the server that was discovered.
	// Example: ok
	// Enum: ["ok","unavailable","slow","expired","undetermined","unreachable"]
	State *string `json:"state,omitempty" yaml:"state,omitempty"`
}

// Validate validates this active directory inline discovered servers inline array item
func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItem) validateNode(formats strfmt.Registry) error {
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

var activeDirectoryInlineDiscoveredServersInlineArrayItemTypePreferencePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","preferred","favored","adequate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		activeDirectoryInlineDiscoveredServersInlineArrayItemTypePreferencePropEnum = append(activeDirectoryInlineDiscoveredServersInlineArrayItemTypePreferencePropEnum, v)
	}
}

const (

	// ActiveDirectoryInlineDiscoveredServersInlineArrayItemPreferenceUnknown captures enum value "unknown"
	ActiveDirectoryInlineDiscoveredServersInlineArrayItemPreferenceUnknown string = "unknown"

	// ActiveDirectoryInlineDiscoveredServersInlineArrayItemPreferencePreferred captures enum value "preferred"
	ActiveDirectoryInlineDiscoveredServersInlineArrayItemPreferencePreferred string = "preferred"

	// ActiveDirectoryInlineDiscoveredServersInlineArrayItemPreferenceFavored captures enum value "favored"
	ActiveDirectoryInlineDiscoveredServersInlineArrayItemPreferenceFavored string = "favored"

	// ActiveDirectoryInlineDiscoveredServersInlineArrayItemPreferenceAdequate captures enum value "adequate"
	ActiveDirectoryInlineDiscoveredServersInlineArrayItemPreferenceAdequate string = "adequate"
)

// prop value enum
func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItem) validatePreferenceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, activeDirectoryInlineDiscoveredServersInlineArrayItemTypePreferencePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItem) validatePreference(formats strfmt.Registry) error {
	if swag.IsZero(m.Preference) { // not required
		return nil
	}

	// value enum
	if err := m.validatePreferenceEnum("preference", "body", *m.Preference); err != nil {
		return err
	}

	return nil
}

func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItem) validateServer(formats strfmt.Registry) error {
	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

var activeDirectoryInlineDiscoveredServersInlineArrayItemTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","unavailable","slow","expired","undetermined","unreachable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		activeDirectoryInlineDiscoveredServersInlineArrayItemTypeStatePropEnum = append(activeDirectoryInlineDiscoveredServersInlineArrayItemTypeStatePropEnum, v)
	}
}

const (

	// ActiveDirectoryInlineDiscoveredServersInlineArrayItemStateOk captures enum value "ok"
	ActiveDirectoryInlineDiscoveredServersInlineArrayItemStateOk string = "ok"

	// ActiveDirectoryInlineDiscoveredServersInlineArrayItemStateUnavailable captures enum value "unavailable"
	ActiveDirectoryInlineDiscoveredServersInlineArrayItemStateUnavailable string = "unavailable"

	// ActiveDirectoryInlineDiscoveredServersInlineArrayItemStateSlow captures enum value "slow"
	ActiveDirectoryInlineDiscoveredServersInlineArrayItemStateSlow string = "slow"

	// ActiveDirectoryInlineDiscoveredServersInlineArrayItemStateExpired captures enum value "expired"
	ActiveDirectoryInlineDiscoveredServersInlineArrayItemStateExpired string = "expired"

	// ActiveDirectoryInlineDiscoveredServersInlineArrayItemStateUndetermined captures enum value "undetermined"
	ActiveDirectoryInlineDiscoveredServersInlineArrayItemStateUndetermined string = "undetermined"

	// ActiveDirectoryInlineDiscoveredServersInlineArrayItemStateUnreachable captures enum value "unreachable"
	ActiveDirectoryInlineDiscoveredServersInlineArrayItemStateUnreachable string = "unreachable"
)

// prop value enum
func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItem) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, activeDirectoryInlineDiscoveredServersInlineArrayItemTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItem) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this active directory inline discovered servers inline array item based on the context it is used
func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItem) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItem) contextValidateServer(ctx context.Context, formats strfmt.Registry) error {

	if m.Server != nil {

		if swag.IsZero(m.Server) { // not required
			return nil
		}

		if err := m.Server.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res ActiveDirectoryInlineDiscoveredServersInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNode active directory inline discovered servers inline array item inline node
//
// swagger:model active_directory_inline_discovered_servers_inline_array_item_inline_node
type ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNode struct {

	// links
	Links *ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNodeInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this active directory inline discovered servers inline array item inline node
func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this active directory inline discovered servers inline array item inline node based on the context it is used
func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNode) UnmarshalBinary(b []byte) error {
	var res ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNodeInlineLinks active directory inline discovered servers inline array item inline node inline links
//
// swagger:model active_directory_inline_discovered_servers_inline_array_item_inline_node_inline__links
type ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this active directory inline discovered servers inline array item inline node inline links
func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this active directory inline discovered servers inline array item inline node inline links based on the context it is used
func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res ActiveDirectoryInlineDiscoveredServersInlineArrayItemInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ActiveDirectoryInlinePreferredDcsInlineArrayItem active directory inline preferred dcs inline array item
//
// swagger:model active_directory_inline_preferred_dcs_inline_array_item
type ActiveDirectoryInlinePreferredDcsInlineArrayItem struct {

	// Fully Qualified Domain Name.
	// Example: test.com
	// Max Length: 254
	Fqdn *string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`

	// IP address of the preferred DC. The address can be either an IPv4 or an IPv6 address.
	// Example: 4.4.4.4
	ServerIP *string `json:"server_ip,omitempty" yaml:"server_ip,omitempty"`
}

// Validate validates this active directory inline preferred dcs inline array item
func (m *ActiveDirectoryInlinePreferredDcsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFqdn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectoryInlinePreferredDcsInlineArrayItem) validateFqdn(formats strfmt.Registry) error {
	if swag.IsZero(m.Fqdn) { // not required
		return nil
	}

	if err := validate.MaxLength("fqdn", "body", *m.Fqdn, 254); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this active directory inline preferred dcs inline array item based on context it is used
func (m *ActiveDirectoryInlinePreferredDcsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActiveDirectoryInlinePreferredDcsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveDirectoryInlinePreferredDcsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res ActiveDirectoryInlinePreferredDcsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ActiveDirectoryInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model active_directory_inline_svm
type ActiveDirectoryInlineSvm struct {

	// links
	Links *ActiveDirectoryInlineSvmInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this active directory inline svm
func (m *ActiveDirectoryInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectoryInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this active directory inline svm based on the context it is used
func (m *ActiveDirectoryInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectoryInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ActiveDirectoryInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveDirectoryInlineSvm) UnmarshalBinary(b []byte) error {
	var res ActiveDirectoryInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ActiveDirectoryInlineSvmInlineLinks active directory inline svm inline links
//
// swagger:model active_directory_inline_svm_inline__links
type ActiveDirectoryInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this active directory inline svm inline links
func (m *ActiveDirectoryInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectoryInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this active directory inline svm inline links based on the context it is used
func (m *ActiveDirectoryInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectoryInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ActiveDirectoryInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveDirectoryInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res ActiveDirectoryInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
