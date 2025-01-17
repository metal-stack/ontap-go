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

// PerformanceMetricRawSvmReference These are raw performance numbers, such as IOPS latency and throughput for SVM protocols. These numbers are aggregated across all nodes in the cluster and increase with the uptime of the cluster.
//
// swagger:model performance_metric_raw_svm_reference
type PerformanceMetricRawSvmReference struct {

	// iops raw
	IopsRaw *PerformanceMetricRawSvmReferenceInlineIopsRaw `json:"iops_raw,omitempty" yaml:"iops_raw,omitempty"`

	// latency raw
	LatencyRaw *PerformanceMetricRawSvmReferenceInlineLatencyRaw `json:"latency_raw,omitempty" yaml:"latency_raw,omitempty"`

	// Any errors associated with the sample. For example, if the aggregation of data over multiple nodes fails then any of the partial errors might be returned, "ok" on success, or "error" on any internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "Inconsistent_delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "Inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: ["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]
	Status *string `json:"status,omitempty" yaml:"status,omitempty"`

	// throughput raw
	ThroughputRaw *PerformanceMetricRawSvmReferenceInlineThroughputRaw `json:"throughput_raw,omitempty" yaml:"throughput_raw,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25 11:20:13
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
}

// Validate validates this performance metric raw svm reference
func (m *PerformanceMetricRawSvmReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIopsRaw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatencyRaw(formats); err != nil {
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

func (m *PerformanceMetricRawSvmReference) validateIopsRaw(formats strfmt.Registry) error {
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

func (m *PerformanceMetricRawSvmReference) validateLatencyRaw(formats strfmt.Registry) error {
	if swag.IsZero(m.LatencyRaw) { // not required
		return nil
	}

	if m.LatencyRaw != nil {
		if err := m.LatencyRaw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latency_raw")
			}
			return err
		}
	}

	return nil
}

var performanceMetricRawSvmReferenceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceMetricRawSvmReferenceTypeStatusPropEnum = append(performanceMetricRawSvmReferenceTypeStatusPropEnum, v)
	}
}

const (

	// PerformanceMetricRawSvmReferenceStatusOk captures enum value "ok"
	PerformanceMetricRawSvmReferenceStatusOk string = "ok"

	// PerformanceMetricRawSvmReferenceStatusError captures enum value "error"
	PerformanceMetricRawSvmReferenceStatusError string = "error"

	// PerformanceMetricRawSvmReferenceStatusPartialNoData captures enum value "partial_no_data"
	PerformanceMetricRawSvmReferenceStatusPartialNoData string = "partial_no_data"

	// PerformanceMetricRawSvmReferenceStatusPartialNoResponse captures enum value "partial_no_response"
	PerformanceMetricRawSvmReferenceStatusPartialNoResponse string = "partial_no_response"

	// PerformanceMetricRawSvmReferenceStatusPartialOtherError captures enum value "partial_other_error"
	PerformanceMetricRawSvmReferenceStatusPartialOtherError string = "partial_other_error"

	// PerformanceMetricRawSvmReferenceStatusNegativeDelta captures enum value "negative_delta"
	PerformanceMetricRawSvmReferenceStatusNegativeDelta string = "negative_delta"

	// PerformanceMetricRawSvmReferenceStatusNotFound captures enum value "not_found"
	PerformanceMetricRawSvmReferenceStatusNotFound string = "not_found"

	// PerformanceMetricRawSvmReferenceStatusBackfilledData captures enum value "backfilled_data"
	PerformanceMetricRawSvmReferenceStatusBackfilledData string = "backfilled_data"

	// PerformanceMetricRawSvmReferenceStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PerformanceMetricRawSvmReferenceStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// PerformanceMetricRawSvmReferenceStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PerformanceMetricRawSvmReferenceStatusInconsistentOldData string = "inconsistent_old_data"

	// PerformanceMetricRawSvmReferenceStatusPartialNoUUID captures enum value "partial_no_uuid"
	PerformanceMetricRawSvmReferenceStatusPartialNoUUID string = "partial_no_uuid"
)

// prop value enum
func (m *PerformanceMetricRawSvmReference) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceMetricRawSvmReferenceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceMetricRawSvmReference) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceMetricRawSvmReference) validateThroughputRaw(formats strfmt.Registry) error {
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

func (m *PerformanceMetricRawSvmReference) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this performance metric raw svm reference based on the context it is used
func (m *PerformanceMetricRawSvmReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIopsRaw(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatencyRaw(ctx, formats); err != nil {
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

func (m *PerformanceMetricRawSvmReference) contextValidateIopsRaw(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceMetricRawSvmReference) contextValidateLatencyRaw(ctx context.Context, formats strfmt.Registry) error {

	if m.LatencyRaw != nil {

		if swag.IsZero(m.LatencyRaw) { // not required
			return nil
		}

		if err := m.LatencyRaw.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency_raw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latency_raw")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceMetricRawSvmReference) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceMetricRawSvmReference) contextValidateThroughputRaw(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceMetricRawSvmReference) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricRawSvmReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricRawSvmReference) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricRawSvmReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricRawSvmReferenceInlineIopsRaw The number of I/O operations observed at the storage object. This should be used along with delta time to calculate the rate of I/O operations per unit of time.
//
// swagger:model performance_metric_raw_svm_reference_inline_iops_raw
type PerformanceMetricRawSvmReferenceInlineIopsRaw struct {

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

// Validate validates this performance metric raw svm reference inline iops raw
func (m *PerformanceMetricRawSvmReferenceInlineIopsRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance metric raw svm reference inline iops raw based on the context it is used
func (m *PerformanceMetricRawSvmReferenceInlineIopsRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricRawSvmReferenceInlineIopsRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricRawSvmReferenceInlineIopsRaw) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricRawSvmReferenceInlineIopsRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricRawSvmReferenceInlineLatencyRaw The raw latency in microseconds observed at the storage object. This should be divided by the raw IOPS value to calculate the average latency per I/O operation.
//
// swagger:model performance_metric_raw_svm_reference_inline_latency_raw
type PerformanceMetricRawSvmReferenceInlineLatencyRaw struct {

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

// Validate validates this performance metric raw svm reference inline latency raw
func (m *PerformanceMetricRawSvmReferenceInlineLatencyRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance metric raw svm reference inline latency raw based on the context it is used
func (m *PerformanceMetricRawSvmReferenceInlineLatencyRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricRawSvmReferenceInlineLatencyRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricRawSvmReferenceInlineLatencyRaw) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricRawSvmReferenceInlineLatencyRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceMetricRawSvmReferenceInlineThroughputRaw Throughput bytes observed at the storage object. This should be used along with delta time to calculate the rate of throughput bytes per unit of time.
//
// swagger:model performance_metric_raw_svm_reference_inline_throughput_raw
type PerformanceMetricRawSvmReferenceInlineThroughputRaw struct {

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

// Validate validates this performance metric raw svm reference inline throughput raw
func (m *PerformanceMetricRawSvmReferenceInlineThroughputRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance metric raw svm reference inline throughput raw based on the context it is used
func (m *PerformanceMetricRawSvmReferenceInlineThroughputRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceMetricRawSvmReferenceInlineThroughputRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceMetricRawSvmReferenceInlineThroughputRaw) UnmarshalBinary(b []byte) error {
	var res PerformanceMetricRawSvmReferenceInlineThroughputRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
