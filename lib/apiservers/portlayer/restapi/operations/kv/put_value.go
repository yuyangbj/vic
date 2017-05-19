package kv

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutValueHandlerFunc turns a function with the right signature into a put value handler
type PutValueHandlerFunc func(PutValueParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutValueHandlerFunc) Handle(params PutValueParams) middleware.Responder {
	return fn(params)
}

// PutValueHandler interface for that can handle valid put value params
type PutValueHandler interface {
	Handle(PutValueParams) middleware.Responder
}

// NewPutValue creates a new http.Handler for the put value operation
func NewPutValue(ctx *middleware.Context, handler PutValueHandler) *PutValue {
	return &PutValue{Context: ctx, Handler: handler}
}

/*PutValue swagger:route PUT /kv/{key} kv putValue

Adds / updates value in k/v store

*/
type PutValue struct {
	Context *middleware.Context
	Handler PutValueHandler
}

func (o *PutValue) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPutValueParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}