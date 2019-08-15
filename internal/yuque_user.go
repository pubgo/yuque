package internal

import (
	"github.com/pubgo/g/errors"
	"github.com/pubgo/yuque/internal/routes"
	"github.com/pubgo/yuque/models"
)

func (t *yqClient) GetMe() (res *models.UserDetail, err error) {
	defer errors.RespErr(&err)

	_dt := make(map[string]*models.UserDetail)
	errors.Panic(t.get(routes.GetMe, nil, nil, &_dt))
	return _dt["data"], nil
}

func (t *yqClient) GetUser(userId string) (res *models.UserDetail, err error) {
	defer errors.RespErr(&err)

	_dt := make(map[string]*models.UserDetail)
	errors.Panic(t.get(routes.GetUser, map[string]string{"user_id": userId}, nil, &_dt))
	return _dt["data"], nil
}

func (t *yqClient) GetUserByName(username string) (res *models.UserDetail, err error) {
	defer errors.RespErr(&err)

	_dt := make(map[string]*models.UserDetail)
	errors.Panic(t.get(routes.GetUserByName, map[string]string{"username": username}, nil, &_dt))
	return _dt["data"], nil
}
