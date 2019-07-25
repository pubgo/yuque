package routes

import (
	"github.com/pubgo/yuque/models"
)

// 获取认证的用户的个人信息
// 获取当前 Token 对应的用户的个人信息
const GetMe = apiProxy + "/user"

// 获取单个用户信息
const GetUser = apiProxy + "/users/{user_id}"
const GetUserByName = apiProxy + "/users/{username}"

type YuqueUser interface {
	GetMe() (res *models.UserDetail)
	GetUser(userId string) (res *models.UserDetail)
	GetUserByName(username string) (res *models.UserDetail)
}
