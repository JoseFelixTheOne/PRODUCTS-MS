package domain

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey;column:CategoryID" json:"id"`
	Name      string    `gorm:"size:100;not null;uniqueIndex" json:"name"`
	Slug      string    `gorm:"size:120;not null;uniqueIndex" json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Category) TableName() string { return "Category" }
