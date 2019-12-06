// doc https://www.yuque.com/yuque/developer/api

package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/pkg"
	"github.com/pubgo/g/xdi"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/abc"
	"time"
)

type yqClient struct {
	abc.YuqueUser
	abc.YuqueGroup

	c *resty.Request
}

type YuQue struct {
	AppName          string
	RetryCount       int
	RetryWaitTime    int
	RetryMaxWaitTime int
	Timeout          int
	Debug            bool
	yqClient         *resty.Client
}

func (t *YuQue) Group() *yqGroup {
	xerror.PanicT(pkg.IsNone(t.yqClient), "yuque client is null")
	return &yqGroup{c: t.yqClient.R()}
}

func (t *YuQue) AddAuth(token string) {
	t.yqClient = resty.New().
		SetDebug(t.Debug).
		SetContentLength(true).
		SetHostURL("https://www.yuque.com").
		SetRetryCount(t.RetryCount).
		SetRetryWaitTime(time.Second * time.Duration(t.RetryWaitTime)).
		SetRetryMaxWaitTime(time.Second * time.Duration(t.RetryMaxWaitTime)).
		SetTimeout(time.Second * time.Duration(t.Timeout)).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"User-Agent":   t.AppName,
			"X-Auth-Token": token,
		})
}

func init() {
	xdi.InitProvide(func() *YuQue {
		return &YuQue{
			AppName:          "test",
			RetryCount:       3,
			RetryWaitTime:    5,
			RetryMaxWaitTime: 20,
			Timeout:          60,
			Debug:            true,
		}

	})
}
