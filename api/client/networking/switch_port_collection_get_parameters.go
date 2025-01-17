// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// NewSwitchPortCollectionGetParams creates a new SwitchPortCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSwitchPortCollectionGetParams() *SwitchPortCollectionGetParams {
	return &SwitchPortCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSwitchPortCollectionGetParamsWithTimeout creates a new SwitchPortCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSwitchPortCollectionGetParamsWithTimeout(timeout time.Duration) *SwitchPortCollectionGetParams {
	return &SwitchPortCollectionGetParams{
		timeout: timeout,
	}
}

// NewSwitchPortCollectionGetParamsWithContext creates a new SwitchPortCollectionGetParams object
// with the ability to set a context for a request.
func NewSwitchPortCollectionGetParamsWithContext(ctx context.Context) *SwitchPortCollectionGetParams {
	return &SwitchPortCollectionGetParams{
		Context: ctx,
	}
}

// NewSwitchPortCollectionGetParamsWithHTTPClient creates a new SwitchPortCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSwitchPortCollectionGetParamsWithHTTPClient(client *http.Client) *SwitchPortCollectionGetParams {
	return &SwitchPortCollectionGetParams{
		HTTPClient: client,
	}
}

/*
SwitchPortCollectionGetParams contains all the parameters to send to the API endpoint

	for the switch port collection get operation.

	Typically these are written to a http.Request.
*/
type SwitchPortCollectionGetParams struct {

	/* Configured.

	   Filter by configured
	*/
	Configured *string

	/* DuplexType.

	   Filter by duplex_type
	*/
	DuplexType *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* IdentityIndex.

	   Filter by identity.index
	*/
	IdentityIndex *int64

	/* IdentityName.

	   Filter by identity.name
	*/
	IdentityName *string

	/* IdentityNumber.

	   Filter by identity.number
	*/
	IdentityNumber *int64

	/* Isl.

	   Filter by isl
	*/
	Isl *bool

	/* MacAddress.

	   Filter by mac_address
	*/
	MacAddress *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Mtu.

	   Filter by mtu
	*/
	Mtu *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* RemotePortDeviceNodeName.

	   Filter by remote_port.device.node.name
	*/
	RemotePortDeviceNodeName *string

	/* RemotePortDeviceNodeUUID.

	   Filter by remote_port.device.node.uuid
	*/
	RemotePortDeviceNodeUUID *string

	/* RemotePortDeviceShelfModule.

	   Filter by remote_port.device.shelf.module
	*/
	RemotePortDeviceShelfModule *string

	/* RemotePortDeviceShelfName.

	   Filter by remote_port.device.shelf.name
	*/
	RemotePortDeviceShelfName *string

	/* RemotePortDeviceShelfUID.

	   Filter by remote_port.device.shelf.uid
	*/
	RemotePortDeviceShelfUID *string

	/* RemotePortMtu.

	   Filter by remote_port.mtu
	*/
	RemotePortMtu *int64

	/* RemotePortName.

	   Filter by remote_port.name
	*/
	RemotePortName *string

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

	/* Speed.

	   Filter by speed
	*/
	Speed *int64

	/* State.

	   Filter by state
	*/
	State *string

	/* StatisticsReceiveRawDiscards.

	   Filter by statistics.receive_raw.discards
	*/
	StatisticsReceiveRawDiscards *int64

	/* StatisticsReceiveRawErrors.

	   Filter by statistics.receive_raw.errors
	*/
	StatisticsReceiveRawErrors *int64

	/* StatisticsReceiveRawPackets.

	   Filter by statistics.receive_raw.packets
	*/
	StatisticsReceiveRawPackets *int64

	/* StatisticsTransmitRawDiscards.

	   Filter by statistics.transmit_raw.discards
	*/
	StatisticsTransmitRawDiscards *int64

	/* StatisticsTransmitRawErrors.

	   Filter by statistics.transmit_raw.errors
	*/
	StatisticsTransmitRawErrors *int64

	/* StatisticsTransmitRawPackets.

	   Filter by statistics.transmit_raw.packets
	*/
	StatisticsTransmitRawPackets *int64

	/* SwitchName.

	   Filter by switch.name
	*/
	SwitchName *string

	/* Type.

	   Filter by type
	*/
	Type *string

	/* VlanID.

	   Filter by vlan_id
	*/
	VlanID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the switch port collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SwitchPortCollectionGetParams) WithDefaults() *SwitchPortCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the switch port collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SwitchPortCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SwitchPortCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithTimeout(timeout time.Duration) *SwitchPortCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithContext(ctx context.Context) *SwitchPortCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithHTTPClient(client *http.Client) *SwitchPortCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigured adds the configured to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithConfigured(configured *string) *SwitchPortCollectionGetParams {
	o.SetConfigured(configured)
	return o
}

