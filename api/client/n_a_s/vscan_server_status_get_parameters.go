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

// NewVscanServerStatusGetParams creates a new VscanServerStatusGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVscanServerStatusGetParams() *VscanServerStatusGetParams {
	return &VscanServerStatusGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVscanServerStatusGetParamsWithTimeout creates a new VscanServerStatusGetParams object
// with the ability to set a timeout on a request.
func NewVscanServerStatusGetParamsWithTimeout(timeout time.Duration) *VscanServerStatusGetParams {
	return &VscanServerStatusGetParams{
		timeout: timeout,
	}
}

// NewVscanServerStatusGetParamsWithContext creates a new VscanServerStatusGetParams object
// with the ability to set a context for a request.
func NewVscanServerStatusGetParamsWithContext(ctx context.Context) *VscanServerStatusGetParams {
	return &VscanServerStatusGetParams{
		Context: ctx,
	}
}

// NewVscanServerStatusGetParamsWithHTTPClient creates a new VscanServerStatusGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewVscanServerStatusGetParamsWithHTTPClient(client *http.Client) *VscanServerStatusGetParams {
	return &VscanServerStatusGetParams{
		HTTPClient: client,
	}
}

/*
VscanServerStatusGetParams contains all the parameters to send to the API endpoint

	for the vscan server status get operation.

	Typically these are written to a http.Request.
*/
type VscanServerStatusGetParams struct {

	/* DisconnectedReason.

	   Filter by disconnected_reason
	*/
	DisconnectedReason *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* InterfaceIPAddress.

	   Filter by interface.ip.address
	*/
	InterfaceIPAddress *string

	/* InterfaceName.

	   Filter by interface.name
	*/
	InterfaceName *string

	/* InterfaceUUID.

	   Filter by interface.uuid
	*/
	InterfaceUUID *string

	/* IP.

	   Filter by ip
	*/
	IP *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* NodeName.

	   Filter by node.name
	*/
	NodeName *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUID *string

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

	/* Type.

	   Filter by type
	*/
	Type *string

	/* UpdateTime.

	   Filter by update_time
	*/
	UpdateTime *string

	/* Vendor.

	   Filter by vendor
	*/
	Vendor *string

	/* Version.

	   Filter by version
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vscan server status get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanServerStatusGetParams) WithDefaults() *VscanServerStatusGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vscan server status get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanServerStatusGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := VscanServerStatusGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the vscan server status get params
func (o *VscanServerStatusGetParams) WithTimeout(timeout time.Duration) *VscanServerStatusGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vscan server status get params
func (o *VscanServerStatusGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vscan server status get params
func (o *VscanServerStatusGetParams) WithContext(ctx context.Context) *VscanServerStatusGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vscan server status get params
func (o *VscanServerStatusGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vscan server status get params
func (o *VscanServerStatusGetParams) WithHTTPClient(client *http.Client) *VscanServerStatusGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vscan server status get params
func (o *VscanServerStatusGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisconnectedReason adds the disconnectedReason to the vscan server status get params
func (o *VscanServerStatusGetParams) WithDisconnectedReason(disconnectedReason *string) *VscanServerStatusGetParams {
	o.SetDisconnectedReason(disconnectedReason)
	return o
}

// SetDisconnectedReason adds the disconnectedReason to the vscan server status get params
func (o *VscanServerStatusGetParams) SetDisconnectedReason(disconnectedReason *string) {
	o.DisconnectedReason = disconnectedReason
}

// WithFields adds the fields to the vscan server status get params
func (o *VscanServerStatusGetParams) WithFields(fields []string) *VscanServerStatusGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the vscan server status get params
func (o *VscanServerStatusGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithInterfaceIPAddress adds the interfaceIPAddress to the vscan server status get params
func (o *VscanServerStatusGetParams) WithInterfaceIPAddress(interfaceIPAddress *string) *VscanServerStatusGetParams {
	o.SetInterfaceIPAddress(interfaceIPAddress)
	return o
}

// SetInterfaceIPAddress adds the interfaceIpAddress to the vscan server status get params
func (o *VscanServerStatusGetParams) SetInterfaceIPAddress(interfaceIPAddress *string) {
	o.InterfaceIPAddress = interfaceIPAddress
}

// WithInterfaceName adds the interfaceName to the vscan server status get params
func (o *VscanServerStatusGetParams) WithInterfaceName(interfaceName *string) *VscanServerStatusGetParams {
	o.SetInterfaceName(interfaceName)
	return o
}

// SetInterfaceName adds the interfaceName to the vscan server status get params
func (o *VscanServerStatusGetParams) SetInterfaceName(interfaceName *string) {
	o.InterfaceName = interfaceName
}

// WithInterfaceUUID adds the interfaceUUID to the vscan server status get params
func (o *VscanServerStatusGetParams) WithInterfaceUUID(interfaceUUID *string) *VscanServerStatusGetParams {
	o.SetInterfaceUUID(interfaceUUID)
	return o
}

// SetInterfaceUUID adds the interfaceUuid to the vscan server status get params
func (o *VscanServerStatusGetParams) SetInterfaceUUID(interfaceUUID *string) {
	o.InterfaceUUID = interfaceUUID
}

// WithIP adds the ip to the vscan server status get params
func (o *VscanServerStatusGetParams) WithIP(ip *string) *VscanServerStatusGetParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the vscan server status get params
func (o *VscanServerStatusGetParams) SetIP(ip *string) {
	o.IP = ip
}

// WithMaxRecords adds the maxRecords to the vscan server status get params
func (o *VscanServerStatusGetParams) WithMaxRecords(maxRecords *int64) *VscanServerStatusGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the vscan server status get params
func (o *VscanServerStatusGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNodeName adds the nodeName to the vscan server status get params
func (o *VscanServerStatusGetParams) WithNodeName(nodeName *string) *VscanServerStatusGetParams {
	o.SetNodeName(nodeName)
	return o
}

// SetNodeName adds the nodeName to the vscan server status get params
func (o *VscanServerStatusGetParams) SetNodeName(nodeName *string) {
	o.NodeName = nodeName
}

// WithNodeUUID adds the nodeUUID to the vscan server status get params
func (o *VscanServerStatusGetParams) WithNodeUUID(nodeUUID *string) *VscanServerStatusGetParams {
	o.SetNodeUUID(nodeUUID)
	return o
}

// SetNodeUUID adds the nodeUuid to the vscan server status get params
func (o *VscanServerStatusGetParams) SetNodeUUID(nodeUUID *string) {
	o.NodeUUID = nodeUUID
}

// WithOrderBy adds the orderBy to the vscan server status get params
func (o *VscanServerStatusGetParams) WithOrderBy(orderBy []string) *VscanServerStatusGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the vscan server status get params
func (o *VscanServerStatusGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the vscan server status get params
func (o *VscanServerStatusGetParams) WithReturnRecords(returnRecords *bool) *VscanServerStatusGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the vscan server status get params
func (o *VscanServerStatusGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the vscan server status get params
func (o *VscanServerStatusGetParams) WithReturnTimeout(returnTimeout *int64) *VscanServerStatusGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the vscan server status get params
func (o *VscanServerStatusGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithState adds the state to the vscan server status get params
func (o *VscanServerStatusGetParams) WithState(state *string) *VscanServerStatusGetParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the vscan server status get params
func (o *VscanServerStatusGetParams) SetState(state *string) {
	o.State = state
}

// WithSvmName adds the svmName to the vscan server status get params
func (o *VscanServerStatusGetParams) WithSvmName(svmName *string) *VscanServerStatusGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the vscan server status get params
func (o *VscanServerStatusGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the vscan server status get params
func (o *VscanServerStatusGetParams) WithSvmUUID(svmUUID *string) *VscanServerStatusGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the vscan server status get params
func (o *VscanServerStatusGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithType adds the typeVar to the vscan server status get params
func (o *VscanServerStatusGetParams) WithType(typeVar *string) *VscanServerStatusGetParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the vscan server status get params
func (o *VscanServerStatusGetParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithUpdateTime adds the updateTime to the vscan server status get params
func (o *VscanServerStatusGetParams) WithUpdateTime(updateTime *string) *VscanServerStatusGetParams {
	o.SetUpdateTime(updateTime)
	return o
}

// SetUpdateTime adds the updateTime to the vscan server status get params
func (o *VscanServerStatusGetParams) SetUpdateTime(updateTime *string) {
	o.UpdateTime = updateTime
}

// WithVendor adds the vendor to the vscan server status get params
func (o *VscanServerStatusGetParams) WithVendor(vendor *string) *VscanServerStatusGetParams {
	o.SetVendor(vendor)
	return o
}

// SetVendor adds the vendor to the vscan server status get params
func (o *VscanServerStatusGetParams) SetVendor(vendor *string) {
	o.Vendor = vendor
}

// WithVersion adds the version to the vscan server status get params
func (o *VscanServerStatusGetParams) WithVersion(version *string) *VscanServerStatusGetParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the vscan server status get params
func (o *VscanServerStatusGetParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *VscanServerStatusGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DisconnectedReason != nil {

		// query param disconnected_reason
		var qrDisconnectedReason string

		if o.DisconnectedReason != nil {
			qrDisconnectedReason = *o.DisconnectedReason
		}
		qDisconnectedReason := qrDisconnectedReason
		if qDisconnectedReason != "" {

			if err := r.SetQueryParam("disconnected_reason", qDisconnectedReason); err != nil {
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

	if o.InterfaceIPAddress != nil {

		// query param interface.ip.address
		var qrInterfaceIPAddress string

		if o.InterfaceIPAddress != nil {
			qrInterfaceIPAddress = *o.InterfaceIPAddress
		}
		qInterfaceIPAddress := qrInterfaceIPAddress
		if qInterfaceIPAddress != "" {

			if err := r.SetQueryParam("interface.ip.address", qInterfaceIPAddress); err != nil {
				return err
			}
		}
	}

	if o.InterfaceName != nil {

		// query param interface.name
		var qrInterfaceName string

		if o.InterfaceName != nil {
			qrInterfaceName = *o.InterfaceName
		}
		qInterfaceName := qrInterfaceName
		if qInterfaceName != "" {

			if err := r.SetQueryParam("interface.name", qInterfaceName); err != nil {
				return err
			}
		}
	}

	if o.InterfaceUUID != nil {

		// query param interface.uuid
		var qrInterfaceUUID string

		if o.InterfaceUUID != nil {
			qrInterfaceUUID = *o.InterfaceUUID
		}
		qInterfaceUUID := qrInterfaceUUID
		if qInterfaceUUID != "" {

			if err := r.SetQueryParam("interface.uuid", qInterfaceUUID); err != nil {
				return err
			}
		}
	}

	if o.IP != nil {

		// query param ip
		var qrIP string

		if o.IP != nil {
			qrIP = *o.IP
		}
		qIP := qrIP
		if qIP != "" {

			if err := r.SetQueryParam("ip", qIP); err != nil {
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

	if o.NodeUUID != nil {

		// query param node.uuid
		var qrNodeUUID string

		if o.NodeUUID != nil {
			qrNodeUUID = *o.NodeUUID
		}
		qNodeUUID := qrNodeUUID
		if qNodeUUID != "" {

			if err := r.SetQueryParam("node.uuid", qNodeUUID); err != nil {
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

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.UpdateTime != nil {

		// query param update_time
		var qrUpdateTime string

		if o.UpdateTime != nil {
			qrUpdateTime = *o.UpdateTime
		}
		qUpdateTime := qrUpdateTime
		if qUpdateTime != "" {

			if err := r.SetQueryParam("update_time", qUpdateTime); err != nil {
				return err
			}
		}
	}

	if o.Vendor != nil {

		// query param vendor
		var qrVendor string

		if o.Vendor != nil {
			qrVendor = *o.Vendor
		}
		qVendor := qrVendor
		if qVendor != "" {

			if err := r.SetQueryParam("vendor", qVendor); err != nil {
				return err
			}
		}
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamVscanServerStatusGet binds the parameter fields
func (o *VscanServerStatusGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamVscanServerStatusGet binds the parameter order_by
func (o *VscanServerStatusGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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