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

// NewTopMetricsSvmUserCollectionGetParams creates a new TopMetricsSvmUserCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTopMetricsSvmUserCollectionGetParams() *TopMetricsSvmUserCollectionGetParams {
	return &TopMetricsSvmUserCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTopMetricsSvmUserCollectionGetParamsWithTimeout creates a new TopMetricsSvmUserCollectionGetParams object
// with the ability to set a timeout on a request.
func NewTopMetricsSvmUserCollectionGetParamsWithTimeout(timeout time.Duration) *TopMetricsSvmUserCollectionGetParams {
	return &TopMetricsSvmUserCollectionGetParams{
		timeout: timeout,
	}
}

// NewTopMetricsSvmUserCollectionGetParamsWithContext creates a new TopMetricsSvmUserCollectionGetParams object
// with the ability to set a context for a request.
func NewTopMetricsSvmUserCollectionGetParamsWithContext(ctx context.Context) *TopMetricsSvmUserCollectionGetParams {
	return &TopMetricsSvmUserCollectionGetParams{
		Context: ctx,
	}
}

// NewTopMetricsSvmUserCollectionGetParamsWithHTTPClient creates a new TopMetricsSvmUserCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewTopMetricsSvmUserCollectionGetParamsWithHTTPClient(client *http.Client) *TopMetricsSvmUserCollectionGetParams {
	return &TopMetricsSvmUserCollectionGetParams{
		HTTPClient: client,
	}
}

/*
TopMetricsSvmUserCollectionGetParams contains all the parameters to send to the API endpoint

	for the top metrics svm user collection get operation.

	Typically these are written to a http.Request.
*/
type TopMetricsSvmUserCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* IopsErrorLowerBound.

	   Filter by iops.error.lower_bound
	*/
	IopsErrorLowerBound *int64

	/* IopsErrorUpperBound.

	   Filter by iops.error.upper_bound
	*/
	IopsErrorUpperBound *int64

	/* IopsRead.

	   Filter by iops.read
	*/
	IopsRead *int64

	/* IopsWrite.

	   Filter by iops.write
	*/
	IopsWrite *int64

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

	   SVM UUID
	*/
	SvmUUID string

	/* ThroughputErrorLowerBound.

	   Filter by throughput.error.lower_bound
	*/
	ThroughputErrorLowerBound *int64

	/* ThroughputErrorUpperBound.

	   Filter by throughput.error.upper_bound
	*/
	ThroughputErrorUpperBound *int64

	/* ThroughputRead.

	   Filter by throughput.read
	*/
	ThroughputRead *int64

	/* ThroughputWrite.

	   Filter by throughput.write
	*/
	ThroughputWrite *int64

	/* TopMetric.

	   I/O activity type

	   Default: "iops.read"
	*/
	TopMetric *string

	/* UserID.

	   Filter by user_id
	*/
	UserID *string

	/* UserName.

	   Filter by user_name
	*/
	UserName *string

	/* VolumesName.

	   Filter by volumes.name
	*/
	VolumesName *string

	/* VolumesUUID.

	   Filter by volumes.uuid
	*/
	VolumesUUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the top metrics svm user collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TopMetricsSvmUserCollectionGetParams) WithDefaults() *TopMetricsSvmUserCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the top metrics svm user collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TopMetricsSvmUserCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		topMetricDefault = string("iops.read")
	)

	val := TopMetricsSvmUserCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
		TopMetric:     &topMetricDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithTimeout(timeout time.Duration) *TopMetricsSvmUserCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithContext(ctx context.Context) *TopMetricsSvmUserCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithHTTPClient(client *http.Client) *TopMetricsSvmUserCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithFields(fields []string) *TopMetricsSvmUserCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIopsErrorLowerBound adds the iopsErrorLowerBound to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithIopsErrorLowerBound(iopsErrorLowerBound *int64) *TopMetricsSvmUserCollectionGetParams {
	o.SetIopsErrorLowerBound(iopsErrorLowerBound)
	return o
}

// SetIopsErrorLowerBound adds the iopsErrorLowerBound to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetIopsErrorLowerBound(iopsErrorLowerBound *int64) {
	o.IopsErrorLowerBound = iopsErrorLowerBound
}

// WithIopsErrorUpperBound adds the iopsErrorUpperBound to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithIopsErrorUpperBound(iopsErrorUpperBound *int64) *TopMetricsSvmUserCollectionGetParams {
	o.SetIopsErrorUpperBound(iopsErrorUpperBound)
	return o
}

// SetIopsErrorUpperBound adds the iopsErrorUpperBound to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetIopsErrorUpperBound(iopsErrorUpperBound *int64) {
	o.IopsErrorUpperBound = iopsErrorUpperBound
}

// WithIopsRead adds the iopsRead to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithIopsRead(iopsRead *int64) *TopMetricsSvmUserCollectionGetParams {
	o.SetIopsRead(iopsRead)
	return o
}

