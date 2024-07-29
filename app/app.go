package app

import (
	"context"
	delivery_http "mnb/users/internal/users/delivery/http"
	"os"
	"os/signal"

	"log"
	"mnb/users/pkg/middleware"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// App - server service app struct
type App struct {
	Handler *delivery_http.Handler
	Server  *http.Server
	MW      *middleware.Middleware
	SignKey []byte
	DB      *gorm.DB
}

// Run - function for run service app
func (app *App) Run() {
	app.Route()
	log.Println("================")
	log.Println("=server=running=")
	log.Println("================")
	log.Println("=MNB=elephants==")
	log.Println("================")

	var wait time.Duration = time.Hour * 100

	go func() {
		if err := app.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	app.Server.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}

// New - function for create new service app
func New(db *gorm.DB, key []byte, addr string) *App {
	app := &App{
		DB:      db,
		SignKey: key,
	}

	app.Server = &http.Server{
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	app.MW = middleware.New(app.SignKey)

	app.InitHandler()

	return app
}
