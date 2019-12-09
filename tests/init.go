package tests

import (
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/yuque"
)

var yq *yuque.YuQue

func init() {
	defer xerror.Assert()
	xerror.Panic(xenv.LoadFile("../.env"))

	yq = yuque.New()
	yq.Debug = xenv.IsDebug()
	yq.AddAuth(xenv.GetEnv("token"))
}
