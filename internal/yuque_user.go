package internal

import (
	"github.com/pubgo/yuque/internal/routes"
	"github.com/pubgo/yuque/models"
)

func (t *yqClient) GetMe() (res *models.UserDetail) {
	_dt := make(map[string]*models.UserDetail)
	t.get(routes.GetMe, nil, nil, &_dt)
	return _dt["data"]
}

func (t *yqClient) GetUser(userId string) (res *models.UserDetail) {
	_dt := make(map[string]*models.UserDetail)
	t.get(routes.GetUser, map[string]string{"user_id": userId}, nil, &_dt)
	return _dt["data"]
}

func (t *yqClient) GetUserByName(username string) (res *models.UserDetail) {
	_dt := make(map[string]*models.UserDetail)
	t.get(routes.GetUserByName, map[string]string{"username": username}, nil, &_dt)
	return _dt["data"]
}

func (t *yqClient) GetUserGroups(userId string) (res []*models.Group) {
	_dt := make(map[string][]*models.Group)
	t.get(routes.GetUserByName, map[string]string{"user_id": userId}, nil, &_dt)
	return _dt["data"]
}

func (t *yqClient) GetUserGroupsByName(username string) (res []*models.Group) {
	_dt := make(map[string][]*models.Group)
	t.get(routes.GetUserGroupsByName, map[string]string{"username": username}, nil, &_dt)
	return _dt["data"]
}
