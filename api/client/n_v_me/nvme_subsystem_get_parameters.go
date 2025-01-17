// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

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

// NewNvmeSubsystemGetParams creates a new NvmeSubsystemGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNvmeSubsystemGetParams() *NvmeSubsystemGetParams {
	return &NvmeSubsystemGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNvmeSubsystemGetParamsWithTimeout creates a new NvmeSubsystemGetParams object
// with the ability to set a timeout on a request.
func NewNvmeSubsystemGetParamsWithTimeout(timeout time.Duration) *NvmeSubsystemGetParams {
	return &NvmeSubsystemGetParams{
		timeout: timeout,
	}
}

// NewNvmeSubsystemGetParamsWithContext creates a new NvmeSubsystemGetParams object
// with the ability to set a context for a request.
func NewNvmeSubsystemGetParamsWithContext(ctx context.Context) *NvmeSubsystemGetParams {
	return &NvmeSubsystemGetParams{
		Context: ctx,
	}
}

// NewNvmeSubsystemGetParamsWithHTTPClient creates a new NvmeSubsystemGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNvmeSubsystemGetParamsWithHTTPClient(client *http.Client) *NvmeSubsystemGetParams {
	return &NvmeSubsystemGetParams{
		HTTPClient: client,
	}
}

/*
NvmeSubsystemGetParams contains all the parameters to send to the API endpoint

	for the nvme subsystem get operation.

	Typically these are written to a http.Request.
*/
type NvmeSubsystemGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* UUID.

	   The unique identifier of the NVMe subsystem.

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nvme subsystem get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemGetParams) WithDefaults() *NvmeSubsystemGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nvme subsystem get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the nvme subsystem get params
func (o *NvmeSubsystemGetParams) WithTimeout(timeout time.Duration) *NvmeSubsystemGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nvme subsystem get params
func (o *NvmeSubsystemGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nvme subsystem get params
func (o *NvmeSubsystemGetParams) WithContext(ctx context.Context) *NvmeSubsystemGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nvme subsystem get params
func (o *NvmeSubsystemGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nvme subsystem get params
func (o *NvmeSubsystemGetParams) WithHTTPClient(client *http.Client) *NvmeSubsystemGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nvme subsystem get params
func (o *NvmeSubsystemGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the nvme subsystem get params
func (o *NvmeSubsystemGetParams) WithFields(fields []string) *NvmeSubsystemGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the nvme subsystem get params
func (o *NvmeSubsystemGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithUUID adds the uuid to the nvme subsystem get params
func (o *NvmeSubsystemGetParams) WithUUID(uuid string) *NvmeSubsystemGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the nvme subsystem get params
func (o *NvmeSubsystemGetParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NvmeSubsystemGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
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

// bindParamNvmeSubsystemGet binds the parameter fields
func (o *NvmeSubsystemGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}