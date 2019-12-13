// doc https://www.yuque.com/yuque/developer/repo

package abc

import (
	"github.com/pubgo/yuque/yuque/models"
)

type YuQueRepo interface {
	// 获取一个仓库的文档列表
	GetUserRepos(userId string) func(string, int) (_ []*models.Book, err error)
	GetUserReposByName(username string) func(string, int) (_ []*models.Book, err error)
	GetGroupRepos(groupId string) func(string, int) (_ []*models.Book, err error)
	GetGroupReposByName(groupName string) func(string, int) (_ []*models.Book, err error)
	// 往自己下面创建仓库
	// 需要 Group 的 abilities.repo.create 权限
	CreateUserRepo(userId string) func(data *models.BookCreate) (_ *models.Book, err error)
	CreateUserRepoByName(username string) func(data *models.BookCreate) (_ *models.Book, err error)
	// 创建新仓库
	// 往组织创建仓库
	CreateGroupRepo(groupId string) func(data *models.BookCreate) (_ *models.Book, err error)
	CreateGroupRepoByName(groupName string) func(data *models.BookCreate) (_ *models.Book, err error)
	// 获取仓库详情
	GetRepo(repoId string) func(string) (_ *models.Book, err error)
	GetRepoByName(repoName string) func(string) (_ *models.Book, err error)
	// 更新仓库信息
	// 需要 Repo 的 abilities.update 权限
	UpdateRepo(repoId string) func(*models.BookCreate) (_ *models.Book, err error)
	UpdateRepoByName(repoName string) func(*models.BookCreate) (_ *models.Book, err error)
	// 删除仓库
	// 需要 Repo 的 abilities.destroy 权限
	DelRepo(repoId string) (_ *models.Book, err error)
	DelByName(repoName string) (_ *models.Book, err error)
	// 获取一个仓库的目录结构
	// 需要 Repo 的 abilities.read 权限
	// 如果是 Group 成员，将能获取到私密文档、未发布的草稿
	GetRepoToc(repoId string) (_ *models.BookToc, err error)
	GetRepoTocByName(repoName string) (_ *models.BookToc, err error)
	// 基于关键字搜索仓库
	// 无法搜索到私密仓库
	SearchRepo(q, typ string) (_ *models.Search, err error)
}
