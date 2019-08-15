package internal

import (
	"github.com/pubgo/g/errors"
	"github.com/pubgo/yuque/internal/routes"
	"github.com/pubgo/yuque/models"
)

func (t *yqClient) GetUserGroups(userId string) (res []*models.Group, err error) {
	defer errors.RespErr(&err)

	_dt := make(map[string][]*models.Group)
	errors.Panic(t.get(routes.GetUserByName, map[string]string{"user_id": userId}, nil, &_dt))
	return _dt["data"], nil
}

func (t *yqClient) GetUserGroupsByName(username string) (res []*models.Group, err error) {
	defer errors.RespErr(&err)

	_dt := make(map[string][]*models.Group)
	errors.Panic(t.get(routes.GetUserGroupsByName, map[string]string{"username": username}, nil, &_dt))
	return _dt["data"], nil
}

func (t *yqClient) GetMyPubGroups() (res []*models.Group, err error) {
	defer errors.RespErr(&err)

	_dt := make(map[string][]*models.Group)
	errors.Panic(t.get(routes.GetUserGroupsByName, map[string]string{"username": username}, nil, &_dt))
	return _dt["data"], nil
}
