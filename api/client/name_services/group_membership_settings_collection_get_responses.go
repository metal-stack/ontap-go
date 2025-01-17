// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// GroupMembershipSettingsCollectionGetReader is a Reader for the GroupMembershipSettingsCollectionGet structure.
type GroupMembershipSettingsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupMembershipSettingsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupMembershipSettingsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupMembershipSettingsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupMembershipSettingsCollectionGetOK creates a GroupMembershipSettingsCollectionGetOK with default headers values
func NewGroupMembershipSettingsCollectionGetOK() *GroupMembershipSettingsCollectionGetOK {
	return &GroupMembershipSettingsCollectionGetOK{}
}

/*
GroupMembershipSettingsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type GroupMembershipSettingsCollectionGetOK struct {
	Payload *models.GroupMembershipSettingsResponse
}

// IsSuccess returns true when this group membership settings collection get o k response has a 2xx status code
func (o *GroupMembershipSettingsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this group membership settings collection get o k response has a 3xx status code
func (o *GroupMembershipSettingsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group membership settings collection get o k response has a 4xx status code
func (o *GroupMembershipSettingsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this group membership settings collection get o k response has a 5xx status code
func (o *GroupMembershipSettingsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this group membership settings collection get o k response a status code equal to that given
func (o *GroupMembershipSettingsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the group membership settings collection get o k response
func (o *GroupMembershipSettingsCollectionGetOK) Code() int {
	return 200
}

func (o *GroupMembershipSettingsCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/cache/group-membership/settings][%d] groupMembershipSettingsCollectionGetOK %s", 200, payload)
}

func (o *GroupMembershipSettingsCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/cache/group-membership/settings][%d] groupMembershipSettingsCollectionGetOK %s", 200, payload)
}

func (o *GroupMembershipSettingsCollectionGetOK) GetPayload() *models.GroupMembershipSettingsResponse {
	return o.Payload
}

func (o *GroupMembershipSettingsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupMembershipSettingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupMembershipSettingsCollectionGetDefault creates a GroupMembershipSettingsCollectionGetDefault with default headers values
func NewGroupMembershipSettingsCollectionGetDefault(code int) *GroupMembershipSettingsCollectionGetDefault {
	return &GroupMembershipSettingsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
GroupMembershipSettingsCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type GroupMembershipSettingsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this group membership settings collection get default response has a 2xx status code
func (o *GroupMembershipSettingsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this group membership settings collection get default response has a 3xx status code
func (o *GroupMembershipSettingsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this group membership settings collection get default response has a 4xx status code
func (o *GroupMembershipSettingsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this group membership settings collection get default response has a 5xx status code
func (o *GroupMembershipSettingsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this group membership settings collection get default response a status code equal to that given
func (o *GroupMembershipSettingsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the group membership settings collection get default response
func (o *GroupMembershipSettingsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *GroupMembershipSettingsCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/cache/group-membership/settings][%d] group_membership_settings_collection_get default %s", o._statusCode, payload)
}

func (o *GroupMembershipSettingsCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /name-services/cache/group-membership/settings][%d] group_membership_settings_collection_get default %s", o._statusCode, payload)
}

func (o *GroupMembershipSettingsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GroupMembershipSettingsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
