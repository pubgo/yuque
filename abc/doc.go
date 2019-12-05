package abc

import "github.com/pubgo/yuque/models"

type YuqueDoc interface {
	GetDocs(RepoId string) (_ []*models.Doc, err error)
	GetDocsByName(RepoName string) (_ []*models.Doc, err error)
	CreateDoc(RepoId string) func(create *models.DocCreate) (_ *models.Doc, err error)
	CreateDocByName(RepoName string) func(create *models.DocCreate) (_ *models.Doc, err error)
	UpdateDoc(RepoId string, DocId string) func(create *models.DocCreate) (_ *models.Doc, err error)
	UpdateDocByName(RepoName string, DocId string) func(create *models.DocCreate) (_ *models.Doc, err error)
	GetDoc(RepoId string, DocId string) (_ *models.Doc, err error)
	GetDocByName(RepoName string, DocId string) (_ *models.Doc, err error)
	DelDoc(RepoId string, DocId string) (_ *models.Doc, err error)
	DelDocByName(RepoName string, DocId string) (_ *models.Doc, err error)
}

// 获取一个仓库的文档列表
var GetDocs = _url("/repos/%s/docs")

// 创建文档
var CreateDoc = _url("/repos/%s/docs")

// 获取单篇文档的详细信息
var GetDoc = _url("/repos/%s/docs/%s")

// 更新文档
var UpdateDoc = _url("/repos/%s/docs/%s")

// 删除文档
var DelDoc = _url("/repos/%s/docs/%s")
