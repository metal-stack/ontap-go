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

// NvmeSubsystemCreateReader is a Reader for the NvmeSubsystemCreate structure.
type NvmeSubsystemCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeSubsystemCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNvmeSubsystemCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeSubsystemCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeSubsystemCreateCreated creates a NvmeSubsystemCreateCreated with default headers values
func NewNvmeSubsystemCreateCreated() *NvmeSubsystemCreateCreated {
	return &NvmeSubsystemCreateCreated{}
}

/*
NvmeSubsystemCreateCreated describes a response with status code 201, with default header values.

Created
*/
type NvmeSubsystemCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.NvmeSubsystemResponse
}

// IsSuccess returns true when this nvme subsystem create created response has a 2xx status code
func (o *NvmeSubsystemCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme subsystem create created response has a 3xx status code
func (o *NvmeSubsystemCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme subsystem create created response has a 4xx status code
func (o *NvmeSubsystemCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme subsystem create created response has a 5xx status code
func (o *NvmeSubsystemCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme subsystem create created response a status code equal to that given
func (o *NvmeSubsystemCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the nvme subsystem create created response
func (o *NvmeSubsystemCreateCreated) Code() int {
	return 201
}

func (o *NvmeSubsystemCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nvme/subsystems][%d] nvmeSubsystemCreateCreated %s", 201, payload)
}

func (o *NvmeSubsystemCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nvme/subsystems][%d] nvmeSubsystemCreateCreated %s", 201, payload)
}

func (o *NvmeSubsystemCreateCreated) GetPayload() *models.NvmeSubsystemResponse {
	return o.Payload
}

func (o *NvmeSubsystemCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.NvmeSubsystemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeSubsystemCreateDefault creates a NvmeSubsystemCreateDefault with default headers values
func NewNvmeSubsystemCreateDefault(code int) *NvmeSubsystemCreateDefault {
	return &NvmeSubsystemCreateDefault{
		_statusCode: code,
	}
}

/*
	NvmeSubsystemCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621462 | The supplied SVM does not exist. |
| 2621706 | The specified `svm.uuid` and `svm.name` do not refer to the same SVM. |
| 2621707 | The svm.uuid or svm.name must be provided. |
| 72089635 | Setting vendor-specific UUIDs on NVMe subsystems is not supported until the effective cluster version is 9.9 or later. |
| 72089636 | Creating NVMe subsystems with `os_type` AIX is not supported until the effective cluster version is 9.13.1 or later. |
| 72089709 | The NVMe subsystem name contains an invalid character. |
| 72089711 | An invalid vendor-specific UUID was specified. |
| 72089712 | A duplicate vendor-specific UUID was specific. |
| 72089713 | Too many vendor UUIDs were supplied. |
| 72089771 | The NQN is invalid. A non-empty qualifier is required after the prefix. An example of a valid NQN is _nqn.1992-01.com.example:string_. |
| 72089772 | The NQN is invalid. Add the prefix _'nqn'_. An example of a valid NQN is _nqn.1992-01.com.example:string_. |
| 72089773 | The NQN is invalid. The date field must be formatted _yyyy-mm_. An example of a valid NQN is _nqn.1992-01.com.example:string_. |
| 72090003 | A host to be added to an NVMe subsystem is missing the "nqn" property. |
| 72090025 | The NVMe subsystem already exists for the SVM. |
| 72090029 | The NVMe service does not exist. |
| 72090030 | A partial success occurred while adding multiple NVMe subsystem hosts to an NVMe subsystem. |
| 72090036 | An NVMe subsystem host NQN was duplicated in the input. |
| 72090042 | The `dh_hmac_chap.host_secret_key` property is required when setting any other NVMe in-band authentication properties for a host. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NvmeSubsystemCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme subsystem create default response has a 2xx status code
func (o *NvmeSubsystemCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme subsystem create default response has a 3xx status code
func (o *NvmeSubsystemCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme subsystem create default response has a 4xx status code
func (o *NvmeSubsystemCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme subsystem create default response has a 5xx status code
func (o *NvmeSubsystemCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme subsystem create default response a status code equal to that given
func (o *NvmeSubsystemCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme subsystem create default response
func (o *NvmeSubsystemCreateDefault) Code() int {
	return o._statusCode
}

func (o *NvmeSubsystemCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nvme/subsystems][%d] nvme_subsystem_create default %s", o._statusCode, payload)
}

func (o *NvmeSubsystemCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nvme/subsystems][%d] nvme_subsystem_create default %s", o._statusCode, payload)
}

func (o *NvmeSubsystemCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeSubsystemCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
