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

// NewClusterLdapGetParams creates a new ClusterLdapGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterLdapGetParams() *ClusterLdapGetParams {
	return &ClusterLdapGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterLdapGetParamsWithTimeout creates a new ClusterLdapGetParams object
// with the ability to set a timeout on a request.
func NewClusterLdapGetParamsWithTimeout(timeout time.Duration) *ClusterLdapGetParams {
	return &ClusterLdapGetParams{
		timeout: timeout,
	}
}

// NewClusterLdapGetParamsWithContext creates a new ClusterLdapGetParams object
// with the ability to set a context for a request.
func NewClusterLdapGetParamsWithContext(ctx context.Context) *ClusterLdapGetParams {
	return &ClusterLdapGetParams{
		Context: ctx,
	}
}

// NewClusterLdapGetParamsWithHTTPClient creates a new ClusterLdapGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterLdapGetParamsWithHTTPClient(client *http.Client) *ClusterLdapGetParams {
	return &ClusterLdapGetParams{
		HTTPClient: client,
	}
}

/*
ClusterLdapGetParams contains all the parameters to send to the API endpoint

	for the cluster ldap get operation.

	Typically these are written to a http.Request.
*/
type ClusterLdapGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster ldap get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterLdapGetParams) WithDefaults() *ClusterLdapGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster ldap get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterLdapGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster ldap get params
func (o *ClusterLdapGetParams) WithTimeout(timeout time.Duration) *ClusterLdapGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster ldap get params
func (o *ClusterLdapGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster ldap get params
func (o *ClusterLdapGetParams) WithContext(ctx context.Context) *ClusterLdapGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster ldap get params
func (o *ClusterLdapGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster ldap get params
func (o *ClusterLdapGetParams) WithHTTPClient(client *http.Client) *ClusterLdapGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster ldap get params
func (o *ClusterLdapGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the cluster ldap get params
func (o *ClusterLdapGetParams) WithFields(fields []string) *ClusterLdapGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the cluster ldap get params
func (o *ClusterLdapGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterLdapGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamClusterLdapGet binds the parameter fields
func (o *ClusterLdapGetParams) bindParamFields(formats strfmt.Registry) []string {
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
