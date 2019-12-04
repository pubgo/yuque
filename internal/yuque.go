package internal

import (
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/errors"
	"github.com/pubgo/g/pkg"
	"github.com/pubgo/yuque/internal/routes"
	"sync"
	"time"
)

type yqClient struct {
	routes.YuqueUser
	routes.YuqueGroup

	c *resty.Request
}

type YuQue struct {
	BaseUrl          string
	AppName          string
	RetryCount       int
	RetryWaitTime    int
	RetryMaxWaitTime int
	Timeout          int
	Debug            bool
	yqClientMap      map[string]*resty.Client
}

func (t *YuQue) R(name string) (c *yqClient, err error) {
	defer errors.RespErr(&err)

	_c, ok := t.yqClientMap[name]
	errors.PanicT(!ok, "%s not found", name)
	errors.PanicT(pkg.IsNone(_c), "%s is null", name)

	return &yqClient{c: _c.R()}, nil
}

func (t *YuQue) AddAuth(name, token string) bool {
	if _, ok := t.yqClientMap[name]; ok {
		return ok
	}

	t.yqClientMap[name] = resty.New().
		SetDebug(t.Debug).
		SetContentLength(true).
		SetHostURL("https://www.yuque.com").
		SetRetryCount(t.RetryCount).
		SetRetryWaitTime(time.Second*time.Duration(t.RetryWaitTime)).
		SetRetryMaxWaitTime(time.Second*time.Duration(t.RetryMaxWaitTime)).
		SetHeader("Content-Type", "application/json").
		SetTimeout(time.Second * time.Duration(t.Timeout)).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"User-Agent":   t.AppName,
			"X-Auth-Token": token,
		})
	return true
}

var _yq *YuQue
var _yqOnce sync.Once

func Default() *YuQue {
	_yqOnce.Do(func() {
		_yq = &YuQue{
			BaseUrl:          "https://www.yuque.com/api/v2",
			AppName:          "test",
			RetryCount:       3,
			RetryWaitTime:    5,
			RetryMaxWaitTime: 20,
			Timeout:          60,
			Debug:            true,
			yqClientMap:      make(map[string]*resty.Client),
		}
	})
	return _yq
}
