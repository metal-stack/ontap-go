// Code generated by go-swagger; DO NOT EDIT.

package security

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

// ClusterNisDeleteReader is a Reader for the ClusterNisDelete structure.
type ClusterNisDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterNisDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterNisDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterNisDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterNisDeleteOK creates a ClusterNisDeleteOK with default headers values
func NewClusterNisDeleteOK() *ClusterNisDeleteOK {
	return &ClusterNisDeleteOK{}
}

/*
ClusterNisDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ClusterNisDeleteOK struct {
}

// IsSuccess returns true when this cluster nis delete o k response has a 2xx status code
func (o *ClusterNisDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster nis delete o k response has a 3xx status code
func (o *ClusterNisDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster nis delete o k response has a 4xx status code
func (o *ClusterNisDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster nis delete o k response has a 5xx status code
func (o *ClusterNisDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster nis delete o k response a status code equal to that given
func (o *ClusterNisDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster nis delete o k response
func (o *ClusterNisDeleteOK) Code() int {
	return 200
}

func (o *ClusterNisDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/authentication/cluster/nis][%d] clusterNisDeleteOK", 200)
}

func (o *ClusterNisDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/authentication/cluster/nis][%d] clusterNisDeleteOK", 200)
}

func (o *ClusterNisDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClusterNisDeleteDefault creates a ClusterNisDeleteDefault with default headers values
func NewClusterNisDeleteDefault(code int) *ClusterNisDeleteDefault {
	return &ClusterNisDeleteDefault{
		_statusCode: code,
	}
}

/*
ClusterNisDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterNisDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster nis delete default response has a 2xx status code
func (o *ClusterNisDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster nis delete default response has a 3xx status code
func (o *ClusterNisDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster nis delete default response has a 4xx status code
func (o *ClusterNisDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster nis delete default response has a 5xx status code
func (o *ClusterNisDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster nis delete default response a status code equal to that given
func (o *ClusterNisDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster nis delete default response
func (o *ClusterNisDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ClusterNisDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/authentication/cluster/nis][%d] cluster_nis_delete default %s", o._statusCode, payload)
}

func (o *ClusterNisDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/authentication/cluster/nis][%d] cluster_nis_delete default %s", o._statusCode, payload)
}

func (o *ClusterNisDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterNisDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
