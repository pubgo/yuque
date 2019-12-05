package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/yuque/abc"
	"github.com/pubgo/yuque/models"
)

var _ abc.YuqueRepo = &yqRepo{}

type yqRepo struct {
	c *resty.Request
}

func (t *yqRepo) GetUserRepos(userId string) func(string, int) (_ []*models.Book, err error) {
	panic("implement me")
}

func (t *yqRepo) GetUserReposByName(username string) func(string, int) (_ []*models.Book, err error) {
	panic("implement me")
}

func (t *yqRepo) GetGroupRepos(groupId string) func(string, int) (_ []*models.Book, err error) {
	panic("implement me")
}

func (t *yqRepo) GetGroupReposByName(groupName string) func(string, int) (_ []*models.Book, err error) {
	panic("implement me")
}

func (t *yqRepo) CreateUserRepo(userId string) (_ *models.Book, err error) {
	panic("implement me")
}

func (t *yqRepo) CreateUserRepoByName(username string) (_ *models.Book, err error) {
	panic("implement me")
}

func (t *yqRepo) CreateGroupRepo(groupId string) (_ *models.Book, err error) {
	panic("implement me")
}

func (t *yqRepo) CreateGroupRepoByName(groupName string) (_ *models.Book, err error) {
	panic("implement me")
}

func (t *yqRepo) GetRepo(repoId string) func(string) (_ *models.Book, err error) {
	panic("implement me")
}

func (t *yqRepo) GetRepoByName(repoName string) func(string) (_ *models.Book, err error) {
	panic("implement me")
}

func (t *yqRepo) UpdateRepo(repoId string) func(*models.BookCreate) (_ *models.Book, err error) {
	panic("implement me")
}

func (t *yqRepo) UpdateRepoByName(repoName string) func(*models.BookCreate) (_ *models.Book, err error) {
	panic("implement me")
}

func (t *yqRepo) DelRepo(repoId string) (_ *models.Book, err error) {
	panic("implement me")
}

func (t *yqRepo) DelByName(repoName string) (_ *models.Book, err error) {
	panic("implement me")
}

func (t *yqRepo) GetRepoToc(repoId string) (_ *models.BookToc, err error) {
	panic("implement me")
}

func (t *yqRepo) GetRepoTocByName(repoName string) (_ *models.BookToc, err error) {
	panic("implement me")
}

func (t *yqRepo) SearchRepo() func(string, string) (err error) {
	panic("implement me")
}
