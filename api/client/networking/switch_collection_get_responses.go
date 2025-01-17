// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// SwitchCollectionGetReader is a Reader for the SwitchCollectionGet structure.
type SwitchCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SwitchCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSwitchCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSwitchCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSwitchCollectionGetOK creates a SwitchCollectionGetOK with default headers values
func NewSwitchCollectionGetOK() *SwitchCollectionGetOK {
	return &SwitchCollectionGetOK{}
}

/*
SwitchCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SwitchCollectionGetOK struct {
	Payload *models.SwitchResponse
}

// IsSuccess returns true when this switch collection get o k response has a 2xx status code
func (o *SwitchCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this switch collection get o k response has a 3xx status code
func (o *SwitchCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this switch collection get o k response has a 4xx status code
func (o *SwitchCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this switch collection get o k response has a 5xx status code
func (o *SwitchCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this switch collection get o k response a status code equal to that given
func (o *SwitchCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the switch collection get o k response
func (o *SwitchCollectionGetOK) Code() int {
	return 200
}

func (o *SwitchCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/switches][%d] switchCollectionGetOK %s", 200, payload)
}

func (o *SwitchCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/switches][%d] switchCollectionGetOK %s", 200, payload)
}

func (o *SwitchCollectionGetOK) GetPayload() *models.SwitchResponse {
	return o.Payload
}

func (o *SwitchCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SwitchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSwitchCollectionGetDefault creates a SwitchCollectionGetDefault with default headers values
func NewSwitchCollectionGetDefault(code int) *SwitchCollectionGetDefault {
	return &SwitchCollectionGetDefault{
		_statusCode: code,
	}
}

/*
SwitchCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SwitchCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this switch collection get default response has a 2xx status code
func (o *SwitchCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this switch collection get default response has a 3xx status code
func (o *SwitchCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this switch collection get default response has a 4xx status code
func (o *SwitchCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this switch collection get default response has a 5xx status code
func (o *SwitchCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this switch collection get default response a status code equal to that given
func (o *SwitchCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the switch collection get default response
func (o *SwitchCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SwitchCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/switches][%d] switch_collection_get default %s", o._statusCode, payload)
}

func (o *SwitchCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/switches][%d] switch_collection_get default %s", o._statusCode, payload)
}

func (o *SwitchCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SwitchCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
