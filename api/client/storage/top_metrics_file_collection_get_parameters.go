// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewTopMetricsFileCollectionGetParams creates a new TopMetricsFileCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTopMetricsFileCollectionGetParams() *TopMetricsFileCollectionGetParams {
	return &TopMetricsFileCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTopMetricsFileCollectionGetParamsWithTimeout creates a new TopMetricsFileCollectionGetParams object
// with the ability to set a timeout on a request.
func NewTopMetricsFileCollectionGetParamsWithTimeout(timeout time.Duration) *TopMetricsFileCollectionGetParams {
	return &TopMetricsFileCollectionGetParams{
		timeout: timeout,
	}
}

// NewTopMetricsFileCollectionGetParamsWithContext creates a new TopMetricsFileCollectionGetParams object
// with the ability to set a context for a request.
func NewTopMetricsFileCollectionGetParamsWithContext(ctx context.Context) *TopMetricsFileCollectionGetParams {
	return &TopMetricsFileCollectionGetParams{
		Context: ctx,
	}
}

// NewTopMetricsFileCollectionGetParamsWithHTTPClient creates a new TopMetricsFileCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewTopMetricsFileCollectionGetParamsWithHTTPClient(client *http.Client) *TopMetricsFileCollectionGetParams {
	return &TopMetricsFileCollectionGetParams{
		HTTPClient: client,
	}
}

/*
TopMetricsFileCollectionGetParams contains all the parameters to send to the API endpoint

	for the top metrics file collection get operation.

	Typically these are written to a http.Request.
*/
type TopMetricsFileCollectionGetParams struct {

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

	/* MaxRecordsPerVolume.

	   Max records per volume.
	*/
	MaxRecordsPerVolume *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* Path.

	   Filter by path
	*/
	Path *string

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

	/* VolumeName.

	   Filter by volume.name
	*/
	VolumeName *string

	/* VolumeUUID.

	   Volume UUID
	*/
	VolumeUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the top metrics file collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TopMetricsFileCollectionGetParams) WithDefaults() *TopMetricsFileCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the top metrics file collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TopMetricsFileCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		topMetricDefault = string("iops.read")
	)

	val := TopMetricsFileCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
		TopMetric:     &topMetricDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithTimeout(timeout time.Duration) *TopMetricsFileCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithContext(ctx context.Context) *TopMetricsFileCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithHTTPClient(client *http.Client) *TopMetricsFileCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithFields(fields []string) *TopMetricsFileCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIopsErrorLowerBound adds the iopsErrorLowerBound to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithIopsErrorLowerBound(iopsErrorLowerBound *int64) *TopMetricsFileCollectionGetParams {
	o.SetIopsErrorLowerBound(iopsErrorLowerBound)
	return o
}

// SetIopsErrorLowerBound adds the iopsErrorLowerBound to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetIopsErrorLowerBound(iopsErrorLowerBound *int64) {
	o.IopsErrorLowerBound = iopsErrorLowerBound
}

// WithIopsErrorUpperBound adds the iopsErrorUpperBound to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithIopsErrorUpperBound(iopsErrorUpperBound *int64) *TopMetricsFileCollectionGetParams {
	o.SetIopsErrorUpperBound(iopsErrorUpperBound)
	return o
}

// SetIopsErrorUpperBound adds the iopsErrorUpperBound to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetIopsErrorUpperBound(iopsErrorUpperBound *int64) {
	o.IopsErrorUpperBound = iopsErrorUpperBound
}

// WithIopsRead adds the iopsRead to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithIopsRead(iopsRead *int64) *TopMetricsFileCollectionGetParams {
	o.SetIopsRead(iopsRead)
	return o
}

// SetIopsRead adds the iopsRead to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetIopsRead(iopsRead *int64) {
	o.IopsRead = iopsRead
}

// WithIopsWrite adds the iopsWrite to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithIopsWrite(iopsWrite *int64) *TopMetricsFileCollectionGetParams {
	o.SetIopsWrite(iopsWrite)
	return o
}

// SetIopsWrite adds the iopsWrite to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetIopsWrite(iopsWrite *int64) {
	o.IopsWrite = iopsWrite
}

// WithMaxRecords adds the maxRecords to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithMaxRecords(maxRecords *int64) *TopMetricsFileCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithMaxRecordsPerVolume adds the maxRecordsPerVolume to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithMaxRecordsPerVolume(maxRecordsPerVolume *int64) *TopMetricsFileCollectionGetParams {
	o.SetMaxRecordsPerVolume(maxRecordsPerVolume)
	return o
}

// SetMaxRecordsPerVolume adds the maxRecordsPerVolume to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetMaxRecordsPerVolume(maxRecordsPerVolume *int64) {
	o.MaxRecordsPerVolume = maxRecordsPerVolume
}

