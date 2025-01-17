// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

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

// NvmeSubsystemControllerGetReader is a Reader for the NvmeSubsystemControllerGet structure.
type NvmeSubsystemControllerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeSubsystemControllerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeSubsystemControllerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeSubsystemControllerGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeSubsystemControllerGetOK creates a NvmeSubsystemControllerGetOK with default headers values
func NewNvmeSubsystemControllerGetOK() *NvmeSubsystemControllerGetOK {
	return &NvmeSubsystemControllerGetOK{}
}

/*
NvmeSubsystemControllerGetOK describes a response with status code 200, with default header values.

OK
*/
type NvmeSubsystemControllerGetOK struct {
	Payload *models.NvmeSubsystemController
}

// IsSuccess returns true when this nvme subsystem controller get o k response has a 2xx status code
func (o *NvmeSubsystemControllerGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme subsystem controller get o k response has a 3xx status code
func (o *NvmeSubsystemControllerGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme subsystem controller get o k response has a 4xx status code
func (o *NvmeSubsystemControllerGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme subsystem controller get o k response has a 5xx status code
func (o *NvmeSubsystemControllerGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme subsystem controller get o k response a status code equal to that given
func (o *NvmeSubsystemControllerGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nvme subsystem controller get o k response
func (o *NvmeSubsystemControllerGetOK) Code() int {
	return 200
}

func (o *NvmeSubsystemControllerGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystem-controllers/{subsystem.uuid}/{id}][%d] nvmeSubsystemControllerGetOK %s", 200, payload)
}

func (o *NvmeSubsystemControllerGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystem-controllers/{subsystem.uuid}/{id}][%d] nvmeSubsystemControllerGetOK %s", 200, payload)
}

func (o *NvmeSubsystemControllerGetOK) GetPayload() *models.NvmeSubsystemController {
	return o.Payload
}

func (o *NvmeSubsystemControllerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmeSubsystemController)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeSubsystemControllerGetDefault creates a NvmeSubsystemControllerGetDefault with default headers values
func NewNvmeSubsystemControllerGetDefault(code int) *NvmeSubsystemControllerGetDefault {
	return &NvmeSubsystemControllerGetDefault{
		_statusCode: code,
	}
}

/*
	NvmeSubsystemControllerGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 72090001 | The supplied subsystem identifier does not exist. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NvmeSubsystemControllerGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme subsystem controller get default response has a 2xx status code
func (o *NvmeSubsystemControllerGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme subsystem controller get default response has a 3xx status code
func (o *NvmeSubsystemControllerGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme subsystem controller get default response has a 4xx status code
func (o *NvmeSubsystemControllerGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme subsystem controller get default response has a 5xx status code
func (o *NvmeSubsystemControllerGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme subsystem controller get default response a status code equal to that given
func (o *NvmeSubsystemControllerGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme subsystem controller get default response
func (o *NvmeSubsystemControllerGetDefault) Code() int {
	return o._statusCode
}

func (o *NvmeSubsystemControllerGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystem-controllers/{subsystem.uuid}/{id}][%d] nvme_subsystem_controller_get default %s", o._statusCode, payload)
}

func (o *NvmeSubsystemControllerGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystem-controllers/{subsystem.uuid}/{id}][%d] nvme_subsystem_controller_get default %s", o._statusCode, payload)
}

func (o *NvmeSubsystemControllerGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeSubsystemControllerGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}