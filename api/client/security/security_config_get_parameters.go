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
)

// NewSecurityConfigGetParams creates a new SecurityConfigGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityConfigGetParams() *SecurityConfigGetParams {
	return &SecurityConfigGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityConfigGetParamsWithTimeout creates a new SecurityConfigGetParams object
// with the ability to set a timeout on a request.
func NewSecurityConfigGetParamsWithTimeout(timeout time.Duration) *SecurityConfigGetParams {
	return &SecurityConfigGetParams{
		timeout: timeout,
	}
}

// NewSecurityConfigGetParamsWithContext creates a new SecurityConfigGetParams object
// with the ability to set a context for a request.
func NewSecurityConfigGetParamsWithContext(ctx context.Context) *SecurityConfigGetParams {
	return &SecurityConfigGetParams{
		Context: ctx,
	}
}

// NewSecurityConfigGetParamsWithHTTPClient creates a new SecurityConfigGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityConfigGetParamsWithHTTPClient(client *http.Client) *SecurityConfigGetParams {
	return &SecurityConfigGetParams{
		HTTPClient: client,
	}
}

/*
SecurityConfigGetParams contains all the parameters to send to the API endpoint

	for the security config get operation.

	Typically these are written to a http.Request.
*/
type SecurityConfigGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security config get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityConfigGetParams) WithDefaults() *SecurityConfigGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security config get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityConfigGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security config get params
func (o *SecurityConfigGetParams) WithTimeout(timeout time.Duration) *SecurityConfigGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security config get params
func (o *SecurityConfigGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security config get params
func (o *SecurityConfigGetParams) WithContext(ctx context.Context) *SecurityConfigGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security config get params
func (o *SecurityConfigGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security config get params
func (o *SecurityConfigGetParams) WithHTTPClient(client *http.Client) *SecurityConfigGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security config get params
func (o *SecurityConfigGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the security config get params
func (o *SecurityConfigGetParams) WithFields(fields []string) *SecurityConfigGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the security config get params
func (o *SecurityConfigGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityConfigGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSecurityConfigGet binds the parameter fields
func (o *SecurityConfigGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
