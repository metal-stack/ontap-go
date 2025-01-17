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

// Chassis chassis
//
// swagger:model chassis
type Chassis struct {

	// List of FRUs in the chassis.
	ChassisInlineFrus []*ChassisInlineFrusInlineArrayItem `json:"frus,omitempty" yaml:"frus,omitempty"`

	// List of nodes in the chassis.
	ChassisInlineNodes []*ChassisInlineNodesInlineArrayItem `json:"nodes,omitempty" yaml:"nodes,omitempty"`

	// List of shelves in chassis.
	ChassisInlineShelves []*ShelfReference `json:"shelves,omitempty" yaml:"shelves,omitempty"`

	// id
	// Example: 21352005981
	ID *string `json:"id,omitempty" yaml:"id,omitempty"`

	// state
	// Enum: ["ok","error"]
	State *string `json:"state,omitempty" yaml:"state,omitempty"`
}

// Validate validates this chassis
func (m *Chassis) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChassisInlineFrus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChassisInlineNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChassisInlineShelves(formats); err != nil {
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

func (m *Chassis) validateChassisInlineFrus(formats strfmt.Registry) error {
	if swag.IsZero(m.ChassisInlineFrus) { // not required
		return nil
	}

	for i := 0; i < len(m.ChassisInlineFrus); i++ {
		if swag.IsZero(m.ChassisInlineFrus[i]) { // not required
			continue
		}

		if m.ChassisInlineFrus[i] != nil {
			if err := m.ChassisInlineFrus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("frus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("frus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Chassis) validateChassisInlineNodes(formats strfmt.Registry) error {
	if swag.IsZero(m.ChassisInlineNodes) { // not required
		return nil
	}

	for i := 0; i < len(m.ChassisInlineNodes); i++ {
		if swag.IsZero(m.ChassisInlineNodes[i]) { // not required
			continue
		}

		if m.ChassisInlineNodes[i] != nil {
			if err := m.ChassisInlineNodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Chassis) validateChassisInlineShelves(formats strfmt.Registry) error {
	if swag.IsZero(m.ChassisInlineShelves) { // not required
		return nil
	}

	for i := 0; i < len(m.ChassisInlineShelves); i++ {
		if swag.IsZero(m.ChassisInlineShelves[i]) { // not required
			continue
		}

		if m.ChassisInlineShelves[i] != nil {
			if err := m.ChassisInlineShelves[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shelves" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shelves" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var chassisTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		chassisTypeStatePropEnum = append(chassisTypeStatePropEnum, v)
	}
}

const (

	// ChassisStateOk captures enum value "ok"
	ChassisStateOk string = "ok"

	// ChassisStateError captures enum value "error"
	ChassisStateError string = "error"
)

// prop value enum
func (m *Chassis) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, chassisTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Chassis) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this chassis based on the context it is used
func (m *Chassis) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChassisInlineFrus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChassisInlineNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChassisInlineShelves(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Chassis) contextValidateChassisInlineFrus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChassisInlineFrus); i++ {

		if m.ChassisInlineFrus[i] != nil {

			if swag.IsZero(m.ChassisInlineFrus[i]) { // not required
				return nil
			}

			if err := m.ChassisInlineFrus[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("frus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("frus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Chassis) contextValidateChassisInlineNodes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChassisInlineNodes); i++ {

		if m.ChassisInlineNodes[i] != nil {

			if swag.IsZero(m.ChassisInlineNodes[i]) { // not required
				return nil
			}

			if err := m.ChassisInlineNodes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Chassis) contextValidateChassisInlineShelves(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChassisInlineShelves); i++ {

		if m.ChassisInlineShelves[i] != nil {

			if swag.IsZero(m.ChassisInlineShelves[i]) { // not required
				return nil
			}

			if err := m.ChassisInlineShelves[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shelves" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shelves" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Chassis) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Chassis) UnmarshalBinary(b []byte) error {
	var res Chassis
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ChassisInlineFrusInlineArrayItem chassis inline frus inline array item
//
// swagger:model chassis_inline_frus_inline_array_item
type ChassisInlineFrusInlineArrayItem struct {

	// id
	ID *string `json:"id,omitempty" yaml:"id,omitempty"`

	// state
	// Enum: ["ok","error"]
	State *string `json:"state,omitempty" yaml:"state,omitempty"`

	// type
	// Enum: ["fan","psu"]
	Type *string `json:"type,omitempty" yaml:"type,omitempty"`
}

// Validate validates this chassis inline frus inline array item
func (m *ChassisInlineFrusInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var chassisInlineFrusInlineArrayItemTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		chassisInlineFrusInlineArrayItemTypeStatePropEnum = append(chassisInlineFrusInlineArrayItemTypeStatePropEnum, v)
	}
}

const (

	// ChassisInlineFrusInlineArrayItemStateOk captures enum value "ok"
	ChassisInlineFrusInlineArrayItemStateOk string = "ok"

	// ChassisInlineFrusInlineArrayItemStateError captures enum value "error"
	ChassisInlineFrusInlineArrayItemStateError string = "error"
)

// prop value enum
func (m *ChassisInlineFrusInlineArrayItem) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, chassisInlineFrusInlineArrayItemTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ChassisInlineFrusInlineArrayItem) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

var chassisInlineFrusInlineArrayItemTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fan","psu"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		chassisInlineFrusInlineArrayItemTypeTypePropEnum = append(chassisInlineFrusInlineArrayItemTypeTypePropEnum, v)
	}
}

const (

	// ChassisInlineFrusInlineArrayItemTypeFan captures enum value "fan"
	ChassisInlineFrusInlineArrayItemTypeFan string = "fan"

	// ChassisInlineFrusInlineArrayItemTypePsu captures enum value "psu"
	ChassisInlineFrusInlineArrayItemTypePsu string = "psu"
)

// prop value enum
func (m *ChassisInlineFrusInlineArrayItem) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, chassisInlineFrusInlineArrayItemTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ChassisInlineFrusInlineArrayItem) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this chassis inline frus inline array item based on context it is used
func (m *ChassisInlineFrusInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChassisInlineFrusInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisInlineFrusInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res ChassisInlineFrusInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ChassisInlineNodesInlineArrayItem List of nodes in chassis.
//
// swagger:model chassis_inline_nodes_inline_array_item
type ChassisInlineNodesInlineArrayItem struct {

	// links
	Links *ChassisInlineNodesInlineArrayItemInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// pcis
	Pcis *ChassisInlineNodesInlineArrayItemInlinePcis `json:"pcis,omitempty" yaml:"pcis,omitempty"`

	// The position of the node in the chassis, when viewed from the rear of the system.
	// Example: top
	// Enum: ["top","bottom","left","right","centre","unknown"]
	Position *string `json:"position,omitempty" yaml:"position,omitempty"`

	// usbs
	Usbs *ChassisInlineNodesInlineArrayItemInlineUsbs `json:"usbs,omitempty" yaml:"usbs,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this chassis inline nodes inline array item
func (m *ChassisInlineNodesInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePcis(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsbs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisInlineNodesInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

func (m *ChassisInlineNodesInlineArrayItem) validatePcis(formats strfmt.Registry) error {
	if swag.IsZero(m.Pcis) { // not required
		return nil
	}

	if m.Pcis != nil {
		if err := m.Pcis.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pcis")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pcis")
			}
			return err
		}
	}

	return nil
}

var chassisInlineNodesInlineArrayItemTypePositionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["top","bottom","left","right","centre","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		chassisInlineNodesInlineArrayItemTypePositionPropEnum = append(chassisInlineNodesInlineArrayItemTypePositionPropEnum, v)
	}
}

const (

	// ChassisInlineNodesInlineArrayItemPositionTop captures enum value "top"
	ChassisInlineNodesInlineArrayItemPositionTop string = "top"

	// ChassisInlineNodesInlineArrayItemPositionBottom captures enum value "bottom"
	ChassisInlineNodesInlineArrayItemPositionBottom string = "bottom"

	// ChassisInlineNodesInlineArrayItemPositionLeft captures enum value "left"
	ChassisInlineNodesInlineArrayItemPositionLeft string = "left"

	// ChassisInlineNodesInlineArrayItemPositionRight captures enum value "right"
	ChassisInlineNodesInlineArrayItemPositionRight string = "right"

	// ChassisInlineNodesInlineArrayItemPositionCentre captures enum value "centre"
	ChassisInlineNodesInlineArrayItemPositionCentre string = "centre"

	// ChassisInlineNodesInlineArrayItemPositionUnknown captures enum value "unknown"
	ChassisInlineNodesInlineArrayItemPositionUnknown string = "unknown"
)

// prop value enum
func (m *ChassisInlineNodesInlineArrayItem) validatePositionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, chassisInlineNodesInlineArrayItemTypePositionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ChassisInlineNodesInlineArrayItem) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	// value enum
	if err := m.validatePositionEnum("position", "body", *m.Position); err != nil {
		return err
	}

	return nil
}

func (m *ChassisInlineNodesInlineArrayItem) validateUsbs(formats strfmt.Registry) error {
	if swag.IsZero(m.Usbs) { // not required
		return nil
	}

	if m.Usbs != nil {
		if err := m.Usbs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usbs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usbs")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this chassis inline nodes inline array item based on the context it is used
func (m *ChassisInlineNodesInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePcis(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsbs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisInlineNodesInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ChassisInlineNodesInlineArrayItem) contextValidatePcis(ctx context.Context, formats strfmt.Registry) error {

	if m.Pcis != nil {

		if swag.IsZero(m.Pcis) { // not required
			return nil
		}

		if err := m.Pcis.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pcis")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pcis")
			}
			return err
		}
	}

	return nil
}

func (m *ChassisInlineNodesInlineArrayItem) contextValidateUsbs(ctx context.Context, formats strfmt.Registry) error {

	if m.Usbs != nil {

		if swag.IsZero(m.Usbs) { // not required
			return nil
		}

		if err := m.Usbs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usbs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usbs")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChassisInlineNodesInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisInlineNodesInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res ChassisInlineNodesInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ChassisInlineNodesInlineArrayItemInlineLinks chassis inline nodes inline array item inline links
//
// swagger:model chassis_inline_nodes_inline_array_item_inline__links
type ChassisInlineNodesInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this chassis inline nodes inline array item inline links
func (m *ChassisInlineNodesInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisInlineNodesInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this chassis inline nodes inline array item inline links based on the context it is used
func (m *ChassisInlineNodesInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisInlineNodesInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ChassisInlineNodesInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisInlineNodesInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res ChassisInlineNodesInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ChassisInlineNodesInlineArrayItemInlinePcis chassis inline nodes inline array item inline pcis
//
// swagger:model chassis_inline_nodes_inline_array_item_inline_pcis
type ChassisInlineNodesInlineArrayItemInlinePcis struct {

	// cards
	Cards []*ChassisNodesItems0PcisCardsItems0 `json:"cards" yaml:"cards"`
}

// Validate validates this chassis inline nodes inline array item inline pcis
func (m *ChassisInlineNodesInlineArrayItemInlinePcis) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCards(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisInlineNodesInlineArrayItemInlinePcis) validateCards(formats strfmt.Registry) error {
	if swag.IsZero(m.Cards) { // not required
		return nil
	}

	for i := 0; i < len(m.Cards); i++ {
		if swag.IsZero(m.Cards[i]) { // not required
			continue
		}

		if m.Cards[i] != nil {
			if err := m.Cards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pcis" + "." + "cards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pcis" + "." + "cards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this chassis inline nodes inline array item inline pcis based on the context it is used
func (m *ChassisInlineNodesInlineArrayItemInlinePcis) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCards(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisInlineNodesInlineArrayItemInlinePcis) contextValidateCards(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Cards); i++ {

		if m.Cards[i] != nil {

			if swag.IsZero(m.Cards[i]) { // not required
				return nil
			}

			if err := m.Cards[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pcis" + "." + "cards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pcis" + "." + "cards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChassisInlineNodesInlineArrayItemInlinePcis) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisInlineNodesInlineArrayItemInlinePcis) UnmarshalBinary(b []byte) error {
	var res ChassisInlineNodesInlineArrayItemInlinePcis
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ChassisNodesItems0PcisCardsItems0 chassis nodes items0 pcis cards items0
//
// swagger:model ChassisNodesItems0PcisCardsItems0
type ChassisNodesItems0PcisCardsItems0 struct {

	// The description of the PCI card.
	// Example: Intel Lewisburg series chipset SATA Controller
	Device *string `json:"device,omitempty" yaml:"device,omitempty"`

	// The info string from the device driver of the PCI card.
	// Example: Additional Info: 0 (0xaaf00000)   SHM2S86Q120GLM22NP FW1146 114473MB 512B/sect (SPG190108GW)
	Info *string `json:"info,omitempty" yaml:"info,omitempty"`

	// The slot where the PCI card is placed. This can sometimes take the form of "6-1" to indicate slot and subslot.
	// Example: 0
	Slot *string `json:"slot,omitempty" yaml:"slot,omitempty"`
}

// Validate validates this chassis nodes items0 pcis cards items0
func (m *ChassisNodesItems0PcisCardsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this chassis nodes items0 pcis cards items0 based on context it is used
func (m *ChassisNodesItems0PcisCardsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChassisNodesItems0PcisCardsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisNodesItems0PcisCardsItems0) UnmarshalBinary(b []byte) error {
	var res ChassisNodesItems0PcisCardsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ChassisInlineNodesInlineArrayItemInlineUsbs The status of the USB ports on the controller.
//
// swagger:model chassis_inline_nodes_inline_array_item_inline_usbs
type ChassisInlineNodesInlineArrayItemInlineUsbs struct {

	// Indicates whether or not the USB ports are enabled.
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// ports
	Ports []*ChassisNodesItems0UsbsPortsItems0 `json:"ports" yaml:"ports"`

	// Indicates whether or not USB ports are supported on the current platform.
	Supported *bool `json:"supported,omitempty" yaml:"supported,omitempty"`
}

// Validate validates this chassis inline nodes inline array item inline usbs
func (m *ChassisInlineNodesInlineArrayItemInlineUsbs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisInlineNodesInlineArrayItemInlineUsbs) validatePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usbs" + "." + "ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usbs" + "." + "ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this chassis inline nodes inline array item inline usbs based on the context it is used
func (m *ChassisInlineNodesInlineArrayItemInlineUsbs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisInlineNodesInlineArrayItemInlineUsbs) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ports); i++ {

		if m.Ports[i] != nil {

			if swag.IsZero(m.Ports[i]) { // not required
				return nil
			}

			if err := m.Ports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usbs" + "." + "ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usbs" + "." + "ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChassisInlineNodesInlineArrayItemInlineUsbs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisInlineNodesInlineArrayItemInlineUsbs) UnmarshalBinary(b []byte) error {
	var res ChassisInlineNodesInlineArrayItemInlineUsbs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ChassisNodesItems0UsbsPortsItems0 chassis nodes items0 usbs ports items0
//
// swagger:model ChassisNodesItems0UsbsPortsItems0
type ChassisNodesItems0UsbsPortsItems0 struct {

	// Indicates whether or not the USB port has a device connected to it.
	Connected *bool `json:"connected,omitempty" yaml:"connected,omitempty"`
}

// Validate validates this chassis nodes items0 usbs ports items0
func (m *ChassisNodesItems0UsbsPortsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this chassis nodes items0 usbs ports items0 based on context it is used
func (m *ChassisNodesItems0UsbsPortsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChassisNodesItems0UsbsPortsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisNodesItems0UsbsPortsItems0) UnmarshalBinary(b []byte) error {
	var res ChassisNodesItems0UsbsPortsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}