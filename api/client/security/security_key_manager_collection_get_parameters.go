// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewSecurityKeyManagerCollectionGetParams creates a new SecurityKeyManagerCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityKeyManagerCollectionGetParams() *SecurityKeyManagerCollectionGetParams {
	return &SecurityKeyManagerCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityKeyManagerCollectionGetParamsWithTimeout creates a new SecurityKeyManagerCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSecurityKeyManagerCollectionGetParamsWithTimeout(timeout time.Duration) *SecurityKeyManagerCollectionGetParams {
	return &SecurityKeyManagerCollectionGetParams{
		timeout: timeout,
	}
}

// NewSecurityKeyManagerCollectionGetParamsWithContext creates a new SecurityKeyManagerCollectionGetParams object
// with the ability to set a context for a request.
func NewSecurityKeyManagerCollectionGetParamsWithContext(ctx context.Context) *SecurityKeyManagerCollectionGetParams {
	return &SecurityKeyManagerCollectionGetParams{
		Context: ctx,
	}
}

// NewSecurityKeyManagerCollectionGetParamsWithHTTPClient creates a new SecurityKeyManagerCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityKeyManagerCollectionGetParamsWithHTTPClient(client *http.Client) *SecurityKeyManagerCollectionGetParams {
	return &SecurityKeyManagerCollectionGetParams{
		HTTPClient: client,
	}
}

/*
SecurityKeyManagerCollectionGetParams contains all the parameters to send to the API endpoint

	for the security key manager collection get operation.

	Typically these are written to a http.Request.
*/
type SecurityKeyManagerCollectionGetParams struct {

	/* ExternalClientCertificateName.

	   Filter by external.client_certificate.name
	*/
	ExternalClientCertificateName *string

	/* ExternalClientCertificateUUID.

	   Filter by external.client_certificate.uuid
	*/
	ExternalClientCertificateUUID *string

	/* ExternalServerCaCertificatesName.

	   Filter by external.server_ca_certificates.name
	*/
	ExternalServerCaCertificatesName *string

	/* ExternalServerCaCertificatesUUID.

	   Filter by external.server_ca_certificates.uuid
	*/
	ExternalServerCaCertificatesUUID *string

	/* ExternalServersConnectivityClusterAvailability.

	   Filter by external.servers.connectivity.cluster_availability
	*/
	ExternalServersConnectivityClusterAvailability *bool

	/* ExternalServersConnectivityNodeStatesNodeName.

	   Filter by external.servers.connectivity.node_states.node.name
	*/
	ExternalServersConnectivityNodeStatesNodeName *string

	/* ExternalServersConnectivityNodeStatesNodeUUID.

	   Filter by external.servers.connectivity.node_states.node.uuid
	*/
	ExternalServersConnectivityNodeStatesNodeUUID *string

	/* ExternalServersConnectivityNodeStatesState.

	   Filter by external.servers.connectivity.node_states.state
	*/
	ExternalServersConnectivityNodeStatesState *string

	/* ExternalServersSecondaryKeyServers.

	   Filter by external.servers.secondary_key_servers
	*/
	ExternalServersSecondaryKeyServers *string

	/* ExternalServersServer.

	   Filter by external.servers.server
	*/
	ExternalServersServer *string

	/* ExternalServersTimeout.

	   Filter by external.servers.timeout
	*/
	ExternalServersTimeout *int64

	/* ExternalServersUsername.

	   Filter by external.servers.username
	*/
	ExternalServersUsername *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* IsDefaultDataAtRestEncryptionDisabled.

	   Filter by is_default_data_at_rest_encryption_disabled
	*/
	IsDefaultDataAtRestEncryptionDisabled *bool

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OnboardEnabled.

	   Filter by onboard.enabled
	*/
	OnboardEnabled *bool

	/* OnboardKeyBackup.

	   Filter by onboard.key_backup
	*/
	OnboardKeyBackup *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* Policy.

	   Filter by policy
	*/
	Policy *string

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

	/* Scope.

	   Filter by scope
	*/
	Scope *string

	/* StatusCode.

	   Filter by status.code
	*/
	StatusCode *int64

	/* StatusMessage.

	   Filter by status.message
	*/
	StatusMessage *string

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

	/* VolumeEncryptionCode.

	   Filter by volume_encryption.code
	*/
	VolumeEncryptionCode *int64

	/* VolumeEncryptionMessage.

	   Filter by volume_encryption.message
	*/
	VolumeEncryptionMessage *string

	/* VolumeEncryptionSupported.

	   Filter by volume_encryption.supported
	*/
	VolumeEncryptionSupported *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security key manager collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerCollectionGetParams) WithDefaults() *SecurityKeyManagerCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security key manager collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SecurityKeyManagerCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithTimeout(timeout time.Duration) *SecurityKeyManagerCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithContext(ctx context.Context) *SecurityKeyManagerCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithHTTPClient(client *http.Client) *SecurityKeyManagerCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExternalClientCertificateName adds the externalClientCertificateName to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalClientCertificateName(externalClientCertificateName *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalClientCertificateName(externalClientCertificateName)
	return o
}

