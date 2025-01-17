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

// NewIPSubnetModifyParams creates a new IPSubnetModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIPSubnetModifyParams() *IPSubnetModifyParams {
	return &IPSubnetModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIPSubnetModifyParamsWithTimeout creates a new IPSubnetModifyParams object
// with the ability to set a timeout on a request.
func NewIPSubnetModifyParamsWithTimeout(timeout time.Duration) *IPSubnetModifyParams {
	return &IPSubnetModifyParams{
		timeout: timeout,
	}
}

// NewIPSubnetModifyParamsWithContext creates a new IPSubnetModifyParams object
// with the ability to set a context for a request.
func NewIPSubnetModifyParamsWithContext(ctx context.Context) *IPSubnetModifyParams {
	return &IPSubnetModifyParams{
		Context: ctx,
	}
}

// NewIPSubnetModifyParamsWithHTTPClient creates a new IPSubnetModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewIPSubnetModifyParamsWithHTTPClient(client *http.Client) *IPSubnetModifyParams {
	return &IPSubnetModifyParams{
		HTTPClient: client,
	}
}

/*
IPSubnetModifyParams contains all the parameters to send to the API endpoint

	for the ip subnet modify operation.

	Typically these are written to a http.Request.
*/
type IPSubnetModifyParams struct {

	// Info.
	Info *models.IPSubnet

	/* UUID.

	   IP subnet UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ip subnet modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IPSubnetModifyParams) WithDefaults() *IPSubnetModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ip subnet modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IPSubnetModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ip subnet modify params
func (o *IPSubnetModifyParams) WithTimeout(timeout time.Duration) *IPSubnetModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ip subnet modify params
func (o *IPSubnetModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ip subnet modify params
func (o *IPSubnetModifyParams) WithContext(ctx context.Context) *IPSubnetModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ip subnet modify params
func (o *IPSubnetModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ip subnet modify params
func (o *IPSubnetModifyParams) WithHTTPClient(client *http.Client) *IPSubnetModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ip subnet modify params
func (o *IPSubnetModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the ip subnet modify params
func (o *IPSubnetModifyParams) WithInfo(info *models.IPSubnet) *IPSubnetModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ip subnet modify params
func (o *IPSubnetModifyParams) SetInfo(info *models.IPSubnet) {
	o.Info = info
}

// WithUUID adds the uuid to the ip subnet modify params
func (o *IPSubnetModifyParams) WithUUID(uuid string) *IPSubnetModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the ip subnet modify params
func (o *IPSubnetModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *IPSubnetModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
