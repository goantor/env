package env

import (
	"os"
	"path"
	"runtime"
	"strings"
)

var (
	vars    = make(t_vars)
	appPath string
)

type t_vars map[string]string

func (vs t_vars) Get(key, def string) (v string) {
	var ok bool
	if v, ok = vs[key]; ok {
		return
	}

	return def
}

func (vs t_vars) Set(key, val string) {
	vs[key] = val
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

func Set(key, val string) {
	vars.Set(key, val)
}

func AppPath() string {
	if 0 == len(appPath) {
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			panic("current file unknown error")
		}

		appPath = path.Dir(path.Dir(filename))
	}
	return appPath
}
