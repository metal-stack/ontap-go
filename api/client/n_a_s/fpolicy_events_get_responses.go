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

// FpolicyEventsGetReader is a Reader for the FpolicyEventsGet structure.
type FpolicyEventsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyEventsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyEventsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyEventsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyEventsGetOK creates a FpolicyEventsGetOK with default headers values
func NewFpolicyEventsGetOK() *FpolicyEventsGetOK {
	return &FpolicyEventsGetOK{}
}

/*
FpolicyEventsGetOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyEventsGetOK struct {
	Payload *models.FpolicyEvent
}

// IsSuccess returns true when this fpolicy events get o k response has a 2xx status code
func (o *FpolicyEventsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy events get o k response has a 3xx status code
func (o *FpolicyEventsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy events get o k response has a 4xx status code
func (o *FpolicyEventsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy events get o k response has a 5xx status code
func (o *FpolicyEventsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy events get o k response a status code equal to that given
func (o *FpolicyEventsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fpolicy events get o k response
func (o *FpolicyEventsGetOK) Code() int {
	return 200
}

func (o *FpolicyEventsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/events/{name}][%d] fpolicyEventsGetOK %s", 200, payload)
}

func (o *FpolicyEventsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/events/{name}][%d] fpolicyEventsGetOK %s", 200, payload)
}

func (o *FpolicyEventsGetOK) GetPayload() *models.FpolicyEvent {
	return o.Payload
}

func (o *FpolicyEventsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FpolicyEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFpolicyEventsGetDefault creates a FpolicyEventsGetDefault with default headers values
func NewFpolicyEventsGetDefault(code int) *FpolicyEventsGetDefault {
	return &FpolicyEventsGetDefault{
		_statusCode: code,
	}
}

/*
FpolicyEventsGetDefault describes a response with status code -1, with default header values.

Error
*/
type FpolicyEventsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fpolicy events get default response has a 2xx status code
func (o *FpolicyEventsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy events get default response has a 3xx status code
func (o *FpolicyEventsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy events get default response has a 4xx status code
func (o *FpolicyEventsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy events get default response has a 5xx status code
func (o *FpolicyEventsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy events get default response a status code equal to that given
func (o *FpolicyEventsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fpolicy events get default response
func (o *FpolicyEventsGetDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyEventsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/events/{name}][%d] fpolicy_events_get default %s", o._statusCode, payload)
}

func (o *FpolicyEventsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/events/{name}][%d] fpolicy_events_get default %s", o._statusCode, payload)
}

func (o *FpolicyEventsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyEventsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
