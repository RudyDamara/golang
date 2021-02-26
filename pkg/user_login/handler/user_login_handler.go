package handler

import (
	"github.com/RudyDamara/golang/pkg/user_login/model"
	"github.com/labstack/echo"
)

type HTTPHandler struct {
	model model.UserLoginModel
}

func NewHTTPHandler(model model.UserLoginModel) *HTTPHandler {
	return &HTTPHandler{model: model}
}

func (h *HTTPHandler) Mount(g *echo.Group, auth echo.MiddlewareFunc) {
	g.GET("/logout", h.ToLogout, auth)
}
