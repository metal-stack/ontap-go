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

// S3BucketLifecycleRuleModifyReader is a Reader for the S3BucketLifecycleRuleModify structure.
type S3BucketLifecycleRuleModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3BucketLifecycleRuleModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3BucketLifecycleRuleModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewS3BucketLifecycleRuleModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3BucketLifecycleRuleModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3BucketLifecycleRuleModifyOK creates a S3BucketLifecycleRuleModifyOK with default headers values
func NewS3BucketLifecycleRuleModifyOK() *S3BucketLifecycleRuleModifyOK {
	return &S3BucketLifecycleRuleModifyOK{}
}

/*
S3BucketLifecycleRuleModifyOK describes a response with status code 200, with default header values.

OK
*/
type S3BucketLifecycleRuleModifyOK struct {
	Payload *models.S3BucketLifecycleRuleResponse
}

// IsSuccess returns true when this s3 bucket lifecycle rule modify o k response has a 2xx status code
func (o *S3BucketLifecycleRuleModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket lifecycle rule modify o k response has a 3xx status code
func (o *S3BucketLifecycleRuleModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket lifecycle rule modify o k response has a 4xx status code
func (o *S3BucketLifecycleRuleModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket lifecycle rule modify o k response has a 5xx status code
func (o *S3BucketLifecycleRuleModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket lifecycle rule modify o k response a status code equal to that given
func (o *S3BucketLifecycleRuleModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the s3 bucket lifecycle rule modify o k response
func (o *S3BucketLifecycleRuleModifyOK) Code() int {
	return 200
}

func (o *S3BucketLifecycleRuleModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules/{name}][%d] s3BucketLifecycleRuleModifyOK %s", 200, payload)
}

func (o *S3BucketLifecycleRuleModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules/{name}][%d] s3BucketLifecycleRuleModifyOK %s", 200, payload)
}

func (o *S3BucketLifecycleRuleModifyOK) GetPayload() *models.S3BucketLifecycleRuleResponse {
	return o.Payload
}

func (o *S3BucketLifecycleRuleModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3BucketLifecycleRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketLifecycleRuleModifyAccepted creates a S3BucketLifecycleRuleModifyAccepted with default headers values
func NewS3BucketLifecycleRuleModifyAccepted() *S3BucketLifecycleRuleModifyAccepted {
	return &S3BucketLifecycleRuleModifyAccepted{}
}

/*
S3BucketLifecycleRuleModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type S3BucketLifecycleRuleModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this s3 bucket lifecycle rule modify accepted response has a 2xx status code
func (o *S3BucketLifecycleRuleModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket lifecycle rule modify accepted response has a 3xx status code
func (o *S3BucketLifecycleRuleModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket lifecycle rule modify accepted response has a 4xx status code
func (o *S3BucketLifecycleRuleModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket lifecycle rule modify accepted response has a 5xx status code
func (o *S3BucketLifecycleRuleModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket lifecycle rule modify accepted response a status code equal to that given
func (o *S3BucketLifecycleRuleModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the s3 bucket lifecycle rule modify accepted response
func (o *S3BucketLifecycleRuleModifyAccepted) Code() int {
	return 202
}

func (o *S3BucketLifecycleRuleModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules/{name}][%d] s3BucketLifecycleRuleModifyAccepted %s", 202, payload)
}

func (o *S3BucketLifecycleRuleModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules/{name}][%d] s3BucketLifecycleRuleModifyAccepted %s", 202, payload)
}

func (o *S3BucketLifecycleRuleModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *S3BucketLifecycleRuleModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketLifecycleRuleModifyDefault creates a S3BucketLifecycleRuleModifyDefault with default headers values
func NewS3BucketLifecycleRuleModifyDefault(code int) *S3BucketLifecycleRuleModifyDefault {
	return &S3BucketLifecycleRuleModifyDefault{
		_statusCode: code,
	}
}

/*
	S3BucketLifecycleRuleModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error code | Message |
| ---------- | ------- |
| 92405861   | "The specified SVM UUID or bucket UUID does not exist.";
| 92406114   | "The Expiration action requires specifying either object_expiry_date, object_age_days or expired_object_delete_marker.";
| 92406115   | "The NonCurrentVersionExpiration action requires either new_non_current_versions or non_current_days.";
| 92406116   | "The AbortIncompleteMultipartUpload action requires after_initiation_days.";
| 92406117   | "The \"Expiration\" action cannot have both an expiry date and an age.";
| 92406118   | "Using Lifecycle Management rules requires an effective cluster version of 9.13.1 .";
| 92406122   | "\"expired_object_delete_marker\" cannot be specified with \"object_age_days\" or \"object_expiry_date\".";
| 92406131   | "Lifecycle Management rule \"testcheck2\" for bucket \"testbuck1\" in SVM \"vs0\" cannot be created because \"non_current_days\" must be specified along with \"new_non_current_versions\".";
| 92406132   | "Lifecycle Management rule \"<rule>\" for bucket \"<bucket>\" in SVM \"<SVM>\" requires \"object_age_days\" to be greater than zero.";
| 92406132   | "Lifecycle Management rule \"<rule>\" for bucket \"<bucket>\" in SVM \"<SVM>\" requires \"new_non_current_versions\" to be greater than zero.";
| 92406132   | "Lifecycle Management rule \"<rule>\" for bucket \"<bucket>\" in SVM \"<SVM>\" requires \"after_initiation_days\" to be greater than zero.";
| 92406133   | "Lifecycle Management rule for bucket in Vserver is invalid. The object_expiry_date must be later than January 1, 1970.";
| 92406135   | "MetroCluster is configured on cluster. Object Expiration is not supported in a MetroCluster configuration.";
| 92406136   | "Lifecycle Management rule \"rule1\" for bucket \"buck1\" in SVM \"vs0\" is invalid. The \"object_expiry_date\" must be at midnight GMT.";
| 92406139   | "Lifecycle Management rule for bucket in Vserver with action is a stale entry. Contact technical support for assistance.";
| 92406142   | "Lifecycle Management rule \"rule1\" for bucket \"buck1\" in SVM \"vs0\" with action \"Expiration\" cannot have \"expired_object_delete_marker\" disabled. To disable \"expired_object_delete_marker\", run the \"vserver object-store-server bucket lifecycle-management-rule delete\" command.";
| 92406146   | "Lifecycle Management rule \"rule1\" for bucket \"buck1\" in SVM \"vs0\" already exists. Delete the rule and then try the operation again.";
| 92406148   | "Lifecycle Management rule \"rule1\" for bucket \"buck1\" in SVM \"vs0\" cannot have \"new_non_current_versions\" more than 100.";
| 92406149   | "Lifecycle Management rule \"rule1\" for bucket \"buck1\" in SVM \"vs0\" requires an action to be specified. Retry the operation after adding an action.";
*/
type S3BucketLifecycleRuleModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 bucket lifecycle rule modify default response has a 2xx status code
func (o *S3BucketLifecycleRuleModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 bucket lifecycle rule modify default response has a 3xx status code
func (o *S3BucketLifecycleRuleModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 bucket lifecycle rule modify default response has a 4xx status code
func (o *S3BucketLifecycleRuleModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 bucket lifecycle rule modify default response has a 5xx status code
func (o *S3BucketLifecycleRuleModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 bucket lifecycle rule modify default response a status code equal to that given
func (o *S3BucketLifecycleRuleModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 bucket lifecycle rule modify default response
func (o *S3BucketLifecycleRuleModifyDefault) Code() int {
	return o._statusCode
}

func (o *S3BucketLifecycleRuleModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules/{name}][%d] s3_bucket_lifecycle_rule_modify default %s", o._statusCode, payload)
}

func (o *S3BucketLifecycleRuleModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/s3/services/{svm.uuid}/buckets/{s3_bucket.uuid}/rules/{name}][%d] s3_bucket_lifecycle_rule_modify default %s", o._statusCode, payload)
}

func (o *S3BucketLifecycleRuleModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3BucketLifecycleRuleModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
