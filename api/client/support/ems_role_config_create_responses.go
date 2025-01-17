// Code generated by go-swagger; DO NOT EDIT.

package support

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

// EmsRoleConfigCreateReader is a Reader for the EmsRoleConfigCreate structure.
type EmsRoleConfigCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsRoleConfigCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewEmsRoleConfigCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsRoleConfigCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsRoleConfigCreateCreated creates a EmsRoleConfigCreateCreated with default headers values
func NewEmsRoleConfigCreateCreated() *EmsRoleConfigCreateCreated {
	return &EmsRoleConfigCreateCreated{}
}

/*
EmsRoleConfigCreateCreated describes a response with status code 201, with default header values.

Created
*/
type EmsRoleConfigCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.EmsRoleConfigResponse
}

// IsSuccess returns true when this ems role config create created response has a 2xx status code
func (o *EmsRoleConfigCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ems role config create created response has a 3xx status code
func (o *EmsRoleConfigCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ems role config create created response has a 4xx status code
func (o *EmsRoleConfigCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ems role config create created response has a 5xx status code
func (o *EmsRoleConfigCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ems role config create created response a status code equal to that given
func (o *EmsRoleConfigCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ems role config create created response
func (o *EmsRoleConfigCreateCreated) Code() int {
	return 201
}

func (o *EmsRoleConfigCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/role-configs][%d] emsRoleConfigCreateCreated %s", 201, payload)
}

func (o *EmsRoleConfigCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/role-configs][%d] emsRoleConfigCreateCreated %s", 201, payload)
}

func (o *EmsRoleConfigCreateCreated) GetPayload() *models.EmsRoleConfigResponse {
	return o.Payload
}

func (o *EmsRoleConfigCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.EmsRoleConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmsRoleConfigCreateDefault creates a EmsRoleConfigCreateDefault with default headers values
func NewEmsRoleConfigCreateDefault(code int) *EmsRoleConfigCreateDefault {
	return &EmsRoleConfigCreateDefault{
		_statusCode: code,
	}
}

/*
	EmsRoleConfigCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 983198 | The event filter provided does not exist. |
| 983199 | The access control role provided does not exist. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type EmsRoleConfigCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ems role config create default response has a 2xx status code
func (o *EmsRoleConfigCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ems role config create default response has a 3xx status code
func (o *EmsRoleConfigCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ems role config create default response has a 4xx status code
func (o *EmsRoleConfigCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ems role config create default response has a 5xx status code
func (o *EmsRoleConfigCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ems role config create default response a status code equal to that given
func (o *EmsRoleConfigCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ems role config create default response
func (o *EmsRoleConfigCreateDefault) Code() int {
	return o._statusCode
}

func (o *EmsRoleConfigCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/role-configs][%d] ems_role_config_create default %s", o._statusCode, payload)
}

func (o *EmsRoleConfigCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/role-configs][%d] ems_role_config_create default %s", o._statusCode, payload)
}

func (o *EmsRoleConfigCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsRoleConfigCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
