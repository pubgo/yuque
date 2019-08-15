package routes

import (
	"github.com/pubgo/yuque/models"
)

// 获取某个用户的加入的组织列表
const GetUserGroups = apiPrefix + "/users/{user_id}/groups"
const GetUserGroupsByName = apiPrefix + "/users/{username}/groups"

// 获取公开组织列表
const GetMyPubGroups = apiPrefix + "/groups"

// 创建组织
const CreateGroup = apiPrefix + "/groups"

// 获取单个组织的详细信息
const GetGroup = apiPrefix + "/groups/{group_id}"
const GetGroupByName = apiPrefix + "/groups/{group_name}"

// 更新单个组织的详细信息
const UpdateGroup = apiPrefix + "/groups/{group_id}"
const UpdateGroupByName = apiPrefix + "/groups/{group_name}"

// 删除组织
// 此接口仅会删除 Group 基本信息，其他有关的仓库、文档、画板均保持不动，以避免误删除以后需要恢复
const DelGroup = apiPrefix + "/groups/{group_id}"
const DelGroupByName = apiPrefix + "/groups/{group_name}"

// 获取组织成员信息
const GetGroupMembers = apiPrefix + "/groups/{group_id}/users"
const GetGroupMembersByName = apiPrefix + "/groups/{group_name}/users"

// 增加或更新组织成员
const UpdateGroupMember = apiPrefix + "/groups/{group_id}/users/{username}"
const UpdateGroupMemberByName = apiPrefix + "/groups/{group_name}/users/{username}"

// 删除组织成员
const DelGroupMember = apiPrefix + "/groups/{group_id}/users/{username}"
const DelGroupMemberByName = apiPrefix + "/groups/{group_name}/users/{username}"

type YuqueGroup interface {
	GetUserGroups(userId string) (res []*models.Group, err error)
	GetUserGroupsByName(username string) (res []*models.Group, err error)
	GetMyPubGroups() ([]*models.Group,error)
	CreateGroup(group *models.GroupCreate) (*models.GroupDetail, error)
	GetGroup(groupId string) (*models.GroupDetail, error)
	GetGroupByName(groupName string) (*models.GroupDetail, error)
	UpdateGroup(groupId string) func(group *models.GroupCreate) (*models.GroupDetail, error)
	UpdateGroupByName(groupName string) func(group *models.GroupCreate) (*models.GroupDetail, error)
	DelGroup(groupId string) (*models.GroupDetail, error)
	DelGroupByName(groupName string) (*models.GroupDetail, error)
	GetGroupMembers(groupId string) ([]*models.GroupUser, error)
	GetGroupMembersByName(groupName string) ([]*models.GroupUser, error)
	UpdateGroupMember(groupId, userName string) func(*models.GroupCreate) (*models.GroupUser, error)
	UpdateGroupMemberByName(groupName, userName string) func(*models.GroupCreate) (*models.GroupUser, error)
	DelGroupMember(groupId, userName string) (*models.GroupUser, error)
	DelGroupMemberByName(groupName, userName string) (*models.GroupUser, error)
}
