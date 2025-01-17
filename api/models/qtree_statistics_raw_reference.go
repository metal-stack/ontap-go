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

// QtreeStatisticsRawReference These are raw IOPS and throughput performance numbers. These numbers are aggregated across all nodes in the cluster and increase with the uptime of the cluster.
//
// swagger:model qtree_statistics_raw_reference
type QtreeStatisticsRawReference struct {

	// iops raw
	IopsRaw *QtreeStatisticsRawReferenceInlineIopsRaw `json:"iops_raw,omitempty" yaml:"iops_raw,omitempty"`

	// Any errors associated with the sample. For example, if the aggregation of data over multiple nodes fails then any of the partial errors might be returned, "ok" on success, or "error" on any internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled with the next closest collection and tagged with "backfilled_data". "inconsistent_delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "negative_delta" is returned when an expected monotonically increasing value has decreased in value. "inconsistent_old_data" is returned when one or more nodes does not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: ["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]
	Status *string `json:"status,omitempty" yaml:"status,omitempty"`

	// throughput raw
	ThroughputRaw *QtreeStatisticsRawReferenceInlineThroughputRaw `json:"throughput_raw,omitempty" yaml:"throughput_raw,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25 11:20:13
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
}

// Validate validates this qtree statistics raw reference
func (m *QtreeStatisticsRawReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIopsRaw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughputRaw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QtreeStatisticsRawReference) validateIopsRaw(formats strfmt.Registry) error {
	if swag.IsZero(m.IopsRaw) { // not required
		return nil
	}

	if m.IopsRaw != nil {
		if err := m.IopsRaw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iops_raw")
			}
			return err
		}
	}

	return nil
}

var qtreeStatisticsRawReferenceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		qtreeStatisticsRawReferenceTypeStatusPropEnum = append(qtreeStatisticsRawReferenceTypeStatusPropEnum, v)
	}
}

const (

	// QtreeStatisticsRawReferenceStatusOk captures enum value "ok"
	QtreeStatisticsRawReferenceStatusOk string = "ok"

	// QtreeStatisticsRawReferenceStatusError captures enum value "error"
	QtreeStatisticsRawReferenceStatusError string = "error"

	// QtreeStatisticsRawReferenceStatusPartialNoData captures enum value "partial_no_data"
	QtreeStatisticsRawReferenceStatusPartialNoData string = "partial_no_data"

	// QtreeStatisticsRawReferenceStatusPartialNoUUID captures enum value "partial_no_uuid"
	QtreeStatisticsRawReferenceStatusPartialNoUUID string = "partial_no_uuid"

	// QtreeStatisticsRawReferenceStatusPartialNoResponse captures enum value "partial_no_response"
	QtreeStatisticsRawReferenceStatusPartialNoResponse string = "partial_no_response"

	// QtreeStatisticsRawReferenceStatusPartialOtherError captures enum value "partial_other_error"
	QtreeStatisticsRawReferenceStatusPartialOtherError string = "partial_other_error"

	// QtreeStatisticsRawReferenceStatusNegativeDelta captures enum value "negative_delta"
	QtreeStatisticsRawReferenceStatusNegativeDelta string = "negative_delta"

	// QtreeStatisticsRawReferenceStatusBackfilledData captures enum value "backfilled_data"
	QtreeStatisticsRawReferenceStatusBackfilledData string = "backfilled_data"

	// QtreeStatisticsRawReferenceStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	QtreeStatisticsRawReferenceStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// QtreeStatisticsRawReferenceStatusInconsistentOldData captures enum value "inconsistent_old_data"
	QtreeStatisticsRawReferenceStatusInconsistentOldData string = "inconsistent_old_data"
)

// prop value enum
func (m *QtreeStatisticsRawReference) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, qtreeStatisticsRawReferenceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *QtreeStatisticsRawReference) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *QtreeStatisticsRawReference) validateThroughputRaw(formats strfmt.Registry) error {
	if swag.IsZero(m.ThroughputRaw) { // not required
		return nil
	}

	if m.ThroughputRaw != nil {
		if err := m.ThroughputRaw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throughput_raw")
			}
			return err
		}
	}

	return nil
}

func (m *QtreeStatisticsRawReference) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this qtree statistics raw reference based on the context it is used
func (m *QtreeStatisticsRawReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIopsRaw(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughputRaw(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QtreeStatisticsRawReference) contextValidateIopsRaw(ctx context.Context, formats strfmt.Registry) error {

	if m.IopsRaw != nil {

		if swag.IsZero(m.IopsRaw) { // not required
			return nil
		}

		if err := m.IopsRaw.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iops_raw")
			}
			return err
		}
	}

	return nil
}

func (m *QtreeStatisticsRawReference) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *QtreeStatisticsRawReference) contextValidateThroughputRaw(ctx context.Context, formats strfmt.Registry) error {

	if m.ThroughputRaw != nil {

		if swag.IsZero(m.ThroughputRaw) { // not required
			return nil
		}

		if err := m.ThroughputRaw.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throughput_raw")
			}
			return err
		}
	}

	return nil
}

func (m *QtreeStatisticsRawReference) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QtreeStatisticsRawReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QtreeStatisticsRawReference) UnmarshalBinary(b []byte) error {
	var res QtreeStatisticsRawReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QtreeStatisticsRawReferenceInlineIopsRaw The number of I/O operations observed at the storage object. This should be used along with delta time to calculate the rate of I/O operations per unit of time.
//
// swagger:model qtree_statistics_raw_reference_inline_iops_raw
type QtreeStatisticsRawReferenceInlineIopsRaw struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other *int64 `json:"other,omitempty" yaml:"other,omitempty"`

	// Performance metric for read I/O operations.
	// Example: 200
	Read *int64 `json:"read,omitempty" yaml:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total *int64 `json:"total,omitempty" yaml:"total,omitempty"`

	// Peformance metric for write I/O operations.
	// Example: 100
	Write *int64 `json:"write,omitempty" yaml:"write,omitempty"`
}

