// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// FcLoginGetReader is a Reader for the FcLoginGet structure.
type FcLoginGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FcLoginGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFcLoginGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFcLoginGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFcLoginGetOK creates a FcLoginGetOK with default headers values
func NewFcLoginGetOK() *FcLoginGetOK {
	return &FcLoginGetOK{}
}

/*
FcLoginGetOK describes a response with status code 200, with default header values.

OK
*/
type FcLoginGetOK struct {
	Payload *models.FcLogin
}

// IsSuccess returns true when this fc login get o k response has a 2xx status code
func (o *FcLoginGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fc login get o k response has a 3xx status code
func (o *FcLoginGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fc login get o k response has a 4xx status code
func (o *FcLoginGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fc login get o k response has a 5xx status code
func (o *FcLoginGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fc login get o k response a status code equal to that given
func (o *FcLoginGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fc login get o k response
func (o *FcLoginGetOK) Code() int {
	return 200
}

func (o *FcLoginGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/logins/{interface.uuid}/{initiator.wwpn}][%d] fcLoginGetOK %s", 200, payload)
}

func (o *FcLoginGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/logins/{interface.uuid}/{initiator.wwpn}][%d] fcLoginGetOK %s", 200, payload)
}

func (o *FcLoginGetOK) GetPayload() *models.FcLogin {
	return o.Payload
}

func (o *FcLoginGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FcLogin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFcLoginGetDefault creates a FcLoginGetDefault with default headers values
func NewFcLoginGetDefault(code int) *FcLoginGetDefault {
	return &FcLoginGetDefault{
		_statusCode: code,
	}
}

/*
	FcLoginGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 4 | The Fibre Channel login specified does not exist. |
| 5373982 | An invalid WWN was specified. The length is incorrect. |
| 5373983 | An invalid WWN was specified. The format is incorrect. |
| 5374881 | The Fibre Channel interface specified does not exist. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type FcLoginGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fc login get default response has a 2xx status code
func (o *FcLoginGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fc login get default response has a 3xx status code
func (o *FcLoginGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fc login get default response has a 4xx status code
func (o *FcLoginGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fc login get default response has a 5xx status code
func (o *FcLoginGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fc login get default response a status code equal to that given
func (o *FcLoginGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fc login get default response
func (o *FcLoginGetDefault) Code() int {
	return o._statusCode
}

func (o *FcLoginGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/logins/{interface.uuid}/{initiator.wwpn}][%d] fc_login_get default %s", o._statusCode, payload)
}

func (o *FcLoginGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/logins/{interface.uuid}/{initiator.wwpn}][%d] fc_login_get default %s", o._statusCode, payload)
}

func (o *FcLoginGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FcLoginGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
