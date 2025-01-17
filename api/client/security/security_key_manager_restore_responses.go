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

// SecurityKeyManagerRestoreReader is a Reader for the SecurityKeyManagerRestore structure.
type SecurityKeyManagerRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeyManagerRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSecurityKeyManagerRestoreCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSecurityKeyManagerRestoreAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeyManagerRestoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeyManagerRestoreCreated creates a SecurityKeyManagerRestoreCreated with default headers values
func NewSecurityKeyManagerRestoreCreated() *SecurityKeyManagerRestoreCreated {
	return &SecurityKeyManagerRestoreCreated{}
}

/*
SecurityKeyManagerRestoreCreated describes a response with status code 201, with default header values.

Created
*/
type SecurityKeyManagerRestoreCreated struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this security key manager restore created response has a 2xx status code
func (o *SecurityKeyManagerRestoreCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security key manager restore created response has a 3xx status code
func (o *SecurityKeyManagerRestoreCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security key manager restore created response has a 4xx status code
func (o *SecurityKeyManagerRestoreCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this security key manager restore created response has a 5xx status code
func (o *SecurityKeyManagerRestoreCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this security key manager restore created response a status code equal to that given
func (o *SecurityKeyManagerRestoreCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the security key manager restore created response
func (o *SecurityKeyManagerRestoreCreated) Code() int {
	return 201
}

func (o *SecurityKeyManagerRestoreCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{security_key_manager.uuid}/restore][%d] securityKeyManagerRestoreCreated %s", 201, payload)
}

func (o *SecurityKeyManagerRestoreCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{security_key_manager.uuid}/restore][%d] securityKeyManagerRestoreCreated %s", 201, payload)
}

func (o *SecurityKeyManagerRestoreCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SecurityKeyManagerRestoreCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeyManagerRestoreAccepted creates a SecurityKeyManagerRestoreAccepted with default headers values
func NewSecurityKeyManagerRestoreAccepted() *SecurityKeyManagerRestoreAccepted {
	return &SecurityKeyManagerRestoreAccepted{}
}

/*
SecurityKeyManagerRestoreAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SecurityKeyManagerRestoreAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this security key manager restore accepted response has a 2xx status code
func (o *SecurityKeyManagerRestoreAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security key manager restore accepted response has a 3xx status code
func (o *SecurityKeyManagerRestoreAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security key manager restore accepted response has a 4xx status code
func (o *SecurityKeyManagerRestoreAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this security key manager restore accepted response has a 5xx status code
func (o *SecurityKeyManagerRestoreAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this security key manager restore accepted response a status code equal to that given
func (o *SecurityKeyManagerRestoreAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the security key manager restore accepted response
func (o *SecurityKeyManagerRestoreAccepted) Code() int {
	return 202
}

func (o *SecurityKeyManagerRestoreAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{security_key_manager.uuid}/restore][%d] securityKeyManagerRestoreAccepted %s", 202, payload)
}

func (o *SecurityKeyManagerRestoreAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{security_key_manager.uuid}/restore][%d] securityKeyManagerRestoreAccepted %s", 202, payload)
}

func (o *SecurityKeyManagerRestoreAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SecurityKeyManagerRestoreAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeyManagerRestoreDefault creates a SecurityKeyManagerRestoreDefault with default headers values
func NewSecurityKeyManagerRestoreDefault(code int) *SecurityKeyManagerRestoreDefault {
	return &SecurityKeyManagerRestoreDefault{
		_statusCode: code,
	}
}

/*
	SecurityKeyManagerRestoreDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536083 | Internal error. Failed to restore the authentication key. |
| 65536843 | The key management server is not configured for the SVM. |
| 65536855 | Internal error. Failed to restore the volume encryption key. |
| 65538500 | This command only restores keys from primary key servers. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityKeyManagerRestoreDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security key manager restore default response has a 2xx status code
func (o *SecurityKeyManagerRestoreDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security key manager restore default response has a 3xx status code
func (o *SecurityKeyManagerRestoreDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security key manager restore default response has a 4xx status code
func (o *SecurityKeyManagerRestoreDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security key manager restore default response has a 5xx status code
func (o *SecurityKeyManagerRestoreDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security key manager restore default response a status code equal to that given
func (o *SecurityKeyManagerRestoreDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security key manager restore default response
func (o *SecurityKeyManagerRestoreDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeyManagerRestoreDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{security_key_manager.uuid}/restore][%d] security_key_manager_restore default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerRestoreDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{security_key_manager.uuid}/restore][%d] security_key_manager_restore default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerRestoreDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeyManagerRestoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
