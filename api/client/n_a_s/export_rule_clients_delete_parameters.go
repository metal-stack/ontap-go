// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewExportRuleClientsDeleteParams creates a new ExportRuleClientsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportRuleClientsDeleteParams() *ExportRuleClientsDeleteParams {
	return &ExportRuleClientsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportRuleClientsDeleteParamsWithTimeout creates a new ExportRuleClientsDeleteParams object
// with the ability to set a timeout on a request.
func NewExportRuleClientsDeleteParamsWithTimeout(timeout time.Duration) *ExportRuleClientsDeleteParams {
	return &ExportRuleClientsDeleteParams{
		timeout: timeout,
	}
}

// NewExportRuleClientsDeleteParamsWithContext creates a new ExportRuleClientsDeleteParams object
// with the ability to set a context for a request.
func NewExportRuleClientsDeleteParamsWithContext(ctx context.Context) *ExportRuleClientsDeleteParams {
	return &ExportRuleClientsDeleteParams{
		Context: ctx,
	}
}

// NewExportRuleClientsDeleteParamsWithHTTPClient creates a new ExportRuleClientsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportRuleClientsDeleteParamsWithHTTPClient(client *http.Client) *ExportRuleClientsDeleteParams {
	return &ExportRuleClientsDeleteParams{
		HTTPClient: client,
	}
}

/*
ExportRuleClientsDeleteParams contains all the parameters to send to the API endpoint

	for the export rule clients delete operation.

	Typically these are written to a http.Request.
*/
type ExportRuleClientsDeleteParams struct {

	/* Index.

	   Export Rule Index
	*/
	Index int64

	/* Match.

	   Export Client Match
	*/
	Match string

	/* PolicyID.

	   Export Policy ID
	*/
	PolicyID int64

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the export rule clients delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleClientsDeleteParams) WithDefaults() *ExportRuleClientsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export rule clients delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleClientsDeleteParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := ExportRuleClientsDeleteParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the export rule clients delete params
func (o *ExportRuleClientsDeleteParams) WithTimeout(timeout time.Duration) *ExportRuleClientsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export rule clients delete params
func (o *ExportRuleClientsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export rule clients delete params
func (o *ExportRuleClientsDeleteParams) WithContext(ctx context.Context) *ExportRuleClientsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export rule clients delete params
func (o *ExportRuleClientsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export rule clients delete params
func (o *ExportRuleClientsDeleteParams) WithHTTPClient(client *http.Client) *ExportRuleClientsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export rule clients delete params
func (o *ExportRuleClientsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the export rule clients delete params
func (o *ExportRuleClientsDeleteParams) WithIndex(index int64) *ExportRuleClientsDeleteParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the export rule clients delete params
func (o *ExportRuleClientsDeleteParams) SetIndex(index int64) {
	o.Index = index
}

// WithMatch adds the match to the export rule clients delete params
func (o *ExportRuleClientsDeleteParams) WithMatch(match string) *ExportRuleClientsDeleteParams {
	o.SetMatch(match)
	return o
}

// SetMatch adds the match to the export rule clients delete params
func (o *ExportRuleClientsDeleteParams) SetMatch(match string) {
	o.Match = match
}

// WithPolicyID adds the policyID to the export rule clients delete params
func (o *ExportRuleClientsDeleteParams) WithPolicyID(policyID int64) *ExportRuleClientsDeleteParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the export rule clients delete params
func (o *ExportRuleClientsDeleteParams) SetPolicyID(policyID int64) {
	o.PolicyID = policyID
}

// WithReturnRecords adds the returnRecords to the export rule clients delete params
func (o *ExportRuleClientsDeleteParams) WithReturnRecords(returnRecords *bool) *ExportRuleClientsDeleteParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the export rule clients delete params
func (o *ExportRuleClientsDeleteParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *ExportRuleClientsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

	// path param match
	if err := r.SetPathParam("match", o.Match); err != nil {
		return err
	}

	// path param policy.id
	if err := r.SetPathParam("policy.id", swag.FormatInt64(o.PolicyID)); err != nil {
		return err
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
