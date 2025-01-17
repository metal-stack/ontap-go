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

// IgroupInitiatorListItem igroup initiator list item
//
// swagger:model igroup_initiator_list_item
type IgroupInitiatorListItem struct {

	// links
	Links *IgroupInitiatorListItemInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// A comment available for use by the administrator. Valid in POST and PATCH.
	//
	// Max Length: 254
	// Min Length: 0
	Comment *string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// connectivity tracking
	ConnectivityTracking *IgroupInitiatorListItemInlineConnectivityTracking `json:"connectivity_tracking,omitempty" yaml:"connectivity_tracking,omitempty"`

	// igroup
	Igroup *IgroupInitiatorListItemInlineIgroup `json:"igroup,omitempty" yaml:"igroup,omitempty"`

	// The FC WWPN, iSCSI IQN, or iSCSI EUI that identifies the host initiator. Valid in POST only and not allowed when the `records` property is used.<br/>
	// An FC WWPN consists of 16 hexadecimal digits grouped as 8 pairs separated by colons. The format for an iSCSI IQN is _iqn.yyyy-mm.reverse_domain_name:any_. The iSCSI EUI format consists of the _eui._ prefix followed by 16 hexadecimal characters.
	//
	// Example: iqn.1998-01.com.corp.iscsi:name1
	// Max Length: 96
	// Min Length: 1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// proximity
	Proximity *IgroupInitiatorListItemInlineProximity `json:"proximity,omitempty" yaml:"proximity,omitempty"`
}

// Validate validates this igroup initiator list item
func (m *IgroupInitiatorListItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectivityTracking(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIgroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProximity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItem) validateLinks(formats strfmt.Registry) error {
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

func (m *IgroupInitiatorListItem) validateComment(formats strfmt.Registry) error {
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

func (m *IgroupInitiatorListItem) validateConnectivityTracking(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectivityTracking) { // not required
		return nil
	}

	if m.ConnectivityTracking != nil {
		if err := m.ConnectivityTracking.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectivity_tracking")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectivity_tracking")
			}
			return err
		}
	}

	return nil
}

func (m *IgroupInitiatorListItem) validateIgroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Igroup) { // not required
		return nil
	}

	if m.Igroup != nil {
		if err := m.Igroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("igroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("igroup")
			}
			return err
		}
	}

	return nil
}

func (m *IgroupInitiatorListItem) validateName(formats strfmt.Registry) error {
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

func (m *IgroupInitiatorListItem) validateProximity(formats strfmt.Registry) error {
	if swag.IsZero(m.Proximity) { // not required
		return nil
	}

	if m.Proximity != nil {
		if err := m.Proximity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proximity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proximity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this igroup initiator list item based on the context it is used
func (m *IgroupInitiatorListItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectivityTracking(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIgroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProximity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IgroupInitiatorListItem) contextValidateConnectivityTracking(ctx context.Context, formats strfmt.Registry) error {

	if m.ConnectivityTracking != nil {

		if swag.IsZero(m.ConnectivityTracking) { // not required
			return nil
		}

		if err := m.ConnectivityTracking.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectivity_tracking")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectivity_tracking")
			}
			return err
		}
	}

	return nil
}

func (m *IgroupInitiatorListItem) contextValidateIgroup(ctx context.Context, formats strfmt.Registry) error {

	if m.Igroup != nil {

		if swag.IsZero(m.Igroup) { // not required
			return nil
		}

		if err := m.Igroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("igroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("igroup")
			}
			return err
		}
	}

	return nil
}

