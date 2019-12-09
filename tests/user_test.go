package tests

import (
	"fmt"
	"github.com/pubgo/g/xerror"
	"testing"
)

func TestGetMe(t *testing.T) {
	defer xerror.Assert()

	xerror.PanicErr(yq.User().GetMe())

}

func TestGetUser(t *testing.T) {
	defer xerror.Assert()

	xerror.PanicErr(yq.User().GetUser("253323"))
}

func TestGetUserReposByName(t *testing.T) {
	defer xerror.Assert()

	_, err := yq.User().GetUser("barry.me")
	fmt.Printf("%s", err)
	xerror.ErrHandle(err, func(err *xerror.Err) {
	})
}
