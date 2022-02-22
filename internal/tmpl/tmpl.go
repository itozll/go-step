package tmpl

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

//go:embed template/*
var f embed.FS

type Provider interface {
	Get(string) []byte
}

type provider struct {
	get    func(string) ([]byte, error)
	prefix string
}

func Buildin() Provider {
	return &provider{
		get:    f.ReadFile,
		prefix: "template/",
	}
}

func WithFS(f embed.FS, prefix string) Provider {
	return &provider{
		get:    f.ReadFile,
		prefix: fill(prefix),
	}
}

func WithPath(prefix string) Provider {
	return &provider{
		get:    os.ReadFile,
		prefix: fill(prefix),
	}
}

func (p *provider) Get(name string) []byte {
	data, _ := get0(p.get, p.prefix+name)
	return data
}

func get0(fnc func(name string) ([]byte, error), name string) ([]byte, error) {
	data, err := fnc(name)
	if err != nil {
		if e, ok := err.(*fs.PathError); ok && e.Err == os.ErrNotExist {
			err = os.ErrNotExist
		}

		if err == os.ErrNotExist {
			// try <name>.tpl
			if !strings.HasSuffix(name, ".tpl") {
				return get0(fnc, name+".tpl")
			}
		}

		fmt.Printf("\n  ** can not get template[%s]: %s.\n", name, err)
		os.Exit(1)
	}

	return data, nil
}

func fill(prefix string) string {
	if !strings.HasSuffix(prefix, "/") {
		return prefix + "/"
	}

	return prefix
}
