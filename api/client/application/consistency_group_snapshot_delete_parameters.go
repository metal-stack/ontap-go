// Code generated by go-swagger; DO NOT EDIT.

package application

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
)

// NewConsistencyGroupSnapshotDeleteParams creates a new ConsistencyGroupSnapshotDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConsistencyGroupSnapshotDeleteParams() *ConsistencyGroupSnapshotDeleteParams {
	return &ConsistencyGroupSnapshotDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConsistencyGroupSnapshotDeleteParamsWithTimeout creates a new ConsistencyGroupSnapshotDeleteParams object
// with the ability to set a timeout on a request.
func NewConsistencyGroupSnapshotDeleteParamsWithTimeout(timeout time.Duration) *ConsistencyGroupSnapshotDeleteParams {
	return &ConsistencyGroupSnapshotDeleteParams{
		timeout: timeout,
	}
}

// NewConsistencyGroupSnapshotDeleteParamsWithContext creates a new ConsistencyGroupSnapshotDeleteParams object
// with the ability to set a context for a request.
func NewConsistencyGroupSnapshotDeleteParamsWithContext(ctx context.Context) *ConsistencyGroupSnapshotDeleteParams {
	return &ConsistencyGroupSnapshotDeleteParams{
		Context: ctx,
	}
}

// NewConsistencyGroupSnapshotDeleteParamsWithHTTPClient creates a new ConsistencyGroupSnapshotDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewConsistencyGroupSnapshotDeleteParamsWithHTTPClient(client *http.Client) *ConsistencyGroupSnapshotDeleteParams {
	return &ConsistencyGroupSnapshotDeleteParams{
		HTTPClient: client,
	}
}

/*
ConsistencyGroupSnapshotDeleteParams contains all the parameters to send to the API endpoint

	for the consistency group snapshot delete operation.

	Typically these are written to a http.Request.
*/
type ConsistencyGroupSnapshotDeleteParams struct {

	/* ConsistencyGroupUUID.

	   The unique identifier of the Snapshot copy of the consistency group to delete.

	*/
	ConsistencyGroupUUID string

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   Snapshot copy UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the consistency group snapshot delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConsistencyGroupSnapshotDeleteParams) WithDefaults() *ConsistencyGroupSnapshotDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the consistency group snapshot delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConsistencyGroupSnapshotDeleteParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := ConsistencyGroupSnapshotDeleteParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the consistency group snapshot delete params
func (o *ConsistencyGroupSnapshotDeleteParams) WithTimeout(timeout time.Duration) *ConsistencyGroupSnapshotDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the consistency group snapshot delete params
func (o *ConsistencyGroupSnapshotDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the consistency group snapshot delete params
func (o *ConsistencyGroupSnapshotDeleteParams) WithContext(ctx context.Context) *ConsistencyGroupSnapshotDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the consistency group snapshot delete params
func (o *ConsistencyGroupSnapshotDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the consistency group snapshot delete params
func (o *ConsistencyGroupSnapshotDeleteParams) WithHTTPClient(client *http.Client) *ConsistencyGroupSnapshotDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the consistency group snapshot delete params
func (o *ConsistencyGroupSnapshotDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConsistencyGroupUUID adds the consistencyGroupUUID to the consistency group snapshot delete params
func (o *ConsistencyGroupSnapshotDeleteParams) WithConsistencyGroupUUID(consistencyGroupUUID string) *ConsistencyGroupSnapshotDeleteParams {
	o.SetConsistencyGroupUUID(consistencyGroupUUID)
	return o
}

// SetConsistencyGroupUUID adds the consistencyGroupUuid to the consistency group snapshot delete params
func (o *ConsistencyGroupSnapshotDeleteParams) SetConsistencyGroupUUID(consistencyGroupUUID string) {
	o.ConsistencyGroupUUID = consistencyGroupUUID
}

// WithReturnTimeout adds the returnTimeout to the consistency group snapshot delete params
func (o *ConsistencyGroupSnapshotDeleteParams) WithReturnTimeout(returnTimeout *int64) *ConsistencyGroupSnapshotDeleteParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the consistency group snapshot delete params
func (o *ConsistencyGroupSnapshotDeleteParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUID adds the uuid to the consistency group snapshot delete params
func (o *ConsistencyGroupSnapshotDeleteParams) WithUUID(uuid string) *ConsistencyGroupSnapshotDeleteParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the consistency group snapshot delete params
func (o *ConsistencyGroupSnapshotDeleteParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ConsistencyGroupSnapshotDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param consistency_group.uuid
	if err := r.SetPathParam("consistency_group.uuid", o.ConsistencyGroupUUID); err != nil {
		return err
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