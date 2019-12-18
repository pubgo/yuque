package yuque

import (
	"fmt"
	"github.com/pubgo/g/xcmd"
	"github.com/pubgo/yuque/version"
	"runtime"
)

func Version() {
	xcmds.AddCommand(&xcmd.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "version info",
		Run: func(cmd *xcmd.Command, args []string) {
			fmt.Println("Version:", version.Version)
			fmt.Println("GitHash:", version.CommitV)
			fmt.Println("BuildDate:", version.BuildV)
			fmt.Println("GoVersion:", runtime.Version())
			fmt.Println("GOOS:", runtime.GOOS)
			fmt.Println("GOARCH:", runtime.GOARCH)
		},
	})
}
