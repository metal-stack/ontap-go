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

// SecurityKeystoreCollectionGetReader is a Reader for the SecurityKeystoreCollectionGet structure.
type SecurityKeystoreCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeystoreCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeystoreCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeystoreCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeystoreCollectionGetOK creates a SecurityKeystoreCollectionGetOK with default headers values
func NewSecurityKeystoreCollectionGetOK() *SecurityKeystoreCollectionGetOK {
	return &SecurityKeystoreCollectionGetOK{}
}

/*
SecurityKeystoreCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeystoreCollectionGetOK struct {
	Payload *models.SecurityKeystoreResponse
}

// IsSuccess returns true when this security keystore collection get o k response has a 2xx status code
func (o *SecurityKeystoreCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security keystore collection get o k response has a 3xx status code
func (o *SecurityKeystoreCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security keystore collection get o k response has a 4xx status code
func (o *SecurityKeystoreCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security keystore collection get o k response has a 5xx status code
func (o *SecurityKeystoreCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security keystore collection get o k response a status code equal to that given
func (o *SecurityKeystoreCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security keystore collection get o k response
func (o *SecurityKeystoreCollectionGetOK) Code() int {
	return 200
}

func (o *SecurityKeystoreCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-stores][%d] securityKeystoreCollectionGetOK %s", 200, payload)
}

func (o *SecurityKeystoreCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-stores][%d] securityKeystoreCollectionGetOK %s", 200, payload)
}

func (o *SecurityKeystoreCollectionGetOK) GetPayload() *models.SecurityKeystoreResponse {
	return o.Payload
}

func (o *SecurityKeystoreCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityKeystoreResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeystoreCollectionGetDefault creates a SecurityKeystoreCollectionGetDefault with default headers values
func NewSecurityKeystoreCollectionGetDefault(code int) *SecurityKeystoreCollectionGetDefault {
	return &SecurityKeystoreCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	SecurityKeystoreCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 13434920 | SVM does not exist. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityKeystoreCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security keystore collection get default response has a 2xx status code
func (o *SecurityKeystoreCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security keystore collection get default response has a 3xx status code
func (o *SecurityKeystoreCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security keystore collection get default response has a 4xx status code
func (o *SecurityKeystoreCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security keystore collection get default response has a 5xx status code
func (o *SecurityKeystoreCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security keystore collection get default response a status code equal to that given
func (o *SecurityKeystoreCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security keystore collection get default response
func (o *SecurityKeystoreCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeystoreCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-stores][%d] security_keystore_collection_get default %s", o._statusCode, payload)
}

func (o *SecurityKeystoreCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-stores][%d] security_keystore_collection_get default %s", o._statusCode, payload)
}

func (o *SecurityKeystoreCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeystoreCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
