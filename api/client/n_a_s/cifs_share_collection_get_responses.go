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

// CifsShareCollectionGetReader is a Reader for the CifsShareCollectionGet structure.
type CifsShareCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsShareCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsShareCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsShareCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsShareCollectionGetOK creates a CifsShareCollectionGetOK with default headers values
func NewCifsShareCollectionGetOK() *CifsShareCollectionGetOK {
	return &CifsShareCollectionGetOK{}
}

/*
CifsShareCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type CifsShareCollectionGetOK struct {
	Payload *models.CifsShareResponse
}

// IsSuccess returns true when this cifs share collection get o k response has a 2xx status code
func (o *CifsShareCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs share collection get o k response has a 3xx status code
func (o *CifsShareCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs share collection get o k response has a 4xx status code
func (o *CifsShareCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs share collection get o k response has a 5xx status code
func (o *CifsShareCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs share collection get o k response a status code equal to that given
func (o *CifsShareCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs share collection get o k response
func (o *CifsShareCollectionGetOK) Code() int {
	return 200
}

func (o *CifsShareCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shares][%d] cifsShareCollectionGetOK %s", 200, payload)
}

func (o *CifsShareCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shares][%d] cifsShareCollectionGetOK %s", 200, payload)
}

func (o *CifsShareCollectionGetOK) GetPayload() *models.CifsShareResponse {
	return o.Payload
}

func (o *CifsShareCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsShareResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsShareCollectionGetDefault creates a CifsShareCollectionGetDefault with default headers values
func NewCifsShareCollectionGetDefault(code int) *CifsShareCollectionGetDefault {
	return &CifsShareCollectionGetDefault{
		_statusCode: code,
	}
}

/*
CifsShareCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type CifsShareCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs share collection get default response has a 2xx status code
func (o *CifsShareCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs share collection get default response has a 3xx status code
func (o *CifsShareCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs share collection get default response has a 4xx status code
func (o *CifsShareCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs share collection get default response has a 5xx status code
func (o *CifsShareCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs share collection get default response a status code equal to that given
func (o *CifsShareCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs share collection get default response
func (o *CifsShareCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *CifsShareCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shares][%d] cifs_share_collection_get default %s", o._statusCode, payload)
}

func (o *CifsShareCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shares][%d] cifs_share_collection_get default %s", o._statusCode, payload)
}

func (o *CifsShareCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsShareCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}