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

// NewSecurityKeyManagerCreateParams creates a new SecurityKeyManagerCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityKeyManagerCreateParams() *SecurityKeyManagerCreateParams {
	return &SecurityKeyManagerCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityKeyManagerCreateParamsWithTimeout creates a new SecurityKeyManagerCreateParams object
// with the ability to set a timeout on a request.
func NewSecurityKeyManagerCreateParamsWithTimeout(timeout time.Duration) *SecurityKeyManagerCreateParams {
	return &SecurityKeyManagerCreateParams{
		timeout: timeout,
	}
}

// NewSecurityKeyManagerCreateParamsWithContext creates a new SecurityKeyManagerCreateParams object
// with the ability to set a context for a request.
func NewSecurityKeyManagerCreateParamsWithContext(ctx context.Context) *SecurityKeyManagerCreateParams {
	return &SecurityKeyManagerCreateParams{
		Context: ctx,
	}
}

// NewSecurityKeyManagerCreateParamsWithHTTPClient creates a new SecurityKeyManagerCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityKeyManagerCreateParamsWithHTTPClient(client *http.Client) *SecurityKeyManagerCreateParams {
	return &SecurityKeyManagerCreateParams{
		HTTPClient: client,
	}
}

/*
SecurityKeyManagerCreateParams contains all the parameters to send to the API endpoint

	for the security key manager create operation.

	Typically these are written to a http.Request.
*/
type SecurityKeyManagerCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.SecurityKeyManager

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security key manager create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerCreateParams) WithDefaults() *SecurityKeyManagerCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security key manager create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := SecurityKeyManagerCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the security key manager create params
func (o *SecurityKeyManagerCreateParams) WithTimeout(timeout time.Duration) *SecurityKeyManagerCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security key manager create params
func (o *SecurityKeyManagerCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security key manager create params
func (o *SecurityKeyManagerCreateParams) WithContext(ctx context.Context) *SecurityKeyManagerCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security key manager create params
func (o *SecurityKeyManagerCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security key manager create params
func (o *SecurityKeyManagerCreateParams) WithHTTPClient(client *http.Client) *SecurityKeyManagerCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security key manager create params
func (o *SecurityKeyManagerCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the security key manager create params
func (o *SecurityKeyManagerCreateParams) WithInfo(info *models.SecurityKeyManager) *SecurityKeyManagerCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security key manager create params
func (o *SecurityKeyManagerCreateParams) SetInfo(info *models.SecurityKeyManager) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the security key manager create params
func (o *SecurityKeyManagerCreateParams) WithReturnRecords(returnRecords *bool) *SecurityKeyManagerCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the security key manager create params
func (o *SecurityKeyManagerCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityKeyManagerCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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