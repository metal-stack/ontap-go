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
	"github.com/go-openapi/swag"

	"github.com/metal-stack/ontap-go/api/models"
)

// NewSecurityCertificateSignParams creates a new SecurityCertificateSignParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityCertificateSignParams() *SecurityCertificateSignParams {
	return &SecurityCertificateSignParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityCertificateSignParamsWithTimeout creates a new SecurityCertificateSignParams object
// with the ability to set a timeout on a request.
func NewSecurityCertificateSignParamsWithTimeout(timeout time.Duration) *SecurityCertificateSignParams {
	return &SecurityCertificateSignParams{
		timeout: timeout,
	}
}

// NewSecurityCertificateSignParamsWithContext creates a new SecurityCertificateSignParams object
// with the ability to set a context for a request.
func NewSecurityCertificateSignParamsWithContext(ctx context.Context) *SecurityCertificateSignParams {
	return &SecurityCertificateSignParams{
		Context: ctx,
	}
}

// NewSecurityCertificateSignParamsWithHTTPClient creates a new SecurityCertificateSignParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityCertificateSignParamsWithHTTPClient(client *http.Client) *SecurityCertificateSignParams {
	return &SecurityCertificateSignParams{
		HTTPClient: client,
	}
}

/*
SecurityCertificateSignParams contains all the parameters to send to the API endpoint

	for the security certificate sign operation.

	Typically these are written to a http.Request.
*/
type SecurityCertificateSignParams struct {

	/* CaUUID.

	   UUID of the existing certificate authority certificate
	*/
	CaUUID string

	/* Info.

	   Certificate sign information specification
	*/
	Info *models.SecurityCertificateSign

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security certificate sign params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityCertificateSignParams) WithDefaults() *SecurityCertificateSignParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security certificate sign params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityCertificateSignParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := SecurityCertificateSignParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the security certificate sign params
func (o *SecurityCertificateSignParams) WithTimeout(timeout time.Duration) *SecurityCertificateSignParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security certificate sign params
func (o *SecurityCertificateSignParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security certificate sign params
func (o *SecurityCertificateSignParams) WithContext(ctx context.Context) *SecurityCertificateSignParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security certificate sign params
func (o *SecurityCertificateSignParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security certificate sign params
func (o *SecurityCertificateSignParams) WithHTTPClient(client *http.Client) *SecurityCertificateSignParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security certificate sign params
func (o *SecurityCertificateSignParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCaUUID adds the caUUID to the security certificate sign params
func (o *SecurityCertificateSignParams) WithCaUUID(caUUID string) *SecurityCertificateSignParams {
	o.SetCaUUID(caUUID)
	return o
}

// SetCaUUID adds the caUuid to the security certificate sign params
func (o *SecurityCertificateSignParams) SetCaUUID(caUUID string) {
	o.CaUUID = caUUID
}

// WithInfo adds the info to the security certificate sign params
func (o *SecurityCertificateSignParams) WithInfo(info *models.SecurityCertificateSign) *SecurityCertificateSignParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security certificate sign params
func (o *SecurityCertificateSignParams) SetInfo(info *models.SecurityCertificateSign) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the security certificate sign params
func (o *SecurityCertificateSignParams) WithReturnRecords(returnRecords *bool) *SecurityCertificateSignParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the security certificate sign params
func (o *SecurityCertificateSignParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityCertificateSignParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ca.uuid
	if err := r.SetPathParam("ca.uuid", o.CaUUID); err != nil {
		return err
	}
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
