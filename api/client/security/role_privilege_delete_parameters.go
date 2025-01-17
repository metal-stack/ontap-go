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
)

// NewRolePrivilegeDeleteParams creates a new RolePrivilegeDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRolePrivilegeDeleteParams() *RolePrivilegeDeleteParams {
	return &RolePrivilegeDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRolePrivilegeDeleteParamsWithTimeout creates a new RolePrivilegeDeleteParams object
// with the ability to set a timeout on a request.
func NewRolePrivilegeDeleteParamsWithTimeout(timeout time.Duration) *RolePrivilegeDeleteParams {
	return &RolePrivilegeDeleteParams{
		timeout: timeout,
	}
}

// NewRolePrivilegeDeleteParamsWithContext creates a new RolePrivilegeDeleteParams object
// with the ability to set a context for a request.
func NewRolePrivilegeDeleteParamsWithContext(ctx context.Context) *RolePrivilegeDeleteParams {
	return &RolePrivilegeDeleteParams{
		Context: ctx,
	}
}

// NewRolePrivilegeDeleteParamsWithHTTPClient creates a new RolePrivilegeDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewRolePrivilegeDeleteParamsWithHTTPClient(client *http.Client) *RolePrivilegeDeleteParams {
	return &RolePrivilegeDeleteParams{
		HTTPClient: client,
	}
}

/*
RolePrivilegeDeleteParams contains all the parameters to send to the API endpoint

	for the role privilege delete operation.

	Typically these are written to a http.Request.
*/
type RolePrivilegeDeleteParams struct {

	/* Name.

	   Role name
	*/
	Name string

	/* OwnerUUID.

	   Role owner UUID
	*/
	OwnerUUID string

	/* Path.

	   REST API path or command/command directory path
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the role privilege delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolePrivilegeDeleteParams) WithDefaults() *RolePrivilegeDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the role privilege delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolePrivilegeDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the role privilege delete params
func (o *RolePrivilegeDeleteParams) WithTimeout(timeout time.Duration) *RolePrivilegeDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the role privilege delete params
func (o *RolePrivilegeDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the role privilege delete params
func (o *RolePrivilegeDeleteParams) WithContext(ctx context.Context) *RolePrivilegeDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the role privilege delete params
func (o *RolePrivilegeDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the role privilege delete params
func (o *RolePrivilegeDeleteParams) WithHTTPClient(client *http.Client) *RolePrivilegeDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the role privilege delete params
func (o *RolePrivilegeDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the role privilege delete params
func (o *RolePrivilegeDeleteParams) WithName(name string) *RolePrivilegeDeleteParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the role privilege delete params
func (o *RolePrivilegeDeleteParams) SetName(name string) {
	o.Name = name
}

// WithOwnerUUID adds the ownerUUID to the role privilege delete params
func (o *RolePrivilegeDeleteParams) WithOwnerUUID(ownerUUID string) *RolePrivilegeDeleteParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the role privilege delete params
func (o *RolePrivilegeDeleteParams) SetOwnerUUID(ownerUUID string) {
	o.OwnerUUID = ownerUUID
}

// WithPath adds the path to the role privilege delete params
func (o *RolePrivilegeDeleteParams) WithPath(path string) *RolePrivilegeDeleteParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the role privilege delete params
func (o *RolePrivilegeDeleteParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *RolePrivilegeDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param owner.uuid
	if err := r.SetPathParam("owner.uuid", o.OwnerUUID); err != nil {
		return err
	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}