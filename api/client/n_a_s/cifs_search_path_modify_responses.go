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

// CifsSearchPathModifyReader is a Reader for the CifsSearchPathModify structure.
type CifsSearchPathModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsSearchPathModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsSearchPathModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsSearchPathModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsSearchPathModifyOK creates a CifsSearchPathModifyOK with default headers values
func NewCifsSearchPathModifyOK() *CifsSearchPathModifyOK {
	return &CifsSearchPathModifyOK{}
}

/*
CifsSearchPathModifyOK describes a response with status code 200, with default header values.

OK
*/
type CifsSearchPathModifyOK struct {
}

// IsSuccess returns true when this cifs search path modify o k response has a 2xx status code
func (o *CifsSearchPathModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs search path modify o k response has a 3xx status code
func (o *CifsSearchPathModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs search path modify o k response has a 4xx status code
func (o *CifsSearchPathModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs search path modify o k response has a 5xx status code
func (o *CifsSearchPathModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs search path modify o k response a status code equal to that given
func (o *CifsSearchPathModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs search path modify o k response
func (o *CifsSearchPathModifyOK) Code() int {
	return 200
}

func (o *CifsSearchPathModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/home-directory/search-paths/{svm.uuid}/{index}][%d] cifsSearchPathModifyOK", 200)
}

func (o *CifsSearchPathModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/home-directory/search-paths/{svm.uuid}/{index}][%d] cifsSearchPathModifyOK", 200)
}

func (o *CifsSearchPathModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCifsSearchPathModifyDefault creates a CifsSearchPathModifyDefault with default headers values
func NewCifsSearchPathModifyDefault(code int) *CifsSearchPathModifyDefault {
	return &CifsSearchPathModifyDefault{
		_statusCode: code,
	}
}

/*
	CifsSearchPathModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655463     | Failed to reorder the search-path because the new-index is invalid. It cannot be '0' and it cannot go beyond the current entries |
*/
type CifsSearchPathModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs search path modify default response has a 2xx status code
func (o *CifsSearchPathModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs search path modify default response has a 3xx status code
func (o *CifsSearchPathModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs search path modify default response has a 4xx status code
func (o *CifsSearchPathModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs search path modify default response has a 5xx status code
func (o *CifsSearchPathModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs search path modify default response a status code equal to that given
func (o *CifsSearchPathModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs search path modify default response
func (o *CifsSearchPathModifyDefault) Code() int {
	return o._statusCode
}

func (o *CifsSearchPathModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/home-directory/search-paths/{svm.uuid}/{index}][%d] cifs_search_path_modify default %s", o._statusCode, payload)
}

func (o *CifsSearchPathModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/cifs/home-directory/search-paths/{svm.uuid}/{index}][%d] cifs_search_path_modify default %s", o._statusCode, payload)
}

func (o *CifsSearchPathModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsSearchPathModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
