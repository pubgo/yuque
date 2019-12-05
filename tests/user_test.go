package tests

import (
	"github.com/pubgo/g/logs"
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/g/xinit"
	"github.com/pubgo/yuque/yuque"
	"testing"
)

var q *yuque.YuQue

func init() {
	defer xerror.Assert()
	xerror.Panic(xenv.SetDebug())
	xerror.Panic(xenv.LoadFile("../.env"))

	xinit.InitInvoke(func(_que *yuque.YuQue) {
		q.AddAuth(xenv.GetEnv("token"))
		q = _que
	})
}

func TestUser(t *testing.T) {
	defer xerror.Assert()

	xerror.Panic(xinit.Start())
	q.Group().CreateGroup()
	m, err := q.GetMe()
	xerror.Panic(err)
	logs.Debug(m)

}