func (m *IgroupInitiatorListItem) contextValidateProximity(ctx context.Context, formats strfmt.Registry) error {

	if m.Proximity != nil {

		if swag.IsZero(m.Proximity) { // not required
			return nil
		}

		if err := m.Proximity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proximity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proximity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IgroupInitiatorListItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItem) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupInitiatorListItemInlineConnectivityTracking Overview of the initiator's connections to ONTAP.
//
// swagger:model igroup_initiator_list_item_inline_connectivity_tracking
type IgroupInitiatorListItemInlineConnectivityTracking struct {

	// Connection state.
	//
	// Read Only: true
	// Enum: ["full","none","partial","no_lun_maps"]
	ConnectionState *string `json:"connection_state,omitempty" yaml:"connection_state,omitempty"`
}

// Validate validates this igroup initiator list item inline connectivity tracking
func (m *IgroupInitiatorListItemInlineConnectivityTracking) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var igroupInitiatorListItemInlineConnectivityTrackingTypeConnectionStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["full","none","partial","no_lun_maps"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		igroupInitiatorListItemInlineConnectivityTrackingTypeConnectionStatePropEnum = append(igroupInitiatorListItemInlineConnectivityTrackingTypeConnectionStatePropEnum, v)
	}
}

const (

	// IgroupInitiatorListItemInlineConnectivityTrackingConnectionStateFull captures enum value "full"
	IgroupInitiatorListItemInlineConnectivityTrackingConnectionStateFull string = "full"

	// IgroupInitiatorListItemInlineConnectivityTrackingConnectionStateNone captures enum value "none"
	IgroupInitiatorListItemInlineConnectivityTrackingConnectionStateNone string = "none"

	// IgroupInitiatorListItemInlineConnectivityTrackingConnectionStatePartial captures enum value "partial"
	IgroupInitiatorListItemInlineConnectivityTrackingConnectionStatePartial string = "partial"

	// IgroupInitiatorListItemInlineConnectivityTrackingConnectionStateNoLunMaps captures enum value "no_lun_maps"
	IgroupInitiatorListItemInlineConnectivityTrackingConnectionStateNoLunMaps string = "no_lun_maps"
)

