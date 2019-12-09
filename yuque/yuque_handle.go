package yuque

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"net/http"
	"strings"
)

const BaseUrlPrefix = "/api/v2"

func yqVerb(c *resty.Request, url string, pathParams map[string]string, queryParams map[string]string, dt interface{}) (err error) {
	defer xerror.RespErr(&err)
	xerror.PanicT(dt == nil, "response error")

	url = BaseUrlPrefix + url
	return
}

func yqPut(c *resty.Request, url string, body interface{}, dt interface{}) (err error) {
	defer xerror.RespErr(&err)

	xerror.PanicT(dt == nil, "response error")
	url = BaseUrlPrefix + url

	c.SetHeader("Content-Type", "application/json")
	return
}

func yqDelete(c *resty.Request, url string, queryParams map[string]string, dt interface{}) (err error) {
	defer xerror.RespErr(&err)

	xerror.PanicT(dt == nil, "response error")

	url = BaseUrlPrefix + url

	_c := c
	if queryParams != nil {
		_c = _c.SetQueryParams(queryParams)
	}

	resp, err := _c.Delete(url)
	xerror.PanicM(err, "%s request error", url)
	xerror.PanicTT(resp.StatusCode() != http.StatusOK, func(err *xerror.Err) {
		err.Msg("%s resp error", url)
		err.M("status_code", resp.StatusCode())
		if queryParams != nil {
			err.M("query_params", queryParams)
		}
		err.M("resp_body", resp.String())
	})
	xerror.PanicM(json.Unmarshal(resp.Body(), dt), "%s resp decode error", url)
	return
}

func yqGet(c *resty.Request, url string, queryParams map[string]string, dt interface{}) (err error) {
	defer xerror.RespErr(&err)

	xerror.PanicT(dt == nil, "response error")

	url = BaseUrlPrefix + url

	_c := c
	if queryParams != nil {
		_c = _c.SetQueryParams(queryParams)
	}

	resp, err := _c.Get(url)
	xerror.PanicM(err, "%s request error", url)
	xerror.PanicTT(resp.StatusCode() != http.StatusOK, func(err *xerror.Err) {
		err.Msg("%s resp error", url)
		err.M("status_code", resp.StatusCode())
		if queryParams != nil {
			err.M("query_params", queryParams)
		}
		err.M("resp_body", resp.String())
	})
	if dt != nil {
		xerror.PanicM(json.Unmarshal(resp.Body(), dt), "%s resp decode error", url)
	}
	return
}

func yqPost(c *resty.Request, url string, body interface{}, dt interface{}) (err error) {
	defer xerror.RespErr(&err)

	xerror.PanicT(dt == nil, "response error")

	url = BaseUrlPrefix + url

	c.SetHeader("Content-Type", "application/json")

	_c := c

	if body != nil {
		_c = _c.SetBody(body)
	}

	resp, err := _c.Post(url)
	xerror.PanicM(err, "%s request error", url)
	xerror.PanicTT(resp.StatusCode() != http.StatusOK, func(err *xerror.Err) {
		err.Msg("%s resp error", url)
		err.M("status_code", resp.StatusCode())
		if body != nil {
			err.M("body", body)
		}
		err.M("resp", resp.String())
	})
	if dt != nil {
		xerror.PanicM(json.Unmarshal(resp.Body(), dt), "%s resp decode error", url)
	}

	return
}

func RateLimit(resp *resty.Response) {
	// 总次数
	resp.Header().Get("X-RateLimit-Limit")
	// 剩余次数
	resp.Header().Get("X-RateLimit-Remaining")
}

func IsRateLimitOver(resp *resty.Response) bool {
	return resp.StatusCode() == 429 && strings.Contains(strings.ToLower(resp.String()), `too many requests`)
}