// SetConfigured adds the configured to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetConfigured(configured *string) {
	o.Configured = configured
}

// WithDuplexType adds the duplexType to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithDuplexType(duplexType *string) *SwitchPortCollectionGetParams {
	o.SetDuplexType(duplexType)
	return o
}

// SetDuplexType adds the duplexType to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetDuplexType(duplexType *string) {
	o.DuplexType = duplexType
}

// WithFields adds the fields to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithFields(fields []string) *SwitchPortCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIdentityIndex adds the identityIndex to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithIdentityIndex(identityIndex *int64) *SwitchPortCollectionGetParams {
	o.SetIdentityIndex(identityIndex)
	return o
}

// SetIdentityIndex adds the identityIndex to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetIdentityIndex(identityIndex *int64) {
	o.IdentityIndex = identityIndex
}

// WithIdentityName adds the identityName to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithIdentityName(identityName *string) *SwitchPortCollectionGetParams {
	o.SetIdentityName(identityName)
	return o
}

// SetIdentityName adds the identityName to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetIdentityName(identityName *string) {
	o.IdentityName = identityName
}

// WithIdentityNumber adds the identityNumber to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithIdentityNumber(identityNumber *int64) *SwitchPortCollectionGetParams {
	o.SetIdentityNumber(identityNumber)
	return o
}

// SetIdentityNumber adds the identityNumber to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetIdentityNumber(identityNumber *int64) {
	o.IdentityNumber = identityNumber
}

// WithIsl adds the isl to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithIsl(isl *bool) *SwitchPortCollectionGetParams {
	o.SetIsl(isl)
	return o
}

// SetIsl adds the isl to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetIsl(isl *bool) {
	o.Isl = isl
}

// WithMacAddress adds the macAddress to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithMacAddress(macAddress *string) *SwitchPortCollectionGetParams {
	o.SetMacAddress(macAddress)
	return o
}

// SetMacAddress adds the macAddress to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetMacAddress(macAddress *string) {
	o.MacAddress = macAddress
}

// WithMaxRecords adds the maxRecords to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithMaxRecords(maxRecords *int64) *SwitchPortCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithMtu adds the mtu to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithMtu(mtu *int64) *SwitchPortCollectionGetParams {
	o.SetMtu(mtu)
	return o
}

// SetMtu adds the mtu to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetMtu(mtu *int64) {
	o.Mtu = mtu
}

// WithOrderBy adds the orderBy to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithOrderBy(orderBy []string) *SwitchPortCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithRemotePortDeviceNodeName adds the remotePortDeviceNodeName to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithRemotePortDeviceNodeName(remotePortDeviceNodeName *string) *SwitchPortCollectionGetParams {
	o.SetRemotePortDeviceNodeName(remotePortDeviceNodeName)
	return o
}

// SetRemotePortDeviceNodeName adds the remotePortDeviceNodeName to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetRemotePortDeviceNodeName(remotePortDeviceNodeName *string) {
	o.RemotePortDeviceNodeName = remotePortDeviceNodeName
}

