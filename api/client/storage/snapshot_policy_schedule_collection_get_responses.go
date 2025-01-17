// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// SnapshotPolicyScheduleCollectionGetReader is a Reader for the SnapshotPolicyScheduleCollectionGet structure.
type SnapshotPolicyScheduleCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotPolicyScheduleCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotPolicyScheduleCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotPolicyScheduleCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotPolicyScheduleCollectionGetOK creates a SnapshotPolicyScheduleCollectionGetOK with default headers values
func NewSnapshotPolicyScheduleCollectionGetOK() *SnapshotPolicyScheduleCollectionGetOK {
	return &SnapshotPolicyScheduleCollectionGetOK{}
}

/*
SnapshotPolicyScheduleCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SnapshotPolicyScheduleCollectionGetOK struct {
	Payload *models.SnapshotPolicyScheduleResponse
}

// IsSuccess returns true when this snapshot policy schedule collection get o k response has a 2xx status code
func (o *SnapshotPolicyScheduleCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapshot policy schedule collection get o k response has a 3xx status code
func (o *SnapshotPolicyScheduleCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot policy schedule collection get o k response has a 4xx status code
func (o *SnapshotPolicyScheduleCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapshot policy schedule collection get o k response has a 5xx status code
func (o *SnapshotPolicyScheduleCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot policy schedule collection get o k response a status code equal to that given
func (o *SnapshotPolicyScheduleCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snapshot policy schedule collection get o k response
func (o *SnapshotPolicyScheduleCollectionGetOK) Code() int {
	return 200
}

func (o *SnapshotPolicyScheduleCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snapshot-policies/{snapshot_policy.uuid}/schedules][%d] snapshotPolicyScheduleCollectionGetOK %s", 200, payload)
}

func (o *SnapshotPolicyScheduleCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snapshot-policies/{snapshot_policy.uuid}/schedules][%d] snapshotPolicyScheduleCollectionGetOK %s", 200, payload)
}

func (o *SnapshotPolicyScheduleCollectionGetOK) GetPayload() *models.SnapshotPolicyScheduleResponse {
	return o.Payload
}

func (o *SnapshotPolicyScheduleCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotPolicyScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotPolicyScheduleCollectionGetDefault creates a SnapshotPolicyScheduleCollectionGetDefault with default headers values
func NewSnapshotPolicyScheduleCollectionGetDefault(code int) *SnapshotPolicyScheduleCollectionGetDefault {
	return &SnapshotPolicyScheduleCollectionGetDefault{
		_statusCode: code,
	}
}

/*
SnapshotPolicyScheduleCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SnapshotPolicyScheduleCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snapshot policy schedule collection get default response has a 2xx status code
func (o *SnapshotPolicyScheduleCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapshot policy schedule collection get default response has a 3xx status code
func (o *SnapshotPolicyScheduleCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapshot policy schedule collection get default response has a 4xx status code
func (o *SnapshotPolicyScheduleCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapshot policy schedule collection get default response has a 5xx status code
func (o *SnapshotPolicyScheduleCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapshot policy schedule collection get default response a status code equal to that given
func (o *SnapshotPolicyScheduleCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snapshot policy schedule collection get default response
func (o *SnapshotPolicyScheduleCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SnapshotPolicyScheduleCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snapshot-policies/{snapshot_policy.uuid}/schedules][%d] snapshot_policy_schedule_collection_get default %s", o._statusCode, payload)
}

func (o *SnapshotPolicyScheduleCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snapshot-policies/{snapshot_policy.uuid}/schedules][%d] snapshot_policy_schedule_collection_get default %s", o._statusCode, payload)
}

func (o *SnapshotPolicyScheduleCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotPolicyScheduleCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}