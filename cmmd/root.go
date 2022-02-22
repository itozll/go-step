package cmmd

import (
	"fmt"
	"os"

	"github.com/itozll/go-step/pkg/rtti"
	"github.com/urfave/cli/v2"
)

var Version string
var Workspace = "go"

var App = &cli.App{
	Name:    rtti.AppName(),
	Usage:   "golang project skeleton",
	Version: Version,
}

func Main() {
	err := App.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, " \033[31;1m** [Err] %s\033[0m\n", err)
	}
}