// WithRemotePortDeviceNodeUUID adds the remotePortDeviceNodeUUID to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithRemotePortDeviceNodeUUID(remotePortDeviceNodeUUID *string) *SwitchPortCollectionGetParams {
	o.SetRemotePortDeviceNodeUUID(remotePortDeviceNodeUUID)
	return o
}

// SetRemotePortDeviceNodeUUID adds the remotePortDeviceNodeUuid to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetRemotePortDeviceNodeUUID(remotePortDeviceNodeUUID *string) {
	o.RemotePortDeviceNodeUUID = remotePortDeviceNodeUUID
}

// WithRemotePortDeviceShelfModule adds the remotePortDeviceShelfModule to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithRemotePortDeviceShelfModule(remotePortDeviceShelfModule *string) *SwitchPortCollectionGetParams {
	o.SetRemotePortDeviceShelfModule(remotePortDeviceShelfModule)
	return o
}

// SetRemotePortDeviceShelfModule adds the remotePortDeviceShelfModule to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetRemotePortDeviceShelfModule(remotePortDeviceShelfModule *string) {
	o.RemotePortDeviceShelfModule = remotePortDeviceShelfModule
}

// WithRemotePortDeviceShelfName adds the remotePortDeviceShelfName to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithRemotePortDeviceShelfName(remotePortDeviceShelfName *string) *SwitchPortCollectionGetParams {
	o.SetRemotePortDeviceShelfName(remotePortDeviceShelfName)
	return o
}

// SetRemotePortDeviceShelfName adds the remotePortDeviceShelfName to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetRemotePortDeviceShelfName(remotePortDeviceShelfName *string) {
	o.RemotePortDeviceShelfName = remotePortDeviceShelfName
}

// WithRemotePortDeviceShelfUID adds the remotePortDeviceShelfUID to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithRemotePortDeviceShelfUID(remotePortDeviceShelfUID *string) *SwitchPortCollectionGetParams {
	o.SetRemotePortDeviceShelfUID(remotePortDeviceShelfUID)
	return o
}

// SetRemotePortDeviceShelfUID adds the remotePortDeviceShelfUid to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetRemotePortDeviceShelfUID(remotePortDeviceShelfUID *string) {
	o.RemotePortDeviceShelfUID = remotePortDeviceShelfUID
}

// WithRemotePortMtu adds the remotePortMtu to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithRemotePortMtu(remotePortMtu *int64) *SwitchPortCollectionGetParams {
	o.SetRemotePortMtu(remotePortMtu)
	return o
}

// SetRemotePortMtu adds the remotePortMtu to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetRemotePortMtu(remotePortMtu *int64) {
	o.RemotePortMtu = remotePortMtu
}

// WithRemotePortName adds the remotePortName to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithRemotePortName(remotePortName *string) *SwitchPortCollectionGetParams {
	o.SetRemotePortName(remotePortName)
	return o
}

// SetRemotePortName adds the remotePortName to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetRemotePortName(remotePortName *string) {
	o.RemotePortName = remotePortName
}

// WithReturnRecords adds the returnRecords to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithReturnRecords(returnRecords *bool) *SwitchPortCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *SwitchPortCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSpeed adds the speed to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithSpeed(speed *int64) *SwitchPortCollectionGetParams {
	o.SetSpeed(speed)
	return o
}

// SetSpeed adds the speed to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetSpeed(speed *int64) {
	o.Speed = speed
}

// WithState adds the state to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithState(state *string) *SwitchPortCollectionGetParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetState(state *string) {
	o.State = state
}

// WithStatisticsReceiveRawDiscards adds the statisticsReceiveRawDiscards to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithStatisticsReceiveRawDiscards(statisticsReceiveRawDiscards *int64) *SwitchPortCollectionGetParams {
	o.SetStatisticsReceiveRawDiscards(statisticsReceiveRawDiscards)
	return o
}

