package models

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `gorm:"not null; index" json:"user_id"`
	Title     string    `gorm:"verchar(200); not null" json:"title"`
	Body      string    `gorm:"not null" json:"body"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}
