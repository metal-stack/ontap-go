// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/metal-stack/ontap-go/api/models"
)

// SecurityKeyManagerMigrateReader is a Reader for the SecurityKeyManagerMigrate structure.
type SecurityKeyManagerMigrateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeyManagerMigrateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSecurityKeyManagerMigrateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSecurityKeyManagerMigrateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeyManagerMigrateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeyManagerMigrateCreated creates a SecurityKeyManagerMigrateCreated with default headers values
func NewSecurityKeyManagerMigrateCreated() *SecurityKeyManagerMigrateCreated {
	return &SecurityKeyManagerMigrateCreated{}
}

/*
SecurityKeyManagerMigrateCreated describes a response with status code 201, with default header values.

Created
*/
type SecurityKeyManagerMigrateCreated struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this security key manager migrate created response has a 2xx status code
func (o *SecurityKeyManagerMigrateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security key manager migrate created response has a 3xx status code
func (o *SecurityKeyManagerMigrateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security key manager migrate created response has a 4xx status code
func (o *SecurityKeyManagerMigrateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this security key manager migrate created response has a 5xx status code
func (o *SecurityKeyManagerMigrateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this security key manager migrate created response a status code equal to that given
func (o *SecurityKeyManagerMigrateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the security key manager migrate created response
func (o *SecurityKeyManagerMigrateCreated) Code() int {
	return 201
}

func (o *SecurityKeyManagerMigrateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{source.uuid}/migrate][%d] securityKeyManagerMigrateCreated %s", 201, payload)
}

func (o *SecurityKeyManagerMigrateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{source.uuid}/migrate][%d] securityKeyManagerMigrateCreated %s", 201, payload)
}

func (o *SecurityKeyManagerMigrateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SecurityKeyManagerMigrateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeyManagerMigrateAccepted creates a SecurityKeyManagerMigrateAccepted with default headers values
func NewSecurityKeyManagerMigrateAccepted() *SecurityKeyManagerMigrateAccepted {
	return &SecurityKeyManagerMigrateAccepted{}
}

/*
SecurityKeyManagerMigrateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SecurityKeyManagerMigrateAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this security key manager migrate accepted response has a 2xx status code
func (o *SecurityKeyManagerMigrateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security key manager migrate accepted response has a 3xx status code
func (o *SecurityKeyManagerMigrateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security key manager migrate accepted response has a 4xx status code
func (o *SecurityKeyManagerMigrateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this security key manager migrate accepted response has a 5xx status code
func (o *SecurityKeyManagerMigrateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this security key manager migrate accepted response a status code equal to that given
func (o *SecurityKeyManagerMigrateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the security key manager migrate accepted response
func (o *SecurityKeyManagerMigrateAccepted) Code() int {
	return 202
}

func (o *SecurityKeyManagerMigrateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{source.uuid}/migrate][%d] securityKeyManagerMigrateAccepted %s", 202, payload)
}

func (o *SecurityKeyManagerMigrateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{source.uuid}/migrate][%d] securityKeyManagerMigrateAccepted %s", 202, payload)
}

func (o *SecurityKeyManagerMigrateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SecurityKeyManagerMigrateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeyManagerMigrateDefault creates a SecurityKeyManagerMigrateDefault with default headers values
func NewSecurityKeyManagerMigrateDefault(code int) *SecurityKeyManagerMigrateDefault {
	return &SecurityKeyManagerMigrateDefault{
		_statusCode: code,
	}
}

/*
	SecurityKeyManagerMigrateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536886 | The specified migration option is not supported in this release. |
| 65536959 | The source-uuid and UUID must be different values. |
| 65536968 | Check that all nodes of the cluster are healthy and retry the operation. |
| 65537117 | The migrate operation cannot be started because a UUID cannot be converted to an SVM name. |
| 65537117 | Cannot start migration because a key manager referenced by a provided UUID does not exist. |
| 65537551 | Top-level internal key protection key (KEK) is unavailable on one or more nodes. |
| 65537552 | Embedded KMIP server status is not available. |
| 65537564 | Check that the Azure Key Vault Service is healthy and retry the operation. |
| 65537720 | Failed to configure the Google Cloud Key Management Service for an SVM because a key manager is already configured. |
| 65537736 | Check that the Google Cloud Key Management Service is healthy and retry the operation. |
| 65538107 | Key migration to an IBM Key Lore key manager is not supported. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityKeyManagerMigrateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security key manager migrate default response has a 2xx status code
func (o *SecurityKeyManagerMigrateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security key manager migrate default response has a 3xx status code
func (o *SecurityKeyManagerMigrateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security key manager migrate default response has a 4xx status code
func (o *SecurityKeyManagerMigrateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security key manager migrate default response has a 5xx status code
func (o *SecurityKeyManagerMigrateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security key manager migrate default response a status code equal to that given
func (o *SecurityKeyManagerMigrateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security key manager migrate default response
func (o *SecurityKeyManagerMigrateDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeyManagerMigrateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{source.uuid}/migrate][%d] security_key_manager_migrate default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerMigrateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{source.uuid}/migrate][%d] security_key_manager_migrate default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerMigrateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeyManagerMigrateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SecurityKeyManagerMigrateBody Migration destination key manager UUID
swagger:model SecurityKeyManagerMigrateBody
*/
type SecurityKeyManagerMigrateBody struct {

	// links
	Links *SecurityKeyManagerMigrateParamsBodyLinks `json:"_links,omitempty" yaml:"_links,omitempty"`

	// Key manager UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563434
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this security key manager migrate body
func (o *SecurityKeyManagerMigrateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeyManagerMigrateBody) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security key manager migrate body based on the context it is used
func (o *SecurityKeyManagerMigrateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeyManagerMigrateBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {

		if swag.IsZero(o.Links) { // not required
			return nil
		}

		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SecurityKeyManagerMigrateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityKeyManagerMigrateBody) UnmarshalBinary(b []byte) error {
	var res SecurityKeyManagerMigrateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SecurityKeyManagerMigrateParamsBodyLinks security key manager migrate params body links
swagger:model SecurityKeyManagerMigrateParamsBodyLinks
*/
type SecurityKeyManagerMigrateParamsBodyLinks struct {

	// self
	Self *models.Href `json:"self,omitempty" yaml:"self,omitempty"`
}

// Validate validates this security key manager migrate params body links
func (o *SecurityKeyManagerMigrateParamsBodyLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeyManagerMigrateParamsBodyLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security key manager migrate params body links based on the context it is used
func (o *SecurityKeyManagerMigrateParamsBodyLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeyManagerMigrateParamsBodyLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {

		if swag.IsZero(o.Self) { // not required
			return nil
		}

		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SecurityKeyManagerMigrateParamsBodyLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityKeyManagerMigrateParamsBodyLinks) UnmarshalBinary(b []byte) error {
	var res SecurityKeyManagerMigrateParamsBodyLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}