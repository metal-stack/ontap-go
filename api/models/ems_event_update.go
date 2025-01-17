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

// EmsEventUpdate ems event update
//
// swagger:model ems_event_update
type EmsEventUpdate struct {

	// A formatted text string about the update.
	// Read Only: true
	LogMessage *string `json:"log_message,omitempty" yaml:"log_message,omitempty"`

	// State of the event instance when the update is raised.
	// Example: resolving
	// Read Only: true
	// Enum: ["opened","resolving","resolved","closed"]
	State *string `json:"state,omitempty" yaml:"state,omitempty"`

	// Timestamp of the update.
	// Read Only: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty" yaml:"update_time,omitempty"`
}

// Validate validates this ems event update
func (m *EmsEventUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var emsEventUpdateTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["opened","resolving","resolved","closed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emsEventUpdateTypeStatePropEnum = append(emsEventUpdateTypeStatePropEnum, v)
	}
}

const (

	// EmsEventUpdateStateOpened captures enum value "opened"
	EmsEventUpdateStateOpened string = "opened"

	// EmsEventUpdateStateResolving captures enum value "resolving"
	EmsEventUpdateStateResolving string = "resolving"

	// EmsEventUpdateStateResolved captures enum value "resolved"
	EmsEventUpdateStateResolved string = "resolved"

	// EmsEventUpdateStateClosed captures enum value "closed"
	EmsEventUpdateStateClosed string = "closed"
)

// prop value enum
func (m *EmsEventUpdate) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emsEventUpdateTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmsEventUpdate) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

func (m *EmsEventUpdate) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ems event update based on the context it is used
func (m *EmsEventUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventUpdate) contextValidateLogMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "log_message", "body", m.LogMessage); err != nil {
		return err
	}

	return nil
}

func (m *EmsEventUpdate) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *EmsEventUpdate) contextValidateUpdateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "update_time", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventUpdate) UnmarshalBinary(b []byte) error {
	var res EmsEventUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}