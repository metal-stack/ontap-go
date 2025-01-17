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

	"github.com/metal-stack/ontap-go/api/models"
)

// NewLocalCifsGroupModifyParams creates a new LocalCifsGroupModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocalCifsGroupModifyParams() *LocalCifsGroupModifyParams {
	return &LocalCifsGroupModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocalCifsGroupModifyParamsWithTimeout creates a new LocalCifsGroupModifyParams object
// with the ability to set a timeout on a request.
func NewLocalCifsGroupModifyParamsWithTimeout(timeout time.Duration) *LocalCifsGroupModifyParams {
	return &LocalCifsGroupModifyParams{
		timeout: timeout,
	}
}

// NewLocalCifsGroupModifyParamsWithContext creates a new LocalCifsGroupModifyParams object
// with the ability to set a context for a request.
func NewLocalCifsGroupModifyParamsWithContext(ctx context.Context) *LocalCifsGroupModifyParams {
	return &LocalCifsGroupModifyParams{
		Context: ctx,
	}
}

// NewLocalCifsGroupModifyParamsWithHTTPClient creates a new LocalCifsGroupModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocalCifsGroupModifyParamsWithHTTPClient(client *http.Client) *LocalCifsGroupModifyParams {
	return &LocalCifsGroupModifyParams{
		HTTPClient: client,
	}
}

/*
LocalCifsGroupModifyParams contains all the parameters to send to the API endpoint

	for the local cifs group modify operation.

	Typically these are written to a http.Request.
*/
type LocalCifsGroupModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.LocalCifsGroup

	/* Sid.

	   Local group SID
	*/
	Sid string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the local cifs group modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalCifsGroupModifyParams) WithDefaults() *LocalCifsGroupModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the local cifs group modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalCifsGroupModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the local cifs group modify params
func (o *LocalCifsGroupModifyParams) WithTimeout(timeout time.Duration) *LocalCifsGroupModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the local cifs group modify params
func (o *LocalCifsGroupModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the local cifs group modify params
func (o *LocalCifsGroupModifyParams) WithContext(ctx context.Context) *LocalCifsGroupModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the local cifs group modify params
func (o *LocalCifsGroupModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the local cifs group modify params
func (o *LocalCifsGroupModifyParams) WithHTTPClient(client *http.Client) *LocalCifsGroupModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the local cifs group modify params
func (o *LocalCifsGroupModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the local cifs group modify params
func (o *LocalCifsGroupModifyParams) WithInfo(info *models.LocalCifsGroup) *LocalCifsGroupModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the local cifs group modify params
func (o *LocalCifsGroupModifyParams) SetInfo(info *models.LocalCifsGroup) {
	o.Info = info
}

// WithSid adds the sid to the local cifs group modify params
func (o *LocalCifsGroupModifyParams) WithSid(sid string) *LocalCifsGroupModifyParams {
	o.SetSid(sid)
	return o
}

// SetSid adds the sid to the local cifs group modify params
func (o *LocalCifsGroupModifyParams) SetSid(sid string) {
	o.Sid = sid
}

// WithSvmUUID adds the svmUUID to the local cifs group modify params
func (o *LocalCifsGroupModifyParams) WithSvmUUID(svmUUID string) *LocalCifsGroupModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the local cifs group modify params
func (o *LocalCifsGroupModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *LocalCifsGroupModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param sid
	if err := r.SetPathParam("sid", o.Sid); err != nil {
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
