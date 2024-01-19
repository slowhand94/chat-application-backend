package router

import (
	"slowhand94/chat-application-backend/controller"

	"github.com/gin-gonic/gin"
)

/*
 *	router object
 */
var router *gin.Engine

/*
 * initialize a router object if not created
 * @TODO : make custom router with specified configurations
 */
func CreateRouter() {
	if router == nil {
		router = gin.Default()
	}
}

/*
 * add routes to the router object
 */
func SetupRoutes() {
	router.GET("/health-check", controller.HealthCheck)
}

/*
 * start the router
 * @TODO : start with custom settings in future
 */
func StartRouter() {
	router.Run()
}
