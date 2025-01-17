// Code generated by go-swagger; DO NOT EDIT.

package s_vm

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

// TopMetricsSvmDirectoryCollectionGetReader is a Reader for the TopMetricsSvmDirectoryCollectionGet structure.
type TopMetricsSvmDirectoryCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TopMetricsSvmDirectoryCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTopMetricsSvmDirectoryCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTopMetricsSvmDirectoryCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTopMetricsSvmDirectoryCollectionGetOK creates a TopMetricsSvmDirectoryCollectionGetOK with default headers values
func NewTopMetricsSvmDirectoryCollectionGetOK() *TopMetricsSvmDirectoryCollectionGetOK {
	return &TopMetricsSvmDirectoryCollectionGetOK{}
}

/*
TopMetricsSvmDirectoryCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type TopMetricsSvmDirectoryCollectionGetOK struct {
	Payload *models.TopMetricsSvmDirectoryResponse
}

// IsSuccess returns true when this top metrics svm directory collection get o k response has a 2xx status code
func (o *TopMetricsSvmDirectoryCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this top metrics svm directory collection get o k response has a 3xx status code
func (o *TopMetricsSvmDirectoryCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this top metrics svm directory collection get o k response has a 4xx status code
func (o *TopMetricsSvmDirectoryCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this top metrics svm directory collection get o k response has a 5xx status code
func (o *TopMetricsSvmDirectoryCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this top metrics svm directory collection get o k response a status code equal to that given
func (o *TopMetricsSvmDirectoryCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the top metrics svm directory collection get o k response
func (o *TopMetricsSvmDirectoryCollectionGetOK) Code() int {
	return 200
}

func (o *TopMetricsSvmDirectoryCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/svms/{svm.uuid}/top-metrics/directories][%d] topMetricsSvmDirectoryCollectionGetOK %s", 200, payload)
}

func (o *TopMetricsSvmDirectoryCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/svms/{svm.uuid}/top-metrics/directories][%d] topMetricsSvmDirectoryCollectionGetOK %s", 200, payload)
}

func (o *TopMetricsSvmDirectoryCollectionGetOK) GetPayload() *models.TopMetricsSvmDirectoryResponse {
	return o.Payload
}

func (o *TopMetricsSvmDirectoryCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TopMetricsSvmDirectoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTopMetricsSvmDirectoryCollectionGetDefault creates a TopMetricsSvmDirectoryCollectionGetDefault with default headers values
func NewTopMetricsSvmDirectoryCollectionGetDefault(code int) *TopMetricsSvmDirectoryCollectionGetDefault {
	return &TopMetricsSvmDirectoryCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	TopMetricsSvmDirectoryCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 124519405 | The activity tracking report for SVM svm.name returned zero records. Check whether the activity tracking enabled volumes belonging to the SVM have read/write traffic. Refer to the REST API documentation for more information on why there might be no records. |
| 124519406 | Failed to get the activity tracking report for SVM svm.name. Reason:<Reason for failure>. |
| 124519407 | SVM wildcard queries are not supported for activity tracking reports. |
| 124519408 | Activity tracking is not supported on SVM svm.name, because it is configured as a destination for SVM DR. |
| 124519409 | Activity tracking is not supported on SVM svm.name, because it is configured as a destination of a MetroCluster SVM relationship and the SVM admin state is stopped. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type TopMetricsSvmDirectoryCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this top metrics svm directory collection get default response has a 2xx status code
func (o *TopMetricsSvmDirectoryCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this top metrics svm directory collection get default response has a 3xx status code
func (o *TopMetricsSvmDirectoryCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this top metrics svm directory collection get default response has a 4xx status code
func (o *TopMetricsSvmDirectoryCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this top metrics svm directory collection get default response has a 5xx status code
func (o *TopMetricsSvmDirectoryCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this top metrics svm directory collection get default response a status code equal to that given
func (o *TopMetricsSvmDirectoryCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the top metrics svm directory collection get default response
func (o *TopMetricsSvmDirectoryCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *TopMetricsSvmDirectoryCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/svms/{svm.uuid}/top-metrics/directories][%d] top_metrics_svm_directory_collection_get default %s", o._statusCode, payload)
}

func (o *TopMetricsSvmDirectoryCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/svms/{svm.uuid}/top-metrics/directories][%d] top_metrics_svm_directory_collection_get default %s", o._statusCode, payload)
}

func (o *TopMetricsSvmDirectoryCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TopMetricsSvmDirectoryCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
