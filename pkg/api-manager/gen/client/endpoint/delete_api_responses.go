///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// DeleteAPIReader is a Reader for the DeleteAPI structure.
type DeleteAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteAPIBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteAPIUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteAPIForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteAPINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteAPIDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAPIOK creates a DeleteAPIOK with default headers values
func NewDeleteAPIOK() *DeleteAPIOK {
	return &DeleteAPIOK{}
}

/*DeleteAPIOK handles this case with default header values.

Successful operation
*/
type DeleteAPIOK struct {
	Payload *v1.API
}

func (o *DeleteAPIOK) Error() string {
	return fmt.Sprintf("[DELETE /{api}][%d] deleteApiOK  %+v", 200, o.Payload)
}

func (o *DeleteAPIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.API)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIBadRequest creates a DeleteAPIBadRequest with default headers values
func NewDeleteAPIBadRequest() *DeleteAPIBadRequest {
	return &DeleteAPIBadRequest{}
}

/*DeleteAPIBadRequest handles this case with default header values.

Invalid Name supplied
*/
type DeleteAPIBadRequest struct {
	Payload *v1.Error
}

func (o *DeleteAPIBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /{api}][%d] deleteApiBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAPIBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIUnauthorized creates a DeleteAPIUnauthorized with default headers values
func NewDeleteAPIUnauthorized() *DeleteAPIUnauthorized {
	return &DeleteAPIUnauthorized{}
}

/*DeleteAPIUnauthorized handles this case with default header values.

Unauthorized Request
*/
type DeleteAPIUnauthorized struct {
	Payload *v1.Error
}

func (o *DeleteAPIUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /{api}][%d] deleteApiUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAPIUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIForbidden creates a DeleteAPIForbidden with default headers values
func NewDeleteAPIForbidden() *DeleteAPIForbidden {
	return &DeleteAPIForbidden{}
}

/*DeleteAPIForbidden handles this case with default header values.

access to this resource is forbidden
*/
type DeleteAPIForbidden struct {
	Payload *v1.Error
}

func (o *DeleteAPIForbidden) Error() string {
	return fmt.Sprintf("[DELETE /{api}][%d] deleteApiForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAPIForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPINotFound creates a DeleteAPINotFound with default headers values
func NewDeleteAPINotFound() *DeleteAPINotFound {
	return &DeleteAPINotFound{}
}

/*DeleteAPINotFound handles this case with default header values.

API not found
*/
type DeleteAPINotFound struct {
	Payload *v1.Error
}

func (o *DeleteAPINotFound) Error() string {
	return fmt.Sprintf("[DELETE /{api}][%d] deleteApiNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAPINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIDefault creates a DeleteAPIDefault with default headers values
func NewDeleteAPIDefault(code int) *DeleteAPIDefault {
	return &DeleteAPIDefault{
		_statusCode: code,
	}
}

/*DeleteAPIDefault handles this case with default header values.

Unknown error
*/
type DeleteAPIDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the delete API default response
func (o *DeleteAPIDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAPIDefault) Error() string {
	return fmt.Sprintf("[DELETE /{api}][%d] deleteAPI default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAPIDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
