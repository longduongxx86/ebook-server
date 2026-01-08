package handlers

import (
	"net/http"
	"strings"

	"main/internal/config"
	"main/internal/models"
	"main/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UploadAvatar uploads user avatar image
func UploadAvatar(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Get file from form
	file, header, err := c.Request.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided or invalid file"})
		return
	}
	defer file.Close()

	// Validate file type
	contentType := header.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File must be an image"})
		return
	}

	// Validate file size (max 5MB)
	if header.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size must be less than 5MB"})
		return
	}

	// Upload to MinIO
	fileURL, err := utils.UploadFile(file, header.Filename, contentType, "avatars")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file: " + err.Error()})
		return
	}

	// Update user avatar in database
	db := config.GetDB()
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		}
		return
	}

	// Delete old avatar if exists (optional - you can implement this)
	// oldAvatar := user.AvatarURL
	// if oldAvatar != "" {
	// 	objectPath := utils.ExtractObjectPathFromURL(oldAvatar)
	// 	utils.DeleteFile(objectPath)
	// }

	// Delete old avatar if exists
	oldAvatar := user.AvatarURL
	if oldAvatar != "" {
		objectPath := utils.ExtractObjectPathFromURL(oldAvatar)
		if objectPath != "" {
			utils.DeleteFile(objectPath) // Best effort, ignore error
		}
	}

	// Update avatar URL
	user.AvatarURL = fileURL
	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update avatar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Avatar uploaded successfully",
		"avatar_url": fileURL,
	})
}

func UploadBookImage(c *gin.Context) {
	_, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Get file from form
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided or invalid file"})
		return
	}
	defer file.Close()

	// Validate file type
	contentType := header.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File must be an image"})
		return
	}

	// Validate file size (max 5MB)
	if header.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size must be less than 5MB"})
		return
	}

	// Upload to MinIO
	fileURL, err := utils.UploadFile(file, header.Filename, contentType, "books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Image uploaded successfully",
		"image_url": fileURL,
	})
}

