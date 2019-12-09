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
			dt, err := json.MarshalIndent(i, "", "\t")
			xerror.PanicM(err, "P json MarshalIndent error")
			fmt.Println(string(dt))
		default:
			fmt.Printf("%#v\n", i)
		}
	}
}
