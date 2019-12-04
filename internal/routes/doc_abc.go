package routes

import "github.com/pubgo/yuque/models"

// 获取一个仓库的文档列表
const GetDocs = apiPrefix + "/repos/{repo_id}/docs"
const GetDocsByName = apiPrefix + "/repos/{repo_name}/docs"

// 创建文档
const CreateDoc = apiPrefix + "/repos/{repo_id}/docs"
const CreateDocByName = apiPrefix + "/repos/{repo_name}/docs"

// 获取单篇文档的详细信息
const GetDoc = apiPrefix + "/repos/{repo_id}/docs/{doc_id}"
const GetDocByName = apiPrefix + "/repos/{repo_name}/docs/{slug}"

// 更新文档
const UpdateDoc = apiPrefix + "/repos/{repo_id}/docs/{doc_id}"
const UpdateDocByName = apiPrefix + "/repos/{repo_name}/docs/{doc_id}"

// 删除文档
const DelDoc = apiPrefix + "/repos/{repo_id}/docs/{doc_id}"
const DelDocByName = apiPrefix + "/repos/{repo_name}/docs/{doc_id}"

type YuqueDoc interface {
	GetDocs(RepoId string) ([]*models.Doc, error)
	GetDocsByName(RepoName string) ([]*models.Doc, error)
	CreateDoc(RepoId string) func(create *models.DocCreate) (*models.Doc, error)
	CreateDocByName(RepoName string) func(create *models.DocCreate) (*models.Doc, error)
	UpdateDoc(RepoId string, DocId string) func(create *models.DocCreate) (*models.Doc, error)
	UpdateDocByName(RepoName string, DocId string) func(create *models.DocCreate) (*models.Doc, error)
	GetDoc(RepoId string, DocId string) (*models.Doc, error)
	GetDocByName(RepoName string, DocId string) (*models.Doc, error)
	DelDoc(RepoId string, DocId string) (*models.Doc, error)
	DelDocByName(RepoName string, DocId string) (*models.Doc, error)
}
