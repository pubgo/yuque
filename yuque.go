package yuque

import (
	"github.com/pubgo/g/xinit"
	"github.com/pubgo/yuque/internal"
)

var Default = internal.Default

func init() {
	xinit.InitProvide(func() *internal.YuQue {


		return nil
	})
}