// SetStatisticsReceiveRawDiscards adds the statisticsReceiveRawDiscards to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetStatisticsReceiveRawDiscards(statisticsReceiveRawDiscards *int64) {
	o.StatisticsReceiveRawDiscards = statisticsReceiveRawDiscards
}

// WithStatisticsReceiveRawErrors adds the statisticsReceiveRawErrors to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithStatisticsReceiveRawErrors(statisticsReceiveRawErrors *int64) *SwitchPortCollectionGetParams {
	o.SetStatisticsReceiveRawErrors(statisticsReceiveRawErrors)
	return o
}

// SetStatisticsReceiveRawErrors adds the statisticsReceiveRawErrors to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetStatisticsReceiveRawErrors(statisticsReceiveRawErrors *int64) {
	o.StatisticsReceiveRawErrors = statisticsReceiveRawErrors
}

// WithStatisticsReceiveRawPackets adds the statisticsReceiveRawPackets to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithStatisticsReceiveRawPackets(statisticsReceiveRawPackets *int64) *SwitchPortCollectionGetParams {
	o.SetStatisticsReceiveRawPackets(statisticsReceiveRawPackets)
	return o
}

// SetStatisticsReceiveRawPackets adds the statisticsReceiveRawPackets to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetStatisticsReceiveRawPackets(statisticsReceiveRawPackets *int64) {
	o.StatisticsReceiveRawPackets = statisticsReceiveRawPackets
}

// WithStatisticsTransmitRawDiscards adds the statisticsTransmitRawDiscards to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithStatisticsTransmitRawDiscards(statisticsTransmitRawDiscards *int64) *SwitchPortCollectionGetParams {
	o.SetStatisticsTransmitRawDiscards(statisticsTransmitRawDiscards)
	return o
}

// SetStatisticsTransmitRawDiscards adds the statisticsTransmitRawDiscards to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetStatisticsTransmitRawDiscards(statisticsTransmitRawDiscards *int64) {
	o.StatisticsTransmitRawDiscards = statisticsTransmitRawDiscards
}

// WithStatisticsTransmitRawErrors adds the statisticsTransmitRawErrors to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithStatisticsTransmitRawErrors(statisticsTransmitRawErrors *int64) *SwitchPortCollectionGetParams {
	o.SetStatisticsTransmitRawErrors(statisticsTransmitRawErrors)
	return o
}

// SetStatisticsTransmitRawErrors adds the statisticsTransmitRawErrors to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetStatisticsTransmitRawErrors(statisticsTransmitRawErrors *int64) {
	o.StatisticsTransmitRawErrors = statisticsTransmitRawErrors
}

// WithStatisticsTransmitRawPackets adds the statisticsTransmitRawPackets to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithStatisticsTransmitRawPackets(statisticsTransmitRawPackets *int64) *SwitchPortCollectionGetParams {
	o.SetStatisticsTransmitRawPackets(statisticsTransmitRawPackets)
	return o
}

// SetStatisticsTransmitRawPackets adds the statisticsTransmitRawPackets to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetStatisticsTransmitRawPackets(statisticsTransmitRawPackets *int64) {
	o.StatisticsTransmitRawPackets = statisticsTransmitRawPackets
}

// WithSwitchName adds the switchName to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithSwitchName(switchName *string) *SwitchPortCollectionGetParams {
	o.SetSwitchName(switchName)
	return o
}

// SetSwitchName adds the switchName to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetSwitchName(switchName *string) {
	o.SwitchName = switchName
}

// WithType adds the typeVar to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithType(typeVar *string) *SwitchPortCollectionGetParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithVlanID adds the vlanID to the switch port collection get params
func (o *SwitchPortCollectionGetParams) WithVlanID(vlanID *int64) *SwitchPortCollectionGetParams {
	o.SetVlanID(vlanID)
	return o
}

// SetVlanID adds the vlanId to the switch port collection get params
func (o *SwitchPortCollectionGetParams) SetVlanID(vlanID *int64) {
	o.VlanID = vlanID
}

