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

// NewAzureKeyVaultModifyParams creates a new AzureKeyVaultModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAzureKeyVaultModifyParams() *AzureKeyVaultModifyParams {
	return &AzureKeyVaultModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAzureKeyVaultModifyParamsWithTimeout creates a new AzureKeyVaultModifyParams object
// with the ability to set a timeout on a request.
func NewAzureKeyVaultModifyParamsWithTimeout(timeout time.Duration) *AzureKeyVaultModifyParams {
	return &AzureKeyVaultModifyParams{
		timeout: timeout,
	}
}

// NewAzureKeyVaultModifyParamsWithContext creates a new AzureKeyVaultModifyParams object
// with the ability to set a context for a request.
func NewAzureKeyVaultModifyParamsWithContext(ctx context.Context) *AzureKeyVaultModifyParams {
	return &AzureKeyVaultModifyParams{
		Context: ctx,
	}
}

// NewAzureKeyVaultModifyParamsWithHTTPClient creates a new AzureKeyVaultModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewAzureKeyVaultModifyParamsWithHTTPClient(client *http.Client) *AzureKeyVaultModifyParams {
	return &AzureKeyVaultModifyParams{
		HTTPClient: client,
	}
}

/*
AzureKeyVaultModifyParams contains all the parameters to send to the API endpoint

	for the azure key vault modify operation.

	Typically these are written to a http.Request.
*/
type AzureKeyVaultModifyParams struct {

	/* Info.

	   AKV information
	*/
	Info *models.AzureKeyVault

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   AKV UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the azure key vault modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureKeyVaultModifyParams) WithDefaults() *AzureKeyVaultModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the azure key vault modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureKeyVaultModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := AzureKeyVaultModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the azure key vault modify params
func (o *AzureKeyVaultModifyParams) WithTimeout(timeout time.Duration) *AzureKeyVaultModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the azure key vault modify params
func (o *AzureKeyVaultModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the azure key vault modify params
func (o *AzureKeyVaultModifyParams) WithContext(ctx context.Context) *AzureKeyVaultModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the azure key vault modify params
func (o *AzureKeyVaultModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the azure key vault modify params
func (o *AzureKeyVaultModifyParams) WithHTTPClient(client *http.Client) *AzureKeyVaultModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the azure key vault modify params
func (o *AzureKeyVaultModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the azure key vault modify params
func (o *AzureKeyVaultModifyParams) WithInfo(info *models.AzureKeyVault) *AzureKeyVaultModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the azure key vault modify params
func (o *AzureKeyVaultModifyParams) SetInfo(info *models.AzureKeyVault) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the azure key vault modify params
func (o *AzureKeyVaultModifyParams) WithReturnTimeout(returnTimeout *int64) *AzureKeyVaultModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the azure key vault modify params
func (o *AzureKeyVaultModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUID adds the uuid to the azure key vault modify params
func (o *AzureKeyVaultModifyParams) WithUUID(uuid string) *AzureKeyVaultModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the azure key vault modify params
func (o *AzureKeyVaultModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *AzureKeyVaultModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
