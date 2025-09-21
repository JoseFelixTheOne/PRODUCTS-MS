package service

import (
	"context"

	"github.com/JoseFelixTheOne/products-ms/internal/repository"
)

type ProductService struct{ repo repository.ProductRepository }

func NewProductService(r repository.ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) List(ctx context.Context, page, pageSize int, filters repository.ProductFilters) (*repository.Page, error) {
	return s.repo.List(ctx, page, pageSize, filters)
}

func (s *ProductService) Categories(ctx context.Context) (any, error) {
	return s.repo.GetCategories(ctx)
}
