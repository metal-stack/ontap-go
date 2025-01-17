// Code generated by go-swagger; DO NOT EDIT.

package object_store

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

// PerformanceS3MetricCollectionGetReader is a Reader for the PerformanceS3MetricCollectionGet structure.
type PerformanceS3MetricCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformanceS3MetricCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformanceS3MetricCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformanceS3MetricCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformanceS3MetricCollectionGetOK creates a PerformanceS3MetricCollectionGetOK with default headers values
func NewPerformanceS3MetricCollectionGetOK() *PerformanceS3MetricCollectionGetOK {
	return &PerformanceS3MetricCollectionGetOK{}
}

/*
PerformanceS3MetricCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type PerformanceS3MetricCollectionGetOK struct {
	Payload *models.PerformanceS3MetricResponse
}

// IsSuccess returns true when this performance s3 metric collection get o k response has a 2xx status code
func (o *PerformanceS3MetricCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance s3 metric collection get o k response has a 3xx status code
func (o *PerformanceS3MetricCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance s3 metric collection get o k response has a 4xx status code
func (o *PerformanceS3MetricCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance s3 metric collection get o k response has a 5xx status code
func (o *PerformanceS3MetricCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance s3 metric collection get o k response a status code equal to that given
func (o *PerformanceS3MetricCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performance s3 metric collection get o k response
func (o *PerformanceS3MetricCollectionGetOK) Code() int {
	return 200
}

func (o *PerformanceS3MetricCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/metrics][%d] performanceS3MetricCollectionGetOK %s", 200, payload)
}

func (o *PerformanceS3MetricCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/metrics][%d] performanceS3MetricCollectionGetOK %s", 200, payload)
}

func (o *PerformanceS3MetricCollectionGetOK) GetPayload() *models.PerformanceS3MetricResponse {
	return o.Payload
}

func (o *PerformanceS3MetricCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceS3MetricResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformanceS3MetricCollectionGetDefault creates a PerformanceS3MetricCollectionGetDefault with default headers values
func NewPerformanceS3MetricCollectionGetDefault(code int) *PerformanceS3MetricCollectionGetDefault {
	return &PerformanceS3MetricCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	PerformanceS3MetricCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8586226    | UUID not found. |
| 8585947    | There are no entries matching your query. |
*/
type PerformanceS3MetricCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this performance s3 metric collection get default response has a 2xx status code
func (o *PerformanceS3MetricCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performance s3 metric collection get default response has a 3xx status code
func (o *PerformanceS3MetricCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performance s3 metric collection get default response has a 4xx status code
func (o *PerformanceS3MetricCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performance s3 metric collection get default response has a 5xx status code
func (o *PerformanceS3MetricCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performance s3 metric collection get default response a status code equal to that given
func (o *PerformanceS3MetricCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the performance s3 metric collection get default response
func (o *PerformanceS3MetricCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *PerformanceS3MetricCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/metrics][%d] performance_s3_metric_collection_get default %s", o._statusCode, payload)
}

func (o *PerformanceS3MetricCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/metrics][%d] performance_s3_metric_collection_get default %s", o._statusCode, payload)
}

func (o *PerformanceS3MetricCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PerformanceS3MetricCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}