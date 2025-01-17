// Code generated by go-swagger; DO NOT EDIT.

package s_vm

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

// SvmPeerPermissionCreateReader is a Reader for the SvmPeerPermissionCreate structure.
type SvmPeerPermissionCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmPeerPermissionCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSvmPeerPermissionCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmPeerPermissionCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmPeerPermissionCreateCreated creates a SvmPeerPermissionCreateCreated with default headers values
func NewSvmPeerPermissionCreateCreated() *SvmPeerPermissionCreateCreated {
	return &SvmPeerPermissionCreateCreated{}
}

/*
SvmPeerPermissionCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SvmPeerPermissionCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.SvmPeerPermission
}

// IsSuccess returns true when this svm peer permission create created response has a 2xx status code
func (o *SvmPeerPermissionCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm peer permission create created response has a 3xx status code
func (o *SvmPeerPermissionCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm peer permission create created response has a 4xx status code
func (o *SvmPeerPermissionCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm peer permission create created response has a 5xx status code
func (o *SvmPeerPermissionCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this svm peer permission create created response a status code equal to that given
func (o *SvmPeerPermissionCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the svm peer permission create created response
func (o *SvmPeerPermissionCreateCreated) Code() int {
	return 201
}

func (o *SvmPeerPermissionCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/peer-permissions][%d] svmPeerPermissionCreateCreated %s", 201, payload)
}

func (o *SvmPeerPermissionCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/peer-permissions][%d] svmPeerPermissionCreateCreated %s", 201, payload)
}

func (o *SvmPeerPermissionCreateCreated) GetPayload() *models.SvmPeerPermission {
	return o.Payload
}

func (o *SvmPeerPermissionCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.SvmPeerPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmPeerPermissionCreateDefault creates a SvmPeerPermissionCreateDefault with default headers values
func NewSvmPeerPermissionCreateDefault(code int) *SvmPeerPermissionCreateDefault {
	return &SvmPeerPermissionCreateDefault{
		_statusCode: code,
	}
}

/*
	SvmPeerPermissionCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

<br/>
```
| Error codes | Description |
| ----------- | ----------- |
| 26345572    | {field} is a required field. |
| 26345573    | Failed to find the SVM or volume UUID with name. |
| 26345574    | Failed to find the SVM or volume name with UUID. |
| 26345575    | The specified peer cluster name and peer cluster UUID do not match. |
| 9896057     | SVM peer permission already exists. |
```
<br/>
*/
type SvmPeerPermissionCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm peer permission create default response has a 2xx status code
func (o *SvmPeerPermissionCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm peer permission create default response has a 3xx status code
func (o *SvmPeerPermissionCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm peer permission create default response has a 4xx status code
func (o *SvmPeerPermissionCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm peer permission create default response has a 5xx status code
func (o *SvmPeerPermissionCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm peer permission create default response a status code equal to that given
func (o *SvmPeerPermissionCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm peer permission create default response
func (o *SvmPeerPermissionCreateDefault) Code() int {
	return o._statusCode
}

func (o *SvmPeerPermissionCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/peer-permissions][%d] svm_peer_permission_create default %s", o._statusCode, payload)
}

func (o *SvmPeerPermissionCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/peer-permissions][%d] svm_peer_permission_create default %s", o._statusCode, payload)
}

func (o *SvmPeerPermissionCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmPeerPermissionCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
