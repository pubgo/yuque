package models

import (
	"time"
)

type Popular struct {
	ID          int         `json:"id"`
	SpaceID     int         `json:"space_id"`
	TargetType  string      `json:"target_type"`
	TargetID    int         `json:"target_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Cover       string      `json:"cover"`
	Week        interface{} `json:"week"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Target      struct {
		ID         int         `json:"id"`
		Title      string      `json:"title"`
		Slug       string      `json:"slug"`
		UserID     int         `json:"user_id"`
		BookID     int         `json:"book_id"`
		LikesCount int         `json:"likes_count"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
		Book       Book  `json:"book"`
		LastEditor interface{} `json:"last_editor"`
		Share      interface{} `json:"share"`
		Serializer string      `json:"_serializer"`
	} `json:"target"`
}
