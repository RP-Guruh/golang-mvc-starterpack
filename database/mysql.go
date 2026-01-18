package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {

	// Get value configuration database from .env
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	// make data source name for connection to mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)

	// make connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Failed connect to MYSQL. Error : %v", err.Error())
	}

	// lifetime for database
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
