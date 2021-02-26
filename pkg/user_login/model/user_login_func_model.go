package model

import (
	"errors"

	"github.com/RudyDamara/golang/db"
	"github.com/RudyDamara/golang/lib/models"
	"github.com/RudyDamara/golang/pkg/user_login/structs"
)

type userLoginModel struct {
	dbConn *db.DbConnection
}

func NewUserLoginModel(dbConn *db.DbConnection) UserLoginModel {
	return &userLoginModel{dbConn: dbConn}
}

var ErrRecordNotFound = errors.New("record not found")

func (r *userLoginModel) Logout(param structs.User) chan models.Result {
	output := make(chan models.Result)
	var users structs.User

	go func() {
		defer close(output)

		p := param
		q := r.dbConn.DbData

		query := "UPDATE users SET islogin = false where id = ? RETURNING *"
		e := q.Raw(query, p.ID).Scan(&users).Error
		if e != nil {
			output <- models.Result{Error: e}
			return
		}

		if users.Islogin == false {
			resp := &models.Response{Code: 200, MessageCode: 10, Message: "Success Logout"}
			output <- models.Result{Data: resp}

		}

	}()

	return output
}
