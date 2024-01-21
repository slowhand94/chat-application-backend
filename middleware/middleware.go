package middleware

import "github.com/gin-gonic/gin"

/*
 * middleware to authenticate user
 */
func Authenticate(context *gin.Context) {
	// @todo : add logic later
	context.Next()
}