// SetIopsRead adds the iopsRead to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetIopsRead(iopsRead *int64) {
	o.IopsRead = iopsRead
}

// WithIopsWrite adds the iopsWrite to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithIopsWrite(iopsWrite *int64) *TopMetricsSvmUserCollectionGetParams {
	o.SetIopsWrite(iopsWrite)
	return o
}

// SetIopsWrite adds the iopsWrite to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetIopsWrite(iopsWrite *int64) {
	o.IopsWrite = iopsWrite
}

// WithMaxRecords adds the maxRecords to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithMaxRecords(maxRecords *int64) *TopMetricsSvmUserCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithOrderBy(orderBy []string) *TopMetricsSvmUserCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithReturnRecords(returnRecords *bool) *TopMetricsSvmUserCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *TopMetricsSvmUserCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmName adds the svmName to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithSvmName(svmName *string) *TopMetricsSvmUserCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithSvmUUID(svmUUID string) *TopMetricsSvmUserCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WithThroughputErrorLowerBound adds the throughputErrorLowerBound to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithThroughputErrorLowerBound(throughputErrorLowerBound *int64) *TopMetricsSvmUserCollectionGetParams {
	o.SetThroughputErrorLowerBound(throughputErrorLowerBound)
	return o
}

// SetThroughputErrorLowerBound adds the throughputErrorLowerBound to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetThroughputErrorLowerBound(throughputErrorLowerBound *int64) {
	o.ThroughputErrorLowerBound = throughputErrorLowerBound
}

// WithThroughputErrorUpperBound adds the throughputErrorUpperBound to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithThroughputErrorUpperBound(throughputErrorUpperBound *int64) *TopMetricsSvmUserCollectionGetParams {
	o.SetThroughputErrorUpperBound(throughputErrorUpperBound)
	return o
}

// SetThroughputErrorUpperBound adds the throughputErrorUpperBound to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetThroughputErrorUpperBound(throughputErrorUpperBound *int64) {
	o.ThroughputErrorUpperBound = throughputErrorUpperBound
}

// WithThroughputRead adds the throughputRead to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithThroughputRead(throughputRead *int64) *TopMetricsSvmUserCollectionGetParams {
	o.SetThroughputRead(throughputRead)
	return o
}

// SetThroughputRead adds the throughputRead to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetThroughputRead(throughputRead *int64) {
	o.ThroughputRead = throughputRead
}

// WithThroughputWrite adds the throughputWrite to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithThroughputWrite(throughputWrite *int64) *TopMetricsSvmUserCollectionGetParams {
	o.SetThroughputWrite(throughputWrite)
	return o
}

// SetThroughputWrite adds the throughputWrite to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetThroughputWrite(throughputWrite *int64) {
	o.ThroughputWrite = throughputWrite
}

// WithTopMetric adds the topMetric to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithTopMetric(topMetric *string) *TopMetricsSvmUserCollectionGetParams {
	o.SetTopMetric(topMetric)
	return o
}

// SetTopMetric adds the topMetric to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetTopMetric(topMetric *string) {
	o.TopMetric = topMetric
}

// WithUserID adds the userID to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithUserID(userID *string) *TopMetricsSvmUserCollectionGetParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WithUserName adds the userName to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithUserName(userName *string) *TopMetricsSvmUserCollectionGetParams {
	o.SetUserName(userName)
	return o
}

// SetUserName adds the userName to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetUserName(userName *string) {
	o.UserName = userName
}

// WithVolumesName adds the volumesName to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithVolumesName(volumesName *string) *TopMetricsSvmUserCollectionGetParams {
	o.SetVolumesName(volumesName)
	return o
}

// SetVolumesName adds the volumesName to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetVolumesName(volumesName *string) {
	o.VolumesName = volumesName
}

// WithVolumesUUID adds the volumesUUID to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) WithVolumesUUID(volumesUUID *string) *TopMetricsSvmUserCollectionGetParams {
	o.SetVolumesUUID(volumesUUID)
	return o
}

// SetVolumesUUID adds the volumesUuid to the top metrics svm user collection get params
func (o *TopMetricsSvmUserCollectionGetParams) SetVolumesUUID(volumesUUID *string) {
	o.VolumesUUID = volumesUUID
}

