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

// MediatorCollectionGetReader is a Reader for the MediatorCollectionGet structure.
type MediatorCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MediatorCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMediatorCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMediatorCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMediatorCollectionGetOK creates a MediatorCollectionGetOK with default headers values
func NewMediatorCollectionGetOK() *MediatorCollectionGetOK {
	return &MediatorCollectionGetOK{}
}

/*
MediatorCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type MediatorCollectionGetOK struct {
	Payload *models.MediatorResponse
}

// IsSuccess returns true when this mediator collection get o k response has a 2xx status code
func (o *MediatorCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this mediator collection get o k response has a 3xx status code
func (o *MediatorCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this mediator collection get o k response has a 4xx status code
func (o *MediatorCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this mediator collection get o k response has a 5xx status code
func (o *MediatorCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this mediator collection get o k response a status code equal to that given
func (o *MediatorCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the mediator collection get o k response
func (o *MediatorCollectionGetOK) Code() int {
	return 200
}

func (o *MediatorCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/mediators][%d] mediatorCollectionGetOK %s", 200, payload)
}

func (o *MediatorCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/mediators][%d] mediatorCollectionGetOK %s", 200, payload)
}

func (o *MediatorCollectionGetOK) GetPayload() *models.MediatorResponse {
	return o.Payload
}

func (o *MediatorCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MediatorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMediatorCollectionGetDefault creates a MediatorCollectionGetDefault with default headers values
func NewMediatorCollectionGetDefault(code int) *MediatorCollectionGetDefault {
	return &MediatorCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	MediatorCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 2430739     | Unable to access Mediator. Reason: Invalid Mediator IP or Mediator does not exist.|
*/
type MediatorCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this mediator collection get default response has a 2xx status code
func (o *MediatorCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this mediator collection get default response has a 3xx status code
func (o *MediatorCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this mediator collection get default response has a 4xx status code
func (o *MediatorCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this mediator collection get default response has a 5xx status code
func (o *MediatorCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this mediator collection get default response a status code equal to that given
func (o *MediatorCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the mediator collection get default response
func (o *MediatorCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *MediatorCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/mediators][%d] mediator_collection_get default %s", o._statusCode, payload)
}

func (o *MediatorCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/mediators][%d] mediator_collection_get default %s", o._statusCode, payload)
}

func (o *MediatorCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MediatorCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
