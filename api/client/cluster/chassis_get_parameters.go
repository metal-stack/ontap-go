// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewChassisGetParams creates a new ChassisGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChassisGetParams() *ChassisGetParams {
	return &ChassisGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChassisGetParamsWithTimeout creates a new ChassisGetParams object
// with the ability to set a timeout on a request.
func NewChassisGetParamsWithTimeout(timeout time.Duration) *ChassisGetParams {
	return &ChassisGetParams{
		timeout: timeout,
	}
}

// NewChassisGetParamsWithContext creates a new ChassisGetParams object
// with the ability to set a context for a request.
func NewChassisGetParamsWithContext(ctx context.Context) *ChassisGetParams {
	return &ChassisGetParams{
		Context: ctx,
	}
}

// NewChassisGetParamsWithHTTPClient creates a new ChassisGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewChassisGetParamsWithHTTPClient(client *http.Client) *ChassisGetParams {
	return &ChassisGetParams{
		HTTPClient: client,
	}
}

/*
ChassisGetParams contains all the parameters to send to the API endpoint

	for the chassis get operation.

	Typically these are written to a http.Request.
*/
type ChassisGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* ID.

	   Chassis ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the chassis get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChassisGetParams) WithDefaults() *ChassisGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the chassis get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChassisGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the chassis get params
func (o *ChassisGetParams) WithTimeout(timeout time.Duration) *ChassisGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the chassis get params
func (o *ChassisGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the chassis get params
func (o *ChassisGetParams) WithContext(ctx context.Context) *ChassisGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the chassis get params
func (o *ChassisGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the chassis get params
func (o *ChassisGetParams) WithHTTPClient(client *http.Client) *ChassisGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the chassis get params
func (o *ChassisGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the chassis get params
func (o *ChassisGetParams) WithFields(fields []string) *ChassisGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the chassis get params
func (o *ChassisGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the chassis get params
func (o *ChassisGetParams) WithID(id string) *ChassisGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the chassis get params
func (o *ChassisGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ChassisGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamChassisGet binds the parameter fields
func (o *ChassisGetParams) bindParamFields(formats strfmt.Registry) []string {
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
