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

// LunAttributeCreateReader is a Reader for the LunAttributeCreate structure.
type LunAttributeCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunAttributeCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewLunAttributeCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunAttributeCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunAttributeCreateCreated creates a LunAttributeCreateCreated with default headers values
func NewLunAttributeCreateCreated() *LunAttributeCreateCreated {
	return &LunAttributeCreateCreated{}
}

/*
LunAttributeCreateCreated describes a response with status code 201, with default header values.

Created
*/
type LunAttributeCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.LunAttributeResponse
}

// IsSuccess returns true when this lun attribute create created response has a 2xx status code
func (o *LunAttributeCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lun attribute create created response has a 3xx status code
func (o *LunAttributeCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lun attribute create created response has a 4xx status code
func (o *LunAttributeCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this lun attribute create created response has a 5xx status code
func (o *LunAttributeCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this lun attribute create created response a status code equal to that given
func (o *LunAttributeCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the lun attribute create created response
func (o *LunAttributeCreateCreated) Code() int {
	return 201
}

func (o *LunAttributeCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/luns/{lun.uuid}/attributes][%d] lunAttributeCreateCreated %s", 201, payload)
}

func (o *LunAttributeCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/luns/{lun.uuid}/attributes][%d] lunAttributeCreateCreated %s", 201, payload)
}

func (o *LunAttributeCreateCreated) GetPayload() *models.LunAttributeResponse {
	return o.Payload
}

func (o *LunAttributeCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.LunAttributeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLunAttributeCreateDefault creates a LunAttributeCreateDefault with default headers values
func NewLunAttributeCreateDefault(code int) *LunAttributeCreateDefault {
	return &LunAttributeCreateDefault{
		_statusCode: code,
	}
}

/*
	LunAttributeCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5374875 | The specified LUN was not found. |
| 5374928 | An incomplete attribute name/value pair was supplied. |
| 5374929 | The combined sizes of an attribute name and value are too large. |
| 5374930 | The attribute already exists for the LUN. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type LunAttributeCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this lun attribute create default response has a 2xx status code
func (o *LunAttributeCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lun attribute create default response has a 3xx status code
func (o *LunAttributeCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lun attribute create default response has a 4xx status code
func (o *LunAttributeCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lun attribute create default response has a 5xx status code
func (o *LunAttributeCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lun attribute create default response a status code equal to that given
func (o *LunAttributeCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lun attribute create default response
func (o *LunAttributeCreateDefault) Code() int {
	return o._statusCode
}

func (o *LunAttributeCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/luns/{lun.uuid}/attributes][%d] lun_attribute_create default %s", o._statusCode, payload)
}

func (o *LunAttributeCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/luns/{lun.uuid}/attributes][%d] lun_attribute_create default %s", o._statusCode, payload)
}

func (o *LunAttributeCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunAttributeCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
