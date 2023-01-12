package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type File struct {
	relPath   string
	name      string
	extension string
}

// struct file ~ model db
func main() {
	vault := os.Getenv("OBSIDIAN_VAULT")
	// files, err := os.ReadDir(vault)
	queue := make([]string, 0)
	queue = append(queue, vault)
	ignorePath, err := regexp.Compile(`(?:^.*imdone-tasks.*$)|(?:^\.)`)
	if err != nil {
		panic(err)
	}
	ignoreLink, err := regexp.Compile(`(?:^/400\sArchives.*$)|(?:^.*@Novall.*$)`)
	if err != nil {
		panic(err)
	}

	result := make([]File, 0)
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		files, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			relPath := path[len(vault):]
			filePath := filepath.Join(path, file.Name())
			if ignorePath.MatchString(file.Name()) || ignoreLink.MatchString(relPath) {
				continue
			} else if file.IsDir() {
				queue = append(queue, filePath)
			} else if filepath.Ext(file.Name()) == ".md" {
				extension := filepath.Ext(file.Name())
				f := File{relPath, file.Name()[:len(file.Name())-len(extension)], filepath.Ext(file.Name())}
				result = append(result, f)
				// fmt.Printf("path=%s\nfile=%s\nextension=%s\n\n", f.relPath, f.name, f.extension)
			}
		}
	}
	for _, value := range result {
		fmt.Printf("%#v\n", value)
	}
}
