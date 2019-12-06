package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/abc"
	"github.com/pubgo/yuque/models"
)

var _ abc.YuqueGroup = (*yqGroup)(nil)

type yqGroup struct {
	c *resty.Request
}

func (t *yqGroup) CreateGroup(data *models.GroupCreate) (_ *models.GroupDetail, err error) {
	_dt := make(map[string]*models.GroupDetail)
	return _dt["data"], xerror.Wrap(yqPost(t.c, abc.CreateGroup(), data, &_dt), "yqGroup.CreateGroup")
}

func (t *yqGroup) GetGroup(groupId string) (_ *models.GroupDetail, err error) {
	_dt := make(map[string]*models.GroupDetail)
	return _dt["data"], xerror.Wrap(yqGet(t.c, abc.UpdateGroupMember(groupId), nil, &_dt), "yqGroup.GetGroup")
}

func (t *yqGroup) GetGroupByName(groupName string) (_ *models.GroupDetail, err error) {
	return t.GetGroup(groupName)
}

func (t *yqGroup) UpdateGroup(groupId string) func(group *models.GroupCreate) (_ *models.GroupDetail, err error) {
	return func(data *models.GroupCreate) (_ *models.GroupDetail, err error) {
		_dt := make(map[string]*models.GroupDetail)
		return _dt["data"], xerror.Wrap(yqPost(t.c, abc.UpdateGroupMember(groupId), data, &_dt), "yqGroup.UpdateGroup")
	}
}

func (t *yqGroup) UpdateGroupByName(groupName string) func(group *models.GroupCreate) (_ *models.GroupDetail, err error) {
	return t.UpdateGroup(groupName)
}

func (t *yqGroup) DelGroup(groupId string) (_ *models.GroupDetail, err error) {
	_dt := make(map[string]*models.GroupDetail)
	return _dt["data"], xerror.Wrap(yqGet(t.c, abc.DelGroup(groupId), nil, &_dt), "yqGroup.DelGroup")
}

func (t *yqGroup) DelGroupByName(groupName string) (_ *models.GroupDetail, err error) {
	return t.DelGroup(groupName)
}

func (t *yqGroup) GetGroupMembers(groupId string) (_ []*models.GroupUser, err error) {
	_dt := make(map[string][]*models.GroupUser)
	return _dt["data"], xerror.Wrap(yqGet(t.c, abc.GetGroupMembers(groupId), nil, &_dt), "yqGroup.GetGroupMembers")
}

func (t *yqGroup) GetGroupMembersByName(groupName string) (_ []*models.GroupUser, err error) {
	return t.GetGroupMembers(groupName)
}

func (t *yqGroup) UpdateGroupMember(groupId, username string) func(_ *models.GroupCreate) (_ *models.GroupUser, err error) {
	return func(data *models.GroupCreate) (_ *models.GroupUser, err error) {
		_dt := make(map[string]*models.GroupUser)
		return _dt["data"], xerror.Wrap(yqPost(t.c, abc.UpdateGroupMember(groupId, username), data, &_dt), "yqGroup.UpdateGroupMember")
	}
}

func (t *yqGroup) UpdateGroupMemberByName(groupName, username string) func(_ *models.GroupCreate) (_ *models.GroupUser, err error) {
	return t.UpdateGroupMember(groupName, username)
}

func (t *yqGroup) DelGroupMember(groupId, username string) (_ *models.GroupUser, err error) {
	_dt := make(map[string]*models.GroupUser)
	return _dt["data"], xerror.Wrap(yqDelete(t.c, abc.DelGroupMember(groupId, username), nil, &_dt), "yqGroup.DelGroupMember")
}

func (t *yqGroup) DelGroupMemberByName(groupName, username string) (_ *models.GroupUser, err error) {
	return t.DelGroupMember(groupName, username)
}

func (t *yqGroup) GetUserGroups(userId string) (res []*models.Group, err error) {
	_dt := make(map[string][]*models.Group)
	return _dt["data"], xerror.Wrap(yqGet(t.c, abc.GetUserGroups(userId), nil, &_dt), "yqGroup.GetUserGroups")
}

func (t *yqGroup) GetUserGroupsByName(username string) (res []*models.Group, err error) {
	return t.GetUserGroups(username)
}

func (t *yqGroup) GetMyPubGroups() (res []*models.Group, err error) {
	_dt := make(map[string][]*models.Group)
	return _dt["data"], xerror.Wrap(yqGet(t.c, abc.GetMyPubGroups(), nil, &_dt), "yqGroup.GetMyPubGroups")
}