// SetExternalClientCertificateName adds the externalClientCertificateName to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalClientCertificateName(externalClientCertificateName *string) {
	o.ExternalClientCertificateName = externalClientCertificateName
}

// WithExternalClientCertificateUUID adds the externalClientCertificateUUID to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalClientCertificateUUID(externalClientCertificateUUID *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalClientCertificateUUID(externalClientCertificateUUID)
	return o
}

// SetExternalClientCertificateUUID adds the externalClientCertificateUuid to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalClientCertificateUUID(externalClientCertificateUUID *string) {
	o.ExternalClientCertificateUUID = externalClientCertificateUUID
}

// WithExternalServerCaCertificatesName adds the externalServerCaCertificatesName to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServerCaCertificatesName(externalServerCaCertificatesName *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServerCaCertificatesName(externalServerCaCertificatesName)
	return o
}

// SetExternalServerCaCertificatesName adds the externalServerCaCertificatesName to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServerCaCertificatesName(externalServerCaCertificatesName *string) {
	o.ExternalServerCaCertificatesName = externalServerCaCertificatesName
}

// WithExternalServerCaCertificatesUUID adds the externalServerCaCertificatesUUID to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServerCaCertificatesUUID(externalServerCaCertificatesUUID *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServerCaCertificatesUUID(externalServerCaCertificatesUUID)
	return o
}

// SetExternalServerCaCertificatesUUID adds the externalServerCaCertificatesUuid to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServerCaCertificatesUUID(externalServerCaCertificatesUUID *string) {
	o.ExternalServerCaCertificatesUUID = externalServerCaCertificatesUUID
}

// WithExternalServersConnectivityClusterAvailability adds the externalServersConnectivityClusterAvailability to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServersConnectivityClusterAvailability(externalServersConnectivityClusterAvailability *bool) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServersConnectivityClusterAvailability(externalServersConnectivityClusterAvailability)
	return o
}

// SetExternalServersConnectivityClusterAvailability adds the externalServersConnectivityClusterAvailability to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServersConnectivityClusterAvailability(externalServersConnectivityClusterAvailability *bool) {
	o.ExternalServersConnectivityClusterAvailability = externalServersConnectivityClusterAvailability
}

// WithExternalServersConnectivityNodeStatesNodeName adds the externalServersConnectivityNodeStatesNodeName to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServersConnectivityNodeStatesNodeName(externalServersConnectivityNodeStatesNodeName *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServersConnectivityNodeStatesNodeName(externalServersConnectivityNodeStatesNodeName)
	return o
}

// SetExternalServersConnectivityNodeStatesNodeName adds the externalServersConnectivityNodeStatesNodeName to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServersConnectivityNodeStatesNodeName(externalServersConnectivityNodeStatesNodeName *string) {
	o.ExternalServersConnectivityNodeStatesNodeName = externalServersConnectivityNodeStatesNodeName
}

