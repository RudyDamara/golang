package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/zerolog/log"
)

type DbConnection struct {
	DbData *gorm.DB
}

func CreateConnection() *DbConnection {
	//connection db_transaction_data
	log.Debug().Msg(os.Getenv("DRIVER_NAME"))
	fmt.Println("DRIVER_NAME", os.Getenv("DRIVER_NAME"))
	db, err := gorm.Open(os.Getenv("DRIVER_NAME"), os.Getenv("CONNECTION_STRING"))
	if err != nil {
		log.Error().Msg(err.Error())
		log.Info().Msg("failed to connect database db_name")
	}

	num1, _ := strconv.Atoi(os.Getenv("MAX_CONNECTION_POOL"))
	db.DB().SetMaxOpenConns(num1)
	db.LogMode(true)

	return &DbConnection{db}
}
