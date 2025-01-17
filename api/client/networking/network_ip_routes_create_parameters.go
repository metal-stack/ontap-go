// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// NewNetworkIPRoutesCreateParams creates a new NetworkIPRoutesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkIPRoutesCreateParams() *NetworkIPRoutesCreateParams {
	return &NetworkIPRoutesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkIPRoutesCreateParamsWithTimeout creates a new NetworkIPRoutesCreateParams object
// with the ability to set a timeout on a request.
func NewNetworkIPRoutesCreateParamsWithTimeout(timeout time.Duration) *NetworkIPRoutesCreateParams {
	return &NetworkIPRoutesCreateParams{
		timeout: timeout,
	}
}

// NewNetworkIPRoutesCreateParamsWithContext creates a new NetworkIPRoutesCreateParams object
// with the ability to set a context for a request.
func NewNetworkIPRoutesCreateParamsWithContext(ctx context.Context) *NetworkIPRoutesCreateParams {
	return &NetworkIPRoutesCreateParams{
		Context: ctx,
	}
}

// NewNetworkIPRoutesCreateParamsWithHTTPClient creates a new NetworkIPRoutesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkIPRoutesCreateParamsWithHTTPClient(client *http.Client) *NetworkIPRoutesCreateParams {
	return &NetworkIPRoutesCreateParams{
		HTTPClient: client,
	}
}

/*
NetworkIPRoutesCreateParams contains all the parameters to send to the API endpoint

	for the network ip routes create operation.

	Typically these are written to a http.Request.
*/
type NetworkIPRoutesCreateParams struct {

	/* Info.

	   Route parameters
	*/
	Info *models.NetworkRoute

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ip routes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPRoutesCreateParams) WithDefaults() *NetworkIPRoutesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ip routes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPRoutesCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := NetworkIPRoutesCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the network ip routes create params
func (o *NetworkIPRoutesCreateParams) WithTimeout(timeout time.Duration) *NetworkIPRoutesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ip routes create params
func (o *NetworkIPRoutesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ip routes create params
func (o *NetworkIPRoutesCreateParams) WithContext(ctx context.Context) *NetworkIPRoutesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ip routes create params
func (o *NetworkIPRoutesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ip routes create params
func (o *NetworkIPRoutesCreateParams) WithHTTPClient(client *http.Client) *NetworkIPRoutesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ip routes create params
func (o *NetworkIPRoutesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the network ip routes create params
func (o *NetworkIPRoutesCreateParams) WithInfo(info *models.NetworkRoute) *NetworkIPRoutesCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the network ip routes create params
func (o *NetworkIPRoutesCreateParams) SetInfo(info *models.NetworkRoute) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the network ip routes create params
func (o *NetworkIPRoutesCreateParams) WithReturnRecords(returnRecords *bool) *NetworkIPRoutesCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the network ip routes create params
func (o *NetworkIPRoutesCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkIPRoutesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
