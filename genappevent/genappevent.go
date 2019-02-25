package genappevent

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	commentMarker = "//event "
)

func Run(targetDir string, renderers ...func(TmplParam) error) error {
	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		return err
	}
	var pkgName string
	eventComments := make([]string, 0, 50)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".go") {
			continue
		}
		if strings.HasSuffix(file.Name(), "_test.go") {
			continue
		}
		filename := filepath.Join(targetDir, file.Name())
		comments, err := extractEventComments(filename)
		if err != nil {
			return err
		}
		eventComments = append(eventComments, comments...)

		if pkgName == "" {
			pkgName, err = extractPkgName(filename)
			if err != nil {
				return err
			}
		}
	}

	events := make([]EventInfo, len(eventComments))
	for i, c := range eventComments {
		c := strings.TrimPrefix(c, commentMarker)
		cs := strings.Split(c, ",")
		info := EventInfo{
			Name:   strings.TrimSpace(cs[0]),
			Params: make([]ParamInfo, len(cs[1:])),
		}
		for i, param := range cs[1:] {
			ss := strings.Fields(param)
			info.Params[i] = ParamInfo{
				Name: ss[0],
				Type: ss[1],
			}
		}
		events[i] = info
	}

	param := TmplParam{
		PackageName: pkgName,
		Events:      events,
	}

	for _, renderer := range renderers {
		if err := renderer(param); err != nil {
			return err
		}
	}
	return nil
}

type TmplParam struct {
	PackageName    string
	Events         []EventInfo
	ImportPackages []string
}

type EventInfo struct {
	Name   string
	Params []ParamInfo
}

type ParamInfo struct {
	Name string
	Type string
}

func extractEventComments(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	results := make([]string, 0, 10)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(text, commentMarker) {
			results = append(results, text)
		}
	}
	return results, nil
}

func extractPkgName(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(text, "package ") {
			return strings.TrimPrefix(text, "package "), nil
		}
	}
	return "", nil
}
