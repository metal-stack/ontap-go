// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// S3ServerWarning s3 server warning
//
// swagger:model s3_server_warning
type S3ServerWarning struct {

	// Warning code of the warning encountered.
	Code *int64 `json:"code,omitempty" yaml:"code,omitempty"`

	// Details of the warning sent from the S3 server.
	Message *string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Validate validates this s3 server warning
func (m *S3ServerWarning) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this s3 server warning based on context it is used
func (m *S3ServerWarning) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *S3ServerWarning) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3ServerWarning) UnmarshalBinary(b []byte) error {
	var res S3ServerWarning
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
