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

// ClusterNtpKeysModifyReader is a Reader for the ClusterNtpKeysModify structure.
type ClusterNtpKeysModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterNtpKeysModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterNtpKeysModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterNtpKeysModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterNtpKeysModifyOK creates a ClusterNtpKeysModifyOK with default headers values
func NewClusterNtpKeysModifyOK() *ClusterNtpKeysModifyOK {
	return &ClusterNtpKeysModifyOK{}
}

/*
ClusterNtpKeysModifyOK describes a response with status code 200, with default header values.

OK
*/
type ClusterNtpKeysModifyOK struct {
}

// IsSuccess returns true when this cluster ntp keys modify o k response has a 2xx status code
func (o *ClusterNtpKeysModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster ntp keys modify o k response has a 3xx status code
func (o *ClusterNtpKeysModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster ntp keys modify o k response has a 4xx status code
func (o *ClusterNtpKeysModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster ntp keys modify o k response has a 5xx status code
func (o *ClusterNtpKeysModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster ntp keys modify o k response a status code equal to that given
func (o *ClusterNtpKeysModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster ntp keys modify o k response
func (o *ClusterNtpKeysModifyOK) Code() int {
	return 200
}

func (o *ClusterNtpKeysModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /cluster/ntp/keys/{id}][%d] clusterNtpKeysModifyOK", 200)
}

func (o *ClusterNtpKeysModifyOK) String() string {
	return fmt.Sprintf("[PATCH /cluster/ntp/keys/{id}][%d] clusterNtpKeysModifyOK", 200)
}

func (o *ClusterNtpKeysModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClusterNtpKeysModifyDefault creates a ClusterNtpKeysModifyDefault with default headers values
func NewClusterNtpKeysModifyDefault(code int) *ClusterNtpKeysModifyDefault {
	return &ClusterNtpKeysModifyDefault{
		_statusCode: code,
	}
}

/*
	ClusterNtpKeysModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2097187 | An invalid SHA1 key was provided. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ClusterNtpKeysModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster ntp keys modify default response has a 2xx status code
func (o *ClusterNtpKeysModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster ntp keys modify default response has a 3xx status code
func (o *ClusterNtpKeysModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster ntp keys modify default response has a 4xx status code
func (o *ClusterNtpKeysModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster ntp keys modify default response has a 5xx status code
func (o *ClusterNtpKeysModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster ntp keys modify default response a status code equal to that given
func (o *ClusterNtpKeysModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster ntp keys modify default response
func (o *ClusterNtpKeysModifyDefault) Code() int {
	return o._statusCode
}

func (o *ClusterNtpKeysModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/ntp/keys/{id}][%d] cluster_ntp_keys_modify default %s", o._statusCode, payload)
}

func (o *ClusterNtpKeysModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/ntp/keys/{id}][%d] cluster_ntp_keys_modify default %s", o._statusCode, payload)
}

func (o *ClusterNtpKeysModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterNtpKeysModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}