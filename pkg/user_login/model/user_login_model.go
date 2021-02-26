package model

import (
	"github.com/RudyDamara/golang/lib/models"
	"github.com/RudyDamara/golang/pkg/user_login/structs"
)

type UserLoginModel interface {
	Logout(structs.User) chan models.Result
}
