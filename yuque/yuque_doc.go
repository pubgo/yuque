package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/yuque/abc"
	"github.com/pubgo/yuque/yuque/models"
)

var (
	_          abc.YuQueDoc = (*YQDoc)(nil)
	_GetDocs                = _url("/repos/%s/docs")
	_CreateDoc              = _url("/repos/%s/docs")
	_GetDoc                 = _url("/repos/%s/docs/%s")
	_UpdateDoc              = _url("/repos/%s/docs/%s")
	_DelDoc                 = _url("/repos/%s/docs/%s")
)

type YQDoc struct {
	c *resty.Request
}

func (t YQDoc) GetDocs(RepoId string) (_ []*models.Doc, err error) {
	_dt := make(map[string][]*models.Doc)
	return _dt["data"], xerror.Wrap(yqGet(t.c, _GetDocs(RepoId), nil, &_dt), "yqRepo.GetDocs")
}

func (t YQDoc) GetDocsByName(RepoName string) (_ []*models.Doc, err error) {
	return t.GetDocs(RepoName)
}

func (t YQDoc) CreateDoc(RepoId string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	return func(data *models.DocCreate) (_ *models.Doc, err error) {
		_dt := make(map[string]*models.Doc)
		return _dt["data"], xerror.Wrap(yqPost(t.c, _CreateDoc(RepoId), data, &_dt), "yqRepo.CreateDoc")
	}
}

func (t YQDoc) CreateDocByName(RepoName string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	return t.CreateDoc(RepoName)
}

func (t YQDoc) UpdateDoc(RepoId string, DocId string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	return func(data *models.DocCreate) (_ *models.Doc, err error) {
		_dt := make(map[string]*models.Doc)
		return _dt["data"], xerror.Wrap(yqPut(t.c, _UpdateDoc(RepoId, DocId), data, &_dt), "yqRepo.UpdateDoc")
	}
}

func (t YQDoc) UpdateDocByName(RepoName string, DocId string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	return t.UpdateDoc(RepoName, DocId)
}

func (t YQDoc) GetDoc(RepoId string, DocId string) (_ *models.Doc, err error) {
	_dt := make(map[string]*models.Doc)
	return _dt["data"], xerror.Wrap(yqGet(t.c, _GetDoc(RepoId, DocId), nil, &_dt), "yqRepo.GetDoc")
}

func (t YQDoc) GetDocByName(RepoName string, DocId string) (_ *models.Doc, err error) {
	return t.GetDoc(RepoName, DocId)
}

func (t YQDoc) DelDoc(RepoId string, DocId string) (_ *models.Doc, err error) {
	_dt := make(map[string]*models.Doc)
	return _dt["data"], xerror.Wrap(yqDelete(t.c, _DelDoc(RepoId, DocId), nil, &_dt), "yqRepo.DelDoc")
}

func (t YQDoc) DelDocByName(RepoName string, DocId string) (_ *models.Doc, err error) {
	return t.DelDoc(RepoName, DocId)
}
