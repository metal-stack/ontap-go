// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// LocalCifsUserCreateReader is a Reader for the LocalCifsUserCreate structure.
type LocalCifsUserCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalCifsUserCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewLocalCifsUserCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalCifsUserCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalCifsUserCreateCreated creates a LocalCifsUserCreateCreated with default headers values
func NewLocalCifsUserCreateCreated() *LocalCifsUserCreateCreated {
	return &LocalCifsUserCreateCreated{}
}

/*
LocalCifsUserCreateCreated describes a response with status code 201, with default header values.

Created
*/
type LocalCifsUserCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.LocalCifsUserResponse
}

// IsSuccess returns true when this local cifs user create created response has a 2xx status code
func (o *LocalCifsUserCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local cifs user create created response has a 3xx status code
func (o *LocalCifsUserCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local cifs user create created response has a 4xx status code
func (o *LocalCifsUserCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this local cifs user create created response has a 5xx status code
func (o *LocalCifsUserCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this local cifs user create created response a status code equal to that given
func (o *LocalCifsUserCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the local cifs user create created response
func (o *LocalCifsUserCreateCreated) Code() int {
	return 201
}

func (o *LocalCifsUserCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/cifs/local-users][%d] localCifsUserCreateCreated %s", 201, payload)
}

func (o *LocalCifsUserCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/cifs/local-users][%d] localCifsUserCreateCreated %s", 201, payload)
}

func (o *LocalCifsUserCreateCreated) GetPayload() *models.LocalCifsUserResponse {
	return o.Payload
}

func (o *LocalCifsUserCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.LocalCifsUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocalCifsUserCreateDefault creates a LocalCifsUserCreateDefault with default headers values
func NewLocalCifsUserCreateDefault(code int) *LocalCifsUserCreateDefault {
	return &LocalCifsUserCreateDefault{
		_statusCode: code,
	}
}

/*
	LocalCifsUserCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262278     | Name and password are required fields. |
| 655399     | CIFS server must exist to create a local user. |
| 655660     | The operation is allowed only on data SVMs. |
| 655661     | The user name should not exceed 20 characters. Also full_name and description should not exceed 256 characters. |
| 655665     | The user name must not match the CIFS server name of the specified SVM. |
| 655668     | The specified user name contains illegal characters. |
| 655675     | The local domain name specified in user name doesn't exist. |
| 655682     | The user name cannot be blank. |
| 655733     | The password does not meet the password complexity requirements. |
| 655736     | The specified user name already exists. |
| 2621706    | The specified SVM UUID is incorrect for the specified SVM name. |
*/
type LocalCifsUserCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this local cifs user create default response has a 2xx status code
func (o *LocalCifsUserCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this local cifs user create default response has a 3xx status code
func (o *LocalCifsUserCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this local cifs user create default response has a 4xx status code
func (o *LocalCifsUserCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this local cifs user create default response has a 5xx status code
func (o *LocalCifsUserCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this local cifs user create default response a status code equal to that given
func (o *LocalCifsUserCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the local cifs user create default response
func (o *LocalCifsUserCreateDefault) Code() int {
	return o._statusCode
}

func (o *LocalCifsUserCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/cifs/local-users][%d] local_cifs_user_create default %s", o._statusCode, payload)
}

func (o *LocalCifsUserCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/cifs/local-users][%d] local_cifs_user_create default %s", o._statusCode, payload)
}

func (o *LocalCifsUserCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalCifsUserCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
