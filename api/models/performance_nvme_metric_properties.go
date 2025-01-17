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

// PerformanceNvmeMetricProperties Performance numbers, such as IOPS latency and throughput, for SVM protocols.
//
// swagger:model performance_nvme_metric_properties
type PerformanceNvmeMetricProperties struct {

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Read Only: true
	// Enum: ["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]
	Duration *string `json:"duration,omitempty" yaml:"duration,omitempty"`

	// iops
	Iops *PerformanceNvmeMetricPropertiesInlineIops `json:"iops,omitempty" yaml:"iops,omitempty"`

	// latency
	Latency *PerformanceNvmeMetricPropertiesInlineLatency `json:"latency,omitempty" yaml:"latency,omitempty"`

	// Any errors associated with the sample. For example, if the aggregation of data over multiple nodes fails then any of the partial errors might be returned, "ok" on success, or "error" on any internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "Inconsistent_ delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "Inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: ["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]
	Status *string `json:"status,omitempty" yaml:"status,omitempty"`

	// throughput
	Throughput *PerformanceNvmeMetricPropertiesInlineThroughput `json:"throughput,omitempty" yaml:"throughput,omitempty"`
}

// Validate validates this performance nvme metric properties
func (m *PerformanceNvmeMetricProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var performanceNvmeMetricPropertiesTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceNvmeMetricPropertiesTypeDurationPropEnum = append(performanceNvmeMetricPropertiesTypeDurationPropEnum, v)
	}
}

const (

	// PerformanceNvmeMetricPropertiesDurationPT15S captures enum value "PT15S"
	PerformanceNvmeMetricPropertiesDurationPT15S string = "PT15S"

	// PerformanceNvmeMetricPropertiesDurationPT4M captures enum value "PT4M"
	PerformanceNvmeMetricPropertiesDurationPT4M string = "PT4M"

	// PerformanceNvmeMetricPropertiesDurationPT30M captures enum value "PT30M"
	PerformanceNvmeMetricPropertiesDurationPT30M string = "PT30M"

	// PerformanceNvmeMetricPropertiesDurationPT2H captures enum value "PT2H"
	PerformanceNvmeMetricPropertiesDurationPT2H string = "PT2H"

	// PerformanceNvmeMetricPropertiesDurationP1D captures enum value "P1D"
	PerformanceNvmeMetricPropertiesDurationP1D string = "P1D"

	// PerformanceNvmeMetricPropertiesDurationPT5M captures enum value "PT5M"
	PerformanceNvmeMetricPropertiesDurationPT5M string = "PT5M"
)

// prop value enum
func (m *PerformanceNvmeMetricProperties) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceNvmeMetricPropertiesTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceNvmeMetricProperties) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", *m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceNvmeMetricProperties) validateIops(formats strfmt.Registry) error {
	if swag.IsZero(m.Iops) { // not required
		return nil
	}

	if m.Iops != nil {
		if err := m.Iops.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceNvmeMetricProperties) validateLatency(formats strfmt.Registry) error {
	if swag.IsZero(m.Latency) { // not required
		return nil
	}

	if m.Latency != nil {
		if err := m.Latency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

var performanceNvmeMetricPropertiesTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceNvmeMetricPropertiesTypeStatusPropEnum = append(performanceNvmeMetricPropertiesTypeStatusPropEnum, v)
	}
}

const (

	// PerformanceNvmeMetricPropertiesStatusOk captures enum value "ok"
	PerformanceNvmeMetricPropertiesStatusOk string = "ok"

	// PerformanceNvmeMetricPropertiesStatusError captures enum value "error"
	PerformanceNvmeMetricPropertiesStatusError string = "error"

	// PerformanceNvmeMetricPropertiesStatusPartialNoData captures enum value "partial_no_data"
	PerformanceNvmeMetricPropertiesStatusPartialNoData string = "partial_no_data"

	// PerformanceNvmeMetricPropertiesStatusPartialNoResponse captures enum value "partial_no_response"
	PerformanceNvmeMetricPropertiesStatusPartialNoResponse string = "partial_no_response"

	// PerformanceNvmeMetricPropertiesStatusPartialOtherError captures enum value "partial_other_error"
	PerformanceNvmeMetricPropertiesStatusPartialOtherError string = "partial_other_error"

	// PerformanceNvmeMetricPropertiesStatusNegativeDelta captures enum value "negative_delta"
	PerformanceNvmeMetricPropertiesStatusNegativeDelta string = "negative_delta"

	// PerformanceNvmeMetricPropertiesStatusNotFound captures enum value "not_found"
	PerformanceNvmeMetricPropertiesStatusNotFound string = "not_found"

	// PerformanceNvmeMetricPropertiesStatusBackfilledData captures enum value "backfilled_data"
	PerformanceNvmeMetricPropertiesStatusBackfilledData string = "backfilled_data"

	// PerformanceNvmeMetricPropertiesStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PerformanceNvmeMetricPropertiesStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// PerformanceNvmeMetricPropertiesStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PerformanceNvmeMetricPropertiesStatusInconsistentOldData string = "inconsistent_old_data"

	// PerformanceNvmeMetricPropertiesStatusPartialNoUUID captures enum value "partial_no_uuid"
	PerformanceNvmeMetricPropertiesStatusPartialNoUUID string = "partial_no_uuid"
)

