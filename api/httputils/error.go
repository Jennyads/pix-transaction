package httputils

import (
	"github.com/valyala/fasthttp"
	"net/http"
)

func IsBackendError(err error) bool {
	_, ok := err.(*backendError)
	return ok
}

type backendError struct {
	msg string
}

func (e *backendError) Error() string {
	return e.msg
}

var (
	ErrMessageNotFound    = backendError{msg: "not found"}
	ErrMessageInternal    = backendError{msg: "internal server error"}
	ErrMessageUnavailable = backendError{msg: "service unavailable"}
	ErrMessageTimeout     = backendError{msg: "gateway timeout"}
)

func BackendErrorFactory(ctx *fasthttp.Response, err error) {
	if IsBackendError(err) {
		switch err {
		case &ErrMessageUnavailable:
			JSONError(ctx, err, http.StatusServiceUnavailable)
			return
		case &ErrMessageNotFound:
			JSONError(ctx, err, http.StatusNotFound)
			return
		case &ErrMessageInternal:
			JSONError(ctx, err, http.StatusInternalServerError)
			return
		case &ErrMessageTimeout:
			JSONError(ctx, err, http.StatusGatewayTimeout)
			return
		}
	}
	JSONError(ctx, err, http.StatusInternalServerError)
	return
}
