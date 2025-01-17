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

// ApplicationProtectionGroups application protection groups
//
// swagger:model application_protection_groups
type ApplicationProtectionGroups struct {

	// Protection group name
	// Read Only: true
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// rpo
	Rpo *ApplicationProtectionGroupsInlineRpo `json:"rpo,omitempty" yaml:"rpo,omitempty"`

	// Protection group UUID
	// Read Only: true
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this application protection groups
func (m *ApplicationProtectionGroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRpo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationProtectionGroups) validateRpo(formats strfmt.Registry) error {
	if swag.IsZero(m.Rpo) { // not required
		return nil
	}

	if m.Rpo != nil {
		if err := m.Rpo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rpo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rpo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application protection groups based on the context it is used
func (m *ApplicationProtectionGroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRpo(ctx, formats); err != nil {
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

func (m *ApplicationProtectionGroups) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationProtectionGroups) contextValidateRpo(ctx context.Context, formats strfmt.Registry) error {

	if m.Rpo != nil {

		if swag.IsZero(m.Rpo) { // not required
			return nil
		}

		if err := m.Rpo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rpo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rpo")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationProtectionGroups) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationProtectionGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationProtectionGroups) UnmarshalBinary(b []byte) error {
	var res ApplicationProtectionGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationProtectionGroupsInlineRpo application protection groups inline rpo
//
// swagger:model application_protection_groups_inline_rpo
type ApplicationProtectionGroupsInlineRpo struct {

	// local
	Local *ApplicationProtectionGroupsInlineRpoInlineLocal `json:"local,omitempty" yaml:"local,omitempty"`

	// remote
	Remote *ApplicationProtectionGroupsInlineRpoInlineRemote `json:"remote,omitempty" yaml:"remote,omitempty"`
}

// Validate validates this application protection groups inline rpo
func (m *ApplicationProtectionGroupsInlineRpo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemote(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationProtectionGroupsInlineRpo) validateLocal(formats strfmt.Registry) error {
	if swag.IsZero(m.Local) { // not required
		return nil
	}

	if m.Local != nil {
		if err := m.Local.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rpo" + "." + "local")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rpo" + "." + "local")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationProtectionGroupsInlineRpo) validateRemote(formats strfmt.Registry) error {
	if swag.IsZero(m.Remote) { // not required
		return nil
	}

	if m.Remote != nil {
		if err := m.Remote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rpo" + "." + "remote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rpo" + "." + "remote")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application protection groups inline rpo based on the context it is used
func (m *ApplicationProtectionGroupsInlineRpo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationProtectionGroupsInlineRpo) contextValidateLocal(ctx context.Context, formats strfmt.Registry) error {

	if m.Local != nil {

		if swag.IsZero(m.Local) { // not required
			return nil
		}

		if err := m.Local.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rpo" + "." + "local")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rpo" + "." + "local")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationProtectionGroupsInlineRpo) contextValidateRemote(ctx context.Context, formats strfmt.Registry) error {

	if m.Remote != nil {

		if swag.IsZero(m.Remote) { // not required
			return nil
		}

		if err := m.Remote.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rpo" + "." + "remote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rpo" + "." + "remote")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationProtectionGroupsInlineRpo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationProtectionGroupsInlineRpo) UnmarshalBinary(b []byte) error {
	var res ApplicationProtectionGroupsInlineRpo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationProtectionGroupsInlineRpoInlineLocal application protection groups inline rpo inline local
//
// swagger:model application_protection_groups_inline_rpo_inline_local
type ApplicationProtectionGroupsInlineRpoInlineLocal struct {

	// A detailed description of the local RPO. This includes details on the Snapshot copy schedule.
	// Read Only: true
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`

	// The local RPO of the component. This indicates how often component Snapshot copies are automatically created.
	// Read Only: true
	// Enum: ["none","hourly","6_hourly","15_minutely"]
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this application protection groups inline rpo inline local
func (m *ApplicationProtectionGroupsInlineRpoInlineLocal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var applicationProtectionGroupsInlineRpoInlineLocalTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","hourly","6_hourly","15_minutely"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationProtectionGroupsInlineRpoInlineLocalTypeNamePropEnum = append(applicationProtectionGroupsInlineRpoInlineLocalTypeNamePropEnum, v)
	}
}

const (

	// ApplicationProtectionGroupsInlineRpoInlineLocalNameNone captures enum value "none"
	ApplicationProtectionGroupsInlineRpoInlineLocalNameNone string = "none"

	// ApplicationProtectionGroupsInlineRpoInlineLocalNameHourly captures enum value "hourly"
	ApplicationProtectionGroupsInlineRpoInlineLocalNameHourly string = "hourly"

	// ApplicationProtectionGroupsInlineRpoInlineLocalNameNr6Hourly captures enum value "6_hourly"
	ApplicationProtectionGroupsInlineRpoInlineLocalNameNr6Hourly string = "6_hourly"

	// ApplicationProtectionGroupsInlineRpoInlineLocalNameNr15Minutely captures enum value "15_minutely"
	ApplicationProtectionGroupsInlineRpoInlineLocalNameNr15Minutely string = "15_minutely"
)

// prop value enum
func (m *ApplicationProtectionGroupsInlineRpoInlineLocal) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, applicationProtectionGroupsInlineRpoInlineLocalTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationProtectionGroupsInlineRpoInlineLocal) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("rpo"+"."+"local"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this application protection groups inline rpo inline local based on the context it is used
func (m *ApplicationProtectionGroupsInlineRpoInlineLocal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationProtectionGroupsInlineRpoInlineLocal) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rpo"+"."+"local"+"."+"description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationProtectionGroupsInlineRpoInlineLocal) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rpo"+"."+"local"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationProtectionGroupsInlineRpoInlineLocal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationProtectionGroupsInlineRpoInlineLocal) UnmarshalBinary(b []byte) error {
	var res ApplicationProtectionGroupsInlineRpoInlineLocal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ApplicationProtectionGroupsInlineRpoInlineRemote application protection groups inline rpo inline remote
//
// swagger:model application_protection_groups_inline_rpo_inline_remote
type ApplicationProtectionGroupsInlineRpoInlineRemote struct {

	// A detailed description of the remote RPO.
	// Read Only: true
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`

	// The remote RPO of the component. A remote RPO of zero indicates that the component is synchronously replicated to another cluster.
	// Read Only: true
	// Enum: ["none","zero","hourly","6_hourly","15_minutely"]
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this application protection groups inline rpo inline remote
func (m *ApplicationProtectionGroupsInlineRpoInlineRemote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var applicationProtectionGroupsInlineRpoInlineRemoteTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","zero","hourly","6_hourly","15_minutely"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationProtectionGroupsInlineRpoInlineRemoteTypeNamePropEnum = append(applicationProtectionGroupsInlineRpoInlineRemoteTypeNamePropEnum, v)
	}
}

const (

	// ApplicationProtectionGroupsInlineRpoInlineRemoteNameNone captures enum value "none"
	ApplicationProtectionGroupsInlineRpoInlineRemoteNameNone string = "none"

	// ApplicationProtectionGroupsInlineRpoInlineRemoteNameZero captures enum value "zero"
	ApplicationProtectionGroupsInlineRpoInlineRemoteNameZero string = "zero"

	// ApplicationProtectionGroupsInlineRpoInlineRemoteNameHourly captures enum value "hourly"
	ApplicationProtectionGroupsInlineRpoInlineRemoteNameHourly string = "hourly"

	// ApplicationProtectionGroupsInlineRpoInlineRemoteNameNr6Hourly captures enum value "6_hourly"
	ApplicationProtectionGroupsInlineRpoInlineRemoteNameNr6Hourly string = "6_hourly"

	// ApplicationProtectionGroupsInlineRpoInlineRemoteNameNr15Minutely captures enum value "15_minutely"
	ApplicationProtectionGroupsInlineRpoInlineRemoteNameNr15Minutely string = "15_minutely"
)

// prop value enum
func (m *ApplicationProtectionGroupsInlineRpoInlineRemote) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, applicationProtectionGroupsInlineRpoInlineRemoteTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationProtectionGroupsInlineRpoInlineRemote) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("rpo"+"."+"remote"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this application protection groups inline rpo inline remote based on the context it is used
func (m *ApplicationProtectionGroupsInlineRpoInlineRemote) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationProtectionGroupsInlineRpoInlineRemote) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rpo"+"."+"remote"+"."+"description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationProtectionGroupsInlineRpoInlineRemote) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rpo"+"."+"remote"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationProtectionGroupsInlineRpoInlineRemote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationProtectionGroupsInlineRpoInlineRemote) UnmarshalBinary(b []byte) error {
	var res ApplicationProtectionGroupsInlineRpoInlineRemote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
