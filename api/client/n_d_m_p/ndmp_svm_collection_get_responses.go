// Code generated by go-swagger; DO NOT EDIT.

package n_d_m_p

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

// NdmpSvmCollectionGetReader is a Reader for the NdmpSvmCollectionGet structure.
type NdmpSvmCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NdmpSvmCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNdmpSvmCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNdmpSvmCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNdmpSvmCollectionGetOK creates a NdmpSvmCollectionGetOK with default headers values
func NewNdmpSvmCollectionGetOK() *NdmpSvmCollectionGetOK {
	return &NdmpSvmCollectionGetOK{}
}

/*
NdmpSvmCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type NdmpSvmCollectionGetOK struct {
	Payload *models.NdmpSvmResponse
}

// IsSuccess returns true when this ndmp svm collection get o k response has a 2xx status code
func (o *NdmpSvmCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ndmp svm collection get o k response has a 3xx status code
func (o *NdmpSvmCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ndmp svm collection get o k response has a 4xx status code
func (o *NdmpSvmCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ndmp svm collection get o k response has a 5xx status code
func (o *NdmpSvmCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ndmp svm collection get o k response a status code equal to that given
func (o *NdmpSvmCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ndmp svm collection get o k response
func (o *NdmpSvmCollectionGetOK) Code() int {
	return 200
}

func (o *NdmpSvmCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp/svms][%d] ndmpSvmCollectionGetOK %s", 200, payload)
}

func (o *NdmpSvmCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp/svms][%d] ndmpSvmCollectionGetOK %s", 200, payload)
}

func (o *NdmpSvmCollectionGetOK) GetPayload() *models.NdmpSvmResponse {
	return o.Payload
}

func (o *NdmpSvmCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NdmpSvmResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNdmpSvmCollectionGetDefault creates a NdmpSvmCollectionGetDefault with default headers values
func NewNdmpSvmCollectionGetDefault(code int) *NdmpSvmCollectionGetDefault {
	return &NdmpSvmCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	NdmpSvmCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 2           | The UUID provided is an invalid value for field \"svm.uuid\".|
| 262197      | The value provided is invalid for field \"fields\".|
| 65601536    | The operation is not supported because NDMP SVM-aware mode is disabled.|
*/
type NdmpSvmCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ndmp svm collection get default response has a 2xx status code
func (o *NdmpSvmCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ndmp svm collection get default response has a 3xx status code
func (o *NdmpSvmCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ndmp svm collection get default response has a 4xx status code
func (o *NdmpSvmCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ndmp svm collection get default response has a 5xx status code
func (o *NdmpSvmCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ndmp svm collection get default response a status code equal to that given
func (o *NdmpSvmCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ndmp svm collection get default response
func (o *NdmpSvmCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *NdmpSvmCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp/svms][%d] ndmp_svm_collection_get default %s", o._statusCode, payload)
}

func (o *NdmpSvmCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp/svms][%d] ndmp_svm_collection_get default %s", o._statusCode, payload)
}

func (o *NdmpSvmCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NdmpSvmCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
