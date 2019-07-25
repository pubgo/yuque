package models

// user_detail
// https://www.yuque.com/yuque/developer/userdetailserializer
type UserDetail struct {
	AccountID        int         `json:"account_id"`
	AvatarURL        string      `json:"avatar_url"`
	BooksCount       int         `json:"books_count"`
	CreatedAt        string      `json:"created_at"`
	Description      interface{} `json:"description"`
	FollowersCount   int         `json:"followers_count"`
	FollowingCount   int         `json:"following_count"`
	ID               int         `json:"id"`
	LargeAvatarURL   string      `json:"large_avatar_url"`
	Login            string      `json:"login"`
	MediumAvatarURL  string      `json:"medium_avatar_url"`
	Name             string      `json:"name"`
	Public           int         `json:"public"`
	PublicBooksCount int         `json:"public_books_count"`
	SmallAvatarURL   string      `json:"small_avatar_url"`
	SpaceID          int         `json:"space_id"`
	Type             string      `json:"type"`
	UpdatedAt        string      `json:"updated_at"`
}
