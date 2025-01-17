// Code generated by go-swagger; DO NOT EDIT.

package support

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

// EmsMessageCollectionGetReader is a Reader for the EmsMessageCollectionGet structure.
type EmsMessageCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsMessageCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmsMessageCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsMessageCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsMessageCollectionGetOK creates a EmsMessageCollectionGetOK with default headers values
func NewEmsMessageCollectionGetOK() *EmsMessageCollectionGetOK {
	return &EmsMessageCollectionGetOK{}
}

/*
EmsMessageCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type EmsMessageCollectionGetOK struct {
	Payload *models.EmsMessageResponse
}

// IsSuccess returns true when this ems message collection get o k response has a 2xx status code
func (o *EmsMessageCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ems message collection get o k response has a 3xx status code
func (o *EmsMessageCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ems message collection get o k response has a 4xx status code
func (o *EmsMessageCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ems message collection get o k response has a 5xx status code
func (o *EmsMessageCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ems message collection get o k response a status code equal to that given
func (o *EmsMessageCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ems message collection get o k response
func (o *EmsMessageCollectionGetOK) Code() int {
	return 200
}

func (o *EmsMessageCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/messages][%d] emsMessageCollectionGetOK %s", 200, payload)
}

func (o *EmsMessageCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/messages][%d] emsMessageCollectionGetOK %s", 200, payload)
}

func (o *EmsMessageCollectionGetOK) GetPayload() *models.EmsMessageResponse {
	return o.Payload
}

func (o *EmsMessageCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmsMessageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmsMessageCollectionGetDefault creates a EmsMessageCollectionGetDefault with default headers values
func NewEmsMessageCollectionGetDefault(code int) *EmsMessageCollectionGetDefault {
	return &EmsMessageCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	EmsMessageCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 983093 | The provided filter is unknown. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type EmsMessageCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ems message collection get default response has a 2xx status code
func (o *EmsMessageCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ems message collection get default response has a 3xx status code
func (o *EmsMessageCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ems message collection get default response has a 4xx status code
func (o *EmsMessageCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ems message collection get default response has a 5xx status code
func (o *EmsMessageCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ems message collection get default response a status code equal to that given
func (o *EmsMessageCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ems message collection get default response
func (o *EmsMessageCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *EmsMessageCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/messages][%d] ems_message_collection_get default %s", o._statusCode, payload)
}

func (o *EmsMessageCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems/messages][%d] ems_message_collection_get default %s", o._statusCode, payload)
}

func (o *EmsMessageCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsMessageCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
