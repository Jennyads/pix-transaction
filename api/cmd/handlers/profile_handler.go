package handlers

import (
	"api/internal/httputils"
	"api/internal/middleware"
	"api/internal/profile"
	"encoding/json"
	"errors"
	"github.com/fasthttp/router"
	"github.com/go-playground/validator/v10"
	"github.com/valyala/fasthttp"
	"net/http"
	"regexp"
)

func ProfileRoutes(routes *router.Router, handler ProfileHandler, middleware middleware.Middleware) *router.Router {
	group := routes.Group("/profile/v1")

	group.Handle(http.MethodPost, "/webhook", middleware.WrapHandler(handler.Webhook))
	group.Handle(http.MethodPost, "/user", middleware.WrapHandler(handler.CreateUser))
	group.Handle(http.MethodPost, "/account", middleware.WrapHandler(handler.CreateAccount))
	group.Handle(http.MethodGet, "/users", middleware.WrapHandler(handler.ListUsers))
	group.Handle(http.MethodGet, "/account", middleware.WrapHandler(handler.FindAccount))
	group.Handle(http.MethodPost, "/pix", middleware.WrapHandler(handler.SendPix))
	group.Handle(http.MethodPost, "/pixWebhook", middleware.WrapHandler(handler.PixWebhook))
	group.Handle(http.MethodPost, "/key", middleware.WrapHandler(handler.CreateKey))

	return routes
}

type ProfileHandler interface {
	Webhook(ctx *fasthttp.RequestCtx)
	CreateUser(ctx *fasthttp.RequestCtx)
	CreateAccount(ctx *fasthttp.RequestCtx)
	ListUsers(ctx *fasthttp.RequestCtx)
	FindAccount(ctx *fasthttp.RequestCtx)
	SendPix(ctx *fasthttp.RequestCtx)
	PixWebhook(ctx *fasthttp.RequestCtx)
	CreateKey(ctx *fasthttp.RequestCtx)
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
	phoneRegex := regexp.MustCompile(`^\+\d{13}$`)
	return phoneRegex.MatchString(phone)
}

func isValidCPF(cpf string) bool {
	cpfRegex := regexp.MustCompile(`^\d{11}$`)
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

	//userId := ctx.UserValue("userId").(string)

	userId := string(ctx.Request.Header.Peek("userId"))
	ctx.SetUserValue("userId", userId)

	if err := json.Unmarshal(ctx.Request.Body(), &body); err != nil {
		httputils.JSONError(&ctx.Response, err, http.StatusBadRequest)
		return
	}
	err := r.backend.CreateAccount(ctx, userId, body)
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

func (r *profileHandler) FindAccount(ctx *fasthttp.RequestCtx) {
	userId := string(ctx.QueryArgs().Peek("userId"))

	if userId == "" {
		httputils.JSONError(&ctx.Response, errors.New("userId cant be empty"), http.StatusBadRequest)
		return
	}

	//account := profile.Account{Name: userId}

	err := r.backend.FindAccount(ctx, userId)
	if err != nil {
		httputils.BackendErrorFactory(&ctx.Response, err)
		return
	}

	httputils.JSON(&ctx.Response, &httputils.Response{Status: http.StatusOK, Msg: "success"}, http.StatusOK)
}

func (r *profileHandler) ListUsers(ctx *fasthttp.RequestCtx) {
	ids := ctx.QueryArgs().PeekMulti("ids")
	if len(ids) == 0 {
		httputils.JSONError(&ctx.Response, errors.New("ids cant be empty"), http.StatusBadRequest)
		return
	}

	var idStrings []string
	for _, id := range ids {
		idStrings = append(idStrings, string(id))
	}

	list, err := r.backend.ListUsers(ctx, idStrings)
	if err != nil {
		httputils.BackendErrorFactory(&ctx.Response, err)
		return
	}
	httputils.JSON(&ctx.Response, list, http.StatusOK)
}

func (r *profileHandler) CreateKey(ctx *fasthttp.RequestCtx) {
	var body profile.Key
	if err := json.Unmarshal(ctx.Request.Body(), &body); err != nil {
		httputils.JSONError(&ctx.Response, err, http.StatusBadRequest)
		return
	}
	err := r.backend.CreateKey(ctx, body)
	if err != nil {
		httputils.BackendErrorFactory(&ctx.Response, err)
		return
	}
	httputils.JSON(&ctx.Response, &httputils.Response{Status: http.StatusOK, Msg: "success"}, http.StatusOK)
}

func NewProfileHandler(backend profile.Backend) ProfileHandler {
	return &profileHandler{
		backend: backend,
	}
}
