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

// NewS3AuditGetParams creates a new S3AuditGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3AuditGetParams() *S3AuditGetParams {
	return &S3AuditGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3AuditGetParamsWithTimeout creates a new S3AuditGetParams object
// with the ability to set a timeout on a request.
func NewS3AuditGetParamsWithTimeout(timeout time.Duration) *S3AuditGetParams {
	return &S3AuditGetParams{
		timeout: timeout,
	}
}

// NewS3AuditGetParamsWithContext creates a new S3AuditGetParams object
// with the ability to set a context for a request.
func NewS3AuditGetParamsWithContext(ctx context.Context) *S3AuditGetParams {
	return &S3AuditGetParams{
		Context: ctx,
	}
}

// NewS3AuditGetParamsWithHTTPClient creates a new S3AuditGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3AuditGetParamsWithHTTPClient(client *http.Client) *S3AuditGetParams {
	return &S3AuditGetParams{
		HTTPClient: client,
	}
}

/*
S3AuditGetParams contains all the parameters to send to the API endpoint

	for the s3 audit get operation.

	Typically these are written to a http.Request.
*/
type S3AuditGetParams struct {

	/* Enabled.

	   Filter by enabled
	*/
	Enabled *bool

	/* EventsData.

	   Filter by events.data
	*/
	EventsData *bool

	/* EventsManagement.

	   Filter by events.management
	*/
	EventsManagement *bool

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* LogFormat.

	   Filter by log.format
	*/
	LogFormat *string

	/* LogRetentionCount.

	   Filter by log.retention.count
	*/
	LogRetentionCount *int64

	/* LogRetentionDuration.

	   Filter by log.retention.duration
	*/
	LogRetentionDuration *string

	/* LogRotationScheduleDays.

	   Filter by log.rotation.schedule.days
	*/
	LogRotationScheduleDays *int64

	/* LogRotationScheduleHours.

	   Filter by log.rotation.schedule.hours
	*/
	LogRotationScheduleHours *int64

	/* LogRotationScheduleMinutes.

	   Filter by log.rotation.schedule.minutes
	*/
	LogRotationScheduleMinutes *int64

	/* LogRotationScheduleMonths.

	   Filter by log.rotation.schedule.months
	*/
	LogRotationScheduleMonths *int64

	/* LogRotationScheduleWeekdays.

	   Filter by log.rotation.schedule.weekdays
	*/
	LogRotationScheduleWeekdays *int64

	/* LogRotationSize.

	   Filter by log.rotation.size
	*/
	LogRotationSize *int64

	/* LogPath.

	   Filter by log_path
	*/
	LogPath *string

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

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 audit get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3AuditGetParams) WithDefaults() *S3AuditGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 audit get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3AuditGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := S3AuditGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the s3 audit get params
func (o *S3AuditGetParams) WithTimeout(timeout time.Duration) *S3AuditGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 audit get params
func (o *S3AuditGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 audit get params
func (o *S3AuditGetParams) WithContext(ctx context.Context) *S3AuditGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 audit get params
func (o *S3AuditGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 audit get params
func (o *S3AuditGetParams) WithHTTPClient(client *http.Client) *S3AuditGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 audit get params
func (o *S3AuditGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabled adds the enabled to the s3 audit get params
func (o *S3AuditGetParams) WithEnabled(enabled *bool) *S3AuditGetParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the s3 audit get params
func (o *S3AuditGetParams) SetEnabled(enabled *bool) {
	o.Enabled = enabled
}

// WithEventsData adds the eventsData to the s3 audit get params
func (o *S3AuditGetParams) WithEventsData(eventsData *bool) *S3AuditGetParams {
	o.SetEventsData(eventsData)
	return o
}

// SetEventsData adds the eventsData to the s3 audit get params
func (o *S3AuditGetParams) SetEventsData(eventsData *bool) {
	o.EventsData = eventsData
}

// WithEventsManagement adds the eventsManagement to the s3 audit get params
func (o *S3AuditGetParams) WithEventsManagement(eventsManagement *bool) *S3AuditGetParams {
	o.SetEventsManagement(eventsManagement)
	return o
}

// SetEventsManagement adds the eventsManagement to the s3 audit get params
func (o *S3AuditGetParams) SetEventsManagement(eventsManagement *bool) {
	o.EventsManagement = eventsManagement
}

// WithFields adds the fields to the s3 audit get params
func (o *S3AuditGetParams) WithFields(fields []string) *S3AuditGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the s3 audit get params
func (o *S3AuditGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithLogFormat adds the logFormat to the s3 audit get params
func (o *S3AuditGetParams) WithLogFormat(logFormat *string) *S3AuditGetParams {
	o.SetLogFormat(logFormat)
	return o
}

// SetLogFormat adds the logFormat to the s3 audit get params
func (o *S3AuditGetParams) SetLogFormat(logFormat *string) {
	o.LogFormat = logFormat
}

// WithLogRetentionCount adds the logRetentionCount to the s3 audit get params
func (o *S3AuditGetParams) WithLogRetentionCount(logRetentionCount *int64) *S3AuditGetParams {
	o.SetLogRetentionCount(logRetentionCount)
	return o
}

// SetLogRetentionCount adds the logRetentionCount to the s3 audit get params
func (o *S3AuditGetParams) SetLogRetentionCount(logRetentionCount *int64) {
	o.LogRetentionCount = logRetentionCount
}

// WithLogRetentionDuration adds the logRetentionDuration to the s3 audit get params
func (o *S3AuditGetParams) WithLogRetentionDuration(logRetentionDuration *string) *S3AuditGetParams {
	o.SetLogRetentionDuration(logRetentionDuration)
	return o
}

// SetLogRetentionDuration adds the logRetentionDuration to the s3 audit get params
func (o *S3AuditGetParams) SetLogRetentionDuration(logRetentionDuration *string) {
	o.LogRetentionDuration = logRetentionDuration
}

// WithLogRotationScheduleDays adds the logRotationScheduleDays to the s3 audit get params
func (o *S3AuditGetParams) WithLogRotationScheduleDays(logRotationScheduleDays *int64) *S3AuditGetParams {
	o.SetLogRotationScheduleDays(logRotationScheduleDays)
	return o
}

// SetLogRotationScheduleDays adds the logRotationScheduleDays to the s3 audit get params
func (o *S3AuditGetParams) SetLogRotationScheduleDays(logRotationScheduleDays *int64) {
	o.LogRotationScheduleDays = logRotationScheduleDays
}

// WithLogRotationScheduleHours adds the logRotationScheduleHours to the s3 audit get params
func (o *S3AuditGetParams) WithLogRotationScheduleHours(logRotationScheduleHours *int64) *S3AuditGetParams {
	o.SetLogRotationScheduleHours(logRotationScheduleHours)
	return o
}

// SetLogRotationScheduleHours adds the logRotationScheduleHours to the s3 audit get params
func (o *S3AuditGetParams) SetLogRotationScheduleHours(logRotationScheduleHours *int64) {
	o.LogRotationScheduleHours = logRotationScheduleHours
}

// WithLogRotationScheduleMinutes adds the logRotationScheduleMinutes to the s3 audit get params
func (o *S3AuditGetParams) WithLogRotationScheduleMinutes(logRotationScheduleMinutes *int64) *S3AuditGetParams {
	o.SetLogRotationScheduleMinutes(logRotationScheduleMinutes)
	return o
}

// SetLogRotationScheduleMinutes adds the logRotationScheduleMinutes to the s3 audit get params
func (o *S3AuditGetParams) SetLogRotationScheduleMinutes(logRotationScheduleMinutes *int64) {
	o.LogRotationScheduleMinutes = logRotationScheduleMinutes
}

// WithLogRotationScheduleMonths adds the logRotationScheduleMonths to the s3 audit get params
func (o *S3AuditGetParams) WithLogRotationScheduleMonths(logRotationScheduleMonths *int64) *S3AuditGetParams {
	o.SetLogRotationScheduleMonths(logRotationScheduleMonths)
	return o
}

// SetLogRotationScheduleMonths adds the logRotationScheduleMonths to the s3 audit get params
func (o *S3AuditGetParams) SetLogRotationScheduleMonths(logRotationScheduleMonths *int64) {
	o.LogRotationScheduleMonths = logRotationScheduleMonths
}

// WithLogRotationScheduleWeekdays adds the logRotationScheduleWeekdays to the s3 audit get params
func (o *S3AuditGetParams) WithLogRotationScheduleWeekdays(logRotationScheduleWeekdays *int64) *S3AuditGetParams {
	o.SetLogRotationScheduleWeekdays(logRotationScheduleWeekdays)
	return o
}

// SetLogRotationScheduleWeekdays adds the logRotationScheduleWeekdays to the s3 audit get params
func (o *S3AuditGetParams) SetLogRotationScheduleWeekdays(logRotationScheduleWeekdays *int64) {
	o.LogRotationScheduleWeekdays = logRotationScheduleWeekdays
}

// WithLogRotationSize adds the logRotationSize to the s3 audit get params
func (o *S3AuditGetParams) WithLogRotationSize(logRotationSize *int64) *S3AuditGetParams {
	o.SetLogRotationSize(logRotationSize)
	return o
}

// SetLogRotationSize adds the logRotationSize to the s3 audit get params
func (o *S3AuditGetParams) SetLogRotationSize(logRotationSize *int64) {
	o.LogRotationSize = logRotationSize
}

// WithLogPath adds the logPath to the s3 audit get params
func (o *S3AuditGetParams) WithLogPath(logPath *string) *S3AuditGetParams {
	o.SetLogPath(logPath)
	return o
}

// SetLogPath adds the logPath to the s3 audit get params
func (o *S3AuditGetParams) SetLogPath(logPath *string) {
	o.LogPath = logPath
}

// WithMaxRecords adds the maxRecords to the s3 audit get params
func (o *S3AuditGetParams) WithMaxRecords(maxRecords *int64) *S3AuditGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the s3 audit get params
func (o *S3AuditGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the s3 audit get params
func (o *S3AuditGetParams) WithOrderBy(orderBy []string) *S3AuditGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the s3 audit get params
func (o *S3AuditGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the s3 audit get params
func (o *S3AuditGetParams) WithReturnRecords(returnRecords *bool) *S3AuditGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the s3 audit get params
func (o *S3AuditGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the s3 audit get params
func (o *S3AuditGetParams) WithReturnTimeout(returnTimeout *int64) *S3AuditGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the s3 audit get params
func (o *S3AuditGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmName adds the svmName to the s3 audit get params
func (o *S3AuditGetParams) WithSvmName(svmName *string) *S3AuditGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the s3 audit get params
func (o *S3AuditGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the s3 audit get params
func (o *S3AuditGetParams) WithSvmUUID(svmUUID string) *S3AuditGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the s3 audit get params
func (o *S3AuditGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *S3AuditGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EventsData != nil {

		// query param events.data
		var qrEventsData bool

		if o.EventsData != nil {
			qrEventsData = *o.EventsData
		}
		qEventsData := swag.FormatBool(qrEventsData)
		if qEventsData != "" {

			if err := r.SetQueryParam("events.data", qEventsData); err != nil {
				return err
			}
		}
	}

	if o.EventsManagement != nil {

		// query param events.management
		var qrEventsManagement bool

		if o.EventsManagement != nil {
			qrEventsManagement = *o.EventsManagement
		}
		qEventsManagement := swag.FormatBool(qrEventsManagement)
		if qEventsManagement != "" {

			if err := r.SetQueryParam("events.management", qEventsManagement); err != nil {
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

	if o.LogFormat != nil {

		// query param log.format
		var qrLogFormat string

		if o.LogFormat != nil {
			qrLogFormat = *o.LogFormat
		}
		qLogFormat := qrLogFormat
		if qLogFormat != "" {

			if err := r.SetQueryParam("log.format", qLogFormat); err != nil {
				return err
			}
		}
	}

	if o.LogRetentionCount != nil {

		// query param log.retention.count
		var qrLogRetentionCount int64

		if o.LogRetentionCount != nil {
			qrLogRetentionCount = *o.LogRetentionCount
		}
		qLogRetentionCount := swag.FormatInt64(qrLogRetentionCount)
		if qLogRetentionCount != "" {

			if err := r.SetQueryParam("log.retention.count", qLogRetentionCount); err != nil {
				return err
			}
		}
	}

	if o.LogRetentionDuration != nil {

		// query param log.retention.duration
		var qrLogRetentionDuration string

		if o.LogRetentionDuration != nil {
			qrLogRetentionDuration = *o.LogRetentionDuration
		}
		qLogRetentionDuration := qrLogRetentionDuration
		if qLogRetentionDuration != "" {

			if err := r.SetQueryParam("log.retention.duration", qLogRetentionDuration); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleDays != nil {

		// query param log.rotation.schedule.days
		var qrLogRotationScheduleDays int64

		if o.LogRotationScheduleDays != nil {
			qrLogRotationScheduleDays = *o.LogRotationScheduleDays
		}
		qLogRotationScheduleDays := swag.FormatInt64(qrLogRotationScheduleDays)
		if qLogRotationScheduleDays != "" {

			if err := r.SetQueryParam("log.rotation.schedule.days", qLogRotationScheduleDays); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleHours != nil {

		// query param log.rotation.schedule.hours
		var qrLogRotationScheduleHours int64

		if o.LogRotationScheduleHours != nil {
			qrLogRotationScheduleHours = *o.LogRotationScheduleHours
		}
		qLogRotationScheduleHours := swag.FormatInt64(qrLogRotationScheduleHours)
		if qLogRotationScheduleHours != "" {

			if err := r.SetQueryParam("log.rotation.schedule.hours", qLogRotationScheduleHours); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleMinutes != nil {

		// query param log.rotation.schedule.minutes
		var qrLogRotationScheduleMinutes int64

		if o.LogRotationScheduleMinutes != nil {
			qrLogRotationScheduleMinutes = *o.LogRotationScheduleMinutes
		}
		qLogRotationScheduleMinutes := swag.FormatInt64(qrLogRotationScheduleMinutes)
		if qLogRotationScheduleMinutes != "" {

			if err := r.SetQueryParam("log.rotation.schedule.minutes", qLogRotationScheduleMinutes); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleMonths != nil {

		// query param log.rotation.schedule.months
		var qrLogRotationScheduleMonths int64

		if o.LogRotationScheduleMonths != nil {
			qrLogRotationScheduleMonths = *o.LogRotationScheduleMonths
		}
		qLogRotationScheduleMonths := swag.FormatInt64(qrLogRotationScheduleMonths)
		if qLogRotationScheduleMonths != "" {

			if err := r.SetQueryParam("log.rotation.schedule.months", qLogRotationScheduleMonths); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleWeekdays != nil {

		// query param log.rotation.schedule.weekdays
		var qrLogRotationScheduleWeekdays int64

		if o.LogRotationScheduleWeekdays != nil {
			qrLogRotationScheduleWeekdays = *o.LogRotationScheduleWeekdays
		}
		qLogRotationScheduleWeekdays := swag.FormatInt64(qrLogRotationScheduleWeekdays)
		if qLogRotationScheduleWeekdays != "" {

			if err := r.SetQueryParam("log.rotation.schedule.weekdays", qLogRotationScheduleWeekdays); err != nil {
				return err
			}
		}
	}

	if o.LogRotationSize != nil {

		// query param log.rotation.size
		var qrLogRotationSize int64

		if o.LogRotationSize != nil {
			qrLogRotationSize = *o.LogRotationSize
		}
		qLogRotationSize := swag.FormatInt64(qrLogRotationSize)
		if qLogRotationSize != "" {

			if err := r.SetQueryParam("log.rotation.size", qLogRotationSize); err != nil {
				return err
			}
		}
	}

	if o.LogPath != nil {

		// query param log_path
		var qrLogPath string

		if o.LogPath != nil {
			qrLogPath = *o.LogPath
		}
		qLogPath := qrLogPath
		if qLogPath != "" {

			if err := r.SetQueryParam("log_path", qLogPath); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamS3AuditGet binds the parameter fields
func (o *S3AuditGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamS3AuditGet binds the parameter order_by
func (o *S3AuditGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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