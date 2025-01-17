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

// InterfaceStatistics These are the total throughput raw performance data for the interface.
//
// swagger:model interface_statistics
type InterfaceStatistics struct {

	// Errors associated with the sample. For example, if the aggregation of data over multiple nodes fails, then any partial errors might return "ok" on success or "error" on an internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "inconsistent_delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Enum: ["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]
	Status *string `json:"status,omitempty" yaml:"status,omitempty"`

	// throughput raw
	ThroughputRaw *InterfaceStatisticsInlineThroughputRaw `json:"throughput_raw,omitempty" yaml:"throughput_raw,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25 11:20:13
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
}

// Validate validates this interface statistics
func (m *InterfaceStatistics) Validate(formats strfmt.Registry) error {
	var res []error

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

var interfaceStatisticsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceStatisticsTypeStatusPropEnum = append(interfaceStatisticsTypeStatusPropEnum, v)
	}
}

const (

	// InterfaceStatisticsStatusOk captures enum value "ok"
	InterfaceStatisticsStatusOk string = "ok"

	// InterfaceStatisticsStatusError captures enum value "error"
	InterfaceStatisticsStatusError string = "error"

	// InterfaceStatisticsStatusPartialNoData captures enum value "partial_no_data"
	InterfaceStatisticsStatusPartialNoData string = "partial_no_data"

	// InterfaceStatisticsStatusPartialNoUUID captures enum value "partial_no_uuid"
	InterfaceStatisticsStatusPartialNoUUID string = "partial_no_uuid"

	// InterfaceStatisticsStatusPartialNoResponse captures enum value "partial_no_response"
	InterfaceStatisticsStatusPartialNoResponse string = "partial_no_response"

	// InterfaceStatisticsStatusPartialOtherError captures enum value "partial_other_error"
	InterfaceStatisticsStatusPartialOtherError string = "partial_other_error"

	// InterfaceStatisticsStatusNegativeDelta captures enum value "negative_delta"
	InterfaceStatisticsStatusNegativeDelta string = "negative_delta"

	// InterfaceStatisticsStatusBackfilledData captures enum value "backfilled_data"
	InterfaceStatisticsStatusBackfilledData string = "backfilled_data"

	// InterfaceStatisticsStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	InterfaceStatisticsStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// InterfaceStatisticsStatusInconsistentOldData captures enum value "inconsistent_old_data"
	InterfaceStatisticsStatusInconsistentOldData string = "inconsistent_old_data"
)

// prop value enum
func (m *InterfaceStatistics) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceStatisticsTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceStatistics) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceStatistics) validateThroughputRaw(formats strfmt.Registry) error {
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

func (m *InterfaceStatistics) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this interface statistics based on the context it is used
func (m *InterfaceStatistics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateThroughputRaw(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceStatistics) contextValidateThroughputRaw(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *InterfaceStatistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceStatistics) UnmarshalBinary(b []byte) error {
	var res InterfaceStatistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InterfaceStatisticsInlineThroughputRaw Throughput bytes observed at the interface. This can be used along with delta time to calculate the rate of throughput bytes per unit of time.
//
// swagger:model interface_statistics_inline_throughput_raw
type InterfaceStatisticsInlineThroughputRaw struct {

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

// Validate validates this interface statistics inline throughput raw
func (m *InterfaceStatisticsInlineThroughputRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this interface statistics inline throughput raw based on context it is used
func (m *InterfaceStatisticsInlineThroughputRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceStatisticsInlineThroughputRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceStatisticsInlineThroughputRaw) UnmarshalBinary(b []byte) error {
	var res InterfaceStatisticsInlineThroughputRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
