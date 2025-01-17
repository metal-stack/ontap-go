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

// VscanOnAccessGetReader is a Reader for the VscanOnAccessGet structure.
type VscanOnAccessGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanOnAccessGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanOnAccessGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanOnAccessGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanOnAccessGetOK creates a VscanOnAccessGetOK with default headers values
func NewVscanOnAccessGetOK() *VscanOnAccessGetOK {
	return &VscanOnAccessGetOK{}
}

/*
VscanOnAccessGetOK describes a response with status code 200, with default header values.

OK
*/
type VscanOnAccessGetOK struct {
	Payload *models.VscanOnAccess
}

// IsSuccess returns true when this vscan on access get o k response has a 2xx status code
func (o *VscanOnAccessGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan on access get o k response has a 3xx status code
func (o *VscanOnAccessGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan on access get o k response has a 4xx status code
func (o *VscanOnAccessGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan on access get o k response has a 5xx status code
func (o *VscanOnAccessGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan on access get o k response a status code equal to that given
func (o *VscanOnAccessGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vscan on access get o k response
func (o *VscanOnAccessGetOK) Code() int {
	return 200
}

func (o *VscanOnAccessGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/on-access-policies/{name}][%d] vscanOnAccessGetOK %s", 200, payload)
}

func (o *VscanOnAccessGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/on-access-policies/{name}][%d] vscanOnAccessGetOK %s", 200, payload)
}

func (o *VscanOnAccessGetOK) GetPayload() *models.VscanOnAccess {
	return o.Payload
}

func (o *VscanOnAccessGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VscanOnAccess)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVscanOnAccessGetDefault creates a VscanOnAccessGetDefault with default headers values
func NewVscanOnAccessGetDefault(code int) *VscanOnAccessGetDefault {
	return &VscanOnAccessGetDefault{
		_statusCode: code,
	}
}

/*
VscanOnAccessGetDefault describes a response with status code -1, with default header values.

Error
*/
type VscanOnAccessGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vscan on access get default response has a 2xx status code
func (o *VscanOnAccessGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan on access get default response has a 3xx status code
func (o *VscanOnAccessGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan on access get default response has a 4xx status code
func (o *VscanOnAccessGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan on access get default response has a 5xx status code
func (o *VscanOnAccessGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan on access get default response a status code equal to that given
func (o *VscanOnAccessGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vscan on access get default response
func (o *VscanOnAccessGetDefault) Code() int {
	return o._statusCode
}

func (o *VscanOnAccessGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/on-access-policies/{name}][%d] vscan_on_access_get default %s", o._statusCode, payload)
}

func (o *VscanOnAccessGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/{svm.uuid}/on-access-policies/{name}][%d] vscan_on_access_get default %s", o._statusCode, payload)
}

func (o *VscanOnAccessGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanOnAccessGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
