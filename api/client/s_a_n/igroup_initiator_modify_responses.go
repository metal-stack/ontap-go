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

// IgroupInitiatorModifyReader is a Reader for the IgroupInitiatorModify structure.
type IgroupInitiatorModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IgroupInitiatorModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIgroupInitiatorModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIgroupInitiatorModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIgroupInitiatorModifyOK creates a IgroupInitiatorModifyOK with default headers values
func NewIgroupInitiatorModifyOK() *IgroupInitiatorModifyOK {
	return &IgroupInitiatorModifyOK{}
}

/*
IgroupInitiatorModifyOK describes a response with status code 200, with default header values.

OK
*/
type IgroupInitiatorModifyOK struct {
}

// IsSuccess returns true when this igroup initiator modify o k response has a 2xx status code
func (o *IgroupInitiatorModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this igroup initiator modify o k response has a 3xx status code
func (o *IgroupInitiatorModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this igroup initiator modify o k response has a 4xx status code
func (o *IgroupInitiatorModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this igroup initiator modify o k response has a 5xx status code
func (o *IgroupInitiatorModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this igroup initiator modify o k response a status code equal to that given
func (o *IgroupInitiatorModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the igroup initiator modify o k response
func (o *IgroupInitiatorModifyOK) Code() int {
	return 200
}

func (o *IgroupInitiatorModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/san/igroups/{igroup.uuid}/initiators/{name}][%d] igroupInitiatorModifyOK", 200)
}

func (o *IgroupInitiatorModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/san/igroups/{igroup.uuid}/initiators/{name}][%d] igroupInitiatorModifyOK", 200)
}

func (o *IgroupInitiatorModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIgroupInitiatorModifyDefault creates a IgroupInitiatorModifyDefault with default headers values
func NewIgroupInitiatorModifyDefault(code int) *IgroupInitiatorModifyDefault {
	return &IgroupInitiatorModifyDefault{
		_statusCode: code,
	}
}

/*
	IgroupInitiatorModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5374034 | An initiator is not a member of the initiator group. |
| 5374744 | The cluster is currently running in a mixed version and the initiators cannot be modified until the effective cluster version reaches 9.9.1. |
| 5374852 | The initiator group does not exist. |
| 5374918 | A subset of the provided list of initiators were modified before a failure occurred. |
| 5375055 | The `local_svm` property of an initiator proximity was not specified. |
| 5375056 | An SVM peering relationship that does not have the initiator group's SVM as the local SVM was specified. |
| 5375261 | Setting initiator proximity is not supported for the SVM type. |
| 5376057 | Setting initiator proximity is not supported for the ONTAP version. |
| 5376059 | Setting initiator proximity to a peer that is either the destination of an SVM DR relationship or in a Metrocluster configuration is not supported. |
| 26345672 | The specified SVM peering relationship was not found. |
| 26345673 | An SVM peering relationship between the initiator group's SVM and specified peer SVM was not found. |
| 26345675 | An SVM peering relationship UUID and name were specified and they do not refer to the same SVM peering relationship. |
| 26345680 | Supplied SVM peer is on the local cluster. The operation requires a peer on a remote cluster. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type IgroupInitiatorModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this igroup initiator modify default response has a 2xx status code
func (o *IgroupInitiatorModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this igroup initiator modify default response has a 3xx status code
func (o *IgroupInitiatorModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this igroup initiator modify default response has a 4xx status code
func (o *IgroupInitiatorModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this igroup initiator modify default response has a 5xx status code
func (o *IgroupInitiatorModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this igroup initiator modify default response a status code equal to that given
func (o *IgroupInitiatorModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the igroup initiator modify default response
func (o *IgroupInitiatorModifyDefault) Code() int {
	return o._statusCode
}

func (o *IgroupInitiatorModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/san/igroups/{igroup.uuid}/initiators/{name}][%d] igroup_initiator_modify default %s", o._statusCode, payload)
}

func (o *IgroupInitiatorModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/san/igroups/{igroup.uuid}/initiators/{name}][%d] igroup_initiator_modify default %s", o._statusCode, payload)
}

func (o *IgroupInitiatorModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IgroupInitiatorModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
