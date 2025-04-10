// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// PostConfigSessionulclNoContentCode is the HTTP code returned for type PostConfigSessionulclNoContent
const PostConfigSessionulclNoContentCode int = 204

/*
PostConfigSessionulclNoContent OK

swagger:response postConfigSessionulclNoContent
*/
type PostConfigSessionulclNoContent struct {
}

// NewPostConfigSessionulclNoContent creates PostConfigSessionulclNoContent with default headers values
func NewPostConfigSessionulclNoContent() *PostConfigSessionulclNoContent {

	return &PostConfigSessionulclNoContent{}
}

// WriteResponse to the client
func (o *PostConfigSessionulclNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostConfigSessionulclBadRequestCode is the HTTP code returned for type PostConfigSessionulclBadRequest
const PostConfigSessionulclBadRequestCode int = 400

/*
PostConfigSessionulclBadRequest Malformed arguments for API call

swagger:response postConfigSessionulclBadRequest
*/
type PostConfigSessionulclBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigSessionulclBadRequest creates PostConfigSessionulclBadRequest with default headers values
func NewPostConfigSessionulclBadRequest() *PostConfigSessionulclBadRequest {

	return &PostConfigSessionulclBadRequest{}
}

// WithPayload adds the payload to the post config sessionulcl bad request response
func (o *PostConfigSessionulclBadRequest) WithPayload(payload *models.Error) *PostConfigSessionulclBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config sessionulcl bad request response
func (o *PostConfigSessionulclBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigSessionulclBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigSessionulclUnauthorizedCode is the HTTP code returned for type PostConfigSessionulclUnauthorized
const PostConfigSessionulclUnauthorizedCode int = 401

/*
PostConfigSessionulclUnauthorized Invalid authentication credentials

swagger:response postConfigSessionulclUnauthorized
*/
type PostConfigSessionulclUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigSessionulclUnauthorized creates PostConfigSessionulclUnauthorized with default headers values
func NewPostConfigSessionulclUnauthorized() *PostConfigSessionulclUnauthorized {

	return &PostConfigSessionulclUnauthorized{}
}

// WithPayload adds the payload to the post config sessionulcl unauthorized response
func (o *PostConfigSessionulclUnauthorized) WithPayload(payload *models.Error) *PostConfigSessionulclUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config sessionulcl unauthorized response
func (o *PostConfigSessionulclUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigSessionulclUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigSessionulclForbiddenCode is the HTTP code returned for type PostConfigSessionulclForbidden
const PostConfigSessionulclForbiddenCode int = 403

/*
PostConfigSessionulclForbidden Capacity insufficient

swagger:response postConfigSessionulclForbidden
*/
type PostConfigSessionulclForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigSessionulclForbidden creates PostConfigSessionulclForbidden with default headers values
func NewPostConfigSessionulclForbidden() *PostConfigSessionulclForbidden {

	return &PostConfigSessionulclForbidden{}
}

// WithPayload adds the payload to the post config sessionulcl forbidden response
func (o *PostConfigSessionulclForbidden) WithPayload(payload *models.Error) *PostConfigSessionulclForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config sessionulcl forbidden response
func (o *PostConfigSessionulclForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigSessionulclForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigSessionulclNotFoundCode is the HTTP code returned for type PostConfigSessionulclNotFound
const PostConfigSessionulclNotFoundCode int = 404

/*
PostConfigSessionulclNotFound Resource not found

swagger:response postConfigSessionulclNotFound
*/
type PostConfigSessionulclNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigSessionulclNotFound creates PostConfigSessionulclNotFound with default headers values
func NewPostConfigSessionulclNotFound() *PostConfigSessionulclNotFound {

	return &PostConfigSessionulclNotFound{}
}

// WithPayload adds the payload to the post config sessionulcl not found response
func (o *PostConfigSessionulclNotFound) WithPayload(payload *models.Error) *PostConfigSessionulclNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config sessionulcl not found response
func (o *PostConfigSessionulclNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigSessionulclNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigSessionulclConflictCode is the HTTP code returned for type PostConfigSessionulclConflict
const PostConfigSessionulclConflictCode int = 409

/*
PostConfigSessionulclConflict Resource Conflict. VLAN already exists OR dependency VRF/VNET not found

swagger:response postConfigSessionulclConflict
*/
type PostConfigSessionulclConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigSessionulclConflict creates PostConfigSessionulclConflict with default headers values
func NewPostConfigSessionulclConflict() *PostConfigSessionulclConflict {

	return &PostConfigSessionulclConflict{}
}

// WithPayload adds the payload to the post config sessionulcl conflict response
func (o *PostConfigSessionulclConflict) WithPayload(payload *models.Error) *PostConfigSessionulclConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config sessionulcl conflict response
func (o *PostConfigSessionulclConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigSessionulclConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigSessionulclInternalServerErrorCode is the HTTP code returned for type PostConfigSessionulclInternalServerError
const PostConfigSessionulclInternalServerErrorCode int = 500

/*
PostConfigSessionulclInternalServerError Internal service error

swagger:response postConfigSessionulclInternalServerError
*/
type PostConfigSessionulclInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigSessionulclInternalServerError creates PostConfigSessionulclInternalServerError with default headers values
func NewPostConfigSessionulclInternalServerError() *PostConfigSessionulclInternalServerError {

	return &PostConfigSessionulclInternalServerError{}
}

// WithPayload adds the payload to the post config sessionulcl internal server error response
func (o *PostConfigSessionulclInternalServerError) WithPayload(payload *models.Error) *PostConfigSessionulclInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config sessionulcl internal server error response
func (o *PostConfigSessionulclInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigSessionulclInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigSessionulclServiceUnavailableCode is the HTTP code returned for type PostConfigSessionulclServiceUnavailable
const PostConfigSessionulclServiceUnavailableCode int = 503

/*
PostConfigSessionulclServiceUnavailable Maintenance mode

swagger:response postConfigSessionulclServiceUnavailable
*/
type PostConfigSessionulclServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigSessionulclServiceUnavailable creates PostConfigSessionulclServiceUnavailable with default headers values
func NewPostConfigSessionulclServiceUnavailable() *PostConfigSessionulclServiceUnavailable {

	return &PostConfigSessionulclServiceUnavailable{}
}

// WithPayload adds the payload to the post config sessionulcl service unavailable response
func (o *PostConfigSessionulclServiceUnavailable) WithPayload(payload *models.Error) *PostConfigSessionulclServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config sessionulcl service unavailable response
func (o *PostConfigSessionulclServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigSessionulclServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
