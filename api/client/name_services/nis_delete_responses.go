// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NisDeleteReader is a Reader for the NisDelete structure.
type NisDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NisDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNisDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNisDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNisDeleteOK creates a NisDeleteOK with default headers values
func NewNisDeleteOK() *NisDeleteOK {
	return &NisDeleteOK{}
}

/*
NisDeleteOK describes a response with status code 200, with default header values.

OK
*/
type NisDeleteOK struct {
}

// IsSuccess returns true when this nis delete o k response has a 2xx status code
func (o *NisDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nis delete o k response has a 3xx status code
func (o *NisDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nis delete o k response has a 4xx status code
func (o *NisDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nis delete o k response has a 5xx status code
func (o *NisDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nis delete o k response a status code equal to that given
func (o *NisDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nis delete o k response
func (o *NisDeleteOK) Code() int {
	return 200
}

func (o *NisDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /name-services/nis/{svm.uuid}][%d] nisDeleteOK", 200)
}

func (o *NisDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /name-services/nis/{svm.uuid}][%d] nisDeleteOK", 200)
}

func (o *NisDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNisDeleteDefault creates a NisDeleteDefault with default headers values
func NewNisDeleteDefault(code int) *NisDeleteDefault {
	return &NisDeleteDefault{
		_statusCode: code,
	}
}

/*
NisDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type NisDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nis delete default response has a 2xx status code
func (o *NisDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nis delete default response has a 3xx status code
func (o *NisDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nis delete default response has a 4xx status code
func (o *NisDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nis delete default response has a 5xx status code
func (o *NisDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nis delete default response a status code equal to that given
func (o *NisDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nis delete default response
func (o *NisDeleteDefault) Code() int {
	return o._statusCode
}

func (o *NisDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /name-services/nis/{svm.uuid}][%d] nis_delete default %s", o._statusCode, payload)
}

func (o *NisDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /name-services/nis/{svm.uuid}][%d] nis_delete default %s", o._statusCode, payload)
}

func (o *NisDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NisDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}