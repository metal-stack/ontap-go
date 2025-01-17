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

// NewVscanOnAccessPolicyCollectionGetParams creates a new VscanOnAccessPolicyCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVscanOnAccessPolicyCollectionGetParams() *VscanOnAccessPolicyCollectionGetParams {
	return &VscanOnAccessPolicyCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVscanOnAccessPolicyCollectionGetParamsWithTimeout creates a new VscanOnAccessPolicyCollectionGetParams object
// with the ability to set a timeout on a request.
func NewVscanOnAccessPolicyCollectionGetParamsWithTimeout(timeout time.Duration) *VscanOnAccessPolicyCollectionGetParams {
	return &VscanOnAccessPolicyCollectionGetParams{
		timeout: timeout,
	}
}

// NewVscanOnAccessPolicyCollectionGetParamsWithContext creates a new VscanOnAccessPolicyCollectionGetParams object
// with the ability to set a context for a request.
func NewVscanOnAccessPolicyCollectionGetParamsWithContext(ctx context.Context) *VscanOnAccessPolicyCollectionGetParams {
	return &VscanOnAccessPolicyCollectionGetParams{
		Context: ctx,
	}
}

// NewVscanOnAccessPolicyCollectionGetParamsWithHTTPClient creates a new VscanOnAccessPolicyCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewVscanOnAccessPolicyCollectionGetParamsWithHTTPClient(client *http.Client) *VscanOnAccessPolicyCollectionGetParams {
	return &VscanOnAccessPolicyCollectionGetParams{
		HTTPClient: client,
	}
}

/*
VscanOnAccessPolicyCollectionGetParams contains all the parameters to send to the API endpoint

	for the vscan on access policy collection get operation.

	Typically these are written to a http.Request.
*/
type VscanOnAccessPolicyCollectionGetParams struct {

	/* Enabled.

	   Filter by enabled
	*/
	Enabled *bool

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Mandatory.

	   Filter by mandatory
	*/
	Mandatory *bool

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Name.

	   Filter by name
	*/
	Name *string

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

	/* ScopeExcludeExtensions.

	   Filter by scope.exclude_extensions
	*/
	ScopeExcludeExtensions *string

	/* ScopeExcludePaths.

	   Filter by scope.exclude_paths
	*/
	ScopeExcludePaths *string

	/* ScopeIncludeExtensions.

	   Filter by scope.include_extensions
	*/
	ScopeIncludeExtensions *string

	/* ScopeMaxFileSize.

	   Filter by scope.max_file_size
	*/
	ScopeMaxFileSize *int64

	/* ScopeOnlyExecuteAccess.

	   Filter by scope.only_execute_access
	*/
	ScopeOnlyExecuteAccess *bool

	/* ScopeScanReadonlyVolumes.

	   Filter by scope.scan_readonly_volumes
	*/
	ScopeScanReadonlyVolumes *bool

	/* ScopeScanWithoutExtension.

	   Filter by scope.scan_without_extension
	*/
	ScopeScanWithoutExtension *bool

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vscan on access policy collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanOnAccessPolicyCollectionGetParams) WithDefaults() *VscanOnAccessPolicyCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vscan on access policy collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanOnAccessPolicyCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := VscanOnAccessPolicyCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithTimeout(timeout time.Duration) *VscanOnAccessPolicyCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithContext(ctx context.Context) *VscanOnAccessPolicyCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithHTTPClient(client *http.Client) *VscanOnAccessPolicyCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabled adds the enabled to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithEnabled(enabled *bool) *VscanOnAccessPolicyCollectionGetParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetEnabled(enabled *bool) {
	o.Enabled = enabled
}

// WithFields adds the fields to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithFields(fields []string) *VscanOnAccessPolicyCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMandatory adds the mandatory to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithMandatory(mandatory *bool) *VscanOnAccessPolicyCollectionGetParams {
	o.SetMandatory(mandatory)
	return o
}

// SetMandatory adds the mandatory to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetMandatory(mandatory *bool) {
	o.Mandatory = mandatory
}

