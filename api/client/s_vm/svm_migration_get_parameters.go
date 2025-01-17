// Code generated by go-swagger; DO NOT EDIT.

package s_vm

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

// NewSvmMigrationGetParams creates a new SvmMigrationGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmMigrationGetParams() *SvmMigrationGetParams {
	return &SvmMigrationGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmMigrationGetParamsWithTimeout creates a new SvmMigrationGetParams object
// with the ability to set a timeout on a request.
func NewSvmMigrationGetParamsWithTimeout(timeout time.Duration) *SvmMigrationGetParams {
	return &SvmMigrationGetParams{
		timeout: timeout,
	}
}

// NewSvmMigrationGetParamsWithContext creates a new SvmMigrationGetParams object
// with the ability to set a context for a request.
func NewSvmMigrationGetParamsWithContext(ctx context.Context) *SvmMigrationGetParams {
	return &SvmMigrationGetParams{
		Context: ctx,
	}
}

// NewSvmMigrationGetParamsWithHTTPClient creates a new SvmMigrationGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmMigrationGetParamsWithHTTPClient(client *http.Client) *SvmMigrationGetParams {
	return &SvmMigrationGetParams{
		HTTPClient: client,
	}
}

/*
SvmMigrationGetParams contains all the parameters to send to the API endpoint

	for the svm migration get operation.

	Typically these are written to a http.Request.
*/
type SvmMigrationGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* UUID.

	   Migration UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the svm migration get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmMigrationGetParams) WithDefaults() *SvmMigrationGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm migration get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmMigrationGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the svm migration get params
func (o *SvmMigrationGetParams) WithTimeout(timeout time.Duration) *SvmMigrationGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm migration get params
func (o *SvmMigrationGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm migration get params
func (o *SvmMigrationGetParams) WithContext(ctx context.Context) *SvmMigrationGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm migration get params
func (o *SvmMigrationGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm migration get params
func (o *SvmMigrationGetParams) WithHTTPClient(client *http.Client) *SvmMigrationGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm migration get params
func (o *SvmMigrationGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the svm migration get params
func (o *SvmMigrationGetParams) WithFields(fields []string) *SvmMigrationGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the svm migration get params
func (o *SvmMigrationGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithUUID adds the uuid to the svm migration get params
func (o *SvmMigrationGetParams) WithUUID(uuid string) *SvmMigrationGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the svm migration get params
func (o *SvmMigrationGetParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SvmMigrationGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamSvmMigrationGet binds the parameter fields
func (o *SvmMigrationGetParams) bindParamFields(formats strfmt.Registry) []string {
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