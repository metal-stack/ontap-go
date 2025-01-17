// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewSnapshotPolicyDeleteParams creates a new SnapshotPolicyDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapshotPolicyDeleteParams() *SnapshotPolicyDeleteParams {
	return &SnapshotPolicyDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotPolicyDeleteParamsWithTimeout creates a new SnapshotPolicyDeleteParams object
// with the ability to set a timeout on a request.
func NewSnapshotPolicyDeleteParamsWithTimeout(timeout time.Duration) *SnapshotPolicyDeleteParams {
	return &SnapshotPolicyDeleteParams{
		timeout: timeout,
	}
}

// NewSnapshotPolicyDeleteParamsWithContext creates a new SnapshotPolicyDeleteParams object
// with the ability to set a context for a request.
func NewSnapshotPolicyDeleteParamsWithContext(ctx context.Context) *SnapshotPolicyDeleteParams {
	return &SnapshotPolicyDeleteParams{
		Context: ctx,
	}
}

// NewSnapshotPolicyDeleteParamsWithHTTPClient creates a new SnapshotPolicyDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapshotPolicyDeleteParamsWithHTTPClient(client *http.Client) *SnapshotPolicyDeleteParams {
	return &SnapshotPolicyDeleteParams{
		HTTPClient: client,
	}
}

/*
SnapshotPolicyDeleteParams contains all the parameters to send to the API endpoint

	for the snapshot policy delete operation.

	Typically these are written to a http.Request.
*/
type SnapshotPolicyDeleteParams struct {

	/* UUID.

	   Snapshot copy policy UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapshot policy delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotPolicyDeleteParams) WithDefaults() *SnapshotPolicyDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapshot policy delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotPolicyDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snapshot policy delete params
func (o *SnapshotPolicyDeleteParams) WithTimeout(timeout time.Duration) *SnapshotPolicyDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot policy delete params
func (o *SnapshotPolicyDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot policy delete params
func (o *SnapshotPolicyDeleteParams) WithContext(ctx context.Context) *SnapshotPolicyDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot policy delete params
func (o *SnapshotPolicyDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot policy delete params
func (o *SnapshotPolicyDeleteParams) WithHTTPClient(client *http.Client) *SnapshotPolicyDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot policy delete params
func (o *SnapshotPolicyDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the snapshot policy delete params
func (o *SnapshotPolicyDeleteParams) WithUUID(uuid string) *SnapshotPolicyDeleteParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the snapshot policy delete params
func (o *SnapshotPolicyDeleteParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotPolicyDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
