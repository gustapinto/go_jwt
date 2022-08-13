package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gustapinto/go_jwt/controller"
)

func main() {
	router := gin.Default()
	controller := controller.NewUserController()

	router.POST("/api/user", controller.Create)
	router.POST("/api/auth", controller.Auth)
	router.Run()
}
