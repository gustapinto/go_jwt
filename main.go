package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gustapinto/go_jwt/controller"
)

func main() {
	router := gin.Default()
	userController := controller.NewUserController()
	fooController := &controller.FooController{}

	router.POST("/api/user", userController.Create)
	router.POST("/api/auth", userController.Auth)
	router.GET("/api/foo", fooController.Foo)

	router.Run()
}
