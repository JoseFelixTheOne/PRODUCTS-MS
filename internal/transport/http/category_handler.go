package http

import (
	"context"
	"net/http"
	"time"

	"github.com/JoseFelixTheOne/products-ms/internal/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct{ svc *service.ProductService }

func NewCategoryHandler(s *service.ProductService) *CategoryHandler { return &CategoryHandler{svc: s} }

func (h *CategoryHandler) List(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	cats, err := h.svc.Categories(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cats)
}
