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

// NewUnixGroupSettingsGetParams creates a new UnixGroupSettingsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnixGroupSettingsGetParams() *UnixGroupSettingsGetParams {
	return &UnixGroupSettingsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnixGroupSettingsGetParamsWithTimeout creates a new UnixGroupSettingsGetParams object
// with the ability to set a timeout on a request.
func NewUnixGroupSettingsGetParamsWithTimeout(timeout time.Duration) *UnixGroupSettingsGetParams {
	return &UnixGroupSettingsGetParams{
		timeout: timeout,
	}
}

// NewUnixGroupSettingsGetParamsWithContext creates a new UnixGroupSettingsGetParams object
// with the ability to set a context for a request.
func NewUnixGroupSettingsGetParamsWithContext(ctx context.Context) *UnixGroupSettingsGetParams {
	return &UnixGroupSettingsGetParams{
		Context: ctx,
	}
}

// NewUnixGroupSettingsGetParamsWithHTTPClient creates a new UnixGroupSettingsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnixGroupSettingsGetParamsWithHTTPClient(client *http.Client) *UnixGroupSettingsGetParams {
	return &UnixGroupSettingsGetParams{
		HTTPClient: client,
	}
}

/*
UnixGroupSettingsGetParams contains all the parameters to send to the API endpoint

	for the unix group settings get operation.

	Typically these are written to a http.Request.
*/
type UnixGroupSettingsGetParams struct {

	/* Enabled.

	   Filter by enabled
	*/
	Enabled *bool

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* NegativeCacheEnabled.

	   Filter by negative_cache_enabled
	*/
	NegativeCacheEnabled *bool

	/* NegativeTTL.

	   Filter by negative_ttl
	*/
	NegativeTTL *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* PropagationEnabled.

	   Filter by propagation_enabled
	*/
	PropagationEnabled *bool

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

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   SVM UUID.
	*/
	SvmUUID string

	/* TTL.

	   Filter by ttl
	*/
	TTL *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unix group settings get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupSettingsGetParams) WithDefaults() *UnixGroupSettingsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unix group settings get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupSettingsGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := UnixGroupSettingsGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithTimeout(timeout time.Duration) *UnixGroupSettingsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithContext(ctx context.Context) *UnixGroupSettingsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithHTTPClient(client *http.Client) *UnixGroupSettingsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabled adds the enabled to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithEnabled(enabled *bool) *UnixGroupSettingsGetParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetEnabled(enabled *bool) {
	o.Enabled = enabled
}

// WithFields adds the fields to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithFields(fields []string) *UnixGroupSettingsGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithMaxRecords(maxRecords *int64) *UnixGroupSettingsGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNegativeCacheEnabled adds the negativeCacheEnabled to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithNegativeCacheEnabled(negativeCacheEnabled *bool) *UnixGroupSettingsGetParams {
	o.SetNegativeCacheEnabled(negativeCacheEnabled)
	return o
}

// SetNegativeCacheEnabled adds the negativeCacheEnabled to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetNegativeCacheEnabled(negativeCacheEnabled *bool) {
	o.NegativeCacheEnabled = negativeCacheEnabled
}

// WithNegativeTTL adds the negativeTTL to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithNegativeTTL(negativeTTL *string) *UnixGroupSettingsGetParams {
	o.SetNegativeTTL(negativeTTL)
	return o
}

// SetNegativeTTL adds the negativeTtl to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetNegativeTTL(negativeTTL *string) {
	o.NegativeTTL = negativeTTL
}

// WithOrderBy adds the orderBy to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithOrderBy(orderBy []string) *UnixGroupSettingsGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPropagationEnabled adds the propagationEnabled to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithPropagationEnabled(propagationEnabled *bool) *UnixGroupSettingsGetParams {
	o.SetPropagationEnabled(propagationEnabled)
	return o
}

// SetPropagationEnabled adds the propagationEnabled to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetPropagationEnabled(propagationEnabled *bool) {
	o.PropagationEnabled = propagationEnabled
}

// WithReturnRecords adds the returnRecords to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithReturnRecords(returnRecords *bool) *UnixGroupSettingsGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithReturnTimeout(returnTimeout *int64) *UnixGroupSettingsGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmName adds the svmName to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithSvmName(svmName *string) *UnixGroupSettingsGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithSvmUUID(svmUUID string) *UnixGroupSettingsGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WithTTL adds the ttl to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithTTL(ttl *string) *UnixGroupSettingsGetParams {
	o.SetTTL(ttl)
	return o
}

// SetTTL adds the ttl to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetTTL(ttl *string) {
	o.TTL = ttl
}

// WriteToRequest writes these params to a swagger request
func (o *UnixGroupSettingsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled bool

		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
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

	if o.NegativeCacheEnabled != nil {

		// query param negative_cache_enabled
		var qrNegativeCacheEnabled bool

		if o.NegativeCacheEnabled != nil {
			qrNegativeCacheEnabled = *o.NegativeCacheEnabled
		}
		qNegativeCacheEnabled := swag.FormatBool(qrNegativeCacheEnabled)
		if qNegativeCacheEnabled != "" {

			if err := r.SetQueryParam("negative_cache_enabled", qNegativeCacheEnabled); err != nil {
				return err
			}
		}
	}

	if o.NegativeTTL != nil {

		// query param negative_ttl
		var qrNegativeTTL string

		if o.NegativeTTL != nil {
			qrNegativeTTL = *o.NegativeTTL
		}
		qNegativeTTL := qrNegativeTTL
		if qNegativeTTL != "" {

			if err := r.SetQueryParam("negative_ttl", qNegativeTTL); err != nil {
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

	if o.PropagationEnabled != nil {

		// query param propagation_enabled
		var qrPropagationEnabled bool

		if o.PropagationEnabled != nil {
			qrPropagationEnabled = *o.PropagationEnabled
		}
		qPropagationEnabled := swag.FormatBool(qrPropagationEnabled)
		if qPropagationEnabled != "" {

			if err := r.SetQueryParam("propagation_enabled", qPropagationEnabled); err != nil {
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

	if o.SvmName != nil {

		// query param svm.name
		var qrSvmName string

		if o.SvmName != nil {
			qrSvmName = *o.SvmName
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if o.TTL != nil {

		// query param ttl
		var qrTTL string

		if o.TTL != nil {
			qrTTL = *o.TTL
		}
		qTTL := qrTTL
		if qTTL != "" {

			if err := r.SetQueryParam("ttl", qTTL); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamUnixGroupSettingsGet binds the parameter fields
func (o *UnixGroupSettingsGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamUnixGroupSettingsGet binds the parameter order_by
func (o *UnixGroupSettingsGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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