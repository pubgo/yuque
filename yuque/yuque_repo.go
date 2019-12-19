package yuque

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/yuque/abc"
	"github.com/pubgo/yuque/yuque/models"
)

var (
	_                abc.YuQueRepo = (*yqRepoImpl)(nil)
	_GetUserRepos                  = _url("/users/%s/repos")
	_GetGroupRepos                 = _url("/groups/%s/repos")
	_CreateGroupRepo               = _url("/groups/%s/repos")
	_CreateUserRepo                = _url("/users/%s/repos")
	_GetRepo                       = _url("/repos/%s")
	_UpdateRepo                    = _url("/repos/%s")
	_DelRepo                       = _url("/repos/%s")
	_GetRepoToc                    = _url("/repos/%s/toc")
	_SearchRepo                    = _url("/search/repos?q=%s&type=%s")
)

type yqRepoImpl struct {
	c *resty.Client
}

func (t yqRepoImpl) CreateUserRepo(userId string) func(data *models.BookCreate) (_ *models.Book, err error) {
	return func(data *models.BookCreate) (_ *models.Book, err error) {
		_dt := make(map[string]*models.Book)
		return _dt["data"], xerror.Wrap(yqPost(t.c.R(), _CreateUserRepo(userId), data, &_dt), "YQRepo.CreateUserRepo")
	}
}

func (t yqRepoImpl) CreateUserRepoByName(username string) func(data *models.BookCreate) (_ *models.Book, err error) {
	return t.CreateUserRepo(username)
}

func (t yqRepoImpl) CreateGroupRepo(groupId string) func(data *models.BookCreate) (_ *models.Book, err error) {
	return func(data *models.BookCreate) (_ *models.Book, err error) {
		_dt := make(map[string]*models.Book)
		return _dt["data"], xerror.Wrap(yqPost(t.c.R(), _CreateGroupRepo(groupId), data, &_dt), "YQRepo.CreateGroupRepo")
	}
}

func (t yqRepoImpl) CreateGroupRepoByName(groupName string) func(data *models.BookCreate) (_ *models.Book, err error) {
	return t.CreateGroupRepo(groupName)
}

func (t yqRepoImpl) GetRepo(repoId string) func(typ string) (_ *models.Book, err error) {
	return func(typ string) (_ *models.Book, err error) {
		_dt := make(map[string]*models.Book)
		return _dt["data"], xerror.Wrap(yqGet(t.c.R(), _GetRepo(repoId), map[string]string{"type": typ}, &_dt), "YQRepo.GetRepo")
	}
}

func (t yqRepoImpl) GetRepoByName(repoName string) func(string) (_ *models.Book, err error) {
	return t.GetRepo(repoName)
}

func (t yqRepoImpl) UpdateRepo(repoId string) func(*models.BookCreate) (_ *models.Book, err error) {
	return func(data *models.BookCreate) (_ *models.Book, err error) {
		_dt := make(map[string]*models.Book)
		return _dt["data"], xerror.Wrap(yqPut(t.c.R(), _UpdateRepo(repoId), data, &_dt), "YQRepo.UpdateRepo")
	}
}

func (t yqRepoImpl) UpdateRepoByName(repoName string) func(*models.BookCreate) (_ *models.Book, err error) {
	return t.UpdateRepo(repoName)
}

func (t yqRepoImpl) DelRepo(repoId string) (_ *models.Book, err error) {
	_dt := make(map[string]*models.Book)
	return _dt["data"], xerror.Wrap(yqDelete(t.c.R(), _DelRepo(repoId), nil, &_dt), "YQRepo.DelRepo")
}

func (t yqRepoImpl) DelByName(repoName string) (_ *models.Book, err error) {
	return t.DelRepo(repoName)
}

func (t yqRepoImpl) GetRepoToc(repoId string) (_ *models.BookToc, err error) {
	_dt := make(map[string]*models.BookToc)
	return _dt["data"], xerror.Wrap(yqGet(t.c.R(), _GetRepoToc(repoId), nil, &_dt), "YQRepo.GetRepoToc")
}

func (t yqRepoImpl) GetRepoTocByName(repoName string) (_ *models.BookToc, err error) {
	return t.GetRepoToc(repoName)
}

func (t yqRepoImpl) SearchRepo(q, typ string) (_ *models.Search, err error) {
	_dt := make(map[string]*models.Search)
	return _dt["data"], xerror.Wrap(yqGet(t.c.R(), _SearchRepo(), map[string]string{"q": q, "type": typ}, &_dt), "YQRepo.SearchRepo")
}

func (t yqRepoImpl) GetUserRepos(userId string) func(typ string, offset int) (_ []*models.Book, err error) {
	return func(typ string, offset int) (_ []*models.Book, err error) {
		_dt := make(map[string][]*models.Book)
		return _dt["data"], xerror.Wrap(yqGet(t.c.R(), _GetUserRepos(userId), map[string]string{
			"type":   typ,
			"offset": fmt.Sprintf("%d", offset),
		}, &_dt), "YQRepo.GetUserRepos")
	}
}

func (t yqRepoImpl) GetUserReposByName(username string) func(typ string, offset int) (_ []*models.Book, err error) {
	return t.GetUserRepos(username)
}

func (t yqRepoImpl) GetGroupRepos(groupId string) func(typ string, offset int) (_ []*models.Book, err error) {
	return func(typ string, offset int) (_ []*models.Book, err error) {
		_dt := make(map[string][]*models.Book)
		return _dt["data"], xerror.Wrap(yqGet(t.c.R(), _GetGroupRepos(groupId), map[string]string{
			"type":   typ,
			"offset": fmt.Sprintf("%d", offset),
		}, &_dt), "YQRepo.GetGroupRepos")
	}
}

func (t yqRepoImpl) GetGroupReposByName(groupName string) func(string, int) (_ []*models.Book, err error) {
	return t.GetGroupRepos(groupName)
}
