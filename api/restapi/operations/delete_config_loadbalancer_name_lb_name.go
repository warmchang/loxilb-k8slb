// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteConfigLoadbalancerNameLbNameHandlerFunc turns a function with the right signature into a delete config loadbalancer name lb name handler
type DeleteConfigLoadbalancerNameLbNameHandlerFunc func(DeleteConfigLoadbalancerNameLbNameParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteConfigLoadbalancerNameLbNameHandlerFunc) Handle(params DeleteConfigLoadbalancerNameLbNameParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteConfigLoadbalancerNameLbNameHandler interface for that can handle valid delete config loadbalancer name lb name params
type DeleteConfigLoadbalancerNameLbNameHandler interface {
	Handle(DeleteConfigLoadbalancerNameLbNameParams, interface{}) middleware.Responder
}

// NewDeleteConfigLoadbalancerNameLbName creates a new http.Handler for the delete config loadbalancer name lb name operation
func NewDeleteConfigLoadbalancerNameLbName(ctx *middleware.Context, handler DeleteConfigLoadbalancerNameLbNameHandler) *DeleteConfigLoadbalancerNameLbName {
	return &DeleteConfigLoadbalancerNameLbName{Context: ctx, Handler: handler}
}

/*
	DeleteConfigLoadbalancerNameLbName swagger:route DELETE /config/loadbalancer/name/{lb_name} deleteConfigLoadbalancerNameLbName

# Delete an existing Load balancer service

Delete an existing load balancer service with name.
*/
type DeleteConfigLoadbalancerNameLbName struct {
	Context *middleware.Context
	Handler DeleteConfigLoadbalancerNameLbNameHandler
}

func (o *DeleteConfigLoadbalancerNameLbName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteConfigLoadbalancerNameLbNameParams()
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
