package router

import (
	handler "pheuanpet-user-management/handle"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/users", handler.GetUsers)
    return r
}