package app

import (
	"github.com/studingprojects/bookstore_user-api/controllers"
	"github.com/studingprojects/bookstore_user-api/controllers/users"
)

func initRoutes() {
	router.GET("/ping", controllers.Ping)
	router.GET("/users", users.Search)
	router.GET("/users/:userId", users.GetUser)
	router.POST("/users", users.Create)
	router.PUT("/users/:userId", users.Update)
	router.PATCH("/users/:userId", users.Update)
	router.DELETE("/users/:userId", users.Delete)
	router.POST("/users/login", users.Login)
}
