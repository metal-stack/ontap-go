// Code generated by go-swagger; DO NOT EDIT.

package snap_mirror

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

// SnapmirrorPolicyCreateReader is a Reader for the SnapmirrorPolicyCreate structure.
type SnapmirrorPolicyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapmirrorPolicyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSnapmirrorPolicyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSnapmirrorPolicyCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapmirrorPolicyCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapmirrorPolicyCreateCreated creates a SnapmirrorPolicyCreateCreated with default headers values
func NewSnapmirrorPolicyCreateCreated() *SnapmirrorPolicyCreateCreated {
	return &SnapmirrorPolicyCreateCreated{}
}

/*
SnapmirrorPolicyCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SnapmirrorPolicyCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snapmirror policy create created response has a 2xx status code
func (o *SnapmirrorPolicyCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapmirror policy create created response has a 3xx status code
func (o *SnapmirrorPolicyCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapmirror policy create created response has a 4xx status code
func (o *SnapmirrorPolicyCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapmirror policy create created response has a 5xx status code
func (o *SnapmirrorPolicyCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this snapmirror policy create created response a status code equal to that given
func (o *SnapmirrorPolicyCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the snapmirror policy create created response
func (o *SnapmirrorPolicyCreateCreated) Code() int {
	return 201
}

func (o *SnapmirrorPolicyCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /snapmirror/policies][%d] snapmirrorPolicyCreateCreated %s", 201, payload)
}

func (o *SnapmirrorPolicyCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /snapmirror/policies][%d] snapmirrorPolicyCreateCreated %s", 201, payload)
}

func (o *SnapmirrorPolicyCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnapmirrorPolicyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapmirrorPolicyCreateAccepted creates a SnapmirrorPolicyCreateAccepted with default headers values
func NewSnapmirrorPolicyCreateAccepted() *SnapmirrorPolicyCreateAccepted {
	return &SnapmirrorPolicyCreateAccepted{}
}

/*
SnapmirrorPolicyCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnapmirrorPolicyCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snapmirror policy create accepted response has a 2xx status code
func (o *SnapmirrorPolicyCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapmirror policy create accepted response has a 3xx status code
func (o *SnapmirrorPolicyCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapmirror policy create accepted response has a 4xx status code
func (o *SnapmirrorPolicyCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapmirror policy create accepted response has a 5xx status code
func (o *SnapmirrorPolicyCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this snapmirror policy create accepted response a status code equal to that given
func (o *SnapmirrorPolicyCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the snapmirror policy create accepted response
func (o *SnapmirrorPolicyCreateAccepted) Code() int {
	return 202
}

func (o *SnapmirrorPolicyCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /snapmirror/policies][%d] snapmirrorPolicyCreateAccepted %s", 202, payload)
}

func (o *SnapmirrorPolicyCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /snapmirror/policies][%d] snapmirrorPolicyCreateAccepted %s", 202, payload)
}

func (o *SnapmirrorPolicyCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnapmirrorPolicyCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapmirrorPolicyCreateDefault creates a SnapmirrorPolicyCreateDefault with default headers values
func NewSnapmirrorPolicyCreateDefault(code int) *SnapmirrorPolicyCreateDefault {
	return &SnapmirrorPolicyCreateDefault{
		_statusCode: code,
	}
}

/*
	SnapmirrorPolicyCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 6619714     | Schedule specified is an interval schedule. SnapMirror does not support interval schedules. |
| 13304126    | Enter a value for "count" in the range indicated in the error message. |
| 13303887    | Failed to create SnapMirror policy. Reason: Maximum number of allowed retention rules reached |
| 13304083    | The specified property is not supported because all nodes in the cluster are not capable of supporting this property. |
| 13304084    | Properties specified are mutually exclusive. Provide only one property. |
| 13304085    | The specified property does not support the specified value. |
| 13304092    | Input value of the retention period property is invalid. For relationships with FlexVol volume or FlexGroup volume destinations, the duration must be in ISO 6801 format or can be infinite. For relationships with object store destinations, only duration values with Y, M or D and supported and must be in the specified range. |

| 6621045     | The property rentention.warn is not supported for a policy when the retention.preserve property is false.
| 13304129    | The property retention.warn value must be less than the property retention.count value for a rule in a policy.
| 13304130    | The total retention.count value for all rules in a policy cannot exceed the value indicated in the error message.
| 6621060     | Failed to create specified policy. Reason: A policy with this name already exists. |
| 13304118    | The specified properties cannot be both set at the same time:  retention.creation_schedule and retention.period |
| 6621972     | Schedule specified is not supported by SnapMirror Synchronous. The allowed time interval for the creation of common Snapshot copies can range between 30 minutes and 24 hours. |
| 13304011    | Failed to create or modify the specified SnapMirror policy. Reason: The property retention.count cannot have a value greater than 1. |
| 6621091     | The specified SnapMirror policy cannot have a rule with a preserve value of true. |
| 6621088     | The specified SnapMirror policy must have at least one rule without a schedule. |
*/
type SnapmirrorPolicyCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snapmirror policy create default response has a 2xx status code
func (o *SnapmirrorPolicyCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapmirror policy create default response has a 3xx status code
func (o *SnapmirrorPolicyCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapmirror policy create default response has a 4xx status code
func (o *SnapmirrorPolicyCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapmirror policy create default response has a 5xx status code
func (o *SnapmirrorPolicyCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapmirror policy create default response a status code equal to that given
func (o *SnapmirrorPolicyCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snapmirror policy create default response
func (o *SnapmirrorPolicyCreateDefault) Code() int {
	return o._statusCode
}

func (o *SnapmirrorPolicyCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /snapmirror/policies][%d] snapmirror_policy_create default %s", o._statusCode, payload)
}

func (o *SnapmirrorPolicyCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /snapmirror/policies][%d] snapmirror_policy_create default %s", o._statusCode, payload)
}

func (o *SnapmirrorPolicyCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapmirrorPolicyCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}