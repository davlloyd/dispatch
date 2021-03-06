///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package service_class

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetServiceClassByNameParams creates a new GetServiceClassByNameParams object
// with the default values initialized.
func NewGetServiceClassByNameParams() *GetServiceClassByNameParams {
	var ()
	return &GetServiceClassByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceClassByNameParamsWithTimeout creates a new GetServiceClassByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServiceClassByNameParamsWithTimeout(timeout time.Duration) *GetServiceClassByNameParams {
	var ()
	return &GetServiceClassByNameParams{

		timeout: timeout,
	}
}

// NewGetServiceClassByNameParamsWithContext creates a new GetServiceClassByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServiceClassByNameParamsWithContext(ctx context.Context) *GetServiceClassByNameParams {
	var ()
	return &GetServiceClassByNameParams{

		Context: ctx,
	}
}

// NewGetServiceClassByNameParamsWithHTTPClient creates a new GetServiceClassByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServiceClassByNameParamsWithHTTPClient(client *http.Client) *GetServiceClassByNameParams {
	var ()
	return &GetServiceClassByNameParams{
		HTTPClient: client,
	}
}

/*GetServiceClassByNameParams contains all the parameters to send to the API endpoint
for the get service class by name operation typically these are written to a http.Request
*/
type GetServiceClassByNameParams struct {

	/*XDispatchOrg*/
	XDispatchOrg string
	/*ServiceClassName
	  Name of service class to return

	*/
	ServiceClassName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service class by name params
func (o *GetServiceClassByNameParams) WithTimeout(timeout time.Duration) *GetServiceClassByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service class by name params
func (o *GetServiceClassByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service class by name params
func (o *GetServiceClassByNameParams) WithContext(ctx context.Context) *GetServiceClassByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service class by name params
func (o *GetServiceClassByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service class by name params
func (o *GetServiceClassByNameParams) WithHTTPClient(client *http.Client) *GetServiceClassByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service class by name params
func (o *GetServiceClassByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDispatchOrg adds the xDispatchOrg to the get service class by name params
func (o *GetServiceClassByNameParams) WithXDispatchOrg(xDispatchOrg string) *GetServiceClassByNameParams {
	o.SetXDispatchOrg(xDispatchOrg)
	return o
}

// SetXDispatchOrg adds the xDispatchOrg to the get service class by name params
func (o *GetServiceClassByNameParams) SetXDispatchOrg(xDispatchOrg string) {
	o.XDispatchOrg = xDispatchOrg
}

// WithServiceClassName adds the serviceClassName to the get service class by name params
func (o *GetServiceClassByNameParams) WithServiceClassName(serviceClassName string) *GetServiceClassByNameParams {
	o.SetServiceClassName(serviceClassName)
	return o
}

// SetServiceClassName adds the serviceClassName to the get service class by name params
func (o *GetServiceClassByNameParams) SetServiceClassName(serviceClassName string) {
	o.ServiceClassName = serviceClassName
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceClassByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Dispatch-Org
	if err := r.SetHeaderParam("X-Dispatch-Org", o.XDispatchOrg); err != nil {
		return err
	}

	// path param serviceClassName
	if err := r.SetPathParam("serviceClassName", o.ServiceClassName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
