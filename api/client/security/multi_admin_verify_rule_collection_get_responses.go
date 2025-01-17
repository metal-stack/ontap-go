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

// MultiAdminVerifyRuleCollectionGetReader is a Reader for the MultiAdminVerifyRuleCollectionGet structure.
type MultiAdminVerifyRuleCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MultiAdminVerifyRuleCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMultiAdminVerifyRuleCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMultiAdminVerifyRuleCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMultiAdminVerifyRuleCollectionGetOK creates a MultiAdminVerifyRuleCollectionGetOK with default headers values
func NewMultiAdminVerifyRuleCollectionGetOK() *MultiAdminVerifyRuleCollectionGetOK {
	return &MultiAdminVerifyRuleCollectionGetOK{}
}

/*
MultiAdminVerifyRuleCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type MultiAdminVerifyRuleCollectionGetOK struct {
	Payload *models.MultiAdminVerifyRuleResponse
}

// IsSuccess returns true when this multi admin verify rule collection get o k response has a 2xx status code
func (o *MultiAdminVerifyRuleCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this multi admin verify rule collection get o k response has a 3xx status code
func (o *MultiAdminVerifyRuleCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this multi admin verify rule collection get o k response has a 4xx status code
func (o *MultiAdminVerifyRuleCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this multi admin verify rule collection get o k response has a 5xx status code
func (o *MultiAdminVerifyRuleCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this multi admin verify rule collection get o k response a status code equal to that given
func (o *MultiAdminVerifyRuleCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the multi admin verify rule collection get o k response
func (o *MultiAdminVerifyRuleCollectionGetOK) Code() int {
	return 200
}

func (o *MultiAdminVerifyRuleCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/multi-admin-verify/rules][%d] multiAdminVerifyRuleCollectionGetOK %s", 200, payload)
}

func (o *MultiAdminVerifyRuleCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/multi-admin-verify/rules][%d] multiAdminVerifyRuleCollectionGetOK %s", 200, payload)
}

func (o *MultiAdminVerifyRuleCollectionGetOK) GetPayload() *models.MultiAdminVerifyRuleResponse {
	return o.Payload
}

func (o *MultiAdminVerifyRuleCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MultiAdminVerifyRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiAdminVerifyRuleCollectionGetDefault creates a MultiAdminVerifyRuleCollectionGetDefault with default headers values
func NewMultiAdminVerifyRuleCollectionGetDefault(code int) *MultiAdminVerifyRuleCollectionGetDefault {
	return &MultiAdminVerifyRuleCollectionGetDefault{
		_statusCode: code,
	}
}

/*
MultiAdminVerifyRuleCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type MultiAdminVerifyRuleCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this multi admin verify rule collection get default response has a 2xx status code
func (o *MultiAdminVerifyRuleCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this multi admin verify rule collection get default response has a 3xx status code
func (o *MultiAdminVerifyRuleCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this multi admin verify rule collection get default response has a 4xx status code
func (o *MultiAdminVerifyRuleCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this multi admin verify rule collection get default response has a 5xx status code
func (o *MultiAdminVerifyRuleCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this multi admin verify rule collection get default response a status code equal to that given
func (o *MultiAdminVerifyRuleCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the multi admin verify rule collection get default response
func (o *MultiAdminVerifyRuleCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *MultiAdminVerifyRuleCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/multi-admin-verify/rules][%d] multi_admin_verify_rule_collection_get default %s", o._statusCode, payload)
}

func (o *MultiAdminVerifyRuleCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/multi-admin-verify/rules][%d] multi_admin_verify_rule_collection_get default %s", o._statusCode, payload)
}

func (o *MultiAdminVerifyRuleCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MultiAdminVerifyRuleCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}