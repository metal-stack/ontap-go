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

// CreateCertificateSigningRequestReader is a Reader for the CreateCertificateSigningRequest structure.
type CreateCertificateSigningRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCertificateSigningRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCertificateSigningRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateCertificateSigningRequestDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCertificateSigningRequestOK creates a CreateCertificateSigningRequestOK with default headers values
func NewCreateCertificateSigningRequestOK() *CreateCertificateSigningRequestOK {
	return &CreateCertificateSigningRequestOK{}
}

/*
CreateCertificateSigningRequestOK describes a response with status code 200, with default header values.

OK
*/
type CreateCertificateSigningRequestOK struct {
	Payload *models.CertificateSigningRequest
}

// IsSuccess returns true when this create certificate signing request o k response has a 2xx status code
func (o *CreateCertificateSigningRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create certificate signing request o k response has a 3xx status code
func (o *CreateCertificateSigningRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create certificate signing request o k response has a 4xx status code
func (o *CreateCertificateSigningRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create certificate signing request o k response has a 5xx status code
func (o *CreateCertificateSigningRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create certificate signing request o k response a status code equal to that given
func (o *CreateCertificateSigningRequestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create certificate signing request o k response
func (o *CreateCertificateSigningRequestOK) Code() int {
	return 200
}

func (o *CreateCertificateSigningRequestOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/certificate-signing-request][%d] createCertificateSigningRequestOK %s", 200, payload)
}

func (o *CreateCertificateSigningRequestOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/certificate-signing-request][%d] createCertificateSigningRequestOK %s", 200, payload)
}

func (o *CreateCertificateSigningRequestOK) GetPayload() *models.CertificateSigningRequest {
	return o.Payload
}

func (o *CreateCertificateSigningRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateSigningRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCertificateSigningRequestDefault creates a CreateCertificateSigningRequestDefault with default headers values
func NewCreateCertificateSigningRequestDefault(code int) *CreateCertificateSigningRequestDefault {
	return &CreateCertificateSigningRequestDefault{
		_statusCode: code,
	}
}

/*
	CreateCertificateSigningRequestDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 3735554 | Certificate signing request failed. |
| 3735664 | Key size is not supported in FIPS mode. |
| 3735665 | Hash function is not supported in FIPS mode. |
| 3735700 | Key size is not supported. |
| 3735713 | Security strength bits length is not supported. |
| 3735714 | Security strength bits length is not supported in FIPS mode. |
| 3735715 | Certificate creation requires a common name or SAN extensions. |
| 3735741 | Key size is not applicable with the EC encryption algorithm. |
| 3735750 | Cannot supply "critical" as the only value in a list of extension values. |
| 52560173 | Hash function is not supported for digital signatures. |
| 52560423 | Failed to read the relative distinguished names. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type CreateCertificateSigningRequestDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create certificate signing request default response has a 2xx status code
func (o *CreateCertificateSigningRequestDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create certificate signing request default response has a 3xx status code
func (o *CreateCertificateSigningRequestDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create certificate signing request default response has a 4xx status code
func (o *CreateCertificateSigningRequestDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create certificate signing request default response has a 5xx status code
func (o *CreateCertificateSigningRequestDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create certificate signing request default response a status code equal to that given
func (o *CreateCertificateSigningRequestDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create certificate signing request default response
func (o *CreateCertificateSigningRequestDefault) Code() int {
	return o._statusCode
}

func (o *CreateCertificateSigningRequestDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/certificate-signing-request][%d] create_certificate_signing_request default %s", o._statusCode, payload)
}

func (o *CreateCertificateSigningRequestDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/certificate-signing-request][%d] create_certificate_signing_request default %s", o._statusCode, payload)
}

func (o *CreateCertificateSigningRequestDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateCertificateSigningRequestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}