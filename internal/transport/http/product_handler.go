package http

import (
	"context"
	"net/http"
	"time"

	"github.com/JoseFelixTheOne/products-ms/internal/domain"
	"github.com/JoseFelixTheOne/products-ms/internal/repository"
	"github.com/JoseFelixTheOne/products-ms/internal/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct{ svc *service.ProductService }

func NewProductHandler(s *service.ProductService) *ProductHandler { return &ProductHandler{svc: s} }

func (h *ProductHandler) List(c *gin.Context) {
	var q ProductQuery
	if err := c.BindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetros inválidos"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	filters := repository.ProductFilters{
		Query:      q.Q,
		CategoryID: q.CategoryID,
		MinPrice:   q.MinPrice,
		MaxPrice:   q.MaxPrice,
		InStock:    q.InStock,
		Active:     q.Active,
		SortBy:     q.SortBy,
		Order:      q.Order,
	}
	page, err := h.svc.List(ctx, q.Page, q.PageSize, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	items, _ := page.Items.([]domain.Product)
	resp := PageResponse[domain.Product]{
		Items:      items,
		Page:       page.Page,
		PageSize:   page.PageSize,
		TotalItems: page.TotalItems,
		TotalPages: page.TotalPages,
		HasNext:    page.HasNext,
		HasPrev:    page.HasPrev,
	}
	c.JSON(http.StatusOK, resp)
}
