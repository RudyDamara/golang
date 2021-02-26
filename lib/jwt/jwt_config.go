package jwt

import (
	"fmt"
	"io/ioutil"

	"github.com/RudyDamara/golang/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/zerolog/log"
)

type registerRepo struct {
	dbConn *db.DbConnection
}

type JwtCustomClaims struct {
	User User `json:"user"`
	jwt.StandardClaims
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Islogin  bool   `json:"islogin"`
}

type GetUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Islogin  bool   `json:"islogin"`
}

func JWTConfig() middleware.JWTConfig {
	keyBytes, err := ioutil.ReadFile("./certificate/jwtRS256.key.pub")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	publicKey, _ := jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	fmt.Println(&JwtCustomClaims{})
	config := middleware.JWTConfig{
		SigningKey:    publicKey,
		SigningMethod: "RS256",
		Claims:        &JwtCustomClaims{},
	}

	return config
}

func GetTokenDecode(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}

func GetDataUser(c echo.Context) *GetUser {
	jwtUser := c.Get("user").(*jwt.Token)
	jwtClaims := jwtUser.Claims.(*JwtCustomClaims)
	jwtGetUsers := jwtClaims.User

	data := GetUser{
		ID:       jwtGetUsers.ID,
		Username: jwtGetUsers.Username,
		Email:    jwtGetUsers.Email,
		Islogin:  jwtGetUsers.Islogin,
	}
	fmt.Println(data)
	return &data
}

func CheckLogin(c echo.Context) bool {
	jwtUser := c.Get("user").(*jwt.Token)
	jwtClaims := jwtUser.Claims.(*JwtCustomClaims)
	jwtGetUsers := jwtClaims.User

	q := db.CreateConnection().DbData
	getUser := GetUser{}
	qr := "Select * from users where id = ?"
	q.Raw(qr, jwtGetUsers.ID).Scan(&getUser)
	// if err := .Error; err != nil {
	// 	return false
	// }
	if getUser.Islogin == true {
		fmt.Println("checkmail", true)
		return true
	} else {
		return false
	}
}
