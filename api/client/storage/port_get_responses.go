// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// PortGetReader is a Reader for the PortGet structure.
type PortGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PortGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPortGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPortGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPortGetOK creates a PortGetOK with default headers values
func NewPortGetOK() *PortGetOK {
	return &PortGetOK{}
}

/*
PortGetOK describes a response with status code 200, with default header values.

OK
*/
type PortGetOK struct {
	Payload *models.StoragePort
}

// IsSuccess returns true when this port get o k response has a 2xx status code
func (o *PortGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this port get o k response has a 3xx status code
func (o *PortGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this port get o k response has a 4xx status code
func (o *PortGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this port get o k response has a 5xx status code
func (o *PortGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this port get o k response a status code equal to that given
func (o *PortGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the port get o k response
func (o *PortGetOK) Code() int {
	return 200
}

func (o *PortGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/ports/{node.uuid}/{name}][%d] portGetOK %s", 200, payload)
}

func (o *PortGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/ports/{node.uuid}/{name}][%d] portGetOK %s", 200, payload)
}

func (o *PortGetOK) GetPayload() *models.StoragePort {
	return o.Payload
}

func (o *PortGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPortGetDefault creates a PortGetDefault with default headers values
func NewPortGetDefault(code int) *PortGetDefault {
	return &PortGetDefault{
		_statusCode: code,
	}
}

/*
PortGetDefault describes a response with status code -1, with default header values.

Error
*/
type PortGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this port get default response has a 2xx status code
func (o *PortGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this port get default response has a 3xx status code
func (o *PortGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this port get default response has a 4xx status code
func (o *PortGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this port get default response has a 5xx status code
func (o *PortGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this port get default response a status code equal to that given
func (o *PortGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the port get default response
func (o *PortGetDefault) Code() int {
	return o._statusCode
}

func (o *PortGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/ports/{node.uuid}/{name}][%d] port_get default %s", o._statusCode, payload)
}

func (o *PortGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/ports/{node.uuid}/{name}][%d] port_get default %s", o._statusCode, payload)
}

func (o *PortGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PortGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
