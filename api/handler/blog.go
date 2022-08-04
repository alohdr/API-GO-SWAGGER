package handler

import (
	"LEARN_API_SPEC_SWAGGER/api/restapi/operations/blog"
	"github.com/go-openapi/runtime/middleware"
)

type Blog struct {
	service service.
}

func (b Blog) BlogCreateHandler() blog.BlogCreateHandlerFunc {
	return func(params blog.BlogCreateParams) middleware.Responder {
		return middleware.NotImplemented("operation blog.BlogCreate has not yet been implemented")
	}
}

func (b Blog) BlogGetHandler() blog.BlogGetHandlerFunc {
	return func(params blog.BlogGetParams) middleware.Responder {
		return middleware.NotImplemented("operation blog.BlogGet has not yet been implemented")
	}
}
