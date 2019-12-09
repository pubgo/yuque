// doc https://www.yuque.com/yuque/developer/api

package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"time"
)

type YuQue struct {
	AppName          string // default test
	RetryCount       int    // default 3
	RetryWaitTime    int    // default 5 second
	RetryMaxWaitTime int    // default 20 second
	Timeout          int    // default 60 second
	Debug            bool   // default true
	yqClient         *resty.Client
}

func (t *YuQue) Group() YQGroup {
	xerror.PanicT(t.yqClient == nil, "yuque client is null")
	return YQGroup{c: t.yqClient.R()}
}

func (t *YuQue) User() YQUser {
	xerror.PanicT(t.yqClient == nil, "yuque client is null")
	return YQUser{c: t.yqClient.R()}
}

func (t *YuQue) Repo() YQRepo {
	xerror.PanicT(t.yqClient == nil, "yuque client is null")
	return YQRepo{c: t.yqClient.R()}
}

func (t *YuQue) Doc() YQDoc {
	xerror.PanicT(t.yqClient == nil, "yuque client is null")
	return YQDoc{c: t.yqClient.R()}
}

func (t *YuQue) _init() {
	if t.AppName == "" {
		t.AppName = "test"
	}

	if t.RetryCount < 1 {
		t.RetryCount = 3
	}

	if t.RetryWaitTime < 1 {
		t.RetryWaitTime = 5
	}

	if t.RetryMaxWaitTime < 1 {
		t.RetryMaxWaitTime = 20
	}

	if t.Timeout < 1 {
		t.Timeout = 60
	}
}

func (t *YuQue) AddAuth(token string) {
	t._init()

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

func New() *YuQue {
	_yq := &YuQue{Debug: true}
	_yq._init()
	return _yq
}
