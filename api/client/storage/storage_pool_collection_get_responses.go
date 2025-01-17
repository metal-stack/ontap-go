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

// StoragePoolCollectionGetReader is a Reader for the StoragePoolCollectionGet structure.
type StoragePoolCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StoragePoolCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStoragePoolCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStoragePoolCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStoragePoolCollectionGetOK creates a StoragePoolCollectionGetOK with default headers values
func NewStoragePoolCollectionGetOK() *StoragePoolCollectionGetOK {
	return &StoragePoolCollectionGetOK{}
}

/*
StoragePoolCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type StoragePoolCollectionGetOK struct {
	Payload *models.StoragePoolResponse
}

// IsSuccess returns true when this storage pool collection get o k response has a 2xx status code
func (o *StoragePoolCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this storage pool collection get o k response has a 3xx status code
func (o *StoragePoolCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this storage pool collection get o k response has a 4xx status code
func (o *StoragePoolCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this storage pool collection get o k response has a 5xx status code
func (o *StoragePoolCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this storage pool collection get o k response a status code equal to that given
func (o *StoragePoolCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the storage pool collection get o k response
func (o *StoragePoolCollectionGetOK) Code() int {
	return 200
}

func (o *StoragePoolCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/pools][%d] storagePoolCollectionGetOK %s", 200, payload)
}

func (o *StoragePoolCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/pools][%d] storagePoolCollectionGetOK %s", 200, payload)
}

func (o *StoragePoolCollectionGetOK) GetPayload() *models.StoragePoolResponse {
	return o.Payload
}

func (o *StoragePoolCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePoolResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStoragePoolCollectionGetDefault creates a StoragePoolCollectionGetDefault with default headers values
func NewStoragePoolCollectionGetDefault(code int) *StoragePoolCollectionGetDefault {
	return &StoragePoolCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	StoragePoolCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 10944513 | Unable to retrieve shared disk capability information. |
| 10944514 | Unable to enable shared disk capability. |
| 10944527 | Storage pools are not supported in MetroCluster configurations. |
| 10944528 | Unable to retrieve MetroCluster configuration information. |
| 11206662 | There is no storage pool matching the specified UUID or name. |
| 11206667 | Storage pool feature is not enabled. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type StoragePoolCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this storage pool collection get default response has a 2xx status code
func (o *StoragePoolCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this storage pool collection get default response has a 3xx status code
func (o *StoragePoolCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this storage pool collection get default response has a 4xx status code
func (o *StoragePoolCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this storage pool collection get default response has a 5xx status code
func (o *StoragePoolCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this storage pool collection get default response a status code equal to that given
func (o *StoragePoolCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the storage pool collection get default response
func (o *StoragePoolCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *StoragePoolCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/pools][%d] storage_pool_collection_get default %s", o._statusCode, payload)
}

func (o *StoragePoolCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/pools][%d] storage_pool_collection_get default %s", o._statusCode, payload)
}

func (o *StoragePoolCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StoragePoolCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
