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
	User        struct {
		AvatarURL string      `json:"avatar_url"`
		CreatedAt string      `json:"created_at"`
		Email     interface{} `json:"email"`
		ID        int         `json:"id"`
		Login     string      `json:"login"`
		Name      string      `json:"name"`
		UpdatedAt string      `json:"updated_at"`
		WorkID    interface{} `json:"work_id"`
	} `json:"user"`
	UserID int `json:"user_id"`
}
