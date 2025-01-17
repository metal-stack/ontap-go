// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TopMetricsDirectory Information about a directory's IO metrics.
// Example: {"_links":{"directory":{"href":"/api/storage/volumes/02178914-5f67-11eb-b987-005056ac5da5/files/dir_1%2fdir_2"},"metadata":{"href":"/api/storage/volumes/02178914-5f67-11eb-b987-005056ac5da5/files/dir_1%2fdir_2?return_metadata=true"}},"iops":{"error":{"lower_bound":2,"upper_bound":5},"read":2,"write":3},"non_recursive_bytes_used":300,"path":"/dir1_/dir_2","svm":{"name":"vserver_2","uuid":"42ee3002-67dd-11ea-8508-005056a7b8ac"},"throughput":{"error":{"lower_bound":3,"upper_bound":6},"read":3,"write":5},"volume":{"name":"vol_8","uuid":"c05eb66a-685f-11ea-8508-005056a7b8ac"}}
//
// swagger:model top_metrics_directory
type TopMetricsDirectory struct {

	// links
	Links *TopMetricsDirectoryInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// iops
	Iops *TopMetricsDirectoryInlineIops `json:"iops,omitempty" yaml:"iops,omitempty"`

	// Non-recursive bytes used by the contents of a directory.
	// Example: 300
	// Read Only: true
	NonRecursiveBytesUsed *int64 `json:"non_recursive_bytes_used,omitempty" yaml:"non_recursive_bytes_used,omitempty"`

	// Path of the directory.
	// Example: /dir_abc/dir_123/dir_20
	// Read Only: true
	Path *string `json:"path,omitempty" yaml:"path,omitempty"`

	// svm
	Svm *TopMetricsDirectoryInlineSvm `json:"svm,omitempty" yaml:"svm,omitempty"`

	// throughput
	Throughput *TopMetricsDirectoryInlineThroughput `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	// volume
	Volume *TopMetricsDirectoryInlineVolume `json:"volume,omitempty" yaml:"volume,omitempty"`
}

// Validate validates this top metrics directory
func (m *TopMetricsDirectory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsDirectory) validateLinks(formats strfmt.Registry) error {
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

func (m *TopMetricsDirectory) validateIops(formats strfmt.Registry) error {
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

func (m *TopMetricsDirectory) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsDirectory) validateThroughput(formats strfmt.Registry) error {
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

func (m *TopMetricsDirectory) validateVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics directory based on the context it is used
func (m *TopMetricsDirectory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNonRecursiveBytesUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsDirectory) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TopMetricsDirectory) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TopMetricsDirectory) contextValidateNonRecursiveBytesUsed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "non_recursive_bytes_used", "body", m.NonRecursiveBytesUsed); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsDirectory) contextValidatePath(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsDirectory) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {

		if swag.IsZero(m.Svm) { // not required
			return nil
		}

		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsDirectory) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TopMetricsDirectory) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.Volume != nil {

		if swag.IsZero(m.Volume) { // not required
			return nil
		}

		if err := m.Volume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsDirectory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsDirectory) UnmarshalBinary(b []byte) error {
	var res TopMetricsDirectory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsDirectoryInlineIops top metrics directory inline iops
//
// swagger:model top_metrics_directory_inline_iops
type TopMetricsDirectoryInlineIops struct {

	// error
	Error *TopMetricValueErrorBounds `json:"error,omitempty" yaml:"error,omitempty"`

	// Average number of read operations per second.
	// Example: 10
	// Read Only: true
	Read *int64 `json:"read,omitempty" yaml:"read,omitempty"`

	// Average number of write operations per second.
	// Example: 5
	// Read Only: true
	Write *int64 `json:"write,omitempty" yaml:"write,omitempty"`
}

// Validate validates this top metrics directory inline iops
func (m *TopMetricsDirectoryInlineIops) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsDirectoryInlineIops) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iops" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics directory inline iops based on the context it is used
func (m *TopMetricsDirectoryInlineIops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRead(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWrite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsDirectoryInlineIops) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {

		if swag.IsZero(m.Error) { // not required
			return nil
		}

		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iops" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsDirectoryInlineIops) contextValidateRead(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "iops"+"."+"read", "body", m.Read); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsDirectoryInlineIops) contextValidateWrite(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "iops"+"."+"write", "body", m.Write); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsDirectoryInlineIops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsDirectoryInlineIops) UnmarshalBinary(b []byte) error {
	var res TopMetricsDirectoryInlineIops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsDirectoryInlineLinks top metrics directory inline links
//
// swagger:model top_metrics_directory_inline__links
type TopMetricsDirectoryInlineLinks struct {

	// metadata
	Metadata *Href `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this top metrics directory inline links
func (m *TopMetricsDirectoryInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
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

func (m *TopMetricsDirectoryInlineLinks) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "metadata")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsDirectoryInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics directory inline links based on the context it is used
func (m *TopMetricsDirectoryInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
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

func (m *TopMetricsDirectoryInlineLinks) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if swag.IsZero(m.Metadata) { // not required
			return nil
		}

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "metadata")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsDirectoryInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsDirectoryInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsDirectoryInlineLinks) UnmarshalBinary(b []byte) error {
	var res TopMetricsDirectoryInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsDirectoryInlineSvm SVM, applies only to SVM-scoped objects.
//
// swagger:model top_metrics_directory_inline_svm
type TopMetricsDirectoryInlineSvm struct {

	// links
	Links *TopMetricsDirectoryInlineSvmInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this top metrics directory inline svm
func (m *TopMetricsDirectoryInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsDirectoryInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics directory inline svm based on the context it is used
func (m *TopMetricsDirectoryInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsDirectoryInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsDirectoryInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsDirectoryInlineSvm) UnmarshalBinary(b []byte) error {
	var res TopMetricsDirectoryInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsDirectoryInlineSvmInlineLinks top metrics directory inline svm inline links
//
// swagger:model top_metrics_directory_inline_svm_inline__links
type TopMetricsDirectoryInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this top metrics directory inline svm inline links
func (m *TopMetricsDirectoryInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsDirectoryInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics directory inline svm inline links based on the context it is used
func (m *TopMetricsDirectoryInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsDirectoryInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsDirectoryInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsDirectoryInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res TopMetricsDirectoryInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsDirectoryInlineThroughput top metrics directory inline throughput
//
// swagger:model top_metrics_directory_inline_throughput
type TopMetricsDirectoryInlineThroughput struct {

	// error
	Error *TopMetricValueErrorBounds `json:"error,omitempty" yaml:"error,omitempty"`

	// Average number of read bytes received per second.
	// Example: 3
	// Read Only: true
	Read *int64 `json:"read,omitempty" yaml:"read,omitempty"`

	// Average number of write bytes received per second.
	// Example: 20
	// Read Only: true
	Write *int64 `json:"write,omitempty" yaml:"write,omitempty"`
}

// Validate validates this top metrics directory inline throughput
func (m *TopMetricsDirectoryInlineThroughput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsDirectoryInlineThroughput) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throughput" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics directory inline throughput based on the context it is used
func (m *TopMetricsDirectoryInlineThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRead(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWrite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsDirectoryInlineThroughput) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {

		if swag.IsZero(m.Error) { // not required
			return nil
		}

		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throughput" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsDirectoryInlineThroughput) contextValidateRead(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "throughput"+"."+"read", "body", m.Read); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsDirectoryInlineThroughput) contextValidateWrite(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "throughput"+"."+"write", "body", m.Write); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsDirectoryInlineThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsDirectoryInlineThroughput) UnmarshalBinary(b []byte) error {
	var res TopMetricsDirectoryInlineThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsDirectoryInlineVolume top metrics directory inline volume
//
// swagger:model top_metrics_directory_inline_volume
type TopMetricsDirectoryInlineVolume struct {

	// links
	Links *TopMetricsDirectoryInlineVolumeInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// The name of the volume. This field cannot be specified in a POST or PATCH method.
	// Example: volume1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	// Example: 028baa66-41bd-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this top metrics directory inline volume
func (m *TopMetricsDirectoryInlineVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsDirectoryInlineVolume) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics directory inline volume based on the context it is used
func (m *TopMetricsDirectoryInlineVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsDirectoryInlineVolume) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsDirectoryInlineVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsDirectoryInlineVolume) UnmarshalBinary(b []byte) error {
	var res TopMetricsDirectoryInlineVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsDirectoryInlineVolumeInlineLinks top metrics directory inline volume inline links
//
// swagger:model top_metrics_directory_inline_volume_inline__links
type TopMetricsDirectoryInlineVolumeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this top metrics directory inline volume inline links
func (m *TopMetricsDirectoryInlineVolumeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsDirectoryInlineVolumeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics directory inline volume inline links based on the context it is used
func (m *TopMetricsDirectoryInlineVolumeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsDirectoryInlineVolumeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsDirectoryInlineVolumeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsDirectoryInlineVolumeInlineLinks) UnmarshalBinary(b []byte) error {
	var res TopMetricsDirectoryInlineVolumeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}