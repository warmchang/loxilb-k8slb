// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNoContentCode is the HTTP code returned for type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNoContent
const DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNoContentCode int = 204

/*
DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNoContent OK

swagger:response deleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNoContent
*/
type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNoContent struct {
}

// NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNoContent creates DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNoContent with default headers values
func NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNoContent() *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNoContent {

	return &DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNoContent{}
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameBadRequestCode is the HTTP code returned for type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameBadRequest
const DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameBadRequestCode int = 400

/*
DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameBadRequest Malformed arguments for API call

swagger:response deleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameBadRequest
*/
type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameBadRequest creates DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameBadRequest with default headers values
func NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameBadRequest() *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameBadRequest {

	return &DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameBadRequest{}
}

// WithPayload adds the payload to the delete config bgp policy definedsets defineset type type name bad request response
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameBadRequest) WithPayload(payload *models.Error) *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bgp policy definedsets defineset type type name bad request response
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameUnauthorizedCode is the HTTP code returned for type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameUnauthorized
const DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameUnauthorizedCode int = 401

/*
DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameUnauthorized Invalid authentication credentials

swagger:response deleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameUnauthorized
*/
type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameUnauthorized creates DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameUnauthorized with default headers values
func NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameUnauthorized() *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameUnauthorized {

	return &DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameUnauthorized{}
}

// WithPayload adds the payload to the delete config bgp policy definedsets defineset type type name unauthorized response
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameUnauthorized) WithPayload(payload *models.Error) *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bgp policy definedsets defineset type type name unauthorized response
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameForbiddenCode is the HTTP code returned for type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameForbidden
const DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameForbiddenCode int = 403

/*
DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameForbidden Capacity insufficient

swagger:response deleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameForbidden
*/
type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameForbidden creates DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameForbidden with default headers values
func NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameForbidden() *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameForbidden {

	return &DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameForbidden{}
}

// WithPayload adds the payload to the delete config bgp policy definedsets defineset type type name forbidden response
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameForbidden) WithPayload(payload *models.Error) *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bgp policy definedsets defineset type type name forbidden response
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNotFoundCode is the HTTP code returned for type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNotFound
const DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNotFoundCode int = 404

/*
DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNotFound Resource not found

swagger:response deleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNotFound
*/
type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNotFound creates DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNotFound with default headers values
func NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNotFound() *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNotFound {

	return &DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNotFound{}
}

// WithPayload adds the payload to the delete config bgp policy definedsets defineset type type name not found response
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNotFound) WithPayload(payload *models.Error) *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bgp policy definedsets defineset type type name not found response
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameConflictCode is the HTTP code returned for type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameConflict
const DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameConflictCode int = 409

/*
DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameConflict Resource Conflict. Neigh already exists

swagger:response deleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameConflict
*/
type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameConflict creates DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameConflict with default headers values
func NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameConflict() *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameConflict {

	return &DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameConflict{}
}

// WithPayload adds the payload to the delete config bgp policy definedsets defineset type type name conflict response
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameConflict) WithPayload(payload *models.Error) *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bgp policy definedsets defineset type type name conflict response
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameInternalServerErrorCode is the HTTP code returned for type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameInternalServerError
const DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameInternalServerErrorCode int = 500

/*
DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameInternalServerError Internal service error

swagger:response deleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameInternalServerError
*/
type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameInternalServerError creates DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameInternalServerError with default headers values
func NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameInternalServerError() *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameInternalServerError {

	return &DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameInternalServerError{}
}

// WithPayload adds the payload to the delete config bgp policy definedsets defineset type type name internal server error response
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameInternalServerError) WithPayload(payload *models.Error) *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bgp policy definedsets defineset type type name internal server error response
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameServiceUnavailableCode is the HTTP code returned for type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameServiceUnavailable
const DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameServiceUnavailableCode int = 503

/*
DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameServiceUnavailable Maintenance mode

swagger:response deleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameServiceUnavailable
*/
type DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameServiceUnavailable creates DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameServiceUnavailable with default headers values
func NewDeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameServiceUnavailable() *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameServiceUnavailable {

	return &DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameServiceUnavailable{}
}

// WithPayload adds the payload to the delete config bgp policy definedsets defineset type type name service unavailable response
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameServiceUnavailable) WithPayload(payload *models.Error) *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bgp policy definedsets defineset type type name service unavailable response
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
