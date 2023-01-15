package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"obsidian-server/repo"
	"obsidian-server/types"
)

func regexWrapper(s string) *regexp.Regexp {
	pattern, err := regexp.Compile(s)
	if err != nil {
		panic(err)
	}
	return pattern
}

func fileFinder(vault string, ch chan<- *types.File) {
	queue := make([]string, 0)
	queue = append(queue, vault)
	ignorePath := regexWrapper(`(?:^.*imdone-tasks.*$)|(?:^\.)`)
	ignoreLink := regexWrapper(`(?:^/400\sArchives.*$)|(?:^.*@Novall.*$)`)

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
				ch <- &types.File{RelPath: relPath, Name: file.Name()[:len(file.Name())-len(extension)], Extension: filepath.Ext(file.Name())}
			}
		}
	}
	close(ch)
}

func main() {
	// var rs []string
	var err error
	var ch = make(chan *types.File, 100)
	done := make(chan bool)

	defer repo.Deinit()
	err = repo.File.DropTable()
	if err != nil {
		panic(err)
	}

	err = repo.File.CreateTable()
	if err != nil {
		panic(err)
	}

	vault := os.Getenv("OBSIDIAN_VAULT")
	if vault == "" {
		panic("OBSIDIAN_VAULT is not set")
	}

	go fileFinder(vault, ch)
	go repo.File.CreateMany(ch, done)
	for {
		closed := <-done
		if closed {
			break
		} else {
			fmt.Println("Sleeping")
			time.Sleep(1 * time.Millisecond)
		}
	}

	// rs, err = repo.File.ReadAllPath()
	// if err != nil {
	// 	panic(err)
	// }
	// for _, r := range rs {
	// 	fmt.Printf("%v\n", r)
	// }

	var rs []int64
	rs, err = repo.File.FindByName(&types.File{RelPath: "/200 Areas", Name: "MOD", Extension: ".md"})
	if err != nil {
		panic(err)
	}
	for _, r := range rs {
		fmt.Printf("id=%d\n", r)
	}
}
