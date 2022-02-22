package tmpl

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/itozll/go-step/internal/tmpl"
)

var (
	TargetPath = ""

	once sync.Once
)

type Command struct {
	Before func() error

	After func() error

	Binder interface{}

	TmplProvider tmpl.Provider

	Actions []*Action
}

func (cmd *Command) Exec() (err error) {
	hasBinder := cmd.Binder != nil

	if cmd.Before != nil {
		if err = cmd.Before(); err != nil {
			return
		}
	}

	fmt.Printf("\033[32m@ START\033[0m\n")
	for _, action := range cmd.Actions {
		path := absPath(action.Path)
		relPath := relPath(path)

		if len(action.Templates) > 0 {
			if !hasBinder {
				return errors.New("need a binder to parse template")
			}

			if err = cmd.parseAndCopy(relPath, action, action.Templates, true); err != nil {
				return
			}
		}

		if len(action.Copy) > 0 {
			if err = cmd.parseAndCopy(relPath, action, action.Copy, false); err != nil {
				return
			}
		}
	}

	if cmd.After != nil {
		return cmd.After()
	}

	return nil
}

func (cmd *Command) parseAndCopy(path string, action *Action, list []string, isTmpl bool) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil && os.IsNotExist(err) {
		return err
	}

	p := action.TmplProvider
	if p == nil {
		p = cmd.TmplProvider
	}

	for _, name := range list {
		if len(name) == 0 {
			return errors.New("source name must not be empty")
		}

		dstName, srcName := splitName(name)
		dstPath := path + "/" + dstName

		dstFd, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFd.Close()

		data := p.Get(srcName)
		if !isTmpl {
			io.Copy(dstFd, strings.NewReader(string(data)))
		} else {
			tpl, err := template.New(dstPath).Parse(string(data))
			if err != nil {
				return err
			}

			binder := action.Binder
			if binder == nil {
				binder = cmd.Binder
			}

			err = tpl.Execute(dstFd, binder)
			if err != nil {
				return err
			}
		}

		fmt.Printf("\033[32m  [GEN]  %s\033[0m\n", dstPath)
	}

	return nil
}

func absPath(path string) string {
	once.Do(func() {
		if len(TargetPath) > 0 && !strings.HasPrefix(TargetPath, "/") {
			TargetPath += "/"
		}
	})

	p, _ := filepath.Abs(TargetPath + path)
	return p
}

func relPath(path string) string {
	currPath, _ := os.Getwd()
	p, _ := filepath.Rel(currPath, path)
	return p
}

func splitName(name string) (dst, src string) {
	names := strings.Split(name, ":")
	src = names[0]
	dst = src

	l := len(names)

	// name
	if l == 1 {
		return
	}

	if l > 2 {
		// name::new-name
		if len(names[2]) > 0 {
			dst = names[2]
			return
		}
	}

	// name:<suffix> -> model:go change to model.go
	if len(names[1]) > 0 {
		dst += "." + names[1]
	}

	return
}
