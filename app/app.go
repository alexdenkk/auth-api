package app

import (
	delivery "alexdenkk/auth-api/internal/auth/delivery/http"
	"alexdenkk/auth-api/internal/auth/repository"
	"alexdenkk/auth-api/internal/auth/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// App - service app struct
type App struct {
	Delivery *delivery.Delivery
	Engine   *gin.Engine
	Host     string
}

// Run - function for running service app
func (app *App) Run() {
	app.Engine.Run(app.Host)
}

// New - function for creating App instance
func New(db *gorm.DB, key []byte, host string) *App {
	app := &App{
		Engine: gin.Default(),
		Host:   host,
	}

	repo := repository.New(db)
	srv := service.New(repo, key)
	dlv := delivery.New(srv)

	app.Delivery = dlv

	app.RegisterEndpoints()

	return app
}
