// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NewGlobalCacheSettingGetParams creates a new GlobalCacheSettingGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGlobalCacheSettingGetParams() *GlobalCacheSettingGetParams {
	return &GlobalCacheSettingGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGlobalCacheSettingGetParamsWithTimeout creates a new GlobalCacheSettingGetParams object
// with the ability to set a timeout on a request.
func NewGlobalCacheSettingGetParamsWithTimeout(timeout time.Duration) *GlobalCacheSettingGetParams {
	return &GlobalCacheSettingGetParams{
		timeout: timeout,
	}
}

// NewGlobalCacheSettingGetParamsWithContext creates a new GlobalCacheSettingGetParams object
// with the ability to set a context for a request.
func NewGlobalCacheSettingGetParamsWithContext(ctx context.Context) *GlobalCacheSettingGetParams {
	return &GlobalCacheSettingGetParams{
		Context: ctx,
	}
}

// NewGlobalCacheSettingGetParamsWithHTTPClient creates a new GlobalCacheSettingGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGlobalCacheSettingGetParamsWithHTTPClient(client *http.Client) *GlobalCacheSettingGetParams {
	return &GlobalCacheSettingGetParams{
		HTTPClient: client,
	}
}

/*
GlobalCacheSettingGetParams contains all the parameters to send to the API endpoint

	for the global cache setting get operation.

	Typically these are written to a http.Request.
*/
type GlobalCacheSettingGetParams struct {

	/* EvictionTimeInterval.

	   Filter by eviction_time_interval
	*/
	EvictionTimeInterval *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* RemoteFetchEnabled.

	   Filter by remote_fetch_enabled
	*/
	RemoteFetchEnabled *bool

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the global cache setting get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GlobalCacheSettingGetParams) WithDefaults() *GlobalCacheSettingGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the global cache setting get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GlobalCacheSettingGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := GlobalCacheSettingGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the global cache setting get params
func (o *GlobalCacheSettingGetParams) WithTimeout(timeout time.Duration) *GlobalCacheSettingGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the global cache setting get params
func (o *GlobalCacheSettingGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the global cache setting get params
func (o *GlobalCacheSettingGetParams) WithContext(ctx context.Context) *GlobalCacheSettingGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the global cache setting get params
func (o *GlobalCacheSettingGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the global cache setting get params
func (o *GlobalCacheSettingGetParams) WithHTTPClient(client *http.Client) *GlobalCacheSettingGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the global cache setting get params
func (o *GlobalCacheSettingGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEvictionTimeInterval adds the evictionTimeInterval to the global cache setting get params
func (o *GlobalCacheSettingGetParams) WithEvictionTimeInterval(evictionTimeInterval *string) *GlobalCacheSettingGetParams {
	o.SetEvictionTimeInterval(evictionTimeInterval)
	return o
}

// SetEvictionTimeInterval adds the evictionTimeInterval to the global cache setting get params
func (o *GlobalCacheSettingGetParams) SetEvictionTimeInterval(evictionTimeInterval *string) {
	o.EvictionTimeInterval = evictionTimeInterval
}

// WithFields adds the fields to the global cache setting get params
func (o *GlobalCacheSettingGetParams) WithFields(fields []string) *GlobalCacheSettingGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the global cache setting get params
func (o *GlobalCacheSettingGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the global cache setting get params
func (o *GlobalCacheSettingGetParams) WithMaxRecords(maxRecords *int64) *GlobalCacheSettingGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the global cache setting get params
func (o *GlobalCacheSettingGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the global cache setting get params
func (o *GlobalCacheSettingGetParams) WithOrderBy(orderBy []string) *GlobalCacheSettingGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the global cache setting get params
func (o *GlobalCacheSettingGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithRemoteFetchEnabled adds the remoteFetchEnabled to the global cache setting get params
func (o *GlobalCacheSettingGetParams) WithRemoteFetchEnabled(remoteFetchEnabled *bool) *GlobalCacheSettingGetParams {
	o.SetRemoteFetchEnabled(remoteFetchEnabled)
	return o
}

// SetRemoteFetchEnabled adds the remoteFetchEnabled to the global cache setting get params
func (o *GlobalCacheSettingGetParams) SetRemoteFetchEnabled(remoteFetchEnabled *bool) {
	o.RemoteFetchEnabled = remoteFetchEnabled
}

// WithReturnRecords adds the returnRecords to the global cache setting get params
func (o *GlobalCacheSettingGetParams) WithReturnRecords(returnRecords *bool) *GlobalCacheSettingGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the global cache setting get params
func (o *GlobalCacheSettingGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the global cache setting get params
func (o *GlobalCacheSettingGetParams) WithReturnTimeout(returnTimeout *int64) *GlobalCacheSettingGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the global cache setting get params
func (o *GlobalCacheSettingGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *GlobalCacheSettingGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EvictionTimeInterval != nil {

		// query param eviction_time_interval
		var qrEvictionTimeInterval string

		if o.EvictionTimeInterval != nil {
			qrEvictionTimeInterval = *o.EvictionTimeInterval
		}
		qEvictionTimeInterval := qrEvictionTimeInterval
		if qEvictionTimeInterval != "" {

			if err := r.SetQueryParam("eviction_time_interval", qEvictionTimeInterval); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.RemoteFetchEnabled != nil {

		// query param remote_fetch_enabled
		var qrRemoteFetchEnabled bool

		if o.RemoteFetchEnabled != nil {
			qrRemoteFetchEnabled = *o.RemoteFetchEnabled
		}
		qRemoteFetchEnabled := swag.FormatBool(qrRemoteFetchEnabled)
		if qRemoteFetchEnabled != "" {

			if err := r.SetQueryParam("remote_fetch_enabled", qRemoteFetchEnabled); err != nil {
				return err
			}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGlobalCacheSettingGet binds the parameter fields
func (o *GlobalCacheSettingGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamGlobalCacheSettingGet binds the parameter order_by
func (o *GlobalCacheSettingGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}