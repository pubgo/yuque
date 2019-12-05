// doc https://www.yuque.com/yuque/developer/group

package abc

import (
	"github.com/pubgo/yuque/models"
)

type YuqueGroup interface {
	GetUserGroups(userId string) (_ []*models.Group, err error)
	GetUserGroupsByName(username string) (_ []*models.Group, err error)
	GetMyPubGroups() (_ []*models.Group, err error)
	CreateGroup(group *models.GroupCreate) (_ *models.GroupDetail, err error)
	GetGroup(groupId string) (_ *models.GroupDetail, err error)
	GetGroupByName(groupName string) (_ *models.GroupDetail, err error)
	UpdateGroup(groupId string) func(group *models.GroupCreate) (_ *models.GroupDetail, err error)
	UpdateGroupByName(groupName string) func(group *models.GroupCreate) (_ *models.GroupDetail, err error)
	DelGroup(groupId string) (_ *models.GroupDetail, err error)
	DelGroupByName(groupName string) (_ *models.GroupDetail, err error)
	GetGroupMembers(groupId string) (_ []*models.GroupUser, err error)
	GetGroupMembersByName(groupName string) (_ []*models.GroupUser, err error)
	UpdateGroupMember(groupId, userName string) func(_ *models.GroupCreate) (_ *models.GroupUser, err error)
	UpdateGroupMemberByName(groupName, userName string) func(_ *models.GroupCreate) (_ *models.GroupUser, err error)
	DelGroupMember(groupId, userName string) (_ *models.GroupUser, err error)
	DelGroupMemberByName(groupName, userName string) (_ *models.GroupUser, err error)
}

// 获取某个用户的加入的组织列表
var GetUserGroups = _url("/users/%s/groups")
var GetUserGroupsByName = _url("/users/%s/groups")

// 获取公开组织列表
var GetMyPubGroups = _url("/groups")

// 创建组织
var CreateGroup = _url("/groups")

// 获取单个组织的详细信息
var GetGroup = _url("/groups/%s")
var GetGroupByName = _url("/groups/%s")

// 更新单个组织的详细信息
var UpdateGroup = _url("/groups/%s)")
var UpdateGroupByName = _url("/groups/%s)")

// 删除组织
// 此接口仅会删除 Group 基本信息，其他有关的仓库、文档、画板均保持不动，以避免误删除以后需要恢复
var DelGroup = _url("/groups/%s)")
var DelGroupByName = _url("/groups/%s)")

// 获取组织成员信息
// 需要 abilities.group.read 权限
var GetGroupMembers = _url("/groups/%s/users")
var GetGroupMembersByName = _url("/groups/%s/users")

// 增加或更新组织成员
// 需要 abilities.group_user.read 权限
// 需要有 abilities.group_user.create 权限
// role 为 0 的时候，需要 abilities.group_user.update 权限
var UpdateGroupMember = _url("/groups/%s/users/%s")
var UpdateGroupMemberByName = "/groups/%s/users/%s"

// 删除组织成员
// 需要有 abilities.group_user.destroy 权限
// 不可以自己删除自己
var DelGroupMember = _url("/groups/%s/users/%s")
var DelGroupMemberByName = _url("/groups/%s/users/%s")
