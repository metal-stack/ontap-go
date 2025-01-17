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

// NetworkEthernetPortGetReader is a Reader for the NetworkEthernetPortGet structure.
type NetworkEthernetPortGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkEthernetPortGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkEthernetPortGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkEthernetPortGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkEthernetPortGetOK creates a NetworkEthernetPortGetOK with default headers values
func NewNetworkEthernetPortGetOK() *NetworkEthernetPortGetOK {
	return &NetworkEthernetPortGetOK{}
}

/*
NetworkEthernetPortGetOK describes a response with status code 200, with default header values.

OK
*/
type NetworkEthernetPortGetOK struct {
	Payload *models.Port
}

// IsSuccess returns true when this network ethernet port get o k response has a 2xx status code
func (o *NetworkEthernetPortGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network ethernet port get o k response has a 3xx status code
func (o *NetworkEthernetPortGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network ethernet port get o k response has a 4xx status code
func (o *NetworkEthernetPortGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network ethernet port get o k response has a 5xx status code
func (o *NetworkEthernetPortGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network ethernet port get o k response a status code equal to that given
func (o *NetworkEthernetPortGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the network ethernet port get o k response
func (o *NetworkEthernetPortGetOK) Code() int {
	return 200
}

func (o *NetworkEthernetPortGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/ports/{uuid}][%d] networkEthernetPortGetOK %s", 200, payload)
}

func (o *NetworkEthernetPortGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/ports/{uuid}][%d] networkEthernetPortGetOK %s", 200, payload)
}

func (o *NetworkEthernetPortGetOK) GetPayload() *models.Port {
	return o.Payload
}

func (o *NetworkEthernetPortGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Port)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkEthernetPortGetDefault creates a NetworkEthernetPortGetDefault with default headers values
func NewNetworkEthernetPortGetDefault(code int) *NetworkEthernetPortGetDefault {
	return &NetworkEthernetPortGetDefault{
		_statusCode: code,
	}
}

/*
NetworkEthernetPortGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetworkEthernetPortGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ethernet port get default response has a 2xx status code
func (o *NetworkEthernetPortGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ethernet port get default response has a 3xx status code
func (o *NetworkEthernetPortGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ethernet port get default response has a 4xx status code
func (o *NetworkEthernetPortGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ethernet port get default response has a 5xx status code
func (o *NetworkEthernetPortGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ethernet port get default response a status code equal to that given
func (o *NetworkEthernetPortGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ethernet port get default response
func (o *NetworkEthernetPortGetDefault) Code() int {
	return o._statusCode
}

func (o *NetworkEthernetPortGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/ports/{uuid}][%d] network_ethernet_port_get default %s", o._statusCode, payload)
}

func (o *NetworkEthernetPortGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/ports/{uuid}][%d] network_ethernet_port_get default %s", o._statusCode, payload)
}

func (o *NetworkEthernetPortGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkEthernetPortGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}