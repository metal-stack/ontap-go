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

// KeyManagerKeysGetReader is a Reader for the KeyManagerKeysGet structure.
type KeyManagerKeysGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyManagerKeysGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyManagerKeysGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKeyManagerKeysGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKeyManagerKeysGetOK creates a KeyManagerKeysGetOK with default headers values
func NewKeyManagerKeysGetOK() *KeyManagerKeysGetOK {
	return &KeyManagerKeysGetOK{}
}

/*
KeyManagerKeysGetOK describes a response with status code 200, with default header values.

OK
*/
type KeyManagerKeysGetOK struct {
	Payload *models.KeyManagerKeys
}

// IsSuccess returns true when this key manager keys get o k response has a 2xx status code
func (o *KeyManagerKeysGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this key manager keys get o k response has a 3xx status code
func (o *KeyManagerKeysGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key manager keys get o k response has a 4xx status code
func (o *KeyManagerKeysGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this key manager keys get o k response has a 5xx status code
func (o *KeyManagerKeysGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this key manager keys get o k response a status code equal to that given
func (o *KeyManagerKeysGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the key manager keys get o k response
func (o *KeyManagerKeysGetOK) Code() int {
	return 200
}

func (o *KeyManagerKeysGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{security_key_manager.uuid}/keys/{node.uuid}/key-ids/{key_id}][%d] keyManagerKeysGetOK %s", 200, payload)
}

func (o *KeyManagerKeysGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{security_key_manager.uuid}/keys/{node.uuid}/key-ids/{key_id}][%d] keyManagerKeysGetOK %s", 200, payload)
}

func (o *KeyManagerKeysGetOK) GetPayload() *models.KeyManagerKeys {
	return o.Payload
}

func (o *KeyManagerKeysGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyManagerKeys)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyManagerKeysGetDefault creates a KeyManagerKeysGetDefault with default headers values
func NewKeyManagerKeysGetDefault(code int) *KeyManagerKeysGetDefault {
	return &KeyManagerKeysGetDefault{
		_statusCode: code,
	}
}

/*
KeyManagerKeysGetDefault describes a response with status code -1, with default header values.

Error
*/
type KeyManagerKeysGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this key manager keys get default response has a 2xx status code
func (o *KeyManagerKeysGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this key manager keys get default response has a 3xx status code
func (o *KeyManagerKeysGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this key manager keys get default response has a 4xx status code
func (o *KeyManagerKeysGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this key manager keys get default response has a 5xx status code
func (o *KeyManagerKeysGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this key manager keys get default response a status code equal to that given
func (o *KeyManagerKeysGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the key manager keys get default response
func (o *KeyManagerKeysGetDefault) Code() int {
	return o._statusCode
}

func (o *KeyManagerKeysGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{security_key_manager.uuid}/keys/{node.uuid}/key-ids/{key_id}][%d] key_manager_keys_get default %s", o._statusCode, payload)
}

func (o *KeyManagerKeysGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-managers/{security_key_manager.uuid}/keys/{node.uuid}/key-ids/{key_id}][%d] key_manager_keys_get default %s", o._statusCode, payload)
}

func (o *KeyManagerKeysGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *KeyManagerKeysGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}