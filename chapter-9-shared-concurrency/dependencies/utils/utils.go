package utils

import (
	"bytes"
	"encoding/json"
	"github.com/lapeko/books__go_programming_language/chapter-9-shared-concurrency/dependencies/utils/slice"
	"os/exec"
	"strings"
)

const modPath = "github.com/lapeko/books__go_programming_language/chapter-9-shared-concurrency/dependencies"

type listItem struct {
	Dir        string
	ImportPath string
	Name       string
	Doc        string
	Target     string
	Goroot     bool
	Standard   bool
	Root       string
	GoFiles    []string
	Imports    []string
	Deps       []string
}

func GetDeps(envDeps, deps []string) ([]string, error) {
	searchDeps := slice.GetIntersection(envDeps, deps)

	if len(searchDeps) == 0 {
		return searchDeps, nil
	}

	items, err := goList(searchDeps...)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		innerDeps, err := GetDeps(envDeps, item.Deps)
		if err != nil {
			return nil, err
		}
		searchDeps = append(searchDeps, innerDeps...)
	}

	return slice.RmDuplicates(searchDeps), nil
}

func GetAllEnvDeps() (deps []string, err error) {
	items, err := goList("./...")

	if err != nil {
		return
	}

	for _, item := range items {
		if strings.HasPrefix(item.ImportPath, modPath) {
			deps = append(deps, item.ImportPath)
		}
	}

	return
}

func goList(deps ...string) (list []listItem, err error) {
	cmd := exec.Command("go", append([]string{"list", "-json"}, deps...)...)

	res, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(bytes.NewReader(res))
	for decoder.More() {
		var l listItem
		if err := decoder.Decode(&l); err != nil {
			return nil, err
		}
		list = append(list, l)
	}
	return
}
