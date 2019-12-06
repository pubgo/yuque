// doc https://www.yuque.com/yuque/developer/repo

package abc

import "github.com/pubgo/yuque/models"

type YuqueRepo interface {
	GetUserRepos(userId string) func(string, int) (_ []*models.Book, err error)
	GetUserReposByName(username string) func(string, int) (_ []*models.Book, err error)
	GetGroupRepos(groupId string) func(string, int) (_ []*models.Book, err error)
	GetGroupReposByName(groupName string) func(string, int) (_ []*models.Book, err error)
	CreateUserRepo(userId string) func(data *models.BookCreate) (_ *models.Book, err error)
	CreateUserRepoByName(username string) func(data *models.BookCreate) (_ *models.Book, err error)
	CreateGroupRepo(groupId string) func(data *models.BookCreate) (_ *models.Book, err error)
	CreateGroupRepoByName(groupName string) func(data *models.BookCreate) (_ *models.Book, err error)
	GetRepo(repoId string) func(string) (_ *models.Book, err error)
	GetRepoByName(repoName string) func(string) (_ *models.Book, err error)
	UpdateRepo(repoId string) func(*models.BookCreate) (_ *models.Book, err error)
	UpdateRepoByName(repoName string) func(*models.BookCreate) (_ *models.Book, err error)
	DelRepo(repoId string) (_ *models.Book, err error)
	DelByName(repoName string) (_ *models.Book, err error)
	GetRepoToc(repoId string) (_ *models.BookToc, err error)
	GetRepoTocByName(repoName string) (_ *models.BookToc, err error)
	SearchRepo(q, typ string) (_ *models.Search, err error)
}

// 获取一个仓库的文档列表
var GetUserRepos = _url("/users/%s/repos")
var GetGroupRepos = _url("/groups/%s/repos")

// 创建新仓库
// 往组织创建仓库
var CreateGroupRepo = _url("/groups/%s/repos")

// 往自己下面创建仓库
// 需要 Group 的 abilities.repo.create 权限
var CreateUserRepo = _url("/users/%s/repos")

// 获取仓库详情
var GetRepo = _url("/repos/%s")

// 更新仓库信息
// 需要 Repo 的 abilities.update 权限
var UpdateRepo = _url("/repos/%s")

// 删除仓库
// 需要 Repo 的 abilities.destroy 权限
var DelRepo = _url("/repos/%s")

// 获取一个仓库的目录结构
// 需要 Repo 的 abilities.read 权限
// 如果是 Group 成员，将能获取到私密文档、未发布的草稿
var GetRepoToc = _url("/repos/%s/toc")

// 基于关键字搜索仓库
// 无法搜索到私密仓库
var SearchRepo = _url("/search/repos?q=%s&type=%s")
