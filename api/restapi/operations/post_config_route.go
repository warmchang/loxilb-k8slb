// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostConfigRouteHandlerFunc turns a function with the right signature into a post config route handler
type PostConfigRouteHandlerFunc func(PostConfigRouteParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostConfigRouteHandlerFunc) Handle(params PostConfigRouteParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostConfigRouteHandler interface for that can handle valid post config route params
type PostConfigRouteHandler interface {
	Handle(PostConfigRouteParams, interface{}) middleware.Responder
}

// NewPostConfigRoute creates a new http.Handler for the post config route operation
func NewPostConfigRoute(ctx *middleware.Context, handler PostConfigRouteHandler) *PostConfigRoute {
	return &PostConfigRoute{Context: ctx, Handler: handler}
}

/*
	PostConfigRoute swagger:route POST /config/route postConfigRoute

# Create a new route config

Create a new route config .
*/
type PostConfigRoute struct {
	Context *middleware.Context
	Handler PostConfigRouteHandler
}

func (o *PostConfigRoute) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostConfigRouteParams()
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
