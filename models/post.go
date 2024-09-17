package models

import "time"

type Post struct {
	ID          uint       `json:"id" gorm:"primaryKey; autoIncrement"`
	UserID      uint       `json:"user_id"` //foreign key
	Title       string     `json:"title"`
	Slug        string     `json:"slug" gorm:"unique"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	User        User       `json:"user" gorm:"foreignKey:UserID"` //association
}
