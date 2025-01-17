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

// EmsConfigModifyReader is a Reader for the EmsConfigModify structure.
type EmsConfigModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsConfigModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmsConfigModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsConfigModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsConfigModifyOK creates a EmsConfigModifyOK with default headers values
func NewEmsConfigModifyOK() *EmsConfigModifyOK {
	return &EmsConfigModifyOK{}
}

/*
EmsConfigModifyOK describes a response with status code 200, with default header values.

OK
*/
type EmsConfigModifyOK struct {
}

// IsSuccess returns true when this ems config modify o k response has a 2xx status code
func (o *EmsConfigModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ems config modify o k response has a 3xx status code
func (o *EmsConfigModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ems config modify o k response has a 4xx status code
func (o *EmsConfigModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ems config modify o k response has a 5xx status code
func (o *EmsConfigModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ems config modify o k response a status code equal to that given
func (o *EmsConfigModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ems config modify o k response
func (o *EmsConfigModifyOK) Code() int {
	return 200
}

func (o *EmsConfigModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /support/ems][%d] emsConfigModifyOK", 200)
}

func (o *EmsConfigModifyOK) String() string {
	return fmt.Sprintf("[PATCH /support/ems][%d] emsConfigModifyOK", 200)
}

func (o *EmsConfigModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmsConfigModifyDefault creates a EmsConfigModifyDefault with default headers values
func NewEmsConfigModifyDefault(code int) *EmsConfigModifyDefault {
	return &EmsConfigModifyDefault{
		_statusCode: code,
	}
}

/*
	EmsConfigModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 983123 | The validation of the mail server provided failed |
| 983136 | The proxy URL cannot contain a username or password |
| 983137 | The proxy URL provided is invalid |
| 983139 | The IPv6 proxy URL provided is invalid |
| 983140 | The proxy URL provided contains an invalid scheme. The only supported scheme is 'http' |
| 983160 | The provided parameter requires an effective cluster version of ONTAP 9.9.1 or later |
| 983220 | The proxy URL provided contains an invalid scheme. The only supported schemes are 'http' and 'https' |
| 983224 | The mail server URL cannot contain a username or password |
| 983225 | The mail server username cannot be set without a mail server URL |
| 983226 | The mail server password cannot be set without a mail server username |
| 983227 | The provided field requires an effective cluster version of ONTAP 9.15.1 or later |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type EmsConfigModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ems config modify default response has a 2xx status code
func (o *EmsConfigModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ems config modify default response has a 3xx status code
func (o *EmsConfigModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ems config modify default response has a 4xx status code
func (o *EmsConfigModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ems config modify default response has a 5xx status code
func (o *EmsConfigModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ems config modify default response a status code equal to that given
func (o *EmsConfigModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ems config modify default response
func (o *EmsConfigModifyDefault) Code() int {
	return o._statusCode
}

func (o *EmsConfigModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /support/ems][%d] ems_config_modify default %s", o._statusCode, payload)
}

func (o *EmsConfigModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /support/ems][%d] ems_config_modify default %s", o._statusCode, payload)
}

func (o *EmsConfigModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsConfigModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
