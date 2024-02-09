package middleware

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

type logger struct {
	enabled bool
}

func (l *logger) WrapHandler(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	if !l.enabled {
		return next
	}

	colorText := "\033[34m"
	return func(ctx *fasthttp.RequestCtx) {

		startTime := time.Now()

		next(ctx)

		responseTime := time.Since(startTime)

		fmt.Printf(
			"%s[%s] %s | Status: %d | Request Time: %s | Response Time: %s\n",
			colorText,
			ctx.Method(),
			ctx.Path(),
			ctx.Response.StatusCode(),
			startTime.Format("2006-01-02 15:04:05"),
			responseTime.String(),
		)
	}
}

func NewLogger(enabled bool) Middleware {
	return &logger{
		enabled: enabled,
	}
}
