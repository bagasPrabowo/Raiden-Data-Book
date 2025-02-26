package controllers

import (
	"github.com/sev-2/raiden"
)

type HelloWordRequest struct {
	Name string `path:"name"`
}

type HelloWordResponse struct {
	Message string `json:"message"`
}

type HelloWordController struct {
	raiden.ControllerBase
	Http	string `path:"/hello/{name}" type:"custom"`
	Payload	*HelloWordRequest
	Result	HelloWordResponse
}

func (c *HelloWordController) Get(ctx raiden.Context) error {
	c.Result.Message = "Welcome to raiden " + c.Payload.Name
	return ctx.SendJson(c.Result)
}
