package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/abc"
	"github.com/pubgo/yuque/models"
)

var _ abc.YuqueDoc = (*yqDoc)(nil)

type yqDoc struct {
	c *resty.Request
}

func (t *yqDoc) GetDocs(RepoId string) (_ []*models.Doc, err error) {
	_dt := make(map[string][]*models.Doc)
	return _dt["data"], xerror.Wrap(yqGet(t.c, abc.GetDocs(RepoId), nil, &_dt), "yqRepo.GetDocs")
}

func (t *yqDoc) GetDocsByName(RepoName string) (_ []*models.Doc, err error) {
	return t.GetDocs(RepoName)
}

func (t *yqDoc) CreateDoc(RepoId string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	return func(data *models.DocCreate) (_ *models.Doc, err error) {
		_dt := make(map[string]*models.Doc)
		return _dt["data"], xerror.Wrap(yqPost(t.c, abc.CreateDoc(RepoId), data, &_dt), "yqRepo.CreateDoc")
	}
}

func (t *yqDoc) CreateDocByName(RepoName string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	return t.CreateDoc(RepoName)
}

func (t *yqDoc) UpdateDoc(RepoId string, DocId string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	return func(data *models.DocCreate) (_ *models.Doc, err error) {
		_dt := make(map[string]*models.Doc)
		return _dt["data"], xerror.Wrap(yqPut(t.c, abc.CreateDoc(RepoId, DocId), data, &_dt), "yqRepo.UpdateDoc")
	}
}

func (t *yqDoc) UpdateDocByName(RepoName string, DocId string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	return t.UpdateDoc(RepoName, DocId)
}

func (t *yqDoc) GetDoc(RepoId string, DocId string) (_ *models.Doc, err error) {
	_dt := make(map[string]*models.Doc)
	return _dt["data"], xerror.Wrap(yqGet(t.c, abc.GetDoc(RepoId, DocId), nil, &_dt), "yqRepo.GetDoc")
}

func (t *yqDoc) GetDocByName(RepoName string, DocId string) (_ *models.Doc, err error) {
	return t.GetDoc(RepoName, DocId)
}

func (t *yqDoc) DelDoc(RepoId string, DocId string) (_ *models.Doc, err error) {
	_dt := make(map[string]*models.Doc)
	return _dt["data"], xerror.Wrap(yqDelete(t.c, abc.DelDoc(RepoId, DocId), nil, &_dt), "yqRepo.DelDoc")
}

func (t *yqDoc) DelDocByName(RepoName string, DocId string) (_ *models.Doc, err error) {
	return t.DelDoc(RepoName, DocId)
}
