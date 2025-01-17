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

// SvmPeerCreateReader is a Reader for the SvmPeerCreate structure.
type SvmPeerCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmPeerCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSvmPeerCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSvmPeerCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmPeerCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmPeerCreateCreated creates a SvmPeerCreateCreated with default headers values
func NewSvmPeerCreateCreated() *SvmPeerCreateCreated {
	return &SvmPeerCreateCreated{}
}

/*
SvmPeerCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SvmPeerCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.SvmPeer
}

// IsSuccess returns true when this svm peer create created response has a 2xx status code
func (o *SvmPeerCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm peer create created response has a 3xx status code
func (o *SvmPeerCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm peer create created response has a 4xx status code
func (o *SvmPeerCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm peer create created response has a 5xx status code
func (o *SvmPeerCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this svm peer create created response a status code equal to that given
func (o *SvmPeerCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the svm peer create created response
func (o *SvmPeerCreateCreated) Code() int {
	return 201
}

func (o *SvmPeerCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/peers][%d] svmPeerCreateCreated %s", 201, payload)
}

func (o *SvmPeerCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/peers][%d] svmPeerCreateCreated %s", 201, payload)
}

func (o *SvmPeerCreateCreated) GetPayload() *models.SvmPeer {
	return o.Payload
}

func (o *SvmPeerCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.SvmPeer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmPeerCreateAccepted creates a SvmPeerCreateAccepted with default headers values
func NewSvmPeerCreateAccepted() *SvmPeerCreateAccepted {
	return &SvmPeerCreateAccepted{}
}

/*
SvmPeerCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SvmPeerCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.SvmPeer
}

// IsSuccess returns true when this svm peer create accepted response has a 2xx status code
func (o *SvmPeerCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm peer create accepted response has a 3xx status code
func (o *SvmPeerCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm peer create accepted response has a 4xx status code
func (o *SvmPeerCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm peer create accepted response has a 5xx status code
func (o *SvmPeerCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this svm peer create accepted response a status code equal to that given
func (o *SvmPeerCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the svm peer create accepted response
func (o *SvmPeerCreateAccepted) Code() int {
	return 202
}

func (o *SvmPeerCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/peers][%d] svmPeerCreateAccepted %s", 202, payload)
}

func (o *SvmPeerCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/peers][%d] svmPeerCreateAccepted %s", 202, payload)
}

func (o *SvmPeerCreateAccepted) GetPayload() *models.SvmPeer {
	return o.Payload
}

func (o *SvmPeerCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.SvmPeer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmPeerCreateDefault creates a SvmPeerCreateDefault with default headers values
func NewSvmPeerCreateDefault(code int) *SvmPeerCreateDefault {
	return &SvmPeerCreateDefault{
		_statusCode: code,
	}
}

/*
	SvmPeerCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

<br/>
```
| Error codes | Description |
| ----------- | ----------- |
| 13434889    | Internal error. Wait and retry. |
| 26345575    | The specified peer cluster name and peer cluster UUID do not match. |
| 26345579    | The specified field is invalid. |
| 26345580    | SVM name or SVM UUID must be provided. |
| 9896086     | Peer SVM name conflicts with one of the following: a peer SVM in an existing SVM peer relationship, a local SVM, or an IPSpace. Use the \"name\" property to uniquely specify the peer SVM alias name. |
| 9895996     | Cannot specify lun-copy as an inter-cluster application. |
```
<br/>
*/
type SvmPeerCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm peer create default response has a 2xx status code
func (o *SvmPeerCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm peer create default response has a 3xx status code
func (o *SvmPeerCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm peer create default response has a 4xx status code
func (o *SvmPeerCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm peer create default response has a 5xx status code
func (o *SvmPeerCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm peer create default response a status code equal to that given
func (o *SvmPeerCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm peer create default response
func (o *SvmPeerCreateDefault) Code() int {
	return o._statusCode
}

func (o *SvmPeerCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/peers][%d] svm_peer_create default %s", o._statusCode, payload)
}

func (o *SvmPeerCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/peers][%d] svm_peer_create default %s", o._statusCode, payload)
}

func (o *SvmPeerCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmPeerCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