// WithExternalServersConnectivityNodeStatesNodeUUID adds the externalServersConnectivityNodeStatesNodeUUID to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServersConnectivityNodeStatesNodeUUID(externalServersConnectivityNodeStatesNodeUUID *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServersConnectivityNodeStatesNodeUUID(externalServersConnectivityNodeStatesNodeUUID)
	return o
}

// SetExternalServersConnectivityNodeStatesNodeUUID adds the externalServersConnectivityNodeStatesNodeUuid to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServersConnectivityNodeStatesNodeUUID(externalServersConnectivityNodeStatesNodeUUID *string) {
	o.ExternalServersConnectivityNodeStatesNodeUUID = externalServersConnectivityNodeStatesNodeUUID
}

// WithExternalServersConnectivityNodeStatesState adds the externalServersConnectivityNodeStatesState to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServersConnectivityNodeStatesState(externalServersConnectivityNodeStatesState *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServersConnectivityNodeStatesState(externalServersConnectivityNodeStatesState)
	return o
}

// SetExternalServersConnectivityNodeStatesState adds the externalServersConnectivityNodeStatesState to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServersConnectivityNodeStatesState(externalServersConnectivityNodeStatesState *string) {
	o.ExternalServersConnectivityNodeStatesState = externalServersConnectivityNodeStatesState
}

// WithExternalServersSecondaryKeyServers adds the externalServersSecondaryKeyServers to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServersSecondaryKeyServers(externalServersSecondaryKeyServers *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServersSecondaryKeyServers(externalServersSecondaryKeyServers)
	return o
}

// SetExternalServersSecondaryKeyServers adds the externalServersSecondaryKeyServers to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServersSecondaryKeyServers(externalServersSecondaryKeyServers *string) {
	o.ExternalServersSecondaryKeyServers = externalServersSecondaryKeyServers
}

// WithExternalServersServer adds the externalServersServer to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServersServer(externalServersServer *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServersServer(externalServersServer)
	return o
}

// SetExternalServersServer adds the externalServersServer to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServersServer(externalServersServer *string) {
	o.ExternalServersServer = externalServersServer
}

// WithExternalServersTimeout adds the externalServersTimeout to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServersTimeout(externalServersTimeout *int64) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServersTimeout(externalServersTimeout)
	return o
}

// SetExternalServersTimeout adds the externalServersTimeout to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServersTimeout(externalServersTimeout *int64) {
	o.ExternalServersTimeout = externalServersTimeout
}

// WithExternalServersUsername adds the externalServersUsername to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithExternalServersUsername(externalServersUsername *string) *SecurityKeyManagerCollectionGetParams {
	o.SetExternalServersUsername(externalServersUsername)
	return o
}

// SetExternalServersUsername adds the externalServersUsername to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetExternalServersUsername(externalServersUsername *string) {
	o.ExternalServersUsername = externalServersUsername
}

// WithFields adds the fields to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithFields(fields []string) *SecurityKeyManagerCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIsDefaultDataAtRestEncryptionDisabled adds the isDefaultDataAtRestEncryptionDisabled to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithIsDefaultDataAtRestEncryptionDisabled(isDefaultDataAtRestEncryptionDisabled *bool) *SecurityKeyManagerCollectionGetParams {
	o.SetIsDefaultDataAtRestEncryptionDisabled(isDefaultDataAtRestEncryptionDisabled)
	return o
}

// SetIsDefaultDataAtRestEncryptionDisabled adds the isDefaultDataAtRestEncryptionDisabled to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetIsDefaultDataAtRestEncryptionDisabled(isDefaultDataAtRestEncryptionDisabled *bool) {
	o.IsDefaultDataAtRestEncryptionDisabled = isDefaultDataAtRestEncryptionDisabled
}

