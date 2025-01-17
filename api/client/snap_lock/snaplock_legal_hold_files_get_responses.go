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

// SnaplockLegalHoldFilesGetReader is a Reader for the SnaplockLegalHoldFilesGet structure.
type SnaplockLegalHoldFilesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockLegalHoldFilesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockLegalHoldFilesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockLegalHoldFilesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockLegalHoldFilesGetOK creates a SnaplockLegalHoldFilesGetOK with default headers values
func NewSnaplockLegalHoldFilesGetOK() *SnaplockLegalHoldFilesGetOK {
	return &SnaplockLegalHoldFilesGetOK{}
}

/*
SnaplockLegalHoldFilesGetOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockLegalHoldFilesGetOK struct {
	Payload *models.SnaplockLitigationFileResponse
}

// IsSuccess returns true when this snaplock legal hold files get o k response has a 2xx status code
func (o *SnaplockLegalHoldFilesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock legal hold files get o k response has a 3xx status code
func (o *SnaplockLegalHoldFilesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock legal hold files get o k response has a 4xx status code
func (o *SnaplockLegalHoldFilesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock legal hold files get o k response has a 5xx status code
func (o *SnaplockLegalHoldFilesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock legal hold files get o k response a status code equal to that given
func (o *SnaplockLegalHoldFilesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snaplock legal hold files get o k response
func (o *SnaplockLegalHoldFilesGetOK) Code() int {
	return 200
}

func (o *SnaplockLegalHoldFilesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/litigations/{litigation.id}/files][%d] snaplockLegalHoldFilesGetOK %s", 200, payload)
}

func (o *SnaplockLegalHoldFilesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/litigations/{litigation.id}/files][%d] snaplockLegalHoldFilesGetOK %s", 200, payload)
}

func (o *SnaplockLegalHoldFilesGetOK) GetPayload() *models.SnaplockLitigationFileResponse {
	return o.Payload
}

func (o *SnaplockLegalHoldFilesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockLitigationFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLegalHoldFilesGetDefault creates a SnaplockLegalHoldFilesGetDefault with default headers values
func NewSnaplockLegalHoldFilesGetDefault(code int) *SnaplockLegalHoldFilesGetDefault {
	return &SnaplockLegalHoldFilesGetDefault{
		_statusCode: code,
	}
}

/*
	SnaplockLegalHoldFilesGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 13763280    | Only a user with security login role \"vsadmin-snaplock\" is allowed to perform this operation.  |
*/
type SnaplockLegalHoldFilesGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snaplock legal hold files get default response has a 2xx status code
func (o *SnaplockLegalHoldFilesGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock legal hold files get default response has a 3xx status code
func (o *SnaplockLegalHoldFilesGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock legal hold files get default response has a 4xx status code
func (o *SnaplockLegalHoldFilesGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock legal hold files get default response has a 5xx status code
func (o *SnaplockLegalHoldFilesGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock legal hold files get default response a status code equal to that given
func (o *SnaplockLegalHoldFilesGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snaplock legal hold files get default response
func (o *SnaplockLegalHoldFilesGetDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockLegalHoldFilesGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/litigations/{litigation.id}/files][%d] snaplock_legal_hold_files_get default %s", o._statusCode, payload)
}

func (o *SnaplockLegalHoldFilesGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/snaplock/litigations/{litigation.id}/files][%d] snaplock_legal_hold_files_get default %s", o._statusCode, payload)
}

func (o *SnaplockLegalHoldFilesGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockLegalHoldFilesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}