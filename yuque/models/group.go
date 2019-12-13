package models

// group
// https://www.yuque.com/yuque/developer/groupuserserializer

type GroupUser struct {
	CreatedAt string           `json:"created_at"`
	Group     interface{}      `json:"group"`
	GroupID   int              `json:"group_id"`
	ID        int              `json:"id"`
	Role      int              `json:"role"`
	Status    int              `json:"status"`
	UpdatedAt string           `json:"updated_at"`
	User      UserDetail `json:"user"`
	UserID    int              `json:"user_id"`
}

type Group struct {
	AvatarURL         string `json:"avatar_url"`
	BooksCount        int    `json:"books_count"`
	CreatedAt         string `json:"created_at"`
	Description       string `json:"description"`
	ID                int    `json:"id"`
	LargeAvatarURL    string `json:"large_avatar_url"`
	Login             string `json:"login"`
	MediumAvatarURL   string `json:"medium_avatar_url"`
	MembersCount      int    `json:"members_count"`
	Name              string `json:"name"`
	Public            int    `json:"public"`
	PublicBooksCount  int    `json:"public_books_count"`
	PublicTopicsCount int    `json:"public_topics_count"`
	SmallAvatarURL    string `json:"small_avatar_url"`
	TopicsCount       int    `json:"topics_count"`
	UpdatedAt         string `json:"updated_at"`
}

type GroupDetail struct {
	AvatarURL         string `json:"avatar_url"`
	BooksCount        int    `json:"books_count"`
	CreatedAt         string `json:"created_at"`
	Description       string `json:"description"`
	ID                int    `json:"id"`
	LargeAvatarURL    string `json:"large_avatar_url"`
	Login             string `json:"login"`
	MediumAvatarURL   string `json:"medium_avatar_url"`
	MembersCount      int    `json:"members_count"`
	Name              string `json:"name"`
	OrganizationID    int    `json:"organization_id"`
	OwnerID           int    `json:"owner_id"`
	Public            int    `json:"public"`
	PublicBooksCount  int    `json:"public_books_count"`
	PublicTopicsCount int    `json:"public_topics_count"`
	SmallAvatarURL    string `json:"small_avatar_url"`
	SpaceID           int    `json:"space_id"`
	TopicsCount       int    `json:"topics_count"`
	UpdatedAt         string `json:"updated_at"`
}

type GroupCreate struct {
	Name        string `json:"name"`
	Login       string `json:"login"`
	Description string `json:"description"`
}
