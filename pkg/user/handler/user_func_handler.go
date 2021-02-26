package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/RudyDamara/golang/lib/models"
	"github.com/RudyDamara/golang/pkg/user/structs"
	"github.com/labstack/echo"
)

func (h *HTTPHandler) ToLogin(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}
	t := structs.ReqLogin{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}
	model := structs.ReqLogin{
		Username: t.Username,
		Password: t.Password,
	}

	result := <-h.model.Login(model)

	if result.Error != nil {
		resp := &models.Response{Code: 400, MessageCode: 0, Message: result.Error.Error()}
		return c.JSON(http.StatusBadRequest, resp)
	}

	return c.JSON(http.StatusOK, result.Data)
}
