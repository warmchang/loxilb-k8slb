// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNoContentCode is the HTTP code returned for type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNoContent
const DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNoContentCode int = 204

/*
DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNoContent OK

swagger:response deleteConfigLoadbalancerHosturlHosturlExternalipaddressIpAddressPortPortProtocolProtoNoContent
*/
type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNoContent struct {
}

// NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNoContent creates DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNoContent with default headers values
func NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNoContent() *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNoContent {

	return &DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNoContent{}
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoBadRequestCode is the HTTP code returned for type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoBadRequest
const DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoBadRequestCode int = 400

/*
DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoBadRequest Malformed arguments for API call

swagger:response deleteConfigLoadbalancerHosturlHosturlExternalipaddressIpAddressPortPortProtocolProtoBadRequest
*/
type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoBadRequest creates DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoBadRequest with default headers values
func NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoBadRequest() *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoBadRequest {

	return &DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoBadRequest{}
}

// WithPayload adds the payload to the delete config loadbalancer hosturl hosturl externalipaddress Ip address port port protocol proto bad request response
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoBadRequest) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer hosturl hosturl externalipaddress Ip address port port protocol proto bad request response
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoUnauthorizedCode is the HTTP code returned for type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoUnauthorized
const DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoUnauthorizedCode int = 401

/*
DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoUnauthorized Invalid authentication credentials

swagger:response deleteConfigLoadbalancerHosturlHosturlExternalipaddressIpAddressPortPortProtocolProtoUnauthorized
*/
type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoUnauthorized creates DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoUnauthorized with default headers values
func NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoUnauthorized() *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoUnauthorized {

	return &DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoUnauthorized{}
}

// WithPayload adds the payload to the delete config loadbalancer hosturl hosturl externalipaddress Ip address port port protocol proto unauthorized response
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoUnauthorized) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer hosturl hosturl externalipaddress Ip address port port protocol proto unauthorized response
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoForbiddenCode is the HTTP code returned for type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoForbidden
const DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoForbiddenCode int = 403

/*
DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoForbidden Capacity insufficient

swagger:response deleteConfigLoadbalancerHosturlHosturlExternalipaddressIpAddressPortPortProtocolProtoForbidden
*/
type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoForbidden creates DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoForbidden with default headers values
func NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoForbidden() *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoForbidden {

	return &DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoForbidden{}
}

// WithPayload adds the payload to the delete config loadbalancer hosturl hosturl externalipaddress Ip address port port protocol proto forbidden response
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoForbidden) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer hosturl hosturl externalipaddress Ip address port port protocol proto forbidden response
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNotFoundCode is the HTTP code returned for type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNotFound
const DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNotFoundCode int = 404

/*
DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNotFound Resource not found

swagger:response deleteConfigLoadbalancerHosturlHosturlExternalipaddressIpAddressPortPortProtocolProtoNotFound
*/
type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNotFound creates DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNotFound with default headers values
func NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNotFound() *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNotFound {

	return &DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNotFound{}
}

// WithPayload adds the payload to the delete config loadbalancer hosturl hosturl externalipaddress Ip address port port protocol proto not found response
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNotFound) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer hosturl hosturl externalipaddress Ip address port port protocol proto not found response
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoConflictCode is the HTTP code returned for type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoConflict
const DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoConflictCode int = 409

/*
DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoConflict Resource Conflict. VLAN already exists OR dependency VRF/VNET not found

swagger:response deleteConfigLoadbalancerHosturlHosturlExternalipaddressIpAddressPortPortProtocolProtoConflict
*/
type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoConflict creates DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoConflict with default headers values
func NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoConflict() *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoConflict {

	return &DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoConflict{}
}

// WithPayload adds the payload to the delete config loadbalancer hosturl hosturl externalipaddress Ip address port port protocol proto conflict response
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoConflict) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer hosturl hosturl externalipaddress Ip address port port protocol proto conflict response
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoInternalServerErrorCode is the HTTP code returned for type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoInternalServerError
const DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoInternalServerErrorCode int = 500

/*
DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoInternalServerError Internal service error

swagger:response deleteConfigLoadbalancerHosturlHosturlExternalipaddressIpAddressPortPortProtocolProtoInternalServerError
*/
type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoInternalServerError creates DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoInternalServerError with default headers values
func NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoInternalServerError() *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoInternalServerError {

	return &DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoInternalServerError{}
}

// WithPayload adds the payload to the delete config loadbalancer hosturl hosturl externalipaddress Ip address port port protocol proto internal server error response
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoInternalServerError) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer hosturl hosturl externalipaddress Ip address port port protocol proto internal server error response
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailableCode is the HTTP code returned for type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable
const DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailableCode int = 503

/*
DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable Maintenance mode

swagger:response deleteConfigLoadbalancerHosturlHosturlExternalipaddressIpAddressPortPortProtocolProtoServiceUnavailable
*/
type DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable creates DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable with default headers values
func NewDeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable() *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable {

	return &DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable{}
}

// WithPayload adds the payload to the delete config loadbalancer hosturl hosturl externalipaddress Ip address port port protocol proto service unavailable response
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable) WithPayload(payload *models.Error) *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config loadbalancer hosturl hosturl externalipaddress Ip address port port protocol proto service unavailable response
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigLoadbalancerHosturlHosturlExternalipaddressIPAddressPortPortProtocolProtoServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
