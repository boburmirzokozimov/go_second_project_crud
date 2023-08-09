package models

import "time"

type Post struct {
	ID      uint      `json:"id" gorm:"primaryKey"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Time    time.Time `json:"created_at"`
}
