package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/yuque/abc"
	"github.com/pubgo/yuque/yuque/models"
)

var (
	_        abc.YuQueUser = (*yqUserImpl)(nil)
	_GetUser               = _url("/users/%s")
	_GetMe                 = _url("/user")
)

type yqUserImpl struct {
	c *resty.Client
}

func (t yqUserImpl) GetMe() (res *models.UserDetail, err error) {
	_dt := make(map[string]*models.UserDetail)
	return _dt["data"], xerror.Wrap(yqGet(t.c.R(), _GetMe(), nil, &_dt), "YQUser.GetMe")
}

func (t yqUserImpl) GetUser(userId string) (res *models.UserDetail, err error) {
	_dt := make(map[string]*models.UserDetail)
	return _dt["data"], xerror.Wrap(yqGet(t.c.R(), _GetUser(userId), nil, &_dt), "YQUser.GetUser")
}

func (t yqUserImpl) GetUserByName(username string) (res *models.UserDetail, err error) {
	return t.GetUser(username)
}
