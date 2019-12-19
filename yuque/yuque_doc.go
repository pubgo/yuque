package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/yuque/abc"
	"github.com/pubgo/yuque/yuque/models"
)

var (
	_          abc.YuQueDoc = (*yqDocImpl)(nil)
	_GetDocs                = _url("/repos/%s/docs")
	_CreateDoc              = _url("/repos/%s/docs")
	_GetDoc                 = _url("/repos/%s/docs/%s")
	_UpdateDoc              = _url("/repos/%s/docs/%s")
	_DelDoc                 = _url("/repos/%s/docs/%s")
)

type yqDocImpl struct {
	c *resty.Client
}

func (t yqDocImpl) GetDocs(RepoId string) (_ []*models.Doc, err error) {
	_dt := make(map[string][]*models.Doc)
	return _dt["data"], xerror.Wrap(yqGet(t.c.R(), _GetDocs(RepoId), nil, &_dt), "yqRepo.GetDocs")
}

func (t yqDocImpl) GetDocsByName(RepoName string) (_ []*models.Doc, err error) {
	return t.GetDocs(RepoName)
}

func (t yqDocImpl) CreateDoc(RepoId string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	return func(data *models.DocCreate) (_ *models.Doc, err error) {
		_dt := make(map[string]*models.Doc)
		return _dt["data"], xerror.Wrap(yqPost(t.c.R(), _CreateDoc(RepoId), data, &_dt), "yqRepo.CreateDoc")
	}
}

func (t yqDocImpl) CreateDocByName(RepoName string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	return t.CreateDoc(RepoName)
}

func (t yqDocImpl) UpdateDoc(RepoId string, DocId string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	return func(data *models.DocCreate) (_ *models.Doc, err error) {
		_dt := make(map[string]*models.Doc)
		return _dt["data"], xerror.Wrap(yqPut(t.c.R(), _UpdateDoc(RepoId, DocId), data, &_dt), "yqRepo.UpdateDoc")
	}
}

func (t yqDocImpl) UpdateDocByName(RepoName string, DocId string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	return t.UpdateDoc(RepoName, DocId)
}

func (t yqDocImpl) GetDoc(RepoId string, DocId string) (_ *models.Doc, err error) {
	_dt := make(map[string]*models.Doc)
	return _dt["data"], xerror.Wrap(yqGet(t.c.R(), _GetDoc(RepoId, DocId), nil, &_dt), "yqRepo.GetDoc")
}

func (t yqDocImpl) GetDocByName(RepoName string, DocId string) (_ *models.Doc, err error) {
	return t.GetDoc(RepoName, DocId)
}

func (t yqDocImpl) DelDoc(RepoId string, DocId string) (_ *models.Doc, err error) {
	_dt := make(map[string]*models.Doc)
	return _dt["data"], xerror.Wrap(yqDelete(t.c.R(), _DelDoc(RepoId, DocId), nil, &_dt), "yqRepo.DelDoc")
}

func (t yqDocImpl) DelDocByName(RepoName string, DocId string) (_ *models.Doc, err error) {
	return t.DelDoc(RepoName, DocId)
}
