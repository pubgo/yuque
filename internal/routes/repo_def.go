package routes

import "github.com/pubgo/yuque/models"

// 获取某个用户|组织的仓库列表
const GetUserRepos = apiPrefix + "/users/{user_id}/repos"
const GetUserReposByName = apiPrefix + "/users/{username}/repos"
const GetGroupRepos = apiPrefix + "/groups/{group_id}/repos"
const GetGroupReposByName = apiPrefix + "/groups/{group_name}/repos"

// 创建新仓库
// 往组织创建仓库
const CreateGroupRepo = apiPrefix + "/groups/{group_id}/repos"
const CreateGroupRepoByName = apiPrefix + "/groups/{}/repos"

// 往自己下面创建仓库
const CreateUserRepo = apiPrefix + "/users/{user_id}/repos"
const CreateUserRepoByName = apiPrefix + "/users/{username}/repos"

// 获取仓库详情
const GetRepo = apiPrefix + "/repos/{repo_id}"
const GetRepoByName = apiPrefix + "/repos/{repo_name}"

// 更新仓库信息
const UpdateRepo = apiPrefix + "/repos/{repo_id}"
const UpdateRepoByName = apiPrefix + "/repos/{repo_name}"

// 删除仓库
const DelRepo = apiPrefix + "/repos/{repo_id}"
const DelRepoByName = apiPrefix + "/repos/{repo_name}"

// 获取一个仓库的目录结构
const GetRepoToc = apiPrefix + "/repos/{repo_id}/toc"
const GetRepoTocByName = apiPrefix + "/repos/{repo_name}/toc"

// 基于关键字搜索仓库
const SearchRepo = apiPrefix + "/search/repos"

type YuqueRepo interface {
	GetUserRepos(userId string) func(string, int) ([]*models.Book, error)
	GetUserReposByName(username string) func(string, int) ([]*models.Book, error)
	GetGroupRepos(groupId string) func(string, int) ([]*models.Book, error)
	GetGroupReposByName(groupName string) func(string, int) ([]*models.Book, error)
	CreateUserRepo(userId string) (*models.Book, error)
	CreateUserRepoByName(username string) (*models.Book, error)
	CreateGroupRepo(groupId string) (*models.Book, error)
	CreateGroupRepoByName(groupName string) (*models.Book, error)
	GetRepo(repoId string) func(string) (*models.Book, error)
	GetRepoByName(repoName string) func(string) (*models.Book, error)
	UpdateRepo(repoId string) func(*models.BookCreate) (*models.Book, error)
	UpdateRepoByName(repoName string) func(*models.BookCreate) (*models.Book, error)
	DelRepo(repoId string) (*models.Book, error)
	DelByName(repoName string) (*models.Book, error)
	GetRepoToc(repoId string) (*models.BookToc, error)
	GetRepoTocByName(repoName string) (*models.BookToc, error)
	SearchRepo() func(string, string) error
}
