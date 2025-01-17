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

// JobGetReader is a Reader for the JobGet structure.
type JobGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewJobGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewJobGetOK creates a JobGetOK with default headers values
func NewJobGetOK() *JobGetOK {
	return &JobGetOK{}
}

/*
JobGetOK describes a response with status code 200, with default header values.

OK
*/
type JobGetOK struct {
	Payload *models.Job
}

// IsSuccess returns true when this job get o k response has a 2xx status code
func (o *JobGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this job get o k response has a 3xx status code
func (o *JobGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this job get o k response has a 4xx status code
func (o *JobGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this job get o k response has a 5xx status code
func (o *JobGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this job get o k response a status code equal to that given
func (o *JobGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the job get o k response
func (o *JobGetOK) Code() int {
	return 200
}

func (o *JobGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/jobs/{uuid}][%d] jobGetOK %s", 200, payload)
}

func (o *JobGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/jobs/{uuid}][%d] jobGetOK %s", 200, payload)
}

func (o *JobGetOK) GetPayload() *models.Job {
	return o.Payload
}

func (o *JobGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJobGetDefault creates a JobGetDefault with default headers values
func NewJobGetDefault(code int) *JobGetDefault {
	return &JobGetDefault{
		_statusCode: code,
	}
}

/*
JobGetDefault describes a response with status code -1, with default header values.

Error
*/
type JobGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this job get default response has a 2xx status code
func (o *JobGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this job get default response has a 3xx status code
func (o *JobGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this job get default response has a 4xx status code
func (o *JobGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this job get default response has a 5xx status code
func (o *JobGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this job get default response a status code equal to that given
func (o *JobGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the job get default response
func (o *JobGetDefault) Code() int {
	return o._statusCode
}

func (o *JobGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/jobs/{uuid}][%d] job_get default %s", o._statusCode, payload)
}

func (o *JobGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/jobs/{uuid}][%d] job_get default %s", o._statusCode, payload)
}

func (o *JobGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *JobGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}