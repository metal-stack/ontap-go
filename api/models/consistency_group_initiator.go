// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsistencyGroupInitiator The initiators that are members of the initiator group.
//
// swagger:model consistency_group_initiator
type ConsistencyGroupInitiator struct {

	// A comment available for use by the administrator.
	//
	// Example: my comment
	// Max Length: 254
	// Min Length: 0
	Comment *string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// Name of initiator that is a member of the initiator group.
	//
	// Example: iqn.1998-01.com.corp.iscsi:name1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this consistency group initiator
func (m *ConsistencyGroupInitiator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupInitiator) validateComment(formats strfmt.Registry) error {
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

// ContextValidate validates this consistency group initiator based on context it is used
func (m *ConsistencyGroupInitiator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupInitiator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupInitiator) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupInitiator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
