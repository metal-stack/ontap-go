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

// SecurityKeystoreDeleteReader is a Reader for the SecurityKeystoreDelete structure.
type SecurityKeystoreDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeystoreDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeystoreDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeystoreDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeystoreDeleteOK creates a SecurityKeystoreDeleteOK with default headers values
func NewSecurityKeystoreDeleteOK() *SecurityKeystoreDeleteOK {
	return &SecurityKeystoreDeleteOK{}
}

/*
SecurityKeystoreDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeystoreDeleteOK struct {
}

// IsSuccess returns true when this security keystore delete o k response has a 2xx status code
func (o *SecurityKeystoreDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security keystore delete o k response has a 3xx status code
func (o *SecurityKeystoreDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security keystore delete o k response has a 4xx status code
func (o *SecurityKeystoreDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security keystore delete o k response has a 5xx status code
func (o *SecurityKeystoreDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security keystore delete o k response a status code equal to that given
func (o *SecurityKeystoreDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security keystore delete o k response
func (o *SecurityKeystoreDeleteOK) Code() int {
	return 200
}

func (o *SecurityKeystoreDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/key-stores/{uuid}][%d] securityKeystoreDeleteOK", 200)
}

func (o *SecurityKeystoreDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/key-stores/{uuid}][%d] securityKeystoreDeleteOK", 200)
}

func (o *SecurityKeystoreDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecurityKeystoreDeleteDefault creates a SecurityKeystoreDeleteDefault with default headers values
func NewSecurityKeystoreDeleteDefault(code int) *SecurityKeystoreDeleteDefault {
	return &SecurityKeystoreDeleteDefault{
		_statusCode: code,
	}
}

/*
	SecurityKeystoreDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262155 | This operation requires an effective cluster version of 9.14.0 or later. |
| 65538905 | The keystore configuration is currently enabled and cannot be deleted. |
| 65538907 | The method is not yet supported for deleting the given UUID's type of configuration. |
| 65538908 | The specified keystore configuration UUID either does not exist or corresponds to a keystore configuration that is not supported by this operation. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityKeystoreDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security keystore delete default response has a 2xx status code
func (o *SecurityKeystoreDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security keystore delete default response has a 3xx status code
func (o *SecurityKeystoreDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security keystore delete default response has a 4xx status code
func (o *SecurityKeystoreDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security keystore delete default response has a 5xx status code
func (o *SecurityKeystoreDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security keystore delete default response a status code equal to that given
func (o *SecurityKeystoreDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security keystore delete default response
func (o *SecurityKeystoreDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeystoreDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/key-stores/{uuid}][%d] security_keystore_delete default %s", o._statusCode, payload)
}

func (o *SecurityKeystoreDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/key-stores/{uuid}][%d] security_keystore_delete default %s", o._statusCode, payload)
}

func (o *SecurityKeystoreDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeystoreDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
