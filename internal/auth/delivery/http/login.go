package http

import (
	"alexdenkk/auth-api/internal/auth/entity"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login - delivery layer login function
func (d *Delivery) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var request entity.LoginRequest

	err := json.NewDecoder(c.Request.Body).Decode(&request)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	response := d.Service.Login(ctx, request)

	if response.Err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": response.Err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"token": response.Token},
	)
}