// WithMaxRecords adds the maxRecords to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithMaxRecords(maxRecords *int64) *SecurityKeyManagerCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOnboardEnabled adds the onboardEnabled to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithOnboardEnabled(onboardEnabled *bool) *SecurityKeyManagerCollectionGetParams {
	o.SetOnboardEnabled(onboardEnabled)
	return o
}

// SetOnboardEnabled adds the onboardEnabled to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetOnboardEnabled(onboardEnabled *bool) {
	o.OnboardEnabled = onboardEnabled
}

// WithOnboardKeyBackup adds the onboardKeyBackup to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithOnboardKeyBackup(onboardKeyBackup *string) *SecurityKeyManagerCollectionGetParams {
	o.SetOnboardKeyBackup(onboardKeyBackup)
	return o
}

// SetOnboardKeyBackup adds the onboardKeyBackup to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetOnboardKeyBackup(onboardKeyBackup *string) {
	o.OnboardKeyBackup = onboardKeyBackup
}

// WithOrderBy adds the orderBy to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithOrderBy(orderBy []string) *SecurityKeyManagerCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPolicy adds the policy to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithPolicy(policy *string) *SecurityKeyManagerCollectionGetParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetPolicy(policy *string) {
	o.Policy = policy
}

// WithReturnRecords adds the returnRecords to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithReturnRecords(returnRecords *bool) *SecurityKeyManagerCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *SecurityKeyManagerCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScope adds the scope to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithScope(scope *string) *SecurityKeyManagerCollectionGetParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithStatusCode adds the statusCode to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithStatusCode(statusCode *int64) *SecurityKeyManagerCollectionGetParams {
	o.SetStatusCode(statusCode)
	return o
}

// SetStatusCode adds the statusCode to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetStatusCode(statusCode *int64) {
	o.StatusCode = statusCode
}

// WithStatusMessage adds the statusMessage to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithStatusMessage(statusMessage *string) *SecurityKeyManagerCollectionGetParams {
	o.SetStatusMessage(statusMessage)
	return o
}

// SetStatusMessage adds the statusMessage to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetStatusMessage(statusMessage *string) {
	o.StatusMessage = statusMessage
}

// WithSvmName adds the svmName to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithSvmName(svmName *string) *SecurityKeyManagerCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithSvmUUID(svmUUID *string) *SecurityKeyManagerCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithUUID adds the uuid to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithUUID(uuid *string) *SecurityKeyManagerCollectionGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WithVolumeEncryptionCode adds the volumeEncryptionCode to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithVolumeEncryptionCode(volumeEncryptionCode *int64) *SecurityKeyManagerCollectionGetParams {
	o.SetVolumeEncryptionCode(volumeEncryptionCode)
	return o
}

// SetVolumeEncryptionCode adds the volumeEncryptionCode to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetVolumeEncryptionCode(volumeEncryptionCode *int64) {
	o.VolumeEncryptionCode = volumeEncryptionCode
}

// WithVolumeEncryptionMessage adds the volumeEncryptionMessage to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithVolumeEncryptionMessage(volumeEncryptionMessage *string) *SecurityKeyManagerCollectionGetParams {
	o.SetVolumeEncryptionMessage(volumeEncryptionMessage)
	return o
}

// SetVolumeEncryptionMessage adds the volumeEncryptionMessage to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetVolumeEncryptionMessage(volumeEncryptionMessage *string) {
	o.VolumeEncryptionMessage = volumeEncryptionMessage
}

// WithVolumeEncryptionSupported adds the volumeEncryptionSupported to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) WithVolumeEncryptionSupported(volumeEncryptionSupported *bool) *SecurityKeyManagerCollectionGetParams {
	o.SetVolumeEncryptionSupported(volumeEncryptionSupported)
	return o
}

