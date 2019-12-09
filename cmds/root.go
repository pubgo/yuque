package cmds

import (
	"github.com/pubgo/g/version"
	"github.com/pubgo/g/xcmds"
	"github.com/pubgo/g/xcmds/xcmd_ss"
	"github.com/pubgo/g/xerror"
)

const Service = "yuque"
const EnvPrefix = "YQ"

// Execute exec
var Execute = xcmds.Init(EnvPrefix, func(cmd *xcmds.Command) {
	defer xerror.Assert()

	cmd.Use = Service
	cmd.Version = version.Version

	xcmd_ss.Init()
})
