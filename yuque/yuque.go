// doc https://www.yuque.com/yuque/developer/api

package yuque

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/yuque/yuque/abc"
	"log"
	"time"
)

func New() abc.IYuque {
	var _ abc.IYuque = (*yuque)(nil)

	_yq := &yuque{}
	_yq._init()
	return _yq
}

type yuque struct {
	AppName          string // default test
	RetryCount       int    // default 3
	RetryWaitTime    int    // default 5 second
	RetryMaxWaitTime int    // default 20 second
	Timeout          int    // default 60 second
	Debug            bool   // default true
	client           *resty.Client
}

func (t *yuque) checkClient() {
	if t.client == nil {
		log.Fatal("yuque client is null")
	}
}

func (t *yuque) Group() abc.YuQueGroup {
	t.checkClient()
	return &yqGroupImpl{c: t.client}
}

func (t *yuque) User() abc.YuQueUser {
	t.checkClient()
	return &yqUserImpl{c: t.client}
}

func (t *yuque) Repo() abc.YuQueRepo {
	t.checkClient()
	return &yqRepoImpl{c: t.client}
}

func (t *yuque) Doc() abc.YuQueDoc {
	t.checkClient()
	return &yqDocImpl{c: t.client}
}

func (t *yuque) _init() {
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

	t.Debug = xenv.IsDebug()
}

func (t *yuque) Auth(token string) abc.IYuque {
	t._init()

	t.client = resty.New().
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

	return t
}
