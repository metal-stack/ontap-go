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

// CifsOpenFileGetReader is a Reader for the CifsOpenFileGet structure.
type CifsOpenFileGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsOpenFileGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsOpenFileGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsOpenFileGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsOpenFileGetOK creates a CifsOpenFileGetOK with default headers values
func NewCifsOpenFileGetOK() *CifsOpenFileGetOK {
	return &CifsOpenFileGetOK{}
}

/*
CifsOpenFileGetOK describes a response with status code 200, with default header values.

OK
*/
type CifsOpenFileGetOK struct {
	Payload *models.CifsOpenFile
}

// IsSuccess returns true when this cifs open file get o k response has a 2xx status code
func (o *CifsOpenFileGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs open file get o k response has a 3xx status code
func (o *CifsOpenFileGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs open file get o k response has a 4xx status code
func (o *CifsOpenFileGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs open file get o k response has a 5xx status code
func (o *CifsOpenFileGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs open file get o k response a status code equal to that given
func (o *CifsOpenFileGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs open file get o k response
func (o *CifsOpenFileGetOK) Code() int {
	return 200
}

func (o *CifsOpenFileGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/session/files/{node.uuid}/{svm.uuid}/{identifier}/{connection.identifier}/{session.identifier}][%d] cifsOpenFileGetOK %s", 200, payload)
}

func (o *CifsOpenFileGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/session/files/{node.uuid}/{svm.uuid}/{identifier}/{connection.identifier}/{session.identifier}][%d] cifsOpenFileGetOK %s", 200, payload)
}

func (o *CifsOpenFileGetOK) GetPayload() *models.CifsOpenFile {
	return o.Payload
}

func (o *CifsOpenFileGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsOpenFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsOpenFileGetDefault creates a CifsOpenFileGetDefault with default headers values
func NewCifsOpenFileGetDefault(code int) *CifsOpenFileGetDefault {
	return &CifsOpenFileGetDefault{
		_statusCode: code,
	}
}

/*
CifsOpenFileGetDefault describes a response with status code -1, with default header values.

Error
*/
type CifsOpenFileGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs open file get default response has a 2xx status code
func (o *CifsOpenFileGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs open file get default response has a 3xx status code
func (o *CifsOpenFileGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs open file get default response has a 4xx status code
func (o *CifsOpenFileGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs open file get default response has a 5xx status code
func (o *CifsOpenFileGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs open file get default response a status code equal to that given
func (o *CifsOpenFileGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs open file get default response
func (o *CifsOpenFileGetDefault) Code() int {
	return o._statusCode
}

func (o *CifsOpenFileGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/session/files/{node.uuid}/{svm.uuid}/{identifier}/{connection.identifier}/{session.identifier}][%d] cifs_open_file_get default %s", o._statusCode, payload)
}

func (o *CifsOpenFileGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/session/files/{node.uuid}/{svm.uuid}/{identifier}/{connection.identifier}/{session.identifier}][%d] cifs_open_file_get default %s", o._statusCode, payload)
}

func (o *CifsOpenFileGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsOpenFileGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