// prop value enum
func (m *IgroupInitiatorListItemInlineConnectivityTracking) validateConnectionStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, igroupInitiatorListItemInlineConnectivityTrackingTypeConnectionStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IgroupInitiatorListItemInlineConnectivityTracking) validateConnectionState(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectionState) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectionStateEnum("connectivity_tracking"+"."+"connection_state", "body", *m.ConnectionState); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this igroup initiator list item inline connectivity tracking based on the context it is used
func (m *IgroupInitiatorListItemInlineConnectivityTracking) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectionState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemInlineConnectivityTracking) contextValidateConnectionState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connectivity_tracking"+"."+"connection_state", "body", m.ConnectionState); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IgroupInitiatorListItemInlineConnectivityTracking) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItemInlineConnectivityTracking) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItemInlineConnectivityTracking
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupInitiatorListItemInlineIgroup The initiator group that directly owns the initiator, which is where modification of the initiator is supported. This property will only be populated when the initiator is a member of a nested initiator group.
//
// swagger:model igroup_initiator_list_item_inline_igroup
type IgroupInitiatorListItemInlineIgroup struct {

	// links
	Links *IgroupInitiatorListItemInlineIgroupInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the initiator group.
	//
	// Example: igroup1
	// Max Length: 96
	// Min Length: 1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// The unique identifier of the initiator group.
	//
	// Example: 4ea7a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this igroup initiator list item inline igroup
func (m *IgroupInitiatorListItemInlineIgroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
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

func (m *IgroupInitiatorListItemInlineIgroup) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("igroup" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("igroup" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *IgroupInitiatorListItemInlineIgroup) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("igroup"+"."+"name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("igroup"+"."+"name", "body", *m.Name, 96); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this igroup initiator list item inline igroup based on the context it is used
func (m *IgroupInitiatorListItemInlineIgroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemInlineIgroup) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("igroup" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("igroup" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IgroupInitiatorListItemInlineIgroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItemInlineIgroup) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItemInlineIgroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupInitiatorListItemInlineIgroupInlineLinks igroup initiator list item inline igroup inline links
//
// swagger:model igroup_initiator_list_item_inline_igroup_inline__links
type IgroupInitiatorListItemInlineIgroupInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this igroup initiator list item inline igroup inline links
func (m *IgroupInitiatorListItemInlineIgroupInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemInlineIgroupInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("igroup" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("igroup" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this igroup initiator list item inline igroup inline links based on the context it is used
func (m *IgroupInitiatorListItemInlineIgroupInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemInlineIgroupInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("igroup" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("igroup" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IgroupInitiatorListItemInlineIgroupInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItemInlineIgroupInlineLinks) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItemInlineIgroupInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupInitiatorListItemInlineLinks igroup initiator list item inline links
//
// swagger:model igroup_initiator_list_item_inline__links
type IgroupInitiatorListItemInlineLinks struct {

	// connectivity tracking
	ConnectivityTracking *IgroupInitiatorListItemInlineLinksInlineConnectivityTracking `json:"connectivity_tracking,omitempty" yaml:"connectivity_tracking,omitempty"`

	// self
	Self *IgroupInitiatorListItemInlineLinksInlineSelf `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this igroup initiator list item inline links
func (m *IgroupInitiatorListItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectivityTracking(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemInlineLinks) validateConnectivityTracking(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectivityTracking) { // not required
		return nil
	}

	if m.ConnectivityTracking != nil {
		if err := m.ConnectivityTracking.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "connectivity_tracking")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "connectivity_tracking")
			}
			return err
		}
	}

	return nil
}

func (m *IgroupInitiatorListItemInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this igroup initiator list item inline links based on the context it is used
func (m *IgroupInitiatorListItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectivityTracking(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemInlineLinks) contextValidateConnectivityTracking(ctx context.Context, formats strfmt.Registry) error {

	if m.ConnectivityTracking != nil {

		if swag.IsZero(m.ConnectivityTracking) { // not required
			return nil
		}

		if err := m.ConnectivityTracking.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "connectivity_tracking")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "connectivity_tracking")
			}
			return err
		}
	}

	return nil
}

func (m *IgroupInitiatorListItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IgroupInitiatorListItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupInitiatorListItemInlineLinksInlineConnectivityTracking A link to the initiator with connectivity information relevant to its membership of this initiator group.
//
// swagger:model igroup_initiator_list_item_inline__links_inline_connectivity_tracking
type IgroupInitiatorListItemInlineLinksInlineConnectivityTracking struct {

	// href
	// Example: /api/resourcelink
	// Read Only: true
	Href *string `json:"href,omitempty" yaml:"href,omitempty"`
}

// Validate validates this igroup initiator list item inline links inline connectivity tracking
func (m *IgroupInitiatorListItemInlineLinksInlineConnectivityTracking) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this igroup initiator list item inline links inline connectivity tracking based on the context it is used
func (m *IgroupInitiatorListItemInlineLinksInlineConnectivityTracking) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHref(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemInlineLinksInlineConnectivityTracking) contextValidateHref(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_links"+"."+"connectivity_tracking"+"."+"href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IgroupInitiatorListItemInlineLinksInlineConnectivityTracking) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItemInlineLinksInlineConnectivityTracking) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItemInlineLinksInlineConnectivityTracking
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupInitiatorListItemInlineLinksInlineSelf A link to the initiator where mutations can be made. If the initiator is inherited from a nested initiator group, the link refers to the initiator in the nested initiator group. In this case, mutations of the initiator will be applied to all initiator groups referencing the same nested initiator group.
//
// swagger:model igroup_initiator_list_item_inline__links_inline_self
type IgroupInitiatorListItemInlineLinksInlineSelf struct {

	// href
	// Example: /api/resourcelink
	// Read Only: true
	Href *string `json:"href,omitempty" yaml:"href,omitempty"`
}

// Validate validates this igroup initiator list item inline links inline self
func (m *IgroupInitiatorListItemInlineLinksInlineSelf) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this igroup initiator list item inline links inline self based on the context it is used
func (m *IgroupInitiatorListItemInlineLinksInlineSelf) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHref(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemInlineLinksInlineSelf) contextValidateHref(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_links"+"."+"self"+"."+"href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IgroupInitiatorListItemInlineLinksInlineSelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItemInlineLinksInlineSelf) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItemInlineLinksInlineSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupInitiatorListItemInlineProximity Properties that define to what SVMs the initiator is proximal. This information is used to properly report active optimized and active non-optimized network paths via ALUA. If no configuration has been specified for an initiator, the sub-object will not be present in GET.<br/>
// These properties can be set via initiator group POST and PATCH and apply to all instances of the initiator in all initiator groups in the SVM and its peers. The `proximity` sub-object for an initiator is set in POST and PATCH in its entirety and replaces any previously set proximity for the initiator within the SVM for the initiator within the SVM. The `local_svm` property must always be set to `true` or `false` when setting the `proximity` property. To clear any previously set proximity, POST or PATCH the `proximity` object to `null`.
//
// swagger:model igroup_initiator_list_item_inline_proximity
type IgroupInitiatorListItemInlineProximity struct {

	// A boolean that indicates if the initiator is proximal to the SVM of the containing initiator group. This is required for any POST or PATCH that includes the `proximity` sub-object.
	//
	LocalSvm *bool `json:"local_svm,omitempty" yaml:"local_svm,omitempty"`

	// An array of remote peer SVMs to which the initiator is proximal.
	//
	PeerSvms []*IgroupInitiatorListItemProximityPeerSvmsItems0 `json:"peer_svms,omitempty" yaml:"peer_svms,omitempty"`
}

// Validate validates this igroup initiator list item inline proximity
func (m *IgroupInitiatorListItemInlineProximity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeerSvms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemInlineProximity) validatePeerSvms(formats strfmt.Registry) error {
	if swag.IsZero(m.PeerSvms) { // not required
		return nil
	}

	for i := 0; i < len(m.PeerSvms); i++ {
		if swag.IsZero(m.PeerSvms[i]) { // not required
			continue
		}

		if m.PeerSvms[i] != nil {
			if err := m.PeerSvms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proximity" + "." + "peer_svms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("proximity" + "." + "peer_svms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this igroup initiator list item inline proximity based on the context it is used
func (m *IgroupInitiatorListItemInlineProximity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePeerSvms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemInlineProximity) contextValidatePeerSvms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PeerSvms); i++ {

		if m.PeerSvms[i] != nil {

			if swag.IsZero(m.PeerSvms[i]) { // not required
				return nil
			}

			if err := m.PeerSvms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proximity" + "." + "peer_svms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("proximity" + "." + "peer_svms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IgroupInitiatorListItemInlineProximity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItemInlineProximity) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItemInlineProximity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupInitiatorListItemProximityPeerSvmsItems0 A reference to an SVM peer relationship.
//
// swagger:model IgroupInitiatorListItemProximityPeerSvmsItems0
type IgroupInitiatorListItemProximityPeerSvmsItems0 struct {

	// links
	Links *IgroupInitiatorListItemProximityPeerSvmsItems0Links `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The local name of the peer SVM. This name is unique among all local and peer SVMs.
	//
	// Example: peer1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// The unique identifier of the SVM peer relationship. This is the UUID of the relationship, not the UUID of the peer SVM itself.
	//
	// Example: 4204cf77-4c82-9bdb-5644-b5a841c097a9
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this igroup initiator list item proximity peer svms items0
func (m *IgroupInitiatorListItemProximityPeerSvmsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemProximityPeerSvmsItems0) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this igroup initiator list item proximity peer svms items0 based on the context it is used
func (m *IgroupInitiatorListItemProximityPeerSvmsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemProximityPeerSvmsItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *IgroupInitiatorListItemProximityPeerSvmsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItemProximityPeerSvmsItems0) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItemProximityPeerSvmsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupInitiatorListItemProximityPeerSvmsItems0Links igroup initiator list item proximity peer svms items0 links
//
// swagger:model IgroupInitiatorListItemProximityPeerSvmsItems0Links
type IgroupInitiatorListItemProximityPeerSvmsItems0Links struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this igroup initiator list item proximity peer svms items0 links
func (m *IgroupInitiatorListItemProximityPeerSvmsItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemProximityPeerSvmsItems0Links) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this igroup initiator list item proximity peer svms items0 links based on the context it is used
func (m *IgroupInitiatorListItemProximityPeerSvmsItems0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemProximityPeerSvmsItems0Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IgroupInitiatorListItemProximityPeerSvmsItems0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItemProximityPeerSvmsItems0Links) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItemProximityPeerSvmsItems0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
