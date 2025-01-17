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

// SecurityKeyManagerDeleteReader is a Reader for the SecurityKeyManagerDelete structure.
type SecurityKeyManagerDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeyManagerDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeyManagerDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeyManagerDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeyManagerDeleteOK creates a SecurityKeyManagerDeleteOK with default headers values
func NewSecurityKeyManagerDeleteOK() *SecurityKeyManagerDeleteOK {
	return &SecurityKeyManagerDeleteOK{}
}

/*
SecurityKeyManagerDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeyManagerDeleteOK struct {
}

// IsSuccess returns true when this security key manager delete o k response has a 2xx status code
func (o *SecurityKeyManagerDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security key manager delete o k response has a 3xx status code
func (o *SecurityKeyManagerDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security key manager delete o k response has a 4xx status code
func (o *SecurityKeyManagerDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security key manager delete o k response has a 5xx status code
func (o *SecurityKeyManagerDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security key manager delete o k response a status code equal to that given
func (o *SecurityKeyManagerDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security key manager delete o k response
func (o *SecurityKeyManagerDeleteOK) Code() int {
	return 200
}

func (o *SecurityKeyManagerDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/key-managers/{uuid}][%d] securityKeyManagerDeleteOK", 200)
}

func (o *SecurityKeyManagerDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/key-managers/{uuid}][%d] securityKeyManagerDeleteOK", 200)
}

func (o *SecurityKeyManagerDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecurityKeyManagerDeleteDefault creates a SecurityKeyManagerDeleteDefault with default headers values
func NewSecurityKeyManagerDeleteDefault(code int) *SecurityKeyManagerDeleteDefault {
	return &SecurityKeyManagerDeleteDefault{
		_statusCode: code,
	}
}

/*
	SecurityKeyManagerDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536208 | Failed to delete the SVM Key ID. |
| 65536233 | Internal error. Deletion of km_wrapped_kdb key database has failed for the Onboard Key Manager. |
| 65536234 | Internal error. Deletion of cluster_kdb key database has failed for the Onboard Key Manager. |
| 65536239 | Encrypted volumes are found for the SVM. |
| 65536242 | One or more self-encrypting drives are assigned an authentication key. |
| 65536243 | Cannot determine authentication key presence on one or more self-encrypting drives. |
| 65536800 | Failed to lookup onboard keys. |
| 65536813 | Encrypted kernel core files found. |
| 65536817 | Failed to determine if key manager is safe to disable. |
| 65536827 | Failed to determine if the SVM has any encrypted volumes. |
| 65536828 | External key management is not enabled for the SVM. |
| 65536867 | Encrypted volumes are found for the SVM. |
| 196608301 | Failed to determine the type of encryption. |
| 196608305 | NAE aggregates are found in the cluster. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
  - name: KEYMANAGER_MESSAGE_ERR_KM_DISABLE_ENC_CORE_CHECK_TIMEOUT
    message: Failed to disable the key manager because of a timeout when checking for encrypted cores.
*/
type SecurityKeyManagerDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security key manager delete default response has a 2xx status code
func (o *SecurityKeyManagerDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security key manager delete default response has a 3xx status code
func (o *SecurityKeyManagerDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security key manager delete default response has a 4xx status code
func (o *SecurityKeyManagerDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security key manager delete default response has a 5xx status code
func (o *SecurityKeyManagerDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security key manager delete default response a status code equal to that given
func (o *SecurityKeyManagerDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security key manager delete default response
func (o *SecurityKeyManagerDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeyManagerDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/key-managers/{uuid}][%d] security_key_manager_delete default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/key-managers/{uuid}][%d] security_key_manager_delete default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeyManagerDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}