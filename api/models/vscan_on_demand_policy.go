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
	"github.com/go-openapi/validate"
)

// VscanOnDemandPolicy Use On-Demand scanning to check files for viruses on a schedule. An On-Demand policy defines the scope of an On-Demand scan.
//
// swagger:model vscan_on_demand_policy
type VscanOnDemandPolicy struct {

	// The path from the Vserver root where the task report is created.
	// Example: /vol0/report_dir
	LogPath *string `json:"log_path,omitempty" yaml:"log_path,omitempty"`

	// On-Demand task name
	// Example: task-1
	// Max Length: 256
	// Min Length: 1
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// schedule
	Schedule *VscanOnDemandPolicyInlineSchedule `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// scope
	Scope *VscanOnDemandPolicyInlineScope `json:"scope,omitempty" yaml:"scope,omitempty"`

	// List of paths that need to be scanned.
	// Example: ["/vol1/","/vol2/cifs/"]
	VscanOnDemandPolicyInlineScanPaths []*string `json:"scan_paths,omitempty" yaml:"scan_paths,omitempty"`
}

// Validate validates this vscan on demand policy
func (m *VscanOnDemandPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VscanOnDemandPolicy) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 256); err != nil {
		return err
	}

	return nil
}

func (m *VscanOnDemandPolicy) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *VscanOnDemandPolicy) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if m.Scope != nil {
		if err := m.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vscan on demand policy based on the context it is used
func (m *VscanOnDemandPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VscanOnDemandPolicy) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {

		if swag.IsZero(m.Schedule) { // not required
			return nil
		}

		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *VscanOnDemandPolicy) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if m.Scope != nil {

		if swag.IsZero(m.Scope) { // not required
			return nil
		}

		if err := m.Scope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VscanOnDemandPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VscanOnDemandPolicy) UnmarshalBinary(b []byte) error {
	var res VscanOnDemandPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VscanOnDemandPolicyInlineSchedule Schedule of the task.
//
// swagger:model vscan_on_demand_policy_inline_schedule
type VscanOnDemandPolicyInlineSchedule struct {

	// links
	Links *VscanOnDemandPolicyInlineScheduleInlineLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// Job schedule name
	// Example: weekly
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`

	// Job schedule UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this vscan on demand policy inline schedule
func (m *VscanOnDemandPolicyInlineSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VscanOnDemandPolicyInlineSchedule) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vscan on demand policy inline schedule based on the context it is used
func (m *VscanOnDemandPolicyInlineSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VscanOnDemandPolicyInlineSchedule) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VscanOnDemandPolicyInlineSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VscanOnDemandPolicyInlineSchedule) UnmarshalBinary(b []byte) error {
	var res VscanOnDemandPolicyInlineSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VscanOnDemandPolicyInlineScheduleInlineLinks vscan on demand policy inline schedule inline links
//
// swagger:model vscan_on_demand_policy_inline_schedule_inline__links
type VscanOnDemandPolicyInlineScheduleInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this vscan on demand policy inline schedule inline links
func (m *VscanOnDemandPolicyInlineScheduleInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VscanOnDemandPolicyInlineScheduleInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vscan on demand policy inline schedule inline links based on the context it is used
func (m *VscanOnDemandPolicyInlineScheduleInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VscanOnDemandPolicyInlineScheduleInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {

		if swag.IsZero(m.Self) { // not required
			return nil
		}

		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VscanOnDemandPolicyInlineScheduleInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VscanOnDemandPolicyInlineScheduleInlineLinks) UnmarshalBinary(b []byte) error {
	var res VscanOnDemandPolicyInlineScheduleInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VscanOnDemandPolicyInlineScope vscan on demand policy inline scope
//
// swagger:model vscan_on_demand_policy_inline_scope
type VscanOnDemandPolicyInlineScope struct {

	// List of file extensions for which scanning is not performed.
	// Example: ["mp3","mp4"]
	// Max Items: 300
	// Min Items: 1
	ExcludeExtensions []*string `json:"exclude_extensions,omitempty" yaml:"exclude_extensions,omitempty"`

	// List of file paths for which scanning must not be performed.
	// Example: ["/vol1/cold-files/","/vol1/cifs/names"]
	// Max Items: 100
	// Min Items: 1
	ExcludePaths []*string `json:"exclude_paths,omitempty" yaml:"exclude_paths,omitempty"`

	// List of file extensions to be scanned.
	// Example: ["vmdk","mp*"]
	// Max Items: 300
	// Min Items: 1
	IncludeExtensions []*string `json:"include_extensions,omitempty" yaml:"include_extensions,omitempty"`

	// Maximum file size, in bytes, allowed for scanning.
	// Example: 10737418240
	// Maximum: 1.099511627776e+12
	// Minimum: 1024
	MaxFileSize *int64 `json:"max_file_size,omitempty" yaml:"max_file_size,omitempty"`

	// Specifies whether or not files without any extension can be scanned.
	ScanWithoutExtension *bool `json:"scan_without_extension,omitempty" yaml:"scan_without_extension,omitempty"`
}

// Validate validates this vscan on demand policy inline scope
func (m *VscanOnDemandPolicyInlineScope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExcludeExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExcludePaths(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncludeExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxFileSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VscanOnDemandPolicyInlineScope) validateExcludeExtensions(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludeExtensions) { // not required
		return nil
	}

	iExcludeExtensionsSize := int64(len(m.ExcludeExtensions))

	if err := validate.MinItems("scope"+"."+"exclude_extensions", "body", iExcludeExtensionsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("scope"+"."+"exclude_extensions", "body", iExcludeExtensionsSize, 300); err != nil {
		return err
	}

	return nil
}

func (m *VscanOnDemandPolicyInlineScope) validateExcludePaths(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludePaths) { // not required
		return nil
	}

	iExcludePathsSize := int64(len(m.ExcludePaths))

	if err := validate.MinItems("scope"+"."+"exclude_paths", "body", iExcludePathsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("scope"+"."+"exclude_paths", "body", iExcludePathsSize, 100); err != nil {
		return err
	}

	for i := 0; i < len(m.ExcludePaths); i++ {
		if swag.IsZero(m.ExcludePaths[i]) { // not required
			continue
		}

		if err := validate.MinLength("scope"+"."+"exclude_paths"+"."+strconv.Itoa(i), "body", *m.ExcludePaths[i], 1); err != nil {
			return err
		}

		if err := validate.MaxLength("scope"+"."+"exclude_paths"+"."+strconv.Itoa(i), "body", *m.ExcludePaths[i], 255); err != nil {
			return err
		}

	}

	return nil
}

func (m *VscanOnDemandPolicyInlineScope) validateIncludeExtensions(formats strfmt.Registry) error {
	if swag.IsZero(m.IncludeExtensions) { // not required
		return nil
	}

	iIncludeExtensionsSize := int64(len(m.IncludeExtensions))

	if err := validate.MinItems("scope"+"."+"include_extensions", "body", iIncludeExtensionsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("scope"+"."+"include_extensions", "body", iIncludeExtensionsSize, 300); err != nil {
		return err
	}

	return nil
}

func (m *VscanOnDemandPolicyInlineScope) validateMaxFileSize(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxFileSize) { // not required
		return nil
	}

	if err := validate.MinimumInt("scope"+"."+"max_file_size", "body", *m.MaxFileSize, 1024, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("scope"+"."+"max_file_size", "body", *m.MaxFileSize, 1.099511627776e+12, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vscan on demand policy inline scope based on context it is used
func (m *VscanOnDemandPolicyInlineScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VscanOnDemandPolicyInlineScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VscanOnDemandPolicyInlineScope) UnmarshalBinary(b []byte) error {
	var res VscanOnDemandPolicyInlineScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
