package tests

import (
	"github.com/pubgo/g/xerror"
	"testing"
)

func TestGetMe(t *testing.T) {
	defer xerror.Assert()

	xerror.PanicErr(yq.User().GetMe())

}

func TestGetUser(t *testing.T) {
	defer xerror.Assert()

	xerror.PanicErr(yq.User().GetUser(userId))
}

func TestGetUserReposByName(t *testing.T) {
	defer xerror.Assert()

	//xerror.PanicErr(yq.User().GetUser("barry.me")).(*models.UserDetail)
	xerror.PanicErr(yq.User().GetUser(username))
}
