package http

import (
	"net/http"

	"github.com/JoseFelixTheOne/products-ms/internal/config"
	"github.com/JoseFelixTheOne/products-ms/internal/repository"
	"github.com/JoseFelixTheOne/products-ms/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Healthcheck
	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })

	// Dependencias
	repo := repository.NewProductRepositoryGorm(db)
	svc := service.NewProductService(repo)

	ph := NewProductHandler(svc)
	ch := NewCategoryHandler(svc)

	api := r.Group("/api/v1")
	{
		api.GET("/products", ph.List)
		api.GET("/categories", ch.List)
	}

	return r
}
