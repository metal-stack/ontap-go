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

// SvmPeerDeleteReader is a Reader for the SvmPeerDelete structure.
type SvmPeerDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmPeerDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvmPeerDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSvmPeerDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmPeerDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmPeerDeleteOK creates a SvmPeerDeleteOK with default headers values
func NewSvmPeerDeleteOK() *SvmPeerDeleteOK {
	return &SvmPeerDeleteOK{}
}

/*
SvmPeerDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SvmPeerDeleteOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this svm peer delete o k response has a 2xx status code
func (o *SvmPeerDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm peer delete o k response has a 3xx status code
func (o *SvmPeerDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm peer delete o k response has a 4xx status code
func (o *SvmPeerDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm peer delete o k response has a 5xx status code
func (o *SvmPeerDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this svm peer delete o k response a status code equal to that given
func (o *SvmPeerDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the svm peer delete o k response
func (o *SvmPeerDeleteOK) Code() int {
	return 200
}

func (o *SvmPeerDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/peers/{uuid}][%d] svmPeerDeleteOK %s", 200, payload)
}

func (o *SvmPeerDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/peers/{uuid}][%d] svmPeerDeleteOK %s", 200, payload)
}

func (o *SvmPeerDeleteOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SvmPeerDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmPeerDeleteAccepted creates a SvmPeerDeleteAccepted with default headers values
func NewSvmPeerDeleteAccepted() *SvmPeerDeleteAccepted {
	return &SvmPeerDeleteAccepted{}
}

/*
SvmPeerDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SvmPeerDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this svm peer delete accepted response has a 2xx status code
func (o *SvmPeerDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm peer delete accepted response has a 3xx status code
func (o *SvmPeerDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm peer delete accepted response has a 4xx status code
func (o *SvmPeerDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm peer delete accepted response has a 5xx status code
func (o *SvmPeerDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this svm peer delete accepted response a status code equal to that given
func (o *SvmPeerDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the svm peer delete accepted response
func (o *SvmPeerDeleteAccepted) Code() int {
	return 202
}

func (o *SvmPeerDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/peers/{uuid}][%d] svmPeerDeleteAccepted %s", 202, payload)
}

func (o *SvmPeerDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/peers/{uuid}][%d] svmPeerDeleteAccepted %s", 202, payload)
}

func (o *SvmPeerDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SvmPeerDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmPeerDeleteDefault creates a SvmPeerDeleteDefault with default headers values
func NewSvmPeerDeleteDefault(code int) *SvmPeerDeleteDefault {
	return &SvmPeerDeleteDefault{
		_statusCode: code,
	}
}

/*
	SvmPeerDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

<br/>
```
| Error codes | Description |
| ----------- | ----------- |
| 26345578    | Internal error. Unable to retrieve local or peer SVM name. |
| 9895956     | Cannot delete an SVM that is part of an SVM peer or transition peer relationship. |
```
<br/>
*/
type SvmPeerDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm peer delete default response has a 2xx status code
func (o *SvmPeerDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm peer delete default response has a 3xx status code
func (o *SvmPeerDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm peer delete default response has a 4xx status code
func (o *SvmPeerDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm peer delete default response has a 5xx status code
func (o *SvmPeerDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm peer delete default response a status code equal to that given
func (o *SvmPeerDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm peer delete default response
func (o *SvmPeerDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SvmPeerDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/peers/{uuid}][%d] svm_peer_delete default %s", o._statusCode, payload)
}

func (o *SvmPeerDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/peers/{uuid}][%d] svm_peer_delete default %s", o._statusCode, payload)
}

func (o *SvmPeerDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmPeerDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
