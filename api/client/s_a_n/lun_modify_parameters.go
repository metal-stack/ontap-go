// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewLunModifyParams creates a new LunModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLunModifyParams() *LunModifyParams {
	return &LunModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLunModifyParamsWithTimeout creates a new LunModifyParams object
// with the ability to set a timeout on a request.
func NewLunModifyParamsWithTimeout(timeout time.Duration) *LunModifyParams {
	return &LunModifyParams{
		timeout: timeout,
	}
}

// NewLunModifyParamsWithContext creates a new LunModifyParams object
// with the ability to set a context for a request.
func NewLunModifyParamsWithContext(ctx context.Context) *LunModifyParams {
	return &LunModifyParams{
		Context: ctx,
	}
}

// NewLunModifyParamsWithHTTPClient creates a new LunModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewLunModifyParamsWithHTTPClient(client *http.Client) *LunModifyParams {
	return &LunModifyParams{
		HTTPClient: client,
	}
}

/*
LunModifyParams contains all the parameters to send to the API endpoint

	for the lun modify operation.

	Typically these are written to a http.Request.
*/
type LunModifyParams struct {

	/* DataOffset.

	     The offset, in bytes, at which to begin writing LUN data.<br/>
	LUN data write requests are distinguished by the header entry `Content-Type: multipart/form-data`. When this header entry is provided, query parameter `data.offset` is required and used to specify the location within the LUN at which to write the data; no other query parameters are allowed. The request body must be `multipart/form-data` content with exactly one form entry containing the data to write. The content type entry of the form data is ignored and always treated as `application/octet-stream`. Writes are limited to one megabyte (1MB) per request.


	     Format: int64
	*/
	DataOffset *int64

	/* Info.

	     The new property values for the LUN.<br/>
	Either `info` or `data` must be supplied.

	*/
	Info *models.Lun

	/* UUID.

	   The unique identifier of the LUN to retrieve.

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lun modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunModifyParams) WithDefaults() *LunModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lun modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lun modify params
func (o *LunModifyParams) WithTimeout(timeout time.Duration) *LunModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lun modify params
func (o *LunModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lun modify params
func (o *LunModifyParams) WithContext(ctx context.Context) *LunModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lun modify params
func (o *LunModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lun modify params
func (o *LunModifyParams) WithHTTPClient(client *http.Client) *LunModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lun modify params
func (o *LunModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataOffset adds the dataOffset to the lun modify params
func (o *LunModifyParams) WithDataOffset(dataOffset *int64) *LunModifyParams {
	o.SetDataOffset(dataOffset)
	return o
}

// SetDataOffset adds the dataOffset to the lun modify params
func (o *LunModifyParams) SetDataOffset(dataOffset *int64) {
	o.DataOffset = dataOffset
}

// WithInfo adds the info to the lun modify params
func (o *LunModifyParams) WithInfo(info *models.Lun) *LunModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the lun modify params
func (o *LunModifyParams) SetInfo(info *models.Lun) {
	o.Info = info
}

// WithUUID adds the uuid to the lun modify params
func (o *LunModifyParams) WithUUID(uuid string) *LunModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the lun modify params
func (o *LunModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *LunModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DataOffset != nil {

		// query param data.offset
		var qrDataOffset int64

		if o.DataOffset != nil {
			qrDataOffset = *o.DataOffset
		}
		qDataOffset := swag.FormatInt64(qrDataOffset)
		if qDataOffset != "" {

			if err := r.SetQueryParam("data.offset", qDataOffset); err != nil {
				return err
			}
		}
	}
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