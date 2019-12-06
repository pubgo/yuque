package yuque

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/abc"
	"github.com/pubgo/yuque/models"
)

var _ abc.YuqueRepo = (*yqRepo)(nil)

type yqRepo struct {
	c *resty.Request
}

func (t *yqRepo) CreateUserRepo(userId string) func(data *models.BookCreate) (_ *models.Book, err error) {
	return func(data *models.BookCreate) (_ *models.Book, err error) {
		_dt := make(map[string]*models.Book)
		return _dt["data"], xerror.Wrap(yqPost(t.c, abc.CreateUserRepo(userId), data, &_dt), "yqRepo.CreateUserRepo")
	}
}

func (t *yqRepo) CreateUserRepoByName(username string) func(data *models.BookCreate) (_ *models.Book, err error) {
	return t.CreateUserRepo(username)
}

func (t *yqRepo) CreateGroupRepo(groupId string) func(data *models.BookCreate) (_ *models.Book, err error) {
	return func(data *models.BookCreate) (_ *models.Book, err error) {
		_dt := make(map[string]*models.Book)
		return _dt["data"], xerror.Wrap(yqPost(t.c, abc.CreateGroupRepo(groupId), data, &_dt), "yqRepo.CreateGroupRepo")
	}
}

func (t *yqRepo) CreateGroupRepoByName(groupName string) func(data *models.BookCreate) (_ *models.Book, err error) {
	return t.CreateGroupRepo(groupName)
}

func (t *yqRepo) GetRepo(repoId string) func(typ string) (_ *models.Book, err error) {
	return func(typ string) (_ *models.Book, err error) {
		_dt := make(map[string]*models.Book)
		return _dt["data"], xerror.Wrap(yqGet(t.c, abc.CreateGroupRepo(repoId), map[string]string{"type": typ}, &_dt), "yqRepo.GetRepo")
	}
}

func (t *yqRepo) GetRepoByName(repoName string) func(string) (_ *models.Book, err error) {
	return t.GetRepo(repoName)
}

func (t *yqRepo) UpdateRepo(repoId string) func(*models.BookCreate) (_ *models.Book, err error) {
	return func(data *models.BookCreate) (_ *models.Book, err error) {
		_dt := make(map[string]*models.Book)
		return _dt["data"], xerror.Wrap(yqPut(t.c, abc.UpdateRepo(repoId), data, &_dt), "yqRepo.UpdateRepo")
	}
}

func (t *yqRepo) UpdateRepoByName(repoName string) func(*models.BookCreate) (_ *models.Book, err error) {
	return t.UpdateRepo(repoName)
}

func (t *yqRepo) DelRepo(repoId string) (_ *models.Book, err error) {
	_dt := make(map[string]*models.Book)
	return _dt["data"], xerror.Wrap(yqDelete(t.c, abc.DelRepo(repoId), nil, &_dt), "yqRepo.DelRepo")
}

func (t *yqRepo) DelByName(repoName string) (_ *models.Book, err error) {
	return t.DelRepo(repoName)
}

func (t *yqRepo) GetRepoToc(repoId string) (_ *models.BookToc, err error) {
	_dt := make(map[string]*models.BookToc)
	return _dt["data"], xerror.Wrap(yqGet(t.c, abc.GetRepoToc(repoId), nil, &_dt), "yqRepo.GetRepoToc")
}

func (t *yqRepo) GetRepoTocByName(repoName string) (_ *models.BookToc, err error) {
	return t.GetRepoToc(repoName)
}

func (t *yqRepo) SearchRepo(q, typ string) (_ *models.Search, err error) {
	_dt := make(map[string]*models.Search)
	return _dt["data"], xerror.Wrap(yqGet(t.c, abc.SearchRepo(), map[string]string{"q": q, "type": typ}, &_dt), "yqRepo.SearchRepo")
}

func (t *yqRepo) GetUserRepos(userId string) func(typ string, offset int) (_ []*models.Book, err error) {
	return func(typ string, offset int) (_ []*models.Book, err error) {
		_dt := make(map[string][]*models.Book)
		return _dt["data"], xerror.Wrap(yqGet(t.c, abc.GetUserRepos(userId), map[string]string{
			"type":   typ,
			"offset": fmt.Sprintf("%d", offset),
		}, &_dt), "yqRepo.GetUserRepos")
	}
}

func (t *yqRepo) GetUserReposByName(username string) func(typ string, offset int) (_ []*models.Book, err error) {
	return t.GetUserRepos(username)
}

func (t *yqRepo) GetGroupRepos(groupId string) func(typ string, offset int) (_ []*models.Book, err error) {
	return func(typ string, offset int) (_ []*models.Book, err error) {
		_dt := make(map[string][]*models.Book)
		return _dt["data"], xerror.Wrap(yqGet(t.c, abc.GetGroupRepos(groupId), map[string]string{
			"type":   typ,
			"offset": fmt.Sprintf("%d", offset),
		}, &_dt), "yqRepo.GetGroupRepos")
	}
}

func (t *yqRepo) GetGroupReposByName(groupName string) func(string, int) (_ []*models.Book, err error) {
	return t.GetGroupRepos(groupName)
}
