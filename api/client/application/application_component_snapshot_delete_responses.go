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

// ApplicationComponentSnapshotDeleteReader is a Reader for the ApplicationComponentSnapshotDelete structure.
type ApplicationComponentSnapshotDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationComponentSnapshotDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationComponentSnapshotDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewApplicationComponentSnapshotDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationComponentSnapshotDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationComponentSnapshotDeleteOK creates a ApplicationComponentSnapshotDeleteOK with default headers values
func NewApplicationComponentSnapshotDeleteOK() *ApplicationComponentSnapshotDeleteOK {
	return &ApplicationComponentSnapshotDeleteOK{}
}

/*
ApplicationComponentSnapshotDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ApplicationComponentSnapshotDeleteOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this application component snapshot delete o k response has a 2xx status code
func (o *ApplicationComponentSnapshotDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application component snapshot delete o k response has a 3xx status code
func (o *ApplicationComponentSnapshotDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application component snapshot delete o k response has a 4xx status code
func (o *ApplicationComponentSnapshotDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this application component snapshot delete o k response has a 5xx status code
func (o *ApplicationComponentSnapshotDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this application component snapshot delete o k response a status code equal to that given
func (o *ApplicationComponentSnapshotDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the application component snapshot delete o k response
func (o *ApplicationComponentSnapshotDeleteOK) Code() int {
	return 200
}

func (o *ApplicationComponentSnapshotDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid}][%d] applicationComponentSnapshotDeleteOK %s", 200, payload)
}

func (o *ApplicationComponentSnapshotDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid}][%d] applicationComponentSnapshotDeleteOK %s", 200, payload)
}

func (o *ApplicationComponentSnapshotDeleteOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ApplicationComponentSnapshotDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationComponentSnapshotDeleteAccepted creates a ApplicationComponentSnapshotDeleteAccepted with default headers values
func NewApplicationComponentSnapshotDeleteAccepted() *ApplicationComponentSnapshotDeleteAccepted {
	return &ApplicationComponentSnapshotDeleteAccepted{}
}

/*
ApplicationComponentSnapshotDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ApplicationComponentSnapshotDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this application component snapshot delete accepted response has a 2xx status code
func (o *ApplicationComponentSnapshotDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application component snapshot delete accepted response has a 3xx status code
func (o *ApplicationComponentSnapshotDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application component snapshot delete accepted response has a 4xx status code
func (o *ApplicationComponentSnapshotDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this application component snapshot delete accepted response has a 5xx status code
func (o *ApplicationComponentSnapshotDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this application component snapshot delete accepted response a status code equal to that given
func (o *ApplicationComponentSnapshotDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the application component snapshot delete accepted response
func (o *ApplicationComponentSnapshotDeleteAccepted) Code() int {
	return 202
}

func (o *ApplicationComponentSnapshotDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid}][%d] applicationComponentSnapshotDeleteAccepted %s", 202, payload)
}

func (o *ApplicationComponentSnapshotDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid}][%d] applicationComponentSnapshotDeleteAccepted %s", 202, payload)
}

func (o *ApplicationComponentSnapshotDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ApplicationComponentSnapshotDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationComponentSnapshotDeleteDefault creates a ApplicationComponentSnapshotDeleteDefault with default headers values
func NewApplicationComponentSnapshotDeleteDefault(code int) *ApplicationComponentSnapshotDeleteDefault {
	return &ApplicationComponentSnapshotDeleteDefault{
		_statusCode: code,
	}
}

/*
ApplicationComponentSnapshotDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationComponentSnapshotDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this application component snapshot delete default response has a 2xx status code
func (o *ApplicationComponentSnapshotDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application component snapshot delete default response has a 3xx status code
func (o *ApplicationComponentSnapshotDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application component snapshot delete default response has a 4xx status code
func (o *ApplicationComponentSnapshotDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application component snapshot delete default response has a 5xx status code
func (o *ApplicationComponentSnapshotDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application component snapshot delete default response a status code equal to that given
func (o *ApplicationComponentSnapshotDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the application component snapshot delete default response
func (o *ApplicationComponentSnapshotDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationComponentSnapshotDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid}][%d] application_component_snapshot_delete default %s", o._statusCode, payload)
}

func (o *ApplicationComponentSnapshotDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/applications/{application.uuid}/components/{component.uuid}/snapshots/{uuid}][%d] application_component_snapshot_delete default %s", o._statusCode, payload)
}

func (o *ApplicationComponentSnapshotDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationComponentSnapshotDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
