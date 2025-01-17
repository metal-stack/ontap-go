// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NodeModifyReader is a Reader for the NodeModify structure.
type NodeModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewNodeModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNodeModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodeModifyOK creates a NodeModifyOK with default headers values
func NewNodeModifyOK() *NodeModifyOK {
	return &NodeModifyOK{}
}

/*
NodeModifyOK describes a response with status code 200, with default header values.

OK
*/
type NodeModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this node modify o k response has a 2xx status code
func (o *NodeModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this node modify o k response has a 3xx status code
func (o *NodeModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node modify o k response has a 4xx status code
func (o *NodeModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this node modify o k response has a 5xx status code
func (o *NodeModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this node modify o k response a status code equal to that given
func (o *NodeModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the node modify o k response
func (o *NodeModifyOK) Code() int {
	return 200
}

func (o *NodeModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/nodes/{uuid}][%d] nodeModifyOK %s", 200, payload)
}

func (o *NodeModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/nodes/{uuid}][%d] nodeModifyOK %s", 200, payload)
}

func (o *NodeModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *NodeModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeModifyAccepted creates a NodeModifyAccepted with default headers values
func NewNodeModifyAccepted() *NodeModifyAccepted {
	return &NodeModifyAccepted{}
}

/*
NodeModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type NodeModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this node modify accepted response has a 2xx status code
func (o *NodeModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this node modify accepted response has a 3xx status code
func (o *NodeModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node modify accepted response has a 4xx status code
func (o *NodeModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this node modify accepted response has a 5xx status code
func (o *NodeModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this node modify accepted response a status code equal to that given
func (o *NodeModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the node modify accepted response
func (o *NodeModifyAccepted) Code() int {
	return 202
}

func (o *NodeModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/nodes/{uuid}][%d] nodeModifyAccepted %s", 202, payload)
}

func (o *NodeModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/nodes/{uuid}][%d] nodeModifyAccepted %s", 202, payload)
}

func (o *NodeModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *NodeModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeModifyDefault creates a NodeModifyDefault with default headers values
func NewNodeModifyDefault(code int) *NodeModifyDefault {
	return &NodeModifyDefault{
		_statusCode: code,
	}
}

/*
	NodeModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 852046 | HA partner node is not running to do takeover. |
| 852115 | The reboot/shutdown is prevented because LIFs cannot be moved away from the node. |
| 3604514 | A reboot or shutdown request is already in progress. |
| 3604515 | Reboot or shutdown of all nodes results in data service failure and client disruption for the entire cluster. Use "allow-data-outage=true" to bypass this check. |
| 9240591 | The name is not valid. The name is already in use by a cluster node, SVM, or it is the name of the local cluster. |
| 9240606 | The reboot/shutdown is prevented due to quorum warnings. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NodeModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this node modify default response has a 2xx status code
func (o *NodeModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this node modify default response has a 3xx status code
func (o *NodeModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this node modify default response has a 4xx status code
func (o *NodeModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this node modify default response has a 5xx status code
func (o *NodeModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this node modify default response a status code equal to that given
func (o *NodeModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the node modify default response
func (o *NodeModifyDefault) Code() int {
	return o._statusCode
}

func (o *NodeModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/nodes/{uuid}][%d] node_modify default %s", o._statusCode, payload)
}

func (o *NodeModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/nodes/{uuid}][%d] node_modify default %s", o._statusCode, payload)
}

func (o *NodeModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
