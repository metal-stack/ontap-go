// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewVvolBindingCollectionGetParams creates a new VvolBindingCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVvolBindingCollectionGetParams() *VvolBindingCollectionGetParams {
	return &VvolBindingCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVvolBindingCollectionGetParamsWithTimeout creates a new VvolBindingCollectionGetParams object
// with the ability to set a timeout on a request.
func NewVvolBindingCollectionGetParamsWithTimeout(timeout time.Duration) *VvolBindingCollectionGetParams {
	return &VvolBindingCollectionGetParams{
		timeout: timeout,
	}
}

// NewVvolBindingCollectionGetParamsWithContext creates a new VvolBindingCollectionGetParams object
// with the ability to set a context for a request.
func NewVvolBindingCollectionGetParamsWithContext(ctx context.Context) *VvolBindingCollectionGetParams {
	return &VvolBindingCollectionGetParams{
		Context: ctx,
	}
}

// NewVvolBindingCollectionGetParamsWithHTTPClient creates a new VvolBindingCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewVvolBindingCollectionGetParamsWithHTTPClient(client *http.Client) *VvolBindingCollectionGetParams {
	return &VvolBindingCollectionGetParams{
		HTTPClient: client,
	}
}

/*
VvolBindingCollectionGetParams contains all the parameters to send to the API endpoint

	for the vvol binding collection get operation.

	Typically these are written to a http.Request.
*/
type VvolBindingCollectionGetParams struct {

	/* Count.

	   Filter by count
	*/
	Count *int64

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* ID.

	   Filter by id
	*/
	ID *int64

	/* IsOptimal.

	   Filter by is_optimal
	*/
	IsOptimal *bool

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* ProtocolEndpointName.

	   Filter by protocol_endpoint.name
	*/
	ProtocolEndpointName *string

	/* ProtocolEndpointUUID.

	   Filter by protocol_endpoint.uuid
	*/
	ProtocolEndpointUUID *string

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

	/* SecondaryID.

	   Filter by secondary_id
	*/
	SecondaryID *string

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	/* VvolName.

	   Filter by vvol.name
	*/
	VvolName *string

	/* VvolUUID.

	   Filter by vvol.uuid
	*/
	VvolUUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vvol binding collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VvolBindingCollectionGetParams) WithDefaults() *VvolBindingCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vvol binding collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VvolBindingCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := VvolBindingCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithTimeout(timeout time.Duration) *VvolBindingCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithContext(ctx context.Context) *VvolBindingCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithHTTPClient(client *http.Client) *VvolBindingCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithCount(count *int64) *VvolBindingCollectionGetParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetCount(count *int64) {
	o.Count = count
}

// WithFields adds the fields to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithFields(fields []string) *VvolBindingCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithID(id *int64) *VvolBindingCollectionGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetID(id *int64) {
	o.ID = id
}

// WithIsOptimal adds the isOptimal to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithIsOptimal(isOptimal *bool) *VvolBindingCollectionGetParams {
	o.SetIsOptimal(isOptimal)
	return o
}

// SetIsOptimal adds the isOptimal to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetIsOptimal(isOptimal *bool) {
	o.IsOptimal = isOptimal
}

// WithMaxRecords adds the maxRecords to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithMaxRecords(maxRecords *int64) *VvolBindingCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithOrderBy(orderBy []string) *VvolBindingCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithProtocolEndpointName adds the protocolEndpointName to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithProtocolEndpointName(protocolEndpointName *string) *VvolBindingCollectionGetParams {
	o.SetProtocolEndpointName(protocolEndpointName)
	return o
}

// SetProtocolEndpointName adds the protocolEndpointName to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetProtocolEndpointName(protocolEndpointName *string) {
	o.ProtocolEndpointName = protocolEndpointName
}

// WithProtocolEndpointUUID adds the protocolEndpointUUID to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithProtocolEndpointUUID(protocolEndpointUUID *string) *VvolBindingCollectionGetParams {
	o.SetProtocolEndpointUUID(protocolEndpointUUID)
	return o
}

