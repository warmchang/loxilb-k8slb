// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetConfigParamsHandlerFunc turns a function with the right signature into a get config params handler
type GetConfigParamsHandlerFunc func(GetConfigParamsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConfigParamsHandlerFunc) Handle(params GetConfigParamsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetConfigParamsHandler interface for that can handle valid get config params params
type GetConfigParamsHandler interface {
	Handle(GetConfigParamsParams, interface{}) middleware.Responder
}

// NewGetConfigParams creates a new http.Handler for the get config params operation
func NewGetConfigParams(ctx *middleware.Context, handler GetConfigParamsHandler) *GetConfigParams {
	return &GetConfigParams{Context: ctx, Handler: handler}
}

/*
	GetConfigParams swagger:route GET /config/params getConfigParams

# Get Operational params of LoxiLB

Get Operational params of LoxiLB
*/
type GetConfigParams struct {
	Context *middleware.Context
	Handler GetConfigParamsHandler
}

func (o *GetConfigParams) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetConfigParamsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
