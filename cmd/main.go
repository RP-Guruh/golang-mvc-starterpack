package main

import (
	"fmt"
	"golang_mvc_starterpack/controllers"
	"golang_mvc_starterpack/database"
	"golang_mvc_starterpack/repositories"
	"golang_mvc_starterpack/routes"
	"golang_mvc_starterpack/services"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// load env
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// make string for custom port running
	port := fmt.Sprintf(":%s", os.Getenv("PORT_RUNNING"))

	// call database connection, from database/mysql.go
	db := database.Connection()
	defer db.Close()

	repo := repositories.NewPeopleRepository(db)
	svc := services.NewPeopleService(repo)
	h := controllers.NewPeopleHandler(svc)

	r := routes.RegisterRouteApi(h)

	r.Run(port)

}
