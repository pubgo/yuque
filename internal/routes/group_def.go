package routes

import (
	"github.com/pubgo/yuque/models"
)

// 获取某个用户的加入的组织列表
const GetUserGroups = apiProxy + "/users/{user_id}/groups"
const GetUserGroupsByName = apiProxy + "/users/{username}/groups"

// 获取公开组织列表
const GetMyPubGroups = apiProxy + "/groups"

// 创建组织
const CreateGroup = apiProxy + "/groups"

// 获取单个组织的详细信息
const GetGroup = apiProxy + "/groups/{group_id}"
const GetGroupByName = apiProxy + "/groups/{group_name}"

// 更新单个组织的详细信息
const UpdateGroup = apiProxy + "/groups/{group_id}"
const UpdateGroupByName = apiProxy + "/groups/{group_name}"

// 删除组织
// 此接口仅会删除 Group 基本信息，其他有关的仓库、文档、画板均保持不动，以避免误删除以后需要恢复
const DelGroup = apiProxy + "/groups/{group_id}"
const DelGroupByName = apiProxy + "/groups/{group_name}"

// 获取组织成员信息
const GetGroupMembers = apiProxy + "/groups/{group_id}/users"
const GetGroupMembersByName = apiProxy + "/groups/{group_name}/users"

// 增加或更新组织成员
const UpdateGroupMember = apiProxy + "/groups/{group_id}/users/{username}"
const UpdateGroupMemberByName = apiProxy + "/groups/{group_name}/users/{username}"

// 删除组织成员
const DelGroupMember = apiProxy + "/groups/{group_id}/users/{username}"
const DelGroupMemberByName = apiProxy + "/groups/{group_name}/users/{username}"

type YuqueGroup interface {
	GetUserGroups(userId string) (res []*models.Group)
	GetUserGroupsByName(username string) (res []*models.Group)
	GetMyPubGroups() []*models.Group
	CreateGroup(group *models.GroupCreate) *models.GroupDetail
	GetGroup(groupId string) *models.GroupDetail
	GetGroupByName(groupName string) *models.GroupDetail
	UpdateGroup(groupId string) func(group *models.GroupCreate) *models.GroupDetail
	UpdateGroupByName(groupName string) func(group *models.GroupCreate) *models.GroupDetail
	DelGroup(groupId string) *models.GroupDetail
	DelGroupByName(groupName string) *models.GroupDetail
	GetGroupMembers(groupId string) []*models.GroupUser
	GetGroupMembersByName(groupName string) []*models.GroupUser
	UpdateGroupMember(groupId, userName string) func(*models.GroupCreate) *models.GroupUser
	UpdateGroupMemberByName(groupName, userName string) func(*models.GroupCreate) *models.GroupUser
	DelGroupMember(groupId, userName string) *models.GroupUser
	DelGroupMemberByName(groupName, userName string) *models.GroupUser
}
