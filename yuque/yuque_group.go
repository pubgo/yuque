package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/yuque/abc"
	"github.com/pubgo/yuque/yuque/models"
)

var (
	_                  abc.YuQueGroup = (*YQGroup)(nil)
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

type YQGroup struct {
	c *resty.Request
}

func (t YQGroup) CreateGroup(data *models.GroupCreate) (_ *models.GroupDetail, err error) {
	_dt := make(map[string]*models.GroupDetail)
	return _dt["data"], xerror.Wrap(yqPost(t.c, _CreateGroup(), data, &_dt), "YQGroup.CreateGroup")
}

func (t YQGroup) GetGroup(groupId string) (_ *models.GroupDetail, err error) {
	_dt := make(map[string]*models.GroupDetail)
	return _dt["data"], xerror.Wrap(yqGet(t.c, _GetGroup(groupId), nil, &_dt), "YQGroup.GetGroup")
}

func (t YQGroup) GetGroupByName(groupName string) (_ *models.GroupDetail, err error) {
	return t.GetGroup(groupName)
}

func (t YQGroup) UpdateGroup(groupId string) func(group *models.GroupCreate) (_ *models.GroupDetail, err error) {
	return func(data *models.GroupCreate) (_ *models.GroupDetail, err error) {
		_dt := make(map[string]*models.GroupDetail)
		return _dt["data"], xerror.Wrap(yqPut(t.c, _UpdateGroup(groupId), data, &_dt), "YQGroup.UpdateGroup")
	}
}

func (t YQGroup) UpdateGroupByName(groupName string) func(group *models.GroupCreate) (_ *models.GroupDetail, err error) {
	return t.UpdateGroup(groupName)
}

func (t YQGroup) DelGroup(groupId string) (_ *models.GroupDetail, err error) {
	_dt := make(map[string]*models.GroupDetail)
	return _dt["data"], xerror.Wrap(yqDelete(t.c, _DelGroup(groupId), nil, &_dt), "YQGroup.DelGroup")
}

func (t YQGroup) DelGroupByName(groupName string) (_ *models.GroupDetail, err error) {
	return t.DelGroup(groupName)
}

func (t YQGroup) GetGroupMembers(groupId string) (_ []*models.GroupUser, err error) {
	_dt := make(map[string][]*models.GroupUser)
	return _dt["data"], xerror.Wrap(yqGet(t.c, _GetGroupMembers(groupId), nil, &_dt), "YQGroup.GetGroupMembers")
}

func (t YQGroup) GetGroupMembersByName(groupName string) (_ []*models.GroupUser, err error) {
	return t.GetGroupMembers(groupName)
}

func (t YQGroup) UpdateGroupMember(groupId, username string) func(_ *models.GroupCreate) (_ *models.GroupUser, err error) {
	return func(data *models.GroupCreate) (_ *models.GroupUser, err error) {
		_dt := make(map[string]*models.GroupUser)
		return _dt["data"], xerror.Wrap(yqPut(t.c, _UpdateGroupMember(groupId, username), data, &_dt), "YQGroup.UpdateGroupMember")
	}
}

func (t YQGroup) UpdateGroupMemberByName(groupName, username string) func(_ *models.GroupCreate) (_ *models.GroupUser, err error) {
	return t.UpdateGroupMember(groupName, username)
}

func (t YQGroup) DelGroupMember(groupId, username string) (_ *models.GroupUser, err error) {
	_dt := make(map[string]*models.GroupUser)
	return _dt["data"], xerror.Wrap(yqDelete(t.c, _DelGroupMember(groupId, username), nil, &_dt), "YQGroup.DelGroupMember")
}

func (t YQGroup) DelGroupMemberByName(groupName, username string) (_ *models.GroupUser, err error) {
	return t.DelGroupMember(groupName, username)
}

func (t YQGroup) GetUserGroups(userId string) (res []*models.Group, err error) {
	_dt := make(map[string][]*models.Group)
	return _dt["data"], xerror.Wrap(yqGet(t.c, _GetUserGroups(userId), nil, &_dt), "YQGroup.GetUserGroups")
}

func (t YQGroup) GetUserGroupsByName(username string) (res []*models.Group, err error) {
	return t.GetUserGroups(username)
}

func (t YQGroup) GetMyPubGroups() (res []*models.Group, err error) {
	_dt := make(map[string][]*models.Group)
	return _dt["data"], xerror.Wrap(yqGet(t.c, _GetMyPubGroups(), nil, &_dt), "YQGroup.GetMyPubGroups")
}
