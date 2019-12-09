package yuque

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/abc"
	"github.com/pubgo/yuque/models"
)

var _ abc.YuqueRepo = (*YQRepo)(nil)

type YQRepo struct {
	c *resty.Request
}

func (t YQRepo) CreateUserRepo(userId string) func(data *models.BookCreate) (_ *models.Book, err error) {
	return func(data *models.BookCreate) (_ *models.Book, err error) {
		_dt := make(map[string]*models.Book)
		return _dt["data"], xerror.Wrap(yqPost(t.c, abc.CreateUserRepo(userId), data, &_dt), "YQRepo.CreateUserRepo")
	}
}

func (t YQRepo) CreateUserRepoByName(username string) func(data *models.BookCreate) (_ *models.Book, err error) {
	return t.CreateUserRepo(username)
}

func (t YQRepo) CreateGroupRepo(groupId string) func(data *models.BookCreate) (_ *models.Book, err error) {
	return func(data *models.BookCreate) (_ *models.Book, err error) {
		_dt := make(map[string]*models.Book)
		return _dt["data"], xerror.Wrap(yqPost(t.c, abc.CreateGroupRepo(groupId), data, &_dt), "YQRepo.CreateGroupRepo")
	}
}

func (t YQRepo) CreateGroupRepoByName(groupName string) func(data *models.BookCreate) (_ *models.Book, err error) {
	return t.CreateGroupRepo(groupName)
}

func (t YQRepo) GetRepo(repoId string) func(typ string) (_ *models.Book, err error) {
	return func(typ string) (_ *models.Book, err error) {
		_dt := make(map[string]*models.Book)
		return _dt["data"], xerror.Wrap(yqGet(t.c, abc.CreateGroupRepo(repoId), map[string]string{"type": typ}, &_dt), "YQRepo.GetRepo")
	}
}

func (t YQRepo) GetRepoByName(repoName string) func(string) (_ *models.Book, err error) {
	return t.GetRepo(repoName)
}

func (t YQRepo) UpdateRepo(repoId string) func(*models.BookCreate) (_ *models.Book, err error) {
	return func(data *models.BookCreate) (_ *models.Book, err error) {
		_dt := make(map[string]*models.Book)
		return _dt["data"], xerror.Wrap(yqPut(t.c, abc.UpdateRepo(repoId), data, &_dt), "YQRepo.UpdateRepo")
	}
}

func (t YQRepo) UpdateRepoByName(repoName string) func(*models.BookCreate) (_ *models.Book, err error) {
	return t.UpdateRepo(repoName)
}

func (t YQRepo) DelRepo(repoId string) (_ *models.Book, err error) {
	_dt := make(map[string]*models.Book)
	return _dt["data"], xerror.Wrap(yqDelete(t.c, abc.DelRepo(repoId), nil, &_dt), "YQRepo.DelRepo")
}

func (t YQRepo) DelByName(repoName string) (_ *models.Book, err error) {
	return t.DelRepo(repoName)
}

func (t YQRepo) GetRepoToc(repoId string) (_ *models.BookToc, err error) {
	_dt := make(map[string]*models.BookToc)
	return _dt["data"], xerror.Wrap(yqGet(t.c, abc.GetRepoToc(repoId), nil, &_dt), "YQRepo.GetRepoToc")
}

func (t YQRepo) GetRepoTocByName(repoName string) (_ *models.BookToc, err error) {
	return t.GetRepoToc(repoName)
}

func (t YQRepo) SearchRepo(q, typ string) (_ *models.Search, err error) {
	_dt := make(map[string]*models.Search)
	return _dt["data"], xerror.Wrap(yqGet(t.c, abc.SearchRepo(), map[string]string{"q": q, "type": typ}, &_dt), "YQRepo.SearchRepo")
}

func (t YQRepo) GetUserRepos(userId string) func(typ string, offset int) (_ []*models.Book, err error) {
	return func(typ string, offset int) (_ []*models.Book, err error) {
		_dt := make(map[string][]*models.Book)
		return _dt["data"], xerror.Wrap(yqGet(t.c, abc.GetUserRepos(userId), map[string]string{
			"type":   typ,
			"offset": fmt.Sprintf("%d", offset),
		}, &_dt), "YQRepo.GetUserRepos")
	}
}

func (t YQRepo) GetUserReposByName(username string) func(typ string, offset int) (_ []*models.Book, err error) {
	return t.GetUserRepos(username)
}

func (t YQRepo) GetGroupRepos(groupId string) func(typ string, offset int) (_ []*models.Book, err error) {
	return func(typ string, offset int) (_ []*models.Book, err error) {
		_dt := make(map[string][]*models.Book)
		return _dt["data"], xerror.Wrap(yqGet(t.c, abc.GetGroupRepos(groupId), map[string]string{
			"type":   typ,
			"offset": fmt.Sprintf("%d", offset),
		}, &_dt), "YQRepo.GetGroupRepos")
	}
}

func (t YQRepo) GetGroupReposByName(groupName string) func(string, int) (_ []*models.Book, err error) {
	return t.GetGroupRepos(groupName)
}
