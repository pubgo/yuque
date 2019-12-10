package models

type Search struct {
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	Public      int    `json:"public"`
	Slug        string `json:"slug"`
	Type        string `json:"type"`
	UpdatedAt   string `json:"updated_at"`
	User        *User  `json:"user"`
	UserID      int    `json:"user_id"`
}
