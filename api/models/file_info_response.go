// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FileInfoResponse file info response
//
// swagger:model file_info_response
type FileInfoResponse struct {

	// links
	Links *FileInfoResponseInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// analytics
	Analytics *FileInfoResponseInlineAnalytics `json:"analytics,omitempty" yaml:"analytics,omitempty"`

	// file info response inline records
	FileInfoResponseInlineRecords []*FileInfo `json:"records,omitempty" yaml:"records,omitempty"`

	// Number of records.
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty" yaml:"num_records,omitempty"`
}

// Validate validates this file info response
func (m *FileInfoResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnalytics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileInfoResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileInfoResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *FileInfoResponse) validateAnalytics(formats strfmt.Registry) error {
	if swag.IsZero(m.Analytics) { // not required
		return nil
	}

	if m.Analytics != nil {
		if err := m.Analytics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("analytics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("analytics")
			}
			return err
		}
	}

	return nil
}

func (m *FileInfoResponse) validateFileInfoResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.FileInfoResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.FileInfoResponseInlineRecords); i++ {
		if swag.IsZero(m.FileInfoResponseInlineRecords[i]) { // not required
			continue
		}

		if m.FileInfoResponseInlineRecords[i] != nil {
			if err := m.FileInfoResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this file info response based on the context it is used
func (m *FileInfoResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAnalytics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileInfoResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileInfoResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *FileInfoResponse) contextValidateAnalytics(ctx context.Context, formats strfmt.Registry) error {

	if m.Analytics != nil {

		if swag.IsZero(m.Analytics) { // not required
			return nil
		}

		if err := m.Analytics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("analytics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("analytics")
			}
			return err
		}
	}

	return nil
}

func (m *FileInfoResponse) contextValidateFileInfoResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FileInfoResponseInlineRecords); i++ {

		if m.FileInfoResponseInlineRecords[i] != nil {

			if swag.IsZero(m.FileInfoResponseInlineRecords[i]) { // not required
				return nil
			}

			if err := m.FileInfoResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileInfoResponse) UnmarshalBinary(b []byte) error {
	var res FileInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FileInfoResponseInlineAnalytics Additional file system analytics information that is invariant amongst all elements in the collection. <br/>
// This property is only populated if file system analytics is enabled on the containing volume. <br/>
// This analytics object captures properties that are invariant amongst all elements included in the `records` array. The invariant properties are included here, rather than within the information for each element, to avoid returning an excessive amount of duplicated information when the collection is large.
//
// swagger:model file_info_response_inline_analytics
type FileInfoResponseInlineAnalytics struct {

	// by accessed time
	ByAccessedTime *FileInfoResponseInlineAnalyticsInlineByAccessedTime `json:"by_accessed_time,omitempty" yaml:"by_accessed_time,omitempty"`

	// by modified time
	ByModifiedTime *FileInfoResponseInlineAnalyticsInlineByModifiedTime `json:"by_modified_time,omitempty" yaml:"by_modified_time,omitempty"`
}

// Validate validates this file info response inline analytics
func (m *FileInfoResponseInlineAnalytics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateByAccessedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateByModifiedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileInfoResponseInlineAnalytics) validateByAccessedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ByAccessedTime) { // not required
		return nil
	}

	if m.ByAccessedTime != nil {
		if err := m.ByAccessedTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("analytics" + "." + "by_accessed_time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("analytics" + "." + "by_accessed_time")
			}
			return err
		}
	}

	return nil
}

