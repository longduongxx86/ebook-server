package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequireRole checks if user has required role
func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User role not found"})
			c.Abort()
			return
		}

		role, ok := userRole.(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user role"})
			c.Abort()
			return
		}

		// Check if user role is in allowed roles
		hasPermission := false
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireCustomer checks if user is a customer
func RequireCustomer() gin.HandlerFunc {
	return RequireRole("customer")
}

// RequireManager checks if user is a manager
func RequireManager() gin.HandlerFunc {
	return RequireRole("manager", "admin")
}

