// Code generated by go-swagger; DO NOT EDIT.

package application

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

// ConsistencyGroupGetReader is a Reader for the ConsistencyGroupGet structure.
type ConsistencyGroupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsistencyGroupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConsistencyGroupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConsistencyGroupGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsistencyGroupGetOK creates a ConsistencyGroupGetOK with default headers values
func NewConsistencyGroupGetOK() *ConsistencyGroupGetOK {
	return &ConsistencyGroupGetOK{}
}

/*
ConsistencyGroupGetOK describes a response with status code 200, with default header values.

OK
*/
type ConsistencyGroupGetOK struct {
	Payload *models.ConsistencyGroup
}

// IsSuccess returns true when this consistency group get o k response has a 2xx status code
func (o *ConsistencyGroupGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this consistency group get o k response has a 3xx status code
func (o *ConsistencyGroupGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this consistency group get o k response has a 4xx status code
func (o *ConsistencyGroupGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this consistency group get o k response has a 5xx status code
func (o *ConsistencyGroupGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this consistency group get o k response a status code equal to that given
func (o *ConsistencyGroupGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the consistency group get o k response
func (o *ConsistencyGroupGetOK) Code() int {
	return 200
}

func (o *ConsistencyGroupGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/consistency-groups/{uuid}][%d] consistencyGroupGetOK %s", 200, payload)
}

func (o *ConsistencyGroupGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/consistency-groups/{uuid}][%d] consistencyGroupGetOK %s", 200, payload)
}

func (o *ConsistencyGroupGetOK) GetPayload() *models.ConsistencyGroup {
	return o.Payload
}

func (o *ConsistencyGroupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsistencyGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsistencyGroupGetDefault creates a ConsistencyGroupGetDefault with default headers values
func NewConsistencyGroupGetDefault(code int) *ConsistencyGroupGetDefault {
	return &ConsistencyGroupGetDefault{
		_statusCode: code,
	}
}

/*
	ConsistencyGroupGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 53411842 | Consistency group does not exist. |
| 53411843 | A consistency group with specified UUID was not found. |
| 53411844 | Specified consistency group was not found in the specified SVM. |
| 53411845 | The specified UUID and name refer to different consistency groups. |
| 53411846 | Either name or UUID must be provided. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ConsistencyGroupGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this consistency group get default response has a 2xx status code
func (o *ConsistencyGroupGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this consistency group get default response has a 3xx status code
func (o *ConsistencyGroupGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this consistency group get default response has a 4xx status code
func (o *ConsistencyGroupGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this consistency group get default response has a 5xx status code
func (o *ConsistencyGroupGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this consistency group get default response a status code equal to that given
func (o *ConsistencyGroupGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the consistency group get default response
func (o *ConsistencyGroupGetDefault) Code() int {
	return o._statusCode
}

func (o *ConsistencyGroupGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/consistency-groups/{uuid}][%d] consistency_group_get default %s", o._statusCode, payload)
}

func (o *ConsistencyGroupGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/consistency-groups/{uuid}][%d] consistency_group_get default %s", o._statusCode, payload)
}

func (o *ConsistencyGroupGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConsistencyGroupGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}