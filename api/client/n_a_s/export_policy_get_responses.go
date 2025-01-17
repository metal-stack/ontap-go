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

// ExportPolicyGetReader is a Reader for the ExportPolicyGet structure.
type ExportPolicyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportPolicyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportPolicyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportPolicyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportPolicyGetOK creates a ExportPolicyGetOK with default headers values
func NewExportPolicyGetOK() *ExportPolicyGetOK {
	return &ExportPolicyGetOK{}
}

/*
ExportPolicyGetOK describes a response with status code 200, with default header values.

OK
*/
type ExportPolicyGetOK struct {
	Payload *models.ExportPolicy
}

// IsSuccess returns true when this export policy get o k response has a 2xx status code
func (o *ExportPolicyGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export policy get o k response has a 3xx status code
func (o *ExportPolicyGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export policy get o k response has a 4xx status code
func (o *ExportPolicyGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export policy get o k response has a 5xx status code
func (o *ExportPolicyGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export policy get o k response a status code equal to that given
func (o *ExportPolicyGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the export policy get o k response
func (o *ExportPolicyGetOK) Code() int {
	return 200
}

func (o *ExportPolicyGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/export-policies/{id}][%d] exportPolicyGetOK %s", 200, payload)
}

func (o *ExportPolicyGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/export-policies/{id}][%d] exportPolicyGetOK %s", 200, payload)
}

func (o *ExportPolicyGetOK) GetPayload() *models.ExportPolicy {
	return o.Payload
}

func (o *ExportPolicyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportPolicyGetDefault creates a ExportPolicyGetDefault with default headers values
func NewExportPolicyGetDefault(code int) *ExportPolicyGetDefault {
	return &ExportPolicyGetDefault{
		_statusCode: code,
	}
}

/*
ExportPolicyGetDefault describes a response with status code -1, with default header values.

Error
*/
type ExportPolicyGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this export policy get default response has a 2xx status code
func (o *ExportPolicyGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this export policy get default response has a 3xx status code
func (o *ExportPolicyGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this export policy get default response has a 4xx status code
func (o *ExportPolicyGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this export policy get default response has a 5xx status code
func (o *ExportPolicyGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this export policy get default response a status code equal to that given
func (o *ExportPolicyGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the export policy get default response
func (o *ExportPolicyGetDefault) Code() int {
	return o._statusCode
}

func (o *ExportPolicyGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/export-policies/{id}][%d] export_policy_get default %s", o._statusCode, payload)
}

func (o *ExportPolicyGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/export-policies/{id}][%d] export_policy_get default %s", o._statusCode, payload)
}

func (o *ExportPolicyGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportPolicyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
