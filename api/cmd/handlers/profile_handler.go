package handlers

import (
	"api/internal/httputils"
	"api/internal/profile"
	"encoding/json"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"net/http"
)

func ProfileRoutes(routes *router.Router, handler ProfileHandler) *router.Router {
	group := routes.Group("/profile/v1")

	group.Handle(http.MethodPost, "/webhook", handler.Webhook)

	return routes
}

type ProfileHandler interface {
	Webhook(ctx *fasthttp.RequestCtx)
}

type profileHandler struct {
	backend profile.Backend
}

func (r *profileHandler) Webhook(ctx *fasthttp.RequestCtx) {
	var body profile.Webhook
	if err := json.Unmarshal(ctx.Request.Body(), &body); err != nil {
		httputils.JSONError(&ctx.Response, err, http.StatusBadRequest)
		return
	}

	err := r.backend.Webhook(ctx, body)
	if err != nil {
		httputils.BackendErrorFactory(&ctx.Response, err)
		return
	}

	httputils.JSON(&ctx.Response, &httputils.Response{Status: http.StatusOK, Msg: "success"}, http.StatusOK)
}

// TODO CRUD CREATE-USER BFF
//func (r *profileHandler) CreateUser(ctx *fasthttp.RequestCtx) {
//	var body profile.CreateUser
//	if err := json.Unmarshal(ctx.Request.Body(), &body); err != nil {
//		httputils.JSONError(&ctx.Response, err, http.StatusBadRequest)
//		return
//	}
//
//	err := r.backend.CreateUser(ctx, body)
//	if err != nil {
//		httputils.BackendErrorFactory(&ctx.Response, err)
//		return
//	}
//
//	httputils.JSON(&ctx.Response, &httputils.Response{Status: http.StatusOK, Msg: "success"}, http.StatusOK)
//}

func NewProfileHandler(backend profile.Backend) ProfileHandler {
	return &profileHandler{
		backend: backend,
	}
}
