package model

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/RudyDamara/golang/db"
	"github.com/RudyDamara/golang/lib/models"
	"github.com/RudyDamara/golang/pkg/user/structs"
	"github.com/dgrijalva/jwt-go"
)

type userModel struct {
	dbConn *db.DbConnection
}

func NewUserModel(dbConn *db.DbConnection) UserModel {
	return &userModel{dbConn: dbConn}
}

func (r *userModel) Login(param structs.ReqLogin) chan models.Result {
	output := make(chan models.Result)
	var users structs.User
	var users2 structs.User

	go func() {
		defer close(output)

		p := param
		q := r.dbConn.DbData

		qr := "SELECT id, username, email, islogin " +
			"FROM users where username = '" + p.Username +
			"' and password = '" + p.Password + "' "
		// "UPDATE users SET islogin == true where "
		e := q.Raw(qr).Scan(&users).Error
		if e != nil {
			output <- models.Result{Error: e}
			return
		}
		if users.ID != 0 {

			query := "UPDATE users SET islogin = true where id = ? RETURNING *"
			q.Raw(query, users.ID).Scan(&users2)
			if users2.Islogin == true {

				getUsers := structs.User{
					ID:       users.ID,
					Username: users.Username,
					Email:    users.Email,
					Islogin:  true,
				}

				token := CreateToken(getUsers)

				accessToken := structs.AccessToken{
					Type:  "bearer",
					Token: token,
				}

				result := structs.ResultLogin{
					User:        users,
					AccessToken: accessToken,
				}

				resp := &models.Response{Code: 200, MessageCode: 0, Message: "Success", Data: result}
				output <- models.Result{Data: resp}

			}
		}

	}()

	return output
}

func CreateToken(user structs.User) (token string) {
	var signKey *rsa.PrivateKey
	signBytes := []byte(os.Getenv("CERTIFICATE"))

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		fmt.Println("err02:", err)
	}

	// Create token
	tok := jwt.New(jwt.GetSigningMethod("RS256"))

	// Set claims
	claims := tok.Claims.(jwt.MapClaims)
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()

	// Generate encoded token and send it as response.
	token, _ = tok.SignedString(signKey)
	return token
}
