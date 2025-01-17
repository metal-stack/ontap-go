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

// NetworkIPBgpPeerGroupDeleteReader is a Reader for the NetworkIPBgpPeerGroupDelete structure.
type NetworkIPBgpPeerGroupDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkIPBgpPeerGroupDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkIPBgpPeerGroupDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkIPBgpPeerGroupDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkIPBgpPeerGroupDeleteOK creates a NetworkIPBgpPeerGroupDeleteOK with default headers values
func NewNetworkIPBgpPeerGroupDeleteOK() *NetworkIPBgpPeerGroupDeleteOK {
	return &NetworkIPBgpPeerGroupDeleteOK{}
}

/*
NetworkIPBgpPeerGroupDeleteOK describes a response with status code 200, with default header values.

OK
*/
type NetworkIPBgpPeerGroupDeleteOK struct {
}

// IsSuccess returns true when this network Ip bgp peer group delete o k response has a 2xx status code
func (o *NetworkIPBgpPeerGroupDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network Ip bgp peer group delete o k response has a 3xx status code
func (o *NetworkIPBgpPeerGroupDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network Ip bgp peer group delete o k response has a 4xx status code
func (o *NetworkIPBgpPeerGroupDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network Ip bgp peer group delete o k response has a 5xx status code
func (o *NetworkIPBgpPeerGroupDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network Ip bgp peer group delete o k response a status code equal to that given
func (o *NetworkIPBgpPeerGroupDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the network Ip bgp peer group delete o k response
func (o *NetworkIPBgpPeerGroupDeleteOK) Code() int {
	return 200
}

func (o *NetworkIPBgpPeerGroupDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /network/ip/bgp/peer-groups/{uuid}][%d] networkIpBgpPeerGroupDeleteOK", 200)
}

func (o *NetworkIPBgpPeerGroupDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /network/ip/bgp/peer-groups/{uuid}][%d] networkIpBgpPeerGroupDeleteOK", 200)
}

func (o *NetworkIPBgpPeerGroupDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkIPBgpPeerGroupDeleteDefault creates a NetworkIPBgpPeerGroupDeleteDefault with default headers values
func NewNetworkIPBgpPeerGroupDeleteDefault(code int) *NetworkIPBgpPeerGroupDeleteDefault {
	return &NetworkIPBgpPeerGroupDeleteDefault{
		_statusCode: code,
	}
}

/*
	NetworkIPBgpPeerGroupDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 53282019 | Internal error. Failed to remove BGP peer group on node. Wait a few minutes and try the command again. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NetworkIPBgpPeerGroupDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ip bgp peer group delete default response has a 2xx status code
func (o *NetworkIPBgpPeerGroupDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ip bgp peer group delete default response has a 3xx status code
func (o *NetworkIPBgpPeerGroupDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ip bgp peer group delete default response has a 4xx status code
func (o *NetworkIPBgpPeerGroupDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ip bgp peer group delete default response has a 5xx status code
func (o *NetworkIPBgpPeerGroupDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ip bgp peer group delete default response a status code equal to that given
func (o *NetworkIPBgpPeerGroupDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ip bgp peer group delete default response
func (o *NetworkIPBgpPeerGroupDeleteDefault) Code() int {
	return o._statusCode
}

func (o *NetworkIPBgpPeerGroupDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ip/bgp/peer-groups/{uuid}][%d] network_ip_bgp_peer_group_delete default %s", o._statusCode, payload)
}

func (o *NetworkIPBgpPeerGroupDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ip/bgp/peer-groups/{uuid}][%d] network_ip_bgp_peer_group_delete default %s", o._statusCode, payload)
}

func (o *NetworkIPBgpPeerGroupDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkIPBgpPeerGroupDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
