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

// PoolHealth pool health
//
// swagger:model pool_health
type PoolHealth struct {

	// Indicates whether the storage pool is able to participate in provisioning operations.
	// Read Only: true
	IsHealthy *bool `json:"is_healthy,omitempty" yaml:"is_healthy,omitempty"`

	// The state of the shared storage pool.
	// Read Only: true
	// Enum: ["normal","degraded","creating","deleting","reassigning","growing"]
	State *string `json:"state,omitempty" yaml:"state,omitempty"`

	// Indicates why the storage pool is unhealthy. This property is not returned for healthy storage pools.
	// Read Only: true
	UnhealthyReason *Error `json:"unhealthy_reason,omitempty" yaml:"unhealthy_reason,omitempty"`
}

// Validate validates this pool health
func (m *PoolHealth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnhealthyReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var poolHealthTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["normal","degraded","creating","deleting","reassigning","growing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		poolHealthTypeStatePropEnum = append(poolHealthTypeStatePropEnum, v)
	}
}

const (

	// PoolHealthStateNormal captures enum value "normal"
	PoolHealthStateNormal string = "normal"

	// PoolHealthStateDegraded captures enum value "degraded"
	PoolHealthStateDegraded string = "degraded"

	// PoolHealthStateCreating captures enum value "creating"
	PoolHealthStateCreating string = "creating"

	// PoolHealthStateDeleting captures enum value "deleting"
	PoolHealthStateDeleting string = "deleting"

	// PoolHealthStateReassigning captures enum value "reassigning"
	PoolHealthStateReassigning string = "reassigning"

	// PoolHealthStateGrowing captures enum value "growing"
	PoolHealthStateGrowing string = "growing"
)

// prop value enum
func (m *PoolHealth) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, poolHealthTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PoolHealth) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

func (m *PoolHealth) validateUnhealthyReason(formats strfmt.Registry) error {
	if swag.IsZero(m.UnhealthyReason) { // not required
		return nil
	}

	if m.UnhealthyReason != nil {
		if err := m.UnhealthyReason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unhealthy_reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unhealthy_reason")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pool health based on the context it is used
func (m *PoolHealth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIsHealthy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnhealthyReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolHealth) contextValidateIsHealthy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_healthy", "body", m.IsHealthy); err != nil {
		return err
	}

	return nil
}

func (m *PoolHealth) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *PoolHealth) contextValidateUnhealthyReason(ctx context.Context, formats strfmt.Registry) error {

	if m.UnhealthyReason != nil {

		if swag.IsZero(m.UnhealthyReason) { // not required
			return nil
		}

		if err := m.UnhealthyReason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unhealthy_reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unhealthy_reason")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolHealth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolHealth) UnmarshalBinary(b []byte) error {
	var res PoolHealth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}