package app

import (
	"mnb/users/pkg/middleware"
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

// Route - function for routing
func (app *App) Route() {
	r := mux.NewRouter()

	r.Use(middleware.LoggerMW)

	r.PathPrefix("/static/").Handler(
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("./web/static/")),
		),
	)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	r.HandleFunc("/login/", app.Handler.Login).
		Methods(http.MethodPost)

	app.Server.Handler = handler
}
