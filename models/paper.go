package models

type Paper struct {
	Data struct {
		Prev struct {
			Title      string `json:"title"`
			Slug       string `json:"slug"`
			Serializer string `json:"_serializer"`
		} `json:"prev"`
		Next Next `json:"next"`
	} `json:"data"`
}

type Next struct {
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Serializer string `json:"_serializer"`
}
