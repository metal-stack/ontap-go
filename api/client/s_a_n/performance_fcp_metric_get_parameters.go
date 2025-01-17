// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewPerformanceFcpMetricGetParams creates a new PerformanceFcpMetricGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformanceFcpMetricGetParams() *PerformanceFcpMetricGetParams {
	return &PerformanceFcpMetricGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformanceFcpMetricGetParamsWithTimeout creates a new PerformanceFcpMetricGetParams object
// with the ability to set a timeout on a request.
func NewPerformanceFcpMetricGetParamsWithTimeout(timeout time.Duration) *PerformanceFcpMetricGetParams {
	return &PerformanceFcpMetricGetParams{
		timeout: timeout,
	}
}

// NewPerformanceFcpMetricGetParamsWithContext creates a new PerformanceFcpMetricGetParams object
// with the ability to set a context for a request.
func NewPerformanceFcpMetricGetParamsWithContext(ctx context.Context) *PerformanceFcpMetricGetParams {
	return &PerformanceFcpMetricGetParams{
		Context: ctx,
	}
}

// NewPerformanceFcpMetricGetParamsWithHTTPClient creates a new PerformanceFcpMetricGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformanceFcpMetricGetParamsWithHTTPClient(client *http.Client) *PerformanceFcpMetricGetParams {
	return &PerformanceFcpMetricGetParams{
		HTTPClient: client,
	}
}

/*
PerformanceFcpMetricGetParams contains all the parameters to send to the API endpoint

	for the performance fcp metric get operation.

	Typically these are written to a http.Request.
*/
type PerformanceFcpMetricGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* SvmUUID.

	   The unique identifier of the SVM.

	*/
	SvmUUID string

	/* Timestamp.

	   The timestamp of the performance data.


	   Format: date-time
	*/
	Timestamp strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the performance fcp metric get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformanceFcpMetricGetParams) WithDefaults() *PerformanceFcpMetricGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the performance fcp metric get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformanceFcpMetricGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the performance fcp metric get params
func (o *PerformanceFcpMetricGetParams) WithTimeout(timeout time.Duration) *PerformanceFcpMetricGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performance fcp metric get params
func (o *PerformanceFcpMetricGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performance fcp metric get params
func (o *PerformanceFcpMetricGetParams) WithContext(ctx context.Context) *PerformanceFcpMetricGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performance fcp metric get params
func (o *PerformanceFcpMetricGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performance fcp metric get params
func (o *PerformanceFcpMetricGetParams) WithHTTPClient(client *http.Client) *PerformanceFcpMetricGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performance fcp metric get params
func (o *PerformanceFcpMetricGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the performance fcp metric get params
func (o *PerformanceFcpMetricGetParams) WithFields(fields []string) *PerformanceFcpMetricGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the performance fcp metric get params
func (o *PerformanceFcpMetricGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithSvmUUID adds the svmUUID to the performance fcp metric get params
func (o *PerformanceFcpMetricGetParams) WithSvmUUID(svmUUID string) *PerformanceFcpMetricGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the performance fcp metric get params
func (o *PerformanceFcpMetricGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WithTimestamp adds the timestamp to the performance fcp metric get params
func (o *PerformanceFcpMetricGetParams) WithTimestamp(timestamp strfmt.DateTime) *PerformanceFcpMetricGetParams {
	o.SetTimestamp(timestamp)
	return o
}

// SetTimestamp adds the timestamp to the performance fcp metric get params
func (o *PerformanceFcpMetricGetParams) SetTimestamp(timestamp strfmt.DateTime) {
	o.Timestamp = timestamp
}

// WriteToRequest writes these params to a swagger request
func (o *PerformanceFcpMetricGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	// path param timestamp
	if err := r.SetPathParam("timestamp", o.Timestamp.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPerformanceFcpMetricGet binds the parameter fields
func (o *PerformanceFcpMetricGetParams) bindParamFields(formats strfmt.Registry) []string {
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