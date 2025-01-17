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

	"github.com/metal-stack/ontap-go/api/models"
)

// NewS3BucketLifecycleRuleModifyParams creates a new S3BucketLifecycleRuleModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3BucketLifecycleRuleModifyParams() *S3BucketLifecycleRuleModifyParams {
	return &S3BucketLifecycleRuleModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3BucketLifecycleRuleModifyParamsWithTimeout creates a new S3BucketLifecycleRuleModifyParams object
// with the ability to set a timeout on a request.
func NewS3BucketLifecycleRuleModifyParamsWithTimeout(timeout time.Duration) *S3BucketLifecycleRuleModifyParams {
	return &S3BucketLifecycleRuleModifyParams{
		timeout: timeout,
	}
}

// NewS3BucketLifecycleRuleModifyParamsWithContext creates a new S3BucketLifecycleRuleModifyParams object
// with the ability to set a context for a request.
func NewS3BucketLifecycleRuleModifyParamsWithContext(ctx context.Context) *S3BucketLifecycleRuleModifyParams {
	return &S3BucketLifecycleRuleModifyParams{
		Context: ctx,
	}
}

// NewS3BucketLifecycleRuleModifyParamsWithHTTPClient creates a new S3BucketLifecycleRuleModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3BucketLifecycleRuleModifyParamsWithHTTPClient(client *http.Client) *S3BucketLifecycleRuleModifyParams {
	return &S3BucketLifecycleRuleModifyParams{
		HTTPClient: client,
	}
}

/*
S3BucketLifecycleRuleModifyParams contains all the parameters to send to the API endpoint

	for the s3 bucket lifecycle rule modify operation.

	Typically these are written to a http.Request.
*/
type S3BucketLifecycleRuleModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.S3BucketLifecycleRule

	/* Name.

	   The name of the lifecycle management rule to be applied on the bucket.
	*/
	Name string

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

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

// WithDefaults hydrates default values in the s3 bucket lifecycle rule modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3BucketLifecycleRuleModifyParams) WithDefaults() *S3BucketLifecycleRuleModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 bucket lifecycle rule modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3BucketLifecycleRuleModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := S3BucketLifecycleRuleModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) WithTimeout(timeout time.Duration) *S3BucketLifecycleRuleModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) WithContext(ctx context.Context) *S3BucketLifecycleRuleModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) WithHTTPClient(client *http.Client) *S3BucketLifecycleRuleModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) WithInfo(info *models.S3BucketLifecycleRule) *S3BucketLifecycleRuleModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) SetInfo(info *models.S3BucketLifecycleRule) {
	o.Info = info
}

// WithName adds the name to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) WithName(name string) *S3BucketLifecycleRuleModifyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) SetName(name string) {
	o.Name = name
}

// WithReturnTimeout adds the returnTimeout to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) WithReturnTimeout(returnTimeout *int64) *S3BucketLifecycleRuleModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithS3BucketUUID adds the s3BucketUUID to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) WithS3BucketUUID(s3BucketUUID string) *S3BucketLifecycleRuleModifyParams {
	o.SetS3BucketUUID(s3BucketUUID)
	return o
}

// SetS3BucketUUID adds the s3BucketUuid to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) SetS3BucketUUID(s3BucketUUID string) {
	o.S3BucketUUID = s3BucketUUID
}

// WithSvmUUID adds the svmUUID to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) WithSvmUUID(svmUUID string) *S3BucketLifecycleRuleModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the s3 bucket lifecycle rule modify params
func (o *S3BucketLifecycleRuleModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *S3BucketLifecycleRuleModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
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