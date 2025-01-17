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

// DuoCollectionGetReader is a Reader for the DuoCollectionGet structure.
type DuoCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DuoCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDuoCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDuoCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDuoCollectionGetOK creates a DuoCollectionGetOK with default headers values
func NewDuoCollectionGetOK() *DuoCollectionGetOK {
	return &DuoCollectionGetOK{}
}

/*
DuoCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type DuoCollectionGetOK struct {
	Payload *models.DuoResponse
}

// IsSuccess returns true when this duo collection get o k response has a 2xx status code
func (o *DuoCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this duo collection get o k response has a 3xx status code
func (o *DuoCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this duo collection get o k response has a 4xx status code
func (o *DuoCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this duo collection get o k response has a 5xx status code
func (o *DuoCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this duo collection get o k response a status code equal to that given
func (o *DuoCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the duo collection get o k response
func (o *DuoCollectionGetOK) Code() int {
	return 200
}

func (o *DuoCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/duo/profiles][%d] duoCollectionGetOK %s", 200, payload)
}

func (o *DuoCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/duo/profiles][%d] duoCollectionGetOK %s", 200, payload)
}

func (o *DuoCollectionGetOK) GetPayload() *models.DuoResponse {
	return o.Payload
}

func (o *DuoCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DuoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDuoCollectionGetDefault creates a DuoCollectionGetDefault with default headers values
func NewDuoCollectionGetDefault(code int) *DuoCollectionGetDefault {
	return &DuoCollectionGetDefault{
		_statusCode: code,
	}
}

/*
DuoCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type DuoCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this duo collection get default response has a 2xx status code
func (o *DuoCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this duo collection get default response has a 3xx status code
func (o *DuoCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this duo collection get default response has a 4xx status code
func (o *DuoCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this duo collection get default response has a 5xx status code
func (o *DuoCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this duo collection get default response a status code equal to that given
func (o *DuoCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the duo collection get default response
func (o *DuoCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *DuoCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/duo/profiles][%d] duo_collection_get default %s", o._statusCode, payload)
}

func (o *DuoCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/duo/profiles][%d] duo_collection_get default %s", o._statusCode, payload)
}

func (o *DuoCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DuoCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
