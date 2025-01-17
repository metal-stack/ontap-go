// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/metal-stack/ontap-go/api/models"
)

// NewAzureKeyVaultCreateParams creates a new AzureKeyVaultCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAzureKeyVaultCreateParams() *AzureKeyVaultCreateParams {
	return &AzureKeyVaultCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAzureKeyVaultCreateParamsWithTimeout creates a new AzureKeyVaultCreateParams object
// with the ability to set a timeout on a request.
func NewAzureKeyVaultCreateParamsWithTimeout(timeout time.Duration) *AzureKeyVaultCreateParams {
	return &AzureKeyVaultCreateParams{
		timeout: timeout,
	}
}

// NewAzureKeyVaultCreateParamsWithContext creates a new AzureKeyVaultCreateParams object
// with the ability to set a context for a request.
func NewAzureKeyVaultCreateParamsWithContext(ctx context.Context) *AzureKeyVaultCreateParams {
	return &AzureKeyVaultCreateParams{
		Context: ctx,
	}
}

// NewAzureKeyVaultCreateParamsWithHTTPClient creates a new AzureKeyVaultCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAzureKeyVaultCreateParamsWithHTTPClient(client *http.Client) *AzureKeyVaultCreateParams {
	return &AzureKeyVaultCreateParams{
		HTTPClient: client,
	}
}

/*
AzureKeyVaultCreateParams contains all the parameters to send to the API endpoint

	for the azure key vault create operation.

	Typically these are written to a http.Request.
*/
type AzureKeyVaultCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.AzureKeyVault

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the azure key vault create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureKeyVaultCreateParams) WithDefaults() *AzureKeyVaultCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the azure key vault create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureKeyVaultCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := AzureKeyVaultCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the azure key vault create params
func (o *AzureKeyVaultCreateParams) WithTimeout(timeout time.Duration) *AzureKeyVaultCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the azure key vault create params
func (o *AzureKeyVaultCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the azure key vault create params
func (o *AzureKeyVaultCreateParams) WithContext(ctx context.Context) *AzureKeyVaultCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the azure key vault create params
func (o *AzureKeyVaultCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the azure key vault create params
func (o *AzureKeyVaultCreateParams) WithHTTPClient(client *http.Client) *AzureKeyVaultCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the azure key vault create params
func (o *AzureKeyVaultCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the azure key vault create params
func (o *AzureKeyVaultCreateParams) WithInfo(info *models.AzureKeyVault) *AzureKeyVaultCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the azure key vault create params
func (o *AzureKeyVaultCreateParams) SetInfo(info *models.AzureKeyVault) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the azure key vault create params
func (o *AzureKeyVaultCreateParams) WithReturnRecords(returnRecords *bool) *AzureKeyVaultCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the azure key vault create params
func (o *AzureKeyVaultCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *AzureKeyVaultCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
