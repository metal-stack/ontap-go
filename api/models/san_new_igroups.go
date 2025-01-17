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

// SanNewIgroups The list of initiator groups to create.
//
// swagger:model san_new_igroups
type SanNewIgroups struct {

	// A comment available for use by the administrator.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// igroups
	Igroups []*SanNewIgroupsIgroupsItems0 `json:"igroups" yaml:"igroups"`

	// initiator objects
	InitiatorObjects []*SanNewIgroupsInitiatorObjectsItems0 `json:"initiator_objects" yaml:"initiator_objects"`

	// initiators
	Initiators []string `json:"initiators" yaml:"initiators"`

	// The name of the new initiator group.
	// Required: true
	// Max Length: 96
	// Min Length: 1
	Name *string `json:"name" yaml:"name"`

	// The name of the host OS accessing the application. The default value is the host OS that is running the application.
	// Enum: ["aix","hpux","hyper_v","linux","netware","openvms","solaris","vmware","windows","xen"]
	OsType string `json:"os_type,omitempty" yaml:"os_type,omitempty"`

	// The protocol of the new initiator group.
	// Enum: ["fcp","iscsi","mixed"]
	Protocol *string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}

// Validate validates this san new igroups
func (m *SanNewIgroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIgroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiatorObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SanNewIgroups) validateIgroups(formats strfmt.Registry) error {
	if swag.IsZero(m.Igroups) { // not required
		return nil
	}

	for i := 0; i < len(m.Igroups); i++ {
		if swag.IsZero(m.Igroups[i]) { // not required
			continue
		}

		if m.Igroups[i] != nil {
			if err := m.Igroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("igroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SanNewIgroups) validateInitiatorObjects(formats strfmt.Registry) error {
	if swag.IsZero(m.InitiatorObjects) { // not required
		return nil
	}

	for i := 0; i < len(m.InitiatorObjects); i++ {
		if swag.IsZero(m.InitiatorObjects[i]) { // not required
			continue
		}

		if m.InitiatorObjects[i] != nil {
			if err := m.InitiatorObjects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initiator_objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("initiator_objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SanNewIgroups) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 96); err != nil {
		return err
	}

	return nil
}

var sanNewIgroupsTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aix","hpux","hyper_v","linux","netware","openvms","solaris","vmware","windows","xen"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sanNewIgroupsTypeOsTypePropEnum = append(sanNewIgroupsTypeOsTypePropEnum, v)
	}
}

const (

	// SanNewIgroupsOsTypeAix captures enum value "aix"
	SanNewIgroupsOsTypeAix string = "aix"

	// SanNewIgroupsOsTypeHpux captures enum value "hpux"
	SanNewIgroupsOsTypeHpux string = "hpux"

	// SanNewIgroupsOsTypeHyperv captures enum value "hyper_v"
	SanNewIgroupsOsTypeHyperv string = "hyper_v"

	// SanNewIgroupsOsTypeLinux captures enum value "linux"
	SanNewIgroupsOsTypeLinux string = "linux"

	// SanNewIgroupsOsTypeNetware captures enum value "netware"
	SanNewIgroupsOsTypeNetware string = "netware"

	// SanNewIgroupsOsTypeOpenvms captures enum value "openvms"
	SanNewIgroupsOsTypeOpenvms string = "openvms"

	// SanNewIgroupsOsTypeSolaris captures enum value "solaris"
	SanNewIgroupsOsTypeSolaris string = "solaris"

	// SanNewIgroupsOsTypeVmware captures enum value "vmware"
	SanNewIgroupsOsTypeVmware string = "vmware"

	// SanNewIgroupsOsTypeWindows captures enum value "windows"
	SanNewIgroupsOsTypeWindows string = "windows"

	// SanNewIgroupsOsTypeXen captures enum value "xen"
	SanNewIgroupsOsTypeXen string = "xen"
)

// prop value enum
func (m *SanNewIgroups) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sanNewIgroupsTypeOsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SanNewIgroups) validateOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsTypeEnum("os_type", "body", m.OsType); err != nil {
		return err
	}

	return nil
}

var sanNewIgroupsTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fcp","iscsi","mixed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sanNewIgroupsTypeProtocolPropEnum = append(sanNewIgroupsTypeProtocolPropEnum, v)
	}
}

const (

	// SanNewIgroupsProtocolFcp captures enum value "fcp"
	SanNewIgroupsProtocolFcp string = "fcp"

	// SanNewIgroupsProtocolIscsi captures enum value "iscsi"
	SanNewIgroupsProtocolIscsi string = "iscsi"

	// SanNewIgroupsProtocolMixed captures enum value "mixed"
	SanNewIgroupsProtocolMixed string = "mixed"
)

// prop value enum
func (m *SanNewIgroups) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sanNewIgroupsTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SanNewIgroups) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this san new igroups based on the context it is used
func (m *SanNewIgroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIgroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInitiatorObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SanNewIgroups) contextValidateIgroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Igroups); i++ {

		if m.Igroups[i] != nil {

			if swag.IsZero(m.Igroups[i]) { // not required
				return nil
			}

			if err := m.Igroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("igroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SanNewIgroups) contextValidateInitiatorObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InitiatorObjects); i++ {

		if m.InitiatorObjects[i] != nil {

			if swag.IsZero(m.InitiatorObjects[i]) { // not required
				return nil
			}

			if err := m.InitiatorObjects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initiator_objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("initiator_objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SanNewIgroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SanNewIgroups) UnmarshalBinary(b []byte) error {
	var res SanNewIgroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SanNewIgroupsIgroupsItems0 san new igroups igroups items0
//
// swagger:model SanNewIgroupsIgroupsItems0
type SanNewIgroupsIgroupsItems0 struct {

	// The name of an igroup to nest within a parent igroup. Mutually exclusive with initiators and initiator_objects.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The UUID of an igroup to nest within a parent igroup Usage: &lt;UUID&gt;
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this san new igroups igroups items0
func (m *SanNewIgroupsIgroupsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this san new igroups igroups items0 based on context it is used
func (m *SanNewIgroupsIgroupsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SanNewIgroupsIgroupsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SanNewIgroupsIgroupsItems0) UnmarshalBinary(b []byte) error {
	var res SanNewIgroupsIgroupsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SanNewIgroupsInitiatorObjectsItems0 san new igroups initiator objects items0
//
// swagger:model SanNewIgroupsInitiatorObjectsItems0
type SanNewIgroupsInitiatorObjectsItems0 struct {

	// A comment available for use by the administrator.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// The WWPN, IQN, or Alias of the initiator. Mutually exclusive with nested igroups and the initiators array.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this san new igroups initiator objects items0
func (m *SanNewIgroupsInitiatorObjectsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this san new igroups initiator objects items0 based on context it is used
func (m *SanNewIgroupsInitiatorObjectsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SanNewIgroupsInitiatorObjectsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SanNewIgroupsInitiatorObjectsItems0) UnmarshalBinary(b []byte) error {
	var res SanNewIgroupsInitiatorObjectsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
