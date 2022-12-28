package env

import (
	"os"
	"strings"
)

var (
	vars = make(t_vars)
)

type t_vars map[string]string

func (vs t_vars) Get(key, def string) (v string) {
	var ok bool
	if v, ok = vs[key]; ok {
		return
	}

	return def
}

func Initialize() {
	var kv []string
	for _, v := range os.Environ() {
		kv = strings.Split(v, "=")
		vars[kv[0]] = kv[1]
	}
}

func Get(key, def string) string {
	return vars.Get(key, def)
}
