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

// NewS3GroupModifyParams creates a new S3GroupModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3GroupModifyParams() *S3GroupModifyParams {
	return &S3GroupModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3GroupModifyParamsWithTimeout creates a new S3GroupModifyParams object
// with the ability to set a timeout on a request.
func NewS3GroupModifyParamsWithTimeout(timeout time.Duration) *S3GroupModifyParams {
	return &S3GroupModifyParams{
		timeout: timeout,
	}
}

// NewS3GroupModifyParamsWithContext creates a new S3GroupModifyParams object
// with the ability to set a context for a request.
func NewS3GroupModifyParamsWithContext(ctx context.Context) *S3GroupModifyParams {
	return &S3GroupModifyParams{
		Context: ctx,
	}
}

// NewS3GroupModifyParamsWithHTTPClient creates a new S3GroupModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3GroupModifyParamsWithHTTPClient(client *http.Client) *S3GroupModifyParams {
	return &S3GroupModifyParams{
		HTTPClient: client,
	}
}

/*
S3GroupModifyParams contains all the parameters to send to the API endpoint

	for the s3 group modify operation.

	Typically these are written to a http.Request.
*/
type S3GroupModifyParams struct {

	/* ID.

	   Group identifier that identifies the unique group.
	*/
	ID int64

	/* Info.

	   Info specification
	*/
	Info *models.S3Group

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 group modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3GroupModifyParams) WithDefaults() *S3GroupModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 group modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3GroupModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the s3 group modify params
func (o *S3GroupModifyParams) WithTimeout(timeout time.Duration) *S3GroupModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 group modify params
func (o *S3GroupModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 group modify params
func (o *S3GroupModifyParams) WithContext(ctx context.Context) *S3GroupModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 group modify params
func (o *S3GroupModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 group modify params
func (o *S3GroupModifyParams) WithHTTPClient(client *http.Client) *S3GroupModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 group modify params
func (o *S3GroupModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the s3 group modify params
func (o *S3GroupModifyParams) WithID(id int64) *S3GroupModifyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the s3 group modify params
func (o *S3GroupModifyParams) SetID(id int64) {
	o.ID = id
}

// WithInfo adds the info to the s3 group modify params
func (o *S3GroupModifyParams) WithInfo(info *models.S3Group) *S3GroupModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the s3 group modify params
func (o *S3GroupModifyParams) SetInfo(info *models.S3Group) {
	o.Info = info
}

// WithSvmUUID adds the svmUUID to the s3 group modify params
func (o *S3GroupModifyParams) WithSvmUUID(svmUUID string) *S3GroupModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the s3 group modify params
func (o *S3GroupModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *S3GroupModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
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