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

// SecurityKeyManagerKeyServersCreateReader is a Reader for the SecurityKeyManagerKeyServersCreate structure.
type SecurityKeyManagerKeyServersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeyManagerKeyServersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSecurityKeyManagerKeyServersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeyManagerKeyServersCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeyManagerKeyServersCreateCreated creates a SecurityKeyManagerKeyServersCreateCreated with default headers values
func NewSecurityKeyManagerKeyServersCreateCreated() *SecurityKeyManagerKeyServersCreateCreated {
	return &SecurityKeyManagerKeyServersCreateCreated{}
}

/*
SecurityKeyManagerKeyServersCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SecurityKeyManagerKeyServersCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.KeyServerResponse
}

// IsSuccess returns true when this security key manager key servers create created response has a 2xx status code
func (o *SecurityKeyManagerKeyServersCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security key manager key servers create created response has a 3xx status code
func (o *SecurityKeyManagerKeyServersCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security key manager key servers create created response has a 4xx status code
func (o *SecurityKeyManagerKeyServersCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this security key manager key servers create created response has a 5xx status code
func (o *SecurityKeyManagerKeyServersCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this security key manager key servers create created response a status code equal to that given
func (o *SecurityKeyManagerKeyServersCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the security key manager key servers create created response
func (o *SecurityKeyManagerKeyServersCreateCreated) Code() int {
	return 201
}

func (o *SecurityKeyManagerKeyServersCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{uuid}/key-servers][%d] securityKeyManagerKeyServersCreateCreated %s", 201, payload)
}

func (o *SecurityKeyManagerKeyServersCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{uuid}/key-servers][%d] securityKeyManagerKeyServersCreateCreated %s", 201, payload)
}

func (o *SecurityKeyManagerKeyServersCreateCreated) GetPayload() *models.KeyServerResponse {
	return o.Payload
}

func (o *SecurityKeyManagerKeyServersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.KeyServerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeyManagerKeyServersCreateDefault creates a SecurityKeyManagerKeyServersCreateDefault with default headers values
func NewSecurityKeyManagerKeyServersCreateDefault(code int) *SecurityKeyManagerKeyServersCreateDefault {
	return &SecurityKeyManagerKeyServersCreateDefault{
		_statusCode: code,
	}
}

/*
	SecurityKeyManagerKeyServersCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536038 | A maximum of 4 active primary key servers are allowed. |
| 65536042 | Cannot add key server because it is already a secondary key server. |
| 65536600 | Cannot add a key server while a node is out quorum. |
| 65536821 | The certificate is not installed. |
| 65536824 | Multitenant key management is not supported in MetroCluster configurations. |
| 65536828 | External key management is not enabled for the SVM. |
| 65536834 | Failed to get existing key-server details for the SVM. |
| 65536852 | Failed to query supported KMIP protocol versions. |
| 65536870 | Key management servers are already configured. |
| 65536870 | The key management servers already exist. |
| 65536871 | Duplicate key management servers exist. |
| 65536921 | Unable to execute the command on the KMIP server. |
| 66060338 | Unable to establish secure connection to KMIP server due to incorrect server_ca certificates. |
| 66060339 | Unable to establish secure connection to KMIP server due to incorrect client certificates. |
| 66060340 | Unable to establish secure connection to KMIP server due to Cryptsoft error. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityKeyManagerKeyServersCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security key manager key servers create default response has a 2xx status code
func (o *SecurityKeyManagerKeyServersCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security key manager key servers create default response has a 3xx status code
func (o *SecurityKeyManagerKeyServersCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security key manager key servers create default response has a 4xx status code
func (o *SecurityKeyManagerKeyServersCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security key manager key servers create default response has a 5xx status code
func (o *SecurityKeyManagerKeyServersCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security key manager key servers create default response a status code equal to that given
func (o *SecurityKeyManagerKeyServersCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security key manager key servers create default response
func (o *SecurityKeyManagerKeyServersCreateDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeyManagerKeyServersCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{uuid}/key-servers][%d] security_key_manager_key_servers_create default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerKeyServersCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/key-managers/{uuid}/key-servers][%d] security_key_manager_key_servers_create default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerKeyServersCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeyManagerKeyServersCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}