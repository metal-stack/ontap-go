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

// FileCopy File copy
//
// swagger:model file_copy
type FileCopy struct {

	// The maximum amount of time (in seconds) that the source can be quiesced before a destination file must be made available for read-write traffic.
	// Example: 10
	CutoverTime *int64 `json:"cutover_time,omitempty" yaml:"cutover_time,omitempty"`

	// A list of source files along with the destinations they are copied to. If the terminal path component of the destination is a directory, then the source file's basename is replicated in that directory.
	FileCopyInlineFilesToCopy []*FileCopyInlineFilesToCopyInlineArrayItem `json:"files_to_copy,omitempty" yaml:"files_to_copy,omitempty"`

	// Specifies whether the source file should be held quiescent for the duration of the copy operation.
	HoldQuiescence *bool `json:"hold_quiescence,omitempty" yaml:"hold_quiescence,omitempty"`

	// Maximum amount of data, in bytes that can be transferred per second in support of this operation. A non-zero value less than 1MB/s is set to 1MB/s. A non-zero value greater than 1MB/s is truncated to the nearest integral megabyte value. If unspecified, the default value is "0" which means no range is set for the data transfer.
	MaxThroughput *int64 `json:"max_throughput,omitempty" yaml:"max_throughput,omitempty"`

	// The maximum amount of time (in seconds) that the source reference file can be quiesced before the corresponding destination file must be made available for read-write traffic.
	// Example: 10
	ReferenceCutoverTime *int64 `json:"reference_cutover_time,omitempty" yaml:"reference_cutover_time,omitempty"`

	// The source reference file. If a reference file is specified, data for other files being copied will be transferred as a difference from the reference file. This can save bandwidth and destination storage if the specified source files share blocks. If provided, this input must match one of the source file paths. This input need not be provided if only one source file is specified.
	// Example: svm1:volume1/file1
	ReferencePath *string `json:"reference_path,omitempty" yaml:"reference_path,omitempty"`
}

// Validate validates this file copy
func (m *FileCopy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileCopyInlineFilesToCopy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileCopy) validateFileCopyInlineFilesToCopy(formats strfmt.Registry) error {
	if swag.IsZero(m.FileCopyInlineFilesToCopy) { // not required
		return nil
	}

	for i := 0; i < len(m.FileCopyInlineFilesToCopy); i++ {
		if swag.IsZero(m.FileCopyInlineFilesToCopy[i]) { // not required
			continue
		}

		if m.FileCopyInlineFilesToCopy[i] != nil {
			if err := m.FileCopyInlineFilesToCopy[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files_to_copy" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files_to_copy" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this file copy based on the context it is used
func (m *FileCopy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFileCopyInlineFilesToCopy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileCopy) contextValidateFileCopyInlineFilesToCopy(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FileCopyInlineFilesToCopy); i++ {

		if m.FileCopyInlineFilesToCopy[i] != nil {

			if swag.IsZero(m.FileCopyInlineFilesToCopy[i]) { // not required
				return nil
			}

			if err := m.FileCopyInlineFilesToCopy[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files_to_copy" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files_to_copy" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileCopy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileCopy) UnmarshalBinary(b []byte) error {
	var res FileCopy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FileCopyInlineFilesToCopyInlineArrayItem file copy inline files to copy inline array item
//
// swagger:model file_copy_inline_files_to_copy_inline_array_item
type FileCopyInlineFilesToCopyInlineArrayItem struct {

	// destination
	Destination *FileReference `json:"destination,omitempty" yaml:"destination,omitempty"`

	// source
	Source *FileReference `json:"source,omitempty" yaml:"source,omitempty"`
}

// Validate validates this file copy inline files to copy inline array item
func (m *FileCopyInlineFilesToCopyInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileCopyInlineFilesToCopyInlineArrayItem) validateDestination(formats strfmt.Registry) error {
	if swag.IsZero(m.Destination) { // not required
		return nil
	}

	if m.Destination != nil {
		if err := m.Destination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

func (m *FileCopyInlineFilesToCopyInlineArrayItem) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this file copy inline files to copy inline array item based on the context it is used
func (m *FileCopyInlineFilesToCopyInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileCopyInlineFilesToCopyInlineArrayItem) contextValidateDestination(ctx context.Context, formats strfmt.Registry) error {

	if m.Destination != nil {

		if swag.IsZero(m.Destination) { // not required
			return nil
		}

		if err := m.Destination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

func (m *FileCopyInlineFilesToCopyInlineArrayItem) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {

		if swag.IsZero(m.Source) { // not required
			return nil
		}

		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileCopyInlineFilesToCopyInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileCopyInlineFilesToCopyInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res FileCopyInlineFilesToCopyInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}