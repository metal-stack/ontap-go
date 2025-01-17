// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/ontap-go/api/models"
)

// SecurityAuditLogCollectionGetReader is a Reader for the SecurityAuditLogCollectionGet structure.
type SecurityAuditLogCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityAuditLogCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityAuditLogCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityAuditLogCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityAuditLogCollectionGetOK creates a SecurityAuditLogCollectionGetOK with default headers values
func NewSecurityAuditLogCollectionGetOK() *SecurityAuditLogCollectionGetOK {
	return &SecurityAuditLogCollectionGetOK{}
}

/*
SecurityAuditLogCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SecurityAuditLogCollectionGetOK struct {
	Payload *models.SecurityAuditLogResponse
}

// IsSuccess returns true when this security audit log collection get o k response has a 2xx status code
func (o *SecurityAuditLogCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security audit log collection get o k response has a 3xx status code
func (o *SecurityAuditLogCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security audit log collection get o k response has a 4xx status code
func (o *SecurityAuditLogCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security audit log collection get o k response has a 5xx status code
func (o *SecurityAuditLogCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security audit log collection get o k response a status code equal to that given
func (o *SecurityAuditLogCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security audit log collection get o k response
func (o *SecurityAuditLogCollectionGetOK) Code() int {
	return 200
}

func (o *SecurityAuditLogCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/audit/messages][%d] securityAuditLogCollectionGetOK %s", 200, payload)
}

func (o *SecurityAuditLogCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/audit/messages][%d] securityAuditLogCollectionGetOK %s", 200, payload)
}

func (o *SecurityAuditLogCollectionGetOK) GetPayload() *models.SecurityAuditLogResponse {
	return o.Payload
}

func (o *SecurityAuditLogCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityAuditLogResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityAuditLogCollectionGetDefault creates a SecurityAuditLogCollectionGetDefault with default headers values
func NewSecurityAuditLogCollectionGetDefault(code int) *SecurityAuditLogCollectionGetDefault {
	return &SecurityAuditLogCollectionGetDefault{
		_statusCode: code,
	}
}

/*
SecurityAuditLogCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SecurityAuditLogCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security audit log collection get default response has a 2xx status code
func (o *SecurityAuditLogCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security audit log collection get default response has a 3xx status code
func (o *SecurityAuditLogCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security audit log collection get default response has a 4xx status code
func (o *SecurityAuditLogCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security audit log collection get default response has a 5xx status code
func (o *SecurityAuditLogCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security audit log collection get default response a status code equal to that given
func (o *SecurityAuditLogCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security audit log collection get default response
func (o *SecurityAuditLogCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SecurityAuditLogCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/audit/messages][%d] security_audit_log_collection_get default %s", o._statusCode, payload)
}

func (o *SecurityAuditLogCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/audit/messages][%d] security_audit_log_collection_get default %s", o._statusCode, payload)
}

func (o *SecurityAuditLogCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityAuditLogCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
