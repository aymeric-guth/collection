package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// struct file ~ model db
func main() {
	vault := os.Getenv("OBSIDIAN_VAULT")
	// files, err := os.ReadDir(vault)
	queue := make([]string, 0)
	queue = append(queue, vault)

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		files, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			if strings.Compare(file.Name()[0:1], ".") == 0 {
				continue
			} else if file.IsDir() {
				queue = append(queue, path+"/"+file.Name())
			} else if filepath.Ext(file.Name()) == ".md" {
				// path + "/" + file.Name()
				fmt.Printf("MD-FILE: %s\n", file.Name())
			} else {
				fmt.Printf("FILE: %s\n", file.Name())
			}
		}
	}
}
