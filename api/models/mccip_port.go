// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MccipPort Port configuration specification.
// l3_config information is only needed when configuring a MetroCluster IP for use in a layer 3 network.
//
// swagger:model mccip_port
type MccipPort struct {

	// l3 config
	L3Config *MccipPortInlineL3Config `json:"l3_config,omitempty" yaml:"l3_config,omitempty"`

	// Port name
	// Example: e1b
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// node
	Node *MccipPortInlineNode `json:"node,omitempty" yaml:"node,omitempty"`

	// Port UUID
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`

	// VLAN ID
	// Example: 200
	// Maximum: 4095
	// Minimum: 101
	VlanID *int64 `json:"vlan_id,omitempty" yaml:"vlan_id,omitempty"`
}

// Validate validates this mccip port
func (m *MccipPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateL3Config(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPort) validateL3Config(formats strfmt.Registry) error {
	if swag.IsZero(m.L3Config) { // not required
		return nil
	}

	if m.L3Config != nil {
		if err := m.L3Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l3_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("l3_config")
			}
			return err
		}
	}

	return nil
}

func (m *MccipPort) validateNode(formats strfmt.Registry) error {
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

func (m *MccipPort) validateVlanID(formats strfmt.Registry) error {
	if swag.IsZero(m.VlanID) { // not required
		return nil
	}

	if err := validate.MinimumInt("vlan_id", "body", *m.VlanID, 101, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vlan_id", "body", *m.VlanID, 4095, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this mccip port based on the context it is used
func (m *MccipPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateL3Config(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPort) contextValidateL3Config(ctx context.Context, formats strfmt.Registry) error {

	if m.L3Config != nil {

		if swag.IsZero(m.L3Config) { // not required
			return nil
		}

		if err := m.L3Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l3_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("l3_config")
			}
			return err
		}
	}

	return nil
}

func (m *MccipPort) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *MccipPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MccipPort) UnmarshalBinary(b []byte) error {
	var res MccipPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MccipPortInlineL3Config mccip port inline l3 config
//
// swagger:model mccip_port_inline_l3_config
type MccipPortInlineL3Config struct {

	// ipv4 interface
	IPV4Interface *MccipPortInlineL3ConfigInlineIPV4Interface `json:"ipv4_interface,omitempty" yaml:"ipv4_interface,omitempty"`
}

// Validate validates this mccip port inline l3 config
func (m *MccipPortInlineL3Config) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPV4Interface(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortInlineL3Config) validateIPV4Interface(formats strfmt.Registry) error {
	if swag.IsZero(m.IPV4Interface) { // not required
		return nil
	}

	if m.IPV4Interface != nil {
		if err := m.IPV4Interface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l3_config" + "." + "ipv4_interface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("l3_config" + "." + "ipv4_interface")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mccip port inline l3 config based on the context it is used
func (m *MccipPortInlineL3Config) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPV4Interface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortInlineL3Config) contextValidateIPV4Interface(ctx context.Context, formats strfmt.Registry) error {

	if m.IPV4Interface != nil {

		if swag.IsZero(m.IPV4Interface) { // not required
			return nil
		}

		if err := m.IPV4Interface.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l3_config" + "." + "ipv4_interface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("l3_config" + "." + "ipv4_interface")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MccipPortInlineL3Config) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MccipPortInlineL3Config) UnmarshalBinary(b []byte) error {
	var res MccipPortInlineL3Config
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MccipPortInlineL3ConfigInlineIPV4Interface Object to setup an interface along with its default router.
//
// swagger:model mccip_port_inline_l3_config_inline_ipv4_interface
type MccipPortInlineL3ConfigInlineIPV4Interface struct {

	// IPv4 or IPv6 address
	// Example: 10.10.10.7
	Address *string `json:"address,omitempty" yaml:"address,omitempty"`

	// The IPv4 or IPv6 address of the default router.
	// Example: 10.1.1.1
	Gateway *string `json:"gateway,omitempty" yaml:"gateway,omitempty"`

	// netmask
	Netmask *IPNetmask `json:"netmask,omitempty" yaml:"netmask,omitempty"`
}

// Validate validates this mccip port inline l3 config inline ipv4 interface
func (m *MccipPortInlineL3ConfigInlineIPV4Interface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetmask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortInlineL3ConfigInlineIPV4Interface) validateNetmask(formats strfmt.Registry) error {
	if swag.IsZero(m.Netmask) { // not required
		return nil
	}

	if m.Netmask != nil {
		if err := m.Netmask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l3_config" + "." + "ipv4_interface" + "." + "netmask")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("l3_config" + "." + "ipv4_interface" + "." + "netmask")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mccip port inline l3 config inline ipv4 interface based on the context it is used
func (m *MccipPortInlineL3ConfigInlineIPV4Interface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetmask(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortInlineL3ConfigInlineIPV4Interface) contextValidateNetmask(ctx context.Context, formats strfmt.Registry) error {

	if m.Netmask != nil {

		if swag.IsZero(m.Netmask) { // not required
			return nil
		}

		if err := m.Netmask.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l3_config" + "." + "ipv4_interface" + "." + "netmask")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("l3_config" + "." + "ipv4_interface" + "." + "netmask")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MccipPortInlineL3ConfigInlineIPV4Interface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MccipPortInlineL3ConfigInlineIPV4Interface) UnmarshalBinary(b []byte) error {
	var res MccipPortInlineL3ConfigInlineIPV4Interface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MccipPortInlineNode Node information
//
// swagger:model mccip_port_inline_node
type MccipPortInlineNode struct {

	// links
	Links *MccipPortInlineNodeInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this mccip port inline node
func (m *MccipPortInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortInlineNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this mccip port inline node based on the context it is used
func (m *MccipPortInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MccipPortInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MccipPortInlineNode) UnmarshalBinary(b []byte) error {
	var res MccipPortInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MccipPortInlineNodeInlineLinks mccip port inline node inline links
//
// swagger:model mccip_port_inline_node_inline__links
type MccipPortInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this mccip port inline node inline links
func (m *MccipPortInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this mccip port inline node inline links based on the context it is used
func (m *MccipPortInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MccipPortInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MccipPortInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res MccipPortInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
