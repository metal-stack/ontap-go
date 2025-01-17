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

// SnaplockRetentionPolicyDeleteReader is a Reader for the SnaplockRetentionPolicyDelete structure.
type SnaplockRetentionPolicyDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockRetentionPolicyDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockRetentionPolicyDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockRetentionPolicyDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockRetentionPolicyDeleteOK creates a SnaplockRetentionPolicyDeleteOK with default headers values
func NewSnaplockRetentionPolicyDeleteOK() *SnaplockRetentionPolicyDeleteOK {
	return &SnaplockRetentionPolicyDeleteOK{}
}

/*
SnaplockRetentionPolicyDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockRetentionPolicyDeleteOK struct {
}

// IsSuccess returns true when this snaplock retention policy delete o k response has a 2xx status code
func (o *SnaplockRetentionPolicyDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock retention policy delete o k response has a 3xx status code
func (o *SnaplockRetentionPolicyDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock retention policy delete o k response has a 4xx status code
func (o *SnaplockRetentionPolicyDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock retention policy delete o k response has a 5xx status code
func (o *SnaplockRetentionPolicyDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock retention policy delete o k response a status code equal to that given
func (o *SnaplockRetentionPolicyDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snaplock retention policy delete o k response
func (o *SnaplockRetentionPolicyDeleteOK) Code() int {
	return 200
}

func (o *SnaplockRetentionPolicyDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /storage/snaplock/event-retention/policies/{policy.name}][%d] snaplockRetentionPolicyDeleteOK", 200)
}

func (o *SnaplockRetentionPolicyDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /storage/snaplock/event-retention/policies/{policy.name}][%d] snaplockRetentionPolicyDeleteOK", 200)
}

func (o *SnaplockRetentionPolicyDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSnaplockRetentionPolicyDeleteDefault creates a SnaplockRetentionPolicyDeleteDefault with default headers values
func NewSnaplockRetentionPolicyDeleteDefault(code int) *SnaplockRetentionPolicyDeleteDefault {
	return &SnaplockRetentionPolicyDeleteDefault{
		_statusCode: code,
	}
}

/*
	SnaplockRetentionPolicyDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 13763280    | Only a user with security login role \"vsadmin-snaplock\" is allowed to perform this operation.  |
*/
type SnaplockRetentionPolicyDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snaplock retention policy delete default response has a 2xx status code
func (o *SnaplockRetentionPolicyDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock retention policy delete default response has a 3xx status code
func (o *SnaplockRetentionPolicyDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock retention policy delete default response has a 4xx status code
func (o *SnaplockRetentionPolicyDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock retention policy delete default response has a 5xx status code
func (o *SnaplockRetentionPolicyDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock retention policy delete default response a status code equal to that given
func (o *SnaplockRetentionPolicyDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snaplock retention policy delete default response
func (o *SnaplockRetentionPolicyDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockRetentionPolicyDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/snaplock/event-retention/policies/{policy.name}][%d] snaplock_retention_policy_delete default %s", o._statusCode, payload)
}

func (o *SnaplockRetentionPolicyDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/snaplock/event-retention/policies/{policy.name}][%d] snaplock_retention_policy_delete default %s", o._statusCode, payload)
}

func (o *SnaplockRetentionPolicyDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockRetentionPolicyDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
