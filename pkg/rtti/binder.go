package rtti

import (
	"errors"
	"strings"

	"github.com/iancoleman/strcase"
)

var Binder *AppBind

type AppBind struct {
	Workspace string // maybe 'AppName' 'app-name' 'app_name'
	AppName   string

	Github         string // github.com/<GroupName>/<KebadWorkspace>
	GroupName      string
	SnakeWorkspace string // app_name
	KebabWorkspace string // app-name
	CamelWorkspace string // AppName

	GoVersion string
}

func Init(workspace string) error {
	Binder = &AppBind{
		AppName:   appName,
		GoVersion: GoVersion,
	}

	if len(Binder.GoVersion) == 0 {
		Binder.GoVersion = DefaultGoVersion
	}

	if len(workspace) == 0 {
		return errors.New("workspace must not be empty")
	}

	if err := split(workspace); err != nil {
		return err
	}

	if len(Binder.Workspace) == 0 {
		return errors.New("workspace must not be empty")
	}

	if len(Binder.GroupName) == 0 {
		return errors.New("group name must not be empty")
	}

	slice(Binder.Workspace)

	Binder.Github = "github.com/" + Binder.GroupName + "/" + Binder.KebabWorkspace
	return nil
}

func slice(workspace string) {
	Binder.SnakeWorkspace = strcase.ToSnake(workspace)
	Binder.KebabWorkspace = strcase.ToKebab(workspace)
	Binder.CamelWorkspace = strcase.ToCamel(workspace)
}

func split(workspace string) error {
	list := strings.Split(workspace, "/")
	if len(list) > 2 {
		return errors.New("error workspace")
	}
	if len(list) == 1 {
		Binder.GroupName = DefaultGroupName
		Binder.Workspace = list[0]
		return nil
	}

	Binder.GroupName = list[0]
	Binder.Workspace = list[1]
	return nil
}
