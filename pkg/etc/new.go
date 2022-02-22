package etc

import (
	"github.com/itozll/go-step/pkg/rtti"
	"github.com/itozll/go-step/pkg/tmpl"
)

var (
	Root = &tmpl.Action{
		Path: "",
		Templates: []string{
			"README.md",
			"Makefile",
			"main:go",
			"gomod::go.mod",
		},
		Copy: []string{
			".gitignore",
			"generator:.sh",
		},
	}

	Cmd = &tmpl.Action{
		Path: "app/cmd",
		Templates: []string{
			"cmdroot::root.go",
			"cmdserver::server.go",
		},
	}

	Flag = &tmpl.Action{
		Path: "app/cmd/flag",
		Templates: []string{
			"option:go",
		},
	}

	Rttype = &tmpl.Action{
		Path: "app/internal/rttype",
		Copy: []string{
			"envgo::env.go",
		},
	}

	NewCommand = func() *tmpl.Command {
		return &tmpl.Command{
			Before: func() error {
				tmpl.TargetPath = rtti.Binder.Workspace
				return nil
			},
			TmplProvider: tmpl.Buildin(),
			Binder:       rtti.Binder,
			Actions:      append([]*tmpl.Action{Root, Cmd, Flag, Rttype}, HiddenAction...),
		}
	}
)
