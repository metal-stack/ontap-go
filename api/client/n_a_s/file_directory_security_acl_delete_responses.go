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

// FileDirectorySecurityACLDeleteReader is a Reader for the FileDirectorySecurityACLDelete structure.
type FileDirectorySecurityACLDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileDirectorySecurityACLDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFileDirectorySecurityACLDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewFileDirectorySecurityACLDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileDirectorySecurityACLDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileDirectorySecurityACLDeleteOK creates a FileDirectorySecurityACLDeleteOK with default headers values
func NewFileDirectorySecurityACLDeleteOK() *FileDirectorySecurityACLDeleteOK {
	return &FileDirectorySecurityACLDeleteOK{}
}

/*
FileDirectorySecurityACLDeleteOK describes a response with status code 200, with default header values.

OK
*/
type FileDirectorySecurityACLDeleteOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this file directory security Acl delete o k response has a 2xx status code
func (o *FileDirectorySecurityACLDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file directory security Acl delete o k response has a 3xx status code
func (o *FileDirectorySecurityACLDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file directory security Acl delete o k response has a 4xx status code
func (o *FileDirectorySecurityACLDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this file directory security Acl delete o k response has a 5xx status code
func (o *FileDirectorySecurityACLDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this file directory security Acl delete o k response a status code equal to that given
func (o *FileDirectorySecurityACLDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the file directory security Acl delete o k response
func (o *FileDirectorySecurityACLDeleteOK) Code() int {
	return 200
}

func (o *FileDirectorySecurityACLDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/file-security/permissions/{svm.uuid}/{path}/acl/{user}][%d] fileDirectorySecurityAclDeleteOK %s", 200, payload)
}

func (o *FileDirectorySecurityACLDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/file-security/permissions/{svm.uuid}/{path}/acl/{user}][%d] fileDirectorySecurityAclDeleteOK %s", 200, payload)
}

func (o *FileDirectorySecurityACLDeleteOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *FileDirectorySecurityACLDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileDirectorySecurityACLDeleteAccepted creates a FileDirectorySecurityACLDeleteAccepted with default headers values
func NewFileDirectorySecurityACLDeleteAccepted() *FileDirectorySecurityACLDeleteAccepted {
	return &FileDirectorySecurityACLDeleteAccepted{}
}

/*
FileDirectorySecurityACLDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type FileDirectorySecurityACLDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this file directory security Acl delete accepted response has a 2xx status code
func (o *FileDirectorySecurityACLDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file directory security Acl delete accepted response has a 3xx status code
func (o *FileDirectorySecurityACLDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file directory security Acl delete accepted response has a 4xx status code
func (o *FileDirectorySecurityACLDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this file directory security Acl delete accepted response has a 5xx status code
func (o *FileDirectorySecurityACLDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this file directory security Acl delete accepted response a status code equal to that given
func (o *FileDirectorySecurityACLDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the file directory security Acl delete accepted response
func (o *FileDirectorySecurityACLDeleteAccepted) Code() int {
	return 202
}

func (o *FileDirectorySecurityACLDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/file-security/permissions/{svm.uuid}/{path}/acl/{user}][%d] fileDirectorySecurityAclDeleteAccepted %s", 202, payload)
}

func (o *FileDirectorySecurityACLDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/file-security/permissions/{svm.uuid}/{path}/acl/{user}][%d] fileDirectorySecurityAclDeleteAccepted %s", 202, payload)
}

func (o *FileDirectorySecurityACLDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *FileDirectorySecurityACLDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileDirectorySecurityACLDeleteDefault creates a FileDirectorySecurityACLDeleteDefault with default headers values
func NewFileDirectorySecurityACLDeleteDefault(code int) *FileDirectorySecurityACLDeleteDefault {
	return &FileDirectorySecurityACLDeleteDefault{
		_statusCode: code,
	}
}

/*
	FileDirectorySecurityACLDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655865     | The specified file or directory does not exist.|
| 10485811   | Access is a required field.|
| 1260882    | Specified SVM not found.|
| 6691623    | User is not authorized.|
| 4849676    | The specified Windows user or group does not exist.|
| 10485813   | All values cannot be false.|
*/
type FileDirectorySecurityACLDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this file directory security acl delete default response has a 2xx status code
func (o *FileDirectorySecurityACLDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this file directory security acl delete default response has a 3xx status code
func (o *FileDirectorySecurityACLDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this file directory security acl delete default response has a 4xx status code
func (o *FileDirectorySecurityACLDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this file directory security acl delete default response has a 5xx status code
func (o *FileDirectorySecurityACLDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this file directory security acl delete default response a status code equal to that given
func (o *FileDirectorySecurityACLDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the file directory security acl delete default response
func (o *FileDirectorySecurityACLDeleteDefault) Code() int {
	return o._statusCode
}

func (o *FileDirectorySecurityACLDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/file-security/permissions/{svm.uuid}/{path}/acl/{user}][%d] file_directory_security_acl_delete default %s", o._statusCode, payload)
}

func (o *FileDirectorySecurityACLDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/file-security/permissions/{svm.uuid}/{path}/acl/{user}][%d] file_directory_security_acl_delete default %s", o._statusCode, payload)
}

func (o *FileDirectorySecurityACLDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileDirectorySecurityACLDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}