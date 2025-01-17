// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// CifsServiceDeleteReader is a Reader for the CifsServiceDelete structure.
type CifsServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCifsServiceDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsServiceDeleteOK creates a CifsServiceDeleteOK with default headers values
func NewCifsServiceDeleteOK() *CifsServiceDeleteOK {
	return &CifsServiceDeleteOK{}
}

/*
CifsServiceDeleteOK describes a response with status code 200, with default header values.

OK
*/
type CifsServiceDeleteOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cifs service delete o k response has a 2xx status code
func (o *CifsServiceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs service delete o k response has a 3xx status code
func (o *CifsServiceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs service delete o k response has a 4xx status code
func (o *CifsServiceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs service delete o k response has a 5xx status code
func (o *CifsServiceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs service delete o k response a status code equal to that given
func (o *CifsServiceDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs service delete o k response
func (o *CifsServiceDeleteOK) Code() int {
	return 200
}

func (o *CifsServiceDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/services/{svm.uuid}][%d] cifsServiceDeleteOK %s", 200, payload)
}

func (o *CifsServiceDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/services/{svm.uuid}][%d] cifsServiceDeleteOK %s", 200, payload)
}

func (o *CifsServiceDeleteOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *CifsServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsServiceDeleteAccepted creates a CifsServiceDeleteAccepted with default headers values
func NewCifsServiceDeleteAccepted() *CifsServiceDeleteAccepted {
	return &CifsServiceDeleteAccepted{}
}

/*
CifsServiceDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type CifsServiceDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cifs service delete accepted response has a 2xx status code
func (o *CifsServiceDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs service delete accepted response has a 3xx status code
func (o *CifsServiceDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs service delete accepted response has a 4xx status code
func (o *CifsServiceDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs service delete accepted response has a 5xx status code
func (o *CifsServiceDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs service delete accepted response a status code equal to that given
func (o *CifsServiceDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the cifs service delete accepted response
func (o *CifsServiceDeleteAccepted) Code() int {
	return 202
}

func (o *CifsServiceDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/services/{svm.uuid}][%d] cifsServiceDeleteAccepted %s", 202, payload)
}

func (o *CifsServiceDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/services/{svm.uuid}][%d] cifsServiceDeleteAccepted %s", 202, payload)
}

func (o *CifsServiceDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *CifsServiceDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsServiceDeleteDefault creates a CifsServiceDeleteDefault with default headers values
func NewCifsServiceDeleteDefault(code int) *CifsServiceDeleteDefault {
	return &CifsServiceDeleteDefault{
		_statusCode: code,
	}
}

/*
	CifsServiceDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655525     | In order to delete an Active Directory machine account for the CIFS server, you must supply the name and password of a Windows account with sufficient privileges.|
| 3735751    | Failed to authenticate and retrieve the access token from the Azure OAuth host. |
| 3735752    | Failed to extract the private key from the Azure Key Vault certificate. |
| 3735753    | Unsupported content_type in the Azure secrets response. |
| 3735754    | Failed to parse the JSON response from Azure Key Vault. |
| 3735755    | REST API call to Azure failed. |
| 3735756    | Invalid client certificate. |
| 3735757    | Failed to generate client assertion. |
| 3735762    | The provided Azure Key Vault configuration is incorrect. |
| 3735763    | The provided Azure Key Vault configuration is incomplete. |
| 3735764    | Request to Azure failed. Reason - Azure error code and Azure error message. |
| 655563     | NetBIOS name contains characters that are not allowed. |
| 655562     | NetBIOS name is longer than 15 characters. |
| 655923     | Retrieving credentials from AKV is not supported because the effective cluster version is not ONTAP 9.15.0 or later. |
*/
type CifsServiceDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs service delete default response has a 2xx status code
func (o *CifsServiceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs service delete default response has a 3xx status code
func (o *CifsServiceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs service delete default response has a 4xx status code
func (o *CifsServiceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs service delete default response has a 5xx status code
func (o *CifsServiceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs service delete default response a status code equal to that given
func (o *CifsServiceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs service delete default response
func (o *CifsServiceDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CifsServiceDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/services/{svm.uuid}][%d] cifs_service_delete default %s", o._statusCode, payload)
}

func (o *CifsServiceDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/services/{svm.uuid}][%d] cifs_service_delete default %s", o._statusCode, payload)
}

func (o *CifsServiceDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
