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

// EmsApplicationLogsCreateReader is a Reader for the EmsApplicationLogsCreate structure.
type EmsApplicationLogsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsApplicationLogsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewEmsApplicationLogsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsApplicationLogsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsApplicationLogsCreateCreated creates a EmsApplicationLogsCreateCreated with default headers values
func NewEmsApplicationLogsCreateCreated() *EmsApplicationLogsCreateCreated {
	return &EmsApplicationLogsCreateCreated{}
}

/*
EmsApplicationLogsCreateCreated describes a response with status code 201, with default header values.

Created
*/
type EmsApplicationLogsCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this ems application logs create created response has a 2xx status code
func (o *EmsApplicationLogsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ems application logs create created response has a 3xx status code
func (o *EmsApplicationLogsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ems application logs create created response has a 4xx status code
func (o *EmsApplicationLogsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ems application logs create created response has a 5xx status code
func (o *EmsApplicationLogsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ems application logs create created response a status code equal to that given
func (o *EmsApplicationLogsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ems application logs create created response
func (o *EmsApplicationLogsCreateCreated) Code() int {
	return 201
}

func (o *EmsApplicationLogsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /support/ems/application-logs][%d] emsApplicationLogsCreateCreated", 201)
}

func (o *EmsApplicationLogsCreateCreated) String() string {
	return fmt.Sprintf("[POST /support/ems/application-logs][%d] emsApplicationLogsCreateCreated", 201)
}

func (o *EmsApplicationLogsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewEmsApplicationLogsCreateDefault creates a EmsApplicationLogsCreateDefault with default headers values
func NewEmsApplicationLogsCreateDefault(code int) *EmsApplicationLogsCreateDefault {
	return &EmsApplicationLogsCreateDefault{
		_statusCode: code,
	}
}

/*
	EmsApplicationLogsCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 983173 | Application log event generation failed. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type EmsApplicationLogsCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ems application logs create default response has a 2xx status code
func (o *EmsApplicationLogsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ems application logs create default response has a 3xx status code
func (o *EmsApplicationLogsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ems application logs create default response has a 4xx status code
func (o *EmsApplicationLogsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ems application logs create default response has a 5xx status code
func (o *EmsApplicationLogsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ems application logs create default response a status code equal to that given
func (o *EmsApplicationLogsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ems application logs create default response
func (o *EmsApplicationLogsCreateDefault) Code() int {
	return o._statusCode
}

func (o *EmsApplicationLogsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/application-logs][%d] ems_application_logs_create default %s", o._statusCode, payload)
}

func (o *EmsApplicationLogsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/application-logs][%d] ems_application_logs_create default %s", o._statusCode, payload)
}

func (o *EmsApplicationLogsCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsApplicationLogsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}