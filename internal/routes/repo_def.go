package routes

import "github.com/pubgo/yuque/models"

// 获取某个用户/组织的仓库列表
const GetUserRepos = apiProxy + "/users/{user_id}/repos"
const GetUserReposByName = apiProxy + "/users/{username}/repos"
const GetGroupRepos = apiProxy + "/groups/{group_id}/repos"
const GetGroupReposByName = apiProxy + "/groups/{group_name}/repos"

// 创建新仓库
// 往组织创建仓库
const CreateGroupRepo = apiProxy + "/groups/{group_id}/repos"
const CreateGroupRepoByName = apiProxy + "/groups/{}/repos"

// 往自己下面创建仓库
const CreateUserRepo = apiProxy + "/users/{user_id}/repos"
const CreateUserRepoByName = apiProxy + "/users/{username}/repos"

// 获取仓库详情
const GetRepo = apiProxy + "/repos/{repo_id}"
const GetRepoByName = apiProxy + "/repos/{repo_name}"

// 更新仓库信息
const UpdateRepo = apiProxy + "/repos/{repo_id}"
const UpdateRepoByName = apiProxy + "/repos/{repo_name}"

// 删除仓库
const DelRepo = apiProxy + "/repos/{repo_id}"
const DelRepoByName = apiProxy + "/repos/{repo_name}"

// 获取一个仓库的目录结构
const GetRepoToc = apiProxy + "/repos/{repo_id}/toc"
const GetRepoTocByName = apiProxy + "/repos/{repo_name}/toc"

// 基于关键字搜索仓库
const SearchRepo = apiProxy + "/search/repos"

type YuqueRepo interface {
	GetUserRepos(userId string) func(string, int) []*models.Book
	GetUserReposByName(username string) func(string, int) []*models.Book
	GetGroupRepos(groupId string) func(string, int) []*models.Book
	GetGroupReposByName(groupName string) func(string, int) []*models.Book
	CreateUserRepo(userId string) *models.Book
	CreateUserRepoByName(username string) *models.Book
	CreateGroupRepo(groupId string) *models.Book
	CreateGroupRepoByName(groupName string) *models.Book
	GetRepo(repoId string) func(string) *models.Book
	GetRepoByName(repoName string) func(string) *models.Book
	UpdateRepo(repoId string) func(*models.BookCreate) *models.Book
	UpdateRepoByName(repoName string) func(*models.BookCreate) *models.Book
	DelRepo(repoId string) *models.Book
	DelByName(repoName string) *models.Book
	GetRepoToc(repoId string) *models.BookToc
	GetRepoTocByName(repoName string) *models.BookToc
	SearchRepo() func(string, string)
}
