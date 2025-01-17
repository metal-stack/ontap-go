// Code generated by go-swagger; DO NOT EDIT.

package security

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

// AccountCreateReader is a Reader for the AccountCreate structure.
type AccountCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAccountCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAccountCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountCreateCreated creates a AccountCreateCreated with default headers values
func NewAccountCreateCreated() *AccountCreateCreated {
	return &AccountCreateCreated{}
}

/*
AccountCreateCreated describes a response with status code 201, with default header values.

Created
*/
type AccountCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this account create created response has a 2xx status code
func (o *AccountCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account create created response has a 3xx status code
func (o *AccountCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account create created response has a 4xx status code
func (o *AccountCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this account create created response has a 5xx status code
func (o *AccountCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this account create created response a status code equal to that given
func (o *AccountCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the account create created response
func (o *AccountCreateCreated) Code() int {
	return 201
}

func (o *AccountCreateCreated) Error() string {
	return fmt.Sprintf("[POST /security/accounts][%d] accountCreateCreated", 201)
}

func (o *AccountCreateCreated) String() string {
	return fmt.Sprintf("[POST /security/accounts][%d] accountCreateCreated", 201)
}

func (o *AccountCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewAccountCreateDefault creates a AccountCreateDefault with default headers values
func NewAccountCreateDefault(code int) *AccountCreateDefault {
	return &AccountCreateDefault{
		_statusCode: code,
	}
}

/*
	AccountCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1261215 | The role was not found. |
| 1261225 | Invalid command directory name. |
| 1263343 | Cannot lock user with password not set or non-password authentication method. |
| 2621475 | This operation is not supported on a node SVM. |
| 2621601 | This operation is not supported on a system SVM. |
| 2621706 | The specified owner.uuid and owner.name refer to different SVMs. |
| 5636099 | User creation with a non-admin role is not supported for service-processor application. |
| 5636121 | The user account name is reserved for use by the system. |
| 5636126 | Cannot create a user with the username or role as AutoSupport because it is reserved by the system. |
| 5636136 | Specifying "is_ns_switch_group" as "true" is supported only for authentication method "nsswitch". |
| 5636140 | Creating a login with application console for a data SVM  is not supported. |
| 5636141 | Creating a login with application service-processor for a data SVM is not supported. |
| 5636154 | The second authentication method parameter is supported for SSH and Service Processor (SP) applications only. |
| 5636155 | The second-authentication-method parameter can be specified only if the authentication-method password or public key nsswitch. |
| 5636156 | The same value cannot be specified for the second-authentication-method and the authentication-method. |
| 5636164 | If the value for either the authentication-method second-authentication-method is nsswitch or password, the other parameter must differ. |
| 5636165 | Second authentication method is not supported for NIS or LDAP group based accounts. |
| 5636176 | The application and authentication-method combination is invalid. |
| 5636178 | An invalid value is specified for field "application". |
| 5636179 | Creating an AMQP application login for a data SVM is not supported. |
| 5636197 | LDAP fastbind combination for application and authentication method is not supported. |
| 5636198 | LDAP fastbind authentication is supported only for nsswitch. |
| 5636206 | Non-domain user cannot have a backslash in the username. |
| 5636207 | If the value for either the authentication-method or second-authentication-method parameters is domain, the other parameter must be publickey or none. |
| 5636212 | TOTP is supported only when the primary authentication method is password or public key. |
| 5636214 | Configuring the user with TOTP as secondary authentication method requires an effective cluster version of 9.13.1 or later |
| 5636223 | Specifying "is_ns_switch_group" as "true" is supported only for SSH, ONTAPI and HTTP applications. |
| 5636224 | Configuring a Service Processor (SP) user with two-factor authentication requires an effective cluster version of 9.15.1 or later. |
| 5636225 | For a Service Processor (SP) user, the second factor of authentication must be one of publickey or none. |
| 5636226 | Internal error. Failed to check for ONTAP capability. |
| 7077897 | Invalid character in username. |
| 7077898 | The username must contain both letters and numbers. |
| 7077899 | The username does not meet length requirements. |
| 7077906 | A role with that name has not been defined for the Vserver. |
| 7077918 | The password cannot contain the username. |
| 7077919 | The minimum length for new password does not meet the policy. |
| 7077920 | A new password must have both letters and numbers. |
| 7077921 | The minimum number of special characters required do not meet the policy. |
| 7077929 | Cannot lock user with password not set or non-password authentication method. |
| 7077940 | The password exceeds the maximum supported length. |
| 7077941 | The defined password composition exceeds the maximum password length of 128 characters. |
| 7078900 | An admin password is not set. Set the password by including it in the request. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AccountCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this account create default response has a 2xx status code
func (o *AccountCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this account create default response has a 3xx status code
func (o *AccountCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this account create default response has a 4xx status code
func (o *AccountCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this account create default response has a 5xx status code
func (o *AccountCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this account create default response a status code equal to that given
func (o *AccountCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the account create default response
func (o *AccountCreateDefault) Code() int {
	return o._statusCode
}

func (o *AccountCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/accounts][%d] account_create default %s", o._statusCode, payload)
}

func (o *AccountCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/accounts][%d] account_create default %s", o._statusCode, payload)
}

func (o *AccountCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AccountCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}