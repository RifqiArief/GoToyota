package model

import (
	"fmt"
	"os"

	"github.com/GoToyota/utils"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func Init() error {
	err := godotenv.Load("config/database.env")
	if err != nil {
		utils.Logging.Println("model/base/line:12")
		return err
	}

	username := os.Getenv("db_username")
	password := os.Getenv("db_password")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbType := os.Getenv("db_type")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	utils.Logging.Println(dbUri)

	//open connection
	conn, err := sqlx.Connect(dbType, dbUri)
	if err != nil {
		utils.Logging.Println("model/base/line:34")
		return err
	}

	err = conn.Ping()
	if err != nil {
		utils.Logging.Println("model/base/line:40")
		return err
	}

	utils.Logging.Println("Success connect database")
	db = conn
	return nil
}
