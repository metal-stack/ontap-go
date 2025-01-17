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

// AutoUpdateStatusCollectionGetReader is a Reader for the AutoUpdateStatusCollectionGet structure.
type AutoUpdateStatusCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutoUpdateStatusCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutoUpdateStatusCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAutoUpdateStatusCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAutoUpdateStatusCollectionGetOK creates a AutoUpdateStatusCollectionGetOK with default headers values
func NewAutoUpdateStatusCollectionGetOK() *AutoUpdateStatusCollectionGetOK {
	return &AutoUpdateStatusCollectionGetOK{}
}

/*
AutoUpdateStatusCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type AutoUpdateStatusCollectionGetOK struct {
	Payload *models.AutoUpdateStatusResponse
}

// IsSuccess returns true when this auto update status collection get o k response has a 2xx status code
func (o *AutoUpdateStatusCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auto update status collection get o k response has a 3xx status code
func (o *AutoUpdateStatusCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auto update status collection get o k response has a 4xx status code
func (o *AutoUpdateStatusCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this auto update status collection get o k response has a 5xx status code
func (o *AutoUpdateStatusCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this auto update status collection get o k response a status code equal to that given
func (o *AutoUpdateStatusCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the auto update status collection get o k response
func (o *AutoUpdateStatusCollectionGetOK) Code() int {
	return 200
}

func (o *AutoUpdateStatusCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/auto-update/updates][%d] autoUpdateStatusCollectionGetOK %s", 200, payload)
}

func (o *AutoUpdateStatusCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/auto-update/updates][%d] autoUpdateStatusCollectionGetOK %s", 200, payload)
}

func (o *AutoUpdateStatusCollectionGetOK) GetPayload() *models.AutoUpdateStatusResponse {
	return o.Payload
}

func (o *AutoUpdateStatusCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AutoUpdateStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoUpdateStatusCollectionGetDefault creates a AutoUpdateStatusCollectionGetDefault with default headers values
func NewAutoUpdateStatusCollectionGetDefault(code int) *AutoUpdateStatusCollectionGetDefault {
	return &AutoUpdateStatusCollectionGetDefault{
		_statusCode: code,
	}
}

/*
AutoUpdateStatusCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type AutoUpdateStatusCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this auto update status collection get default response has a 2xx status code
func (o *AutoUpdateStatusCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this auto update status collection get default response has a 3xx status code
func (o *AutoUpdateStatusCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this auto update status collection get default response has a 4xx status code
func (o *AutoUpdateStatusCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this auto update status collection get default response has a 5xx status code
func (o *AutoUpdateStatusCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this auto update status collection get default response a status code equal to that given
func (o *AutoUpdateStatusCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the auto update status collection get default response
func (o *AutoUpdateStatusCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *AutoUpdateStatusCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/auto-update/updates][%d] auto_update_status_collection_get default %s", o._statusCode, payload)
}

func (o *AutoUpdateStatusCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/auto-update/updates][%d] auto_update_status_collection_get default %s", o._statusCode, payload)
}

func (o *AutoUpdateStatusCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutoUpdateStatusCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}