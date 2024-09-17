package models

import "time"

type Comment struct {
	ID        uint       `json:"id" gorm:"primaryKey; autoIncrement"`
	Comment   string     `json:"comment"`
	UserID    uint       `json:"user_id"`
	PostID    uint       `json:"post_id"`
	User      User       `json:"user" gorm:"foreignKey:UserID"`
	Post      Post       `json:"post" gorm:"foreignKey:PostID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
