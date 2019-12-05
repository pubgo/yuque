package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/abc"
	"github.com/pubgo/yuque/models"
)

var _ abc.YuqueGroup = &yqGroup{}

type yqGroup struct {
	c *resty.Request
}

func (t *yqGroup) CreateGroup(group *models.GroupCreate) (_ *models.GroupDetail, err error) {
	panic("implement me")
}

func (t *yqGroup) GetGroup(groupId string) (_ *models.GroupDetail, err error) {
	panic("implement me")
}

func (t *yqGroup) GetGroupByName(groupName string) (_ *models.GroupDetail, err error) {
	panic("implement me")
}

func (t *yqGroup) UpdateGroup(groupId string) func(group *models.GroupCreate) (_ *models.GroupDetail, err error) {
	panic("implement me")
}

func (t *yqGroup) UpdateGroupByName(groupName string) func(group *models.GroupCreate) (_ *models.GroupDetail, err error) {
	panic("implement me")
}

func (t *yqGroup) DelGroup(groupId string) (_ *models.GroupDetail, err error) {
	panic("implement me")
}

func (t *yqGroup) DelGroupByName(groupName string) (_ *models.GroupDetail, err error) {
	panic("implement me")
}

func (t *yqGroup) GetGroupMembers(groupId string) (_ []*models.GroupUser, err error) {
	panic("implement me")
}

func (t *yqGroup) GetGroupMembersByName(groupName string) (_ []*models.GroupUser, err error) {
	panic("implement me")
}

func (t *yqGroup) UpdateGroupMember(groupId, userName string) func(_ *models.GroupCreate) (_ *models.GroupUser, err error) {
	panic("implement me")
}

func (t *yqGroup) UpdateGroupMemberByName(groupName, userName string) func(_ *models.GroupCreate) (_ *models.GroupUser, err error) {
	panic("implement me")
}

func (t *yqGroup) DelGroupMember(groupId, userName string) (_ *models.GroupUser, err error) {
	panic("implement me")
}

func (t *yqGroup) DelGroupMemberByName(groupName, userName string) (_ *models.GroupUser, err error) {
	panic("implement me")
}

func (t *yqGroup) GetUserGroups(userId string) (res []*models.Group, err error) {
	defer xerror.RespErr(&err)

	_dt := make(map[string][]*models.Group)
	xerror.Panic(yqGet(t.c, abc.GetUserGroups(userId), nil, &_dt))
	return _dt["data"], nil
}

func (t *yqGroup) GetUserGroupsByName(username string) (res []*models.Group, err error) {
	return t.GetUserGroups(username)
}

func (t *yqGroup) GetMyPubGroups() (res []*models.Group, err error) {
	defer xerror.RespErr(&err)

	_dt := make(map[string][]*models.Group)
	xerror.Panic(yqGet(t.c, abc.GetMyPubGroups(), nil, &_dt))
	return _dt["data"], nil
}
