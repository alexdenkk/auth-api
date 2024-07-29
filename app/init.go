package app

import (
	delivery_http "mnb/users/internal/users/delivery/http"
	"mnb/users/internal/users/repository"
	"mnb/users/internal/users/service"
)

// InitHandler - function for initializing handler
func (app *App) InitHandler() {
	// repository
	repository := repository.New(app.DB)
	// service
	service := service.New(repository, app.SignKey)
	// handler
	h := delivery_http.New(service, app.SignKey)
	app.Handler = h
}
