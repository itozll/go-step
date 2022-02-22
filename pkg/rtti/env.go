package rtti

import (
	"os"
	"path/filepath"
)

const DefaultGoVersion = "1.17"

var (
	appName = filepath.Base(os.Args[0])

	Version      string
	GoVersion    = DefaultGoVersion
	TemplatePath string

	DefaultGroupName = ""

	DryRun  bool
	Verbose bool
	Commit  bool
)

func AppName() string { return appName }
