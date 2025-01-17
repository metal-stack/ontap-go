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

// VscanServerStatusGetReader is a Reader for the VscanServerStatusGet structure.
type VscanServerStatusGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanServerStatusGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanServerStatusGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanServerStatusGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanServerStatusGetOK creates a VscanServerStatusGetOK with default headers values
func NewVscanServerStatusGetOK() *VscanServerStatusGetOK {
	return &VscanServerStatusGetOK{}
}

/*
VscanServerStatusGetOK describes a response with status code 200, with default header values.

OK
*/
type VscanServerStatusGetOK struct {
	Payload *models.VscanServerStatusResponse
}

// IsSuccess returns true when this vscan server status get o k response has a 2xx status code
func (o *VscanServerStatusGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan server status get o k response has a 3xx status code
func (o *VscanServerStatusGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan server status get o k response has a 4xx status code
func (o *VscanServerStatusGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan server status get o k response has a 5xx status code
func (o *VscanServerStatusGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan server status get o k response a status code equal to that given
func (o *VscanServerStatusGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vscan server status get o k response
func (o *VscanServerStatusGetOK) Code() int {
	return 200
}

func (o *VscanServerStatusGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/server-status][%d] vscanServerStatusGetOK %s", 200, payload)
}

func (o *VscanServerStatusGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/server-status][%d] vscanServerStatusGetOK %s", 200, payload)
}

func (o *VscanServerStatusGetOK) GetPayload() *models.VscanServerStatusResponse {
	return o.Payload
}

func (o *VscanServerStatusGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VscanServerStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVscanServerStatusGetDefault creates a VscanServerStatusGetDefault with default headers values
func NewVscanServerStatusGetDefault(code int) *VscanServerStatusGetDefault {
	return &VscanServerStatusGetDefault{
		_statusCode: code,
	}
}

/*
VscanServerStatusGetDefault describes a response with status code -1, with default header values.

Error
*/
type VscanServerStatusGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vscan server status get default response has a 2xx status code
func (o *VscanServerStatusGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan server status get default response has a 3xx status code
func (o *VscanServerStatusGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan server status get default response has a 4xx status code
func (o *VscanServerStatusGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan server status get default response has a 5xx status code
func (o *VscanServerStatusGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan server status get default response a status code equal to that given
func (o *VscanServerStatusGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vscan server status get default response
func (o *VscanServerStatusGetDefault) Code() int {
	return o._statusCode
}

func (o *VscanServerStatusGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/server-status][%d] vscan_server_status_get default %s", o._statusCode, payload)
}

func (o *VscanServerStatusGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/vscan/server-status][%d] vscan_server_status_get default %s", o._statusCode, payload)
}

func (o *VscanServerStatusGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanServerStatusGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
