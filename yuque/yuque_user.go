package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/abc"
	"github.com/pubgo/yuque/models"
)

var _ abc.YuqueUser = (*YQUser)(nil)

type YQUser struct {
	c *resty.Request
}

func (t YQUser) GetMe() (res *models.UserDetail, err error) {
	_dt := make(map[string]*models.UserDetail)
	return _dt["data"], xerror.Wrap(yqGet(t.c, abc.GetMe(), nil, &_dt), "YQUser.GetMe")
}

func (t YQUser) GetUser(userId string) (res *models.UserDetail, err error) {
	_dt := make(map[string]*models.UserDetail)
	return _dt["data"], xerror.Wrap(yqGet(t.c, abc.GetUser(userId), nil, &_dt), "YQUser.GetUser")
}

func (t YQUser) GetUserByName(username string) (res *models.UserDetail, err error) {
	return t.GetUser(username)
}
