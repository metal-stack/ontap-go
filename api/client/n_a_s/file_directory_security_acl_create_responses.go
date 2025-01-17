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

// FileDirectorySecurityACLCreateReader is a Reader for the FileDirectorySecurityACLCreate structure.
type FileDirectorySecurityACLCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileDirectorySecurityACLCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewFileDirectorySecurityACLCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewFileDirectorySecurityACLCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileDirectorySecurityACLCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileDirectorySecurityACLCreateCreated creates a FileDirectorySecurityACLCreateCreated with default headers values
func NewFileDirectorySecurityACLCreateCreated() *FileDirectorySecurityACLCreateCreated {
	return &FileDirectorySecurityACLCreateCreated{}
}

/*
FileDirectorySecurityACLCreateCreated describes a response with status code 201, with default header values.

Created
*/
type FileDirectorySecurityACLCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this file directory security Acl create created response has a 2xx status code
func (o *FileDirectorySecurityACLCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file directory security Acl create created response has a 3xx status code
func (o *FileDirectorySecurityACLCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file directory security Acl create created response has a 4xx status code
func (o *FileDirectorySecurityACLCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this file directory security Acl create created response has a 5xx status code
func (o *FileDirectorySecurityACLCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this file directory security Acl create created response a status code equal to that given
func (o *FileDirectorySecurityACLCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the file directory security Acl create created response
func (o *FileDirectorySecurityACLCreateCreated) Code() int {
	return 201
}

func (o *FileDirectorySecurityACLCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/file-security/permissions/{svm.uuid}/{path}/acl][%d] fileDirectorySecurityAclCreateCreated %s", 201, payload)
}

func (o *FileDirectorySecurityACLCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/file-security/permissions/{svm.uuid}/{path}/acl][%d] fileDirectorySecurityAclCreateCreated %s", 201, payload)
}

func (o *FileDirectorySecurityACLCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *FileDirectorySecurityACLCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewFileDirectorySecurityACLCreateAccepted creates a FileDirectorySecurityACLCreateAccepted with default headers values
func NewFileDirectorySecurityACLCreateAccepted() *FileDirectorySecurityACLCreateAccepted {
	return &FileDirectorySecurityACLCreateAccepted{}
}

/*
FileDirectorySecurityACLCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type FileDirectorySecurityACLCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this file directory security Acl create accepted response has a 2xx status code
func (o *FileDirectorySecurityACLCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file directory security Acl create accepted response has a 3xx status code
func (o *FileDirectorySecurityACLCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file directory security Acl create accepted response has a 4xx status code
func (o *FileDirectorySecurityACLCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this file directory security Acl create accepted response has a 5xx status code
func (o *FileDirectorySecurityACLCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this file directory security Acl create accepted response a status code equal to that given
func (o *FileDirectorySecurityACLCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the file directory security Acl create accepted response
func (o *FileDirectorySecurityACLCreateAccepted) Code() int {
	return 202
}

func (o *FileDirectorySecurityACLCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/file-security/permissions/{svm.uuid}/{path}/acl][%d] fileDirectorySecurityAclCreateAccepted %s", 202, payload)
}

func (o *FileDirectorySecurityACLCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/file-security/permissions/{svm.uuid}/{path}/acl][%d] fileDirectorySecurityAclCreateAccepted %s", 202, payload)
}

func (o *FileDirectorySecurityACLCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *FileDirectorySecurityACLCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewFileDirectorySecurityACLCreateDefault creates a FileDirectorySecurityACLCreateDefault with default headers values
func NewFileDirectorySecurityACLCreateDefault(code int) *FileDirectorySecurityACLCreateDefault {
	return &FileDirectorySecurityACLCreateDefault{
		_statusCode: code,
	}
}

/*
	FileDirectorySecurityACLCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655865     | The specified file or directory does not exist.|
| 10485811   | Access is a required field.|
| 1260882    | Specified SVM not found.|
| 6691623    | User is not authorized.|
| 4849676    | The specified Windows user or group does not exist.|
| 10485810   | User is a required field.|
| 10485813   | All values cannot be false.|
*/
type FileDirectorySecurityACLCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this file directory security acl create default response has a 2xx status code
func (o *FileDirectorySecurityACLCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this file directory security acl create default response has a 3xx status code
func (o *FileDirectorySecurityACLCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this file directory security acl create default response has a 4xx status code
func (o *FileDirectorySecurityACLCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this file directory security acl create default response has a 5xx status code
func (o *FileDirectorySecurityACLCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this file directory security acl create default response a status code equal to that given
func (o *FileDirectorySecurityACLCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the file directory security acl create default response
func (o *FileDirectorySecurityACLCreateDefault) Code() int {
	return o._statusCode
}

func (o *FileDirectorySecurityACLCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/file-security/permissions/{svm.uuid}/{path}/acl][%d] file_directory_security_acl_create default %s", o._statusCode, payload)
}

func (o *FileDirectorySecurityACLCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/file-security/permissions/{svm.uuid}/{path}/acl][%d] file_directory_security_acl_create default %s", o._statusCode, payload)
}

func (o *FileDirectorySecurityACLCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileDirectorySecurityACLCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
