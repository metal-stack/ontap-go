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

// SecurityOauth2GetReader is a Reader for the SecurityOauth2Get structure.
type SecurityOauth2GetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityOauth2GetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityOauth2GetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityOauth2GetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityOauth2GetOK creates a SecurityOauth2GetOK with default headers values
func NewSecurityOauth2GetOK() *SecurityOauth2GetOK {
	return &SecurityOauth2GetOK{}
}

/*
SecurityOauth2GetOK describes a response with status code 200, with default header values.

OK
*/
type SecurityOauth2GetOK struct {
	Payload *models.SecurityOauth2
}

// IsSuccess returns true when this security oauth2 get o k response has a 2xx status code
func (o *SecurityOauth2GetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security oauth2 get o k response has a 3xx status code
func (o *SecurityOauth2GetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security oauth2 get o k response has a 4xx status code
func (o *SecurityOauth2GetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security oauth2 get o k response has a 5xx status code
func (o *SecurityOauth2GetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security oauth2 get o k response a status code equal to that given
func (o *SecurityOauth2GetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security oauth2 get o k response
func (o *SecurityOauth2GetOK) Code() int {
	return 200
}

func (o *SecurityOauth2GetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/oauth2/clients/{name}][%d] securityOauth2GetOK %s", 200, payload)
}

func (o *SecurityOauth2GetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/oauth2/clients/{name}][%d] securityOauth2GetOK %s", 200, payload)
}

func (o *SecurityOauth2GetOK) GetPayload() *models.SecurityOauth2 {
	return o.Payload
}

func (o *SecurityOauth2GetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityOauth2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityOauth2GetDefault creates a SecurityOauth2GetDefault with default headers values
func NewSecurityOauth2GetDefault(code int) *SecurityOauth2GetDefault {
	return &SecurityOauth2GetDefault{
		_statusCode: code,
	}
}

/*
SecurityOauth2GetDefault describes a response with status code -1, with default header values.

Error
*/
type SecurityOauth2GetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security oauth2 get default response has a 2xx status code
func (o *SecurityOauth2GetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security oauth2 get default response has a 3xx status code
func (o *SecurityOauth2GetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security oauth2 get default response has a 4xx status code
func (o *SecurityOauth2GetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security oauth2 get default response has a 5xx status code
func (o *SecurityOauth2GetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security oauth2 get default response a status code equal to that given
func (o *SecurityOauth2GetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security oauth2 get default response
func (o *SecurityOauth2GetDefault) Code() int {
	return o._statusCode
}

func (o *SecurityOauth2GetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/oauth2/clients/{name}][%d] security_oauth2_get default %s", o._statusCode, payload)
}

func (o *SecurityOauth2GetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/oauth2/clients/{name}][%d] security_oauth2_get default %s", o._statusCode, payload)
}

func (o *SecurityOauth2GetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityOauth2GetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}