// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// ChassisCollectionGetReader is a Reader for the ChassisCollectionGet structure.
type ChassisCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChassisCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChassisCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChassisCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChassisCollectionGetOK creates a ChassisCollectionGetOK with default headers values
func NewChassisCollectionGetOK() *ChassisCollectionGetOK {
	return &ChassisCollectionGetOK{}
}

/*
ChassisCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ChassisCollectionGetOK struct {
	Payload *models.ChassisResponse
}

// IsSuccess returns true when this chassis collection get o k response has a 2xx status code
func (o *ChassisCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this chassis collection get o k response has a 3xx status code
func (o *ChassisCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chassis collection get o k response has a 4xx status code
func (o *ChassisCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this chassis collection get o k response has a 5xx status code
func (o *ChassisCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this chassis collection get o k response a status code equal to that given
func (o *ChassisCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the chassis collection get o k response
func (o *ChassisCollectionGetOK) Code() int {
	return 200
}

func (o *ChassisCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/chassis][%d] chassisCollectionGetOK %s", 200, payload)
}

func (o *ChassisCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/chassis][%d] chassisCollectionGetOK %s", 200, payload)
}

func (o *ChassisCollectionGetOK) GetPayload() *models.ChassisResponse {
	return o.Payload
}

func (o *ChassisCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChassisResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChassisCollectionGetDefault creates a ChassisCollectionGetDefault with default headers values
func NewChassisCollectionGetDefault(code int) *ChassisCollectionGetDefault {
	return &ChassisCollectionGetDefault{
		_statusCode: code,
	}
}

/*
ChassisCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ChassisCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this chassis collection get default response has a 2xx status code
func (o *ChassisCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this chassis collection get default response has a 3xx status code
func (o *ChassisCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this chassis collection get default response has a 4xx status code
func (o *ChassisCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this chassis collection get default response has a 5xx status code
func (o *ChassisCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this chassis collection get default response a status code equal to that given
func (o *ChassisCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the chassis collection get default response
func (o *ChassisCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ChassisCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/chassis][%d] chassis_collection_get default %s", o._statusCode, payload)
}

func (o *ChassisCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/chassis][%d] chassis_collection_get default %s", o._statusCode, payload)
}

func (o *ChassisCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ChassisCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
