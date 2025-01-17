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

// AzureKeyVaultDeleteReader is a Reader for the AzureKeyVaultDelete structure.
type AzureKeyVaultDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzureKeyVaultDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAzureKeyVaultDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAzureKeyVaultDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAzureKeyVaultDeleteOK creates a AzureKeyVaultDeleteOK with default headers values
func NewAzureKeyVaultDeleteOK() *AzureKeyVaultDeleteOK {
	return &AzureKeyVaultDeleteOK{}
}

/*
AzureKeyVaultDeleteOK describes a response with status code 200, with default header values.

OK
*/
type AzureKeyVaultDeleteOK struct {
}

// IsSuccess returns true when this azure key vault delete o k response has a 2xx status code
func (o *AzureKeyVaultDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure key vault delete o k response has a 3xx status code
func (o *AzureKeyVaultDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure key vault delete o k response has a 4xx status code
func (o *AzureKeyVaultDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure key vault delete o k response has a 5xx status code
func (o *AzureKeyVaultDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this azure key vault delete o k response a status code equal to that given
func (o *AzureKeyVaultDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the azure key vault delete o k response
func (o *AzureKeyVaultDeleteOK) Code() int {
	return 200
}

func (o *AzureKeyVaultDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/azure-key-vaults/{uuid}][%d] azureKeyVaultDeleteOK", 200)
}

func (o *AzureKeyVaultDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/azure-key-vaults/{uuid}][%d] azureKeyVaultDeleteOK", 200)
}

func (o *AzureKeyVaultDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAzureKeyVaultDeleteDefault creates a AzureKeyVaultDeleteDefault with default headers values
func NewAzureKeyVaultDeleteDefault(code int) *AzureKeyVaultDeleteDefault {
	return &AzureKeyVaultDeleteDefault{
		_statusCode: code,
	}
}

/*
	AzureKeyVaultDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536242 | One or more self-encrypting drives are assigned an authentication key. |
| 65536243 | Cannot determine authentication key presence on one or more self-encrypting drives. |
| 65536817 | Internal error. Failed to determine if key manager is safe to disable. |
| 65536827 | Internal error. Failed to determine if the given SVM has any encrypted volumes. |
| 65536834 | Internal error. Failed to get existing key-server details for the given SVM. |
| 65536867 | Volume encryption keys (VEK) for one or more encrypted volumes are stored on the key manager configured for the given SVM. |
| 65536883 | Internal error. Volume encryption key is missing for a volume. |
| 65536884 | Internal error. Volume encryption key is invalid for a volume. |
| 65536924 | Cannot remove key manager that still contains one or more NSE authentication keys. |
| 65537120 | Azure Key Vault is not configured for the given SVM. |
| 196608080 | One or more nodes in the cluster have the root volume encrypted using NVE (NetApp Volume Encryption). |
| 196608301 | Internal error. Failed to get encryption type. |
| 196608305 | NAE aggregates found in the cluster. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AzureKeyVaultDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this azure key vault delete default response has a 2xx status code
func (o *AzureKeyVaultDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this azure key vault delete default response has a 3xx status code
func (o *AzureKeyVaultDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this azure key vault delete default response has a 4xx status code
func (o *AzureKeyVaultDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this azure key vault delete default response has a 5xx status code
func (o *AzureKeyVaultDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this azure key vault delete default response a status code equal to that given
func (o *AzureKeyVaultDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the azure key vault delete default response
func (o *AzureKeyVaultDeleteDefault) Code() int {
	return o._statusCode
}

func (o *AzureKeyVaultDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/azure-key-vaults/{uuid}][%d] azure_key_vault_delete default %s", o._statusCode, payload)
}

func (o *AzureKeyVaultDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/azure-key-vaults/{uuid}][%d] azure_key_vault_delete default %s", o._statusCode, payload)
}

func (o *AzureKeyVaultDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AzureKeyVaultDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
