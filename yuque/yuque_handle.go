package yuque

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/pubgo/g/xerror"
	"net/http"
	"strings"
)

func yqVerb(c *resty.Request, r string, pathParams map[string]string, queryParams map[string]string, dt interface{}) (err error) {

}

func yqPut(c *resty.Request, r string, pathParams map[string]string, queryParams map[string]string, dt interface{}) (err error) {
	c.SetHeader("Content-Type", "application/json")
}

func yqDelete(c *resty.Request, r string, pathParams map[string]string, queryParams map[string]string, dt interface{}) (err error) {

}

func yqGet(c *resty.Request, r string, queryParams map[string]string, dt interface{}) (err error) {
	defer xerror.RespErr(&err)

	_c := c
	if queryParams != nil {
		_c = _c.SetQueryParams(queryParams)
	}

	resp, err := _c.Get(r)
	xerror.PanicM(err, "%s request error", r)
	xerror.PanicTT(resp.StatusCode() != http.StatusOK, func(err *xerror.Err) {
		err.Msg("%s resp error", r)
		err.M("status_code", resp.StatusCode())
		if queryParams != nil {
			err.M("query_params", queryParams)
		}
		err.M("resp_body", resp.String())
	})
	xerror.PanicM(json.Unmarshal(resp.Body(), &dt), "%s resp decode error", r)
	return
}

func yqPost(c *resty.Request, r string, pathParams map[string]string, body interface{}, dt interface{}) (err error) {
	defer xerror.RespErr(&err)

	c.SetHeader("Content-Type", "application/json")

	_c := c
	if pathParams != nil {
		_c = _c.SetPathParams(pathParams)
	}

	if body != nil {
		_c = _c.SetBody(body)
	}

	resp, err := _c.Post(r)
	xerror.PanicM(err, "%s request error", r)
	xerror.PanicTT(resp.StatusCode() != http.StatusOK, func(err *xerror.Err) {
		err.Msg("%s resp error", r)
		err.M("status_code", resp.StatusCode())
		if pathParams != nil {
			err.M("path_params", pathParams)
		}
		if body != nil {
			err.M("body", body)
		}
		err.M("resp", resp.String())
	})
	xerror.PanicM(json.Unmarshal(resp.Body(), &dt), "%s resp decode error", r)
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
