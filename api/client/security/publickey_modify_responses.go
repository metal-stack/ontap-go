// Code generated by go-swagger; DO NOT EDIT.

package security

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

// PublickeyModifyReader is a Reader for the PublickeyModify structure.
type PublickeyModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublickeyModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublickeyModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPublickeyModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPublickeyModifyOK creates a PublickeyModifyOK with default headers values
func NewPublickeyModifyOK() *PublickeyModifyOK {
	return &PublickeyModifyOK{}
}

/*
PublickeyModifyOK describes a response with status code 200, with default header values.

OK
*/
type PublickeyModifyOK struct {
}

// IsSuccess returns true when this publickey modify o k response has a 2xx status code
func (o *PublickeyModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this publickey modify o k response has a 3xx status code
func (o *PublickeyModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this publickey modify o k response has a 4xx status code
func (o *PublickeyModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this publickey modify o k response has a 5xx status code
func (o *PublickeyModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this publickey modify o k response a status code equal to that given
func (o *PublickeyModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the publickey modify o k response
func (o *PublickeyModifyOK) Code() int {
	return 200
}

func (o *PublickeyModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /security/authentication/publickeys/{owner.uuid}/{account.name}/{index}][%d] publickeyModifyOK", 200)
}

func (o *PublickeyModifyOK) String() string {
	return fmt.Sprintf("[PATCH /security/authentication/publickeys/{owner.uuid}/{account.name}/{index}][%d] publickeyModifyOK", 200)
}

func (o *PublickeyModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublickeyModifyDefault creates a PublickeyModifyDefault with default headers values
func NewPublickeyModifyDefault(code int) *PublickeyModifyDefault {
	return &PublickeyModifyDefault{
		_statusCode: code,
	}
}

/*
	PublickeyModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5832707 | Failed to generate fingerprint for the public key. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type PublickeyModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this publickey modify default response has a 2xx status code
func (o *PublickeyModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this publickey modify default response has a 3xx status code
func (o *PublickeyModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this publickey modify default response has a 4xx status code
func (o *PublickeyModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this publickey modify default response has a 5xx status code
func (o *PublickeyModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this publickey modify default response a status code equal to that given
func (o *PublickeyModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the publickey modify default response
func (o *PublickeyModifyDefault) Code() int {
	return o._statusCode
}

func (o *PublickeyModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/authentication/publickeys/{owner.uuid}/{account.name}/{index}][%d] publickey_modify default %s", o._statusCode, payload)
}

func (o *PublickeyModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/authentication/publickeys/{owner.uuid}/{account.name}/{index}][%d] publickey_modify default %s", o._statusCode, payload)
}

func (o *PublickeyModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PublickeyModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}