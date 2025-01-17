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

	"github.com/metal-stack/ontap-go/api/models"
)

// NewSnapshotPolicyScheduleModifyParams creates a new SnapshotPolicyScheduleModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapshotPolicyScheduleModifyParams() *SnapshotPolicyScheduleModifyParams {
	return &SnapshotPolicyScheduleModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotPolicyScheduleModifyParamsWithTimeout creates a new SnapshotPolicyScheduleModifyParams object
// with the ability to set a timeout on a request.
func NewSnapshotPolicyScheduleModifyParamsWithTimeout(timeout time.Duration) *SnapshotPolicyScheduleModifyParams {
	return &SnapshotPolicyScheduleModifyParams{
		timeout: timeout,
	}
}

// NewSnapshotPolicyScheduleModifyParamsWithContext creates a new SnapshotPolicyScheduleModifyParams object
// with the ability to set a context for a request.
func NewSnapshotPolicyScheduleModifyParamsWithContext(ctx context.Context) *SnapshotPolicyScheduleModifyParams {
	return &SnapshotPolicyScheduleModifyParams{
		Context: ctx,
	}
}

// NewSnapshotPolicyScheduleModifyParamsWithHTTPClient creates a new SnapshotPolicyScheduleModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapshotPolicyScheduleModifyParamsWithHTTPClient(client *http.Client) *SnapshotPolicyScheduleModifyParams {
	return &SnapshotPolicyScheduleModifyParams{
		HTTPClient: client,
	}
}

/*
SnapshotPolicyScheduleModifyParams contains all the parameters to send to the API endpoint

	for the snapshot policy schedule modify operation.

	Typically these are written to a http.Request.
*/
type SnapshotPolicyScheduleModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.SnapshotPolicySchedule

	/* ScheduleUUID.

	   Snapshot copy policy schedule UUID
	*/
	ScheduleUUID string

	/* SnapshotPolicyUUID.

	   Snapshot copy policy UUID
	*/
	SnapshotPolicyUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapshot policy schedule modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotPolicyScheduleModifyParams) WithDefaults() *SnapshotPolicyScheduleModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapshot policy schedule modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotPolicyScheduleModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snapshot policy schedule modify params
func (o *SnapshotPolicyScheduleModifyParams) WithTimeout(timeout time.Duration) *SnapshotPolicyScheduleModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot policy schedule modify params
func (o *SnapshotPolicyScheduleModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot policy schedule modify params
func (o *SnapshotPolicyScheduleModifyParams) WithContext(ctx context.Context) *SnapshotPolicyScheduleModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot policy schedule modify params
func (o *SnapshotPolicyScheduleModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot policy schedule modify params
func (o *SnapshotPolicyScheduleModifyParams) WithHTTPClient(client *http.Client) *SnapshotPolicyScheduleModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot policy schedule modify params
func (o *SnapshotPolicyScheduleModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the snapshot policy schedule modify params
func (o *SnapshotPolicyScheduleModifyParams) WithInfo(info *models.SnapshotPolicySchedule) *SnapshotPolicyScheduleModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the snapshot policy schedule modify params
func (o *SnapshotPolicyScheduleModifyParams) SetInfo(info *models.SnapshotPolicySchedule) {
	o.Info = info
}

// WithScheduleUUID adds the scheduleUUID to the snapshot policy schedule modify params
func (o *SnapshotPolicyScheduleModifyParams) WithScheduleUUID(scheduleUUID string) *SnapshotPolicyScheduleModifyParams {
	o.SetScheduleUUID(scheduleUUID)
	return o
}

// SetScheduleUUID adds the scheduleUuid to the snapshot policy schedule modify params
func (o *SnapshotPolicyScheduleModifyParams) SetScheduleUUID(scheduleUUID string) {
	o.ScheduleUUID = scheduleUUID
}

// WithSnapshotPolicyUUID adds the snapshotPolicyUUID to the snapshot policy schedule modify params
func (o *SnapshotPolicyScheduleModifyParams) WithSnapshotPolicyUUID(snapshotPolicyUUID string) *SnapshotPolicyScheduleModifyParams {
	o.SetSnapshotPolicyUUID(snapshotPolicyUUID)
	return o
}

// SetSnapshotPolicyUUID adds the snapshotPolicyUuid to the snapshot policy schedule modify params
func (o *SnapshotPolicyScheduleModifyParams) SetSnapshotPolicyUUID(snapshotPolicyUUID string) {
	o.SnapshotPolicyUUID = snapshotPolicyUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotPolicyScheduleModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param schedule.uuid
	if err := r.SetPathParam("schedule.uuid", o.ScheduleUUID); err != nil {
		return err
	}

	// path param snapshot_policy.uuid
	if err := r.SetPathParam("snapshot_policy.uuid", o.SnapshotPolicyUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
