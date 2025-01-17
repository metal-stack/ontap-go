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

// ClusterNtpServersDeleteReader is a Reader for the ClusterNtpServersDelete structure.
type ClusterNtpServersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterNtpServersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterNtpServersDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewClusterNtpServersDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterNtpServersDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterNtpServersDeleteOK creates a ClusterNtpServersDeleteOK with default headers values
func NewClusterNtpServersDeleteOK() *ClusterNtpServersDeleteOK {
	return &ClusterNtpServersDeleteOK{}
}

/*
ClusterNtpServersDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ClusterNtpServersDeleteOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cluster ntp servers delete o k response has a 2xx status code
func (o *ClusterNtpServersDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster ntp servers delete o k response has a 3xx status code
func (o *ClusterNtpServersDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster ntp servers delete o k response has a 4xx status code
func (o *ClusterNtpServersDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster ntp servers delete o k response has a 5xx status code
func (o *ClusterNtpServersDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster ntp servers delete o k response a status code equal to that given
func (o *ClusterNtpServersDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster ntp servers delete o k response
func (o *ClusterNtpServersDeleteOK) Code() int {
	return 200
}

func (o *ClusterNtpServersDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/ntp/servers/{server}][%d] clusterNtpServersDeleteOK %s", 200, payload)
}

func (o *ClusterNtpServersDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/ntp/servers/{server}][%d] clusterNtpServersDeleteOK %s", 200, payload)
}

func (o *ClusterNtpServersDeleteOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ClusterNtpServersDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterNtpServersDeleteAccepted creates a ClusterNtpServersDeleteAccepted with default headers values
func NewClusterNtpServersDeleteAccepted() *ClusterNtpServersDeleteAccepted {
	return &ClusterNtpServersDeleteAccepted{}
}

/*
ClusterNtpServersDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ClusterNtpServersDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cluster ntp servers delete accepted response has a 2xx status code
func (o *ClusterNtpServersDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster ntp servers delete accepted response has a 3xx status code
func (o *ClusterNtpServersDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster ntp servers delete accepted response has a 4xx status code
func (o *ClusterNtpServersDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster ntp servers delete accepted response has a 5xx status code
func (o *ClusterNtpServersDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster ntp servers delete accepted response a status code equal to that given
func (o *ClusterNtpServersDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the cluster ntp servers delete accepted response
func (o *ClusterNtpServersDeleteAccepted) Code() int {
	return 202
}

func (o *ClusterNtpServersDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/ntp/servers/{server}][%d] clusterNtpServersDeleteAccepted %s", 202, payload)
}

func (o *ClusterNtpServersDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/ntp/servers/{server}][%d] clusterNtpServersDeleteAccepted %s", 202, payload)
}

func (o *ClusterNtpServersDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ClusterNtpServersDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterNtpServersDeleteDefault creates a ClusterNtpServersDeleteDefault with default headers values
func NewClusterNtpServersDeleteDefault(code int) *ClusterNtpServersDeleteDefault {
	return &ClusterNtpServersDeleteDefault{
		_statusCode: code,
	}
}

/*
ClusterNtpServersDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterNtpServersDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster ntp servers delete default response has a 2xx status code
func (o *ClusterNtpServersDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster ntp servers delete default response has a 3xx status code
func (o *ClusterNtpServersDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster ntp servers delete default response has a 4xx status code
func (o *ClusterNtpServersDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster ntp servers delete default response has a 5xx status code
func (o *ClusterNtpServersDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster ntp servers delete default response a status code equal to that given
func (o *ClusterNtpServersDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster ntp servers delete default response
func (o *ClusterNtpServersDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ClusterNtpServersDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/ntp/servers/{server}][%d] cluster_ntp_servers_delete default %s", o._statusCode, payload)
}

func (o *ClusterNtpServersDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/ntp/servers/{server}][%d] cluster_ntp_servers_delete default %s", o._statusCode, payload)
}

func (o *ClusterNtpServersDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterNtpServersDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}