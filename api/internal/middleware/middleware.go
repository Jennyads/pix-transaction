package middleware

import (
	"github.com/valyala/fasthttp"
)

type Middleware interface {
	WrapHandler(next fasthttp.RequestHandler) fasthttp.RequestHandler
}

type Middlewares interface {
	Apply(handler fasthttp.RequestHandler) fasthttp.RequestHandler
}

type middlewaresImpl struct {
	middlewares []Middleware
}

func (m *middlewaresImpl) Apply(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	for i := len(m.middlewares) - 1; i >= 0; i-- {
		handler = m.middlewares[i].WrapHandler(handler)
	}
	return handler
}

func New(middlewares ...Middleware) Middlewares {
	return &middlewaresImpl{
		middlewares: middlewares,
	}
}
