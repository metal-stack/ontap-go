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

// ClusterAccountAdProxyDeleteReader is a Reader for the ClusterAccountAdProxyDelete structure.
type ClusterAccountAdProxyDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterAccountAdProxyDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterAccountAdProxyDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterAccountAdProxyDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterAccountAdProxyDeleteOK creates a ClusterAccountAdProxyDeleteOK with default headers values
func NewClusterAccountAdProxyDeleteOK() *ClusterAccountAdProxyDeleteOK {
	return &ClusterAccountAdProxyDeleteOK{}
}

/*
ClusterAccountAdProxyDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ClusterAccountAdProxyDeleteOK struct {
}

// IsSuccess returns true when this cluster account ad proxy delete o k response has a 2xx status code
func (o *ClusterAccountAdProxyDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster account ad proxy delete o k response has a 3xx status code
func (o *ClusterAccountAdProxyDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster account ad proxy delete o k response has a 4xx status code
func (o *ClusterAccountAdProxyDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster account ad proxy delete o k response has a 5xx status code
func (o *ClusterAccountAdProxyDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster account ad proxy delete o k response a status code equal to that given
func (o *ClusterAccountAdProxyDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster account ad proxy delete o k response
func (o *ClusterAccountAdProxyDeleteOK) Code() int {
	return 200
}

func (o *ClusterAccountAdProxyDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/authentication/cluster/ad-proxy][%d] clusterAccountAdProxyDeleteOK", 200)
}

func (o *ClusterAccountAdProxyDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/authentication/cluster/ad-proxy][%d] clusterAccountAdProxyDeleteOK", 200)
}

func (o *ClusterAccountAdProxyDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClusterAccountAdProxyDeleteDefault creates a ClusterAccountAdProxyDeleteDefault with default headers values
func NewClusterAccountAdProxyDeleteDefault(code int) *ClusterAccountAdProxyDeleteDefault {
	return &ClusterAccountAdProxyDeleteDefault{
		_statusCode: code,
	}
}

/*
ClusterAccountAdProxyDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterAccountAdProxyDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster account ad proxy delete default response has a 2xx status code
func (o *ClusterAccountAdProxyDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster account ad proxy delete default response has a 3xx status code
func (o *ClusterAccountAdProxyDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster account ad proxy delete default response has a 4xx status code
func (o *ClusterAccountAdProxyDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster account ad proxy delete default response has a 5xx status code
func (o *ClusterAccountAdProxyDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster account ad proxy delete default response a status code equal to that given
func (o *ClusterAccountAdProxyDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster account ad proxy delete default response
func (o *ClusterAccountAdProxyDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ClusterAccountAdProxyDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/authentication/cluster/ad-proxy][%d] cluster_account_ad_proxy_delete default %s", o._statusCode, payload)
}

func (o *ClusterAccountAdProxyDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/authentication/cluster/ad-proxy][%d] cluster_account_ad_proxy_delete default %s", o._statusCode, payload)
}

func (o *ClusterAccountAdProxyDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterAccountAdProxyDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