// WithMaxRecords adds the maxRecords to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithMaxRecords(maxRecords *int64) *VscanOnAccessPolicyCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithName adds the name to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithName(name *string) *VscanOnAccessPolicyCollectionGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetName(name *string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithOrderBy(orderBy []string) *VscanOnAccessPolicyCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithReturnRecords(returnRecords *bool) *VscanOnAccessPolicyCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *VscanOnAccessPolicyCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScopeExcludeExtensions adds the scopeExcludeExtensions to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithScopeExcludeExtensions(scopeExcludeExtensions *string) *VscanOnAccessPolicyCollectionGetParams {
	o.SetScopeExcludeExtensions(scopeExcludeExtensions)
	return o
}

// SetScopeExcludeExtensions adds the scopeExcludeExtensions to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetScopeExcludeExtensions(scopeExcludeExtensions *string) {
	o.ScopeExcludeExtensions = scopeExcludeExtensions
}

// WithScopeExcludePaths adds the scopeExcludePaths to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithScopeExcludePaths(scopeExcludePaths *string) *VscanOnAccessPolicyCollectionGetParams {
	o.SetScopeExcludePaths(scopeExcludePaths)
	return o
}

// SetScopeExcludePaths adds the scopeExcludePaths to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetScopeExcludePaths(scopeExcludePaths *string) {
	o.ScopeExcludePaths = scopeExcludePaths
}

// WithScopeIncludeExtensions adds the scopeIncludeExtensions to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithScopeIncludeExtensions(scopeIncludeExtensions *string) *VscanOnAccessPolicyCollectionGetParams {
	o.SetScopeIncludeExtensions(scopeIncludeExtensions)
	return o
}

// SetScopeIncludeExtensions adds the scopeIncludeExtensions to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetScopeIncludeExtensions(scopeIncludeExtensions *string) {
	o.ScopeIncludeExtensions = scopeIncludeExtensions
}

// WithScopeMaxFileSize adds the scopeMaxFileSize to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithScopeMaxFileSize(scopeMaxFileSize *int64) *VscanOnAccessPolicyCollectionGetParams {
	o.SetScopeMaxFileSize(scopeMaxFileSize)
	return o
}

// SetScopeMaxFileSize adds the scopeMaxFileSize to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetScopeMaxFileSize(scopeMaxFileSize *int64) {
	o.ScopeMaxFileSize = scopeMaxFileSize
}

// WithScopeOnlyExecuteAccess adds the scopeOnlyExecuteAccess to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithScopeOnlyExecuteAccess(scopeOnlyExecuteAccess *bool) *VscanOnAccessPolicyCollectionGetParams {
	o.SetScopeOnlyExecuteAccess(scopeOnlyExecuteAccess)
	return o
}

// SetScopeOnlyExecuteAccess adds the scopeOnlyExecuteAccess to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetScopeOnlyExecuteAccess(scopeOnlyExecuteAccess *bool) {
	o.ScopeOnlyExecuteAccess = scopeOnlyExecuteAccess
}

// WithScopeScanReadonlyVolumes adds the scopeScanReadonlyVolumes to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithScopeScanReadonlyVolumes(scopeScanReadonlyVolumes *bool) *VscanOnAccessPolicyCollectionGetParams {
	o.SetScopeScanReadonlyVolumes(scopeScanReadonlyVolumes)
	return o
}

// SetScopeScanReadonlyVolumes adds the scopeScanReadonlyVolumes to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetScopeScanReadonlyVolumes(scopeScanReadonlyVolumes *bool) {
	o.ScopeScanReadonlyVolumes = scopeScanReadonlyVolumes
}

// WithScopeScanWithoutExtension adds the scopeScanWithoutExtension to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithScopeScanWithoutExtension(scopeScanWithoutExtension *bool) *VscanOnAccessPolicyCollectionGetParams {
	o.SetScopeScanWithoutExtension(scopeScanWithoutExtension)
	return o
}

// SetScopeScanWithoutExtension adds the scopeScanWithoutExtension to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetScopeScanWithoutExtension(scopeScanWithoutExtension *bool) {
	o.ScopeScanWithoutExtension = scopeScanWithoutExtension
}

