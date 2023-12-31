package models

import "time"

type Article struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Slug        string    `gorm:"size:255;not null;unique" json:"slug"`
	Description string    `gorm:"size:255;not null" json:"description"`
	Content     string    `gorm:"type:text" json:"content"`
	Status      uint8     `json:"status"`
	Author      string    `gorm:"size:255;not null" json:"author"`
	PublishedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"published_at"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
