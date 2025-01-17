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

// SnapshotCollectionGetReader is a Reader for the SnapshotCollectionGet structure.
type SnapshotCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotCollectionGetOK creates a SnapshotCollectionGetOK with default headers values
func NewSnapshotCollectionGetOK() *SnapshotCollectionGetOK {
	return &SnapshotCollectionGetOK{}
}

/*
SnapshotCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SnapshotCollectionGetOK struct {
	Payload *models.SnapshotResponse
}

// IsSuccess returns true when this snapshot collection get o k response has a 2xx status code
func (o *SnapshotCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapshot collection get o k response has a 3xx status code
func (o *SnapshotCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot collection get o k response has a 4xx status code
func (o *SnapshotCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapshot collection get o k response has a 5xx status code
func (o *SnapshotCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot collection get o k response a status code equal to that given
func (o *SnapshotCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snapshot collection get o k response
func (o *SnapshotCollectionGetOK) Code() int {
	return 200
}

func (o *SnapshotCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes/{volume.uuid}/snapshots][%d] snapshotCollectionGetOK %s", 200, payload)
}

func (o *SnapshotCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes/{volume.uuid}/snapshots][%d] snapshotCollectionGetOK %s", 200, payload)
}

func (o *SnapshotCollectionGetOK) GetPayload() *models.SnapshotResponse {
	return o.Payload
}

func (o *SnapshotCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotCollectionGetDefault creates a SnapshotCollectionGetDefault with default headers values
func NewSnapshotCollectionGetDefault(code int) *SnapshotCollectionGetDefault {
	return &SnapshotCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	SnapshotCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Code

| Error Code | Description |
| ---------- | ----------- |
| 918235     | The specified volume is invalid. |
| 9437613    | The operation is not supported on FlexCache Volumes. |
*/
type SnapshotCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snapshot collection get default response has a 2xx status code
func (o *SnapshotCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapshot collection get default response has a 3xx status code
func (o *SnapshotCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapshot collection get default response has a 4xx status code
func (o *SnapshotCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapshot collection get default response has a 5xx status code
func (o *SnapshotCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapshot collection get default response a status code equal to that given
func (o *SnapshotCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snapshot collection get default response
func (o *SnapshotCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SnapshotCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes/{volume.uuid}/snapshots][%d] snapshot_collection_get default %s", o._statusCode, payload)
}

func (o *SnapshotCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes/{volume.uuid}/snapshots][%d] snapshot_collection_get default %s", o._statusCode, payload)
}

func (o *SnapshotCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
