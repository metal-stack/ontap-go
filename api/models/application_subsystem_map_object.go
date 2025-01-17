// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ApplicationSubsystemMapObject Subsystem map object
//
// swagger:model application_subsystem_map_object
type ApplicationSubsystemMapObject struct {

	// Subsystem ANA group ID
	// Read Only: true
	Anagrpid *string `json:"anagrpid,omitempty" yaml:"anagrpid,omitempty"`

	// Subsystem namespace ID
	// Read Only: true
	Nsid *string `json:"nsid,omitempty" yaml:"nsid,omitempty"`

	// subsystem
	Subsystem *ApplicationSubsystemMapObjectInlineSubsystem `json:"subsystem,omitempty" yaml:"subsystem,omitempty"`
}

// Validate validates this application subsystem map object
func (m *ApplicationSubsystemMapObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubsystem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSubsystemMapObject) validateSubsystem(formats strfmt.Registry) error {
	if swag.IsZero(m.Subsystem) { // not required
		return nil
	}

	if m.Subsystem != nil {
		if err := m.Subsystem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subsystem")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application subsystem map object based on the context it is used
func (m *ApplicationSubsystemMapObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAnagrpid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNsid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubsystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSubsystemMapObject) contextValidateAnagrpid(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "anagrpid", "body", m.Anagrpid); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationSubsystemMapObject) contextValidateNsid(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nsid", "body", m.Nsid); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationSubsystemMapObject) contextValidateSubsystem(ctx context.Context, formats strfmt.Registry) error {

	if m.Subsystem != nil {

		if swag.IsZero(m.Subsystem) { // not required
			return nil
		}

		if err := m.Subsystem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subsystem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationSubsystemMapObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSubsystemMapObject) UnmarshalBinary(b []byte) error {
	var res ApplicationSubsystemMapObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationSubsystemMapObjectInlineSubsystem application subsystem map object inline subsystem
//
// swagger:model application_subsystem_map_object_inline_subsystem
type ApplicationSubsystemMapObjectInlineSubsystem struct {

	// links
	Links *ApplicationSubsystemMapObjectInlineSubsystemInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// hosts
	// Read Only: true
	Hosts []*ApplicationSubsystemMapObjectSubsystemHostsItems0 `json:"hosts,omitempty" yaml:"hosts,omitempty"`

	// Subsystem name
	// Read Only: true
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// Subsystem UUID
	// Read Only: true
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this application subsystem map object inline subsystem
func (m *ApplicationSubsystemMapObjectInlineSubsystem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSubsystemMapObjectInlineSubsystem) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subsystem" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationSubsystemMapObjectInlineSubsystem) validateHosts(formats strfmt.Registry) error {
	if swag.IsZero(m.Hosts) { // not required
		return nil
	}

	for i := 0; i < len(m.Hosts); i++ {
		if swag.IsZero(m.Hosts[i]) { // not required
			continue
		}

		if m.Hosts[i] != nil {
			if err := m.Hosts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subsystem" + "." + "hosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subsystem" + "." + "hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this application subsystem map object inline subsystem based on the context it is used
func (m *ApplicationSubsystemMapObjectInlineSubsystem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSubsystemMapObjectInlineSubsystem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subsystem" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationSubsystemMapObjectInlineSubsystem) contextValidateHosts(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "subsystem"+"."+"hosts", "body", []*ApplicationSubsystemMapObjectSubsystemHostsItems0(m.Hosts)); err != nil {
		return err
	}

	for i := 0; i < len(m.Hosts); i++ {

		if m.Hosts[i] != nil {

			if swag.IsZero(m.Hosts[i]) { // not required
				return nil
			}

			if err := m.Hosts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subsystem" + "." + "hosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subsystem" + "." + "hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationSubsystemMapObjectInlineSubsystem) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "subsystem"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationSubsystemMapObjectInlineSubsystem) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "subsystem"+"."+"uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationSubsystemMapObjectInlineSubsystem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSubsystemMapObjectInlineSubsystem) UnmarshalBinary(b []byte) error {
	var res ApplicationSubsystemMapObjectInlineSubsystem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationSubsystemMapObjectSubsystemHostsItems0 application subsystem map object subsystem hosts items0
//
// swagger:model ApplicationSubsystemMapObjectSubsystemHostsItems0
type ApplicationSubsystemMapObjectSubsystemHostsItems0 struct {

	// links
	Links *ApplicationSubsystemMapObjectSubsystemHostsItems0Links `json:"_links,omitempty" yaml:"_links,omitempty"`

	// Host
	// Read Only: true
	Nqn *string `json:"nqn,omitempty" yaml:"nqn,omitempty"`
}

// Validate validates this application subsystem map object subsystem hosts items0
func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this application subsystem map object subsystem hosts items0 based on the context it is used
func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNqn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0) contextValidateNqn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nqn", "body", m.Nqn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0) UnmarshalBinary(b []byte) error {
	var res ApplicationSubsystemMapObjectSubsystemHostsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationSubsystemMapObjectSubsystemHostsItems0Links application subsystem map object subsystem hosts items0 links
//
// swagger:model ApplicationSubsystemMapObjectSubsystemHostsItems0Links
type ApplicationSubsystemMapObjectSubsystemHostsItems0Links struct {

	// self
	Self *ApplicationSubsystemMapObjectSubsystemHostsItems0LinksSelf `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this application subsystem map object subsystem hosts items0 links
func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0Links) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this application subsystem map object subsystem hosts items0 links based on the context it is used
func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0Links) UnmarshalBinary(b []byte) error {
	var res ApplicationSubsystemMapObjectSubsystemHostsItems0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationSubsystemMapObjectSubsystemHostsItems0LinksSelf application subsystem map object subsystem hosts items0 links self
//
// swagger:model ApplicationSubsystemMapObjectSubsystemHostsItems0LinksSelf
type ApplicationSubsystemMapObjectSubsystemHostsItems0LinksSelf struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this application subsystem map object subsystem hosts items0 links self
func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0LinksSelf) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0LinksSelf) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application subsystem map object subsystem hosts items0 links self based on the context it is used
func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0LinksSelf) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0LinksSelf) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0LinksSelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSubsystemMapObjectSubsystemHostsItems0LinksSelf) UnmarshalBinary(b []byte) error {
	var res ApplicationSubsystemMapObjectSubsystemHostsItems0LinksSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationSubsystemMapObjectInlineSubsystemInlineLinks application subsystem map object inline subsystem inline links
//
// swagger:model application_subsystem_map_object_inline_subsystem_inline__links
type ApplicationSubsystemMapObjectInlineSubsystemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this application subsystem map object inline subsystem inline links
func (m *ApplicationSubsystemMapObjectInlineSubsystemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSubsystemMapObjectInlineSubsystemInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subsystem" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application subsystem map object inline subsystem inline links based on the context it is used
func (m *ApplicationSubsystemMapObjectInlineSubsystemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSubsystemMapObjectInlineSubsystemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subsystem" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationSubsystemMapObjectInlineSubsystemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSubsystemMapObjectInlineSubsystemInlineLinks) UnmarshalBinary(b []byte) error {
	var res ApplicationSubsystemMapObjectInlineSubsystemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