// SetProtocolEndpointUUID adds the protocolEndpointUuid to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetProtocolEndpointUUID(protocolEndpointUUID *string) {
	o.ProtocolEndpointUUID = protocolEndpointUUID
}

// WithReturnRecords adds the returnRecords to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithReturnRecords(returnRecords *bool) *VvolBindingCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *VvolBindingCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSecondaryID adds the secondaryID to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithSecondaryID(secondaryID *string) *VvolBindingCollectionGetParams {
	o.SetSecondaryID(secondaryID)
	return o
}

// SetSecondaryID adds the secondaryId to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetSecondaryID(secondaryID *string) {
	o.SecondaryID = secondaryID
}

// WithSvmName adds the svmName to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithSvmName(svmName *string) *VvolBindingCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithSvmUUID(svmUUID *string) *VvolBindingCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithVvolName adds the vvolName to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithVvolName(vvolName *string) *VvolBindingCollectionGetParams {
	o.SetVvolName(vvolName)
	return o
}

// SetVvolName adds the vvolName to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetVvolName(vvolName *string) {
	o.VvolName = vvolName
}

// WithVvolUUID adds the vvolUUID to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithVvolUUID(vvolUUID *string) *VvolBindingCollectionGetParams {
	o.SetVvolUUID(vvolUUID)
	return o
}

// SetVvolUUID adds the vvolUuid to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetVvolUUID(vvolUUID *string) {
	o.VvolUUID = vvolUUID
}

// WriteToRequest writes these params to a swagger request
func (o *VvolBindingCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int64

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
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

	if o.ID != nil {

		// query param id
		var qrID int64

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt64(qrID)
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IsOptimal != nil {

		// query param is_optimal
		var qrIsOptimal bool

		if o.IsOptimal != nil {
			qrIsOptimal = *o.IsOptimal
		}
		qIsOptimal := swag.FormatBool(qrIsOptimal)
		if qIsOptimal != "" {

			if err := r.SetQueryParam("is_optimal", qIsOptimal); err != nil {
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

	if o.ProtocolEndpointName != nil {

		// query param protocol_endpoint.name
		var qrProtocolEndpointName string

		if o.ProtocolEndpointName != nil {
			qrProtocolEndpointName = *o.ProtocolEndpointName
		}
		qProtocolEndpointName := qrProtocolEndpointName
		if qProtocolEndpointName != "" {

			if err := r.SetQueryParam("protocol_endpoint.name", qProtocolEndpointName); err != nil {
				return err
			}
		}
	}

	if o.ProtocolEndpointUUID != nil {

		// query param protocol_endpoint.uuid
		var qrProtocolEndpointUUID string

		if o.ProtocolEndpointUUID != nil {
			qrProtocolEndpointUUID = *o.ProtocolEndpointUUID
		}
		qProtocolEndpointUUID := qrProtocolEndpointUUID
		if qProtocolEndpointUUID != "" {

			if err := r.SetQueryParam("protocol_endpoint.uuid", qProtocolEndpointUUID); err != nil {
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

	if o.SecondaryID != nil {

		// query param secondary_id
		var qrSecondaryID string

		if o.SecondaryID != nil {
			qrSecondaryID = *o.SecondaryID
		}
		qSecondaryID := qrSecondaryID
		if qSecondaryID != "" {

			if err := r.SetQueryParam("secondary_id", qSecondaryID); err != nil {
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

	if o.VvolName != nil {

		// query param vvol.name
		var qrVvolName string

		if o.VvolName != nil {
			qrVvolName = *o.VvolName
		}
		qVvolName := qrVvolName
		if qVvolName != "" {

			if err := r.SetQueryParam("vvol.name", qVvolName); err != nil {
				return err
			}
		}
	}

	if o.VvolUUID != nil {

		// query param vvol.uuid
		var qrVvolUUID string

		if o.VvolUUID != nil {
			qrVvolUUID = *o.VvolUUID
		}
		qVvolUUID := qrVvolUUID
		if qVvolUUID != "" {

			if err := r.SetQueryParam("vvol.uuid", qVvolUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamVvolBindingCollectionGet binds the parameter fields
func (o *VvolBindingCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamVvolBindingCollectionGet binds the parameter order_by
func (o *VvolBindingCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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