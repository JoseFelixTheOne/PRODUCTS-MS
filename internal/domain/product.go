package domain

import "time"

type Product struct {
	ID          uint    `gorm:"primaryKey;column:ProductID" json:"id"`
	SKU         string  `gorm:"size:64;not null;uniqueIndex" json:"sku"`
	Name        string  `gorm:"size:200;not null;index" json:"name"`
	Description string  `gorm:"type:nvarchar(max)" json:"description"`
	Price       float64 `gorm:"type:decimal(18,2);not null;index" json:"price"`
	Stock       int     `gorm:"not null;index" json:"stock"`
	Active      bool    `gorm:"not null;default:true;index" json:"active"`

	CategoryID uint     `gorm:"not null;index" json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"category"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Product) TableName() string { return "Product" }
