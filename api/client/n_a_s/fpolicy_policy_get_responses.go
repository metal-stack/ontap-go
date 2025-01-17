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

// FpolicyPolicyGetReader is a Reader for the FpolicyPolicyGet structure.
type FpolicyPolicyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyPolicyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyPolicyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyPolicyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyPolicyGetOK creates a FpolicyPolicyGetOK with default headers values
func NewFpolicyPolicyGetOK() *FpolicyPolicyGetOK {
	return &FpolicyPolicyGetOK{}
}

/*
FpolicyPolicyGetOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyPolicyGetOK struct {
	Payload *models.FpolicyPolicy
}

// IsSuccess returns true when this fpolicy policy get o k response has a 2xx status code
func (o *FpolicyPolicyGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy policy get o k response has a 3xx status code
func (o *FpolicyPolicyGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy policy get o k response has a 4xx status code
func (o *FpolicyPolicyGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy policy get o k response has a 5xx status code
func (o *FpolicyPolicyGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy policy get o k response a status code equal to that given
func (o *FpolicyPolicyGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fpolicy policy get o k response
func (o *FpolicyPolicyGetOK) Code() int {
	return 200
}

func (o *FpolicyPolicyGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/policies/{name}][%d] fpolicyPolicyGetOK %s", 200, payload)
}

func (o *FpolicyPolicyGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/policies/{name}][%d] fpolicyPolicyGetOK %s", 200, payload)
}

func (o *FpolicyPolicyGetOK) GetPayload() *models.FpolicyPolicy {
	return o.Payload
}

func (o *FpolicyPolicyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FpolicyPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFpolicyPolicyGetDefault creates a FpolicyPolicyGetDefault with default headers values
func NewFpolicyPolicyGetDefault(code int) *FpolicyPolicyGetDefault {
	return &FpolicyPolicyGetDefault{
		_statusCode: code,
	}
}

/*
FpolicyPolicyGetDefault describes a response with status code -1, with default header values.

Error
*/
type FpolicyPolicyGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fpolicy policy get default response has a 2xx status code
func (o *FpolicyPolicyGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy policy get default response has a 3xx status code
func (o *FpolicyPolicyGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy policy get default response has a 4xx status code
func (o *FpolicyPolicyGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy policy get default response has a 5xx status code
func (o *FpolicyPolicyGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy policy get default response a status code equal to that given
func (o *FpolicyPolicyGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fpolicy policy get default response
func (o *FpolicyPolicyGetDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyPolicyGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/policies/{name}][%d] fpolicy_policy_get default %s", o._statusCode, payload)
}

func (o *FpolicyPolicyGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/policies/{name}][%d] fpolicy_policy_get default %s", o._statusCode, payload)
}

func (o *FpolicyPolicyGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyPolicyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
