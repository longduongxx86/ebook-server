package handlers

import (
	"net/http"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryRepo *repositories.CategoryRepository
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		categoryRepo: repositories.NewCategoryRepository(),
	}
}

// GetListCategory handles GET /categories
func (h *CategoryHandler) GetListCategory(c *gin.Context) {
	categories, err := h.categoryRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}