// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigLoadbalancerAllOKCode is the HTTP code returned for type GetConfigLoadbalancerAllOK
const GetConfigLoadbalancerAllOKCode int = 200

/*
GetConfigLoadbalancerAllOK OK

swagger:response getConfigLoadbalancerAllOK
*/
type GetConfigLoadbalancerAllOK struct {

	/*
	  In: Body
	*/
	Payload *GetConfigLoadbalancerAllOKBody `json:"body,omitempty"`
}

// NewGetConfigLoadbalancerAllOK creates GetConfigLoadbalancerAllOK with default headers values
func NewGetConfigLoadbalancerAllOK() *GetConfigLoadbalancerAllOK {

	return &GetConfigLoadbalancerAllOK{}
}

// WithPayload adds the payload to the get config loadbalancer all o k response
func (o *GetConfigLoadbalancerAllOK) WithPayload(payload *GetConfigLoadbalancerAllOKBody) *GetConfigLoadbalancerAllOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config loadbalancer all o k response
func (o *GetConfigLoadbalancerAllOK) SetPayload(payload *GetConfigLoadbalancerAllOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigLoadbalancerAllOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigLoadbalancerAllUnauthorizedCode is the HTTP code returned for type GetConfigLoadbalancerAllUnauthorized
const GetConfigLoadbalancerAllUnauthorizedCode int = 401

/*
GetConfigLoadbalancerAllUnauthorized Invalid authentication credentials

swagger:response getConfigLoadbalancerAllUnauthorized
*/
type GetConfigLoadbalancerAllUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigLoadbalancerAllUnauthorized creates GetConfigLoadbalancerAllUnauthorized with default headers values
func NewGetConfigLoadbalancerAllUnauthorized() *GetConfigLoadbalancerAllUnauthorized {

	return &GetConfigLoadbalancerAllUnauthorized{}
}

// WithPayload adds the payload to the get config loadbalancer all unauthorized response
func (o *GetConfigLoadbalancerAllUnauthorized) WithPayload(payload *models.Error) *GetConfigLoadbalancerAllUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config loadbalancer all unauthorized response
func (o *GetConfigLoadbalancerAllUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigLoadbalancerAllUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigLoadbalancerAllInternalServerErrorCode is the HTTP code returned for type GetConfigLoadbalancerAllInternalServerError
const GetConfigLoadbalancerAllInternalServerErrorCode int = 500

/*
GetConfigLoadbalancerAllInternalServerError Internal service error

swagger:response getConfigLoadbalancerAllInternalServerError
*/
type GetConfigLoadbalancerAllInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigLoadbalancerAllInternalServerError creates GetConfigLoadbalancerAllInternalServerError with default headers values
func NewGetConfigLoadbalancerAllInternalServerError() *GetConfigLoadbalancerAllInternalServerError {

	return &GetConfigLoadbalancerAllInternalServerError{}
}

// WithPayload adds the payload to the get config loadbalancer all internal server error response
func (o *GetConfigLoadbalancerAllInternalServerError) WithPayload(payload *models.Error) *GetConfigLoadbalancerAllInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config loadbalancer all internal server error response
func (o *GetConfigLoadbalancerAllInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigLoadbalancerAllInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigLoadbalancerAllServiceUnavailableCode is the HTTP code returned for type GetConfigLoadbalancerAllServiceUnavailable
const GetConfigLoadbalancerAllServiceUnavailableCode int = 503

/*
GetConfigLoadbalancerAllServiceUnavailable Maintenance mode

swagger:response getConfigLoadbalancerAllServiceUnavailable
*/
type GetConfigLoadbalancerAllServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigLoadbalancerAllServiceUnavailable creates GetConfigLoadbalancerAllServiceUnavailable with default headers values
func NewGetConfigLoadbalancerAllServiceUnavailable() *GetConfigLoadbalancerAllServiceUnavailable {

	return &GetConfigLoadbalancerAllServiceUnavailable{}
}

// WithPayload adds the payload to the get config loadbalancer all service unavailable response
func (o *GetConfigLoadbalancerAllServiceUnavailable) WithPayload(payload *models.Error) *GetConfigLoadbalancerAllServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config loadbalancer all service unavailable response
func (o *GetConfigLoadbalancerAllServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigLoadbalancerAllServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
