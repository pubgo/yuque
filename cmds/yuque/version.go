package yuque

import (
	"fmt"
	"github.com/pubgo/g/xcmds"
	"github.com/pubgo/yuque/version"
	"runtime"
)

func Version() {
	xcmds.AddCommand(&xcmds.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "version info",
		Run: func(cmd *xcmds.Command, args []string) {
			fmt.Println("Version:", version.Version)
			fmt.Println("GitHash:", version.CommitV)
			fmt.Println("BuildDate:", version.BuildV)
			fmt.Println("GoVersion:", runtime.Version())
			fmt.Println("GOOS:", runtime.GOOS)
			fmt.Println("GOARCH:", runtime.GOARCH)
		},
	})
}
