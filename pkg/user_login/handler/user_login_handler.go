package handler

import (
	"github.com/RudyDamara/golang/pkg/user_login/model"
	"github.com/labstack/echo"
)

type HTTPHandler struct {
	model model.UserBalanceModel
}

func NewHTTPHandler(model model.UserBalanceModel) *HTTPHandler {
	return &HTTPHandler{model: model}
}

func (h *HTTPHandler) Mount(g *echo.Group, auth echo.MiddlewareFunc) {
	g.GET("/logout", h.ToLogout, auth)
}
