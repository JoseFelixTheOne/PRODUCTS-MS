package http

type ProductQuery struct {
	Page       int      `form:"page"`
	PageSize   int      `form:"page_size"`
	Q          string   `form:"q"`
	CategoryID *uint    `form:"category_id"`
	MinPrice   *float64 `form:"min_price"`
	MaxPrice   *float64 `form:"max_price"`
	InStock    *bool    `form:"in_stock"`
	Active     *bool    `form:"active"`
	SortBy     string   `form:"sort"`  // name|price|created_at
	Order      string   `form:"order"` // asc|desc
}

// Respuesta paginada estándar

type PageResponse[T any] struct {
	Items      []T   `json:"items"`
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
	HasNext    bool  `json:"has_next"`
	HasPrev    bool  `json:"has_prev"`
}