// WithOrderBy adds the orderBy to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithOrderBy(orderBy []string) *TopMetricsFileCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPath adds the path to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithPath(path *string) *TopMetricsFileCollectionGetParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetPath(path *string) {
	o.Path = path
}

// WithReturnRecords adds the returnRecords to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithReturnRecords(returnRecords *bool) *TopMetricsFileCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *TopMetricsFileCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmName adds the svmName to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithSvmName(svmName *string) *TopMetricsFileCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithSvmUUID(svmUUID *string) *TopMetricsFileCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithThroughputErrorLowerBound adds the throughputErrorLowerBound to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithThroughputErrorLowerBound(throughputErrorLowerBound *int64) *TopMetricsFileCollectionGetParams {
	o.SetThroughputErrorLowerBound(throughputErrorLowerBound)
	return o
}

// SetThroughputErrorLowerBound adds the throughputErrorLowerBound to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetThroughputErrorLowerBound(throughputErrorLowerBound *int64) {
	o.ThroughputErrorLowerBound = throughputErrorLowerBound
}

// WithThroughputErrorUpperBound adds the throughputErrorUpperBound to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithThroughputErrorUpperBound(throughputErrorUpperBound *int64) *TopMetricsFileCollectionGetParams {
	o.SetThroughputErrorUpperBound(throughputErrorUpperBound)
	return o
}

// SetThroughputErrorUpperBound adds the throughputErrorUpperBound to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetThroughputErrorUpperBound(throughputErrorUpperBound *int64) {
	o.ThroughputErrorUpperBound = throughputErrorUpperBound
}

// WithThroughputRead adds the throughputRead to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithThroughputRead(throughputRead *int64) *TopMetricsFileCollectionGetParams {
	o.SetThroughputRead(throughputRead)
	return o
}

// SetThroughputRead adds the throughputRead to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetThroughputRead(throughputRead *int64) {
	o.ThroughputRead = throughputRead
}

// WithThroughputWrite adds the throughputWrite to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithThroughputWrite(throughputWrite *int64) *TopMetricsFileCollectionGetParams {
	o.SetThroughputWrite(throughputWrite)
	return o
}

// SetThroughputWrite adds the throughputWrite to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetThroughputWrite(throughputWrite *int64) {
	o.ThroughputWrite = throughputWrite
}

// WithTopMetric adds the topMetric to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithTopMetric(topMetric *string) *TopMetricsFileCollectionGetParams {
	o.SetTopMetric(topMetric)
	return o
}

// SetTopMetric adds the topMetric to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetTopMetric(topMetric *string) {
	o.TopMetric = topMetric
}

// WithVolumeName adds the volumeName to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithVolumeName(volumeName *string) *TopMetricsFileCollectionGetParams {
	o.SetVolumeName(volumeName)
	return o
}

// SetVolumeName adds the volumeName to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetVolumeName(volumeName *string) {
	o.VolumeName = volumeName
}

// WithVolumeUUID adds the volumeUUID to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) WithVolumeUUID(volumeUUID string) *TopMetricsFileCollectionGetParams {
	o.SetVolumeUUID(volumeUUID)
	return o
}

// SetVolumeUUID adds the volumeUuid to the top metrics file collection get params
func (o *TopMetricsFileCollectionGetParams) SetVolumeUUID(volumeUUID string) {
	o.VolumeUUID = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *TopMetricsFileCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.MaxRecordsPerVolume != nil {

		// query param max_records_per_volume
		var qrMaxRecordsPerVolume int64

		if o.MaxRecordsPerVolume != nil {
			qrMaxRecordsPerVolume = *o.MaxRecordsPerVolume
		}
		qMaxRecordsPerVolume := swag.FormatInt64(qrMaxRecordsPerVolume)
		if qMaxRecordsPerVolume != "" {

			if err := r.SetQueryParam("max_records_per_volume", qMaxRecordsPerVolume); err != nil {
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

	if o.Path != nil {

		// query param path
		var qrPath string

		if o.Path != nil {
			qrPath = *o.Path
		}
		qPath := qrPath
		if qPath != "" {

			if err := r.SetQueryParam("path", qPath); err != nil {
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

	if o.VolumeName != nil {

		// query param volume.name
		var qrVolumeName string

		if o.VolumeName != nil {
			qrVolumeName = *o.VolumeName
		}
		qVolumeName := qrVolumeName
		if qVolumeName != "" {

			if err := r.SetQueryParam("volume.name", qVolumeName); err != nil {
				return err
			}
		}
	}

	// path param volume.uuid
	if err := r.SetPathParam("volume.uuid", o.VolumeUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamTopMetricsFileCollectionGet binds the parameter fields
func (o *TopMetricsFileCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamTopMetricsFileCollectionGet binds the parameter order_by
func (o *TopMetricsFileCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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