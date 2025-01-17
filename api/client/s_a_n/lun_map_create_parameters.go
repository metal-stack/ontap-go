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

	"github.com/metal-stack/ontap-go/api/models"
)

// NewLunMapCreateParams creates a new LunMapCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLunMapCreateParams() *LunMapCreateParams {
	return &LunMapCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLunMapCreateParamsWithTimeout creates a new LunMapCreateParams object
// with the ability to set a timeout on a request.
func NewLunMapCreateParamsWithTimeout(timeout time.Duration) *LunMapCreateParams {
	return &LunMapCreateParams{
		timeout: timeout,
	}
}

// NewLunMapCreateParamsWithContext creates a new LunMapCreateParams object
// with the ability to set a context for a request.
func NewLunMapCreateParamsWithContext(ctx context.Context) *LunMapCreateParams {
	return &LunMapCreateParams{
		Context: ctx,
	}
}

// NewLunMapCreateParamsWithHTTPClient creates a new LunMapCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewLunMapCreateParamsWithHTTPClient(client *http.Client) *LunMapCreateParams {
	return &LunMapCreateParams{
		HTTPClient: client,
	}
}

/*
LunMapCreateParams contains all the parameters to send to the API endpoint

	for the lun map create operation.

	Typically these are written to a http.Request.
*/
type LunMapCreateParams struct {

	/* AdditionalReportingNodeName.

	   The name of an ONTAP cluster node to add to the default reporting nodes for the LUN map. The HA partner for the node is also added.

	*/
	AdditionalReportingNodeName *string

	/* AdditionalReportingNodeUUID.

	   The unique identifier of an ONTAP cluster node to add to the default reporting nodes for the LUN map. The HA partner for the node is also added.

	*/
	AdditionalReportingNodeUUID *string

	/* Info.

	   The property values for the new LUN map.

	*/
	Info *models.LunMap

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lun map create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunMapCreateParams) WithDefaults() *LunMapCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lun map create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunMapCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := LunMapCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the lun map create params
func (o *LunMapCreateParams) WithTimeout(timeout time.Duration) *LunMapCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lun map create params
func (o *LunMapCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lun map create params
func (o *LunMapCreateParams) WithContext(ctx context.Context) *LunMapCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lun map create params
func (o *LunMapCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lun map create params
func (o *LunMapCreateParams) WithHTTPClient(client *http.Client) *LunMapCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lun map create params
func (o *LunMapCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdditionalReportingNodeName adds the additionalReportingNodeName to the lun map create params
func (o *LunMapCreateParams) WithAdditionalReportingNodeName(additionalReportingNodeName *string) *LunMapCreateParams {
	o.SetAdditionalReportingNodeName(additionalReportingNodeName)
	return o
}

// SetAdditionalReportingNodeName adds the additionalReportingNodeName to the lun map create params
func (o *LunMapCreateParams) SetAdditionalReportingNodeName(additionalReportingNodeName *string) {
	o.AdditionalReportingNodeName = additionalReportingNodeName
}

// WithAdditionalReportingNodeUUID adds the additionalReportingNodeUUID to the lun map create params
func (o *LunMapCreateParams) WithAdditionalReportingNodeUUID(additionalReportingNodeUUID *string) *LunMapCreateParams {
	o.SetAdditionalReportingNodeUUID(additionalReportingNodeUUID)
	return o
}

// SetAdditionalReportingNodeUUID adds the additionalReportingNodeUuid to the lun map create params
func (o *LunMapCreateParams) SetAdditionalReportingNodeUUID(additionalReportingNodeUUID *string) {
	o.AdditionalReportingNodeUUID = additionalReportingNodeUUID
}

// WithInfo adds the info to the lun map create params
func (o *LunMapCreateParams) WithInfo(info *models.LunMap) *LunMapCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the lun map create params
func (o *LunMapCreateParams) SetInfo(info *models.LunMap) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the lun map create params
func (o *LunMapCreateParams) WithReturnRecords(returnRecords *bool) *LunMapCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the lun map create params
func (o *LunMapCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *LunMapCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AdditionalReportingNodeName != nil {

		// query param additional_reporting_node.name
		var qrAdditionalReportingNodeName string

		if o.AdditionalReportingNodeName != nil {
			qrAdditionalReportingNodeName = *o.AdditionalReportingNodeName
		}
		qAdditionalReportingNodeName := qrAdditionalReportingNodeName
		if qAdditionalReportingNodeName != "" {

			if err := r.SetQueryParam("additional_reporting_node.name", qAdditionalReportingNodeName); err != nil {
				return err
			}
		}
	}

	if o.AdditionalReportingNodeUUID != nil {

		// query param additional_reporting_node.uuid
		var qrAdditionalReportingNodeUUID string

		if o.AdditionalReportingNodeUUID != nil {
			qrAdditionalReportingNodeUUID = *o.AdditionalReportingNodeUUID
		}
		qAdditionalReportingNodeUUID := qrAdditionalReportingNodeUUID
		if qAdditionalReportingNodeUUID != "" {

			if err := r.SetQueryParam("additional_reporting_node.uuid", qAdditionalReportingNodeUUID); err != nil {
				return err
			}
		}
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
