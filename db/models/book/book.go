package book

import (
	"time"
)

type BookModel struct {
	ID        uint    `gorm:"primary_key"`
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	Price     float64 `json:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (BookModel) TableName() string {
	return "book"
}
