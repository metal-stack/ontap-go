// Code generated by go-swagger; DO NOT EDIT.

package snap_lock

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

// NewSnaplockLegalHoldInstanceGetParams creates a new SnaplockLegalHoldInstanceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockLegalHoldInstanceGetParams() *SnaplockLegalHoldInstanceGetParams {
	return &SnaplockLegalHoldInstanceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockLegalHoldInstanceGetParamsWithTimeout creates a new SnaplockLegalHoldInstanceGetParams object
// with the ability to set a timeout on a request.
func NewSnaplockLegalHoldInstanceGetParamsWithTimeout(timeout time.Duration) *SnaplockLegalHoldInstanceGetParams {
	return &SnaplockLegalHoldInstanceGetParams{
		timeout: timeout,
	}
}

// NewSnaplockLegalHoldInstanceGetParamsWithContext creates a new SnaplockLegalHoldInstanceGetParams object
// with the ability to set a context for a request.
func NewSnaplockLegalHoldInstanceGetParamsWithContext(ctx context.Context) *SnaplockLegalHoldInstanceGetParams {
	return &SnaplockLegalHoldInstanceGetParams{
		Context: ctx,
	}
}

// NewSnaplockLegalHoldInstanceGetParamsWithHTTPClient creates a new SnaplockLegalHoldInstanceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockLegalHoldInstanceGetParamsWithHTTPClient(client *http.Client) *SnaplockLegalHoldInstanceGetParams {
	return &SnaplockLegalHoldInstanceGetParams{
		HTTPClient: client,
	}
}

/*
SnaplockLegalHoldInstanceGetParams contains all the parameters to send to the API endpoint

	for the snaplock legal hold instance get operation.

	Typically these are written to a http.Request.
*/
type SnaplockLegalHoldInstanceGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* ID.

	   Litigation ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock legal hold instance get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldInstanceGetParams) WithDefaults() *SnaplockLegalHoldInstanceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock legal hold instance get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldInstanceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snaplock legal hold instance get params
func (o *SnaplockLegalHoldInstanceGetParams) WithTimeout(timeout time.Duration) *SnaplockLegalHoldInstanceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock legal hold instance get params
func (o *SnaplockLegalHoldInstanceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock legal hold instance get params
func (o *SnaplockLegalHoldInstanceGetParams) WithContext(ctx context.Context) *SnaplockLegalHoldInstanceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock legal hold instance get params
func (o *SnaplockLegalHoldInstanceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock legal hold instance get params
func (o *SnaplockLegalHoldInstanceGetParams) WithHTTPClient(client *http.Client) *SnaplockLegalHoldInstanceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock legal hold instance get params
func (o *SnaplockLegalHoldInstanceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the snaplock legal hold instance get params
func (o *SnaplockLegalHoldInstanceGetParams) WithFields(fields []string) *SnaplockLegalHoldInstanceGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the snaplock legal hold instance get params
func (o *SnaplockLegalHoldInstanceGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the snaplock legal hold instance get params
func (o *SnaplockLegalHoldInstanceGetParams) WithID(id string) *SnaplockLegalHoldInstanceGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the snaplock legal hold instance get params
func (o *SnaplockLegalHoldInstanceGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockLegalHoldInstanceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSnaplockLegalHoldInstanceGet binds the parameter fields
func (o *SnaplockLegalHoldInstanceGetParams) bindParamFields(formats strfmt.Registry) []string {
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
