package main

import (
	"alexdenkk/auth-api/app"
	"alexdenkk/auth-api/model"
	"alexdenkk/auth-api/pkg/db"
	"log"
	"os"
)

func main() {
	// connecting to db
	appDB, err := db.Connect(
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("db connected")

	// migration
	model.Migrate(appDB)
	log.Println("migration complete")

	// create app
	app := app.New(appDB, []byte(os.Getenv("JWT")), os.Getenv("HOST"))
	log.Println("server initialized")

	// run
	app.Run()
}
