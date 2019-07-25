package routes

import "github.com/pubgo/yuque/models"

// 获取一个仓库的文档列表
const GetDocs = apiProxy + "/repos/{repo_id}/docs"
const GetDocsByName = apiProxy + "/repos/{repo_name}/docs"

// 创建文档
const CreateDoc = apiProxy + "/repos/{repo_id}/docs"
const CreateDocByName = apiProxy + "/repos/{repo_name}/docs"

// 获取单篇文档的详细信息
const GetDoc = apiProxy + "/repos/{repo_id}/docs/{doc_id}"
const GetDocByName = apiProxy + "/repos/{repo_name}/docs/{slug}"

// 更新文档
const UpdateDoc = apiProxy + "/repos/{repo_id}/docs/{doc_id}"
const UpdateDocByName = apiProxy + "/repos/{repo_name}/docs/{doc_id}"

// 删除文档
const DelDoc = apiProxy + "/repos/{repo_id}/docs/{doc_id}"
const DelDocByName = apiProxy + "/repos/{repo_name}/docs/{doc_id}"

type YuqueDoc interface {
	GetDocs(RepoId string) []*models.Doc
	GetDocsByName(RepoName string) []*models.Doc
	CreateDoc(RepoId string) func(create *models.DocCreate) *models.Doc
	CreateDocByName(RepoName string) func(create *models.DocCreate) *models.Doc
	UpdateDoc(RepoId string, DocId string) func(create *models.DocCreate) *models.Doc
	UpdateDocByName(RepoName string, DocId string) func(create *models.DocCreate) *models.Doc
	GetDoc(RepoId string, DocId string) *models.Doc
	GetDocByName(RepoName string, DocId string) *models.Doc
	DelDoc(RepoId string, DocId string) *models.Doc
	DelDocByName(RepoName string, DocId string) *models.Doc
}
