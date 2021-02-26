package model

import (
	"github.com/RudyDamara/golang/lib/models"
	"github.com/RudyDamara/golang/pkg/user/structs"
)

type UserModel interface {
	Login(structs.ReqLogin) chan models.Result
}
