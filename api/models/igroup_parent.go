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

// IgroupParent igroup parent
//
// swagger:model igroup_parent
type IgroupParent struct {

	// links
	Links *IgroupParentInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// A comment available for use by the administrator.
	//
	// Read Only: true
	// Max Length: 254
	// Min Length: 0
	Comment *string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// The initiator groups that contain this initiator group as a member.
	//
	// Read Only: true
	IgroupParentInlineParentIgroups []*IgroupParent `json:"parent_igroups,omitempty" yaml:"parent_igroups,omitempty"`

	// The name of the initiator group.
	//
	// Example: igroup1
	// Read Only: true
	// Max Length: 96
	// Min Length: 1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// The unique identifier of the initiator group.
	//
	// Example: 4ea7a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this igroup parent
func (m *IgroupParent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIgroupParentInlineParentIgroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupParent) validateLinks(formats strfmt.Registry) error {
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

func (m *IgroupParent) validateComment(formats strfmt.Registry) error {
	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if err := validate.MinLength("comment", "body", *m.Comment, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("comment", "body", *m.Comment, 254); err != nil {
		return err
	}

	return nil
}

func (m *IgroupParent) validateIgroupParentInlineParentIgroups(formats strfmt.Registry) error {
	if swag.IsZero(m.IgroupParentInlineParentIgroups) { // not required
		return nil
	}

	for i := 0; i < len(m.IgroupParentInlineParentIgroups); i++ {
		if swag.IsZero(m.IgroupParentInlineParentIgroups[i]) { // not required
			continue
		}

		if m.IgroupParentInlineParentIgroups[i] != nil {
			if err := m.IgroupParentInlineParentIgroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parent_igroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parent_igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IgroupParent) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 96); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this igroup parent based on the context it is used
func (m *IgroupParent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIgroupParentInlineParentIgroups(ctx, formats); err != nil {
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

func (m *IgroupParent) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IgroupParent) contextValidateComment(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "comment", "body", m.Comment); err != nil {
		return err
	}

	return nil
}

func (m *IgroupParent) contextValidateIgroupParentInlineParentIgroups(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "parent_igroups", "body", []*IgroupParent(m.IgroupParentInlineParentIgroups)); err != nil {
		return err
	}

	for i := 0; i < len(m.IgroupParentInlineParentIgroups); i++ {

		if m.IgroupParentInlineParentIgroups[i] != nil {

			if swag.IsZero(m.IgroupParentInlineParentIgroups[i]) { // not required
				return nil
			}

			if err := m.IgroupParentInlineParentIgroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parent_igroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parent_igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IgroupParent) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IgroupParent) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IgroupParent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupParent) UnmarshalBinary(b []byte) error {
	var res IgroupParent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupParentInlineLinks igroup parent inline links
//
// swagger:model igroup_parent_inline__links
type IgroupParentInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this igroup parent inline links
func (m *IgroupParentInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupParentInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this igroup parent inline links based on the context it is used
func (m *IgroupParentInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupParentInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IgroupParentInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupParentInlineLinks) UnmarshalBinary(b []byte) error {
	var res IgroupParentInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
