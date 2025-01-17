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

	"github.com/metal-stack/ontap-go/api/models"
)

// NewRolePrivilegeModifyParams creates a new RolePrivilegeModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRolePrivilegeModifyParams() *RolePrivilegeModifyParams {
	return &RolePrivilegeModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRolePrivilegeModifyParamsWithTimeout creates a new RolePrivilegeModifyParams object
// with the ability to set a timeout on a request.
func NewRolePrivilegeModifyParamsWithTimeout(timeout time.Duration) *RolePrivilegeModifyParams {
	return &RolePrivilegeModifyParams{
		timeout: timeout,
	}
}

// NewRolePrivilegeModifyParamsWithContext creates a new RolePrivilegeModifyParams object
// with the ability to set a context for a request.
func NewRolePrivilegeModifyParamsWithContext(ctx context.Context) *RolePrivilegeModifyParams {
	return &RolePrivilegeModifyParams{
		Context: ctx,
	}
}

// NewRolePrivilegeModifyParamsWithHTTPClient creates a new RolePrivilegeModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewRolePrivilegeModifyParamsWithHTTPClient(client *http.Client) *RolePrivilegeModifyParams {
	return &RolePrivilegeModifyParams{
		HTTPClient: client,
	}
}

/*
RolePrivilegeModifyParams contains all the parameters to send to the API endpoint

	for the role privilege modify operation.

	Typically these are written to a http.Request.
*/
type RolePrivilegeModifyParams struct {

	// Info.
	Info *models.RolePrivilege

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

// WithDefaults hydrates default values in the role privilege modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolePrivilegeModifyParams) WithDefaults() *RolePrivilegeModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the role privilege modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolePrivilegeModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the role privilege modify params
func (o *RolePrivilegeModifyParams) WithTimeout(timeout time.Duration) *RolePrivilegeModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the role privilege modify params
func (o *RolePrivilegeModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the role privilege modify params
func (o *RolePrivilegeModifyParams) WithContext(ctx context.Context) *RolePrivilegeModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the role privilege modify params
func (o *RolePrivilegeModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the role privilege modify params
func (o *RolePrivilegeModifyParams) WithHTTPClient(client *http.Client) *RolePrivilegeModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the role privilege modify params
func (o *RolePrivilegeModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the role privilege modify params
func (o *RolePrivilegeModifyParams) WithInfo(info *models.RolePrivilege) *RolePrivilegeModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the role privilege modify params
func (o *RolePrivilegeModifyParams) SetInfo(info *models.RolePrivilege) {
	o.Info = info
}

// WithName adds the name to the role privilege modify params
func (o *RolePrivilegeModifyParams) WithName(name string) *RolePrivilegeModifyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the role privilege modify params
func (o *RolePrivilegeModifyParams) SetName(name string) {
	o.Name = name
}

// WithOwnerUUID adds the ownerUUID to the role privilege modify params
func (o *RolePrivilegeModifyParams) WithOwnerUUID(ownerUUID string) *RolePrivilegeModifyParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the role privilege modify params
func (o *RolePrivilegeModifyParams) SetOwnerUUID(ownerUUID string) {
	o.OwnerUUID = ownerUUID
}

// WithPath adds the path to the role privilege modify params
func (o *RolePrivilegeModifyParams) WithPath(path string) *RolePrivilegeModifyParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the role privilege modify params
func (o *RolePrivilegeModifyParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *RolePrivilegeModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

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
