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

// KeyManagerConfigGetReader is a Reader for the KeyManagerConfigGet structure.
type KeyManagerConfigGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyManagerConfigGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyManagerConfigGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKeyManagerConfigGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKeyManagerConfigGetOK creates a KeyManagerConfigGetOK with default headers values
func NewKeyManagerConfigGetOK() *KeyManagerConfigGetOK {
	return &KeyManagerConfigGetOK{}
}

/*
KeyManagerConfigGetOK describes a response with status code 200, with default header values.

OK
*/
type KeyManagerConfigGetOK struct {
	Payload *models.KeyManagerConfig
}

// IsSuccess returns true when this key manager config get o k response has a 2xx status code
func (o *KeyManagerConfigGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this key manager config get o k response has a 3xx status code
func (o *KeyManagerConfigGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this key manager config get o k response has a 4xx status code
func (o *KeyManagerConfigGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this key manager config get o k response has a 5xx status code
func (o *KeyManagerConfigGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this key manager config get o k response a status code equal to that given
func (o *KeyManagerConfigGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the key manager config get o k response
func (o *KeyManagerConfigGetOK) Code() int {
	return 200
}

func (o *KeyManagerConfigGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-manager-configs][%d] keyManagerConfigGetOK %s", 200, payload)
}

func (o *KeyManagerConfigGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-manager-configs][%d] keyManagerConfigGetOK %s", 200, payload)
}

func (o *KeyManagerConfigGetOK) GetPayload() *models.KeyManagerConfig {
	return o.Payload
}

func (o *KeyManagerConfigGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyManagerConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyManagerConfigGetDefault creates a KeyManagerConfigGetDefault with default headers values
func NewKeyManagerConfigGetDefault(code int) *KeyManagerConfigGetDefault {
	return &KeyManagerConfigGetDefault{
		_statusCode: code,
	}
}

/*
KeyManagerConfigGetDefault describes a response with status code -1, with default header values.

Error
*/
type KeyManagerConfigGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this key manager config get default response has a 2xx status code
func (o *KeyManagerConfigGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this key manager config get default response has a 3xx status code
func (o *KeyManagerConfigGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this key manager config get default response has a 4xx status code
func (o *KeyManagerConfigGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this key manager config get default response has a 5xx status code
func (o *KeyManagerConfigGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this key manager config get default response a status code equal to that given
func (o *KeyManagerConfigGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the key manager config get default response
func (o *KeyManagerConfigGetDefault) Code() int {
	return o._statusCode
}

func (o *KeyManagerConfigGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-manager-configs][%d] key_manager_config_get default %s", o._statusCode, payload)
}

func (o *KeyManagerConfigGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/key-manager-configs][%d] key_manager_config_get default %s", o._statusCode, payload)
}

func (o *KeyManagerConfigGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *KeyManagerConfigGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}