package router

import (
	"slowhand94/chat-application-backend/controller"
	"slowhand94/chat-application-backend/middleware"

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
	// health check routes
	router.GET("/health-check", controller.HealthCheck)

	// user routes
	router.POST("/user/sign-up", controller.SignUp)
	router.POST("/user/login", controller.Login)
	router.GET("/user/logout", middleware.Authenticate, controller.Logout)

	// chat specific routes
	router.POST("/chat/create-room", middleware.Authenticate, controller.CreateRoom)
	router.GET("/chat/join-room/:roomCode", middleware.Authenticate, controller.JoinRoom)
}

/*
 * start the router
 * @TODO : start with custom settings in future
 */
func StartRouter() {
	router.Run(":8080")
}
