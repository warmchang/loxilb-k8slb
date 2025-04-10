// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetMetricsEpdisttrafficOKCode is the HTTP code returned for type GetMetricsEpdisttrafficOK
const GetMetricsEpdisttrafficOKCode int = 200

/*
GetMetricsEpdisttrafficOK OK

swagger:response getMetricsEpdisttrafficOK
*/
type GetMetricsEpdisttrafficOK struct {

	/*
	  In: Body
	*/
	Payload models.EpDistTrafficMetrics `json:"body,omitempty"`
}

// NewGetMetricsEpdisttrafficOK creates GetMetricsEpdisttrafficOK with default headers values
func NewGetMetricsEpdisttrafficOK() *GetMetricsEpdisttrafficOK {

	return &GetMetricsEpdisttrafficOK{}
}

// WithPayload adds the payload to the get metrics epdisttraffic o k response
func (o *GetMetricsEpdisttrafficOK) WithPayload(payload models.EpDistTrafficMetrics) *GetMetricsEpdisttrafficOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get metrics epdisttraffic o k response
func (o *GetMetricsEpdisttrafficOK) SetPayload(payload models.EpDistTrafficMetrics) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMetricsEpdisttrafficOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty map
		payload = models.EpDistTrafficMetrics{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetMetricsEpdisttrafficInternalServerErrorCode is the HTTP code returned for type GetMetricsEpdisttrafficInternalServerError
const GetMetricsEpdisttrafficInternalServerErrorCode int = 500

/*
GetMetricsEpdisttrafficInternalServerError Internal service error

swagger:response getMetricsEpdisttrafficInternalServerError
*/
type GetMetricsEpdisttrafficInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMetricsEpdisttrafficInternalServerError creates GetMetricsEpdisttrafficInternalServerError with default headers values
func NewGetMetricsEpdisttrafficInternalServerError() *GetMetricsEpdisttrafficInternalServerError {

	return &GetMetricsEpdisttrafficInternalServerError{}
}

// WithPayload adds the payload to the get metrics epdisttraffic internal server error response
func (o *GetMetricsEpdisttrafficInternalServerError) WithPayload(payload *models.Error) *GetMetricsEpdisttrafficInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get metrics epdisttraffic internal server error response
func (o *GetMetricsEpdisttrafficInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMetricsEpdisttrafficInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
