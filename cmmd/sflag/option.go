package sflag

import (
	"github.com/itozll/go-step/pkg/rtti"
	"github.com/urfave/cli/v2"
)

var (
	FlagCommit = &cli.BoolFlag{
		Name:        "commit",
		Usage:       "initial git repository commit information.",
		Value:       rtti.Commit,
		DefaultText: boolToString(rtti.Commit),
		Destination: &rtti.Commit,
	}

	FlagGoVersion = &cli.StringFlag{
		Name:        "go-version",
		Usage:       "the golang version used by the project.",
		Value:       rtti.GoVersion,
		DefaultText: rtti.GoVersion,
		Destination: &rtti.GoVersion,
	}

	FlagVerbose = &cli.BoolFlag{
		Name:        "verbose",
		Usage:       "add more details to output logging.",
		Value:       rtti.Verbose,
		DefaultText: boolToString(rtti.Verbose),
		Destination: &rtti.Verbose,
	}

	FlagTemplatePath = &cli.StringFlag{
		Name:        "template-path",
		Usage:       "path of custom templates.",
		Value:       rtti.TemplatePath,
		DefaultText: rtti.TemplatePath,
		Destination: &rtti.TemplatePath,
	}

	FlagDryrun = &cli.BoolFlag{
		Name:        "dry-run",
		Usage:       "Run through and reports activity without writing out results.",
		Value:       rtti.DryRun,
		DefaultText: boolToString(rtti.DryRun),
		Destination: &rtti.DryRun,
	}
)

func boolToString(o bool) string {
	if o {
		return "true"
	}

	return "false"
}
