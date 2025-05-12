package router

import (
	"user-management/functions/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/", handler.HealthCheck)
    r.GET("/users", handler.GetUsers)
    r.POST("/users", handler.CreateUser)
    return r
}