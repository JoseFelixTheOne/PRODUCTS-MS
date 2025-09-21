package repository

import (
	"context"

	"github.com/JoseFelixTheOne/products-ms/internal/domain"
)

type ProductFilters struct {
	Query      string
	CategoryID *uint
	MinPrice   *float64
	MaxPrice   *float64
	InStock    *bool
	Active     *bool
	SortBy     string // name|price|created_at
	Order      string // asc|desc
}

type Page struct {
	Items      any
	Page       int
	PageSize   int
	TotalItems int64
	TotalPages int
	HasNext    bool
	HasPrev    bool
}

type ProductRepository interface {
	List(ctx context.Context, page, pageSize int, filters ProductFilters) (*Page, error)
	GetCategories(ctx context.Context) ([]domain.Category, error)
}
