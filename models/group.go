package models

// group
// https://www.yuque.com/yuque/developer/groupuserserializer

type GroupUser struct {
	CreatedAt string      `json:"created_at"`
	Group     interface{} `json:"group"`
	GroupID   int         `json:"group_id"`
	ID        int         `json:"id"`
	Role      int         `json:"role"`
	Status    int         `json:"status"`
	UpdatedAt string      `json:"updated_at"`
	User      UserDetail  `json:"user"`
	UserID    int         `json:"user_id"`
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
	Abilities struct {
		Destroy   bool `json:"destroy"`
		GroupUser struct {
			Create  bool `json:"create"`
			Destroy bool `json:"destroy"`
			Update  bool `json:"update"`
		} `json:"group_user"`
		Read bool `json:"read"`
		Repo struct {
			Create  bool `json:"create"`
			Destroy bool `json:"destroy"`
			Update  bool `json:"update"`
		} `json:"repo"`
		Update bool `json:"update"`
	} `json:"abilities"`
	Data Group `json:"data"`
	Meta struct {
		TopicEnable int `json:"topic_enable"`
	} `json:"meta"`
}

type GroupCreate struct {
	Name        string `json:"name"`
	Login       string `json:"login"`
	Description string `json:"description"`
}

type Group struct {
	Id *string `json:"id" xml:"id"`
	Login *string `json:"login" xml:"login" require:"true"`
	Name *string `json:"name" xml:"name" require:"true"`
	AvatarUrl *string `json:"avatar_url" xml:"avatar_url"`
	OwnerId *string `json:"owner_id" xml:"owner_id" require:"true"`
	Description *string `json:"description" xml:"description" require:"true"`
	CreatedAt *string `json:"created_at" xml:"created_at" require:"true"`
	UpdatedAt *string `json:"updated_at" xml:"updated_at" require:"true"`
}
