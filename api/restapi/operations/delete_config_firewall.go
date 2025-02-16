// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteConfigFirewallHandlerFunc turns a function with the right signature into a delete config firewall handler
type DeleteConfigFirewallHandlerFunc func(DeleteConfigFirewallParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteConfigFirewallHandlerFunc) Handle(params DeleteConfigFirewallParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteConfigFirewallHandler interface for that can handle valid delete config firewall params
type DeleteConfigFirewallHandler interface {
	Handle(DeleteConfigFirewallParams, interface{}) middleware.Responder
}

// NewDeleteConfigFirewall creates a new http.Handler for the delete config firewall operation
func NewDeleteConfigFirewall(ctx *middleware.Context, handler DeleteConfigFirewallHandler) *DeleteConfigFirewall {
	return &DeleteConfigFirewall{Context: ctx, Handler: handler}
}

/*
	DeleteConfigFirewall swagger:route DELETE /config/firewall deleteConfigFirewall

# Delete of the firewall service

Delete of the firewall service.
*/
type DeleteConfigFirewall struct {
	Context *middleware.Context
	Handler DeleteConfigFirewallHandler
}

func (o *DeleteConfigFirewall) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteConfigFirewallParams()
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
