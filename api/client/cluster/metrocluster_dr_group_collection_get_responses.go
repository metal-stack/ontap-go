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

// MetroclusterDrGroupCollectionGetReader is a Reader for the MetroclusterDrGroupCollectionGet structure.
type MetroclusterDrGroupCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroclusterDrGroupCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroclusterDrGroupCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMetroclusterDrGroupCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMetroclusterDrGroupCollectionGetOK creates a MetroclusterDrGroupCollectionGetOK with default headers values
func NewMetroclusterDrGroupCollectionGetOK() *MetroclusterDrGroupCollectionGetOK {
	return &MetroclusterDrGroupCollectionGetOK{}
}

/*
MetroclusterDrGroupCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type MetroclusterDrGroupCollectionGetOK struct {
	Payload *models.MetroclusterDrGroupResponse
}

// IsSuccess returns true when this metrocluster dr group collection get o k response has a 2xx status code
func (o *MetroclusterDrGroupCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this metrocluster dr group collection get o k response has a 3xx status code
func (o *MetroclusterDrGroupCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this metrocluster dr group collection get o k response has a 4xx status code
func (o *MetroclusterDrGroupCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this metrocluster dr group collection get o k response has a 5xx status code
func (o *MetroclusterDrGroupCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this metrocluster dr group collection get o k response a status code equal to that given
func (o *MetroclusterDrGroupCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the metrocluster dr group collection get o k response
func (o *MetroclusterDrGroupCollectionGetOK) Code() int {
	return 200
}

func (o *MetroclusterDrGroupCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/dr-groups][%d] metroclusterDrGroupCollectionGetOK %s", 200, payload)
}

func (o *MetroclusterDrGroupCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/dr-groups][%d] metroclusterDrGroupCollectionGetOK %s", 200, payload)
}

func (o *MetroclusterDrGroupCollectionGetOK) GetPayload() *models.MetroclusterDrGroupResponse {
	return o.Payload
}

func (o *MetroclusterDrGroupCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetroclusterDrGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroclusterDrGroupCollectionGetDefault creates a MetroclusterDrGroupCollectionGetDefault with default headers values
func NewMetroclusterDrGroupCollectionGetDefault(code int) *MetroclusterDrGroupCollectionGetDefault {
	return &MetroclusterDrGroupCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	MetroclusterDrGroupCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2425734 | An internal error occurred. Wait a few minutes, and try the operation again. For further assistance, contact technical support. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
  - name: MCC_CONFIG_SETTINGS_UNABLE_TO_REACH_PARTNER_CLUSTER
    message: Unable to reach the partner cluster "cluster"
*/
type MetroclusterDrGroupCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this metrocluster dr group collection get default response has a 2xx status code
func (o *MetroclusterDrGroupCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this metrocluster dr group collection get default response has a 3xx status code
func (o *MetroclusterDrGroupCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this metrocluster dr group collection get default response has a 4xx status code
func (o *MetroclusterDrGroupCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this metrocluster dr group collection get default response has a 5xx status code
func (o *MetroclusterDrGroupCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this metrocluster dr group collection get default response a status code equal to that given
func (o *MetroclusterDrGroupCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the metrocluster dr group collection get default response
func (o *MetroclusterDrGroupCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *MetroclusterDrGroupCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/dr-groups][%d] metrocluster_dr_group_collection_get default %s", o._statusCode, payload)
}

func (o *MetroclusterDrGroupCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/metrocluster/dr-groups][%d] metrocluster_dr_group_collection_get default %s", o._statusCode, payload)
}

func (o *MetroclusterDrGroupCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MetroclusterDrGroupCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}