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

// MultiAdminVerifyRequestGetReader is a Reader for the MultiAdminVerifyRequestGet structure.
type MultiAdminVerifyRequestGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MultiAdminVerifyRequestGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMultiAdminVerifyRequestGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMultiAdminVerifyRequestGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMultiAdminVerifyRequestGetOK creates a MultiAdminVerifyRequestGetOK with default headers values
func NewMultiAdminVerifyRequestGetOK() *MultiAdminVerifyRequestGetOK {
	return &MultiAdminVerifyRequestGetOK{}
}

/*
MultiAdminVerifyRequestGetOK describes a response with status code 200, with default header values.

OK
*/
type MultiAdminVerifyRequestGetOK struct {
	Payload *models.MultiAdminVerifyRequest
}

// IsSuccess returns true when this multi admin verify request get o k response has a 2xx status code
func (o *MultiAdminVerifyRequestGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this multi admin verify request get o k response has a 3xx status code
func (o *MultiAdminVerifyRequestGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this multi admin verify request get o k response has a 4xx status code
func (o *MultiAdminVerifyRequestGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this multi admin verify request get o k response has a 5xx status code
func (o *MultiAdminVerifyRequestGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this multi admin verify request get o k response a status code equal to that given
func (o *MultiAdminVerifyRequestGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the multi admin verify request get o k response
func (o *MultiAdminVerifyRequestGetOK) Code() int {
	return 200
}

func (o *MultiAdminVerifyRequestGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/multi-admin-verify/requests/{index}][%d] multiAdminVerifyRequestGetOK %s", 200, payload)
}

func (o *MultiAdminVerifyRequestGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/multi-admin-verify/requests/{index}][%d] multiAdminVerifyRequestGetOK %s", 200, payload)
}

func (o *MultiAdminVerifyRequestGetOK) GetPayload() *models.MultiAdminVerifyRequest {
	return o.Payload
}

func (o *MultiAdminVerifyRequestGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MultiAdminVerifyRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiAdminVerifyRequestGetDefault creates a MultiAdminVerifyRequestGetDefault with default headers values
func NewMultiAdminVerifyRequestGetDefault(code int) *MultiAdminVerifyRequestGetDefault {
	return &MultiAdminVerifyRequestGetDefault{
		_statusCode: code,
	}
}

/*
MultiAdminVerifyRequestGetDefault describes a response with status code -1, with default header values.

Error
*/
type MultiAdminVerifyRequestGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this multi admin verify request get default response has a 2xx status code
func (o *MultiAdminVerifyRequestGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this multi admin verify request get default response has a 3xx status code
func (o *MultiAdminVerifyRequestGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this multi admin verify request get default response has a 4xx status code
func (o *MultiAdminVerifyRequestGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this multi admin verify request get default response has a 5xx status code
func (o *MultiAdminVerifyRequestGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this multi admin verify request get default response a status code equal to that given
func (o *MultiAdminVerifyRequestGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the multi admin verify request get default response
func (o *MultiAdminVerifyRequestGetDefault) Code() int {
	return o._statusCode
}

func (o *MultiAdminVerifyRequestGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/multi-admin-verify/requests/{index}][%d] multi_admin_verify_request_get default %s", o._statusCode, payload)
}

func (o *MultiAdminVerifyRequestGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/multi-admin-verify/requests/{index}][%d] multi_admin_verify_request_get default %s", o._statusCode, payload)
}

func (o *MultiAdminVerifyRequestGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MultiAdminVerifyRequestGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
