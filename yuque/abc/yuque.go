package abc

type IYuque interface {
	Auth(token string) IYuque
	Group() YuQueGroup
	User() YuQueUser
	Repo() YuQueRepo
	Doc() YuQueDoc
}
