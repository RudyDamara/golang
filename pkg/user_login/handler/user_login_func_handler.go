package handler

import (
	"net/http"

	jwtConfig "github.com/RudyDamara/golang/lib/jwt"

	"github.com/RudyDamara/golang/lib/models"
	"github.com/RudyDamara/golang/pkg/user_login/structs"
	"github.com/labstack/echo"
)

func (h *HTTPHandler) ToLogout(c echo.Context) error {

	dataUser := jwtConfig.GetDataUser(c)

	model := structs.User{
		ID: dataUser.ID,
	}

	result := <-h.model.Logout(model)

	if result.Error != nil {
		resp := &models.Response{Code: 400, MessageCode: 0, Message: result.Error.Error()}
		return c.JSON(http.StatusBadRequest, resp)
	}

	return c.JSON(http.StatusOK, result.Data)
}
