// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewClusterAccountAdProxyDeleteParams creates a new ClusterAccountAdProxyDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterAccountAdProxyDeleteParams() *ClusterAccountAdProxyDeleteParams {
	return &ClusterAccountAdProxyDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterAccountAdProxyDeleteParamsWithTimeout creates a new ClusterAccountAdProxyDeleteParams object
// with the ability to set a timeout on a request.
func NewClusterAccountAdProxyDeleteParamsWithTimeout(timeout time.Duration) *ClusterAccountAdProxyDeleteParams {
	return &ClusterAccountAdProxyDeleteParams{
		timeout: timeout,
	}
}

// NewClusterAccountAdProxyDeleteParamsWithContext creates a new ClusterAccountAdProxyDeleteParams object
// with the ability to set a context for a request.
func NewClusterAccountAdProxyDeleteParamsWithContext(ctx context.Context) *ClusterAccountAdProxyDeleteParams {
	return &ClusterAccountAdProxyDeleteParams{
		Context: ctx,
	}
}

// NewClusterAccountAdProxyDeleteParamsWithHTTPClient creates a new ClusterAccountAdProxyDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterAccountAdProxyDeleteParamsWithHTTPClient(client *http.Client) *ClusterAccountAdProxyDeleteParams {
	return &ClusterAccountAdProxyDeleteParams{
		HTTPClient: client,
	}
}

/*
ClusterAccountAdProxyDeleteParams contains all the parameters to send to the API endpoint

	for the cluster account ad proxy delete operation.

	Typically these are written to a http.Request.
*/
type ClusterAccountAdProxyDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster account ad proxy delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterAccountAdProxyDeleteParams) WithDefaults() *ClusterAccountAdProxyDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster account ad proxy delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterAccountAdProxyDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster account ad proxy delete params
func (o *ClusterAccountAdProxyDeleteParams) WithTimeout(timeout time.Duration) *ClusterAccountAdProxyDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster account ad proxy delete params
func (o *ClusterAccountAdProxyDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster account ad proxy delete params
func (o *ClusterAccountAdProxyDeleteParams) WithContext(ctx context.Context) *ClusterAccountAdProxyDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster account ad proxy delete params
func (o *ClusterAccountAdProxyDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster account ad proxy delete params
func (o *ClusterAccountAdProxyDeleteParams) WithHTTPClient(client *http.Client) *ClusterAccountAdProxyDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster account ad proxy delete params
func (o *ClusterAccountAdProxyDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterAccountAdProxyDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}