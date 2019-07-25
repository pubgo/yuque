package models

type Book struct {
	ContentUpdatedAt string     `json:"content_updated_at"`
	CreatedAt        string     `json:"created_at"`
	CreatorID        int        `json:"creator_id"`
	Description      string     `json:"description"`
	ID               int        `json:"id"`
	ItemsCount       int        `json:"items_count"`
	LikesCount       int        `json:"likes_count"`
	Name             string     `json:"name"`
	Namespace        string     `json:"namespace"`
	Public           int        `json:"public"`
	Slug             string     `json:"slug"`
	Type             string     `json:"type"`
	UpdatedAt        string     `json:"updated_at"`
	User             UserDetail `json:"user"`
	UserID           int        `json:"user_id"`
	WatchesCount     int        `json:"watches_count"`
}

type BookCreate struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	Public      int    `json:"public"`
	Slug        string `json:"slug"`
	Type        string `json:"type"`
}

type BookToc struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
	Depth string `json:"depth"`
}
