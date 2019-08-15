package routes

import (
	"github.com/pubgo/yuque/models"
)

// 获取认证的用户的个人信息
// 获取当前 Token 对应的用户的个人信息
const GetMe = apiPrefix + "/user"

// 获取单个用户信息
const GetUser = apiPrefix + "/users/{user_id}"
const GetUserByName = apiPrefix + "/users/{username}"

type YuqueUser interface {
	GetMe() (res *models.UserDetail, err error)
	GetUser(userId string) (res *models.UserDetail, err error)
	GetUserByName(username string) (res *models.UserDetail, err error)
}
