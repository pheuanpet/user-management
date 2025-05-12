package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func GetUsers(c *gin.Context) {
    // ดึงข้อมูล user ทั้งหมด
}

func CreateUser(c *gin.Context) {
    // สร้าง user ใหม่
}