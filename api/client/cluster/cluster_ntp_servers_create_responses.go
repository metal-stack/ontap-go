// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// ClusterNtpServersCreateReader is a Reader for the ClusterNtpServersCreate structure.
type ClusterNtpServersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterNtpServersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewClusterNtpServersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewClusterNtpServersCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterNtpServersCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterNtpServersCreateCreated creates a ClusterNtpServersCreateCreated with default headers values
func NewClusterNtpServersCreateCreated() *ClusterNtpServersCreateCreated {
	return &ClusterNtpServersCreateCreated{}
}

/*
ClusterNtpServersCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ClusterNtpServersCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cluster ntp servers create created response has a 2xx status code
func (o *ClusterNtpServersCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster ntp servers create created response has a 3xx status code
func (o *ClusterNtpServersCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster ntp servers create created response has a 4xx status code
func (o *ClusterNtpServersCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster ntp servers create created response has a 5xx status code
func (o *ClusterNtpServersCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster ntp servers create created response a status code equal to that given
func (o *ClusterNtpServersCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the cluster ntp servers create created response
func (o *ClusterNtpServersCreateCreated) Code() int {
	return 201
}

func (o *ClusterNtpServersCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/ntp/servers][%d] clusterNtpServersCreateCreated %s", 201, payload)
}

func (o *ClusterNtpServersCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/ntp/servers][%d] clusterNtpServersCreateCreated %s", 201, payload)
}

func (o *ClusterNtpServersCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ClusterNtpServersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterNtpServersCreateAccepted creates a ClusterNtpServersCreateAccepted with default headers values
func NewClusterNtpServersCreateAccepted() *ClusterNtpServersCreateAccepted {
	return &ClusterNtpServersCreateAccepted{}
}

/*
ClusterNtpServersCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ClusterNtpServersCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cluster ntp servers create accepted response has a 2xx status code
func (o *ClusterNtpServersCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster ntp servers create accepted response has a 3xx status code
func (o *ClusterNtpServersCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster ntp servers create accepted response has a 4xx status code
func (o *ClusterNtpServersCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster ntp servers create accepted response has a 5xx status code
func (o *ClusterNtpServersCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster ntp servers create accepted response a status code equal to that given
func (o *ClusterNtpServersCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the cluster ntp servers create accepted response
func (o *ClusterNtpServersCreateAccepted) Code() int {
	return 202
}

func (o *ClusterNtpServersCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/ntp/servers][%d] clusterNtpServersCreateAccepted %s", 202, payload)
}

func (o *ClusterNtpServersCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/ntp/servers][%d] clusterNtpServersCreateAccepted %s", 202, payload)
}

func (o *ClusterNtpServersCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ClusterNtpServersCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterNtpServersCreateDefault creates a ClusterNtpServersCreateDefault with default headers values
func NewClusterNtpServersCreateDefault(code int) *ClusterNtpServersCreateDefault {
	return &ClusterNtpServersCreateDefault{
		_statusCode: code,
	}
}

/*
	ClusterNtpServersCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2097163 | NTP server IPv4 address was invalid. |
| 2097164 | NTP server IPv6 address was invalid. |
| 2097165 | Cannot resolve NTP server name. |
| 2097166 | NTP server address query returned no valid IP addresses. |
| 2097167 | Failed to connect to NTP server. |
| 2097168 | Unable to query NTP server. |
| 2097169 | NTP server provided was not synchronized with a clock or another NTP server. |
| 2097174 | NTP server provided had too high of root distance. |
| 2097177 | NTP server provided an invalid stratum. |
| 2097179 | Too many NTP servers have been configured. |
| 2097181 | NTP server address was invalid. It is a special purpose address such as loopback, multicast, or broadcast address. |
| 2097182 | NTP server address was invalid. The address is neither an IPv4 or IPv6. |
| 2097183 | NTP symmetric key authentication cannot be used for a node not in a cluster. |
| 2097185 | NTP key authentication failed for the provided key. |
| 2097193 | An unknown NTP key was provided. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ClusterNtpServersCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster ntp servers create default response has a 2xx status code
func (o *ClusterNtpServersCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster ntp servers create default response has a 3xx status code
func (o *ClusterNtpServersCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster ntp servers create default response has a 4xx status code
func (o *ClusterNtpServersCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster ntp servers create default response has a 5xx status code
func (o *ClusterNtpServersCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster ntp servers create default response a status code equal to that given
func (o *ClusterNtpServersCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster ntp servers create default response
func (o *ClusterNtpServersCreateDefault) Code() int {
	return o._statusCode
}

func (o *ClusterNtpServersCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/ntp/servers][%d] cluster_ntp_servers_create default %s", o._statusCode, payload)
}

func (o *ClusterNtpServersCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/ntp/servers][%d] cluster_ntp_servers_create default %s", o._statusCode, payload)
}

func (o *ClusterNtpServersCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterNtpServersCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}