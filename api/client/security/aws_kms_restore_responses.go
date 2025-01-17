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

// AwsKmsRestoreReader is a Reader for the AwsKmsRestore structure.
type AwsKmsRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AwsKmsRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAwsKmsRestoreCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewAwsKmsRestoreAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAwsKmsRestoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAwsKmsRestoreCreated creates a AwsKmsRestoreCreated with default headers values
func NewAwsKmsRestoreCreated() *AwsKmsRestoreCreated {
	return &AwsKmsRestoreCreated{}
}

/*
AwsKmsRestoreCreated describes a response with status code 201, with default header values.

Created
*/
type AwsKmsRestoreCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this aws kms restore created response has a 2xx status code
func (o *AwsKmsRestoreCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aws kms restore created response has a 3xx status code
func (o *AwsKmsRestoreCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws kms restore created response has a 4xx status code
func (o *AwsKmsRestoreCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this aws kms restore created response has a 5xx status code
func (o *AwsKmsRestoreCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this aws kms restore created response a status code equal to that given
func (o *AwsKmsRestoreCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the aws kms restore created response
func (o *AwsKmsRestoreCreated) Code() int {
	return 201
}

func (o *AwsKmsRestoreCreated) Error() string {
	return fmt.Sprintf("[POST /security/aws-kms/{aws_kms.uuid}/restore][%d] awsKmsRestoreCreated", 201)
}

func (o *AwsKmsRestoreCreated) String() string {
	return fmt.Sprintf("[POST /security/aws-kms/{aws_kms.uuid}/restore][%d] awsKmsRestoreCreated", 201)
}

func (o *AwsKmsRestoreCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewAwsKmsRestoreAccepted creates a AwsKmsRestoreAccepted with default headers values
func NewAwsKmsRestoreAccepted() *AwsKmsRestoreAccepted {
	return &AwsKmsRestoreAccepted{}
}

/*
AwsKmsRestoreAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AwsKmsRestoreAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this aws kms restore accepted response has a 2xx status code
func (o *AwsKmsRestoreAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aws kms restore accepted response has a 3xx status code
func (o *AwsKmsRestoreAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws kms restore accepted response has a 4xx status code
func (o *AwsKmsRestoreAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this aws kms restore accepted response has a 5xx status code
func (o *AwsKmsRestoreAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this aws kms restore accepted response a status code equal to that given
func (o *AwsKmsRestoreAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the aws kms restore accepted response
func (o *AwsKmsRestoreAccepted) Code() int {
	return 202
}

func (o *AwsKmsRestoreAccepted) Error() string {
	return fmt.Sprintf("[POST /security/aws-kms/{aws_kms.uuid}/restore][%d] awsKmsRestoreAccepted", 202)
}

func (o *AwsKmsRestoreAccepted) String() string {
	return fmt.Sprintf("[POST /security/aws-kms/{aws_kms.uuid}/restore][%d] awsKmsRestoreAccepted", 202)
}

func (o *AwsKmsRestoreAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewAwsKmsRestoreDefault creates a AwsKmsRestoreDefault with default headers values
func NewAwsKmsRestoreDefault(code int) *AwsKmsRestoreDefault {
	return &AwsKmsRestoreDefault{
		_statusCode: code,
	}
}

/*
	AwsKmsRestoreDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536082 | Unable to restore all keys. |
| 65537544 | Missing wrapped top-level internal key protection key (KEK) from internal database. |
| 65537926 | The Amazon Web Service Key Management Service is not configured for the given SVM. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AwsKmsRestoreDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this aws kms restore default response has a 2xx status code
func (o *AwsKmsRestoreDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this aws kms restore default response has a 3xx status code
func (o *AwsKmsRestoreDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this aws kms restore default response has a 4xx status code
func (o *AwsKmsRestoreDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this aws kms restore default response has a 5xx status code
func (o *AwsKmsRestoreDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this aws kms restore default response a status code equal to that given
func (o *AwsKmsRestoreDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the aws kms restore default response
func (o *AwsKmsRestoreDefault) Code() int {
	return o._statusCode
}

func (o *AwsKmsRestoreDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/aws-kms/{aws_kms.uuid}/restore][%d] aws_kms_restore default %s", o._statusCode, payload)
}

func (o *AwsKmsRestoreDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/aws-kms/{aws_kms.uuid}/restore][%d] aws_kms_restore default %s", o._statusCode, payload)
}

func (o *AwsKmsRestoreDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AwsKmsRestoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}