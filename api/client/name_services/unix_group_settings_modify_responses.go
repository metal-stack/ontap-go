// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// UnixGroupSettingsModifyReader is a Reader for the UnixGroupSettingsModify structure.
type UnixGroupSettingsModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnixGroupSettingsModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnixGroupSettingsModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnixGroupSettingsModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnixGroupSettingsModifyOK creates a UnixGroupSettingsModifyOK with default headers values
func NewUnixGroupSettingsModifyOK() *UnixGroupSettingsModifyOK {
	return &UnixGroupSettingsModifyOK{}
}

/*
UnixGroupSettingsModifyOK describes a response with status code 200, with default header values.

OK
*/
type UnixGroupSettingsModifyOK struct {
}

// IsSuccess returns true when this unix group settings modify o k response has a 2xx status code
func (o *UnixGroupSettingsModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unix group settings modify o k response has a 3xx status code
func (o *UnixGroupSettingsModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unix group settings modify o k response has a 4xx status code
func (o *UnixGroupSettingsModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unix group settings modify o k response has a 5xx status code
func (o *UnixGroupSettingsModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unix group settings modify o k response a status code equal to that given
func (o *UnixGroupSettingsModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the unix group settings modify o k response
func (o *UnixGroupSettingsModifyOK) Code() int {
	return 200
}

func (o *UnixGroupSettingsModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /name-services/cache/unix-group/settings/{svm.uuid}][%d] unixGroupSettingsModifyOK", 200)
}

func (o *UnixGroupSettingsModifyOK) String() string {
	return fmt.Sprintf("[PATCH /name-services/cache/unix-group/settings/{svm.uuid}][%d] unixGroupSettingsModifyOK", 200)
}

func (o *UnixGroupSettingsModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnixGroupSettingsModifyDefault creates a UnixGroupSettingsModifyDefault with default headers values
func NewUnixGroupSettingsModifyDefault(code int) *UnixGroupSettingsModifyDefault {
	return &UnixGroupSettingsModifyDefault{
		_statusCode: code,
	}
}

/*
	UnixGroupSettingsModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 23724055 | Internal error. Configuration for Vserver failed. Verify that the cluster is healthy, then try the command again. For further assistance, contact technical support. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type UnixGroupSettingsModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this unix group settings modify default response has a 2xx status code
func (o *UnixGroupSettingsModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unix group settings modify default response has a 3xx status code
func (o *UnixGroupSettingsModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unix group settings modify default response has a 4xx status code
func (o *UnixGroupSettingsModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unix group settings modify default response has a 5xx status code
func (o *UnixGroupSettingsModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unix group settings modify default response a status code equal to that given
func (o *UnixGroupSettingsModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unix group settings modify default response
func (o *UnixGroupSettingsModifyDefault) Code() int {
	return o._statusCode
}

func (o *UnixGroupSettingsModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/cache/unix-group/settings/{svm.uuid}][%d] unix_group_settings_modify default %s", o._statusCode, payload)
}

func (o *UnixGroupSettingsModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/cache/unix-group/settings/{svm.uuid}][%d] unix_group_settings_modify default %s", o._statusCode, payload)
}

func (o *UnixGroupSettingsModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnixGroupSettingsModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
