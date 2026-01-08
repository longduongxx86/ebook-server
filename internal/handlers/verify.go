package handlers

import (
	"net/http"

	"main/internal/config"
	"main/internal/models"

	"github.com/gin-gonic/gin"
)

// VerifyEmail marks user email as verified using token
// GET /api/v1/verify-email?token=...
func VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token is required"})
		return
	}

	db := config.GetDB()
	var user models.User
	if err := db.Where("verification_token = ?", token).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}

	user.EmailVerified = true
	user.VerificationToken = ""
	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to verify email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}
