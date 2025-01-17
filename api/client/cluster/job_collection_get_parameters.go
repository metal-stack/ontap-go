// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewJobCollectionGetParams creates a new JobCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewJobCollectionGetParams() *JobCollectionGetParams {
	return &JobCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewJobCollectionGetParamsWithTimeout creates a new JobCollectionGetParams object
// with the ability to set a timeout on a request.
func NewJobCollectionGetParamsWithTimeout(timeout time.Duration) *JobCollectionGetParams {
	return &JobCollectionGetParams{
		timeout: timeout,
	}
}

// NewJobCollectionGetParamsWithContext creates a new JobCollectionGetParams object
// with the ability to set a context for a request.
func NewJobCollectionGetParamsWithContext(ctx context.Context) *JobCollectionGetParams {
	return &JobCollectionGetParams{
		Context: ctx,
	}
}

// NewJobCollectionGetParamsWithHTTPClient creates a new JobCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewJobCollectionGetParamsWithHTTPClient(client *http.Client) *JobCollectionGetParams {
	return &JobCollectionGetParams{
		HTTPClient: client,
	}
}

/*
JobCollectionGetParams contains all the parameters to send to the API endpoint

	for the job collection get operation.

	Typically these are written to a http.Request.
*/
type JobCollectionGetParams struct {

	/* Code.

	   Filter by code
	*/
	Code *int64

	/* Description.

	   Filter by description
	*/
	Description *string

	/* EndTime.

	   Filter by end_time
	*/
	EndTime *string

	/* ErrorArgumentsCode.

	   Filter by error.arguments.code
	*/
	ErrorArgumentsCode *string

	/* ErrorArgumentsMessage.

	   Filter by error.arguments.message
	*/
	ErrorArgumentsMessage *string

	/* ErrorCode.

	   Filter by error.code
	*/
	ErrorCode *string

	/* ErrorMessage.

	   Filter by error.message
	*/
	ErrorMessage *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Message.

	   Filter by message
	*/
	Message *string

	/* NodeName.

	   Filter by node.name
	*/
	NodeName *string

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

	/* StartTime.

	   Filter by start_time
	*/
	StartTime *string

	/* State.

	   Filter by state
	*/
	State *string

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the job collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobCollectionGetParams) WithDefaults() *JobCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the job collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := JobCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the job collection get params
func (o *JobCollectionGetParams) WithTimeout(timeout time.Duration) *JobCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job collection get params
func (o *JobCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job collection get params
func (o *JobCollectionGetParams) WithContext(ctx context.Context) *JobCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job collection get params
func (o *JobCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job collection get params
func (o *JobCollectionGetParams) WithHTTPClient(client *http.Client) *JobCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job collection get params
func (o *JobCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCode adds the code to the job collection get params
func (o *JobCollectionGetParams) WithCode(code *int64) *JobCollectionGetParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the job collection get params
func (o *JobCollectionGetParams) SetCode(code *int64) {
	o.Code = code
}

// WithDescription adds the description to the job collection get params
func (o *JobCollectionGetParams) WithDescription(description *string) *JobCollectionGetParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the job collection get params
func (o *JobCollectionGetParams) SetDescription(description *string) {
	o.Description = description
}

// WithEndTime adds the endTime to the job collection get params
func (o *JobCollectionGetParams) WithEndTime(endTime *string) *JobCollectionGetParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the job collection get params
func (o *JobCollectionGetParams) SetEndTime(endTime *string) {
	o.EndTime = endTime
}

// WithErrorArgumentsCode adds the errorArgumentsCode to the job collection get params
func (o *JobCollectionGetParams) WithErrorArgumentsCode(errorArgumentsCode *string) *JobCollectionGetParams {
	o.SetErrorArgumentsCode(errorArgumentsCode)
	return o
}

// SetErrorArgumentsCode adds the errorArgumentsCode to the job collection get params
func (o *JobCollectionGetParams) SetErrorArgumentsCode(errorArgumentsCode *string) {
	o.ErrorArgumentsCode = errorArgumentsCode
}

// WithErrorArgumentsMessage adds the errorArgumentsMessage to the job collection get params
func (o *JobCollectionGetParams) WithErrorArgumentsMessage(errorArgumentsMessage *string) *JobCollectionGetParams {
	o.SetErrorArgumentsMessage(errorArgumentsMessage)
	return o
}

// SetErrorArgumentsMessage adds the errorArgumentsMessage to the job collection get params
func (o *JobCollectionGetParams) SetErrorArgumentsMessage(errorArgumentsMessage *string) {
	o.ErrorArgumentsMessage = errorArgumentsMessage
}

// WithErrorCode adds the errorCode to the job collection get params
func (o *JobCollectionGetParams) WithErrorCode(errorCode *string) *JobCollectionGetParams {
	o.SetErrorCode(errorCode)
	return o
}

// SetErrorCode adds the errorCode to the job collection get params
func (o *JobCollectionGetParams) SetErrorCode(errorCode *string) {
	o.ErrorCode = errorCode
}

// WithErrorMessage adds the errorMessage to the job collection get params
func (o *JobCollectionGetParams) WithErrorMessage(errorMessage *string) *JobCollectionGetParams {
	o.SetErrorMessage(errorMessage)
	return o
}

// SetErrorMessage adds the errorMessage to the job collection get params
func (o *JobCollectionGetParams) SetErrorMessage(errorMessage *string) {
	o.ErrorMessage = errorMessage
}

// WithFields adds the fields to the job collection get params
func (o *JobCollectionGetParams) WithFields(fields []string) *JobCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the job collection get params
func (o *JobCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the job collection get params
func (o *JobCollectionGetParams) WithMaxRecords(maxRecords *int64) *JobCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the job collection get params
func (o *JobCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithMessage adds the message to the job collection get params
func (o *JobCollectionGetParams) WithMessage(message *string) *JobCollectionGetParams {
	o.SetMessage(message)
	return o
}

// SetMessage adds the message to the job collection get params
func (o *JobCollectionGetParams) SetMessage(message *string) {
	o.Message = message
}

// WithNodeName adds the nodeName to the job collection get params
func (o *JobCollectionGetParams) WithNodeName(nodeName *string) *JobCollectionGetParams {
	o.SetNodeName(nodeName)
	return o
}

// SetNodeName adds the nodeName to the job collection get params
func (o *JobCollectionGetParams) SetNodeName(nodeName *string) {
	o.NodeName = nodeName
}

// WithOrderBy adds the orderBy to the job collection get params
func (o *JobCollectionGetParams) WithOrderBy(orderBy []string) *JobCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the job collection get params
func (o *JobCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the job collection get params
func (o *JobCollectionGetParams) WithReturnRecords(returnRecords *bool) *JobCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the job collection get params
func (o *JobCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the job collection get params
func (o *JobCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *JobCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the job collection get params
func (o *JobCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithStartTime adds the startTime to the job collection get params
func (o *JobCollectionGetParams) WithStartTime(startTime *string) *JobCollectionGetParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the job collection get params
func (o *JobCollectionGetParams) SetStartTime(startTime *string) {
	o.StartTime = startTime
}

// WithState adds the state to the job collection get params
func (o *JobCollectionGetParams) WithState(state *string) *JobCollectionGetParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the job collection get params
func (o *JobCollectionGetParams) SetState(state *string) {
	o.State = state
}

// WithSvmName adds the svmName to the job collection get params
func (o *JobCollectionGetParams) WithSvmName(svmName *string) *JobCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the job collection get params
func (o *JobCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the job collection get params
func (o *JobCollectionGetParams) WithSvmUUID(svmUUID *string) *JobCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the job collection get params
func (o *JobCollectionGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithUUID adds the uuid to the job collection get params
func (o *JobCollectionGetParams) WithUUID(uuid *string) *JobCollectionGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the job collection get params
func (o *JobCollectionGetParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *JobCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Code != nil {

		// query param code
		var qrCode int64

		if o.Code != nil {
			qrCode = *o.Code
		}
		qCode := swag.FormatInt64(qrCode)
		if qCode != "" {

			if err := r.SetQueryParam("code", qCode); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.EndTime != nil {

		// query param end_time
		var qrEndTime string

		if o.EndTime != nil {
			qrEndTime = *o.EndTime
		}
		qEndTime := qrEndTime
		if qEndTime != "" {

			if err := r.SetQueryParam("end_time", qEndTime); err != nil {
				return err
			}
		}
	}

	if o.ErrorArgumentsCode != nil {

		// query param error.arguments.code
		var qrErrorArgumentsCode string

		if o.ErrorArgumentsCode != nil {
			qrErrorArgumentsCode = *o.ErrorArgumentsCode
		}
		qErrorArgumentsCode := qrErrorArgumentsCode
		if qErrorArgumentsCode != "" {

			if err := r.SetQueryParam("error.arguments.code", qErrorArgumentsCode); err != nil {
				return err
			}
		}
	}

	if o.ErrorArgumentsMessage != nil {

		// query param error.arguments.message
		var qrErrorArgumentsMessage string

		if o.ErrorArgumentsMessage != nil {
			qrErrorArgumentsMessage = *o.ErrorArgumentsMessage
		}
		qErrorArgumentsMessage := qrErrorArgumentsMessage
		if qErrorArgumentsMessage != "" {

			if err := r.SetQueryParam("error.arguments.message", qErrorArgumentsMessage); err != nil {
				return err
			}
		}
	}

	if o.ErrorCode != nil {

		// query param error.code
		var qrErrorCode string

		if o.ErrorCode != nil {
			qrErrorCode = *o.ErrorCode
		}
		qErrorCode := qrErrorCode
		if qErrorCode != "" {

			if err := r.SetQueryParam("error.code", qErrorCode); err != nil {
				return err
			}
		}
	}

	if o.ErrorMessage != nil {

		// query param error.message
		var qrErrorMessage string

		if o.ErrorMessage != nil {
			qrErrorMessage = *o.ErrorMessage
		}
		qErrorMessage := qrErrorMessage
		if qErrorMessage != "" {

			if err := r.SetQueryParam("error.message", qErrorMessage); err != nil {
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

	if o.Message != nil {

		// query param message
		var qrMessage string

		if o.Message != nil {
			qrMessage = *o.Message
		}
		qMessage := qrMessage
		if qMessage != "" {

			if err := r.SetQueryParam("message", qMessage); err != nil {
				return err
			}
		}
	}

	if o.NodeName != nil {

		// query param node.name
		var qrNodeName string

		if o.NodeName != nil {
			qrNodeName = *o.NodeName
		}
		qNodeName := qrNodeName
		if qNodeName != "" {

			if err := r.SetQueryParam("node.name", qNodeName); err != nil {
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

	if o.StartTime != nil {

		// query param start_time
		var qrStartTime string

		if o.StartTime != nil {
			qrStartTime = *o.StartTime
		}
		qStartTime := qrStartTime
		if qStartTime != "" {

			if err := r.SetQueryParam("start_time", qStartTime); err != nil {
				return err
			}
		}
	}

	if o.State != nil {

		// query param state
		var qrState string

		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
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

	if o.UUID != nil {

		// query param uuid
		var qrUUID string

		if o.UUID != nil {
			qrUUID = *o.UUID
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamJobCollectionGet binds the parameter fields
func (o *JobCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamJobCollectionGet binds the parameter order_by
func (o *JobCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
