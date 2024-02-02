package handlers

import (
	"api/internal/httputils"
	"api/internal/profile"
	"encoding/json"
	"errors"
	"github.com/fasthttp/router"
	"github.com/go-playground/validator/v10"
	"github.com/valyala/fasthttp"
	"net/http"
	"regexp"
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

func ValidateCreateUser(user *profile.User) error {
	validate := validator.New()
	if err := validate.RegisterValidation("validateData", validateData); err != nil {
		return errors.New("internal server error")
	}
	err := validate.Struct(user)
	if err != nil {
		return errors.New("invalid user data: " + err.Error())
	}

	return nil
}
func validateData(fl validator.FieldLevel) bool {
	data := fl.Field().String()
	switch {
	case isValidEmail(data), isValidPhoneNumber(data), isValidCPF(data):
		return true
	default:
		return false
	}
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func isValidPhoneNumber(phone string) bool {
	phoneRegex := regexp.MustCompile(`^\+[0-9]+$`)
	return phoneRegex.MatchString(phone)
}

func isValidCPF(cpf string) bool {
	cpfRegex := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`)
	return cpfRegex.MatchString(cpf)
}

func (r *profileHandler) CreateUser(ctx *fasthttp.RequestCtx) {
	var body profile.User
	if err := json.Unmarshal(ctx.Request.Body(), &body); err != nil {
		httputils.JSONError(&ctx.Response, err, http.StatusBadRequest)
		return
	}

	if err := ValidateCreateUser(&body); err != nil {
		httputils.JSONError(&ctx.Response, err, http.StatusBadRequest)
		return
	}

	err := r.backend.CreateUser(ctx, body)
	if err != nil {
		httputils.BackendErrorFactory(&ctx.Response, err)
		return
	}

	httputils.JSON(&ctx.Response, &httputils.Response{Status: http.StatusOK, Msg: "success"}, http.StatusOK)
}

func (r *profileHandler) CreateAccount(ctx *fasthttp.RequestCtx) {
	var body profile.Account
	if err := json.Unmarshal(ctx.Request.Body(), &body); err != nil {
		httputils.JSONError(&ctx.Response, err, http.StatusBadRequest)
		return
	}
	err := r.backend.CreateAccount(ctx, body)
	if err != nil {
		httputils.BackendErrorFactory(&ctx.Response, err)
		return
	}
	httputils.JSON(&ctx.Response, &httputils.Response{Status: http.StatusOK, Msg: "success"}, http.StatusOK)
}

func (r *profileHandler) FindAccount(ctx *fasthttp.RequestCtx) {
	var body profile.Account
	if err := json.Unmarshal(ctx.Request.Body(), &body); err != nil {
		httputils.JSONError(&ctx.Response, err, http.StatusBadRequest)
		return
	}
	err := r.backend.FindAccount(ctx, body)
	if err != nil {
		httputils.BackendErrorFactory(&ctx.Response, err)
		return
	}
	httputils.JSON(&ctx.Response, &httputils.Response{Status: http.StatusOK, Msg: "success"}, http.StatusOK)
}

func (r *profileHandler) PixWebhook(ctx *fasthttp.RequestCtx) {
	var body profile.Webhook
	if err := json.Unmarshal(ctx.Request.Body(), &body); err != nil {
		httputils.JSONError(&ctx.Response, err, http.StatusBadRequest)
		return
	}
	err := r.backend.PixWebhook(ctx, body)
	if err != nil {
		httputils.BackendErrorFactory(&ctx.Response, err)
		return
	}
	httputils.JSON(&ctx.Response, &httputils.Response{Status: http.StatusOK, Msg: "success"}, http.StatusOK)
}

func (r *profileHandler) SendPix(ctx *fasthttp.RequestCtx) {
	var body profile.PixTransaction
	if err := json.Unmarshal(ctx.Request.Body(), &body); err != nil {
		httputils.JSONError(&ctx.Response, err, http.StatusBadRequest)
		return
	}
	err := r.backend.SendPix(ctx, body)
	if err != nil {
		httputils.BackendErrorFactory(&ctx.Response, err)
		return
	}
	httputils.JSON(&ctx.Response, &httputils.Response{Status: http.StatusOK, Msg: "success"}, http.StatusOK)
}

func (r *profileHandler) ListUsers(ctx *fasthttp.RequestCtx) {
	var body []string

	if err := json.Unmarshal(ctx.Request., &body); err != nil {
		httputils.JSONError(&ctx.Response, err, http.StatusBadRequest)
		return
	}
	list, err := r.backend.ListUsers(ctx,)
	if err != nil {
		httputils.BackendErrorFactory(&ctx.Response, err)
		return

	}
}
func NewProfileHandler(backend profile.Backend) ProfileHandler {
	return &profileHandler{
		backend: backend,
	}
}
