// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// QuotaRuleDeleteReader is a Reader for the QuotaRuleDelete structure.
type QuotaRuleDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuotaRuleDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuotaRuleDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewQuotaRuleDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQuotaRuleDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuotaRuleDeleteOK creates a QuotaRuleDeleteOK with default headers values
func NewQuotaRuleDeleteOK() *QuotaRuleDeleteOK {
	return &QuotaRuleDeleteOK{}
}

/*
QuotaRuleDeleteOK describes a response with status code 200, with default header values.

OK
*/
type QuotaRuleDeleteOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this quota rule delete o k response has a 2xx status code
func (o *QuotaRuleDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this quota rule delete o k response has a 3xx status code
func (o *QuotaRuleDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this quota rule delete o k response has a 4xx status code
func (o *QuotaRuleDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this quota rule delete o k response has a 5xx status code
func (o *QuotaRuleDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this quota rule delete o k response a status code equal to that given
func (o *QuotaRuleDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the quota rule delete o k response
func (o *QuotaRuleDeleteOK) Code() int {
	return 200
}

func (o *QuotaRuleDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/quota/rules/{uuid}][%d] quotaRuleDeleteOK %s", 200, payload)
}

func (o *QuotaRuleDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/quota/rules/{uuid}][%d] quotaRuleDeleteOK %s", 200, payload)
}

func (o *QuotaRuleDeleteOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *QuotaRuleDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotaRuleDeleteAccepted creates a QuotaRuleDeleteAccepted with default headers values
func NewQuotaRuleDeleteAccepted() *QuotaRuleDeleteAccepted {
	return &QuotaRuleDeleteAccepted{}
}

/*
QuotaRuleDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type QuotaRuleDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this quota rule delete accepted response has a 2xx status code
func (o *QuotaRuleDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this quota rule delete accepted response has a 3xx status code
func (o *QuotaRuleDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this quota rule delete accepted response has a 4xx status code
func (o *QuotaRuleDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this quota rule delete accepted response has a 5xx status code
func (o *QuotaRuleDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this quota rule delete accepted response a status code equal to that given
func (o *QuotaRuleDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the quota rule delete accepted response
func (o *QuotaRuleDeleteAccepted) Code() int {
	return 202
}

func (o *QuotaRuleDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/quota/rules/{uuid}][%d] quotaRuleDeleteAccepted %s", 202, payload)
}

func (o *QuotaRuleDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/quota/rules/{uuid}][%d] quotaRuleDeleteAccepted %s", 202, payload)
}

func (o *QuotaRuleDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *QuotaRuleDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotaRuleDeleteDefault creates a QuotaRuleDeleteDefault with default headers values
func NewQuotaRuleDeleteDefault(code int) *QuotaRuleDeleteDefault {
	return &QuotaRuleDeleteDefault{
		_statusCode: code,
	}
}

/*
	QuotaRuleDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5308545 | The specified quota rule UUID is invalid. |
| 5308561 | Failed to obtain volume quota state or invalid quota state obtained for volume. |
| 5308569 | Quota policy rule delete operation succeeded, but quota resize failed due to internal error. |
| 5308572 | Quota policy rule delete operation succeeded, however the rule is still being enforced. To stop enforcing the rule, disable quotas and enable them again for this volume. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type QuotaRuleDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this quota rule delete default response has a 2xx status code
func (o *QuotaRuleDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this quota rule delete default response has a 3xx status code
func (o *QuotaRuleDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this quota rule delete default response has a 4xx status code
func (o *QuotaRuleDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this quota rule delete default response has a 5xx status code
func (o *QuotaRuleDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this quota rule delete default response a status code equal to that given
func (o *QuotaRuleDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the quota rule delete default response
func (o *QuotaRuleDeleteDefault) Code() int {
	return o._statusCode
}

func (o *QuotaRuleDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/quota/rules/{uuid}][%d] quota_rule_delete default %s", o._statusCode, payload)
}

func (o *QuotaRuleDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/quota/rules/{uuid}][%d] quota_rule_delete default %s", o._statusCode, payload)
}

func (o *QuotaRuleDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QuotaRuleDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}