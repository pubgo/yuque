package tests

import (
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/yuque"
)

var yq *yuque.YuQue
var username = "barry.bai"
var userId = "253323"

func init() {
	defer xerror.Assert()
	xerror.Panic(xenv.LoadFile("../.env"))

	yq = yuque.New()
	yq.AddAuth(xenv.GetEnv("token"))
}
