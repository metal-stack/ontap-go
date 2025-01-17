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

// LunMapGetReader is a Reader for the LunMapGet structure.
type LunMapGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunMapGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLunMapGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunMapGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunMapGetOK creates a LunMapGetOK with default headers values
func NewLunMapGetOK() *LunMapGetOK {
	return &LunMapGetOK{}
}

/*
LunMapGetOK describes a response with status code 200, with default header values.

OK
*/
type LunMapGetOK struct {
	Payload *models.LunMap
}

// IsSuccess returns true when this lun map get o k response has a 2xx status code
func (o *LunMapGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lun map get o k response has a 3xx status code
func (o *LunMapGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lun map get o k response has a 4xx status code
func (o *LunMapGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lun map get o k response has a 5xx status code
func (o *LunMapGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lun map get o k response a status code equal to that given
func (o *LunMapGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lun map get o k response
func (o *LunMapGetOK) Code() int {
	return 200
}

func (o *LunMapGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}][%d] lunMapGetOK %s", 200, payload)
}

func (o *LunMapGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}][%d] lunMapGetOK %s", 200, payload)
}

func (o *LunMapGetOK) GetPayload() *models.LunMap {
	return o.Payload
}

func (o *LunMapGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LunMap)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLunMapGetDefault creates a LunMapGetDefault with default headers values
func NewLunMapGetDefault(code int) *LunMapGetDefault {
	return &LunMapGetDefault{
		_statusCode: code,
	}
}

/*
	LunMapGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5374852 | The initiator group does not exist or is not accessible to the caller. |
| 5374875 | The LUN does not exist or is not accessible to the caller. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type LunMapGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this lun map get default response has a 2xx status code
func (o *LunMapGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lun map get default response has a 3xx status code
func (o *LunMapGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lun map get default response has a 4xx status code
func (o *LunMapGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lun map get default response has a 5xx status code
func (o *LunMapGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lun map get default response a status code equal to that given
func (o *LunMapGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lun map get default response
func (o *LunMapGetDefault) Code() int {
	return o._statusCode
}

func (o *LunMapGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}][%d] lun_map_get default %s", o._statusCode, payload)
}

func (o *LunMapGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}][%d] lun_map_get default %s", o._statusCode, payload)
}

func (o *LunMapGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunMapGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
