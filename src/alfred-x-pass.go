package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"

	aw "github.com/deanishe/awgo"
)

type file struct {
	Title    string
	Value    string
	Subtitle string
	Match    string
}

var (
	query             string
	wf                *aw.Workflow
	passwordStorePath = os.ExpandEnv("${HOME}/.password-store")
)

func init() {
	wf = aw.New()
}

func removePrefix(s string, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

func fileNameWithoutExtTrimSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func walkDir(dir string) (files []file) {
	filepath.WalkDir(dir, func(path string, entry os.DirEntry, err error) error {
		if entry.IsDir() {
			return nil
		}

		if filepath.Ext(entry.Name()) != ".gpg" {
			return nil
		}

		relativePath := fileNameWithoutExtTrimSuffix(removePrefix(path, passwordStorePath+"/"))
		title := fileNameWithoutExtTrimSuffix(entry.Name())
		match := strings.Join(strings.Split(relativePath, "/"), " ") + " " + title + " " + relativePath

		files = append(files, file{
			Title:    title,
			Value:    relativePath,
			Subtitle: relativePath,
			Match:    match,
		})
		return nil
	})

	return files
}

func run() {
	wf.Args() // call to handle magic actions
	flag.Parse()
	query = flag.Arg(0)

	for _, file := range walkDir(passwordStorePath) {

		if strings.HasSuffix(file.Value, "/") || file.Title == "" {
			continue
		}

		// ignore files that start with "."
		if strings.HasPrefix(file.Value, ".") {
			continue
		}

		// ignore files that start with "__" in some parent directory
		if strings.Contains(file.Value, "/__") {
			continue
		}

		// ignore .git directories
		if strings.HasPrefix(file.Value, ".git") {
			continue
		}

		it := wf.NewItem(file.Title)
		it.Arg(file.Value)
		it.Subtitle(file.Value)
		it.Match(file.Match)
		it.Valid(true)
	}

	wf.Filter(query)

	if wf.IsEmpty() {
		wf.NewWarningItem("No matches found", "Try a different query?")
	}

	// Send results to Alfred
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
