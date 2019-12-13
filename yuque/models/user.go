package models

// 一般在列表的场景返回的用户信息
// doc https://www.yuque.com/yuque/developer/userserializer
type User struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	login     string
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// 用户/团队详细信息
// doc https://www.yuque.com/yuque/developer/userdetailserializer
type UserDetail struct {
	AccountID        int    `json:"account_id"`
	AvatarURL        string `json:"avatar_url"`
	BooksCount       int    `json:"books_count"`
	CreatedAt        string `json:"created_at"`
	Description      string `json:"description"`
	FollowersCount   int    `json:"followers_count"`
	FollowingCount   int    `json:"following_count"`
	ID               int    `json:"id"`
	LargeAvatarURL   string `json:"large_avatar_url"`
	Login            string `json:"login"`
	MediumAvatarURL  string `json:"medium_avatar_url"`
	Name             string `json:"name"`
	Public           int    `json:"public"`
	PublicBooksCount int    `json:"public_books_count"`
	SmallAvatarURL   string `json:"small_avatar_url"`
	SpaceID          int    `json:"space_id"`
	Type             string `json:"type"`
	UpdatedAt        string `json:"updated_at"`
}
