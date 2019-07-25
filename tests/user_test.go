package tests

import (
	"github.com/pubgo/errors"
	"github.com/pubgo/yuque"
	"os"
	"testing"
)

func init() {
	errors.Wrap(os.Setenv("debug", "true"), "set debug env error")
	errors.LoadEnvFile("../.env")
}

func TestUser(t *testing.T) {
	defer errors.Assert()

	yq := yuque.Default()
	yq.AddAuth("test", os.Getenv("token"))
	errors.P("get me", yq.R("test").GetMe())
}
