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

// IPSubnetCollectionGetReader is a Reader for the IPSubnetCollectionGet structure.
type IPSubnetCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPSubnetCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIPSubnetCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIPSubnetCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIPSubnetCollectionGetOK creates a IPSubnetCollectionGetOK with default headers values
func NewIPSubnetCollectionGetOK() *IPSubnetCollectionGetOK {
	return &IPSubnetCollectionGetOK{}
}

/*
IPSubnetCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type IPSubnetCollectionGetOK struct {
	Payload *models.IPSubnetResponse
}

// IsSuccess returns true when this ip subnet collection get o k response has a 2xx status code
func (o *IPSubnetCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ip subnet collection get o k response has a 3xx status code
func (o *IPSubnetCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ip subnet collection get o k response has a 4xx status code
func (o *IPSubnetCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ip subnet collection get o k response has a 5xx status code
func (o *IPSubnetCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ip subnet collection get o k response a status code equal to that given
func (o *IPSubnetCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ip subnet collection get o k response
func (o *IPSubnetCollectionGetOK) Code() int {
	return 200
}

func (o *IPSubnetCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/subnets][%d] ipSubnetCollectionGetOK %s", 200, payload)
}

func (o *IPSubnetCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/subnets][%d] ipSubnetCollectionGetOK %s", 200, payload)
}

func (o *IPSubnetCollectionGetOK) GetPayload() *models.IPSubnetResponse {
	return o.Payload
}

func (o *IPSubnetCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPSubnetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIPSubnetCollectionGetDefault creates a IPSubnetCollectionGetDefault with default headers values
func NewIPSubnetCollectionGetDefault(code int) *IPSubnetCollectionGetDefault {
	return &IPSubnetCollectionGetDefault{
		_statusCode: code,
	}
}

/*
IPSubnetCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type IPSubnetCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ip subnet collection get default response has a 2xx status code
func (o *IPSubnetCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ip subnet collection get default response has a 3xx status code
func (o *IPSubnetCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ip subnet collection get default response has a 4xx status code
func (o *IPSubnetCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ip subnet collection get default response has a 5xx status code
func (o *IPSubnetCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ip subnet collection get default response a status code equal to that given
func (o *IPSubnetCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ip subnet collection get default response
func (o *IPSubnetCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *IPSubnetCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/subnets][%d] ip_subnet_collection_get default %s", o._statusCode, payload)
}

func (o *IPSubnetCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/subnets][%d] ip_subnet_collection_get default %s", o._statusCode, payload)
}

func (o *IPSubnetCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IPSubnetCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}