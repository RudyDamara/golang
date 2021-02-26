package handler

import (
	"github.com/RudyDamara/golang/pkg/user/model"
	"github.com/labstack/echo"
)

type HTTPHandler struct {
	model model.UserModel
}

func NewHTTPHandler(model model.UserModel) *HTTPHandler {
	return &HTTPHandler{model: model}
}

func (h *HTTPHandler) Mount(g *echo.Group) {
	g.POST("/login", h.ToLogin)
}
