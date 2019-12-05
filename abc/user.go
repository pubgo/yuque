// doc https://www.yuque.com/yuque/developer/user

package abc

import (
	"fmt"
	"github.com/pubgo/yuque/models"
)

type YuqueUser interface {
	GetMe() (res *models.UserDetail, err error)
	GetUser(userId string) (res *models.UserDetail, err error)
	GetUserByName(username string) (res *models.UserDetail, err error)
}

var _url = func(url string) func(...interface{}) string {
	return func(params ...interface{}) string {
		return fmt.Sprintf(url, params...)
	}
}

// 获取单个用户信息
// 基于用户 login 或 id 获取一个用户的基本信息
var GetUser = _url("/users/%s")

// 获取认证的用户的个人信息 [需要认证]
// 获取当前 Token 对应的用户的个人信息
// 返回结果如同 GET /users/:login
var GetMe = _url("/user")
