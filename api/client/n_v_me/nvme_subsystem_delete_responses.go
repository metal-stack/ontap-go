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

// NvmeSubsystemDeleteReader is a Reader for the NvmeSubsystemDelete structure.
type NvmeSubsystemDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeSubsystemDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeSubsystemDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeSubsystemDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeSubsystemDeleteOK creates a NvmeSubsystemDeleteOK with default headers values
func NewNvmeSubsystemDeleteOK() *NvmeSubsystemDeleteOK {
	return &NvmeSubsystemDeleteOK{}
}

/*
NvmeSubsystemDeleteOK describes a response with status code 200, with default header values.

OK
*/
type NvmeSubsystemDeleteOK struct {
}

// IsSuccess returns true when this nvme subsystem delete o k response has a 2xx status code
func (o *NvmeSubsystemDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme subsystem delete o k response has a 3xx status code
func (o *NvmeSubsystemDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme subsystem delete o k response has a 4xx status code
func (o *NvmeSubsystemDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme subsystem delete o k response has a 5xx status code
func (o *NvmeSubsystemDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme subsystem delete o k response a status code equal to that given
func (o *NvmeSubsystemDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nvme subsystem delete o k response
func (o *NvmeSubsystemDeleteOK) Code() int {
	return 200
}

func (o *NvmeSubsystemDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/nvme/subsystems/{uuid}][%d] nvmeSubsystemDeleteOK", 200)
}

func (o *NvmeSubsystemDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/nvme/subsystems/{uuid}][%d] nvmeSubsystemDeleteOK", 200)
}

func (o *NvmeSubsystemDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNvmeSubsystemDeleteDefault creates a NvmeSubsystemDeleteDefault with default headers values
func NewNvmeSubsystemDeleteDefault(code int) *NvmeSubsystemDeleteDefault {
	return &NvmeSubsystemDeleteDefault{
		_statusCode: code,
	}
}

/*
	NvmeSubsystemDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 72090001 | The NVMe subsystem does not exist. |
| 72090023 | The NVMe subsystem contains one or more mapped namespaces. Use the `allow_delete_while_mapped` query parameter to delete an NVMe subsystem with mapped NVMe namespaces. |
| 72090024 | The NVMe subsystem contains one or more NVMe hosts. Use the `allow_delete_with_hosts` query parameter to delete an NVMe subsystem with NVMe hosts. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NvmeSubsystemDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme subsystem delete default response has a 2xx status code
func (o *NvmeSubsystemDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme subsystem delete default response has a 3xx status code
func (o *NvmeSubsystemDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme subsystem delete default response has a 4xx status code
func (o *NvmeSubsystemDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme subsystem delete default response has a 5xx status code
func (o *NvmeSubsystemDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme subsystem delete default response a status code equal to that given
func (o *NvmeSubsystemDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme subsystem delete default response
func (o *NvmeSubsystemDeleteDefault) Code() int {
	return o._statusCode
}

func (o *NvmeSubsystemDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/nvme/subsystems/{uuid}][%d] nvme_subsystem_delete default %s", o._statusCode, payload)
}

func (o *NvmeSubsystemDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/nvme/subsystems/{uuid}][%d] nvme_subsystem_delete default %s", o._statusCode, payload)
}

func (o *NvmeSubsystemDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeSubsystemDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
