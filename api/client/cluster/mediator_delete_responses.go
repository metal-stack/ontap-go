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

// MediatorDeleteReader is a Reader for the MediatorDelete structure.
type MediatorDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MediatorDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMediatorDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewMediatorDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMediatorDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMediatorDeleteOK creates a MediatorDeleteOK with default headers values
func NewMediatorDeleteOK() *MediatorDeleteOK {
	return &MediatorDeleteOK{}
}

/*
MediatorDeleteOK describes a response with status code 200, with default header values.

OK
*/
type MediatorDeleteOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this mediator delete o k response has a 2xx status code
func (o *MediatorDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this mediator delete o k response has a 3xx status code
func (o *MediatorDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this mediator delete o k response has a 4xx status code
func (o *MediatorDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this mediator delete o k response has a 5xx status code
func (o *MediatorDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this mediator delete o k response a status code equal to that given
func (o *MediatorDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the mediator delete o k response
func (o *MediatorDeleteOK) Code() int {
	return 200
}

func (o *MediatorDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/mediators/{uuid}][%d] mediatorDeleteOK %s", 200, payload)
}

func (o *MediatorDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/mediators/{uuid}][%d] mediatorDeleteOK %s", 200, payload)
}

func (o *MediatorDeleteOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *MediatorDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMediatorDeleteAccepted creates a MediatorDeleteAccepted with default headers values
func NewMediatorDeleteAccepted() *MediatorDeleteAccepted {
	return &MediatorDeleteAccepted{}
}

/*
MediatorDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type MediatorDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this mediator delete accepted response has a 2xx status code
func (o *MediatorDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this mediator delete accepted response has a 3xx status code
func (o *MediatorDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this mediator delete accepted response has a 4xx status code
func (o *MediatorDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this mediator delete accepted response has a 5xx status code
func (o *MediatorDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this mediator delete accepted response a status code equal to that given
func (o *MediatorDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the mediator delete accepted response
func (o *MediatorDeleteAccepted) Code() int {
	return 202
}

func (o *MediatorDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/mediators/{uuid}][%d] mediatorDeleteAccepted %s", 202, payload)
}

func (o *MediatorDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/mediators/{uuid}][%d] mediatorDeleteAccepted %s", 202, payload)
}

func (o *MediatorDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *MediatorDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMediatorDeleteDefault creates a MediatorDeleteDefault with default headers values
func NewMediatorDeleteDefault(code int) *MediatorDeleteDefault {
	return &MediatorDeleteDefault{
		_statusCode: code,
	}
}

/*
	MediatorDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 13369377    | Mediator field "mediator.id" does not exist.|
*/
type MediatorDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this mediator delete default response has a 2xx status code
func (o *MediatorDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this mediator delete default response has a 3xx status code
func (o *MediatorDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this mediator delete default response has a 4xx status code
func (o *MediatorDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this mediator delete default response has a 5xx status code
func (o *MediatorDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this mediator delete default response a status code equal to that given
func (o *MediatorDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the mediator delete default response
func (o *MediatorDeleteDefault) Code() int {
	return o._statusCode
}

func (o *MediatorDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/mediators/{uuid}][%d] mediator_delete default %s", o._statusCode, payload)
}

func (o *MediatorDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/mediators/{uuid}][%d] mediator_delete default %s", o._statusCode, payload)
}

func (o *MediatorDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MediatorDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
