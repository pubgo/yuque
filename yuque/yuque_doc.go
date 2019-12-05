package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/yuque/abc"
	"github.com/pubgo/yuque/models"
)

var _ abc.YuqueDoc = &yqDoc{}

type yqDoc struct {
	c *resty.Request
}

func (t *yqDoc) GetDocs(RepoId string) (_ []*models.Doc, err error) {
	panic("implement me")
}

func (t *yqDoc) GetDocsByName(RepoName string) (_ []*models.Doc, err error) {
	panic("implement me")
}

func (t *yqDoc) CreateDoc(RepoId string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	panic("implement me")
}

func (t *yqDoc) CreateDocByName(RepoName string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	panic("implement me")
}

func (t *yqDoc) UpdateDoc(RepoId string, DocId string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	panic("implement me")
}

func (t *yqDoc) UpdateDocByName(RepoName string, DocId string) func(create *models.DocCreate) (_ *models.Doc, err error) {
	panic("implement me")
}

func (t *yqDoc) GetDoc(RepoId string, DocId string) (_ *models.Doc, err error) {
	panic("implement me")
}

func (t *yqDoc) GetDocByName(RepoName string, DocId string) (_ *models.Doc, err error) {
	panic("implement me")
}

func (t *yqDoc) DelDoc(RepoId string, DocId string) (_ *models.Doc, err error) {
	panic("implement me")
}

func (t *yqDoc) DelDocByName(RepoName string, DocId string) (_ *models.Doc, err error) {
	panic("implement me")
}

func (t *yqDoc) SearchRepo() func(string, string) (err error) {
	panic("implement me")
}
