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

// CifsServiceCollectionGetReader is a Reader for the CifsServiceCollectionGet structure.
type CifsServiceCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsServiceCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsServiceCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsServiceCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsServiceCollectionGetOK creates a CifsServiceCollectionGetOK with default headers values
func NewCifsServiceCollectionGetOK() *CifsServiceCollectionGetOK {
	return &CifsServiceCollectionGetOK{}
}

/*
CifsServiceCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type CifsServiceCollectionGetOK struct {
	Payload *models.CifsServiceResponse
}

// IsSuccess returns true when this cifs service collection get o k response has a 2xx status code
func (o *CifsServiceCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs service collection get o k response has a 3xx status code
func (o *CifsServiceCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs service collection get o k response has a 4xx status code
func (o *CifsServiceCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs service collection get o k response has a 5xx status code
func (o *CifsServiceCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs service collection get o k response a status code equal to that given
func (o *CifsServiceCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs service collection get o k response
func (o *CifsServiceCollectionGetOK) Code() int {
	return 200
}

func (o *CifsServiceCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/services][%d] cifsServiceCollectionGetOK %s", 200, payload)
}

func (o *CifsServiceCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/services][%d] cifsServiceCollectionGetOK %s", 200, payload)
}

func (o *CifsServiceCollectionGetOK) GetPayload() *models.CifsServiceResponse {
	return o.Payload
}

func (o *CifsServiceCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsServiceCollectionGetDefault creates a CifsServiceCollectionGetDefault with default headers values
func NewCifsServiceCollectionGetDefault(code int) *CifsServiceCollectionGetDefault {
	return &CifsServiceCollectionGetDefault{
		_statusCode: code,
	}
}

/*
CifsServiceCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type CifsServiceCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs service collection get default response has a 2xx status code
func (o *CifsServiceCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs service collection get default response has a 3xx status code
func (o *CifsServiceCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs service collection get default response has a 4xx status code
func (o *CifsServiceCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs service collection get default response has a 5xx status code
func (o *CifsServiceCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs service collection get default response a status code equal to that given
func (o *CifsServiceCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs service collection get default response
func (o *CifsServiceCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *CifsServiceCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/services][%d] cifs_service_collection_get default %s", o._statusCode, payload)
}

func (o *CifsServiceCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/services][%d] cifs_service_collection_get default %s", o._statusCode, payload)
}

func (o *CifsServiceCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsServiceCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
