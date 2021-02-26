package main

import (
	"fmt"
	"net/http"
	"os"

	jwtConfig "github.com/RudyDamara/golang/lib/jwt"
	config "github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/zerolog/log"

	"github.com/RudyDamara/golang/db"

	userHandler "github.com/RudyDamara/golang/pkg/user/handler"
	userModel "github.com/RudyDamara/golang/pkg/user/model"

	userBalanceHandler "github.com/RudyDamara/golang/pkg/user_login/handler"
	userBalanceModel "github.com/RudyDamara/golang/pkg/user_login/model"
)

func main() {

	if err := config.Load(".env"); err != nil {
		fmt.Println(".env is not loaded properly")
		fmt.Println(err)
		os.Exit(2)
	}

	dbConn := db.CreateConnection()
	authMiddleware := middleware.JWTWithConfig(jwtConfig.JWTConfig())

	r := echo.New()
	r.Debug = true
	r.Use(middleware.Recover())
	r.Use(middleware.Logger())
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	r.GET("/", func(context echo.Context) error {
		return context.HTML(http.StatusOK, "<strong>Test API</strong>")
	})

	apiV1 := r.Group("/api")

	userModel := userModel.NewUserModel(dbConn)
	userHandler.NewHTTPHandler(userModel).Mount(apiV1)

	userBalanceModel := userBalanceModel.NewUserBalanceModel(dbConn)
	userBalanceHandler.NewHTTPHandler(userBalanceModel).Mount(apiV1, authMiddleware)

	err := r.Start(":" + os.Getenv("PORT"))
	if err != nil {
		log.Error().Msg(err.Error())
	}
}