// prop value enum
func (m *PerformanceNvmeMetricProperties) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceNvmeMetricPropertiesTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceNvmeMetricProperties) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceNvmeMetricProperties) validateThroughput(formats strfmt.Registry) error {
	if swag.IsZero(m.Throughput) { // not required
		return nil
	}

	if m.Throughput != nil {
		if err := m.Throughput.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this performance nvme metric properties based on the context it is used
func (m *PerformanceNvmeMetricProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceNvmeMetricProperties) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceNvmeMetricProperties) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

	if m.Iops != nil {

		if swag.IsZero(m.Iops) { // not required
			return nil
		}

		if err := m.Iops.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceNvmeMetricProperties) contextValidateLatency(ctx context.Context, formats strfmt.Registry) error {

	if m.Latency != nil {

		if swag.IsZero(m.Latency) { // not required
			return nil
		}

		if err := m.Latency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceNvmeMetricProperties) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceNvmeMetricProperties) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

	if m.Throughput != nil {

		if swag.IsZero(m.Throughput) { // not required
			return nil
		}

		if err := m.Throughput.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceNvmeMetricProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceNvmeMetricProperties) UnmarshalBinary(b []byte) error {
	var res PerformanceNvmeMetricProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceNvmeMetricPropertiesInlineIops The rate of I/O operations observed at the storage object.
//
// swagger:model performance_nvme_metric_properties_inline_iops
type PerformanceNvmeMetricPropertiesInlineIops struct {

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

// Validate validates this performance nvme metric properties inline iops
func (m *PerformanceNvmeMetricPropertiesInlineIops) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance nvme metric properties inline iops based on the context it is used
func (m *PerformanceNvmeMetricPropertiesInlineIops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceNvmeMetricPropertiesInlineIops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceNvmeMetricPropertiesInlineIops) UnmarshalBinary(b []byte) error {
	var res PerformanceNvmeMetricPropertiesInlineIops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceNvmeMetricPropertiesInlineLatency The round trip latency in microseconds observed at the storage object.
//
// swagger:model performance_nvme_metric_properties_inline_latency
type PerformanceNvmeMetricPropertiesInlineLatency struct {

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

// Validate validates this performance nvme metric properties inline latency
func (m *PerformanceNvmeMetricPropertiesInlineLatency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance nvme metric properties inline latency based on the context it is used
func (m *PerformanceNvmeMetricPropertiesInlineLatency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceNvmeMetricPropertiesInlineLatency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceNvmeMetricPropertiesInlineLatency) UnmarshalBinary(b []byte) error {
	var res PerformanceNvmeMetricPropertiesInlineLatency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceNvmeMetricPropertiesInlineThroughput The rate of throughput bytes per second observed at the storage object.
//
// swagger:model performance_nvme_metric_properties_inline_throughput
type PerformanceNvmeMetricPropertiesInlineThroughput struct {

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

// Validate validates this performance nvme metric properties inline throughput
func (m *PerformanceNvmeMetricPropertiesInlineThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance nvme metric properties inline throughput based on the context it is used
func (m *PerformanceNvmeMetricPropertiesInlineThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceNvmeMetricPropertiesInlineThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceNvmeMetricPropertiesInlineThroughput) UnmarshalBinary(b []byte) error {
	var res PerformanceNvmeMetricPropertiesInlineThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}