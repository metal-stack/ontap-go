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

	"github.com/metal-stack/ontap-go/api/models"
)

// NewPublickeyModifyParams creates a new PublickeyModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPublickeyModifyParams() *PublickeyModifyParams {
	return &PublickeyModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPublickeyModifyParamsWithTimeout creates a new PublickeyModifyParams object
// with the ability to set a timeout on a request.
func NewPublickeyModifyParamsWithTimeout(timeout time.Duration) *PublickeyModifyParams {
	return &PublickeyModifyParams{
		timeout: timeout,
	}
}

// NewPublickeyModifyParamsWithContext creates a new PublickeyModifyParams object
// with the ability to set a context for a request.
func NewPublickeyModifyParamsWithContext(ctx context.Context) *PublickeyModifyParams {
	return &PublickeyModifyParams{
		Context: ctx,
	}
}

// NewPublickeyModifyParamsWithHTTPClient creates a new PublickeyModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewPublickeyModifyParamsWithHTTPClient(client *http.Client) *PublickeyModifyParams {
	return &PublickeyModifyParams{
		HTTPClient: client,
	}
}

/*
PublickeyModifyParams contains all the parameters to send to the API endpoint

	for the publickey modify operation.

	Typically these are written to a http.Request.
*/
type PublickeyModifyParams struct {

	/* AccountName.

	   User account name
	*/
	AccountName string

	/* Index.

	   Index number for the public key (where there are multiple keys for the same account).
	*/
	Index int64

	/* Info.

	   Public key modification details.
	*/
	Info *models.Publickey

	/* OwnerUUID.

	   Account owner UUID
	*/
	OwnerUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the publickey modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublickeyModifyParams) WithDefaults() *PublickeyModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the publickey modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublickeyModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the publickey modify params
func (o *PublickeyModifyParams) WithTimeout(timeout time.Duration) *PublickeyModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the publickey modify params
func (o *PublickeyModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the publickey modify params
func (o *PublickeyModifyParams) WithContext(ctx context.Context) *PublickeyModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the publickey modify params
func (o *PublickeyModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the publickey modify params
func (o *PublickeyModifyParams) WithHTTPClient(client *http.Client) *PublickeyModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the publickey modify params
func (o *PublickeyModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountName adds the accountName to the publickey modify params
func (o *PublickeyModifyParams) WithAccountName(accountName string) *PublickeyModifyParams {
	o.SetAccountName(accountName)
	return o
}

// SetAccountName adds the accountName to the publickey modify params
func (o *PublickeyModifyParams) SetAccountName(accountName string) {
	o.AccountName = accountName
}

// WithIndex adds the index to the publickey modify params
func (o *PublickeyModifyParams) WithIndex(index int64) *PublickeyModifyParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the publickey modify params
func (o *PublickeyModifyParams) SetIndex(index int64) {
	o.Index = index
}

// WithInfo adds the info to the publickey modify params
func (o *PublickeyModifyParams) WithInfo(info *models.Publickey) *PublickeyModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the publickey modify params
func (o *PublickeyModifyParams) SetInfo(info *models.Publickey) {
	o.Info = info
}

// WithOwnerUUID adds the ownerUUID to the publickey modify params
func (o *PublickeyModifyParams) WithOwnerUUID(ownerUUID string) *PublickeyModifyParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the publickey modify params
func (o *PublickeyModifyParams) SetOwnerUUID(ownerUUID string) {
	o.OwnerUUID = ownerUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PublickeyModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account.name
	if err := r.SetPathParam("account.name", o.AccountName); err != nil {
		return err
	}

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param owner.uuid
	if err := r.SetPathParam("owner.uuid", o.OwnerUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
