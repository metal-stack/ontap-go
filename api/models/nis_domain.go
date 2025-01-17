// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NisDomain The NIS domain to which this configuration belongs.
//
// swagger:model nis_domain
type NisDomain string

// Validate validates this nis domain
func (m NisDomain) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinLength("", "body", string(m), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("", "body", string(m), 64); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this nis domain based on context it is used
func (m NisDomain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
