package tests

import (
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/yuque"
	"github.com/pubgo/yuque/yuque/models"
	"testing"
)

func TestGetUserGroups(t *testing.T) {
	defer xerror.Assert()

	xerror.PanicErr(yq.Group().GetUserGroups("kooksee"))
}

func TestGetMyPubGroups(t *testing.T) {
	defer xerror.Assert()
	xerror.PanicErr(yq.Group().GetMyPubGroups())
}

func TestGetGroup(t *testing.T) {
	defer xerror.Assert()
	yuque.Debug(xerror.PanicErr(yq.Group().GetGroup("kooksee")))
}

func TestCreateGroup(t *testing.T) {
	defer xerror.Resp(func(err xerror.IErr) {
		if err.Is(yuque.ErrParametersIncorrectOrMissing) {
			err.P()
		}
	})

	xerror.PanicErr(yq.Group().CreateGroup(&models.GroupCreate{
		Name:        "version111",
		Login:       username,
		Description: "测试",
	}))
}

func TestUpdateGroup(t *testing.T) {
	defer xerror.Assert()
	yuque.Debug(xerror.PanicErr(yq.Group().UpdateGroup("253324")(&models.GroupCreate{
		Name:        "kooksee",
		Login:       "kooksee",
		Description: "个人知识",
	})))
}
