package models

import (
	"time"
)

type Comment struct {
	ID        uint      `json:"id" bson:"id"`
	Comment   string    `json:"comment" bson:"comment"`
	PostID    uint      `json:"-" bson:"post_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Post struct {
	ID        uint      `json:"id" bson:"id"`
	Title     string    `json:"title" bson:"title"`
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	Comments  []Comment `json:"comments" bson:"comments"`
}