func (m *FileInfoResponseInlineAnalytics) validateByModifiedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ByModifiedTime) { // not required
		return nil
	}

	if m.ByModifiedTime != nil {
		if err := m.ByModifiedTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("analytics" + "." + "by_modified_time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("analytics" + "." + "by_modified_time")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this file info response inline analytics based on the context it is used
func (m *FileInfoResponseInlineAnalytics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateByAccessedTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateByModifiedTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileInfoResponseInlineAnalytics) contextValidateByAccessedTime(ctx context.Context, formats strfmt.Registry) error {

	if m.ByAccessedTime != nil {

		if swag.IsZero(m.ByAccessedTime) { // not required
			return nil
		}

		if err := m.ByAccessedTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("analytics" + "." + "by_accessed_time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("analytics" + "." + "by_accessed_time")
			}
			return err
		}
	}

	return nil
}

func (m *FileInfoResponseInlineAnalytics) contextValidateByModifiedTime(ctx context.Context, formats strfmt.Registry) error {

	if m.ByModifiedTime != nil {

		if swag.IsZero(m.ByModifiedTime) { // not required
			return nil
		}

		if err := m.ByModifiedTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("analytics" + "." + "by_modified_time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("analytics" + "." + "by_modified_time")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileInfoResponseInlineAnalytics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileInfoResponseInlineAnalytics) UnmarshalBinary(b []byte) error {
	var res FileInfoResponseInlineAnalytics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FileInfoResponseInlineAnalyticsInlineByAccessedTime File system analytics information, broken down by date of last access.
//
// swagger:model file_info_response_inline_analytics_inline_by_accessed_time
type FileInfoResponseInlineAnalyticsInlineByAccessedTime struct {

	// bytes used
	BytesUsed *FileInfoResponseInlineAnalyticsInlineByAccessedTimeInlineBytesUsed `json:"bytes_used,omitempty" yaml:"bytes_used,omitempty"`
}

// Validate validates this file info response inline analytics inline by accessed time
func (m *FileInfoResponseInlineAnalyticsInlineByAccessedTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBytesUsed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileInfoResponseInlineAnalyticsInlineByAccessedTime) validateBytesUsed(formats strfmt.Registry) error {
	if swag.IsZero(m.BytesUsed) { // not required
		return nil
	}

	if m.BytesUsed != nil {
		if err := m.BytesUsed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("analytics" + "." + "by_accessed_time" + "." + "bytes_used")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("analytics" + "." + "by_accessed_time" + "." + "bytes_used")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this file info response inline analytics inline by accessed time based on the context it is used
func (m *FileInfoResponseInlineAnalyticsInlineByAccessedTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBytesUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileInfoResponseInlineAnalyticsInlineByAccessedTime) contextValidateBytesUsed(ctx context.Context, formats strfmt.Registry) error {

	if m.BytesUsed != nil {

		if swag.IsZero(m.BytesUsed) { // not required
			return nil
		}

		if err := m.BytesUsed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("analytics" + "." + "by_accessed_time" + "." + "bytes_used")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("analytics" + "." + "by_accessed_time" + "." + "bytes_used")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileInfoResponseInlineAnalyticsInlineByAccessedTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileInfoResponseInlineAnalyticsInlineByAccessedTime) UnmarshalBinary(b []byte) error {
	var res FileInfoResponseInlineAnalyticsInlineByAccessedTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FileInfoResponseInlineAnalyticsInlineByAccessedTimeInlineBytesUsed Number of bytes used on-disk, broken down by date of last access.
//
// swagger:model file_info_response_inline_analytics_inline_by_accessed_time_inline_bytes_used
type FileInfoResponseInlineAnalyticsInlineByAccessedTimeInlineBytesUsed struct {

	// labels
	Labels AnalyticsHistogramByTimeLabelsArrayInline `json:"labels,omitempty" yaml:"labels,omitempty"`
}

// Validate validates this file info response inline analytics inline by accessed time inline bytes used
func (m *FileInfoResponseInlineAnalyticsInlineByAccessedTimeInlineBytesUsed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileInfoResponseInlineAnalyticsInlineByAccessedTimeInlineBytesUsed) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if err := m.Labels.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("analytics" + "." + "by_accessed_time" + "." + "bytes_used" + "." + "labels")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("analytics" + "." + "by_accessed_time" + "." + "bytes_used" + "." + "labels")
		}
		return err
	}

	return nil
}

// ContextValidate validate this file info response inline analytics inline by accessed time inline bytes used based on the context it is used
func (m *FileInfoResponseInlineAnalyticsInlineByAccessedTimeInlineBytesUsed) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileInfoResponseInlineAnalyticsInlineByAccessedTimeInlineBytesUsed) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Labels.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("analytics" + "." + "by_accessed_time" + "." + "bytes_used" + "." + "labels")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("analytics" + "." + "by_accessed_time" + "." + "bytes_used" + "." + "labels")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileInfoResponseInlineAnalyticsInlineByAccessedTimeInlineBytesUsed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileInfoResponseInlineAnalyticsInlineByAccessedTimeInlineBytesUsed) UnmarshalBinary(b []byte) error {
	var res FileInfoResponseInlineAnalyticsInlineByAccessedTimeInlineBytesUsed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FileInfoResponseInlineAnalyticsInlineByModifiedTime File system analytics information, broken down by date of last modification.
//
// swagger:model file_info_response_inline_analytics_inline_by_modified_time
type FileInfoResponseInlineAnalyticsInlineByModifiedTime struct {

	// bytes used
	BytesUsed *FileInfoResponseInlineAnalyticsInlineByModifiedTimeInlineBytesUsed `json:"bytes_used,omitempty" yaml:"bytes_used,omitempty"`
}

// Validate validates this file info response inline analytics inline by modified time
func (m *FileInfoResponseInlineAnalyticsInlineByModifiedTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBytesUsed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileInfoResponseInlineAnalyticsInlineByModifiedTime) validateBytesUsed(formats strfmt.Registry) error {
	if swag.IsZero(m.BytesUsed) { // not required
		return nil
	}

	if m.BytesUsed != nil {
		if err := m.BytesUsed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("analytics" + "." + "by_modified_time" + "." + "bytes_used")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("analytics" + "." + "by_modified_time" + "." + "bytes_used")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this file info response inline analytics inline by modified time based on the context it is used
func (m *FileInfoResponseInlineAnalyticsInlineByModifiedTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBytesUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileInfoResponseInlineAnalyticsInlineByModifiedTime) contextValidateBytesUsed(ctx context.Context, formats strfmt.Registry) error {

	if m.BytesUsed != nil {

		if swag.IsZero(m.BytesUsed) { // not required
			return nil
		}

		if err := m.BytesUsed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("analytics" + "." + "by_modified_time" + "." + "bytes_used")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("analytics" + "." + "by_modified_time" + "." + "bytes_used")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileInfoResponseInlineAnalyticsInlineByModifiedTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileInfoResponseInlineAnalyticsInlineByModifiedTime) UnmarshalBinary(b []byte) error {
	var res FileInfoResponseInlineAnalyticsInlineByModifiedTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FileInfoResponseInlineAnalyticsInlineByModifiedTimeInlineBytesUsed Number of bytes used on-disk, broken down by date of last modification.
//
// swagger:model file_info_response_inline_analytics_inline_by_modified_time_inline_bytes_used
type FileInfoResponseInlineAnalyticsInlineByModifiedTimeInlineBytesUsed struct {

	// labels
	Labels AnalyticsHistogramByTimeLabelsArrayInline `json:"labels,omitempty" yaml:"labels,omitempty"`
}

// Validate validates this file info response inline analytics inline by modified time inline bytes used
func (m *FileInfoResponseInlineAnalyticsInlineByModifiedTimeInlineBytesUsed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileInfoResponseInlineAnalyticsInlineByModifiedTimeInlineBytesUsed) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if err := m.Labels.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("analytics" + "." + "by_modified_time" + "." + "bytes_used" + "." + "labels")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("analytics" + "." + "by_modified_time" + "." + "bytes_used" + "." + "labels")
		}
		return err
	}

	return nil
}

// ContextValidate validate this file info response inline analytics inline by modified time inline bytes used based on the context it is used
func (m *FileInfoResponseInlineAnalyticsInlineByModifiedTimeInlineBytesUsed) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileInfoResponseInlineAnalyticsInlineByModifiedTimeInlineBytesUsed) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Labels.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("analytics" + "." + "by_modified_time" + "." + "bytes_used" + "." + "labels")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("analytics" + "." + "by_modified_time" + "." + "bytes_used" + "." + "labels")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileInfoResponseInlineAnalyticsInlineByModifiedTimeInlineBytesUsed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileInfoResponseInlineAnalyticsInlineByModifiedTimeInlineBytesUsed) UnmarshalBinary(b []byte) error {
	var res FileInfoResponseInlineAnalyticsInlineByModifiedTimeInlineBytesUsed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FileInfoResponseInlineLinks file info response inline links
//
// swagger:model file_info_response_inline__links
type FileInfoResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty" yaml:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this file info response inline links
func (m *FileInfoResponseInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileInfoResponseInlineLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *FileInfoResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this file info response inline links based on the context it is used
func (m *FileInfoResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileInfoResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {

		if swag.IsZero(m.Next) { // not required
			return nil
		}

		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *FileInfoResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *FileInfoResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileInfoResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res FileInfoResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
