// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigVlanAllOKCode is the HTTP code returned for type GetConfigVlanAllOK
const GetConfigVlanAllOKCode int = 200

/*
GetConfigVlanAllOK OK

swagger:response getConfigVlanAllOK
*/
type GetConfigVlanAllOK struct {

	/*
	  In: Body
	*/
	Payload *GetConfigVlanAllOKBody `json:"body,omitempty"`
}

// NewGetConfigVlanAllOK creates GetConfigVlanAllOK with default headers values
func NewGetConfigVlanAllOK() *GetConfigVlanAllOK {

	return &GetConfigVlanAllOK{}
}

// WithPayload adds the payload to the get config vlan all o k response
func (o *GetConfigVlanAllOK) WithPayload(payload *GetConfigVlanAllOKBody) *GetConfigVlanAllOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config vlan all o k response
func (o *GetConfigVlanAllOK) SetPayload(payload *GetConfigVlanAllOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigVlanAllOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigVlanAllUnauthorizedCode is the HTTP code returned for type GetConfigVlanAllUnauthorized
const GetConfigVlanAllUnauthorizedCode int = 401

/*
GetConfigVlanAllUnauthorized Invalid authentication credentials

swagger:response getConfigVlanAllUnauthorized
*/
type GetConfigVlanAllUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigVlanAllUnauthorized creates GetConfigVlanAllUnauthorized with default headers values
func NewGetConfigVlanAllUnauthorized() *GetConfigVlanAllUnauthorized {

	return &GetConfigVlanAllUnauthorized{}
}

// WithPayload adds the payload to the get config vlan all unauthorized response
func (o *GetConfigVlanAllUnauthorized) WithPayload(payload *models.Error) *GetConfigVlanAllUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config vlan all unauthorized response
func (o *GetConfigVlanAllUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigVlanAllUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigVlanAllInternalServerErrorCode is the HTTP code returned for type GetConfigVlanAllInternalServerError
const GetConfigVlanAllInternalServerErrorCode int = 500

/*
GetConfigVlanAllInternalServerError Internal service error

swagger:response getConfigVlanAllInternalServerError
*/
type GetConfigVlanAllInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigVlanAllInternalServerError creates GetConfigVlanAllInternalServerError with default headers values
func NewGetConfigVlanAllInternalServerError() *GetConfigVlanAllInternalServerError {

	return &GetConfigVlanAllInternalServerError{}
}

// WithPayload adds the payload to the get config vlan all internal server error response
func (o *GetConfigVlanAllInternalServerError) WithPayload(payload *models.Error) *GetConfigVlanAllInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config vlan all internal server error response
func (o *GetConfigVlanAllInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigVlanAllInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigVlanAllServiceUnavailableCode is the HTTP code returned for type GetConfigVlanAllServiceUnavailable
const GetConfigVlanAllServiceUnavailableCode int = 503

/*
GetConfigVlanAllServiceUnavailable Maintenance mode

swagger:response getConfigVlanAllServiceUnavailable
*/
type GetConfigVlanAllServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigVlanAllServiceUnavailable creates GetConfigVlanAllServiceUnavailable with default headers values
func NewGetConfigVlanAllServiceUnavailable() *GetConfigVlanAllServiceUnavailable {

	return &GetConfigVlanAllServiceUnavailable{}
}

// WithPayload adds the payload to the get config vlan all service unavailable response
func (o *GetConfigVlanAllServiceUnavailable) WithPayload(payload *models.Error) *GetConfigVlanAllServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config vlan all service unavailable response
func (o *GetConfigVlanAllServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigVlanAllServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
