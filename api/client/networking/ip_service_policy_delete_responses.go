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

// IPServicePolicyDeleteReader is a Reader for the IPServicePolicyDelete structure.
type IPServicePolicyDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPServicePolicyDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIPServicePolicyDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIPServicePolicyDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIPServicePolicyDeleteOK creates a IPServicePolicyDeleteOK with default headers values
func NewIPServicePolicyDeleteOK() *IPServicePolicyDeleteOK {
	return &IPServicePolicyDeleteOK{}
}

/*
IPServicePolicyDeleteOK describes a response with status code 200, with default header values.

OK
*/
type IPServicePolicyDeleteOK struct {
}

// IsSuccess returns true when this ip service policy delete o k response has a 2xx status code
func (o *IPServicePolicyDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ip service policy delete o k response has a 3xx status code
func (o *IPServicePolicyDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ip service policy delete o k response has a 4xx status code
func (o *IPServicePolicyDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ip service policy delete o k response has a 5xx status code
func (o *IPServicePolicyDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ip service policy delete o k response a status code equal to that given
func (o *IPServicePolicyDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ip service policy delete o k response
func (o *IPServicePolicyDeleteOK) Code() int {
	return 200
}

func (o *IPServicePolicyDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /network/ip/service-policies/{uuid}][%d] ipServicePolicyDeleteOK", 200)
}

func (o *IPServicePolicyDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /network/ip/service-policies/{uuid}][%d] ipServicePolicyDeleteOK", 200)
}

func (o *IPServicePolicyDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIPServicePolicyDeleteDefault creates a IPServicePolicyDeleteDefault with default headers values
func NewIPServicePolicyDeleteDefault(code int) *IPServicePolicyDeleteDefault {
	return &IPServicePolicyDeleteDefault{
		_statusCode: code,
	}
}

/*
	IPServicePolicyDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621740 | An unexpected error when trying to determine whether the target Vserver was locked or not on this cluster. |
| 53281927 | Service policies owned by the system cannot be deleted. |
| 53281928 | Service policies assigned to LIFs cannot be deleted. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type IPServicePolicyDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ip service policy delete default response has a 2xx status code
func (o *IPServicePolicyDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ip service policy delete default response has a 3xx status code
func (o *IPServicePolicyDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ip service policy delete default response has a 4xx status code
func (o *IPServicePolicyDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ip service policy delete default response has a 5xx status code
func (o *IPServicePolicyDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ip service policy delete default response a status code equal to that given
func (o *IPServicePolicyDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ip service policy delete default response
func (o *IPServicePolicyDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IPServicePolicyDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ip/service-policies/{uuid}][%d] ip_service_policy_delete default %s", o._statusCode, payload)
}

func (o *IPServicePolicyDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ip/service-policies/{uuid}][%d] ip_service_policy_delete default %s", o._statusCode, payload)
}

func (o *IPServicePolicyDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IPServicePolicyDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}