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

// Log log
//
// swagger:model log
type Log struct {

	// links
	Links *LogInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The format in which the logs are generated by consolidation process.
	//   Possible values are:
	//   * xml  - Data ONTAP-specific XML log format
	//   * evtx - Microsoft Windows EVTX log format
	//
	// Enum: ["xml","evtx"]
	Format *string `json:"format,omitempty" yaml:"format,omitempty"`

	// retention
	Retention *LogInlineRetention `json:"retention,omitempty" yaml:"retention,omitempty"`

	// rotation
	Rotation *Rotation `json:"rotation,omitempty" yaml:"rotation,omitempty"`
}

// Validate validates this log
func (m *Log) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetention(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRotation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Log) validateLinks(formats strfmt.Registry) error {
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

var logTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["xml","evtx"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		logTypeFormatPropEnum = append(logTypeFormatPropEnum, v)
	}
}

const (

	// LogFormatXML captures enum value "xml"
	LogFormatXML string = "xml"

	// LogFormatEvtx captures enum value "evtx"
	LogFormatEvtx string = "evtx"
)

// prop value enum
func (m *Log) validateFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, logTypeFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Log) validateFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", *m.Format); err != nil {
		return err
	}

	return nil
}

func (m *Log) validateRetention(formats strfmt.Registry) error {
	if swag.IsZero(m.Retention) { // not required
		return nil
	}

	if m.Retention != nil {
		if err := m.Retention.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retention")
			}
			return err
		}
	}

	return nil
}

func (m *Log) validateRotation(formats strfmt.Registry) error {
	if swag.IsZero(m.Rotation) { // not required
		return nil
	}

	if m.Rotation != nil {
		if err := m.Rotation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rotation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rotation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this log based on the context it is used
func (m *Log) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRetention(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRotation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Log) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Log) contextValidateRetention(ctx context.Context, formats strfmt.Registry) error {

	if m.Retention != nil {

		if swag.IsZero(m.Retention) { // not required
			return nil
		}

		if err := m.Retention.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retention")
			}
			return err
		}
	}

	return nil
}

func (m *Log) contextValidateRotation(ctx context.Context, formats strfmt.Registry) error {

	if m.Rotation != nil {

		if swag.IsZero(m.Rotation) { // not required
			return nil
		}

		if err := m.Rotation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rotation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rotation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Log) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Log) UnmarshalBinary(b []byte) error {
	var res Log
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LogInlineLinks log inline links
//
// swagger:model log_inline__links
type LogInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this log inline links
func (m *LogInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this log inline links based on the context it is used
func (m *LogInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *LogInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogInlineLinks) UnmarshalBinary(b []byte) error {
	var res LogInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LogInlineRetention log inline retention
//
// swagger:model log_inline_retention
type LogInlineRetention struct {

	// Determines how many audit log files to retain before
	// rotating the oldest log file out. This is mutually exclusive with
	// duration.
	//
	Count *int64 `json:"count,omitempty" yaml:"count,omitempty"`

	// Specifies an ISO-8601 format date and time to retain the audit log file. The audit log files are
	// deleted once they reach the specified date/time. This is mutually exclusive with count.
	//
	// Example: P4DT12H30M5S
	Duration *string `json:"duration,omitempty" yaml:"duration,omitempty"`
}

// Validate validates this log inline retention
func (m *LogInlineRetention) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this log inline retention based on context it is used
func (m *LogInlineRetention) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogInlineRetention) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogInlineRetention) UnmarshalBinary(b []byte) error {
	var res LogInlineRetention
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
