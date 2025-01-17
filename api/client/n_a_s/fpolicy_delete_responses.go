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

// FpolicyDeleteReader is a Reader for the FpolicyDelete structure.
type FpolicyDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyDeleteOK creates a FpolicyDeleteOK with default headers values
func NewFpolicyDeleteOK() *FpolicyDeleteOK {
	return &FpolicyDeleteOK{}
}

/*
FpolicyDeleteOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyDeleteOK struct {
}

// IsSuccess returns true when this fpolicy delete o k response has a 2xx status code
func (o *FpolicyDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy delete o k response has a 3xx status code
func (o *FpolicyDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy delete o k response has a 4xx status code
func (o *FpolicyDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy delete o k response has a 5xx status code
func (o *FpolicyDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy delete o k response a status code equal to that given
func (o *FpolicyDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fpolicy delete o k response
func (o *FpolicyDeleteOK) Code() int {
	return 200
}

func (o *FpolicyDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/fpolicy/{svm.uuid}][%d] fpolicyDeleteOK", 200)
}

func (o *FpolicyDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/fpolicy/{svm.uuid}][%d] fpolicyDeleteOK", 200)
}

func (o *FpolicyDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFpolicyDeleteDefault creates a FpolicyDeleteDefault with default headers values
func NewFpolicyDeleteDefault(code int) *FpolicyDeleteDefault {
	return &FpolicyDeleteDefault{
		_statusCode: code,
	}
}

/*
	FpolicyDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 9765031    | If any of the FPolicy engine, FPolicy event or FPolicy policy deletion fails due to a systemic error or hardware failure, the cause of the failure is detailed in the error message. |
*/
type FpolicyDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fpolicy delete default response has a 2xx status code
func (o *FpolicyDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy delete default response has a 3xx status code
func (o *FpolicyDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy delete default response has a 4xx status code
func (o *FpolicyDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy delete default response has a 5xx status code
func (o *FpolicyDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy delete default response a status code equal to that given
func (o *FpolicyDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fpolicy delete default response
func (o *FpolicyDeleteDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/fpolicy/{svm.uuid}][%d] fpolicy_delete default %s", o._statusCode, payload)
}

func (o *FpolicyDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/fpolicy/{svm.uuid}][%d] fpolicy_delete default %s", o._statusCode, payload)
}

func (o *FpolicyDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
