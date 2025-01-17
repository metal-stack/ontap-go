// Code generated by go-swagger; DO NOT EDIT.

package s_vm

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

// NewSvmPeerPermissionCollectionGetParams creates a new SvmPeerPermissionCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmPeerPermissionCollectionGetParams() *SvmPeerPermissionCollectionGetParams {
	return &SvmPeerPermissionCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmPeerPermissionCollectionGetParamsWithTimeout creates a new SvmPeerPermissionCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSvmPeerPermissionCollectionGetParamsWithTimeout(timeout time.Duration) *SvmPeerPermissionCollectionGetParams {
	return &SvmPeerPermissionCollectionGetParams{
		timeout: timeout,
	}
}

// NewSvmPeerPermissionCollectionGetParamsWithContext creates a new SvmPeerPermissionCollectionGetParams object
// with the ability to set a context for a request.
func NewSvmPeerPermissionCollectionGetParamsWithContext(ctx context.Context) *SvmPeerPermissionCollectionGetParams {
	return &SvmPeerPermissionCollectionGetParams{
		Context: ctx,
	}
}

// NewSvmPeerPermissionCollectionGetParamsWithHTTPClient creates a new SvmPeerPermissionCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmPeerPermissionCollectionGetParamsWithHTTPClient(client *http.Client) *SvmPeerPermissionCollectionGetParams {
	return &SvmPeerPermissionCollectionGetParams{
		HTTPClient: client,
	}
}

/*
SvmPeerPermissionCollectionGetParams contains all the parameters to send to the API endpoint

	for the svm peer permission collection get operation.

	Typically these are written to a http.Request.
*/
type SvmPeerPermissionCollectionGetParams struct {

	/* Applications.

	   Filter by applications
	*/
	Applications *string

	/* ClusterPeerName.

	   Filter by cluster_peer.name
	*/
	ClusterPeerName *string

	/* ClusterPeerUUID.

	   Filter by cluster_peer.uuid
	*/
	ClusterPeerUUID *string

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

// WithDefaults hydrates default values in the svm peer permission collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmPeerPermissionCollectionGetParams) WithDefaults() *SvmPeerPermissionCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm peer permission collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmPeerPermissionCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SvmPeerPermissionCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) WithTimeout(timeout time.Duration) *SvmPeerPermissionCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) WithContext(ctx context.Context) *SvmPeerPermissionCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) WithHTTPClient(client *http.Client) *SvmPeerPermissionCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplications adds the applications to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) WithApplications(applications *string) *SvmPeerPermissionCollectionGetParams {
	o.SetApplications(applications)
	return o
}

// SetApplications adds the applications to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) SetApplications(applications *string) {
	o.Applications = applications
}

// WithClusterPeerName adds the clusterPeerName to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) WithClusterPeerName(clusterPeerName *string) *SvmPeerPermissionCollectionGetParams {
	o.SetClusterPeerName(clusterPeerName)
	return o
}

// SetClusterPeerName adds the clusterPeerName to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) SetClusterPeerName(clusterPeerName *string) {
	o.ClusterPeerName = clusterPeerName
}

// WithClusterPeerUUID adds the clusterPeerUUID to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) WithClusterPeerUUID(clusterPeerUUID *string) *SvmPeerPermissionCollectionGetParams {
	o.SetClusterPeerUUID(clusterPeerUUID)
	return o
}

// SetClusterPeerUUID adds the clusterPeerUuid to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) SetClusterPeerUUID(clusterPeerUUID *string) {
	o.ClusterPeerUUID = clusterPeerUUID
}

// WithFields adds the fields to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) WithFields(fields []string) *SvmPeerPermissionCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) WithMaxRecords(maxRecords *int64) *SvmPeerPermissionCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) WithOrderBy(orderBy []string) *SvmPeerPermissionCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) WithReturnRecords(returnRecords *bool) *SvmPeerPermissionCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *SvmPeerPermissionCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmName adds the svmName to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) WithSvmName(svmName *string) *SvmPeerPermissionCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) WithSvmUUID(svmUUID *string) *SvmPeerPermissionCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the svm peer permission collection get params
func (o *SvmPeerPermissionCollectionGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SvmPeerPermissionCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Applications != nil {

		// query param applications
		var qrApplications string

		if o.Applications != nil {
			qrApplications = *o.Applications
		}
		qApplications := qrApplications
		if qApplications != "" {

			if err := r.SetQueryParam("applications", qApplications); err != nil {
				return err
			}
		}
	}

	if o.ClusterPeerName != nil {

		// query param cluster_peer.name
		var qrClusterPeerName string

		if o.ClusterPeerName != nil {
			qrClusterPeerName = *o.ClusterPeerName
		}
		qClusterPeerName := qrClusterPeerName
		if qClusterPeerName != "" {

			if err := r.SetQueryParam("cluster_peer.name", qClusterPeerName); err != nil {
				return err
			}
		}
	}

	if o.ClusterPeerUUID != nil {

		// query param cluster_peer.uuid
		var qrClusterPeerUUID string

		if o.ClusterPeerUUID != nil {
			qrClusterPeerUUID = *o.ClusterPeerUUID
		}
		qClusterPeerUUID := qrClusterPeerUUID
		if qClusterPeerUUID != "" {

			if err := r.SetQueryParam("cluster_peer.uuid", qClusterPeerUUID); err != nil {
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

// bindParamSvmPeerPermissionCollectionGet binds the parameter fields
func (o *SvmPeerPermissionCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSvmPeerPermissionCollectionGet binds the parameter order_by
func (o *SvmPeerPermissionCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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