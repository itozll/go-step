package tmpl

import (
	"embed"

	"github.com/itozll/go-step/internal/tmpl"
)

func Buildin() tmpl.Provider {
	return tmpl.Buildin()
}

func WithFS(f embed.FS, prefix string) tmpl.Provider {
	return tmpl.WithFS(f, prefix)
}

func WithPath(prefix string) tmpl.Provider {
	return tmpl.WithPath(prefix)
}
