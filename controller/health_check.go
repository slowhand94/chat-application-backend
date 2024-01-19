package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 * health check API
 */
func HealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, map[string]any{
		"message": "I am UP !",
	})
}
