// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MongoDbOnSan MongoDB using SAN.
//
// swagger:model mongo_db_on_san
type MongoDbOnSan struct {

	// dataset
	// Required: true
	Dataset *MongoDbOnSanDataset `json:"dataset" yaml:"dataset"`

	// The list of initiator groups to create.
	// Max Items: 20
	// Min Items: 0
	NewIgroups []*MongoDbOnSanNewIgroups `json:"new_igroups" yaml:"new_igroups"`

	// The name of the host OS running the application.
	// Enum: ["hyper_v","linux","solaris","solaris_efi","vmware","windows","windows_2008","windows_gpt","xen"]
	OsType *string `json:"os_type,omitempty" yaml:"os_type,omitempty"`

	// The initiator group for the primary.
	// Required: true
	// Max Length: 96
	// Min Length: 1
	PrimaryIgroupName *string `json:"primary_igroup_name" yaml:"primary_igroup_name"`

	// protection type
	ProtectionType *MongoDbOnSanProtectionType `json:"protection_type,omitempty" yaml:"protection_type,omitempty"`

	// secondary igroups
	// Max Items: 19
	// Min Items: 0
	SecondaryIgroups []*MongoDbOnSanSecondaryIgroupsItems0 `json:"secondary_igroups" yaml:"secondary_igroups"`
}

// Validate validates this mongo db on san
func (m *MongoDbOnSan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewIgroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIgroupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryIgroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MongoDbOnSan) validateDataset(formats strfmt.Registry) error {

	if err := validate.Required("dataset", "body", m.Dataset); err != nil {
		return err
	}

	if m.Dataset != nil {
		if err := m.Dataset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataset")
			}
			return err
		}
	}

	return nil
}

func (m *MongoDbOnSan) validateNewIgroups(formats strfmt.Registry) error {
	if swag.IsZero(m.NewIgroups) { // not required
		return nil
	}

	iNewIgroupsSize := int64(len(m.NewIgroups))

	if err := validate.MinItems("new_igroups", "body", iNewIgroupsSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("new_igroups", "body", iNewIgroupsSize, 20); err != nil {
		return err
	}

	for i := 0; i < len(m.NewIgroups); i++ {
		if swag.IsZero(m.NewIgroups[i]) { // not required
			continue
		}

		if m.NewIgroups[i] != nil {
			if err := m.NewIgroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("new_igroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("new_igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var mongoDbOnSanTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hyper_v","linux","solaris","solaris_efi","vmware","windows","windows_2008","windows_gpt","xen"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mongoDbOnSanTypeOsTypePropEnum = append(mongoDbOnSanTypeOsTypePropEnum, v)
	}
}

const (

	// MongoDbOnSanOsTypeHyperv captures enum value "hyper_v"
	MongoDbOnSanOsTypeHyperv string = "hyper_v"

	// MongoDbOnSanOsTypeLinux captures enum value "linux"
	MongoDbOnSanOsTypeLinux string = "linux"

	// MongoDbOnSanOsTypeSolaris captures enum value "solaris"
	MongoDbOnSanOsTypeSolaris string = "solaris"

	// MongoDbOnSanOsTypeSolarisEfi captures enum value "solaris_efi"
	MongoDbOnSanOsTypeSolarisEfi string = "solaris_efi"

	// MongoDbOnSanOsTypeVmware captures enum value "vmware"
	MongoDbOnSanOsTypeVmware string = "vmware"

	// MongoDbOnSanOsTypeWindows captures enum value "windows"
	MongoDbOnSanOsTypeWindows string = "windows"

	// MongoDbOnSanOsTypeWindows2008 captures enum value "windows_2008"
	MongoDbOnSanOsTypeWindows2008 string = "windows_2008"

	// MongoDbOnSanOsTypeWindowsGpt captures enum value "windows_gpt"
	MongoDbOnSanOsTypeWindowsGpt string = "windows_gpt"

	// MongoDbOnSanOsTypeXen captures enum value "xen"
	MongoDbOnSanOsTypeXen string = "xen"
)

// prop value enum
func (m *MongoDbOnSan) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mongoDbOnSanTypeOsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MongoDbOnSan) validateOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsTypeEnum("os_type", "body", *m.OsType); err != nil {
		return err
	}

	return nil
}

func (m *MongoDbOnSan) validatePrimaryIgroupName(formats strfmt.Registry) error {

	if err := validate.Required("primary_igroup_name", "body", m.PrimaryIgroupName); err != nil {
		return err
	}

	if err := validate.MinLength("primary_igroup_name", "body", *m.PrimaryIgroupName, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("primary_igroup_name", "body", *m.PrimaryIgroupName, 96); err != nil {
		return err
	}

	return nil
}

func (m *MongoDbOnSan) validateProtectionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionType) { // not required
		return nil
	}

	if m.ProtectionType != nil {
		if err := m.ProtectionType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protection_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protection_type")
			}
			return err
		}
	}

	return nil
}

