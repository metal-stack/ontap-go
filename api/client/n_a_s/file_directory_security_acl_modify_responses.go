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

// FileDirectorySecurityACLModifyReader is a Reader for the FileDirectorySecurityACLModify structure.
type FileDirectorySecurityACLModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileDirectorySecurityACLModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFileDirectorySecurityACLModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewFileDirectorySecurityACLModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileDirectorySecurityACLModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileDirectorySecurityACLModifyOK creates a FileDirectorySecurityACLModifyOK with default headers values
func NewFileDirectorySecurityACLModifyOK() *FileDirectorySecurityACLModifyOK {
	return &FileDirectorySecurityACLModifyOK{}
}

/*
FileDirectorySecurityACLModifyOK describes a response with status code 200, with default header values.

OK
*/
type FileDirectorySecurityACLModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this file directory security Acl modify o k response has a 2xx status code
func (o *FileDirectorySecurityACLModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file directory security Acl modify o k response has a 3xx status code
func (o *FileDirectorySecurityACLModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file directory security Acl modify o k response has a 4xx status code
func (o *FileDirectorySecurityACLModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this file directory security Acl modify o k response has a 5xx status code
func (o *FileDirectorySecurityACLModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this file directory security Acl modify o k response a status code equal to that given
func (o *FileDirectorySecurityACLModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the file directory security Acl modify o k response
func (o *FileDirectorySecurityACLModifyOK) Code() int {
	return 200
}

func (o *FileDirectorySecurityACLModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/file-security/permissions/{svm.uuid}/{path}/acl/{user}][%d] fileDirectorySecurityAclModifyOK %s", 200, payload)
}

func (o *FileDirectorySecurityACLModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/file-security/permissions/{svm.uuid}/{path}/acl/{user}][%d] fileDirectorySecurityAclModifyOK %s", 200, payload)
}

func (o *FileDirectorySecurityACLModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *FileDirectorySecurityACLModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileDirectorySecurityACLModifyAccepted creates a FileDirectorySecurityACLModifyAccepted with default headers values
func NewFileDirectorySecurityACLModifyAccepted() *FileDirectorySecurityACLModifyAccepted {
	return &FileDirectorySecurityACLModifyAccepted{}
}

/*
FileDirectorySecurityACLModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type FileDirectorySecurityACLModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this file directory security Acl modify accepted response has a 2xx status code
func (o *FileDirectorySecurityACLModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file directory security Acl modify accepted response has a 3xx status code
func (o *FileDirectorySecurityACLModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file directory security Acl modify accepted response has a 4xx status code
func (o *FileDirectorySecurityACLModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this file directory security Acl modify accepted response has a 5xx status code
func (o *FileDirectorySecurityACLModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this file directory security Acl modify accepted response a status code equal to that given
func (o *FileDirectorySecurityACLModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the file directory security Acl modify accepted response
func (o *FileDirectorySecurityACLModifyAccepted) Code() int {
	return 202
}

func (o *FileDirectorySecurityACLModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/file-security/permissions/{svm.uuid}/{path}/acl/{user}][%d] fileDirectorySecurityAclModifyAccepted %s", 202, payload)
}

func (o *FileDirectorySecurityACLModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/file-security/permissions/{svm.uuid}/{path}/acl/{user}][%d] fileDirectorySecurityAclModifyAccepted %s", 202, payload)
}

func (o *FileDirectorySecurityACLModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *FileDirectorySecurityACLModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileDirectorySecurityACLModifyDefault creates a FileDirectorySecurityACLModifyDefault with default headers values
func NewFileDirectorySecurityACLModifyDefault(code int) *FileDirectorySecurityACLModifyDefault {
	return &FileDirectorySecurityACLModifyDefault{
		_statusCode: code,
	}
}

/*
	FileDirectorySecurityACLModifyDefault describes a response with status code -1, with default header values.

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
type FileDirectorySecurityACLModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this file directory security acl modify default response has a 2xx status code
func (o *FileDirectorySecurityACLModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this file directory security acl modify default response has a 3xx status code
func (o *FileDirectorySecurityACLModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this file directory security acl modify default response has a 4xx status code
func (o *FileDirectorySecurityACLModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this file directory security acl modify default response has a 5xx status code
func (o *FileDirectorySecurityACLModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this file directory security acl modify default response a status code equal to that given
func (o *FileDirectorySecurityACLModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the file directory security acl modify default response
func (o *FileDirectorySecurityACLModifyDefault) Code() int {
	return o._statusCode
}

func (o *FileDirectorySecurityACLModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/file-security/permissions/{svm.uuid}/{path}/acl/{user}][%d] file_directory_security_acl_modify default %s", o._statusCode, payload)
}

func (o *FileDirectorySecurityACLModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/file-security/permissions/{svm.uuid}/{path}/acl/{user}][%d] file_directory_security_acl_modify default %s", o._statusCode, payload)
}

func (o *FileDirectorySecurityACLModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileDirectorySecurityACLModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
