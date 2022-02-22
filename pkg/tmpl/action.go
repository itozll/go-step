package tmpl

import "github.com/itozll/go-step/internal/tmpl"

type Action struct {
	Binder interface{}

	TmplProvider tmpl.Provider

	Path string

	Templates []string

	Copy []string
}
