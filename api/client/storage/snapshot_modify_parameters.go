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
	"github.com/go-openapi/swag"

	"github.com/metal-stack/ontap-go/api/models"
)

// NewSnapshotModifyParams creates a new SnapshotModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapshotModifyParams() *SnapshotModifyParams {
	return &SnapshotModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotModifyParamsWithTimeout creates a new SnapshotModifyParams object
// with the ability to set a timeout on a request.
func NewSnapshotModifyParamsWithTimeout(timeout time.Duration) *SnapshotModifyParams {
	return &SnapshotModifyParams{
		timeout: timeout,
	}
}

// NewSnapshotModifyParamsWithContext creates a new SnapshotModifyParams object
// with the ability to set a context for a request.
func NewSnapshotModifyParamsWithContext(ctx context.Context) *SnapshotModifyParams {
	return &SnapshotModifyParams{
		Context: ctx,
	}
}

// NewSnapshotModifyParamsWithHTTPClient creates a new SnapshotModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapshotModifyParamsWithHTTPClient(client *http.Client) *SnapshotModifyParams {
	return &SnapshotModifyParams{
		HTTPClient: client,
	}
}

/*
SnapshotModifyParams contains all the parameters to send to the API endpoint

	for the snapshot modify operation.

	Typically these are written to a http.Request.
*/
type SnapshotModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.Snapshot

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   Snapshot copy UUID
	*/
	UUID string

	/* VolumeUUID.

	   Volume UUID
	*/
	VolumeUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapshot modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotModifyParams) WithDefaults() *SnapshotModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapshot modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := SnapshotModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snapshot modify params
func (o *SnapshotModifyParams) WithTimeout(timeout time.Duration) *SnapshotModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot modify params
func (o *SnapshotModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot modify params
func (o *SnapshotModifyParams) WithContext(ctx context.Context) *SnapshotModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot modify params
func (o *SnapshotModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot modify params
func (o *SnapshotModifyParams) WithHTTPClient(client *http.Client) *SnapshotModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot modify params
func (o *SnapshotModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the snapshot modify params
func (o *SnapshotModifyParams) WithInfo(info *models.Snapshot) *SnapshotModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the snapshot modify params
func (o *SnapshotModifyParams) SetInfo(info *models.Snapshot) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the snapshot modify params
func (o *SnapshotModifyParams) WithReturnTimeout(returnTimeout *int64) *SnapshotModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the snapshot modify params
func (o *SnapshotModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUID adds the uuid to the snapshot modify params
func (o *SnapshotModifyParams) WithUUID(uuid string) *SnapshotModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the snapshot modify params
func (o *SnapshotModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WithVolumeUUID adds the volumeUUID to the snapshot modify params
func (o *SnapshotModifyParams) WithVolumeUUID(volumeUUID string) *SnapshotModifyParams {
	o.SetVolumeUUID(volumeUUID)
	return o
}

// SetVolumeUUID adds the volumeUuid to the snapshot modify params
func (o *SnapshotModifyParams) SetVolumeUUID(volumeUUID string) {
	o.VolumeUUID = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param volume.uuid
	if err := r.SetPathParam("volume.uuid", o.VolumeUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
