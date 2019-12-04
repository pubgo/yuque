package tests

import (
	"github.com/pubgo/g/logs"
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque"
	"os"
	"testing"
)

func init() {
	xerror.Panic(xenv.SetDebug())
	xerror.Panic(xenv.LoadFile("../.env"))
}

func TestUser(t *testing.T) {
	defer xerror.Assert()

	yq := yuque.Default()
	yq.AddAuth("test", os.Getenv("token"))

	c, err := yq.R("test")
	xerror.Panic(err)

	m, err := c.GetMe()
	xerror.Panic(err)
	logs.Debug(m)
}
