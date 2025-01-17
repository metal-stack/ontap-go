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

// CifsShareModifyReader is a Reader for the CifsShareModify structure.
type CifsShareModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsShareModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsShareModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsShareModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsShareModifyOK creates a CifsShareModifyOK with default headers values
func NewCifsShareModifyOK() *CifsShareModifyOK {
	return &CifsShareModifyOK{}
}

/*
CifsShareModifyOK describes a response with status code 200, with default header values.

OK
*/
type CifsShareModifyOK struct {
}

// IsSuccess returns true when this cifs share modify o k response has a 2xx status code
func (o *CifsShareModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs share modify o k response has a 3xx status code
func (o *CifsShareModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs share modify o k response has a 4xx status code
func (o *CifsShareModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs share modify o k response has a 5xx status code
func (o *CifsShareModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs share modify o k response a status code equal to that given
func (o *CifsShareModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs share modify o k response
func (o *CifsShareModifyOK) Code() int {
	return 200
}

func (o *CifsShareModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/shares/{svm.uuid}/{name}][%d] cifsShareModifyOK", 200)
}

func (o *CifsShareModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/shares/{svm.uuid}/{name}][%d] cifsShareModifyOK", 200)
}

func (o *CifsShareModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCifsShareModifyDefault creates a CifsShareModifyDefault with default headers values
func NewCifsShareModifyDefault(code int) *CifsShareModifyDefault {
	return &CifsShareModifyDefault{
		_statusCode: code,
	}
}

/*
	CifsShareModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655628     | 'SMB_ENCRYPTION' property cannot be set on CIFS share because the CIFS server does not support SMB3.0 |
| 655551     | CIFS Share modification failed because the specified path does not exist |
| 655620     | Cannot set symlink properties for admin shares                           |
| 656420     | Cannot modify the standard share ipc$                                    |
| 656421     | Cannot modify the standard share admin$                                  |
| 656422     | Failed to modify the home directory share because the directory shares must specify a path relative to one or more home directory search paths |
| 656423     | Failed to modify CIFS share. The Shares must define an absolute share path |
| 656425     | Failed to modify the CIFS share because the path for an administrative share cannot be modified |
| 655395     | Failed to modify the CIFS share because share cannot be made continuously available unless running SMB3 or later. |
| 4849678    | Failed to modify the CIFS share because the specified UNIX group does not exist |
| 655622     | Invalid value for parameter {max-connections-per-share}. Maximum connections on CIFS share {name} must be between 1 to 4294967295.|
*/
type CifsShareModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs share modify default response has a 2xx status code
func (o *CifsShareModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs share modify default response has a 3xx status code
func (o *CifsShareModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs share modify default response has a 4xx status code
func (o *CifsShareModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs share modify default response has a 5xx status code
func (o *CifsShareModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs share modify default response a status code equal to that given
func (o *CifsShareModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs share modify default response
func (o *CifsShareModifyDefault) Code() int {
	return o._statusCode
}

func (o *CifsShareModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/shares/{svm.uuid}/{name}][%d] cifs_share_modify default %s", o._statusCode, payload)
}

func (o *CifsShareModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/shares/{svm.uuid}/{name}][%d] cifs_share_modify default %s", o._statusCode, payload)
}

func (o *CifsShareModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsShareModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}