// Validate validates this qtree statistics raw reference inline iops raw
func (m *QtreeStatisticsRawReferenceInlineIopsRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this qtree statistics raw reference inline iops raw based on the context it is used
func (m *QtreeStatisticsRawReferenceInlineIopsRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *QtreeStatisticsRawReferenceInlineIopsRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QtreeStatisticsRawReferenceInlineIopsRaw) UnmarshalBinary(b []byte) error {
	var res QtreeStatisticsRawReferenceInlineIopsRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QtreeStatisticsRawReferenceInlineThroughputRaw Throughput bytes observed at the storage object. This should be used along with delta time to calculate the rate of throughput bytes per unit of time.
//
// swagger:model qtree_statistics_raw_reference_inline_throughput_raw
type QtreeStatisticsRawReferenceInlineThroughputRaw struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other *int64 `json:"other,omitempty" yaml:"other,omitempty"`

	// Performance metric for read I/O operations.
	// Example: 200
	Read *int64 `json:"read,omitempty" yaml:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total *int64 `json:"total,omitempty" yaml:"total,omitempty"`

	// Peformance metric for write I/O operations.
	// Example: 100
	Write *int64 `json:"write,omitempty" yaml:"write,omitempty"`
}

// Validate validates this qtree statistics raw reference inline throughput raw
func (m *QtreeStatisticsRawReferenceInlineThroughputRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this qtree statistics raw reference inline throughput raw based on the context it is used
func (m *QtreeStatisticsRawReferenceInlineThroughputRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *QtreeStatisticsRawReferenceInlineThroughputRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QtreeStatisticsRawReferenceInlineThroughputRaw) UnmarshalBinary(b []byte) error {
	var res QtreeStatisticsRawReferenceInlineThroughputRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}