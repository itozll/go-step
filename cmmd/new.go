package cmmd

import (
	"errors"
	"fmt"

	"github.com/itozll/go-step/cmmd/sflag"
	"github.com/itozll/go-step/pkg/etc"
	"github.com/itozll/go-step/pkg/rtti"
	"github.com/urfave/cli/v2"
)

var NewCmd = &cli.Command{
	Name:      "new",
	Aliases:   []string{"n"},
	Usage:     fmt.Sprintf("create an %s workspace", Workspace),
	UsageText: fmt.Sprintf("%s new [options] <project-name>", rtti.AppName()),
	Flags: []cli.Flag{
		// sflag.FlagCommit,
		sflag.FlagGoVersion,
		// sflag.FlagVerbose,
	},
	// After: func(c *cli.Context) error {
	// 	fmt.Printf("\n \033[33;1mEXECUTE THE FOLLOWING COMMAND TO START YOUR PROJECT: \n\033[31;1m  cd " +
	// 		rtti.Binder.KebabWorkspace + " && go mod tidy\033[0m\n\n")
	// 	return nil
	// },
	Action: func(c *cli.Context) error {
		if c.NArg() == 0 {
			return errors.New("need the <project-name> to continue")
		}

		if c.NArg() != 1 {
			return errors.New("<project-name> must be at the end")
		}

		if err := rtti.Init(c.Args().Get(0)); err != nil {
			return err
		}

		if err := etc.NewCommand().Exec(); err != nil {
			return err
		}

		fmt.Printf("\n \033[33mEXECUTE THE FOLLOWING COMMAND TO START YOUR PROJECT: \n\033[33;1m  cd " +
			rtti.Binder.KebabWorkspace + " && go mod tidy\033[0m\n\n")

		return nil
	},
}

func init() {
	App.Commands = append(App.Commands, NewCmd)
}
