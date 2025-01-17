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

// NewVscanOnDemandModifyParams creates a new VscanOnDemandModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVscanOnDemandModifyParams() *VscanOnDemandModifyParams {
	return &VscanOnDemandModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVscanOnDemandModifyParamsWithTimeout creates a new VscanOnDemandModifyParams object
// with the ability to set a timeout on a request.
func NewVscanOnDemandModifyParamsWithTimeout(timeout time.Duration) *VscanOnDemandModifyParams {
	return &VscanOnDemandModifyParams{
		timeout: timeout,
	}
}

// NewVscanOnDemandModifyParamsWithContext creates a new VscanOnDemandModifyParams object
// with the ability to set a context for a request.
func NewVscanOnDemandModifyParamsWithContext(ctx context.Context) *VscanOnDemandModifyParams {
	return &VscanOnDemandModifyParams{
		Context: ctx,
	}
}

// NewVscanOnDemandModifyParamsWithHTTPClient creates a new VscanOnDemandModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewVscanOnDemandModifyParamsWithHTTPClient(client *http.Client) *VscanOnDemandModifyParams {
	return &VscanOnDemandModifyParams{
		HTTPClient: client,
	}
}

/*
VscanOnDemandModifyParams contains all the parameters to send to the API endpoint

	for the vscan on demand modify operation.

	Typically these are written to a http.Request.
*/
type VscanOnDemandModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.VscanOnDemand

	// Name.
	Name string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vscan on demand modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanOnDemandModifyParams) WithDefaults() *VscanOnDemandModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vscan on demand modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanOnDemandModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vscan on demand modify params
func (o *VscanOnDemandModifyParams) WithTimeout(timeout time.Duration) *VscanOnDemandModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vscan on demand modify params
func (o *VscanOnDemandModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vscan on demand modify params
func (o *VscanOnDemandModifyParams) WithContext(ctx context.Context) *VscanOnDemandModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vscan on demand modify params
func (o *VscanOnDemandModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vscan on demand modify params
func (o *VscanOnDemandModifyParams) WithHTTPClient(client *http.Client) *VscanOnDemandModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vscan on demand modify params
func (o *VscanOnDemandModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the vscan on demand modify params
func (o *VscanOnDemandModifyParams) WithInfo(info *models.VscanOnDemand) *VscanOnDemandModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the vscan on demand modify params
func (o *VscanOnDemandModifyParams) SetInfo(info *models.VscanOnDemand) {
	o.Info = info
}

// WithName adds the name to the vscan on demand modify params
func (o *VscanOnDemandModifyParams) WithName(name string) *VscanOnDemandModifyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the vscan on demand modify params
func (o *VscanOnDemandModifyParams) SetName(name string) {
	o.Name = name
}

// WithSvmUUID adds the svmUUID to the vscan on demand modify params
func (o *VscanOnDemandModifyParams) WithSvmUUID(svmUUID string) *VscanOnDemandModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the vscan on demand modify params
func (o *VscanOnDemandModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *VscanOnDemandModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}