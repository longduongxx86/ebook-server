package handlers

import (
	"net/http"
	"strings"

	"main/internal/config"
	"main/internal/models"
	"main/internal/utils"

	"github.com/gin-gonic/gin"
)

// GoogleAuth handles Google Sign-In using an ID token
// POST /api/v1/auth/google
// Body: { "id_token": "..." }
func GoogleAuth(c *gin.Context) {
	var body struct {
		IDToken string `json:"id_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email, sub, name, err := utils.VerifyGoogleIDToken(strings.TrimSpace(body.IDToken))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Google token"})
		return
	}

	db := config.GetDB()

	// Find user by GoogleID or by Email
	var user models.User
	if err := db.Where("google_id = ?", sub).First(&user).Error; err != nil {
		// Try by email
		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			// Generate username from email (use part before @ as username)
			// If username already exists, append Google ID
			usernameBase := strings.Split(email, "@")[0]
			username := usernameBase

			// Check if username exists and generate unique one
			var existingUser models.User
			retryCount := 0
			for db.Where("username = ?", username).First(&existingUser).Error == nil {
				retryCount++
				username = usernameBase + "_" + sub[:8] // Use first 8 chars of Google ID
				if retryCount > 10 {
					// Fallback: use full sub as username
					username = "user_" + sub
					break
				}
			}

			// Create new user
			user = models.User{
				Email:    email,
				FullName: name,
				Role:     "customer", // Default role for new registrations
			}

			// create verification token and send email (best-effort)
			token, _ := utils.GenerateToken(16)
			user.VerificationToken = token

			if err := db.Create(&user).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
				return
			}

			_ = utils.SendVerificationEmail(user.Email, token)
		}
	}

	// Ensure role is set (for existing users without role)
	if user.Role == "" {
		user.Role = "customer"
		db.Save(&user)
	}

	// Issue JWT with role
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login with Google successful",
		"token":   token,
		"user":    user,
		"role":    user.Role, // Return role explicitly for frontend to determine UI
	})
}
