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

// LocalCifsUserGetReader is a Reader for the LocalCifsUserGet structure.
type LocalCifsUserGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalCifsUserGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalCifsUserGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalCifsUserGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalCifsUserGetOK creates a LocalCifsUserGetOK with default headers values
func NewLocalCifsUserGetOK() *LocalCifsUserGetOK {
	return &LocalCifsUserGetOK{}
}

/*
LocalCifsUserGetOK describes a response with status code 200, with default header values.

OK
*/
type LocalCifsUserGetOK struct {
	Payload *models.LocalCifsUser
}

// IsSuccess returns true when this local cifs user get o k response has a 2xx status code
func (o *LocalCifsUserGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local cifs user get o k response has a 3xx status code
func (o *LocalCifsUserGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local cifs user get o k response has a 4xx status code
func (o *LocalCifsUserGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this local cifs user get o k response has a 5xx status code
func (o *LocalCifsUserGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this local cifs user get o k response a status code equal to that given
func (o *LocalCifsUserGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the local cifs user get o k response
func (o *LocalCifsUserGetOK) Code() int {
	return 200
}

func (o *LocalCifsUserGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/local-users/{svm.uuid}/{sid}][%d] localCifsUserGetOK %s", 200, payload)
}

func (o *LocalCifsUserGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/local-users/{svm.uuid}/{sid}][%d] localCifsUserGetOK %s", 200, payload)
}

func (o *LocalCifsUserGetOK) GetPayload() *models.LocalCifsUser {
	return o.Payload
}

func (o *LocalCifsUserGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocalCifsUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocalCifsUserGetDefault creates a LocalCifsUserGetDefault with default headers values
func NewLocalCifsUserGetDefault(code int) *LocalCifsUserGetDefault {
	return &LocalCifsUserGetDefault{
		_statusCode: code,
	}
}

/*
LocalCifsUserGetDefault describes a response with status code -1, with default header values.

Error
*/
type LocalCifsUserGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this local cifs user get default response has a 2xx status code
func (o *LocalCifsUserGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this local cifs user get default response has a 3xx status code
func (o *LocalCifsUserGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this local cifs user get default response has a 4xx status code
func (o *LocalCifsUserGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this local cifs user get default response has a 5xx status code
func (o *LocalCifsUserGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this local cifs user get default response a status code equal to that given
func (o *LocalCifsUserGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the local cifs user get default response
func (o *LocalCifsUserGetDefault) Code() int {
	return o._statusCode
}

func (o *LocalCifsUserGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/local-users/{svm.uuid}/{sid}][%d] local_cifs_user_get default %s", o._statusCode, payload)
}

func (o *LocalCifsUserGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/local-users/{svm.uuid}/{sid}][%d] local_cifs_user_get default %s", o._statusCode, payload)
}

func (o *LocalCifsUserGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalCifsUserGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
