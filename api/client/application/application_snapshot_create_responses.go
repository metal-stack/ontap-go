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

// ApplicationSnapshotCreateReader is a Reader for the ApplicationSnapshotCreate structure.
type ApplicationSnapshotCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationSnapshotCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewApplicationSnapshotCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewApplicationSnapshotCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationSnapshotCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationSnapshotCreateCreated creates a ApplicationSnapshotCreateCreated with default headers values
func NewApplicationSnapshotCreateCreated() *ApplicationSnapshotCreateCreated {
	return &ApplicationSnapshotCreateCreated{}
}

/*
ApplicationSnapshotCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ApplicationSnapshotCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this application snapshot create created response has a 2xx status code
func (o *ApplicationSnapshotCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application snapshot create created response has a 3xx status code
func (o *ApplicationSnapshotCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application snapshot create created response has a 4xx status code
func (o *ApplicationSnapshotCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this application snapshot create created response has a 5xx status code
func (o *ApplicationSnapshotCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this application snapshot create created response a status code equal to that given
func (o *ApplicationSnapshotCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the application snapshot create created response
func (o *ApplicationSnapshotCreateCreated) Code() int {
	return 201
}

func (o *ApplicationSnapshotCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/snapshots][%d] applicationSnapshotCreateCreated %s", 201, payload)
}

func (o *ApplicationSnapshotCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/snapshots][%d] applicationSnapshotCreateCreated %s", 201, payload)
}

func (o *ApplicationSnapshotCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ApplicationSnapshotCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationSnapshotCreateAccepted creates a ApplicationSnapshotCreateAccepted with default headers values
func NewApplicationSnapshotCreateAccepted() *ApplicationSnapshotCreateAccepted {
	return &ApplicationSnapshotCreateAccepted{}
}

/*
ApplicationSnapshotCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ApplicationSnapshotCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this application snapshot create accepted response has a 2xx status code
func (o *ApplicationSnapshotCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application snapshot create accepted response has a 3xx status code
func (o *ApplicationSnapshotCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application snapshot create accepted response has a 4xx status code
func (o *ApplicationSnapshotCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this application snapshot create accepted response has a 5xx status code
func (o *ApplicationSnapshotCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this application snapshot create accepted response a status code equal to that given
func (o *ApplicationSnapshotCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the application snapshot create accepted response
func (o *ApplicationSnapshotCreateAccepted) Code() int {
	return 202
}

func (o *ApplicationSnapshotCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/snapshots][%d] applicationSnapshotCreateAccepted %s", 202, payload)
}

func (o *ApplicationSnapshotCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/snapshots][%d] applicationSnapshotCreateAccepted %s", 202, payload)
}

func (o *ApplicationSnapshotCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ApplicationSnapshotCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationSnapshotCreateDefault creates a ApplicationSnapshotCreateDefault with default headers values
func NewApplicationSnapshotCreateDefault(code int) *ApplicationSnapshotCreateDefault {
	return &ApplicationSnapshotCreateDefault{
		_statusCode: code,
	}
}

/*
ApplicationSnapshotCreateDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationSnapshotCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this application snapshot create default response has a 2xx status code
func (o *ApplicationSnapshotCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application snapshot create default response has a 3xx status code
func (o *ApplicationSnapshotCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application snapshot create default response has a 4xx status code
func (o *ApplicationSnapshotCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application snapshot create default response has a 5xx status code
func (o *ApplicationSnapshotCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application snapshot create default response a status code equal to that given
func (o *ApplicationSnapshotCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the application snapshot create default response
func (o *ApplicationSnapshotCreateDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationSnapshotCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/snapshots][%d] application_snapshot_create default %s", o._statusCode, payload)
}

func (o *ApplicationSnapshotCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/applications/{application.uuid}/snapshots][%d] application_snapshot_create default %s", o._statusCode, payload)
}

func (o *ApplicationSnapshotCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationSnapshotCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}