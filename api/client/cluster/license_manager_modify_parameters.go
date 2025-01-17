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

	"github.com/metal-stack/ontap-go/api/models"
)

// NewLicenseManagerModifyParams creates a new LicenseManagerModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLicenseManagerModifyParams() *LicenseManagerModifyParams {
	return &LicenseManagerModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLicenseManagerModifyParamsWithTimeout creates a new LicenseManagerModifyParams object
// with the ability to set a timeout on a request.
func NewLicenseManagerModifyParamsWithTimeout(timeout time.Duration) *LicenseManagerModifyParams {
	return &LicenseManagerModifyParams{
		timeout: timeout,
	}
}

// NewLicenseManagerModifyParamsWithContext creates a new LicenseManagerModifyParams object
// with the ability to set a context for a request.
func NewLicenseManagerModifyParamsWithContext(ctx context.Context) *LicenseManagerModifyParams {
	return &LicenseManagerModifyParams{
		Context: ctx,
	}
}

// NewLicenseManagerModifyParamsWithHTTPClient creates a new LicenseManagerModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewLicenseManagerModifyParamsWithHTTPClient(client *http.Client) *LicenseManagerModifyParams {
	return &LicenseManagerModifyParams{
		HTTPClient: client,
	}
}

/*
LicenseManagerModifyParams contains all the parameters to send to the API endpoint

	for the license manager modify operation.

	Typically these are written to a http.Request.
*/
type LicenseManagerModifyParams struct {

	/* Info.

	   Request specification
	*/
	Info *models.LicenseManager

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	// UUID.
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the license manager modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LicenseManagerModifyParams) WithDefaults() *LicenseManagerModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the license manager modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LicenseManagerModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := LicenseManagerModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the license manager modify params
func (o *LicenseManagerModifyParams) WithTimeout(timeout time.Duration) *LicenseManagerModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the license manager modify params
func (o *LicenseManagerModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the license manager modify params
func (o *LicenseManagerModifyParams) WithContext(ctx context.Context) *LicenseManagerModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the license manager modify params
func (o *LicenseManagerModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the license manager modify params
func (o *LicenseManagerModifyParams) WithHTTPClient(client *http.Client) *LicenseManagerModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the license manager modify params
func (o *LicenseManagerModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the license manager modify params
func (o *LicenseManagerModifyParams) WithInfo(info *models.LicenseManager) *LicenseManagerModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the license manager modify params
func (o *LicenseManagerModifyParams) SetInfo(info *models.LicenseManager) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the license manager modify params
func (o *LicenseManagerModifyParams) WithReturnTimeout(returnTimeout *int64) *LicenseManagerModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the license manager modify params
func (o *LicenseManagerModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUID adds the uuid to the license manager modify params
func (o *LicenseManagerModifyParams) WithUUID(uuid string) *LicenseManagerModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the license manager modify params
func (o *LicenseManagerModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *LicenseManagerModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
