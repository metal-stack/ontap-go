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

// SecurityOauth2GlobalModifyReader is a Reader for the SecurityOauth2GlobalModify structure.
type SecurityOauth2GlobalModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityOauth2GlobalModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityOauth2GlobalModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityOauth2GlobalModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityOauth2GlobalModifyOK creates a SecurityOauth2GlobalModifyOK with default headers values
func NewSecurityOauth2GlobalModifyOK() *SecurityOauth2GlobalModifyOK {
	return &SecurityOauth2GlobalModifyOK{}
}

/*
SecurityOauth2GlobalModifyOK describes a response with status code 200, with default header values.

OK
*/
type SecurityOauth2GlobalModifyOK struct {
}

// IsSuccess returns true when this security oauth2 global modify o k response has a 2xx status code
func (o *SecurityOauth2GlobalModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security oauth2 global modify o k response has a 3xx status code
func (o *SecurityOauth2GlobalModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security oauth2 global modify o k response has a 4xx status code
func (o *SecurityOauth2GlobalModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security oauth2 global modify o k response has a 5xx status code
func (o *SecurityOauth2GlobalModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security oauth2 global modify o k response a status code equal to that given
func (o *SecurityOauth2GlobalModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security oauth2 global modify o k response
func (o *SecurityOauth2GlobalModifyOK) Code() int {
	return 200
}

func (o *SecurityOauth2GlobalModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /security/authentication/cluster/oauth2][%d] securityOauth2GlobalModifyOK", 200)
}

func (o *SecurityOauth2GlobalModifyOK) String() string {
	return fmt.Sprintf("[PATCH /security/authentication/cluster/oauth2][%d] securityOauth2GlobalModifyOK", 200)
}

func (o *SecurityOauth2GlobalModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecurityOauth2GlobalModifyDefault creates a SecurityOauth2GlobalModifyDefault with default headers values
func NewSecurityOauth2GlobalModifyDefault(code int) *SecurityOauth2GlobalModifyDefault {
	return &SecurityOauth2GlobalModifyDefault{
		_statusCode: code,
	}
}

/*
SecurityOauth2GlobalModifyDefault describes a response with status code -1, with default header values.

SecurityOauth2GlobalModifyDefault security oauth2 global modify default
*/
type SecurityOauth2GlobalModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security oauth2 global modify default response has a 2xx status code
func (o *SecurityOauth2GlobalModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security oauth2 global modify default response has a 3xx status code
func (o *SecurityOauth2GlobalModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security oauth2 global modify default response has a 4xx status code
func (o *SecurityOauth2GlobalModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security oauth2 global modify default response has a 5xx status code
func (o *SecurityOauth2GlobalModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security oauth2 global modify default response a status code equal to that given
func (o *SecurityOauth2GlobalModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security oauth2 global modify default response
func (o *SecurityOauth2GlobalModifyDefault) Code() int {
	return o._statusCode
}

func (o *SecurityOauth2GlobalModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/authentication/cluster/oauth2][%d] security_oauth2_global_modify default %s", o._statusCode, payload)
}

func (o *SecurityOauth2GlobalModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/authentication/cluster/oauth2][%d] security_oauth2_global_modify default %s", o._statusCode, payload)
}

func (o *SecurityOauth2GlobalModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityOauth2GlobalModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
