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

// IgroupCollectionGetReader is a Reader for the IgroupCollectionGet structure.
type IgroupCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IgroupCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIgroupCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIgroupCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIgroupCollectionGetOK creates a IgroupCollectionGetOK with default headers values
func NewIgroupCollectionGetOK() *IgroupCollectionGetOK {
	return &IgroupCollectionGetOK{}
}

/*
IgroupCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type IgroupCollectionGetOK struct {
	Payload *models.IgroupResponse
}

// IsSuccess returns true when this igroup collection get o k response has a 2xx status code
func (o *IgroupCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this igroup collection get o k response has a 3xx status code
func (o *IgroupCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this igroup collection get o k response has a 4xx status code
func (o *IgroupCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this igroup collection get o k response has a 5xx status code
func (o *IgroupCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this igroup collection get o k response a status code equal to that given
func (o *IgroupCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the igroup collection get o k response
func (o *IgroupCollectionGetOK) Code() int {
	return 200
}

func (o *IgroupCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/igroups][%d] igroupCollectionGetOK %s", 200, payload)
}

func (o *IgroupCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/igroups][%d] igroupCollectionGetOK %s", 200, payload)
}

func (o *IgroupCollectionGetOK) GetPayload() *models.IgroupResponse {
	return o.Payload
}

func (o *IgroupCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IgroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIgroupCollectionGetDefault creates a IgroupCollectionGetDefault with default headers values
func NewIgroupCollectionGetDefault(code int) *IgroupCollectionGetDefault {
	return &IgroupCollectionGetDefault{
		_statusCode: code,
	}
}

/*
IgroupCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type IgroupCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this igroup collection get default response has a 2xx status code
func (o *IgroupCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this igroup collection get default response has a 3xx status code
func (o *IgroupCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this igroup collection get default response has a 4xx status code
func (o *IgroupCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this igroup collection get default response has a 5xx status code
func (o *IgroupCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this igroup collection get default response a status code equal to that given
func (o *IgroupCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the igroup collection get default response
func (o *IgroupCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *IgroupCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/igroups][%d] igroup_collection_get default %s", o._statusCode, payload)
}

func (o *IgroupCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/igroups][%d] igroup_collection_get default %s", o._statusCode, payload)
}

func (o *IgroupCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IgroupCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}