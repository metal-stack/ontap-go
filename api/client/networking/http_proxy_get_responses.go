// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// HTTPProxyGetReader is a Reader for the HTTPProxyGet structure.
type HTTPProxyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HTTPProxyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHTTPProxyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHTTPProxyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHTTPProxyGetOK creates a HTTPProxyGetOK with default headers values
func NewHTTPProxyGetOK() *HTTPProxyGetOK {
	return &HTTPProxyGetOK{}
}

/*
HTTPProxyGetOK describes a response with status code 200, with default header values.

OK
*/
type HTTPProxyGetOK struct {
	Payload *models.NetworkHTTPProxy
}

// IsSuccess returns true when this http proxy get o k response has a 2xx status code
func (o *HTTPProxyGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this http proxy get o k response has a 3xx status code
func (o *HTTPProxyGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this http proxy get o k response has a 4xx status code
func (o *HTTPProxyGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this http proxy get o k response has a 5xx status code
func (o *HTTPProxyGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this http proxy get o k response a status code equal to that given
func (o *HTTPProxyGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the http proxy get o k response
func (o *HTTPProxyGetOK) Code() int {
	return 200
}

func (o *HTTPProxyGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/http-proxy/{uuid}][%d] httpProxyGetOK %s", 200, payload)
}

func (o *HTTPProxyGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/http-proxy/{uuid}][%d] httpProxyGetOK %s", 200, payload)
}

func (o *HTTPProxyGetOK) GetPayload() *models.NetworkHTTPProxy {
	return o.Payload
}

func (o *HTTPProxyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkHTTPProxy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHTTPProxyGetDefault creates a HTTPProxyGetDefault with default headers values
func NewHTTPProxyGetDefault(code int) *HTTPProxyGetDefault {
	return &HTTPProxyGetDefault{
		_statusCode: code,
	}
}

/*
HTTPProxyGetDefault describes a response with status code -1, with default header values.

Error
*/
type HTTPProxyGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this http proxy get default response has a 2xx status code
func (o *HTTPProxyGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this http proxy get default response has a 3xx status code
func (o *HTTPProxyGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this http proxy get default response has a 4xx status code
func (o *HTTPProxyGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this http proxy get default response has a 5xx status code
func (o *HTTPProxyGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this http proxy get default response a status code equal to that given
func (o *HTTPProxyGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the http proxy get default response
func (o *HTTPProxyGetDefault) Code() int {
	return o._statusCode
}

func (o *HTTPProxyGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/http-proxy/{uuid}][%d] http_proxy_get default %s", o._statusCode, payload)
}

func (o *HTTPProxyGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/http-proxy/{uuid}][%d] http_proxy_get default %s", o._statusCode, payload)
}

func (o *HTTPProxyGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *HTTPProxyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}