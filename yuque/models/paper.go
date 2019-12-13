package models

type Paper struct {
	Prev struct {
		Title      string `json:"title"`
		Slug       string `json:"slug"`
		Serializer string `json:"_serializer"`
	} `json:"prev"`
	Next Next `json:"next"`
}

type Next struct {
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Serializer string `json:"_serializer"`
}
