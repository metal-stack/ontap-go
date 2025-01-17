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

// NodesCreateReader is a Reader for the NodesCreate structure.
type NodesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNodesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewNodesCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNodesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodesCreateCreated creates a NodesCreateCreated with default headers values
func NewNodesCreateCreated() *NodesCreateCreated {
	return &NodesCreateCreated{}
}

/*
NodesCreateCreated describes a response with status code 201, with default header values.

Created
*/
type NodesCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this nodes create created response has a 2xx status code
func (o *NodesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nodes create created response has a 3xx status code
func (o *NodesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes create created response has a 4xx status code
func (o *NodesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this nodes create created response has a 5xx status code
func (o *NodesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this nodes create created response a status code equal to that given
func (o *NodesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the nodes create created response
func (o *NodesCreateCreated) Code() int {
	return 201
}

func (o *NodesCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/nodes][%d] nodesCreateCreated %s", 201, payload)
}

func (o *NodesCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/nodes][%d] nodesCreateCreated %s", 201, payload)
}

func (o *NodesCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *NodesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesCreateAccepted creates a NodesCreateAccepted with default headers values
func NewNodesCreateAccepted() *NodesCreateAccepted {
	return &NodesCreateAccepted{}
}

/*
NodesCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type NodesCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this nodes create accepted response has a 2xx status code
func (o *NodesCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nodes create accepted response has a 3xx status code
func (o *NodesCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes create accepted response has a 4xx status code
func (o *NodesCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this nodes create accepted response has a 5xx status code
func (o *NodesCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this nodes create accepted response a status code equal to that given
func (o *NodesCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the nodes create accepted response
func (o *NodesCreateAccepted) Code() int {
	return 202
}

func (o *NodesCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/nodes][%d] nodesCreateAccepted %s", 202, payload)
}

func (o *NodesCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/nodes][%d] nodesCreateAccepted %s", 202, payload)
}

func (o *NodesCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *NodesCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesCreateDefault creates a NodesCreateDefault with default headers values
func NewNodesCreateDefault(code int) *NodesCreateDefault {
	return &NodesCreateDefault{
		_statusCode: code,
	}
}

/*
	NodesCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262245 | The value provided was invalid. |
| 1179795 | A node being added is already in the cluster. |
| 1179813 | Fields set for one node must be set for all nodes. |
| 1179817 | The IP address, subnet mask, and gateway must all be provided for cluster manangement interface. |
| 1179818 | The IP address and gateway must be of the same family. |
| 1179821 | An IP address and subnet mask conflicts with an existing entry. |
| 9240591 | The name is not valid. The name is already in use by a cluster node, SVM, or it is the name of the local cluster. |
| 9240594 | An invalid name was provided. |
| 11862025 | The specified node management address and existing node management address belong to two subnets. Create a node management interface after creating the cluster or provide cluster and node management interfaces from the same subnet as the cluster management interface. |
| 131727360 | A node cannot be added to the cluster. This is a generic code, see response message for details. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NodesCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nodes create default response has a 2xx status code
func (o *NodesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nodes create default response has a 3xx status code
func (o *NodesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nodes create default response has a 4xx status code
func (o *NodesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nodes create default response has a 5xx status code
func (o *NodesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nodes create default response a status code equal to that given
func (o *NodesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nodes create default response
func (o *NodesCreateDefault) Code() int {
	return o._statusCode
}

func (o *NodesCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/nodes][%d] nodes_create default %s", o._statusCode, payload)
}

func (o *NodesCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/nodes][%d] nodes_create default %s", o._statusCode, payload)
}

func (o *NodesCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}