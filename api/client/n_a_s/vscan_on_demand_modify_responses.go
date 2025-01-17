// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// VscanOnDemandModifyReader is a Reader for the VscanOnDemandModify structure.
type VscanOnDemandModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanOnDemandModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanOnDemandModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanOnDemandModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanOnDemandModifyOK creates a VscanOnDemandModifyOK with default headers values
func NewVscanOnDemandModifyOK() *VscanOnDemandModifyOK {
	return &VscanOnDemandModifyOK{}
}

/*
VscanOnDemandModifyOK describes a response with status code 200, with default header values.

OK
*/
type VscanOnDemandModifyOK struct {
}

// IsSuccess returns true when this vscan on demand modify o k response has a 2xx status code
func (o *VscanOnDemandModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan on demand modify o k response has a 3xx status code
func (o *VscanOnDemandModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan on demand modify o k response has a 4xx status code
func (o *VscanOnDemandModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan on demand modify o k response has a 5xx status code
func (o *VscanOnDemandModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan on demand modify o k response a status code equal to that given
func (o *VscanOnDemandModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vscan on demand modify o k response
func (o *VscanOnDemandModifyOK) Code() int {
	return 200
}

func (o *VscanOnDemandModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/vscan/{svm.uuid}/on-demand-policies/{name}][%d] vscanOnDemandModifyOK", 200)
}

func (o *VscanOnDemandModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/vscan/{svm.uuid}/on-demand-policies/{name}][%d] vscanOnDemandModifyOK", 200)
}

func (o *VscanOnDemandModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVscanOnDemandModifyDefault creates a VscanOnDemandModifyDefault with default headers values
func NewVscanOnDemandModifyDefault(code int) *VscanOnDemandModifyDefault {
	return &VscanOnDemandModifyDefault{
		_statusCode: code,
	}
}

/*
	VscanOnDemandModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 10027101   | The file size must be in the range 1KB to 1TB |
| 10027107   | The include extensions list cannot be empty. Specify at least one extension for inclusion. |
| 10027164   | An On-Demand policy cannot be scheduled, as the Vscan is disabled. Enable the Vscan and retry the operation. |
| 10027167   | The specified schedule does not exist. Create the schedule or create a policy without specifying the schedule. |
| 10027168   | The specified scan path does not exist. The scan path must be specified from the root of the SVM, and must begin with UNIX path delimiters (use “/” not “\\”) |
| 10027169   | The specified scan path is not supported for scanning. |
| 10027174   | The specified exclude path is invalid. The path must be specified from the root of the SVM, and must begin with UNIX path delimiters (use "/" not "\\") |
| 10027175   | An On-Demand policy cannot be scheduled as the SVM is not in an operational state. |
| 10027176   | The log-path specified does not exist. The log path must be specified from the root of the SVM, and must begin with UNIX path delimiters (use “/” not “\\”) |
| 10027177   | The log path specified is not supported. |
| 10027253   | The number of paths specified exceeds the configured maximum number of paths. You cannot specify more than the maximum number of configured paths. |
| 10027254   | The number of extensions specified exceeds the configured maximum number of extensions. You cannot specify more than the maximum number of configured extensions. |
| 10027255   | Another policy is already scheduled. Only one policy per SVM is allowed to be scheduled at any one time. Update a policy without specifying a schedule. |
*/
type VscanOnDemandModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vscan on demand modify default response has a 2xx status code
func (o *VscanOnDemandModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan on demand modify default response has a 3xx status code
func (o *VscanOnDemandModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan on demand modify default response has a 4xx status code
func (o *VscanOnDemandModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan on demand modify default response has a 5xx status code
func (o *VscanOnDemandModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan on demand modify default response a status code equal to that given
func (o *VscanOnDemandModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vscan on demand modify default response
func (o *VscanOnDemandModifyDefault) Code() int {
	return o._statusCode
}

func (o *VscanOnDemandModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/vscan/{svm.uuid}/on-demand-policies/{name}][%d] vscan_on_demand_modify default %s", o._statusCode, payload)
}

func (o *VscanOnDemandModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/vscan/{svm.uuid}/on-demand-policies/{name}][%d] vscan_on_demand_modify default %s", o._statusCode, payload)
}

func (o *VscanOnDemandModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanOnDemandModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
