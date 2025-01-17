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

// NewIpsecPolicyCreateParams creates a new IpsecPolicyCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpsecPolicyCreateParams() *IpsecPolicyCreateParams {
	return &IpsecPolicyCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpsecPolicyCreateParamsWithTimeout creates a new IpsecPolicyCreateParams object
// with the ability to set a timeout on a request.
func NewIpsecPolicyCreateParamsWithTimeout(timeout time.Duration) *IpsecPolicyCreateParams {
	return &IpsecPolicyCreateParams{
		timeout: timeout,
	}
}

// NewIpsecPolicyCreateParamsWithContext creates a new IpsecPolicyCreateParams object
// with the ability to set a context for a request.
func NewIpsecPolicyCreateParamsWithContext(ctx context.Context) *IpsecPolicyCreateParams {
	return &IpsecPolicyCreateParams{
		Context: ctx,
	}
}

// NewIpsecPolicyCreateParamsWithHTTPClient creates a new IpsecPolicyCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpsecPolicyCreateParamsWithHTTPClient(client *http.Client) *IpsecPolicyCreateParams {
	return &IpsecPolicyCreateParams{
		HTTPClient: client,
	}
}

/*
IpsecPolicyCreateParams contains all the parameters to send to the API endpoint

	for the ipsec policy create operation.

	Typically these are written to a http.Request.
*/
type IpsecPolicyCreateParams struct {

	/* Info.

	   IPsec policy parameters
	*/
	Info *models.IpsecPolicy

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipsec policy create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpsecPolicyCreateParams) WithDefaults() *IpsecPolicyCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipsec policy create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpsecPolicyCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := IpsecPolicyCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ipsec policy create params
func (o *IpsecPolicyCreateParams) WithTimeout(timeout time.Duration) *IpsecPolicyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipsec policy create params
func (o *IpsecPolicyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipsec policy create params
func (o *IpsecPolicyCreateParams) WithContext(ctx context.Context) *IpsecPolicyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipsec policy create params
func (o *IpsecPolicyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipsec policy create params
func (o *IpsecPolicyCreateParams) WithHTTPClient(client *http.Client) *IpsecPolicyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipsec policy create params
func (o *IpsecPolicyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the ipsec policy create params
func (o *IpsecPolicyCreateParams) WithInfo(info *models.IpsecPolicy) *IpsecPolicyCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ipsec policy create params
func (o *IpsecPolicyCreateParams) SetInfo(info *models.IpsecPolicy) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the ipsec policy create params
func (o *IpsecPolicyCreateParams) WithReturnRecords(returnRecords *bool) *IpsecPolicyCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ipsec policy create params
func (o *IpsecPolicyCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *IpsecPolicyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
