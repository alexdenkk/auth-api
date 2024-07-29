package main

import (
	"flag"
	"log"
	"mnb/users/app"
	"mnb/users/model"
	"mnb/users/pkg/db"
	"mnb/users/pkg/hash"
)

var (
	addr string
	jwt  string

	dbName string
	dbPort string
	dbUser string
	dbPswd string

	adminLogin    string
	adminPassword string
)

func main() {

	parseFlags()

	// connecting to db
	appDB, err := db.Connect(
		dbName,
		dbPort,
		dbUser,
		dbPswd,
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("db connected")

	// migration
	model.Migrate(appDB)
	log.Println("migration complete")

	// creating admin from server
	if adminLogin != "" && adminPassword != "" {
		result := appDB.Create(&model.User{
			Login:    adminLogin,
			Password: hash.Hash(adminPassword),
		})

		if result.Error != nil {
			log.Fatal(err)
			return
		}

		log.Println("user created")
		return
	}

	// create app
	usersApp := app.New(appDB, []byte(jwt), addr)
	log.Println("server initialized")

	// run
	usersApp.Run()
}

func parseFlags() {
	flag.StringVar(&addr, "address", ":8001", "address and(or) port for app")
	flag.StringVar(&jwt, "jwt", "", "jwt sign key for user tokens")

	flag.StringVar(&dbName, "db-name", "", "database name")
	flag.StringVar(&dbPort, "db-port", "5432", "database port")
	flag.StringVar(&dbUser, "db-user", "", "database user")
	flag.StringVar(&dbPswd, "db-password", "", "database password")

	flag.StringVar(&adminLogin, "admin-login", "", "admin login")
	flag.StringVar(&adminPassword, "admin-pass", "", "admin password")
	flag.Parse()
}
