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
)

// NewSvmSSHServerCollectionGetParams creates a new SvmSSHServerCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmSSHServerCollectionGetParams() *SvmSSHServerCollectionGetParams {
	return &SvmSSHServerCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmSSHServerCollectionGetParamsWithTimeout creates a new SvmSSHServerCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSvmSSHServerCollectionGetParamsWithTimeout(timeout time.Duration) *SvmSSHServerCollectionGetParams {
	return &SvmSSHServerCollectionGetParams{
		timeout: timeout,
	}
}

// NewSvmSSHServerCollectionGetParamsWithContext creates a new SvmSSHServerCollectionGetParams object
// with the ability to set a context for a request.
func NewSvmSSHServerCollectionGetParamsWithContext(ctx context.Context) *SvmSSHServerCollectionGetParams {
	return &SvmSSHServerCollectionGetParams{
		Context: ctx,
	}
}

// NewSvmSSHServerCollectionGetParamsWithHTTPClient creates a new SvmSSHServerCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmSSHServerCollectionGetParamsWithHTTPClient(client *http.Client) *SvmSSHServerCollectionGetParams {
	return &SvmSSHServerCollectionGetParams{
		HTTPClient: client,
	}
}

/*
SvmSSHServerCollectionGetParams contains all the parameters to send to the API endpoint

	for the svm ssh server collection get operation.

	Typically these are written to a http.Request.
*/
type SvmSSHServerCollectionGetParams struct {

	/* Ciphers.

	   Filter by ciphers
	*/
	Ciphers *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* KeyExchangeAlgorithms.

	   Filter by key_exchange_algorithms
	*/
	KeyExchangeAlgorithms *string

	/* MacAlgorithms.

	   Filter by mac_algorithms
	*/
	MacAlgorithms *string

	/* MaxAuthenticationRetryCount.

	   Filter by max_authentication_retry_count
	*/
	MaxAuthenticationRetryCount *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

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

	   Filter by svm.uuid
	*/
	SvmUUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the svm ssh server collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmSSHServerCollectionGetParams) WithDefaults() *SvmSSHServerCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm ssh server collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmSSHServerCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SvmSSHServerCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithTimeout(timeout time.Duration) *SvmSSHServerCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithContext(ctx context.Context) *SvmSSHServerCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithHTTPClient(client *http.Client) *SvmSSHServerCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCiphers adds the ciphers to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithCiphers(ciphers *string) *SvmSSHServerCollectionGetParams {
	o.SetCiphers(ciphers)
	return o
}

// SetCiphers adds the ciphers to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetCiphers(ciphers *string) {
	o.Ciphers = ciphers
}

// WithFields adds the fields to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithFields(fields []string) *SvmSSHServerCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithKeyExchangeAlgorithms adds the keyExchangeAlgorithms to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithKeyExchangeAlgorithms(keyExchangeAlgorithms *string) *SvmSSHServerCollectionGetParams {
	o.SetKeyExchangeAlgorithms(keyExchangeAlgorithms)
	return o
}

// SetKeyExchangeAlgorithms adds the keyExchangeAlgorithms to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetKeyExchangeAlgorithms(keyExchangeAlgorithms *string) {
	o.KeyExchangeAlgorithms = keyExchangeAlgorithms
}

// WithMacAlgorithms adds the macAlgorithms to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithMacAlgorithms(macAlgorithms *string) *SvmSSHServerCollectionGetParams {
	o.SetMacAlgorithms(macAlgorithms)
	return o
}

// SetMacAlgorithms adds the macAlgorithms to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetMacAlgorithms(macAlgorithms *string) {
	o.MacAlgorithms = macAlgorithms
}

// WithMaxAuthenticationRetryCount adds the maxAuthenticationRetryCount to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithMaxAuthenticationRetryCount(maxAuthenticationRetryCount *int64) *SvmSSHServerCollectionGetParams {
	o.SetMaxAuthenticationRetryCount(maxAuthenticationRetryCount)
	return o
}

// SetMaxAuthenticationRetryCount adds the maxAuthenticationRetryCount to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetMaxAuthenticationRetryCount(maxAuthenticationRetryCount *int64) {
	o.MaxAuthenticationRetryCount = maxAuthenticationRetryCount
}

// WithMaxRecords adds the maxRecords to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithMaxRecords(maxRecords *int64) *SvmSSHServerCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithOrderBy(orderBy []string) *SvmSSHServerCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithReturnRecords(returnRecords *bool) *SvmSSHServerCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *SvmSSHServerCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmName adds the svmName to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithSvmName(svmName *string) *SvmSSHServerCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithSvmUUID(svmUUID *string) *SvmSSHServerCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SvmSSHServerCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ciphers != nil {

		// query param ciphers
		var qrCiphers string

		if o.Ciphers != nil {
			qrCiphers = *o.Ciphers
		}
		qCiphers := qrCiphers
		if qCiphers != "" {

			if err := r.SetQueryParam("ciphers", qCiphers); err != nil {
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

	if o.KeyExchangeAlgorithms != nil {

		// query param key_exchange_algorithms
		var qrKeyExchangeAlgorithms string

		if o.KeyExchangeAlgorithms != nil {
			qrKeyExchangeAlgorithms = *o.KeyExchangeAlgorithms
		}
		qKeyExchangeAlgorithms := qrKeyExchangeAlgorithms
		if qKeyExchangeAlgorithms != "" {

			if err := r.SetQueryParam("key_exchange_algorithms", qKeyExchangeAlgorithms); err != nil {
				return err
			}
		}
	}

	if o.MacAlgorithms != nil {

		// query param mac_algorithms
		var qrMacAlgorithms string

		if o.MacAlgorithms != nil {
			qrMacAlgorithms = *o.MacAlgorithms
		}
		qMacAlgorithms := qrMacAlgorithms
		if qMacAlgorithms != "" {

			if err := r.SetQueryParam("mac_algorithms", qMacAlgorithms); err != nil {
				return err
			}
		}
	}

	if o.MaxAuthenticationRetryCount != nil {

		// query param max_authentication_retry_count
		var qrMaxAuthenticationRetryCount int64

		if o.MaxAuthenticationRetryCount != nil {
			qrMaxAuthenticationRetryCount = *o.MaxAuthenticationRetryCount
		}
		qMaxAuthenticationRetryCount := swag.FormatInt64(qrMaxAuthenticationRetryCount)
		if qMaxAuthenticationRetryCount != "" {

			if err := r.SetQueryParam("max_authentication_retry_count", qMaxAuthenticationRetryCount); err != nil {
				return err
			}
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

	if o.SvmUUID != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SvmUUID != nil {
			qrSvmUUID = *o.SvmUUID
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSvmSSHServerCollectionGet binds the parameter fields
func (o *SvmSSHServerCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSvmSSHServerCollectionGet binds the parameter order_by
func (o *SvmSSHServerCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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