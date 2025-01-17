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

// SecurityKeyManagerKeyServersGetReader is a Reader for the SecurityKeyManagerKeyServersGet structure.
type SecurityKeyManagerKeyServersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeyManagerKeyServersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeyManagerKeyServersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeyManagerKeyServersGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeyManagerKeyServersGetOK creates a SecurityKeyManagerKeyServersGetOK with default headers values
func NewSecurityKeyManagerKeyServersGetOK() *SecurityKeyManagerKeyServersGetOK {
	return &SecurityKeyManagerKeyServersGetOK{}
}

/*
SecurityKeyManagerKeyServersGetOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeyManagerKeyServersGetOK struct {
	Payload *models.KeyServer
}

// IsSuccess returns true when this security key manager key servers get o k response has a 2xx status code
func (o *SecurityKeyManagerKeyServersGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security key manager key servers get o k response has a 3xx status code
func (o *SecurityKeyManagerKeyServersGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security key manager key servers get o k response has a 4xx status code
func (o *SecurityKeyManagerKeyServersGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security key manager key servers get o k response has a 5xx status code
func (o *SecurityKeyManagerKeyServersGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security key manager key servers get o k response a status code equal to that given
func (o *SecurityKeyManagerKeyServersGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security key manager key servers get o k response
func (o *SecurityKeyManagerKeyServersGetOK) Code() int {
	return 200
}

func (o *SecurityKeyManagerKeyServersGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{uuid}/key-servers/{server}][%d] securityKeyManagerKeyServersGetOK %s", 200, payload)
}

func (o *SecurityKeyManagerKeyServersGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{uuid}/key-servers/{server}][%d] securityKeyManagerKeyServersGetOK %s", 200, payload)
}

func (o *SecurityKeyManagerKeyServersGetOK) GetPayload() *models.KeyServer {
	return o.Payload
}

func (o *SecurityKeyManagerKeyServersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeyManagerKeyServersGetDefault creates a SecurityKeyManagerKeyServersGetDefault with default headers values
func NewSecurityKeyManagerKeyServersGetDefault(code int) *SecurityKeyManagerKeyServersGetDefault {
	return &SecurityKeyManagerKeyServersGetDefault{
		_statusCode: code,
	}
}

/*
SecurityKeyManagerKeyServersGetDefault describes a response with status code -1, with default header values.

Error
*/
type SecurityKeyManagerKeyServersGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security key manager key servers get default response has a 2xx status code
func (o *SecurityKeyManagerKeyServersGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security key manager key servers get default response has a 3xx status code
func (o *SecurityKeyManagerKeyServersGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security key manager key servers get default response has a 4xx status code
func (o *SecurityKeyManagerKeyServersGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security key manager key servers get default response has a 5xx status code
func (o *SecurityKeyManagerKeyServersGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security key manager key servers get default response a status code equal to that given
func (o *SecurityKeyManagerKeyServersGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security key manager key servers get default response
func (o *SecurityKeyManagerKeyServersGetDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeyManagerKeyServersGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{uuid}/key-servers/{server}][%d] security_key_manager_key_servers_get default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerKeyServersGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{uuid}/key-servers/{server}][%d] security_key_manager_key_servers_get default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerKeyServersGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeyManagerKeyServersGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}