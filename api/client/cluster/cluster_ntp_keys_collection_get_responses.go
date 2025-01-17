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

// ClusterNtpKeysCollectionGetReader is a Reader for the ClusterNtpKeysCollectionGet structure.
type ClusterNtpKeysCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterNtpKeysCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterNtpKeysCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterNtpKeysCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterNtpKeysCollectionGetOK creates a ClusterNtpKeysCollectionGetOK with default headers values
func NewClusterNtpKeysCollectionGetOK() *ClusterNtpKeysCollectionGetOK {
	return &ClusterNtpKeysCollectionGetOK{}
}

/*
ClusterNtpKeysCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ClusterNtpKeysCollectionGetOK struct {
	Payload *models.NtpKeyResponse
}

// IsSuccess returns true when this cluster ntp keys collection get o k response has a 2xx status code
func (o *ClusterNtpKeysCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster ntp keys collection get o k response has a 3xx status code
func (o *ClusterNtpKeysCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster ntp keys collection get o k response has a 4xx status code
func (o *ClusterNtpKeysCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster ntp keys collection get o k response has a 5xx status code
func (o *ClusterNtpKeysCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster ntp keys collection get o k response a status code equal to that given
func (o *ClusterNtpKeysCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster ntp keys collection get o k response
func (o *ClusterNtpKeysCollectionGetOK) Code() int {
	return 200
}

func (o *ClusterNtpKeysCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/ntp/keys][%d] clusterNtpKeysCollectionGetOK %s", 200, payload)
}

func (o *ClusterNtpKeysCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/ntp/keys][%d] clusterNtpKeysCollectionGetOK %s", 200, payload)
}

func (o *ClusterNtpKeysCollectionGetOK) GetPayload() *models.NtpKeyResponse {
	return o.Payload
}

func (o *ClusterNtpKeysCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NtpKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterNtpKeysCollectionGetDefault creates a ClusterNtpKeysCollectionGetDefault with default headers values
func NewClusterNtpKeysCollectionGetDefault(code int) *ClusterNtpKeysCollectionGetDefault {
	return &ClusterNtpKeysCollectionGetDefault{
		_statusCode: code,
	}
}

/*
ClusterNtpKeysCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterNtpKeysCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster ntp keys collection get default response has a 2xx status code
func (o *ClusterNtpKeysCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster ntp keys collection get default response has a 3xx status code
func (o *ClusterNtpKeysCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster ntp keys collection get default response has a 4xx status code
func (o *ClusterNtpKeysCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster ntp keys collection get default response has a 5xx status code
func (o *ClusterNtpKeysCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster ntp keys collection get default response a status code equal to that given
func (o *ClusterNtpKeysCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster ntp keys collection get default response
func (o *ClusterNtpKeysCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ClusterNtpKeysCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/ntp/keys][%d] cluster_ntp_keys_collection_get default %s", o._statusCode, payload)
}

func (o *ClusterNtpKeysCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/ntp/keys][%d] cluster_ntp_keys_collection_get default %s", o._statusCode, payload)
}

func (o *ClusterNtpKeysCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterNtpKeysCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
