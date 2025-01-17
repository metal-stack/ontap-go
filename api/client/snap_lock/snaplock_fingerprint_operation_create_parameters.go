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
	"github.com/go-openapi/swag"

	"github.com/metal-stack/ontap-go/api/models"
)

// NewSnaplockFingerprintOperationCreateParams creates a new SnaplockFingerprintOperationCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockFingerprintOperationCreateParams() *SnaplockFingerprintOperationCreateParams {
	return &SnaplockFingerprintOperationCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockFingerprintOperationCreateParamsWithTimeout creates a new SnaplockFingerprintOperationCreateParams object
// with the ability to set a timeout on a request.
func NewSnaplockFingerprintOperationCreateParamsWithTimeout(timeout time.Duration) *SnaplockFingerprintOperationCreateParams {
	return &SnaplockFingerprintOperationCreateParams{
		timeout: timeout,
	}
}

// NewSnaplockFingerprintOperationCreateParamsWithContext creates a new SnaplockFingerprintOperationCreateParams object
// with the ability to set a context for a request.
func NewSnaplockFingerprintOperationCreateParamsWithContext(ctx context.Context) *SnaplockFingerprintOperationCreateParams {
	return &SnaplockFingerprintOperationCreateParams{
		Context: ctx,
	}
}

// NewSnaplockFingerprintOperationCreateParamsWithHTTPClient creates a new SnaplockFingerprintOperationCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockFingerprintOperationCreateParamsWithHTTPClient(client *http.Client) *SnaplockFingerprintOperationCreateParams {
	return &SnaplockFingerprintOperationCreateParams{
		HTTPClient: client,
	}
}

/*
SnaplockFingerprintOperationCreateParams contains all the parameters to send to the API endpoint

	for the snaplock fingerprint operation create operation.

	Typically these are written to a http.Request.
*/
type SnaplockFingerprintOperationCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.SnaplockFileFingerprint

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock fingerprint operation create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockFingerprintOperationCreateParams) WithDefaults() *SnaplockFingerprintOperationCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock fingerprint operation create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockFingerprintOperationCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := SnaplockFingerprintOperationCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snaplock fingerprint operation create params
func (o *SnaplockFingerprintOperationCreateParams) WithTimeout(timeout time.Duration) *SnaplockFingerprintOperationCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock fingerprint operation create params
func (o *SnaplockFingerprintOperationCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock fingerprint operation create params
func (o *SnaplockFingerprintOperationCreateParams) WithContext(ctx context.Context) *SnaplockFingerprintOperationCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock fingerprint operation create params
func (o *SnaplockFingerprintOperationCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock fingerprint operation create params
func (o *SnaplockFingerprintOperationCreateParams) WithHTTPClient(client *http.Client) *SnaplockFingerprintOperationCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock fingerprint operation create params
func (o *SnaplockFingerprintOperationCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the snaplock fingerprint operation create params
func (o *SnaplockFingerprintOperationCreateParams) WithInfo(info *models.SnaplockFileFingerprint) *SnaplockFingerprintOperationCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the snaplock fingerprint operation create params
func (o *SnaplockFingerprintOperationCreateParams) SetInfo(info *models.SnaplockFileFingerprint) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the snaplock fingerprint operation create params
func (o *SnaplockFingerprintOperationCreateParams) WithReturnRecords(returnRecords *bool) *SnaplockFingerprintOperationCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the snaplock fingerprint operation create params
func (o *SnaplockFingerprintOperationCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockFingerprintOperationCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