// WithSvmUUID adds the svmUUID to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) WithSvmUUID(svmUUID string) *VscanOnAccessPolicyCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the vscan on access policy collection get params
func (o *VscanOnAccessPolicyCollectionGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *VscanOnAccessPolicyCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Mandatory != nil {

		// query param mandatory
		var qrMandatory bool

		if o.Mandatory != nil {
			qrMandatory = *o.Mandatory
		}
		qMandatory := swag.FormatBool(qrMandatory)
		if qMandatory != "" {

			if err := r.SetQueryParam("mandatory", qMandatory); err != nil {
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

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
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

	if o.ScopeExcludeExtensions != nil {

		// query param scope.exclude_extensions
		var qrScopeExcludeExtensions string

		if o.ScopeExcludeExtensions != nil {
			qrScopeExcludeExtensions = *o.ScopeExcludeExtensions
		}
		qScopeExcludeExtensions := qrScopeExcludeExtensions
		if qScopeExcludeExtensions != "" {

			if err := r.SetQueryParam("scope.exclude_extensions", qScopeExcludeExtensions); err != nil {
				return err
			}
		}
	}

	if o.ScopeExcludePaths != nil {

		// query param scope.exclude_paths
		var qrScopeExcludePaths string

		if o.ScopeExcludePaths != nil {
			qrScopeExcludePaths = *o.ScopeExcludePaths
		}
		qScopeExcludePaths := qrScopeExcludePaths
		if qScopeExcludePaths != "" {

			if err := r.SetQueryParam("scope.exclude_paths", qScopeExcludePaths); err != nil {
				return err
			}
		}
	}

	if o.ScopeIncludeExtensions != nil {

		// query param scope.include_extensions
		var qrScopeIncludeExtensions string

		if o.ScopeIncludeExtensions != nil {
			qrScopeIncludeExtensions = *o.ScopeIncludeExtensions
		}
		qScopeIncludeExtensions := qrScopeIncludeExtensions
		if qScopeIncludeExtensions != "" {

			if err := r.SetQueryParam("scope.include_extensions", qScopeIncludeExtensions); err != nil {
				return err
			}
		}
	}

	if o.ScopeMaxFileSize != nil {

		// query param scope.max_file_size
		var qrScopeMaxFileSize int64

		if o.ScopeMaxFileSize != nil {
			qrScopeMaxFileSize = *o.ScopeMaxFileSize
		}
		qScopeMaxFileSize := swag.FormatInt64(qrScopeMaxFileSize)
		if qScopeMaxFileSize != "" {

			if err := r.SetQueryParam("scope.max_file_size", qScopeMaxFileSize); err != nil {
				return err
			}
		}
	}

	if o.ScopeOnlyExecuteAccess != nil {

		// query param scope.only_execute_access
		var qrScopeOnlyExecuteAccess bool

		if o.ScopeOnlyExecuteAccess != nil {
			qrScopeOnlyExecuteAccess = *o.ScopeOnlyExecuteAccess
		}
		qScopeOnlyExecuteAccess := swag.FormatBool(qrScopeOnlyExecuteAccess)
		if qScopeOnlyExecuteAccess != "" {

			if err := r.SetQueryParam("scope.only_execute_access", qScopeOnlyExecuteAccess); err != nil {
				return err
			}
		}
	}

	if o.ScopeScanReadonlyVolumes != nil {

		// query param scope.scan_readonly_volumes
		var qrScopeScanReadonlyVolumes bool

		if o.ScopeScanReadonlyVolumes != nil {
			qrScopeScanReadonlyVolumes = *o.ScopeScanReadonlyVolumes
		}
		qScopeScanReadonlyVolumes := swag.FormatBool(qrScopeScanReadonlyVolumes)
		if qScopeScanReadonlyVolumes != "" {

			if err := r.SetQueryParam("scope.scan_readonly_volumes", qScopeScanReadonlyVolumes); err != nil {
				return err
			}
		}
	}

	if o.ScopeScanWithoutExtension != nil {

		// query param scope.scan_without_extension
		var qrScopeScanWithoutExtension bool

		if o.ScopeScanWithoutExtension != nil {
			qrScopeScanWithoutExtension = *o.ScopeScanWithoutExtension
		}
		qScopeScanWithoutExtension := swag.FormatBool(qrScopeScanWithoutExtension)
		if qScopeScanWithoutExtension != "" {

			if err := r.SetQueryParam("scope.scan_without_extension", qScopeScanWithoutExtension); err != nil {
				return err
			}
		}
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamVscanOnAccessPolicyCollectionGet binds the parameter fields
func (o *VscanOnAccessPolicyCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamVscanOnAccessPolicyCollectionGet binds the parameter order_by
func (o *VscanOnAccessPolicyCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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