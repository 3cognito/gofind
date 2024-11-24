package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func worker(path, text string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	var counter int
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", path, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, word := range strings.Fields(line) {
			if strings.Contains(strings.ToLower(word), strings.ToLower(text)) {
				// fmt.Printf("Found '%s' in file %s\n", text, path)
				counter++
			}
		}
	}

	ch <- counter
}

func CountOccurrences(searchPath, text, exclude string) int {
	exclusionList := make([]string, 0)
	exclusionList = append(exclusionList, strings.Split(exclude, ",")...)
	ch := make(chan int)
	wg := sync.WaitGroup{}

	go func() {
		filepath.WalkDir(searchPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				fmt.Printf("Error accessing %s: %v\n", path, err)
				return nil
			}

			if !d.IsDir() {
				if contains(exclusionList, filepath.Ext(path)) {
					return nil
				}
				wg.Add(1)
				go worker(path, text, ch, &wg)
			}
			return nil
		})

		wg.Wait()
		close(ch)
	}()

	var total int
	for count := range ch {
		total += count
	}

	fmt.Printf("Total occurrences of '%s': %d\n", text, total)
	return total
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <search_path> <text> <exclusion_list (comma separated)>")
		return
	}

	searchPath := os.Args[1]
	exclude := os.Args[2]
	text := os.Args[3]

	CountOccurrences(searchPath, text, exclude)
}
