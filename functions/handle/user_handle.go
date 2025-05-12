package handler

import (
	"net/http"
	"pheuanpet-user-management/service"

	"github.com/gin-gonic/gin"
)
var userService *service.UserService

func HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func GetUsers(c *gin.Context) {
    // ดึงข้อมูล user ทั้งหมด
    users, err := userService.GetAllUsers(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
        return
    }
    c.JSON(http.StatusOK, users)

}

// GetUserDetailsHandler handles the request to get user details
func GetUserDetailsHandler(c *gin.Context) {
    userID := c.Param("id")
    userDetails, err := userService.GetUserDetails(c.Request.Context(), userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user details"})
        return
    }
    c.JSON(http.StatusOK, userDetails)
}
