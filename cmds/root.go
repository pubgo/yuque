package cmds

import (
	"github.com/pubgo/g/version"
	"github.com/pubgo/g/xcmds"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/yuque/cmds/yuque"
)

const Service = "yuque"
const EnvPrefix = "YQ"

// Execute exec
var Execute = xcmds.Init(EnvPrefix, func(cmd *xcmds.Command) {
	defer xerror.Assert()

	cmd.Use = Service
	cmd.Version = version.Version

	yuque.Version()
})
