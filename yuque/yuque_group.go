package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/yuque/abc"
	"github.com/pubgo/yuque/yuque/models"
)

var (
	_                  abc.YuQueGroup = (*yqGroupImpl)(nil)
	_GetUserGroups                    = _url("/users/%s/groups")
	_GetMyPubGroups                   = _url("/groups")
	_CreateGroup                      = _url("/groups")
	_GetGroup                         = _url("/groups/%s")
	_UpdateGroup                      = _url("/groups/%s)")
	_DelGroup                         = _url("/groups/%s)")
	_GetGroupMembers                  = _url("/groups/%s/users")
	_UpdateGroupMember                = _url("/groups/%s/users/%s")
	_DelGroupMember                   = _url("/groups/%s/users/%s")
)

type yqGroupImpl struct {
	c *resty.Client
}

func (t yqGroupImpl) CreateGroup(data *models.GroupCreate) (_ *models.GroupDetail, err error) {
	_dt := make(map[string]*models.GroupDetail)
	return _dt["data"], xerror.Wrap(yqPost(t.c.R(), _CreateGroup(), data, &_dt), "YQGroup.CreateGroup")
}

func (t yqGroupImpl) GetGroup(groupId string) (_ *models.GroupDetail, err error) {
	_dt := make(map[string]*models.GroupDetail)
	return _dt["data"], xerror.Wrap(yqGet(t.c.R(), _GetGroup(groupId), nil, &_dt), "YQGroup.GetGroup")
}

func (t yqGroupImpl) GetGroupByName(groupName string) (_ *models.GroupDetail, err error) {
	return t.GetGroup(groupName)
}

func (t yqGroupImpl) UpdateGroup(groupId string) func(group *models.GroupCreate) (_ *models.GroupDetail, err error) {
	return func(data *models.GroupCreate) (_ *models.GroupDetail, err error) {
		_dt := make(map[string]*models.GroupDetail)
		return _dt["data"], xerror.Wrap(yqPut(t.c.R(), _UpdateGroup(groupId), data, &_dt), "YQGroup.UpdateGroup")
	}
}

func (t yqGroupImpl) UpdateGroupByName(groupName string) func(group *models.GroupCreate) (_ *models.GroupDetail, err error) {
	return t.UpdateGroup(groupName)
}

func (t yqGroupImpl) DelGroup(groupId string) (_ *models.GroupDetail, err error) {
	_dt := make(map[string]*models.GroupDetail)
	return _dt["data"], xerror.Wrap(yqDelete(t.c.R(), _DelGroup(groupId), nil, &_dt), "YQGroup.DelGroup")
}

func (t yqGroupImpl) DelGroupByName(groupName string) (_ *models.GroupDetail, err error) {
	return t.DelGroup(groupName)
}

func (t yqGroupImpl) GetGroupMembers(groupId string) (_ []*models.GroupUser, err error) {
	_dt := make(map[string][]*models.GroupUser)
	return _dt["data"], xerror.Wrap(yqGet(t.c.R(), _GetGroupMembers(groupId), nil, &_dt), "YQGroup.GetGroupMembers")
}

func (t yqGroupImpl) GetGroupMembersByName(groupName string) (_ []*models.GroupUser, err error) {
	return t.GetGroupMembers(groupName)
}

func (t yqGroupImpl) UpdateGroupMember(groupId, username string) func(_ *models.GroupCreate) (_ *models.GroupUser, err error) {
	return func(data *models.GroupCreate) (_ *models.GroupUser, err error) {
		_dt := make(map[string]*models.GroupUser)
		return _dt["data"], xerror.Wrap(yqPut(t.c.R(), _UpdateGroupMember(groupId, username), data, &_dt), "YQGroup.UpdateGroupMember")
	}
}

func (t yqGroupImpl) UpdateGroupMemberByName(groupName, username string) func(_ *models.GroupCreate) (_ *models.GroupUser, err error) {
	return t.UpdateGroupMember(groupName, username)
}

func (t yqGroupImpl) DelGroupMember(groupId, username string) (_ *models.GroupUser, err error) {
	_dt := make(map[string]*models.GroupUser)
	return _dt["data"], xerror.Wrap(yqDelete(t.c.R(), _DelGroupMember(groupId, username), nil, &_dt), "YQGroup.DelGroupMember")
}

func (t yqGroupImpl) DelGroupMemberByName(groupName, username string) (_ *models.GroupUser, err error) {
	return t.DelGroupMember(groupName, username)
}

func (t yqGroupImpl) GetUserGroups(userId string) (res []*models.Group, err error) {
	_dt := make(map[string][]*models.Group)
	return _dt["data"], xerror.Wrap(yqGet(t.c.R(), _GetUserGroups(userId), nil, &_dt), "YQGroup.GetUserGroups")
}

func (t yqGroupImpl) GetUserGroupsByName(username string) (res []*models.Group, err error) {
	return t.GetUserGroups(username)
}

func (t yqGroupImpl) GetMyPubGroups() (res []*models.Group, err error) {
	_dt := make(map[string][]*models.Group)
	return _dt["data"], xerror.Wrap(yqGet(t.c.R(), _GetMyPubGroups(), nil, &_dt), "YQGroup.GetMyPubGroups")
}
