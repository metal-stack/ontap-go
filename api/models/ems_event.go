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

// EmsEvent ems event
//
// swagger:model ems_event
type EmsEvent struct {

	// links
	Links *EmsEventInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// A list of parameters provided with the EMS event.
	// Read Only: true
	EmsEventInlineParameters []*EmsEventInlineParametersInlineArrayItem `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Index of the event. Returned by default.
	// Example: 1
	// Read Only: true
	Index *int64 `json:"index,omitempty" yaml:"index,omitempty"`

	// A formatted text string populated with parameter details. Returned by default.
	LogMessage *string `json:"log_message,omitempty" yaml:"log_message,omitempty"`

	// message
	Message *EmsEventInlineMessage `json:"message,omitempty" yaml:"message,omitempty"`

	// node
	Node *EmsEventInlineNode `json:"node,omitempty" yaml:"node,omitempty"`

	// Source
	// Read Only: true
	Source *string `json:"source,omitempty" yaml:"source,omitempty"`

	// Timestamp of the event. Returned by default.
	// Read Only: true
	// Format: date-time
	Time *strfmt.DateTime `json:"time,omitempty" yaml:"time,omitempty"`
}

// Validate validates this ems event
func (m *EmsEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmsEventInlineParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEvent) validateLinks(formats strfmt.Registry) error {
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

func (m *EmsEvent) validateEmsEventInlineParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.EmsEventInlineParameters) { // not required
		return nil
	}

	for i := 0; i < len(m.EmsEventInlineParameters); i++ {
		if swag.IsZero(m.EmsEventInlineParameters[i]) { // not required
			continue
		}

		if m.EmsEventInlineParameters[i] != nil {
			if err := m.EmsEventInlineParameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmsEvent) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.Message) { // not required
		return nil
	}

	if m.Message != nil {
		if err := m.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEvent) validateNode(formats strfmt.Registry) error {
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

func (m *EmsEvent) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ems event based on the context it is used
func (m *EmsEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmsEventInlineParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIndex(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEvent) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EmsEvent) contextValidateEmsEventInlineParameters(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "parameters", "body", []*EmsEventInlineParametersInlineArrayItem(m.EmsEventInlineParameters)); err != nil {
		return err
	}

	for i := 0; i < len(m.EmsEventInlineParameters); i++ {

		if m.EmsEventInlineParameters[i] != nil {

			if swag.IsZero(m.EmsEventInlineParameters[i]) { // not required
				return nil
			}

			if err := m.EmsEventInlineParameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmsEvent) contextValidateIndex(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *EmsEvent) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if m.Message != nil {

		if swag.IsZero(m.Message) { // not required
			return nil
		}

		if err := m.Message.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEvent) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EmsEvent) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *EmsEvent) contextValidateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "time", "body", m.Time); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEvent) UnmarshalBinary(b []byte) error {
	var res EmsEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventInlineLinks ems event inline links
//
// swagger:model ems_event_inline__links
type EmsEventInlineLinks struct {

	// self
	Self *EmsEventInlineLinksInlineSelf `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this ems event inline links
func (m *EmsEventInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ems event inline links based on the context it is used
func (m *EmsEventInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EmsEventInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventInlineLinks) UnmarshalBinary(b []byte) error {
	var res EmsEventInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventInlineLinksInlineSelf ems event inline links inline self
//
// swagger:model ems_event_inline__links_inline_self
type EmsEventInlineLinksInlineSelf struct {

	// href
	// Example: /api/resourcelink
	// Read Only: true
	Href *string `json:"href,omitempty" yaml:"href,omitempty"`
}

// Validate validates this ems event inline links inline self
func (m *EmsEventInlineLinksInlineSelf) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this ems event inline links inline self based on the context it is used
func (m *EmsEventInlineLinksInlineSelf) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHref(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventInlineLinksInlineSelf) contextValidateHref(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_links"+"."+"self"+"."+"href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventInlineLinksInlineSelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventInlineLinksInlineSelf) UnmarshalBinary(b []byte) error {
	var res EmsEventInlineLinksInlineSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventInlineMessage ems event inline message
//
// swagger:model ems_event_inline_message
type EmsEventInlineMessage struct {

	// links
	Links *EmsEventInlineMessageInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// Message name of the event. Returned by default.
	// Example: callhome.spares.low
	// Read Only: true
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// Severity of the event. Returned by default.
	// Example: emergency
	// Read Only: true
	// Enum: ["emergency","alert","error","notice","informational","debug"]
	Severity *string `json:"severity,omitempty" yaml:"severity,omitempty"`
}

// Validate validates this ems event inline message
func (m *EmsEventInlineMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventInlineMessage) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("message" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

var emsEventInlineMessageTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["emergency","alert","error","notice","informational","debug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emsEventInlineMessageTypeSeverityPropEnum = append(emsEventInlineMessageTypeSeverityPropEnum, v)
	}
}

const (

	// EmsEventInlineMessageSeverityEmergency captures enum value "emergency"
	EmsEventInlineMessageSeverityEmergency string = "emergency"

	// EmsEventInlineMessageSeverityAlert captures enum value "alert"
	EmsEventInlineMessageSeverityAlert string = "alert"

	// EmsEventInlineMessageSeverityError captures enum value "error"
	EmsEventInlineMessageSeverityError string = "error"

	// EmsEventInlineMessageSeverityNotice captures enum value "notice"
	EmsEventInlineMessageSeverityNotice string = "notice"

	// EmsEventInlineMessageSeverityInformational captures enum value "informational"
	EmsEventInlineMessageSeverityInformational string = "informational"

	// EmsEventInlineMessageSeverityDebug captures enum value "debug"
	EmsEventInlineMessageSeverityDebug string = "debug"
)

// prop value enum
func (m *EmsEventInlineMessage) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emsEventInlineMessageTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmsEventInlineMessage) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityEnum("message"+"."+"severity", "body", *m.Severity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ems event inline message based on the context it is used
func (m *EmsEventInlineMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeverity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventInlineMessage) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("message" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEventInlineMessage) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EmsEventInlineMessage) contextValidateSeverity(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message"+"."+"severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventInlineMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventInlineMessage) UnmarshalBinary(b []byte) error {
	var res EmsEventInlineMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventInlineMessageInlineLinks ems event inline message inline links
//
// swagger:model ems_event_inline_message_inline__links
type EmsEventInlineMessageInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this ems event inline message inline links
func (m *EmsEventInlineMessageInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventInlineMessageInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("message" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems event inline message inline links based on the context it is used
func (m *EmsEventInlineMessageInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventInlineMessageInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("message" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventInlineMessageInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventInlineMessageInlineLinks) UnmarshalBinary(b []byte) error {
	var res EmsEventInlineMessageInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventInlineNode ems event inline node
//
// swagger:model ems_event_inline_node
type EmsEventInlineNode struct {

	// links
	Links *EmsEventInlineNodeInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this ems event inline node
func (m *EmsEventInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventInlineNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this ems event inline node based on the context it is used
func (m *EmsEventInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EmsEventInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventInlineNode) UnmarshalBinary(b []byte) error {
	var res EmsEventInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventInlineNodeInlineLinks ems event inline node inline links
//
// swagger:model ems_event_inline_node_inline__links
type EmsEventInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this ems event inline node inline links
func (m *EmsEventInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ems event inline node inline links based on the context it is used
func (m *EmsEventInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EmsEventInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res EmsEventInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventInlineParametersInlineArrayItem ems event inline parameters inline array item
//
// swagger:model ems_event_inline_parameters_inline_array_item
type EmsEventInlineParametersInlineArrayItem struct {

	// Name of parameter
	// Example: numOps
	// Read Only: true
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// Value of parameter
	// Example: 123
	// Read Only: true
	Value *string `json:"value,omitempty" yaml:"value,omitempty"`
}

// Validate validates this ems event inline parameters inline array item
func (m *EmsEventInlineParametersInlineArrayItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this ems event inline parameters inline array item based on the context it is used
func (m *EmsEventInlineParametersInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventInlineParametersInlineArrayItem) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EmsEventInlineParametersInlineArrayItem) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventInlineParametersInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventInlineParametersInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res EmsEventInlineParametersInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
