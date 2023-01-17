package main

import (
	"collection/node"
	"collection/queue"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type File struct {
	RelPath   string
	Name      string
	Extension string
}

func regexWrapper(s string) *regexp.Regexp {
	pattern, err := regexp.Compile(s)
	if err != nil {
		panic(err)
	}
	return pattern
}

func fileFinder(vault string) []string {
	q := queue.New[string](vault)
	// queue := make([]string, 0)
	// queue = append(queue, vault)
	// q.Push(vault)
	ignorePath := regexWrapper(`(?:^.*imdone-tasks.*$)|(?:^\.)`)
	ignoreLink := regexWrapper(`(?:^/400\sArchives.*$)|(?:^.*@Novall.*$)`)
	result := make([]string, 0)

	for q.Size() > 0 {
		// for len(queue) > 0 {
		path := q.Pop()
		// path := queue[0]
		// queue = queue[1:]
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
				q.Push(filePath)
				// queue = append(queue, filePath)
			} else if filepath.Ext(file.Name()) == ".md" {
				extension := filepath.Ext(file.Name())
				result = append(result, relPath+"/"+file.Name()[:len(file.Name())-len(extension)]+filepath.Ext(file.Name()))
			}
		}
	}
	return result
}

func main() {
	vault := os.Getenv("OBSIDIAN_VAULT")
	files := fileFinder(vault)

	tree := node.NewTree()
	for _, file := range files {
		tree.Update(file)
	}

	tree.DFS()
	path := "/001 Zettelkasten/audio/Blarg Audio Libraries.md"
	fmt.Printf("%s %t\n", path, tree.Search(path))
}
