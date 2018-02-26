package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		var info string
		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			if len(pair) > 1 {
				info = fmt.Sprintf("%s\n%s=%s", info, pair[0], pair[1])
			}
			info = info + "\nngocnd:build-300"
		}
		return info
	})
	m.Get("/merge-branch", func() string {
		var info string
		info = "branch is merged"
		return info
	})
	m.Get("/new-build", func() string {
		var info string
		info = "branch is merged"
		info = info + "\n ngocnd"
		return info
	})
	m.Run()
}
