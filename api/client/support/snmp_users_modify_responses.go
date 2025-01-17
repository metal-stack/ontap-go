// Code generated by go-swagger; DO NOT EDIT.

package support

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

// SnmpUsersModifyReader is a Reader for the SnmpUsersModify structure.
type SnmpUsersModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnmpUsersModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnmpUsersModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnmpUsersModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnmpUsersModifyOK creates a SnmpUsersModifyOK with default headers values
func NewSnmpUsersModifyOK() *SnmpUsersModifyOK {
	return &SnmpUsersModifyOK{}
}

/*
SnmpUsersModifyOK describes a response with status code 200, with default header values.

OK
*/
type SnmpUsersModifyOK struct {
}

// IsSuccess returns true when this snmp users modify o k response has a 2xx status code
func (o *SnmpUsersModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snmp users modify o k response has a 3xx status code
func (o *SnmpUsersModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snmp users modify o k response has a 4xx status code
func (o *SnmpUsersModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snmp users modify o k response has a 5xx status code
func (o *SnmpUsersModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snmp users modify o k response a status code equal to that given
func (o *SnmpUsersModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snmp users modify o k response
func (o *SnmpUsersModifyOK) Code() int {
	return 200
}

func (o *SnmpUsersModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /support/snmp/users/{engine_id}/{name}][%d] snmpUsersModifyOK", 200)
}

func (o *SnmpUsersModifyOK) String() string {
	return fmt.Sprintf("[PATCH /support/snmp/users/{engine_id}/{name}][%d] snmpUsersModifyOK", 200)
}

func (o *SnmpUsersModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSnmpUsersModifyDefault creates a SnmpUsersModifyDefault with default headers values
func NewSnmpUsersModifyDefault(code int) *SnmpUsersModifyDefault {
	return &SnmpUsersModifyDefault{
		_statusCode: code,
	}
}

/*
	SnmpUsersModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621475 | This operation is not allowed on a node SVM. |
| 2621699 | This operation is not allowed on a system SVM. |
| 5636123 | Cannot create an SNMP user with a role other than readonly, none, or admin. |
| 5636124 | Cannot create an SNMP user with a role other than vsadmin-readonly, none, or vsadmin. |
| 5832712 | Cannot modify attributes for user \"diag.\" |
| 7077906 | Cannot use given role with this SVM because a role with that name has not been defined for the SVM. |
| 9043999 | ONTAP failed to create an SNMPv3 user because SNMPv3 is disabled in the cluster. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SnmpUsersModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snmp users modify default response has a 2xx status code
func (o *SnmpUsersModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snmp users modify default response has a 3xx status code
func (o *SnmpUsersModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snmp users modify default response has a 4xx status code
func (o *SnmpUsersModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snmp users modify default response has a 5xx status code
func (o *SnmpUsersModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snmp users modify default response a status code equal to that given
func (o *SnmpUsersModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snmp users modify default response
func (o *SnmpUsersModifyDefault) Code() int {
	return o._statusCode
}

func (o *SnmpUsersModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /support/snmp/users/{engine_id}/{name}][%d] snmp_users_modify default %s", o._statusCode, payload)
}

func (o *SnmpUsersModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /support/snmp/users/{engine_id}/{name}][%d] snmp_users_modify default %s", o._statusCode, payload)
}

func (o *SnmpUsersModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnmpUsersModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
