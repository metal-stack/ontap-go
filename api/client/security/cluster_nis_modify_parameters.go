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

	"github.com/metal-stack/ontap-go/api/models"
)

// NewClusterNisModifyParams creates a new ClusterNisModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterNisModifyParams() *ClusterNisModifyParams {
	return &ClusterNisModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterNisModifyParamsWithTimeout creates a new ClusterNisModifyParams object
// with the ability to set a timeout on a request.
func NewClusterNisModifyParamsWithTimeout(timeout time.Duration) *ClusterNisModifyParams {
	return &ClusterNisModifyParams{
		timeout: timeout,
	}
}

// NewClusterNisModifyParamsWithContext creates a new ClusterNisModifyParams object
// with the ability to set a context for a request.
func NewClusterNisModifyParamsWithContext(ctx context.Context) *ClusterNisModifyParams {
	return &ClusterNisModifyParams{
		Context: ctx,
	}
}

// NewClusterNisModifyParamsWithHTTPClient creates a new ClusterNisModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterNisModifyParamsWithHTTPClient(client *http.Client) *ClusterNisModifyParams {
	return &ClusterNisModifyParams{
		HTTPClient: client,
	}
}

/*
ClusterNisModifyParams contains all the parameters to send to the API endpoint

	for the cluster nis modify operation.

	Typically these are written to a http.Request.
*/
type ClusterNisModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.ClusterNisService

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster nis modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterNisModifyParams) WithDefaults() *ClusterNisModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster nis modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterNisModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster nis modify params
func (o *ClusterNisModifyParams) WithTimeout(timeout time.Duration) *ClusterNisModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster nis modify params
func (o *ClusterNisModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster nis modify params
func (o *ClusterNisModifyParams) WithContext(ctx context.Context) *ClusterNisModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster nis modify params
func (o *ClusterNisModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster nis modify params
func (o *ClusterNisModifyParams) WithHTTPClient(client *http.Client) *ClusterNisModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster nis modify params
func (o *ClusterNisModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the cluster nis modify params
func (o *ClusterNisModifyParams) WithInfo(info *models.ClusterNisService) *ClusterNisModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cluster nis modify params
func (o *ClusterNisModifyParams) SetInfo(info *models.ClusterNisService) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterNisModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}