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

// VvolBindingDeleteReader is a Reader for the VvolBindingDelete structure.
type VvolBindingDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VvolBindingDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVvolBindingDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVvolBindingDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVvolBindingDeleteOK creates a VvolBindingDeleteOK with default headers values
func NewVvolBindingDeleteOK() *VvolBindingDeleteOK {
	return &VvolBindingDeleteOK{}
}

/*
VvolBindingDeleteOK describes a response with status code 200, with default header values.

OK
*/
type VvolBindingDeleteOK struct {
}

// IsSuccess returns true when this vvol binding delete o k response has a 2xx status code
func (o *VvolBindingDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vvol binding delete o k response has a 3xx status code
func (o *VvolBindingDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vvol binding delete o k response has a 4xx status code
func (o *VvolBindingDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vvol binding delete o k response has a 5xx status code
func (o *VvolBindingDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vvol binding delete o k response a status code equal to that given
func (o *VvolBindingDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vvol binding delete o k response
func (o *VvolBindingDeleteOK) Code() int {
	return 200
}

func (o *VvolBindingDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/san/vvol-bindings/{protocol_endpoint.uuid}/{vvol.uuid}][%d] vvolBindingDeleteOK", 200)
}

func (o *VvolBindingDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/san/vvol-bindings/{protocol_endpoint.uuid}/{vvol.uuid}][%d] vvolBindingDeleteOK", 200)
}

func (o *VvolBindingDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVvolBindingDeleteDefault creates a VvolBindingDeleteDefault with default headers values
func NewVvolBindingDeleteDefault(code int) *VvolBindingDeleteDefault {
	return &VvolBindingDeleteDefault{
		_statusCode: code,
	}
}

/*
	VvolBindingDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5374875 | The vVol binding was not found because the protocol endpoint or vVol LUN was not found. Use to the `target` property of the error object to differentiate between the protocol endpoint LUN and the vVol LUN. |
| 5374926 | The vVol binding was not found. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type VvolBindingDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vvol binding delete default response has a 2xx status code
func (o *VvolBindingDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vvol binding delete default response has a 3xx status code
func (o *VvolBindingDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vvol binding delete default response has a 4xx status code
func (o *VvolBindingDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vvol binding delete default response has a 5xx status code
func (o *VvolBindingDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vvol binding delete default response a status code equal to that given
func (o *VvolBindingDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vvol binding delete default response
func (o *VvolBindingDeleteDefault) Code() int {
	return o._statusCode
}

func (o *VvolBindingDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/san/vvol-bindings/{protocol_endpoint.uuid}/{vvol.uuid}][%d] vvol_binding_delete default %s", o._statusCode, payload)
}

func (o *VvolBindingDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/san/vvol-bindings/{protocol_endpoint.uuid}/{vvol.uuid}][%d] vvol_binding_delete default %s", o._statusCode, payload)
}

func (o *VvolBindingDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VvolBindingDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
