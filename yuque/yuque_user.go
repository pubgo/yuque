package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/abc"
	"github.com/pubgo/yuque/models"
)

var _ abc.YuqueUser = (*yqUser)(nil)

type yqUser struct {
	c *resty.Request
}

func (t *yqUser) GetMe() (res *models.UserDetail, err error) {
	_dt := make(map[string]*models.UserDetail)
	return _dt["data"], xerror.Wrap(yqGet(t.c, abc.GetMe(), nil, &_dt), "yqUser.GetMe")
}

func (t *yqUser) GetUser(userId string) (res *models.UserDetail, err error) {
	_dt := make(map[string]*models.UserDetail)
	return _dt["data"], xerror.Wrap(yqGet(t.c, abc.GetUser(userId), nil, &_dt), "yqUser.GetUser")
}

func (t *yqUser) GetUserByName(username string) (res *models.UserDetail, err error) {
	return t.GetUser(username)
}
