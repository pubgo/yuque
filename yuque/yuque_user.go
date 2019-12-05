package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/abc"
	"github.com/pubgo/yuque/models"
)

var _ abc.YuqueUser = &yqUser{}

type yqUser struct {
	c *resty.Request
}

func (t *yqUser) GetMe() (res *models.UserDetail, err error) {
	defer xerror.RespErr(&err)

	_dt := make(map[string]*models.UserDetail)
	xerror.Panic(yqGet(t.c, abc.GetMe(), nil, &_dt))
	return _dt["data"], nil
}

func (t *yqUser) GetUser(userId string) (res *models.UserDetail, err error) {
	defer xerror.RespErr(&err)

	_dt := make(map[string]*models.UserDetail)
	xerror.Panic(yqGet(t.c, abc.GetUser(userId), nil, &_dt))
	return _dt["data"], nil
}

func (t *yqUser) GetUserByName(username string) (res *models.UserDetail, err error) {
	return t.GetUser(username)
}
