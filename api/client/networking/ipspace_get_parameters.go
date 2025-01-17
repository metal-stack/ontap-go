// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// NewIpspaceGetParams creates a new IpspaceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpspaceGetParams() *IpspaceGetParams {
	return &IpspaceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpspaceGetParamsWithTimeout creates a new IpspaceGetParams object
// with the ability to set a timeout on a request.
func NewIpspaceGetParamsWithTimeout(timeout time.Duration) *IpspaceGetParams {
	return &IpspaceGetParams{
		timeout: timeout,
	}
}

// NewIpspaceGetParamsWithContext creates a new IpspaceGetParams object
// with the ability to set a context for a request.
func NewIpspaceGetParamsWithContext(ctx context.Context) *IpspaceGetParams {
	return &IpspaceGetParams{
		Context: ctx,
	}
}

// NewIpspaceGetParamsWithHTTPClient creates a new IpspaceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpspaceGetParamsWithHTTPClient(client *http.Client) *IpspaceGetParams {
	return &IpspaceGetParams{
		HTTPClient: client,
	}
}

/*
IpspaceGetParams contains all the parameters to send to the API endpoint

	for the ipspace get operation.

	Typically these are written to a http.Request.
*/
type IpspaceGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* UUID.

	   IPspace UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipspace get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpspaceGetParams) WithDefaults() *IpspaceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipspace get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpspaceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipspace get params
func (o *IpspaceGetParams) WithTimeout(timeout time.Duration) *IpspaceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipspace get params
func (o *IpspaceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipspace get params
func (o *IpspaceGetParams) WithContext(ctx context.Context) *IpspaceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipspace get params
func (o *IpspaceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipspace get params
func (o *IpspaceGetParams) WithHTTPClient(client *http.Client) *IpspaceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipspace get params
func (o *IpspaceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the ipspace get params
func (o *IpspaceGetParams) WithFields(fields []string) *IpspaceGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ipspace get params
func (o *IpspaceGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithUUID adds the uuid to the ipspace get params
func (o *IpspaceGetParams) WithUUID(uuid string) *IpspaceGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the ipspace get params
func (o *IpspaceGetParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *IpspaceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamIpspaceGet binds the parameter fields
func (o *IpspaceGetParams) bindParamFields(formats strfmt.Registry) []string {
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