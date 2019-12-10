package yuque

import (
	"encoding/json"
	"fmt"
	"github.com/pubgo/g/pkg"
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"reflect"
)

func Debug(d ...interface{}) {
	if !xenv.IsDebug() {
		return
	}

	for _, i := range d {
		if i == nil || pkg.IsNone(i) {
			continue
		}

		switch reflect.TypeOf(i).Kind() {
		case reflect.String, reflect.Slice:
			fmt.Println(i)
		case reflect.Struct, reflect.Ptr, reflect.Map:
			fmt.Printf("%s", xerror.PanicErr(json.MarshalIndent(i, "", "\t")))
		default:
			fmt.Printf("%#v\n", i)
		}
	}
}

var _url = func(url string) func(...interface{}) string {
	return func(params ...interface{}) string {
		return fmt.Sprintf(url, params...)
	}
}
