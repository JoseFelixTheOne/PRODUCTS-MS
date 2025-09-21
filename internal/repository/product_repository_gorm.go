package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/JoseFelixTheOne/products-ms/internal/domain"
	"gorm.io/gorm"
)

type productRepositoryGorm struct {
	db *gorm.DB
}

func NewProductRepositoryGorm(db *gorm.DB) ProductRepository {
	return &productRepositoryGorm{db: db}
}

func (r *productRepositoryGorm) baseQuery(filters ProductFilters) *gorm.DB {
	q := r.db.Model(&domain.Product{}).Preload("Category")

	if filters.Query != "" {
		like := fmt.Sprintf("%%%s%%", strings.TrimSpace(filters.Query))
		q = q.Where("Name LIKE ? OR SKU LIKE ?", like, like)
	}
	if filters.CategoryID != nil {
		q = q.Where("CategoryID = ?", *filters.CategoryID)
	}
	if filters.MinPrice != nil {
		q = q.Where("Price >= ?", *filters.MinPrice)
	}
	if filters.MaxPrice != nil {
		q = q.Where("Price <= ?", *filters.MaxPrice)
	}
	if filters.InStock != nil {
		if *filters.InStock {
			q = q.Where("Stock > 0")
		} else {
			q = q.Where("Stock <= 0")
		}
	}
	if filters.Active != nil {
		q = q.Where("Active = ?", *filters.Active)
	}

	// Orden
	sortCol := map[string]string{
		"name":       "Name",
		"price":      "Price",
		"created_at": "CreatedAt",
	}[strings.ToLower(filters.SortBy)]
	if sortCol == "" {
		sortCol = "CreatedAt"
	}
	order := "ASC"
	if strings.ToLower(filters.Order) == "desc" {
		order = "DESC"
	}
	q = q.Order(sortCol + " " + order)

	return q
}

func (r *productRepositoryGorm) List(ctx context.Context, page, pageSize int, filters ProductFilters) (*Page, error) {
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 200 {
		pageSize = 20
	}

	q := r.baseQuery(filters)

	var total int64
	if err := q.WithContext(ctx).Count(&total).Error; err != nil {
		return nil, err
	}

	var products []domain.Product
	offset := (page - 1) * pageSize
	if err := q.WithContext(ctx).Limit(pageSize).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}

	pages := int((total + int64(pageSize) - 1) / int64(pageSize))
	p := &Page{
		Items:      products,
		Page:       page,
		PageSize:   pageSize,
		TotalItems: total,
		TotalPages: pages,
		HasNext:    page < pages,
		HasPrev:    page > 1,
	}
	return p, nil
}

func (r *productRepositoryGorm) GetCategories(ctx context.Context) ([]domain.Category, error) {
	var cats []domain.Category
	if err := r.db.WithContext(ctx).Order("Name ASC").Find(&cats).Error; err != nil {
		return nil, err
	}
	return cats, nil
}
