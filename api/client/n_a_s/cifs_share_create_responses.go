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

// CifsShareCreateReader is a Reader for the CifsShareCreate structure.
type CifsShareCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsShareCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCifsShareCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsShareCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsShareCreateCreated creates a CifsShareCreateCreated with default headers values
func NewCifsShareCreateCreated() *CifsShareCreateCreated {
	return &CifsShareCreateCreated{}
}

/*
CifsShareCreateCreated describes a response with status code 201, with default header values.

Created
*/
type CifsShareCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this cifs share create created response has a 2xx status code
func (o *CifsShareCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs share create created response has a 3xx status code
func (o *CifsShareCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs share create created response has a 4xx status code
func (o *CifsShareCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs share create created response has a 5xx status code
func (o *CifsShareCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs share create created response a status code equal to that given
func (o *CifsShareCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the cifs share create created response
func (o *CifsShareCreateCreated) Code() int {
	return 201
}

func (o *CifsShareCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/cifs/shares][%d] cifsShareCreateCreated", 201)
}

func (o *CifsShareCreateCreated) String() string {
	return fmt.Sprintf("[POST /protocols/cifs/shares][%d] cifsShareCreateCreated", 201)
}

func (o *CifsShareCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewCifsShareCreateDefault creates a CifsShareCreateDefault with default headers values
func NewCifsShareCreateDefault(code int) *CifsShareCreateDefault {
	return &CifsShareCreateDefault{
		_statusCode: code,
	}
}

/*
	CifsShareCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655628     | CIFS Share Creation with property 'SMB_ENCRYPTION' failed because the CIFS server does not support SMB3.0 |
| 655551     | CIFS Share Creation failed because the specified path does not exist |
| 655577     | The CIFS share name cannot be more than 80 characters long |
| 655399     | Failed to create CIFS share. The CIFS server does not exist for specified SVM |
| 656422     | Failed to create the home directory share because the directory shares must specify a path relative to one or more home directory search paths |
| 656423     | Failed to create CIFS share. The Shares must define an absolute share path |
| 656424     | Failed to create CIFS the administrator share 'c$' because you are not permitted to created any admin shares |
| 655625     | Failed to create CIFS share. The Shares path is not a valid file-type for CIFS share |
| 656426     | CIFS Share Creation failed because the share name is invalid |
| 655655     | no-strict-security should be set to true only if unix_symlink is configured as "local" or "widelink" |
| 655394     | Failed to create CIFS share because share cannot be made continuously available unless running SMB3 or later. |
| 4849678    | Failed to create CIFS share because the specified UNIX group does not exist |
| 655622     | Invalid value for parameter {max-connections-per-share}. Maximum connections on CIFS share {name} must be between 1 to 4294967295.|
| 655446     | Failed to resolve the security identifier (SID) for the account named {account_name}. Reason: {Reason for failure}.|
| 2621706    | The specified SVM UUID is incorrect for the specified SVM name |
*/
type CifsShareCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs share create default response has a 2xx status code
func (o *CifsShareCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs share create default response has a 3xx status code
func (o *CifsShareCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs share create default response has a 4xx status code
func (o *CifsShareCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs share create default response has a 5xx status code
func (o *CifsShareCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs share create default response a status code equal to that given
func (o *CifsShareCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs share create default response
func (o *CifsShareCreateDefault) Code() int {
	return o._statusCode
}

func (o *CifsShareCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/cifs/shares][%d] cifs_share_create default %s", o._statusCode, payload)
}

func (o *CifsShareCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/cifs/shares][%d] cifs_share_create default %s", o._statusCode, payload)
}

func (o *CifsShareCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsShareCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
