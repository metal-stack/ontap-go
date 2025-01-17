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

// SnapshotModifyReader is a Reader for the SnapshotModify structure.
type SnapshotModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSnapshotModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotModifyOK creates a SnapshotModifyOK with default headers values
func NewSnapshotModifyOK() *SnapshotModifyOK {
	return &SnapshotModifyOK{}
}

/*
SnapshotModifyOK describes a response with status code 200, with default header values.

OK
*/
type SnapshotModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snapshot modify o k response has a 2xx status code
func (o *SnapshotModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapshot modify o k response has a 3xx status code
func (o *SnapshotModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot modify o k response has a 4xx status code
func (o *SnapshotModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapshot modify o k response has a 5xx status code
func (o *SnapshotModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot modify o k response a status code equal to that given
func (o *SnapshotModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snapshot modify o k response
func (o *SnapshotModifyOK) Code() int {
	return 200
}

func (o *SnapshotModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/volumes/{volume.uuid}/snapshots/{uuid}][%d] snapshotModifyOK %s", 200, payload)
}

func (o *SnapshotModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/volumes/{volume.uuid}/snapshots/{uuid}][%d] snapshotModifyOK %s", 200, payload)
}

func (o *SnapshotModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnapshotModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotModifyAccepted creates a SnapshotModifyAccepted with default headers values
func NewSnapshotModifyAccepted() *SnapshotModifyAccepted {
	return &SnapshotModifyAccepted{}
}

/*
SnapshotModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnapshotModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snapshot modify accepted response has a 2xx status code
func (o *SnapshotModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapshot modify accepted response has a 3xx status code
func (o *SnapshotModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot modify accepted response has a 4xx status code
func (o *SnapshotModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapshot modify accepted response has a 5xx status code
func (o *SnapshotModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot modify accepted response a status code equal to that given
func (o *SnapshotModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the snapshot modify accepted response
func (o *SnapshotModifyAccepted) Code() int {
	return 202
}

func (o *SnapshotModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/volumes/{volume.uuid}/snapshots/{uuid}][%d] snapshotModifyAccepted %s", 202, payload)
}

func (o *SnapshotModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/volumes/{volume.uuid}/snapshots/{uuid}][%d] snapshotModifyAccepted %s", 202, payload)
}

func (o *SnapshotModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnapshotModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotModifyDefault creates a SnapshotModifyDefault with default headers values
func NewSnapshotModifyDefault(code int) *SnapshotModifyDefault {
	return &SnapshotModifyDefault{
		_statusCode: code,
	}
}

/*
	SnapshotModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Code

| Error Code | Description |
| ---------- | ----------- |
| 524508     | The Snapshot copy was not renamed because the name entered is not valid. |
| 542797     | The specified file or Snapshot copy does not exist. |
| 1638455    | Failed to set comment for Snapshot copy. |
| 1638476    | You cannot rename a Snapshot copy created for use as a reference Snapshot copy by other jobs. |
| 1638477    | User-created Snapshot copy names cannot begin with the specified prefix. |
| 1638518    | The specified Snapshot copy name is invalid. |
| 1638522    | Snapshot copies can only be renamed on read/write (RW) volumes. |
| 1638523    | Failed to set the specified SnapMirror label for the Snapshot copy. |
| 1638524    | Adding SnapMirror labels is not allowed in a mixed version cluster. |
| 1638539    | Cannot determine the status of the Snapshot copy rename operation for the specified volume. |
| 1638554    | Failed to set expiry time for the Snapshot copy. |
| 1638600    | The Snapshot copy does not exist. |
*/
type SnapshotModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snapshot modify default response has a 2xx status code
func (o *SnapshotModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapshot modify default response has a 3xx status code
func (o *SnapshotModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapshot modify default response has a 4xx status code
func (o *SnapshotModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapshot modify default response has a 5xx status code
func (o *SnapshotModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapshot modify default response a status code equal to that given
func (o *SnapshotModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snapshot modify default response
func (o *SnapshotModifyDefault) Code() int {
	return o._statusCode
}

func (o *SnapshotModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/volumes/{volume.uuid}/snapshots/{uuid}][%d] snapshot_modify default %s", o._statusCode, payload)
}

func (o *SnapshotModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/volumes/{volume.uuid}/snapshots/{uuid}][%d] snapshot_modify default %s", o._statusCode, payload)
}

func (o *SnapshotModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}