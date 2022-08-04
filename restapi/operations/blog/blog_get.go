// Code generated by go-swagger; DO NOT EDIT.

package blog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// BlogGetHandlerFunc turns a function with the right signature into a blog get handler
type BlogGetHandlerFunc func(BlogGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BlogGetHandlerFunc) Handle(params BlogGetParams) middleware.Responder {
	return fn(params)
}

// BlogGetHandler interface for that can handle valid blog get params
type BlogGetHandler interface {
	Handle(BlogGetParams) middleware.Responder
}

// NewBlogGet creates a new http.Handler for the blog get operation
func NewBlogGet(ctx *middleware.Context, handler BlogGetHandler) *BlogGet {
	return &BlogGet{Context: ctx, Handler: handler}
}

/* BlogGet swagger:route GET /api/v1/blog/{blog_id} Blog blogGet

get the blog of service

*/
type BlogGet struct {
	Context *middleware.Context
	Handler BlogGetHandler
}

func (o *BlogGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewBlogGetParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
