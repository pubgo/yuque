// doc https://www.yuque.com/yuque/developer/doc

package abc

import "github.com/pubgo/yuque/models"

type YuqueDoc interface {
	// 获取一个仓库的文档列表
	GetDocs(RepoId string) (_ []*models.Doc, err error)
	GetDocsByName(RepoName string) (_ []*models.Doc, err error)
	// 创建文档
	CreateDoc(RepoId string) func(create *models.DocCreate) (_ *models.Doc, err error)
	CreateDocByName(RepoName string) func(create *models.DocCreate) (_ *models.Doc, err error)
	// 更新文档
	UpdateDoc(RepoId string, DocId string) func(create *models.DocCreate) (_ *models.Doc, err error)
	UpdateDocByName(RepoName string, DocId string) func(create *models.DocCreate) (_ *models.Doc, err error)
	// 获取单篇文档的详细信息
	GetDoc(RepoId string, DocId string) (_ *models.Doc, err error)
	GetDocByName(RepoName string, DocId string) (_ *models.Doc, err error)
	// 删除文档
	DelDoc(RepoId string, DocId string) (_ *models.Doc, err error)
	DelDocByName(RepoName string, DocId string) (_ *models.Doc, err error)
}
