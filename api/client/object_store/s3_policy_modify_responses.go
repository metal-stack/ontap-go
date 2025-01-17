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

// S3PolicyModifyReader is a Reader for the S3PolicyModify structure.
type S3PolicyModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3PolicyModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3PolicyModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3PolicyModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3PolicyModifyOK creates a S3PolicyModifyOK with default headers values
func NewS3PolicyModifyOK() *S3PolicyModifyOK {
	return &S3PolicyModifyOK{}
}

/*
S3PolicyModifyOK describes a response with status code 200, with default header values.

OK
*/
type S3PolicyModifyOK struct {
}

// IsSuccess returns true when this s3 policy modify o k response has a 2xx status code
func (o *S3PolicyModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 policy modify o k response has a 3xx status code
func (o *S3PolicyModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 policy modify o k response has a 4xx status code
func (o *S3PolicyModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 policy modify o k response has a 5xx status code
func (o *S3PolicyModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 policy modify o k response a status code equal to that given
func (o *S3PolicyModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s3 policy modify o k response
func (o *S3PolicyModifyOK) Code() int {
	return 200
}

func (o *S3PolicyModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/policies/{name}][%d] s3PolicyModifyOK", 200)
}

func (o *S3PolicyModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/policies/{name}][%d] s3PolicyModifyOK", 200)
}

func (o *S3PolicyModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewS3PolicyModifyDefault creates a S3PolicyModifyDefault with default headers values
func NewS3PolicyModifyDefault(code int) *S3PolicyModifyDefault {
	return &S3PolicyModifyDefault{
		_statusCode: code,
	}
}

/*
	S3PolicyModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 92405906   | The specified action name is invalid.
| 92405963   | Failed to create s3 policy statements \\\"{policy name}\\\". Reason: "{reason of failure}". Resolve all issues and retry the operation.
| 92405953   | Object store server read-only policies do not support create, modify, delete, add-statement, delete-statement and modify-statement operations.
| 92406075   | Failed to modify policy statement for policy \\\"{policy name}\\\". Reason: "{reason of failure}". Valid ways to specify a resource are \"*\", \"<bucket-name>\", \"<bucket-name>/.../...\".\".
*/
type S3PolicyModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 policy modify default response has a 2xx status code
func (o *S3PolicyModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 policy modify default response has a 3xx status code
func (o *S3PolicyModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 policy modify default response has a 4xx status code
func (o *S3PolicyModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 policy modify default response has a 5xx status code
func (o *S3PolicyModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 policy modify default response a status code equal to that given
func (o *S3PolicyModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 policy modify default response
func (o *S3PolicyModifyDefault) Code() int {
	return o._statusCode
}

func (o *S3PolicyModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/policies/{name}][%d] s3_policy_modify default %s", o._statusCode, payload)
}

func (o *S3PolicyModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/policies/{name}][%d] s3_policy_modify default %s", o._statusCode, payload)
}

func (o *S3PolicyModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3PolicyModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}