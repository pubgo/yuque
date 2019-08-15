package tests

import (
	"github.com/pubgo/g/envs"
	"github.com/pubgo/g/errors"
	"github.com/pubgo/g/logs"
	"github.com/pubgo/yuque"
	"os"
	"testing"
)

func init() {
	errors.Panic(envs.SetDebug())
	errors.Panic(envs.LoadFile("../.env"))
}

func TestUser(t *testing.T) {
	defer errors.Assert()

	yq := yuque.Default()
	yq.AddAuth("test", os.Getenv("token"))

	c, err := yq.R("test")
	errors.Panic(err)

	m, err := c.GetMe()
	errors.Panic(err)
	logs.P("get me", m)
}