// WriteToRequest writes these params to a swagger request
func (o *TopMetricsSvmUserCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IopsErrorLowerBound != nil {

		// query param iops.error.lower_bound
		var qrIopsErrorLowerBound int64

		if o.IopsErrorLowerBound != nil {
			qrIopsErrorLowerBound = *o.IopsErrorLowerBound
		}
		qIopsErrorLowerBound := swag.FormatInt64(qrIopsErrorLowerBound)
		if qIopsErrorLowerBound != "" {

			if err := r.SetQueryParam("iops.error.lower_bound", qIopsErrorLowerBound); err != nil {
				return err
			}
		}
	}

	if o.IopsErrorUpperBound != nil {

		// query param iops.error.upper_bound
		var qrIopsErrorUpperBound int64

		if o.IopsErrorUpperBound != nil {
			qrIopsErrorUpperBound = *o.IopsErrorUpperBound
		}
		qIopsErrorUpperBound := swag.FormatInt64(qrIopsErrorUpperBound)
		if qIopsErrorUpperBound != "" {

			if err := r.SetQueryParam("iops.error.upper_bound", qIopsErrorUpperBound); err != nil {
				return err
			}
		}
	}

	if o.IopsRead != nil {

		// query param iops.read
		var qrIopsRead int64

		if o.IopsRead != nil {
			qrIopsRead = *o.IopsRead
		}
		qIopsRead := swag.FormatInt64(qrIopsRead)
		if qIopsRead != "" {

			if err := r.SetQueryParam("iops.read", qIopsRead); err != nil {
				return err
			}
		}
	}

	if o.IopsWrite != nil {

		// query param iops.write
		var qrIopsWrite int64

		if o.IopsWrite != nil {
			qrIopsWrite = *o.IopsWrite
		}
		qIopsWrite := swag.FormatInt64(qrIopsWrite)
		if qIopsWrite != "" {

			if err := r.SetQueryParam("iops.write", qIopsWrite); err != nil {
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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if o.ThroughputErrorLowerBound != nil {

		// query param throughput.error.lower_bound
		var qrThroughputErrorLowerBound int64

		if o.ThroughputErrorLowerBound != nil {
			qrThroughputErrorLowerBound = *o.ThroughputErrorLowerBound
		}
		qThroughputErrorLowerBound := swag.FormatInt64(qrThroughputErrorLowerBound)
		if qThroughputErrorLowerBound != "" {

			if err := r.SetQueryParam("throughput.error.lower_bound", qThroughputErrorLowerBound); err != nil {
				return err
			}
		}
	}

	if o.ThroughputErrorUpperBound != nil {

		// query param throughput.error.upper_bound
		var qrThroughputErrorUpperBound int64

		if o.ThroughputErrorUpperBound != nil {
			qrThroughputErrorUpperBound = *o.ThroughputErrorUpperBound
		}
		qThroughputErrorUpperBound := swag.FormatInt64(qrThroughputErrorUpperBound)
		if qThroughputErrorUpperBound != "" {

			if err := r.SetQueryParam("throughput.error.upper_bound", qThroughputErrorUpperBound); err != nil {
				return err
			}
		}
	}

	if o.ThroughputRead != nil {

		// query param throughput.read
		var qrThroughputRead int64

		if o.ThroughputRead != nil {
			qrThroughputRead = *o.ThroughputRead
		}
		qThroughputRead := swag.FormatInt64(qrThroughputRead)
		if qThroughputRead != "" {

			if err := r.SetQueryParam("throughput.read", qThroughputRead); err != nil {
				return err
			}
		}
	}

	if o.ThroughputWrite != nil {

		// query param throughput.write
		var qrThroughputWrite int64

		if o.ThroughputWrite != nil {
			qrThroughputWrite = *o.ThroughputWrite
		}
		qThroughputWrite := swag.FormatInt64(qrThroughputWrite)
		if qThroughputWrite != "" {

			if err := r.SetQueryParam("throughput.write", qThroughputWrite); err != nil {
				return err
			}
		}
	}

	if o.TopMetric != nil {

		// query param top_metric
		var qrTopMetric string

		if o.TopMetric != nil {
			qrTopMetric = *o.TopMetric
		}
		qTopMetric := qrTopMetric
		if qTopMetric != "" {

			if err := r.SetQueryParam("top_metric", qTopMetric); err != nil {
				return err
			}
		}
	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID string

		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {

			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}
	}

	if o.UserName != nil {

		// query param user_name
		var qrUserName string

		if o.UserName != nil {
			qrUserName = *o.UserName
		}
		qUserName := qrUserName
		if qUserName != "" {

			if err := r.SetQueryParam("user_name", qUserName); err != nil {
				return err
			}
		}
	}

	if o.VolumesName != nil {

		// query param volumes.name
		var qrVolumesName string

		if o.VolumesName != nil {
			qrVolumesName = *o.VolumesName
		}
		qVolumesName := qrVolumesName
		if qVolumesName != "" {

			if err := r.SetQueryParam("volumes.name", qVolumesName); err != nil {
				return err
			}
		}
	}

	if o.VolumesUUID != nil {

		// query param volumes.uuid
		var qrVolumesUUID string

		if o.VolumesUUID != nil {
			qrVolumesUUID = *o.VolumesUUID
		}
		qVolumesUUID := qrVolumesUUID
		if qVolumesUUID != "" {

			if err := r.SetQueryParam("volumes.uuid", qVolumesUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamTopMetricsSvmUserCollectionGet binds the parameter fields
func (o *TopMetricsSvmUserCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamTopMetricsSvmUserCollectionGet binds the parameter order_by
func (o *TopMetricsSvmUserCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
