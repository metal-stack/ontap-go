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

// NetworkEthernetBroadcastDomainsGetReader is a Reader for the NetworkEthernetBroadcastDomainsGet structure.
type NetworkEthernetBroadcastDomainsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkEthernetBroadcastDomainsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkEthernetBroadcastDomainsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkEthernetBroadcastDomainsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkEthernetBroadcastDomainsGetOK creates a NetworkEthernetBroadcastDomainsGetOK with default headers values
func NewNetworkEthernetBroadcastDomainsGetOK() *NetworkEthernetBroadcastDomainsGetOK {
	return &NetworkEthernetBroadcastDomainsGetOK{}
}

/*
NetworkEthernetBroadcastDomainsGetOK describes a response with status code 200, with default header values.

OK
*/
type NetworkEthernetBroadcastDomainsGetOK struct {
	Payload *models.BroadcastDomainResponse
}

// IsSuccess returns true when this network ethernet broadcast domains get o k response has a 2xx status code
func (o *NetworkEthernetBroadcastDomainsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network ethernet broadcast domains get o k response has a 3xx status code
func (o *NetworkEthernetBroadcastDomainsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network ethernet broadcast domains get o k response has a 4xx status code
func (o *NetworkEthernetBroadcastDomainsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network ethernet broadcast domains get o k response has a 5xx status code
func (o *NetworkEthernetBroadcastDomainsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network ethernet broadcast domains get o k response a status code equal to that given
func (o *NetworkEthernetBroadcastDomainsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the network ethernet broadcast domains get o k response
func (o *NetworkEthernetBroadcastDomainsGetOK) Code() int {
	return 200
}

func (o *NetworkEthernetBroadcastDomainsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/broadcast-domains][%d] networkEthernetBroadcastDomainsGetOK %s", 200, payload)
}

func (o *NetworkEthernetBroadcastDomainsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/broadcast-domains][%d] networkEthernetBroadcastDomainsGetOK %s", 200, payload)
}

func (o *NetworkEthernetBroadcastDomainsGetOK) GetPayload() *models.BroadcastDomainResponse {
	return o.Payload
}

func (o *NetworkEthernetBroadcastDomainsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BroadcastDomainResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkEthernetBroadcastDomainsGetDefault creates a NetworkEthernetBroadcastDomainsGetDefault with default headers values
func NewNetworkEthernetBroadcastDomainsGetDefault(code int) *NetworkEthernetBroadcastDomainsGetDefault {
	return &NetworkEthernetBroadcastDomainsGetDefault{
		_statusCode: code,
	}
}

/*
NetworkEthernetBroadcastDomainsGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetworkEthernetBroadcastDomainsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ethernet broadcast domains get default response has a 2xx status code
func (o *NetworkEthernetBroadcastDomainsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ethernet broadcast domains get default response has a 3xx status code
func (o *NetworkEthernetBroadcastDomainsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ethernet broadcast domains get default response has a 4xx status code
func (o *NetworkEthernetBroadcastDomainsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ethernet broadcast domains get default response has a 5xx status code
func (o *NetworkEthernetBroadcastDomainsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ethernet broadcast domains get default response a status code equal to that given
func (o *NetworkEthernetBroadcastDomainsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ethernet broadcast domains get default response
func (o *NetworkEthernetBroadcastDomainsGetDefault) Code() int {
	return o._statusCode
}

func (o *NetworkEthernetBroadcastDomainsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/broadcast-domains][%d] network_ethernet_broadcast_domains_get default %s", o._statusCode, payload)
}

func (o *NetworkEthernetBroadcastDomainsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ethernet/broadcast-domains][%d] network_ethernet_broadcast_domains_get default %s", o._statusCode, payload)
}

func (o *NetworkEthernetBroadcastDomainsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkEthernetBroadcastDomainsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}