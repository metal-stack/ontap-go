// Code generated by go-swagger; DO NOT EDIT.

package snap_lock

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
)

// NewSnaplockLegalHoldDeleteParams creates a new SnaplockLegalHoldDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockLegalHoldDeleteParams() *SnaplockLegalHoldDeleteParams {
	return &SnaplockLegalHoldDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockLegalHoldDeleteParamsWithTimeout creates a new SnaplockLegalHoldDeleteParams object
// with the ability to set a timeout on a request.
func NewSnaplockLegalHoldDeleteParamsWithTimeout(timeout time.Duration) *SnaplockLegalHoldDeleteParams {
	return &SnaplockLegalHoldDeleteParams{
		timeout: timeout,
	}
}

// NewSnaplockLegalHoldDeleteParamsWithContext creates a new SnaplockLegalHoldDeleteParams object
// with the ability to set a context for a request.
func NewSnaplockLegalHoldDeleteParamsWithContext(ctx context.Context) *SnaplockLegalHoldDeleteParams {
	return &SnaplockLegalHoldDeleteParams{
		Context: ctx,
	}
}

// NewSnaplockLegalHoldDeleteParamsWithHTTPClient creates a new SnaplockLegalHoldDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockLegalHoldDeleteParamsWithHTTPClient(client *http.Client) *SnaplockLegalHoldDeleteParams {
	return &SnaplockLegalHoldDeleteParams{
		HTTPClient: client,
	}
}

/*
SnaplockLegalHoldDeleteParams contains all the parameters to send to the API endpoint

	for the snaplock legal hold delete operation.

	Typically these are written to a http.Request.
*/
type SnaplockLegalHoldDeleteParams struct {

	/* ID.

	   Litigation ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock legal hold delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldDeleteParams) WithDefaults() *SnaplockLegalHoldDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock legal hold delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snaplock legal hold delete params
func (o *SnaplockLegalHoldDeleteParams) WithTimeout(timeout time.Duration) *SnaplockLegalHoldDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock legal hold delete params
func (o *SnaplockLegalHoldDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock legal hold delete params
func (o *SnaplockLegalHoldDeleteParams) WithContext(ctx context.Context) *SnaplockLegalHoldDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock legal hold delete params
func (o *SnaplockLegalHoldDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock legal hold delete params
func (o *SnaplockLegalHoldDeleteParams) WithHTTPClient(client *http.Client) *SnaplockLegalHoldDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock legal hold delete params
func (o *SnaplockLegalHoldDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the snaplock legal hold delete params
func (o *SnaplockLegalHoldDeleteParams) WithID(id string) *SnaplockLegalHoldDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the snaplock legal hold delete params
func (o *SnaplockLegalHoldDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockLegalHoldDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
