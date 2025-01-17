// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewCifsHomedirSearchPathGetParams creates a new CifsHomedirSearchPathGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsHomedirSearchPathGetParams() *CifsHomedirSearchPathGetParams {
	return &CifsHomedirSearchPathGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsHomedirSearchPathGetParamsWithTimeout creates a new CifsHomedirSearchPathGetParams object
// with the ability to set a timeout on a request.
func NewCifsHomedirSearchPathGetParamsWithTimeout(timeout time.Duration) *CifsHomedirSearchPathGetParams {
	return &CifsHomedirSearchPathGetParams{
		timeout: timeout,
	}
}

// NewCifsHomedirSearchPathGetParamsWithContext creates a new CifsHomedirSearchPathGetParams object
// with the ability to set a context for a request.
func NewCifsHomedirSearchPathGetParamsWithContext(ctx context.Context) *CifsHomedirSearchPathGetParams {
	return &CifsHomedirSearchPathGetParams{
		Context: ctx,
	}
}

// NewCifsHomedirSearchPathGetParamsWithHTTPClient creates a new CifsHomedirSearchPathGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsHomedirSearchPathGetParamsWithHTTPClient(client *http.Client) *CifsHomedirSearchPathGetParams {
	return &CifsHomedirSearchPathGetParams{
		HTTPClient: client,
	}
}

/*
CifsHomedirSearchPathGetParams contains all the parameters to send to the API endpoint

	for the cifs homedir search path get operation.

	Typically these are written to a http.Request.
*/
type CifsHomedirSearchPathGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Index.

	   Home directory search path index
	*/
	Index int64

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs homedir search path get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsHomedirSearchPathGetParams) WithDefaults() *CifsHomedirSearchPathGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs homedir search path get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsHomedirSearchPathGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cifs homedir search path get params
func (o *CifsHomedirSearchPathGetParams) WithTimeout(timeout time.Duration) *CifsHomedirSearchPathGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs homedir search path get params
func (o *CifsHomedirSearchPathGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs homedir search path get params
func (o *CifsHomedirSearchPathGetParams) WithContext(ctx context.Context) *CifsHomedirSearchPathGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs homedir search path get params
func (o *CifsHomedirSearchPathGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs homedir search path get params
func (o *CifsHomedirSearchPathGetParams) WithHTTPClient(client *http.Client) *CifsHomedirSearchPathGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs homedir search path get params
func (o *CifsHomedirSearchPathGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the cifs homedir search path get params
func (o *CifsHomedirSearchPathGetParams) WithFields(fields []string) *CifsHomedirSearchPathGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the cifs homedir search path get params
func (o *CifsHomedirSearchPathGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIndex adds the index to the cifs homedir search path get params
func (o *CifsHomedirSearchPathGetParams) WithIndex(index int64) *CifsHomedirSearchPathGetParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the cifs homedir search path get params
func (o *CifsHomedirSearchPathGetParams) SetIndex(index int64) {
	o.Index = index
}

// WithSvmUUID adds the svmUUID to the cifs homedir search path get params
func (o *CifsHomedirSearchPathGetParams) WithSvmUUID(svmUUID string) *CifsHomedirSearchPathGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the cifs homedir search path get params
func (o *CifsHomedirSearchPathGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CifsHomedirSearchPathGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCifsHomedirSearchPathGet binds the parameter fields
func (o *CifsHomedirSearchPathGetParams) bindParamFields(formats strfmt.Registry) []string {
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
