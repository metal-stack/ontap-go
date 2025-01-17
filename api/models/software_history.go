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

// SoftwareHistory software history
//
// swagger:model software_history
type SoftwareHistory struct {

	// Completion time of this installation request.
	// Example: 2019-02-02 20:00:00
	// Read Only: true
	// Format: date-time
	EndTime *strfmt.DateTime `json:"end_time,omitempty" yaml:"end_time,omitempty"`

	// Previous version of node
	// Example: ONTAP_X1
	// Read Only: true
	FromVersion *string `json:"from_version,omitempty" yaml:"from_version,omitempty"`

	// node
	Node *SoftwareHistoryInlineNode `json:"node,omitempty" yaml:"node,omitempty"`

	// Start time of this installation request.
	// Example: 2019-02-02 19:00:00
	// Read Only: true
	// Format: date-time
	StartTime *strfmt.DateTime `json:"start_time,omitempty" yaml:"start_time,omitempty"`

	// Status of this installation request.
	// Example: successful
	// Read Only: true
	// Enum: ["successful","canceled"]
	State *string `json:"state,omitempty" yaml:"state,omitempty"`

	// Updated version of node
	// Example: ONTAP_X2
	// Read Only: true
	ToVersion *string `json:"to_version,omitempty" yaml:"to_version,omitempty"`
}

// Validate validates this software history
func (m *SoftwareHistory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
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

func (m *SoftwareHistory) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareHistory) validateNode(formats strfmt.Registry) error {
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

func (m *SoftwareHistory) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var softwareHistoryTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["successful","canceled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		softwareHistoryTypeStatePropEnum = append(softwareHistoryTypeStatePropEnum, v)
	}
}

const (

	// SoftwareHistoryStateSuccessful captures enum value "successful"
	SoftwareHistoryStateSuccessful string = "successful"

	// SoftwareHistoryStateCanceled captures enum value "canceled"
	SoftwareHistoryStateCanceled string = "canceled"
)

// prop value enum
func (m *SoftwareHistory) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, softwareHistoryTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SoftwareHistory) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this software history based on the context it is used
func (m *SoftwareHistory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFromVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateToVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareHistory) contextValidateEndTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "end_time", "body", m.EndTime); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareHistory) contextValidateFromVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "from_version", "body", m.FromVersion); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareHistory) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SoftwareHistory) contextValidateStartTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "start_time", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareHistory) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareHistory) contextValidateToVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "to_version", "body", m.ToVersion); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareHistory) UnmarshalBinary(b []byte) error {
	var res SoftwareHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SoftwareHistoryInlineNode software history inline node
//
// swagger:model software_history_inline_node
type SoftwareHistoryInlineNode struct {

	// links
	Links *SoftwareHistoryInlineNodeInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this software history inline node
func (m *SoftwareHistoryInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareHistoryInlineNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this software history inline node based on the context it is used
func (m *SoftwareHistoryInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareHistoryInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SoftwareHistoryInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareHistoryInlineNode) UnmarshalBinary(b []byte) error {
	var res SoftwareHistoryInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SoftwareHistoryInlineNodeInlineLinks software history inline node inline links
//
// swagger:model software_history_inline_node_inline__links
type SoftwareHistoryInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this software history inline node inline links
func (m *SoftwareHistoryInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareHistoryInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this software history inline node inline links based on the context it is used
func (m *SoftwareHistoryInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareHistoryInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SoftwareHistoryInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareHistoryInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res SoftwareHistoryInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}