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

// SecurityKeystoreModifyReader is a Reader for the SecurityKeystoreModify structure.
type SecurityKeystoreModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeystoreModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeystoreModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSecurityKeystoreModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeystoreModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeystoreModifyOK creates a SecurityKeystoreModifyOK with default headers values
func NewSecurityKeystoreModifyOK() *SecurityKeystoreModifyOK {
	return &SecurityKeystoreModifyOK{}
}

/*
SecurityKeystoreModifyOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeystoreModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this security keystore modify o k response has a 2xx status code
func (o *SecurityKeystoreModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security keystore modify o k response has a 3xx status code
func (o *SecurityKeystoreModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security keystore modify o k response has a 4xx status code
func (o *SecurityKeystoreModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security keystore modify o k response has a 5xx status code
func (o *SecurityKeystoreModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security keystore modify o k response a status code equal to that given
func (o *SecurityKeystoreModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security keystore modify o k response
func (o *SecurityKeystoreModifyOK) Code() int {
	return 200
}

func (o *SecurityKeystoreModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-stores/{uuid}][%d] securityKeystoreModifyOK %s", 200, payload)
}

func (o *SecurityKeystoreModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-stores/{uuid}][%d] securityKeystoreModifyOK %s", 200, payload)
}

func (o *SecurityKeystoreModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SecurityKeystoreModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeystoreModifyAccepted creates a SecurityKeystoreModifyAccepted with default headers values
func NewSecurityKeystoreModifyAccepted() *SecurityKeystoreModifyAccepted {
	return &SecurityKeystoreModifyAccepted{}
}

/*
SecurityKeystoreModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SecurityKeystoreModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this security keystore modify accepted response has a 2xx status code
func (o *SecurityKeystoreModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security keystore modify accepted response has a 3xx status code
func (o *SecurityKeystoreModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security keystore modify accepted response has a 4xx status code
func (o *SecurityKeystoreModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this security keystore modify accepted response has a 5xx status code
func (o *SecurityKeystoreModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this security keystore modify accepted response a status code equal to that given
func (o *SecurityKeystoreModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the security keystore modify accepted response
func (o *SecurityKeystoreModifyAccepted) Code() int {
	return 202
}

func (o *SecurityKeystoreModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-stores/{uuid}][%d] securityKeystoreModifyAccepted %s", 202, payload)
}

func (o *SecurityKeystoreModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-stores/{uuid}][%d] securityKeystoreModifyAccepted %s", 202, payload)
}

func (o *SecurityKeystoreModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SecurityKeystoreModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeystoreModifyDefault creates a SecurityKeystoreModifyDefault with default headers values
func NewSecurityKeystoreModifyDefault(code int) *SecurityKeystoreModifyDefault {
	return &SecurityKeystoreModifyDefault{
		_statusCode: code,
	}
}

/*
	SecurityKeystoreModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262155 | This operation requires an effective cluster version of 9.14.0 or later. |
| 65537605 | Failed to establish connectivity with the cloud key management service. |
| 65538908 | The specified keystore configuration UUID either does not exist or corresponds to a keystore configuration that is not supported by this operation. |
| 65538909 | A value for enabled is required. |
| 65538910 | Disabling an enabled configuration through this method is currently not supported. |
| 65539206 | The SVM associated with the supplied keystore UUID already has a keystore configuration enabled. This command does not support migrating from configurations of that keystore type". |
| 65539212 | Cannot switch the enabled keystore configuration when it is not in the 'active' or 'switching' state. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityKeystoreModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security keystore modify default response has a 2xx status code
func (o *SecurityKeystoreModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security keystore modify default response has a 3xx status code
func (o *SecurityKeystoreModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security keystore modify default response has a 4xx status code
func (o *SecurityKeystoreModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security keystore modify default response has a 5xx status code
func (o *SecurityKeystoreModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security keystore modify default response a status code equal to that given
func (o *SecurityKeystoreModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security keystore modify default response
func (o *SecurityKeystoreModifyDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeystoreModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-stores/{uuid}][%d] security_keystore_modify default %s", o._statusCode, payload)
}

func (o *SecurityKeystoreModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-stores/{uuid}][%d] security_keystore_modify default %s", o._statusCode, payload)
}

func (o *SecurityKeystoreModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeystoreModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
