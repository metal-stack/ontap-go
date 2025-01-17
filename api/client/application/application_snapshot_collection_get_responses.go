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

// ApplicationSnapshotCollectionGetReader is a Reader for the ApplicationSnapshotCollectionGet structure.
type ApplicationSnapshotCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationSnapshotCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationSnapshotCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationSnapshotCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationSnapshotCollectionGetOK creates a ApplicationSnapshotCollectionGetOK with default headers values
func NewApplicationSnapshotCollectionGetOK() *ApplicationSnapshotCollectionGetOK {
	return &ApplicationSnapshotCollectionGetOK{}
}

/*
ApplicationSnapshotCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ApplicationSnapshotCollectionGetOK struct {
	Payload *models.ApplicationSnapshotResponse
}

// IsSuccess returns true when this application snapshot collection get o k response has a 2xx status code
func (o *ApplicationSnapshotCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application snapshot collection get o k response has a 3xx status code
func (o *ApplicationSnapshotCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application snapshot collection get o k response has a 4xx status code
func (o *ApplicationSnapshotCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this application snapshot collection get o k response has a 5xx status code
func (o *ApplicationSnapshotCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this application snapshot collection get o k response a status code equal to that given
func (o *ApplicationSnapshotCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the application snapshot collection get o k response
func (o *ApplicationSnapshotCollectionGetOK) Code() int {
	return 200
}

func (o *ApplicationSnapshotCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/snapshots][%d] applicationSnapshotCollectionGetOK %s", 200, payload)
}

func (o *ApplicationSnapshotCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/snapshots][%d] applicationSnapshotCollectionGetOK %s", 200, payload)
}

func (o *ApplicationSnapshotCollectionGetOK) GetPayload() *models.ApplicationSnapshotResponse {
	return o.Payload
}

func (o *ApplicationSnapshotCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationSnapshotCollectionGetDefault creates a ApplicationSnapshotCollectionGetDefault with default headers values
func NewApplicationSnapshotCollectionGetDefault(code int) *ApplicationSnapshotCollectionGetDefault {
	return &ApplicationSnapshotCollectionGetDefault{
		_statusCode: code,
	}
}

/*
ApplicationSnapshotCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationSnapshotCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this application snapshot collection get default response has a 2xx status code
func (o *ApplicationSnapshotCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application snapshot collection get default response has a 3xx status code
func (o *ApplicationSnapshotCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application snapshot collection get default response has a 4xx status code
func (o *ApplicationSnapshotCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application snapshot collection get default response has a 5xx status code
func (o *ApplicationSnapshotCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application snapshot collection get default response a status code equal to that given
func (o *ApplicationSnapshotCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the application snapshot collection get default response
func (o *ApplicationSnapshotCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationSnapshotCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/snapshots][%d] application_snapshot_collection_get default %s", o._statusCode, payload)
}

func (o *ApplicationSnapshotCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/snapshots][%d] application_snapshot_collection_get default %s", o._statusCode, payload)
}

func (o *ApplicationSnapshotCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationSnapshotCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}