package tests

import (
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/yuque"
	"github.com/pubgo/yuque/yuque/abc"
)

var yq abc.IYuque
var username = "barry.bai"
var userId = "253323"

func init() {
	defer xerror.Assert()
	xerror.Panic(xenv.LoadFile("../.env"))

	yq = yuque.New()
	yq = yq.Auth(xenv.GetEnv("token"))
}
