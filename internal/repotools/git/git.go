package git

import (
	"golang.org/x/mod/semver"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func Tags(path string) ([]string, error) {
	output, err := Git(path, "tag -l")
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(output), "\n")
	return split, nil
}

func Fetch(path string) error {
	_, err := Git(path, "fetch --all")
	return err
}

func Git(path string, arguments ...string) (output []byte, err error) {
	cmd := exec.Command("git", arguments...)
	if len(path) == 0 {
		path, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	}
	cmd.Dir = path
	cmd.Env = append(os.Environ(), "PWD="+path)

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	return cmd.Output()
}

type ModuleTags map[string][]string

func (r ModuleTags) Latest(module string) (string, bool) {
	_, ok := r[module]
	if !ok {
		return "", false
	}
	return r[module][0], true
}

func ParseTags(tags []string) ModuleTags {
	modules := make(map[string][]string)

	for _, tag := range tags {
		idx := strings.LastIndex(tag, "/")

		module := "/"
		version := tag

		if idx != -1 {
			module = tag[:idx]
			version = tag[idx+1:]
		}

		if !semver.IsValid(version) {
			// Ignore invalid Go semver tags
			continue
		}

		modules[module] = append(modules[module], version)
	}

	for _, versions := range modules {
		sort.Slice(versions, func(i, j int) bool {
			// We want to sort higher versions first
			return semver.Compare(versions[i], versions[j]) > 0
		})
	}

	return modules
}