// WriteToRequest writes these params to a swagger request
func (o *SwitchPortCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Configured != nil {

		// query param configured
		var qrConfigured string

		if o.Configured != nil {
			qrConfigured = *o.Configured
		}
		qConfigured := qrConfigured
		if qConfigured != "" {

			if err := r.SetQueryParam("configured", qConfigured); err != nil {
				return err
			}
		}
	}

	if o.DuplexType != nil {

		// query param duplex_type
		var qrDuplexType string

		if o.DuplexType != nil {
			qrDuplexType = *o.DuplexType
		}
		qDuplexType := qrDuplexType
		if qDuplexType != "" {

			if err := r.SetQueryParam("duplex_type", qDuplexType); err != nil {
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

	if o.IdentityIndex != nil {

		// query param identity.index
		var qrIdentityIndex int64

		if o.IdentityIndex != nil {
			qrIdentityIndex = *o.IdentityIndex
		}
		qIdentityIndex := swag.FormatInt64(qrIdentityIndex)
		if qIdentityIndex != "" {

			if err := r.SetQueryParam("identity.index", qIdentityIndex); err != nil {
				return err
			}
		}
	}

	if o.IdentityName != nil {

		// query param identity.name
		var qrIdentityName string

		if o.IdentityName != nil {
			qrIdentityName = *o.IdentityName
		}
		qIdentityName := qrIdentityName
		if qIdentityName != "" {

			if err := r.SetQueryParam("identity.name", qIdentityName); err != nil {
				return err
			}
		}
	}

	if o.IdentityNumber != nil {

		// query param identity.number
		var qrIdentityNumber int64

		if o.IdentityNumber != nil {
			qrIdentityNumber = *o.IdentityNumber
		}
		qIdentityNumber := swag.FormatInt64(qrIdentityNumber)
		if qIdentityNumber != "" {

			if err := r.SetQueryParam("identity.number", qIdentityNumber); err != nil {
				return err
			}
		}
	}

	if o.Isl != nil {

		// query param isl
		var qrIsl bool

		if o.Isl != nil {
			qrIsl = *o.Isl
		}
		qIsl := swag.FormatBool(qrIsl)
		if qIsl != "" {

			if err := r.SetQueryParam("isl", qIsl); err != nil {
				return err
			}
		}
	}

	if o.MacAddress != nil {

		// query param mac_address
		var qrMacAddress string

		if o.MacAddress != nil {
			qrMacAddress = *o.MacAddress
		}
		qMacAddress := qrMacAddress
		if qMacAddress != "" {

			if err := r.SetQueryParam("mac_address", qMacAddress); err != nil {
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

	if o.Mtu != nil {

		// query param mtu
		var qrMtu int64

		if o.Mtu != nil {
			qrMtu = *o.Mtu
		}
		qMtu := swag.FormatInt64(qrMtu)
		if qMtu != "" {

			if err := r.SetQueryParam("mtu", qMtu); err != nil {
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

	if o.RemotePortDeviceNodeName != nil {

		// query param remote_port.device.node.name
		var qrRemotePortDeviceNodeName string

		if o.RemotePortDeviceNodeName != nil {
			qrRemotePortDeviceNodeName = *o.RemotePortDeviceNodeName
		}
		qRemotePortDeviceNodeName := qrRemotePortDeviceNodeName
		if qRemotePortDeviceNodeName != "" {

			if err := r.SetQueryParam("remote_port.device.node.name", qRemotePortDeviceNodeName); err != nil {
				return err
			}
		}
	}

	if o.RemotePortDeviceNodeUUID != nil {

		// query param remote_port.device.node.uuid
		var qrRemotePortDeviceNodeUUID string

		if o.RemotePortDeviceNodeUUID != nil {
			qrRemotePortDeviceNodeUUID = *o.RemotePortDeviceNodeUUID
		}
		qRemotePortDeviceNodeUUID := qrRemotePortDeviceNodeUUID
		if qRemotePortDeviceNodeUUID != "" {

			if err := r.SetQueryParam("remote_port.device.node.uuid", qRemotePortDeviceNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.RemotePortDeviceShelfModule != nil {

		// query param remote_port.device.shelf.module
		var qrRemotePortDeviceShelfModule string

		if o.RemotePortDeviceShelfModule != nil {
			qrRemotePortDeviceShelfModule = *o.RemotePortDeviceShelfModule
		}
		qRemotePortDeviceShelfModule := qrRemotePortDeviceShelfModule
		if qRemotePortDeviceShelfModule != "" {

			if err := r.SetQueryParam("remote_port.device.shelf.module", qRemotePortDeviceShelfModule); err != nil {
				return err
			}
		}
	}

	if o.RemotePortDeviceShelfName != nil {

		// query param remote_port.device.shelf.name
		var qrRemotePortDeviceShelfName string

		if o.RemotePortDeviceShelfName != nil {
			qrRemotePortDeviceShelfName = *o.RemotePortDeviceShelfName
		}
		qRemotePortDeviceShelfName := qrRemotePortDeviceShelfName
		if qRemotePortDeviceShelfName != "" {

			if err := r.SetQueryParam("remote_port.device.shelf.name", qRemotePortDeviceShelfName); err != nil {
				return err
			}
		}
	}

	if o.RemotePortDeviceShelfUID != nil {

		// query param remote_port.device.shelf.uid
		var qrRemotePortDeviceShelfUID string

		if o.RemotePortDeviceShelfUID != nil {
			qrRemotePortDeviceShelfUID = *o.RemotePortDeviceShelfUID
		}
		qRemotePortDeviceShelfUID := qrRemotePortDeviceShelfUID
		if qRemotePortDeviceShelfUID != "" {

			if err := r.SetQueryParam("remote_port.device.shelf.uid", qRemotePortDeviceShelfUID); err != nil {
				return err
			}
		}
	}

	if o.RemotePortMtu != nil {

		// query param remote_port.mtu
		var qrRemotePortMtu int64

		if o.RemotePortMtu != nil {
			qrRemotePortMtu = *o.RemotePortMtu
		}
		qRemotePortMtu := swag.FormatInt64(qrRemotePortMtu)
		if qRemotePortMtu != "" {

			if err := r.SetQueryParam("remote_port.mtu", qRemotePortMtu); err != nil {
				return err
			}
		}
	}

	if o.RemotePortName != nil {

		// query param remote_port.name
		var qrRemotePortName string

		if o.RemotePortName != nil {
			qrRemotePortName = *o.RemotePortName
		}
		qRemotePortName := qrRemotePortName
		if qRemotePortName != "" {

			if err := r.SetQueryParam("remote_port.name", qRemotePortName); err != nil {
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

	if o.Speed != nil {

		// query param speed
		var qrSpeed int64

		if o.Speed != nil {
			qrSpeed = *o.Speed
		}
		qSpeed := swag.FormatInt64(qrSpeed)
		if qSpeed != "" {

			if err := r.SetQueryParam("speed", qSpeed); err != nil {
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

	if o.StatisticsReceiveRawDiscards != nil {

		// query param statistics.receive_raw.discards
		var qrStatisticsReceiveRawDiscards int64

		if o.StatisticsReceiveRawDiscards != nil {
			qrStatisticsReceiveRawDiscards = *o.StatisticsReceiveRawDiscards
		}
		qStatisticsReceiveRawDiscards := swag.FormatInt64(qrStatisticsReceiveRawDiscards)
		if qStatisticsReceiveRawDiscards != "" {

			if err := r.SetQueryParam("statistics.receive_raw.discards", qStatisticsReceiveRawDiscards); err != nil {
				return err
			}
		}
	}

	if o.StatisticsReceiveRawErrors != nil {

		// query param statistics.receive_raw.errors
		var qrStatisticsReceiveRawErrors int64

		if o.StatisticsReceiveRawErrors != nil {
			qrStatisticsReceiveRawErrors = *o.StatisticsReceiveRawErrors
		}
		qStatisticsReceiveRawErrors := swag.FormatInt64(qrStatisticsReceiveRawErrors)
		if qStatisticsReceiveRawErrors != "" {

			if err := r.SetQueryParam("statistics.receive_raw.errors", qStatisticsReceiveRawErrors); err != nil {
				return err
			}
		}
	}

	if o.StatisticsReceiveRawPackets != nil {

		// query param statistics.receive_raw.packets
		var qrStatisticsReceiveRawPackets int64

		if o.StatisticsReceiveRawPackets != nil {
			qrStatisticsReceiveRawPackets = *o.StatisticsReceiveRawPackets
		}
		qStatisticsReceiveRawPackets := swag.FormatInt64(qrStatisticsReceiveRawPackets)
		if qStatisticsReceiveRawPackets != "" {

			if err := r.SetQueryParam("statistics.receive_raw.packets", qStatisticsReceiveRawPackets); err != nil {
				return err
			}
		}
	}

	if o.StatisticsTransmitRawDiscards != nil {

		// query param statistics.transmit_raw.discards
		var qrStatisticsTransmitRawDiscards int64

		if o.StatisticsTransmitRawDiscards != nil {
			qrStatisticsTransmitRawDiscards = *o.StatisticsTransmitRawDiscards
		}
		qStatisticsTransmitRawDiscards := swag.FormatInt64(qrStatisticsTransmitRawDiscards)
		if qStatisticsTransmitRawDiscards != "" {

			if err := r.SetQueryParam("statistics.transmit_raw.discards", qStatisticsTransmitRawDiscards); err != nil {
				return err
			}
		}
	}

	if o.StatisticsTransmitRawErrors != nil {

		// query param statistics.transmit_raw.errors
		var qrStatisticsTransmitRawErrors int64

		if o.StatisticsTransmitRawErrors != nil {
			qrStatisticsTransmitRawErrors = *o.StatisticsTransmitRawErrors
		}
		qStatisticsTransmitRawErrors := swag.FormatInt64(qrStatisticsTransmitRawErrors)
		if qStatisticsTransmitRawErrors != "" {

			if err := r.SetQueryParam("statistics.transmit_raw.errors", qStatisticsTransmitRawErrors); err != nil {
				return err
			}
		}
	}

	if o.StatisticsTransmitRawPackets != nil {

		// query param statistics.transmit_raw.packets
		var qrStatisticsTransmitRawPackets int64

		if o.StatisticsTransmitRawPackets != nil {
			qrStatisticsTransmitRawPackets = *o.StatisticsTransmitRawPackets
		}
		qStatisticsTransmitRawPackets := swag.FormatInt64(qrStatisticsTransmitRawPackets)
		if qStatisticsTransmitRawPackets != "" {

			if err := r.SetQueryParam("statistics.transmit_raw.packets", qStatisticsTransmitRawPackets); err != nil {
				return err
			}
		}
	}

	if o.SwitchName != nil {

		// query param switch.name
		var qrSwitchName string

		if o.SwitchName != nil {
			qrSwitchName = *o.SwitchName
		}
		qSwitchName := qrSwitchName
		if qSwitchName != "" {

			if err := r.SetQueryParam("switch.name", qSwitchName); err != nil {
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

	if o.VlanID != nil {

		// query param vlan_id
		var qrVlanID int64

		if o.VlanID != nil {
			qrVlanID = *o.VlanID
		}
		qVlanID := swag.FormatInt64(qrVlanID)
		if qVlanID != "" {

			if err := r.SetQueryParam("vlan_id", qVlanID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSwitchPortCollectionGet binds the parameter fields
func (o *SwitchPortCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSwitchPortCollectionGet binds the parameter order_by
func (o *SwitchPortCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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