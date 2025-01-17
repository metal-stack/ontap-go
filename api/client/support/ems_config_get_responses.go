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

// EmsConfigGetReader is a Reader for the EmsConfigGet structure.
type EmsConfigGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsConfigGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmsConfigGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsConfigGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsConfigGetOK creates a EmsConfigGetOK with default headers values
func NewEmsConfigGetOK() *EmsConfigGetOK {
	return &EmsConfigGetOK{}
}

/*
EmsConfigGetOK describes a response with status code 200, with default header values.

OK
*/
type EmsConfigGetOK struct {
	Payload *models.EmsConfig
}

// IsSuccess returns true when this ems config get o k response has a 2xx status code
func (o *EmsConfigGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ems config get o k response has a 3xx status code
func (o *EmsConfigGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ems config get o k response has a 4xx status code
func (o *EmsConfigGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ems config get o k response has a 5xx status code
func (o *EmsConfigGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ems config get o k response a status code equal to that given
func (o *EmsConfigGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ems config get o k response
func (o *EmsConfigGetOK) Code() int {
	return 200
}

func (o *EmsConfigGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems][%d] emsConfigGetOK %s", 200, payload)
}

func (o *EmsConfigGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems][%d] emsConfigGetOK %s", 200, payload)
}

func (o *EmsConfigGetOK) GetPayload() *models.EmsConfig {
	return o.Payload
}

func (o *EmsConfigGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmsConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmsConfigGetDefault creates a EmsConfigGetDefault with default headers values
func NewEmsConfigGetDefault(code int) *EmsConfigGetDefault {
	return &EmsConfigGetDefault{
		_statusCode: code,
	}
}

/*
EmsConfigGetDefault describes a response with status code -1, with default header values.

Error
*/
type EmsConfigGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ems config get default response has a 2xx status code
func (o *EmsConfigGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ems config get default response has a 3xx status code
func (o *EmsConfigGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ems config get default response has a 4xx status code
func (o *EmsConfigGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ems config get default response has a 5xx status code
func (o *EmsConfigGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ems config get default response a status code equal to that given
func (o *EmsConfigGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ems config get default response
func (o *EmsConfigGetDefault) Code() int {
	return o._statusCode
}

func (o *EmsConfigGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems][%d] ems_config_get default %s", o._statusCode, payload)
}

func (o *EmsConfigGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/ems][%d] ems_config_get default %s", o._statusCode, payload)
}

func (o *EmsConfigGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsConfigGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