func (m *MongoDbOnSan) validateSecondaryIgroups(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryIgroups) { // not required
		return nil
	}

	iSecondaryIgroupsSize := int64(len(m.SecondaryIgroups))

	if err := validate.MinItems("secondary_igroups", "body", iSecondaryIgroupsSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("secondary_igroups", "body", iSecondaryIgroupsSize, 19); err != nil {
		return err
	}

	for i := 0; i < len(m.SecondaryIgroups); i++ {
		if swag.IsZero(m.SecondaryIgroups[i]) { // not required
			continue
		}

		if m.SecondaryIgroups[i] != nil {
			if err := m.SecondaryIgroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secondary_igroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("secondary_igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this mongo db on san based on the context it is used
func (m *MongoDbOnSan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNewIgroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtectionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecondaryIgroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MongoDbOnSan) contextValidateDataset(ctx context.Context, formats strfmt.Registry) error {

	if m.Dataset != nil {

		if err := m.Dataset.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataset")
			}
			return err
		}
	}

	return nil
}

func (m *MongoDbOnSan) contextValidateNewIgroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NewIgroups); i++ {

		if m.NewIgroups[i] != nil {

			if swag.IsZero(m.NewIgroups[i]) { // not required
				return nil
			}

			if err := m.NewIgroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("new_igroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("new_igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MongoDbOnSan) contextValidateProtectionType(ctx context.Context, formats strfmt.Registry) error {

	if m.ProtectionType != nil {

		if swag.IsZero(m.ProtectionType) { // not required
			return nil
		}

		if err := m.ProtectionType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protection_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protection_type")
			}
			return err
		}
	}

	return nil
}

func (m *MongoDbOnSan) contextValidateSecondaryIgroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SecondaryIgroups); i++ {

		if m.SecondaryIgroups[i] != nil {

			if swag.IsZero(m.SecondaryIgroups[i]) { // not required
				return nil
			}

			if err := m.SecondaryIgroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secondary_igroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("secondary_igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MongoDbOnSan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoDbOnSan) UnmarshalBinary(b []byte) error {
	var res MongoDbOnSan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MongoDbOnSanDataset mongo db on san dataset
//
// swagger:model MongoDbOnSanDataset
type MongoDbOnSanDataset struct {

	// The number of storage elements (LUNs for SAN) of the database to maintain.  Must be an even number between 2 and 16.  Odd numbers will be rounded up to the next even number within range.
	// Maximum: 16
	// Minimum: 2
	ElementCount int64 `json:"element_count,omitempty" yaml:"element_count,omitempty"`

	// The number of data bearing members of the replicaset, including 1 primary and at least 1 secondary.
	// Maximum: 20
	// Minimum: 2
	ReplicationFactor int64 `json:"replication_factor,omitempty" yaml:"replication_factor,omitempty"`

	// The size of the database. Usage: {&lt;integer&gt;[KB|MB|GB|TB|PB]}
	// Required: true
	Size *int64 `json:"size" yaml:"size"`

	// storage service
	StorageService *MongoDbOnSanDatasetStorageService `json:"storage_service,omitempty" yaml:"storage_service,omitempty"`
}

// Validate validates this mongo db on san dataset
func (m *MongoDbOnSanDataset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElementCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicationFactor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MongoDbOnSanDataset) validateElementCount(formats strfmt.Registry) error {
	if swag.IsZero(m.ElementCount) { // not required
		return nil
	}

	if err := validate.MinimumInt("dataset"+"."+"element_count", "body", m.ElementCount, 2, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("dataset"+"."+"element_count", "body", m.ElementCount, 16, false); err != nil {
		return err
	}

	return nil
}

func (m *MongoDbOnSanDataset) validateReplicationFactor(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplicationFactor) { // not required
		return nil
	}

	if err := validate.MinimumInt("dataset"+"."+"replication_factor", "body", m.ReplicationFactor, 2, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("dataset"+"."+"replication_factor", "body", m.ReplicationFactor, 20, false); err != nil {
		return err
	}

	return nil
}

func (m *MongoDbOnSanDataset) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("dataset"+"."+"size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *MongoDbOnSanDataset) validateStorageService(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageService) { // not required
		return nil
	}

	if m.StorageService != nil {
		if err := m.StorageService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataset" + "." + "storage_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataset" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mongo db on san dataset based on the context it is used
func (m *MongoDbOnSanDataset) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MongoDbOnSanDataset) contextValidateStorageService(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageService != nil {

		if swag.IsZero(m.StorageService) { // not required
			return nil
		}

		if err := m.StorageService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataset" + "." + "storage_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataset" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MongoDbOnSanDataset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoDbOnSanDataset) UnmarshalBinary(b []byte) error {
	var res MongoDbOnSanDataset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MongoDbOnSanDatasetStorageService mongo db on san dataset storage service
//
// swagger:model MongoDbOnSanDatasetStorageService
type MongoDbOnSanDatasetStorageService struct {

	// The storage service of the database.
	// Enum: ["extreme","performance","value"]
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this mongo db on san dataset storage service
func (m *MongoDbOnSanDatasetStorageService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var mongoDbOnSanDatasetStorageServiceTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extreme","performance","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mongoDbOnSanDatasetStorageServiceTypeNamePropEnum = append(mongoDbOnSanDatasetStorageServiceTypeNamePropEnum, v)
	}
}

const (

	// MongoDbOnSanDatasetStorageServiceNameExtreme captures enum value "extreme"
	MongoDbOnSanDatasetStorageServiceNameExtreme string = "extreme"

	// MongoDbOnSanDatasetStorageServiceNamePerformance captures enum value "performance"
	MongoDbOnSanDatasetStorageServiceNamePerformance string = "performance"

	// MongoDbOnSanDatasetStorageServiceNameValue captures enum value "value"
	MongoDbOnSanDatasetStorageServiceNameValue string = "value"
)

// prop value enum
func (m *MongoDbOnSanDatasetStorageService) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mongoDbOnSanDatasetStorageServiceTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MongoDbOnSanDatasetStorageService) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("dataset"+"."+"storage_service"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mongo db on san dataset storage service based on context it is used
func (m *MongoDbOnSanDatasetStorageService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MongoDbOnSanDatasetStorageService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoDbOnSanDatasetStorageService) UnmarshalBinary(b []byte) error {
	var res MongoDbOnSanDatasetStorageService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MongoDbOnSanProtectionType mongo db on san protection type
//
// swagger:model MongoDbOnSanProtectionType
type MongoDbOnSanProtectionType struct {

	// The local RPO of the application.
	// Enum: ["hourly","none"]
	LocalRpo string `json:"local_rpo,omitempty" yaml:"local_rpo,omitempty"`

	// The remote RPO of the application.
	// Enum: ["none","zero"]
	RemoteRpo string `json:"remote_rpo,omitempty" yaml:"remote_rpo,omitempty"`
}

// Validate validates this mongo db on san protection type
func (m *MongoDbOnSanProtectionType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocalRpo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteRpo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var mongoDbOnSanProtectionTypeTypeLocalRpoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hourly","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mongoDbOnSanProtectionTypeTypeLocalRpoPropEnum = append(mongoDbOnSanProtectionTypeTypeLocalRpoPropEnum, v)
	}
}

const (

	// MongoDbOnSanProtectionTypeLocalRpoHourly captures enum value "hourly"
	MongoDbOnSanProtectionTypeLocalRpoHourly string = "hourly"

	// MongoDbOnSanProtectionTypeLocalRpoNone captures enum value "none"
	MongoDbOnSanProtectionTypeLocalRpoNone string = "none"
)

// prop value enum
func (m *MongoDbOnSanProtectionType) validateLocalRpoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mongoDbOnSanProtectionTypeTypeLocalRpoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MongoDbOnSanProtectionType) validateLocalRpo(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalRpo) { // not required
		return nil
	}

	// value enum
	if err := m.validateLocalRpoEnum("protection_type"+"."+"local_rpo", "body", m.LocalRpo); err != nil {
		return err
	}

	return nil
}

var mongoDbOnSanProtectionTypeTypeRemoteRpoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","zero"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mongoDbOnSanProtectionTypeTypeRemoteRpoPropEnum = append(mongoDbOnSanProtectionTypeTypeRemoteRpoPropEnum, v)
	}
}

const (

	// MongoDbOnSanProtectionTypeRemoteRpoNone captures enum value "none"
	MongoDbOnSanProtectionTypeRemoteRpoNone string = "none"

	// MongoDbOnSanProtectionTypeRemoteRpoZero captures enum value "zero"
	MongoDbOnSanProtectionTypeRemoteRpoZero string = "zero"
)

// prop value enum
func (m *MongoDbOnSanProtectionType) validateRemoteRpoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mongoDbOnSanProtectionTypeTypeRemoteRpoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MongoDbOnSanProtectionType) validateRemoteRpo(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteRpo) { // not required
		return nil
	}

	// value enum
	if err := m.validateRemoteRpoEnum("protection_type"+"."+"remote_rpo", "body", m.RemoteRpo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mongo db on san protection type based on context it is used
func (m *MongoDbOnSanProtectionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MongoDbOnSanProtectionType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoDbOnSanProtectionType) UnmarshalBinary(b []byte) error {
	var res MongoDbOnSanProtectionType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MongoDbOnSanSecondaryIgroupsItems0 mongo db on san secondary igroups items0
//
// swagger:model MongoDbOnSanSecondaryIgroupsItems0
type MongoDbOnSanSecondaryIgroupsItems0 struct {

	// The name of the initiator group for each secondary.
	// Max Length: 96
	// Min Length: 1
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Validate validates this mongo db on san secondary igroups items0
func (m *MongoDbOnSanSecondaryIgroupsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MongoDbOnSanSecondaryIgroupsItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 96); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mongo db on san secondary igroups items0 based on context it is used
func (m *MongoDbOnSanSecondaryIgroupsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MongoDbOnSanSecondaryIgroupsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoDbOnSanSecondaryIgroupsItems0) UnmarshalBinary(b []byte) error {
	var res MongoDbOnSanSecondaryIgroupsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
