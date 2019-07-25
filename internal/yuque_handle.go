package internal

import (
	"encoding/json"
	"github.com/pubgo/errors"
	"net/http"
)

func (t *yqClient) get(r string, pathParams map[string]string, queryParams map[string]string, dt interface{}) {
	_c := t.c
	if pathParams != nil {
		_c = _c.SetPathParams(pathParams)
	}
	if queryParams != nil {
		_c = _c.SetQueryParams(queryParams)
	}
	resp, err := _c.Get(r)
	errors.Wrap(err, "%s request error", r)
	errors.TT(resp.StatusCode() != http.StatusOK, func(err *errors.Err) {
		err.Msg("%s resp error", r)
		err.M("status_code", resp.StatusCode())
		if pathParams != nil {
			err.M("path_params", pathParams)
		}
		if queryParams != nil {
			err.M("query_params", queryParams)
		}
		err.M("resp", resp.String())
	})
	errors.Wrap(json.Unmarshal(resp.Body(), &dt), "%s resp decode error", r)
}

func (t *yqClient) post(r string, pathParams map[string]string, body interface{}, dt interface{}) {
	_c := t.c
	if pathParams != nil {
		_c = _c.SetPathParams(pathParams)
	}

	if body != nil {
		_c = _c.SetBody(body)
	}

	resp, err := _c.Post(r)
	errors.Wrap(err, "%s request error", r)
	errors.TT(resp.StatusCode() != http.StatusOK, func(err *errors.Err) {
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
	errors.Wrap(json.Unmarshal(resp.Body(), &dt), "%s resp decode error", r)
}