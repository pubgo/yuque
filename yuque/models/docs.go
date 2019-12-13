package models

type Doc struct {
	Book              interface{} `json:"book"`
	BookID            int         `json:"book_id"`
	CommentsCount     int         `json:"comments_count"`
	ContentUpdatedAt  string      `json:"content_updated_at"`
	Cover             interface{} `json:"cover"`
	CreatedAt         string      `json:"created_at"`
	CustomDescription interface{} `json:"custom_description"`
	Description       string      `json:"description"`
	DraftVersion      int         `json:"draft_version"`
	FirstPublishedAt  string      `json:"first_published_at"`
	Format            string      `json:"format"`
	ID                int         `json:"id"`
	LastEditor        struct {
		AvatarURL       string `json:"avatar_url"`
		CreatedAt       string `json:"created_at"`
		Description     string `json:"description"`
		FollowersCount  int    `json:"followers_count"`
		FollowingCount  int    `json:"following_count"`
		ID              int    `json:"id"`
		LargeAvatarURL  string `json:"large_avatar_url"`
		Login           string `json:"login"`
		MediumAvatarURL string `json:"medium_avatar_url"`
		Name            string `json:"name"`
		SmallAvatarURL  string `json:"small_avatar_url"`
		Type            string `json:"type"`
		UpdatedAt       string `json:"updated_at"`
	} `json:"last_editor"`
	LastEditorID int    `json:"last_editor_id"`
	LikesCount   int    `json:"likes_count"`
	Public       int    `json:"public"`
	PublishedAt  string `json:"published_at"`
	Slug         string `json:"slug"`
	Status       int    `json:"status"`
	Title        string `json:"title"`
	UpdatedAt    string `json:"updated_at"`
	UserID       int    `json:"user_id"`
	WordCount    int    `json:"word_count"`
}

type DocCreate struct {
	Body   string `json:"body"`
	Format string `json:"format"`
	Public int    `json:"public"`
	Slug   string `json:"slug"`
	Title  string `json:"title"`
}
