// Code generated by go-swagger; DO NOT EDIT.

package object_store

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

// NewS3BucketLifecycleRuleGetParams creates a new S3BucketLifecycleRuleGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3BucketLifecycleRuleGetParams() *S3BucketLifecycleRuleGetParams {
	return &S3BucketLifecycleRuleGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3BucketLifecycleRuleGetParamsWithTimeout creates a new S3BucketLifecycleRuleGetParams object
// with the ability to set a timeout on a request.
func NewS3BucketLifecycleRuleGetParamsWithTimeout(timeout time.Duration) *S3BucketLifecycleRuleGetParams {
	return &S3BucketLifecycleRuleGetParams{
		timeout: timeout,
	}
}

// NewS3BucketLifecycleRuleGetParamsWithContext creates a new S3BucketLifecycleRuleGetParams object
// with the ability to set a context for a request.
func NewS3BucketLifecycleRuleGetParamsWithContext(ctx context.Context) *S3BucketLifecycleRuleGetParams {
	return &S3BucketLifecycleRuleGetParams{
		Context: ctx,
	}
}

// NewS3BucketLifecycleRuleGetParamsWithHTTPClient creates a new S3BucketLifecycleRuleGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3BucketLifecycleRuleGetParamsWithHTTPClient(client *http.Client) *S3BucketLifecycleRuleGetParams {
	return &S3BucketLifecycleRuleGetParams{
		HTTPClient: client,
	}
}

/*
S3BucketLifecycleRuleGetParams contains all the parameters to send to the API endpoint

	for the s3 bucket lifecycle rule get operation.

	Typically these are written to a http.Request.
*/
type S3BucketLifecycleRuleGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Name.

	   The name of the lifecycle management rule to be applied on the bucket.
	*/
	Name string

	/* S3BucketUUID.

	   The unique identifier of the bucket.
	*/
	S3BucketUUID string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 bucket lifecycle rule get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3BucketLifecycleRuleGetParams) WithDefaults() *S3BucketLifecycleRuleGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 bucket lifecycle rule get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3BucketLifecycleRuleGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the s3 bucket lifecycle rule get params
func (o *S3BucketLifecycleRuleGetParams) WithTimeout(timeout time.Duration) *S3BucketLifecycleRuleGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 bucket lifecycle rule get params
func (o *S3BucketLifecycleRuleGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 bucket lifecycle rule get params
func (o *S3BucketLifecycleRuleGetParams) WithContext(ctx context.Context) *S3BucketLifecycleRuleGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 bucket lifecycle rule get params
func (o *S3BucketLifecycleRuleGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 bucket lifecycle rule get params
func (o *S3BucketLifecycleRuleGetParams) WithHTTPClient(client *http.Client) *S3BucketLifecycleRuleGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 bucket lifecycle rule get params
func (o *S3BucketLifecycleRuleGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the s3 bucket lifecycle rule get params
func (o *S3BucketLifecycleRuleGetParams) WithFields(fields []string) *S3BucketLifecycleRuleGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the s3 bucket lifecycle rule get params
func (o *S3BucketLifecycleRuleGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithName adds the name to the s3 bucket lifecycle rule get params
func (o *S3BucketLifecycleRuleGetParams) WithName(name string) *S3BucketLifecycleRuleGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the s3 bucket lifecycle rule get params
func (o *S3BucketLifecycleRuleGetParams) SetName(name string) {
	o.Name = name
}

// WithS3BucketUUID adds the s3BucketUUID to the s3 bucket lifecycle rule get params
func (o *S3BucketLifecycleRuleGetParams) WithS3BucketUUID(s3BucketUUID string) *S3BucketLifecycleRuleGetParams {
	o.SetS3BucketUUID(s3BucketUUID)
	return o
}

// SetS3BucketUUID adds the s3BucketUuid to the s3 bucket lifecycle rule get params
func (o *S3BucketLifecycleRuleGetParams) SetS3BucketUUID(s3BucketUUID string) {
	o.S3BucketUUID = s3BucketUUID
}

// WithSvmUUID adds the svmUUID to the s3 bucket lifecycle rule get params
func (o *S3BucketLifecycleRuleGetParams) WithSvmUUID(svmUUID string) *S3BucketLifecycleRuleGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the s3 bucket lifecycle rule get params
func (o *S3BucketLifecycleRuleGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *S3BucketLifecycleRuleGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param s3_bucket.uuid
	if err := r.SetPathParam("s3_bucket.uuid", o.S3BucketUUID); err != nil {
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

// bindParamS3BucketLifecycleRuleGet binds the parameter fields
func (o *S3BucketLifecycleRuleGetParams) bindParamFields(formats strfmt.Registry) []string {
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