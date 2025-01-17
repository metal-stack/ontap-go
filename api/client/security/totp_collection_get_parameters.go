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

// NewTotpCollectionGetParams creates a new TotpCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTotpCollectionGetParams() *TotpCollectionGetParams {
	return &TotpCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTotpCollectionGetParamsWithTimeout creates a new TotpCollectionGetParams object
// with the ability to set a timeout on a request.
func NewTotpCollectionGetParamsWithTimeout(timeout time.Duration) *TotpCollectionGetParams {
	return &TotpCollectionGetParams{
		timeout: timeout,
	}
}

// NewTotpCollectionGetParamsWithContext creates a new TotpCollectionGetParams object
// with the ability to set a context for a request.
func NewTotpCollectionGetParamsWithContext(ctx context.Context) *TotpCollectionGetParams {
	return &TotpCollectionGetParams{
		Context: ctx,
	}
}

// NewTotpCollectionGetParamsWithHTTPClient creates a new TotpCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewTotpCollectionGetParamsWithHTTPClient(client *http.Client) *TotpCollectionGetParams {
	return &TotpCollectionGetParams{
		HTTPClient: client,
	}
}

/*
TotpCollectionGetParams contains all the parameters to send to the API endpoint

	for the totp collection get operation.

	Typically these are written to a http.Request.
*/
type TotpCollectionGetParams struct {

	/* AccountName.

	   Filter by account.name
	*/
	AccountName *string

	/* Comment.

	   Filter by comment
	*/
	Comment *string

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

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* OwnerName.

	   Filter by owner.name
	*/
	OwnerName *string

	/* OwnerUUID.

	   Filter by owner.uuid
	*/
	OwnerUUID *string

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

	/* Scope.

	   Filter by scope
	*/
	Scope *string

	/* ShaFingerprint.

	   Filter by sha_fingerprint
	*/
	ShaFingerprint *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the totp collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TotpCollectionGetParams) WithDefaults() *TotpCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the totp collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TotpCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := TotpCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the totp collection get params
func (o *TotpCollectionGetParams) WithTimeout(timeout time.Duration) *TotpCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the totp collection get params
func (o *TotpCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the totp collection get params
func (o *TotpCollectionGetParams) WithContext(ctx context.Context) *TotpCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the totp collection get params
func (o *TotpCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the totp collection get params
func (o *TotpCollectionGetParams) WithHTTPClient(client *http.Client) *TotpCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the totp collection get params
func (o *TotpCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountName adds the accountName to the totp collection get params
func (o *TotpCollectionGetParams) WithAccountName(accountName *string) *TotpCollectionGetParams {
	o.SetAccountName(accountName)
	return o
}

// SetAccountName adds the accountName to the totp collection get params
func (o *TotpCollectionGetParams) SetAccountName(accountName *string) {
	o.AccountName = accountName
}

// WithComment adds the comment to the totp collection get params
func (o *TotpCollectionGetParams) WithComment(comment *string) *TotpCollectionGetParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the totp collection get params
func (o *TotpCollectionGetParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithEnabled adds the enabled to the totp collection get params
func (o *TotpCollectionGetParams) WithEnabled(enabled *bool) *TotpCollectionGetParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the totp collection get params
func (o *TotpCollectionGetParams) SetEnabled(enabled *bool) {
	o.Enabled = enabled
}

// WithFields adds the fields to the totp collection get params
func (o *TotpCollectionGetParams) WithFields(fields []string) *TotpCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the totp collection get params
func (o *TotpCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the totp collection get params
func (o *TotpCollectionGetParams) WithMaxRecords(maxRecords *int64) *TotpCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the totp collection get params
func (o *TotpCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the totp collection get params
func (o *TotpCollectionGetParams) WithOrderBy(orderBy []string) *TotpCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the totp collection get params
func (o *TotpCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithOwnerName adds the ownerName to the totp collection get params
func (o *TotpCollectionGetParams) WithOwnerName(ownerName *string) *TotpCollectionGetParams {
	o.SetOwnerName(ownerName)
	return o
}

// SetOwnerName adds the ownerName to the totp collection get params
func (o *TotpCollectionGetParams) SetOwnerName(ownerName *string) {
	o.OwnerName = ownerName
}

// WithOwnerUUID adds the ownerUUID to the totp collection get params
func (o *TotpCollectionGetParams) WithOwnerUUID(ownerUUID *string) *TotpCollectionGetParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the totp collection get params
func (o *TotpCollectionGetParams) SetOwnerUUID(ownerUUID *string) {
	o.OwnerUUID = ownerUUID
}

// WithReturnRecords adds the returnRecords to the totp collection get params
func (o *TotpCollectionGetParams) WithReturnRecords(returnRecords *bool) *TotpCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the totp collection get params
func (o *TotpCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the totp collection get params
func (o *TotpCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *TotpCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the totp collection get params
func (o *TotpCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScope adds the scope to the totp collection get params
func (o *TotpCollectionGetParams) WithScope(scope *string) *TotpCollectionGetParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the totp collection get params
func (o *TotpCollectionGetParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithShaFingerprint adds the shaFingerprint to the totp collection get params
func (o *TotpCollectionGetParams) WithShaFingerprint(shaFingerprint *string) *TotpCollectionGetParams {
	o.SetShaFingerprint(shaFingerprint)
	return o
}

// SetShaFingerprint adds the shaFingerprint to the totp collection get params
func (o *TotpCollectionGetParams) SetShaFingerprint(shaFingerprint *string) {
	o.ShaFingerprint = shaFingerprint
}

// WriteToRequest writes these params to a swagger request
func (o *TotpCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountName != nil {

		// query param account.name
		var qrAccountName string

		if o.AccountName != nil {
			qrAccountName = *o.AccountName
		}
		qAccountName := qrAccountName
		if qAccountName != "" {

			if err := r.SetQueryParam("account.name", qAccountName); err != nil {
				return err
			}
		}
	}

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

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

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.OwnerName != nil {

		// query param owner.name
		var qrOwnerName string

		if o.OwnerName != nil {
			qrOwnerName = *o.OwnerName
		}
		qOwnerName := qrOwnerName
		if qOwnerName != "" {

			if err := r.SetQueryParam("owner.name", qOwnerName); err != nil {
				return err
			}
		}
	}

	if o.OwnerUUID != nil {

		// query param owner.uuid
		var qrOwnerUUID string

		if o.OwnerUUID != nil {
			qrOwnerUUID = *o.OwnerUUID
		}
		qOwnerUUID := qrOwnerUUID
		if qOwnerUUID != "" {

			if err := r.SetQueryParam("owner.uuid", qOwnerUUID); err != nil {
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

	if o.Scope != nil {

		// query param scope
		var qrScope string

		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}
	}

	if o.ShaFingerprint != nil {

		// query param sha_fingerprint
		var qrShaFingerprint string

		if o.ShaFingerprint != nil {
			qrShaFingerprint = *o.ShaFingerprint
		}
		qShaFingerprint := qrShaFingerprint
		if qShaFingerprint != "" {

			if err := r.SetQueryParam("sha_fingerprint", qShaFingerprint); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamTotpCollectionGet binds the parameter fields
func (o *TotpCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamTotpCollectionGet binds the parameter order_by
func (o *TotpCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
