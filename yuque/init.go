package yuque

import "fmt"

const BaseUrlPrefix = "/api/v2"

var _url = func(url string) func(...interface{}) string {
	return func(params ...interface{}) string {
		return fmt.Sprintf(BaseUrlPrefix+url, params...)
	}
}
