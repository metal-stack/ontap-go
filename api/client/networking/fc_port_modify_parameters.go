// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// NewFcPortModifyParams creates a new FcPortModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFcPortModifyParams() *FcPortModifyParams {
	return &FcPortModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFcPortModifyParamsWithTimeout creates a new FcPortModifyParams object
// with the ability to set a timeout on a request.
func NewFcPortModifyParamsWithTimeout(timeout time.Duration) *FcPortModifyParams {
	return &FcPortModifyParams{
		timeout: timeout,
	}
}

// NewFcPortModifyParamsWithContext creates a new FcPortModifyParams object
// with the ability to set a context for a request.
func NewFcPortModifyParamsWithContext(ctx context.Context) *FcPortModifyParams {
	return &FcPortModifyParams{
		Context: ctx,
	}
}

// NewFcPortModifyParamsWithHTTPClient creates a new FcPortModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewFcPortModifyParamsWithHTTPClient(client *http.Client) *FcPortModifyParams {
	return &FcPortModifyParams{
		HTTPClient: client,
	}
}

/*
FcPortModifyParams contains all the parameters to send to the API endpoint

	for the fc port modify operation.

	Typically these are written to a http.Request.
*/
type FcPortModifyParams struct {

	/* Info.

	   The new property values for the FC port.

	*/
	Info *models.FcPort

	/* UUID.

	   The unique identifier for the FC port.

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fc port modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcPortModifyParams) WithDefaults() *FcPortModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fc port modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcPortModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fc port modify params
func (o *FcPortModifyParams) WithTimeout(timeout time.Duration) *FcPortModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fc port modify params
func (o *FcPortModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fc port modify params
func (o *FcPortModifyParams) WithContext(ctx context.Context) *FcPortModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fc port modify params
func (o *FcPortModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fc port modify params
func (o *FcPortModifyParams) WithHTTPClient(client *http.Client) *FcPortModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fc port modify params
func (o *FcPortModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the fc port modify params
func (o *FcPortModifyParams) WithInfo(info *models.FcPort) *FcPortModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the fc port modify params
func (o *FcPortModifyParams) SetInfo(info *models.FcPort) {
	o.Info = info
}

// WithUUID adds the uuid to the fc port modify params
func (o *FcPortModifyParams) WithUUID(uuid string) *FcPortModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the fc port modify params
func (o *FcPortModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *FcPortModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
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
