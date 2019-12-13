// doc https://www.yuque.com/yuque/developer/user

package abc

import (
	"github.com/pubgo/yuque/yuque/models"
)

type YuQueUser interface {
	// 获取认证的用户的个人信息 [需要认证]
	// 获取当前 Token 对应的用户的个人信息
	// 返回结果如同 GET /users/:login
	GetMe() (res *models.UserDetail, err error)

	// 获取单个用户信息
	// 基于用户 login 或 id 获取一个用户的基本信息
	GetUser(userId string) (res *models.UserDetail, err error)
	GetUserByName(username string) (res *models.UserDetail, err error)
}
