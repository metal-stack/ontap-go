// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// IscsiCredentialsDeleteReader is a Reader for the IscsiCredentialsDelete structure.
type IscsiCredentialsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IscsiCredentialsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIscsiCredentialsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIscsiCredentialsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIscsiCredentialsDeleteOK creates a IscsiCredentialsDeleteOK with default headers values
func NewIscsiCredentialsDeleteOK() *IscsiCredentialsDeleteOK {
	return &IscsiCredentialsDeleteOK{}
}

/*
IscsiCredentialsDeleteOK describes a response with status code 200, with default header values.

OK
*/
type IscsiCredentialsDeleteOK struct {
}

// IsSuccess returns true when this iscsi credentials delete o k response has a 2xx status code
func (o *IscsiCredentialsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iscsi credentials delete o k response has a 3xx status code
func (o *IscsiCredentialsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iscsi credentials delete o k response has a 4xx status code
func (o *IscsiCredentialsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iscsi credentials delete o k response has a 5xx status code
func (o *IscsiCredentialsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iscsi credentials delete o k response a status code equal to that given
func (o *IscsiCredentialsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the iscsi credentials delete o k response
func (o *IscsiCredentialsDeleteOK) Code() int {
	return 200
}

func (o *IscsiCredentialsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/san/iscsi/credentials/{svm.uuid}/{initiator}][%d] iscsiCredentialsDeleteOK", 200)
}

func (o *IscsiCredentialsDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/san/iscsi/credentials/{svm.uuid}/{initiator}][%d] iscsiCredentialsDeleteOK", 200)
}

func (o *IscsiCredentialsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIscsiCredentialsDeleteDefault creates a IscsiCredentialsDeleteDefault with default headers values
func NewIscsiCredentialsDeleteDefault(code int) *IscsiCredentialsDeleteDefault {
	return &IscsiCredentialsDeleteDefault{
		_statusCode: code,
	}
}

/*
	IscsiCredentialsDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621462 | An SVM with the specified UUID does not exist. |
| 2621706 | Both the SVM UUID and SVM name were supplied, but they do not refer to the same SVM. |
| 2621707 | No SVM was specified. Either `svm.name` or `svm.uuid` must be supplied. |
| 5374148 | The default security credential cannot be deleted for an SVM. |
| 5374895 | The iSCSI security credential does not exist on the specified SVM. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type IscsiCredentialsDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this iscsi credentials delete default response has a 2xx status code
func (o *IscsiCredentialsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iscsi credentials delete default response has a 3xx status code
func (o *IscsiCredentialsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iscsi credentials delete default response has a 4xx status code
func (o *IscsiCredentialsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iscsi credentials delete default response has a 5xx status code
func (o *IscsiCredentialsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iscsi credentials delete default response a status code equal to that given
func (o *IscsiCredentialsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the iscsi credentials delete default response
func (o *IscsiCredentialsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IscsiCredentialsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/san/iscsi/credentials/{svm.uuid}/{initiator}][%d] iscsi_credentials_delete default %s", o._statusCode, payload)
}

func (o *IscsiCredentialsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/san/iscsi/credentials/{svm.uuid}/{initiator}][%d] iscsi_credentials_delete default %s", o._statusCode, payload)
}

func (o *IscsiCredentialsDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IscsiCredentialsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
