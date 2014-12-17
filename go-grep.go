package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var root, query string
var found = 1
var wg sync.WaitGroup

func readFile(wg *sync.WaitGroup, path string) {
	defer wg.Done()

	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		if strings.Contains(scanner.Text(), query) {
			found = 0
			fmt.Printf("%s/%s:%d: %s\n", root, path, i, scanner.Text())
		}
	}
}

func main() {
	flag.Parse()
	query = flag.Arg(0)
	root = flag.Arg(1)

	filepath.Walk(root, func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			wg.Add(1)
			go readFile(&wg, path)
		}
		return nil
	})
	wg.Wait()
	defer os.Exit(found)
}
