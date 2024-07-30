package app

// RegisterEndpoints - app function for registering endpoints
func (app *App) RegisterEndpoints() {
	authEndpoints := app.Engine.Group("/auth")

	authEndpoints.POST("/login/", app.Delivery.Login)
}
