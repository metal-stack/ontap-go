// Code generated by go-swagger; DO NOT EDIT.

package snap_lock

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

// SnaplockLogModifyReader is a Reader for the SnaplockLogModify structure.
type SnaplockLogModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockLogModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockLogModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSnaplockLogModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockLogModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockLogModifyOK creates a SnaplockLogModifyOK with default headers values
func NewSnaplockLogModifyOK() *SnaplockLogModifyOK {
	return &SnaplockLogModifyOK{}
}

/*
SnaplockLogModifyOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockLogModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snaplock log modify o k response has a 2xx status code
func (o *SnaplockLogModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock log modify o k response has a 3xx status code
func (o *SnaplockLogModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock log modify o k response has a 4xx status code
func (o *SnaplockLogModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock log modify o k response has a 5xx status code
func (o *SnaplockLogModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock log modify o k response a status code equal to that given
func (o *SnaplockLogModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snaplock log modify o k response
func (o *SnaplockLogModifyOK) Code() int {
	return 200
}

func (o *SnaplockLogModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplockLogModifyOK %s", 200, payload)
}

func (o *SnaplockLogModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplockLogModifyOK %s", 200, payload)
}

func (o *SnaplockLogModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnaplockLogModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLogModifyAccepted creates a SnaplockLogModifyAccepted with default headers values
func NewSnaplockLogModifyAccepted() *SnaplockLogModifyAccepted {
	return &SnaplockLogModifyAccepted{}
}

/*
SnaplockLogModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnaplockLogModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snaplock log modify accepted response has a 2xx status code
func (o *SnaplockLogModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock log modify accepted response has a 3xx status code
func (o *SnaplockLogModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock log modify accepted response has a 4xx status code
func (o *SnaplockLogModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock log modify accepted response has a 5xx status code
func (o *SnaplockLogModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock log modify accepted response a status code equal to that given
func (o *SnaplockLogModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the snaplock log modify accepted response
func (o *SnaplockLogModifyAccepted) Code() int {
	return 202
}

func (o *SnaplockLogModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplockLogModifyAccepted %s", 202, payload)
}

func (o *SnaplockLogModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplockLogModifyAccepted %s", 202, payload)
}

func (o *SnaplockLogModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnaplockLogModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLogModifyDefault creates a SnaplockLogModifyDefault with default headers values
func NewSnaplockLogModifyDefault(code int) *SnaplockLogModifyDefault {
	return &SnaplockLogModifyDefault{
		_statusCode: code,
	}
}

/*
	SnaplockLogModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 14090344    | If log_volume is specified, then log_archive must not be specified  |
| 14090345    | If log_archive.base_name is specified, then log_archive.archive must also be specified  |
| 14090346    | Internal Error. Wait a few minutes, then try the command again  |
*/
type SnaplockLogModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snaplock log modify default response has a 2xx status code
func (o *SnaplockLogModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock log modify default response has a 3xx status code
func (o *SnaplockLogModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock log modify default response has a 4xx status code
func (o *SnaplockLogModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock log modify default response has a 5xx status code
func (o *SnaplockLogModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock log modify default response a status code equal to that given
func (o *SnaplockLogModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snaplock log modify default response
func (o *SnaplockLogModifyDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockLogModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplock_log_modify default %s", o._statusCode, payload)
}

func (o *SnaplockLogModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplock_log_modify default %s", o._statusCode, payload)
}

func (o *SnaplockLogModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockLogModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}