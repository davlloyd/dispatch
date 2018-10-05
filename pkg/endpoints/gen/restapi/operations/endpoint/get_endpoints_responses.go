///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetEndpointsOKCode is the HTTP code returned for type GetEndpointsOK
const GetEndpointsOKCode int = 200

/*GetEndpointsOK Successful operation

swagger:response getEndpointsOK
*/
type GetEndpointsOK struct {

	/*
	  In: Body
	*/
	Payload []*v1.Endpoint `json:"body,omitempty"`
}

// NewGetEndpointsOK creates GetEndpointsOK with default headers values
func NewGetEndpointsOK() *GetEndpointsOK {

	return &GetEndpointsOK{}
}

// WithPayload adds the payload to the get endpoints o k response
func (o *GetEndpointsOK) WithPayload(payload []*v1.Endpoint) *GetEndpointsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get endpoints o k response
func (o *GetEndpointsOK) SetPayload(payload []*v1.Endpoint) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEndpointsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*v1.Endpoint, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetEndpointsUnauthorizedCode is the HTTP code returned for type GetEndpointsUnauthorized
const GetEndpointsUnauthorizedCode int = 401

/*GetEndpointsUnauthorized Unauthorized Request

swagger:response getEndpointsUnauthorized
*/
type GetEndpointsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetEndpointsUnauthorized creates GetEndpointsUnauthorized with default headers values
func NewGetEndpointsUnauthorized() *GetEndpointsUnauthorized {

	return &GetEndpointsUnauthorized{}
}

// WithPayload adds the payload to the get endpoints unauthorized response
func (o *GetEndpointsUnauthorized) WithPayload(payload *v1.Error) *GetEndpointsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get endpoints unauthorized response
func (o *GetEndpointsUnauthorized) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEndpointsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEndpointsForbiddenCode is the HTTP code returned for type GetEndpointsForbidden
const GetEndpointsForbiddenCode int = 403

/*GetEndpointsForbidden access to this resource is forbidden

swagger:response getEndpointsForbidden
*/
type GetEndpointsForbidden struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetEndpointsForbidden creates GetEndpointsForbidden with default headers values
func NewGetEndpointsForbidden() *GetEndpointsForbidden {

	return &GetEndpointsForbidden{}
}

// WithPayload adds the payload to the get endpoints forbidden response
func (o *GetEndpointsForbidden) WithPayload(payload *v1.Error) *GetEndpointsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get endpoints forbidden response
func (o *GetEndpointsForbidden) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEndpointsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetEndpointsDefault Unexpected Error

swagger:response getEndpointsDefault
*/
type GetEndpointsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetEndpointsDefault creates GetEndpointsDefault with default headers values
func NewGetEndpointsDefault(code int) *GetEndpointsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetEndpointsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get endpoints default response
func (o *GetEndpointsDefault) WithStatusCode(code int) *GetEndpointsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get endpoints default response
func (o *GetEndpointsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get endpoints default response
func (o *GetEndpointsDefault) WithPayload(payload *v1.Error) *GetEndpointsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get endpoints default response
func (o *GetEndpointsDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEndpointsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}