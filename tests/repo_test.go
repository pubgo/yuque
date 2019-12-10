package tests

import (
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/yuque"
	"testing"
)

func TestGetUserRepos(t *testing.T) {
	defer xerror.Assert()
	//yuque.Debug(xerror.PanicErr(yq.Repo().GetUserRepos(username)("all", 0)))
	//yuque.Debug(xerror.PanicErr(yq.Repo().GetUserRepos("kooksee")("all", 0)))
	//yuque.Debug(xerror.PanicErr(yq.Repo().GetRepo("kooksee/readme")("Book")))
	yuque.Debug(xerror.PanicErr(yq.Doc().GetDoc("kooksee/readme","readme")))
}
