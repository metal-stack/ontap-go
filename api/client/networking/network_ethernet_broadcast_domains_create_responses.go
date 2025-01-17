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

// NetworkEthernetBroadcastDomainsCreateReader is a Reader for the NetworkEthernetBroadcastDomainsCreate structure.
type NetworkEthernetBroadcastDomainsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkEthernetBroadcastDomainsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNetworkEthernetBroadcastDomainsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkEthernetBroadcastDomainsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkEthernetBroadcastDomainsCreateCreated creates a NetworkEthernetBroadcastDomainsCreateCreated with default headers values
func NewNetworkEthernetBroadcastDomainsCreateCreated() *NetworkEthernetBroadcastDomainsCreateCreated {
	return &NetworkEthernetBroadcastDomainsCreateCreated{}
}

/*
NetworkEthernetBroadcastDomainsCreateCreated describes a response with status code 201, with default header values.

Created
*/
type NetworkEthernetBroadcastDomainsCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this network ethernet broadcast domains create created response has a 2xx status code
func (o *NetworkEthernetBroadcastDomainsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network ethernet broadcast domains create created response has a 3xx status code
func (o *NetworkEthernetBroadcastDomainsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network ethernet broadcast domains create created response has a 4xx status code
func (o *NetworkEthernetBroadcastDomainsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this network ethernet broadcast domains create created response has a 5xx status code
func (o *NetworkEthernetBroadcastDomainsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this network ethernet broadcast domains create created response a status code equal to that given
func (o *NetworkEthernetBroadcastDomainsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the network ethernet broadcast domains create created response
func (o *NetworkEthernetBroadcastDomainsCreateCreated) Code() int {
	return 201
}

func (o *NetworkEthernetBroadcastDomainsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /network/ethernet/broadcast-domains][%d] networkEthernetBroadcastDomainsCreateCreated", 201)
}

func (o *NetworkEthernetBroadcastDomainsCreateCreated) String() string {
	return fmt.Sprintf("[POST /network/ethernet/broadcast-domains][%d] networkEthernetBroadcastDomainsCreateCreated", 201)
}

func (o *NetworkEthernetBroadcastDomainsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewNetworkEthernetBroadcastDomainsCreateDefault creates a NetworkEthernetBroadcastDomainsCreateDefault with default headers values
func NewNetworkEthernetBroadcastDomainsCreateDefault(code int) *NetworkEthernetBroadcastDomainsCreateDefault {
	return &NetworkEthernetBroadcastDomainsCreateDefault{
		_statusCode: code,
	}
}

/*
	NetworkEthernetBroadcastDomainsCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1377267 | The specified IPspace does not exist. |
| 1377596 | The Cluster IPspace cannot contain more than one broadcast domain. Modifying the system IPspace Cluster is not supported. |
| 1966460 | The provided MTU was either too large or too small. |
| 1967082 | The specified ipspace.name does not match the IPspace name of ipspace.uuid. |
| 1967102 | A POST operation might have left the configuration in an inconsistent state. Check the configuration. |
| 1967104 | Invalid IPspace UUID. |
| 53281982 | The specified broadcast domain name is reserved by the system. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NetworkEthernetBroadcastDomainsCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ethernet broadcast domains create default response has a 2xx status code
func (o *NetworkEthernetBroadcastDomainsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ethernet broadcast domains create default response has a 3xx status code
func (o *NetworkEthernetBroadcastDomainsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ethernet broadcast domains create default response has a 4xx status code
func (o *NetworkEthernetBroadcastDomainsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ethernet broadcast domains create default response has a 5xx status code
func (o *NetworkEthernetBroadcastDomainsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ethernet broadcast domains create default response a status code equal to that given
func (o *NetworkEthernetBroadcastDomainsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ethernet broadcast domains create default response
func (o *NetworkEthernetBroadcastDomainsCreateDefault) Code() int {
	return o._statusCode
}

func (o *NetworkEthernetBroadcastDomainsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/ethernet/broadcast-domains][%d] network_ethernet_broadcast_domains_create default %s", o._statusCode, payload)
}

func (o *NetworkEthernetBroadcastDomainsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/ethernet/broadcast-domains][%d] network_ethernet_broadcast_domains_create default %s", o._statusCode, payload)
}

func (o *NetworkEthernetBroadcastDomainsCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkEthernetBroadcastDomainsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
