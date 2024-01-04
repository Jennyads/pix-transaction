package httputils

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func JSON(response *fasthttp.Response, body any, status int) {
	b, _ := json.Marshal(body)
	response.Header.Set(fasthttp.HeaderContentType, "application/json")
	response.SetBody(b)
	response.SetStatusCode(status)
}

func JSONError(response *fasthttp.Response, err error, status int) {
	JSON(response, &Response{Msg: err.Error(), Status: status}, status)
}
