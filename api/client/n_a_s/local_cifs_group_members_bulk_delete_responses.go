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

// LocalCifsGroupMembersBulkDeleteReader is a Reader for the LocalCifsGroupMembersBulkDelete structure.
type LocalCifsGroupMembersBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalCifsGroupMembersBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalCifsGroupMembersBulkDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalCifsGroupMembersBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalCifsGroupMembersBulkDeleteOK creates a LocalCifsGroupMembersBulkDeleteOK with default headers values
func NewLocalCifsGroupMembersBulkDeleteOK() *LocalCifsGroupMembersBulkDeleteOK {
	return &LocalCifsGroupMembersBulkDeleteOK{}
}

/*
LocalCifsGroupMembersBulkDeleteOK describes a response with status code 200, with default header values.

OK
*/
type LocalCifsGroupMembersBulkDeleteOK struct {
}

// IsSuccess returns true when this local cifs group members bulk delete o k response has a 2xx status code
func (o *LocalCifsGroupMembersBulkDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local cifs group members bulk delete o k response has a 3xx status code
func (o *LocalCifsGroupMembersBulkDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local cifs group members bulk delete o k response has a 4xx status code
func (o *LocalCifsGroupMembersBulkDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this local cifs group members bulk delete o k response has a 5xx status code
func (o *LocalCifsGroupMembersBulkDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this local cifs group members bulk delete o k response a status code equal to that given
func (o *LocalCifsGroupMembersBulkDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the local cifs group members bulk delete o k response
func (o *LocalCifsGroupMembersBulkDeleteOK) Code() int {
	return 200
}

func (o *LocalCifsGroupMembersBulkDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/local-groups/{svm.uuid}/{local_cifs_group.sid}/members][%d] localCifsGroupMembersBulkDeleteOK", 200)
}

func (o *LocalCifsGroupMembersBulkDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/local-groups/{svm.uuid}/{local_cifs_group.sid}/members][%d] localCifsGroupMembersBulkDeleteOK", 200)
}

func (o *LocalCifsGroupMembersBulkDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLocalCifsGroupMembersBulkDeleteDefault creates a LocalCifsGroupMembersBulkDeleteDefault with default headers values
func NewLocalCifsGroupMembersBulkDeleteDefault(code int) *LocalCifsGroupMembersBulkDeleteDefault {
	return &LocalCifsGroupMembersBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
LocalCifsGroupMembersBulkDeleteDefault describes a response with status code -1, with default header values.

Error ONTAP Error Response Codes | Error Code | Description | | ---------- | ----------- | | 655673     | Failed to resolve the member to be deleted from the specified group. | | 655719     | Failed to delete a member from the specified group. The error code returned details the failure along with the reason for the failure. Take corrective actions as per the specified reason. | | 655742     | Records must not be specified when member name is specified. | | 655743     | SVM UUID and CIFS local group SID are invalid fields for the "records" parameter. |
*/
type LocalCifsGroupMembersBulkDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this local cifs group members bulk delete default response has a 2xx status code
func (o *LocalCifsGroupMembersBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this local cifs group members bulk delete default response has a 3xx status code
func (o *LocalCifsGroupMembersBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this local cifs group members bulk delete default response has a 4xx status code
func (o *LocalCifsGroupMembersBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this local cifs group members bulk delete default response has a 5xx status code
func (o *LocalCifsGroupMembersBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this local cifs group members bulk delete default response a status code equal to that given
func (o *LocalCifsGroupMembersBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the local cifs group members bulk delete default response
func (o *LocalCifsGroupMembersBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *LocalCifsGroupMembersBulkDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/local-groups/{svm.uuid}/{local_cifs_group.sid}/members][%d] local_cifs_group_members_bulk_delete default %s", o._statusCode, payload)
}

func (o *LocalCifsGroupMembersBulkDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/local-groups/{svm.uuid}/{local_cifs_group.sid}/members][%d] local_cifs_group_members_bulk_delete default %s", o._statusCode, payload)
}

func (o *LocalCifsGroupMembersBulkDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalCifsGroupMembersBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}