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

// ShadowcopySetGetReader is a Reader for the ShadowcopySetGet structure.
type ShadowcopySetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShadowcopySetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShadowcopySetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewShadowcopySetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShadowcopySetGetOK creates a ShadowcopySetGetOK with default headers values
func NewShadowcopySetGetOK() *ShadowcopySetGetOK {
	return &ShadowcopySetGetOK{}
}

/*
ShadowcopySetGetOK describes a response with status code 200, with default header values.

OK
*/
type ShadowcopySetGetOK struct {
	Payload *models.ShadowcopySet
}

// IsSuccess returns true when this shadowcopy set get o k response has a 2xx status code
func (o *ShadowcopySetGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this shadowcopy set get o k response has a 3xx status code
func (o *ShadowcopySetGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this shadowcopy set get o k response has a 4xx status code
func (o *ShadowcopySetGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this shadowcopy set get o k response has a 5xx status code
func (o *ShadowcopySetGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this shadowcopy set get o k response a status code equal to that given
func (o *ShadowcopySetGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the shadowcopy set get o k response
func (o *ShadowcopySetGetOK) Code() int {
	return 200
}

func (o *ShadowcopySetGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shadowcopy-sets/{uuid}][%d] shadowcopySetGetOK %s", 200, payload)
}

func (o *ShadowcopySetGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shadowcopy-sets/{uuid}][%d] shadowcopySetGetOK %s", 200, payload)
}

func (o *ShadowcopySetGetOK) GetPayload() *models.ShadowcopySet {
	return o.Payload
}

func (o *ShadowcopySetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShadowcopySet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShadowcopySetGetDefault creates a ShadowcopySetGetDefault with default headers values
func NewShadowcopySetGetDefault(code int) *ShadowcopySetGetDefault {
	return &ShadowcopySetGetDefault{
		_statusCode: code,
	}
}

/*
ShadowcopySetGetDefault describes a response with status code -1, with default header values.

Error
*/
type ShadowcopySetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this shadowcopy set get default response has a 2xx status code
func (o *ShadowcopySetGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this shadowcopy set get default response has a 3xx status code
func (o *ShadowcopySetGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this shadowcopy set get default response has a 4xx status code
func (o *ShadowcopySetGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this shadowcopy set get default response has a 5xx status code
func (o *ShadowcopySetGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this shadowcopy set get default response a status code equal to that given
func (o *ShadowcopySetGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the shadowcopy set get default response
func (o *ShadowcopySetGetDefault) Code() int {
	return o._statusCode
}

func (o *ShadowcopySetGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shadowcopy-sets/{uuid}][%d] shadowcopy_set_get default %s", o._statusCode, payload)
}

func (o *ShadowcopySetGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shadowcopy-sets/{uuid}][%d] shadowcopy_set_get default %s", o._statusCode, payload)
}

func (o *ShadowcopySetGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ShadowcopySetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