// SetVolumeEncryptionSupported adds the volumeEncryptionSupported to the security key manager collection get params
func (o *SecurityKeyManagerCollectionGetParams) SetVolumeEncryptionSupported(volumeEncryptionSupported *bool) {
	o.VolumeEncryptionSupported = volumeEncryptionSupported
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityKeyManagerCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExternalClientCertificateName != nil {

		// query param external.client_certificate.name
		var qrExternalClientCertificateName string

		if o.ExternalClientCertificateName != nil {
			qrExternalClientCertificateName = *o.ExternalClientCertificateName
		}
		qExternalClientCertificateName := qrExternalClientCertificateName
		if qExternalClientCertificateName != "" {

			if err := r.SetQueryParam("external.client_certificate.name", qExternalClientCertificateName); err != nil {
				return err
			}
		}
	}

	if o.ExternalClientCertificateUUID != nil {

		// query param external.client_certificate.uuid
		var qrExternalClientCertificateUUID string

		if o.ExternalClientCertificateUUID != nil {
			qrExternalClientCertificateUUID = *o.ExternalClientCertificateUUID
		}
		qExternalClientCertificateUUID := qrExternalClientCertificateUUID
		if qExternalClientCertificateUUID != "" {

			if err := r.SetQueryParam("external.client_certificate.uuid", qExternalClientCertificateUUID); err != nil {
				return err
			}
		}
	}

	if o.ExternalServerCaCertificatesName != nil {

		// query param external.server_ca_certificates.name
		var qrExternalServerCaCertificatesName string

		if o.ExternalServerCaCertificatesName != nil {
			qrExternalServerCaCertificatesName = *o.ExternalServerCaCertificatesName
		}
		qExternalServerCaCertificatesName := qrExternalServerCaCertificatesName
		if qExternalServerCaCertificatesName != "" {

			if err := r.SetQueryParam("external.server_ca_certificates.name", qExternalServerCaCertificatesName); err != nil {
				return err
			}
		}
	}

	if o.ExternalServerCaCertificatesUUID != nil {

		// query param external.server_ca_certificates.uuid
		var qrExternalServerCaCertificatesUUID string

		if o.ExternalServerCaCertificatesUUID != nil {
			qrExternalServerCaCertificatesUUID = *o.ExternalServerCaCertificatesUUID
		}
		qExternalServerCaCertificatesUUID := qrExternalServerCaCertificatesUUID
		if qExternalServerCaCertificatesUUID != "" {

			if err := r.SetQueryParam("external.server_ca_certificates.uuid", qExternalServerCaCertificatesUUID); err != nil {
				return err
			}
		}
	}

	if o.ExternalServersConnectivityClusterAvailability != nil {

		// query param external.servers.connectivity.cluster_availability
		var qrExternalServersConnectivityClusterAvailability bool

		if o.ExternalServersConnectivityClusterAvailability != nil {
			qrExternalServersConnectivityClusterAvailability = *o.ExternalServersConnectivityClusterAvailability
		}
		qExternalServersConnectivityClusterAvailability := swag.FormatBool(qrExternalServersConnectivityClusterAvailability)
		if qExternalServersConnectivityClusterAvailability != "" {

			if err := r.SetQueryParam("external.servers.connectivity.cluster_availability", qExternalServersConnectivityClusterAvailability); err != nil {
				return err
			}
		}
	}

	if o.ExternalServersConnectivityNodeStatesNodeName != nil {

		// query param external.servers.connectivity.node_states.node.name
		var qrExternalServersConnectivityNodeStatesNodeName string

		if o.ExternalServersConnectivityNodeStatesNodeName != nil {
			qrExternalServersConnectivityNodeStatesNodeName = *o.ExternalServersConnectivityNodeStatesNodeName
		}
		qExternalServersConnectivityNodeStatesNodeName := qrExternalServersConnectivityNodeStatesNodeName
		if qExternalServersConnectivityNodeStatesNodeName != "" {

			if err := r.SetQueryParam("external.servers.connectivity.node_states.node.name", qExternalServersConnectivityNodeStatesNodeName); err != nil {
				return err
			}
		}
	}

	if o.ExternalServersConnectivityNodeStatesNodeUUID != nil {

		// query param external.servers.connectivity.node_states.node.uuid
		var qrExternalServersConnectivityNodeStatesNodeUUID string

		if o.ExternalServersConnectivityNodeStatesNodeUUID != nil {
			qrExternalServersConnectivityNodeStatesNodeUUID = *o.ExternalServersConnectivityNodeStatesNodeUUID
		}
		qExternalServersConnectivityNodeStatesNodeUUID := qrExternalServersConnectivityNodeStatesNodeUUID
		if qExternalServersConnectivityNodeStatesNodeUUID != "" {

			if err := r.SetQueryParam("external.servers.connectivity.node_states.node.uuid", qExternalServersConnectivityNodeStatesNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.ExternalServersConnectivityNodeStatesState != nil {

		// query param external.servers.connectivity.node_states.state
		var qrExternalServersConnectivityNodeStatesState string

		if o.ExternalServersConnectivityNodeStatesState != nil {
			qrExternalServersConnectivityNodeStatesState = *o.ExternalServersConnectivityNodeStatesState
		}
		qExternalServersConnectivityNodeStatesState := qrExternalServersConnectivityNodeStatesState
		if qExternalServersConnectivityNodeStatesState != "" {

			if err := r.SetQueryParam("external.servers.connectivity.node_states.state", qExternalServersConnectivityNodeStatesState); err != nil {
				return err
			}
		}
	}

	if o.ExternalServersSecondaryKeyServers != nil {

		// query param external.servers.secondary_key_servers
		var qrExternalServersSecondaryKeyServers string

		if o.ExternalServersSecondaryKeyServers != nil {
			qrExternalServersSecondaryKeyServers = *o.ExternalServersSecondaryKeyServers
		}
		qExternalServersSecondaryKeyServers := qrExternalServersSecondaryKeyServers
		if qExternalServersSecondaryKeyServers != "" {

			if err := r.SetQueryParam("external.servers.secondary_key_servers", qExternalServersSecondaryKeyServers); err != nil {
				return err
			}
		}
	}

	if o.ExternalServersServer != nil {

		// query param external.servers.server
		var qrExternalServersServer string

		if o.ExternalServersServer != nil {
			qrExternalServersServer = *o.ExternalServersServer
		}
		qExternalServersServer := qrExternalServersServer
		if qExternalServersServer != "" {

			if err := r.SetQueryParam("external.servers.server", qExternalServersServer); err != nil {
				return err
			}
		}
	}

	if o.ExternalServersTimeout != nil {

		// query param external.servers.timeout
		var qrExternalServersTimeout int64

		if o.ExternalServersTimeout != nil {
			qrExternalServersTimeout = *o.ExternalServersTimeout
		}
		qExternalServersTimeout := swag.FormatInt64(qrExternalServersTimeout)
		if qExternalServersTimeout != "" {

			if err := r.SetQueryParam("external.servers.timeout", qExternalServersTimeout); err != nil {
				return err
			}
		}
	}

	if o.ExternalServersUsername != nil {

		// query param external.servers.username
		var qrExternalServersUsername string

		if o.ExternalServersUsername != nil {
			qrExternalServersUsername = *o.ExternalServersUsername
		}
		qExternalServersUsername := qrExternalServersUsername
		if qExternalServersUsername != "" {

			if err := r.SetQueryParam("external.servers.username", qExternalServersUsername); err != nil {
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

	if o.IsDefaultDataAtRestEncryptionDisabled != nil {

		// query param is_default_data_at_rest_encryption_disabled
		var qrIsDefaultDataAtRestEncryptionDisabled bool

		if o.IsDefaultDataAtRestEncryptionDisabled != nil {
			qrIsDefaultDataAtRestEncryptionDisabled = *o.IsDefaultDataAtRestEncryptionDisabled
		}
		qIsDefaultDataAtRestEncryptionDisabled := swag.FormatBool(qrIsDefaultDataAtRestEncryptionDisabled)
		if qIsDefaultDataAtRestEncryptionDisabled != "" {

			if err := r.SetQueryParam("is_default_data_at_rest_encryption_disabled", qIsDefaultDataAtRestEncryptionDisabled); err != nil {
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

	if o.OnboardEnabled != nil {

		// query param onboard.enabled
		var qrOnboardEnabled bool

		if o.OnboardEnabled != nil {
			qrOnboardEnabled = *o.OnboardEnabled
		}
		qOnboardEnabled := swag.FormatBool(qrOnboardEnabled)
		if qOnboardEnabled != "" {

			if err := r.SetQueryParam("onboard.enabled", qOnboardEnabled); err != nil {
				return err
			}
		}
	}

	if o.OnboardKeyBackup != nil {

		// query param onboard.key_backup
		var qrOnboardKeyBackup string

		if o.OnboardKeyBackup != nil {
			qrOnboardKeyBackup = *o.OnboardKeyBackup
		}
		qOnboardKeyBackup := qrOnboardKeyBackup
		if qOnboardKeyBackup != "" {

			if err := r.SetQueryParam("onboard.key_backup", qOnboardKeyBackup); err != nil {
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

	if o.Policy != nil {

		// query param policy
		var qrPolicy string

		if o.Policy != nil {
			qrPolicy = *o.Policy
		}
		qPolicy := qrPolicy
		if qPolicy != "" {

			if err := r.SetQueryParam("policy", qPolicy); err != nil {
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

	if o.Scope != nil {

		// query param scope
		var qrScope string

		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}
	}

	if o.StatusCode != nil {

		// query param status.code
		var qrStatusCode int64

		if o.StatusCode != nil {
			qrStatusCode = *o.StatusCode
		}
		qStatusCode := swag.FormatInt64(qrStatusCode)
		if qStatusCode != "" {

			if err := r.SetQueryParam("status.code", qStatusCode); err != nil {
				return err
			}
		}
	}

	if o.StatusMessage != nil {

		// query param status.message
		var qrStatusMessage string

		if o.StatusMessage != nil {
			qrStatusMessage = *o.StatusMessage
		}
		qStatusMessage := qrStatusMessage
		if qStatusMessage != "" {

			if err := r.SetQueryParam("status.message", qStatusMessage); err != nil {
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

	if o.VolumeEncryptionCode != nil {

		// query param volume_encryption.code
		var qrVolumeEncryptionCode int64

		if o.VolumeEncryptionCode != nil {
			qrVolumeEncryptionCode = *o.VolumeEncryptionCode
		}
		qVolumeEncryptionCode := swag.FormatInt64(qrVolumeEncryptionCode)
		if qVolumeEncryptionCode != "" {

			if err := r.SetQueryParam("volume_encryption.code", qVolumeEncryptionCode); err != nil {
				return err
			}
		}
	}

	if o.VolumeEncryptionMessage != nil {

		// query param volume_encryption.message
		var qrVolumeEncryptionMessage string

		if o.VolumeEncryptionMessage != nil {
			qrVolumeEncryptionMessage = *o.VolumeEncryptionMessage
		}
		qVolumeEncryptionMessage := qrVolumeEncryptionMessage
		if qVolumeEncryptionMessage != "" {

			if err := r.SetQueryParam("volume_encryption.message", qVolumeEncryptionMessage); err != nil {
				return err
			}
		}
	}

	if o.VolumeEncryptionSupported != nil {

		// query param volume_encryption.supported
		var qrVolumeEncryptionSupported bool

		if o.VolumeEncryptionSupported != nil {
			qrVolumeEncryptionSupported = *o.VolumeEncryptionSupported
		}
		qVolumeEncryptionSupported := swag.FormatBool(qrVolumeEncryptionSupported)
		if qVolumeEncryptionSupported != "" {

			if err := r.SetQueryParam("volume_encryption.supported", qVolumeEncryptionSupported); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSecurityKeyManagerCollectionGet binds the parameter fields
func (o *SecurityKeyManagerCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSecurityKeyManagerCollectionGet binds the parameter order_by
func (o *SecurityKeyManagerCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
