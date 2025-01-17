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

// PerformanceLunMetric Performance numbers, such as IOPS latency and throughput.
//
// swagger:model performance_lun_metric
type PerformanceLunMetric struct {

	// links
	Links *PerformanceLunMetricInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Read Only: true
	// Enum: ["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]
	Duration *string `json:"duration,omitempty" yaml:"duration,omitempty"`

	// iops
	Iops *PerformanceLunMetricInlineIops `json:"iops,omitempty" yaml:"iops,omitempty"`

	// latency
	Latency *PerformanceLunMetricInlineLatency `json:"latency,omitempty" yaml:"latency,omitempty"`

	// Errors associated with the sample. For example, if the aggregation of data over multiple nodes fails, then any partial errors might return "ok" on success or "error" on an internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "Inconsistent_ delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "Inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: ["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]
	Status *string `json:"status,omitempty" yaml:"status,omitempty"`

	// throughput
	Throughput *PerformanceLunMetricInlineThroughput `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25 11:20:13
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`

	// The unique identifier of the LUN.
	//
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this performance lun metric
func (m *PerformanceLunMetric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceLunMetric) validateLinks(formats strfmt.Registry) error {
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

var performanceLunMetricTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceLunMetricTypeDurationPropEnum = append(performanceLunMetricTypeDurationPropEnum, v)
	}
}

const (

	// PerformanceLunMetricDurationPT15S captures enum value "PT15S"
	PerformanceLunMetricDurationPT15S string = "PT15S"

	// PerformanceLunMetricDurationPT4M captures enum value "PT4M"
	PerformanceLunMetricDurationPT4M string = "PT4M"

	// PerformanceLunMetricDurationPT30M captures enum value "PT30M"
	PerformanceLunMetricDurationPT30M string = "PT30M"

	// PerformanceLunMetricDurationPT2H captures enum value "PT2H"
	PerformanceLunMetricDurationPT2H string = "PT2H"

	// PerformanceLunMetricDurationP1D captures enum value "P1D"
	PerformanceLunMetricDurationP1D string = "P1D"

	// PerformanceLunMetricDurationPT5M captures enum value "PT5M"
	PerformanceLunMetricDurationPT5M string = "PT5M"
)

// prop value enum
func (m *PerformanceLunMetric) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceLunMetricTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceLunMetric) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", *m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceLunMetric) validateIops(formats strfmt.Registry) error {
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

func (m *PerformanceLunMetric) validateLatency(formats strfmt.Registry) error {
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

var performanceLunMetricTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceLunMetricTypeStatusPropEnum = append(performanceLunMetricTypeStatusPropEnum, v)
	}
}

const (

	// PerformanceLunMetricStatusOk captures enum value "ok"
	PerformanceLunMetricStatusOk string = "ok"

	// PerformanceLunMetricStatusError captures enum value "error"
	PerformanceLunMetricStatusError string = "error"

	// PerformanceLunMetricStatusPartialNoData captures enum value "partial_no_data"
	PerformanceLunMetricStatusPartialNoData string = "partial_no_data"

	// PerformanceLunMetricStatusPartialNoResponse captures enum value "partial_no_response"
	PerformanceLunMetricStatusPartialNoResponse string = "partial_no_response"

	// PerformanceLunMetricStatusPartialOtherError captures enum value "partial_other_error"
	PerformanceLunMetricStatusPartialOtherError string = "partial_other_error"

	// PerformanceLunMetricStatusNegativeDelta captures enum value "negative_delta"
	PerformanceLunMetricStatusNegativeDelta string = "negative_delta"

	// PerformanceLunMetricStatusNotFound captures enum value "not_found"
	PerformanceLunMetricStatusNotFound string = "not_found"

	// PerformanceLunMetricStatusBackfilledData captures enum value "backfilled_data"
	PerformanceLunMetricStatusBackfilledData string = "backfilled_data"

	// PerformanceLunMetricStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PerformanceLunMetricStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// PerformanceLunMetricStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PerformanceLunMetricStatusInconsistentOldData string = "inconsistent_old_data"

	// PerformanceLunMetricStatusPartialNoUUID captures enum value "partial_no_uuid"
	PerformanceLunMetricStatusPartialNoUUID string = "partial_no_uuid"
)

// prop value enum
func (m *PerformanceLunMetric) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceLunMetricTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceLunMetric) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceLunMetric) validateThroughput(formats strfmt.Registry) error {
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

func (m *PerformanceLunMetric) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this performance lun metric based on the context it is used
func (m *PerformanceLunMetric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceLunMetric) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceLunMetric) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceLunMetric) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceLunMetric) contextValidateLatency(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceLunMetric) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceLunMetric) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceLunMetric) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceLunMetric) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceLunMetric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceLunMetric) UnmarshalBinary(b []byte) error {
	var res PerformanceLunMetric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceLunMetricInlineIops The rate of I/O operations observed at the storage object.
//
// swagger:model performance_lun_metric_inline_iops
type PerformanceLunMetricInlineIops struct {

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

// Validate validates this performance lun metric inline iops
func (m *PerformanceLunMetricInlineIops) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance lun metric inline iops based on the context it is used
func (m *PerformanceLunMetricInlineIops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceLunMetricInlineIops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceLunMetricInlineIops) UnmarshalBinary(b []byte) error {
	var res PerformanceLunMetricInlineIops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceLunMetricInlineLatency The round trip latency in microseconds observed at the storage object.
//
// swagger:model performance_lun_metric_inline_latency
type PerformanceLunMetricInlineLatency struct {

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

// Validate validates this performance lun metric inline latency
func (m *PerformanceLunMetricInlineLatency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance lun metric inline latency based on the context it is used
func (m *PerformanceLunMetricInlineLatency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceLunMetricInlineLatency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceLunMetricInlineLatency) UnmarshalBinary(b []byte) error {
	var res PerformanceLunMetricInlineLatency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceLunMetricInlineLinks performance lun metric inline links
//
// swagger:model performance_lun_metric_inline__links
type PerformanceLunMetricInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this performance lun metric inline links
func (m *PerformanceLunMetricInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceLunMetricInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this performance lun metric inline links based on the context it is used
func (m *PerformanceLunMetricInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceLunMetricInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PerformanceLunMetricInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceLunMetricInlineLinks) UnmarshalBinary(b []byte) error {
	var res PerformanceLunMetricInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceLunMetricInlineThroughput The rate of throughput bytes per second observed at the storage object.
//
// swagger:model performance_lun_metric_inline_throughput
type PerformanceLunMetricInlineThroughput struct {

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

// Validate validates this performance lun metric inline throughput
func (m *PerformanceLunMetricInlineThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance lun metric inline throughput based on the context it is used
func (m *PerformanceLunMetricInlineThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceLunMetricInlineThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceLunMetricInlineThroughput) UnmarshalBinary(b []byte) error {
	var res PerformanceLunMetricInlineThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}