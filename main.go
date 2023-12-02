package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/dustin/go-humanize"
)

var wg sync.WaitGroup

func DirSize(path string) {

	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	if err != nil {
		fmt.Println("read dir err", "err", err)
	} else {
		fmt.Printf(" %-3s %-10s %s\n", "d", humanize.IBytes(uint64(size)), path)
	}

	wg.Done()

}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf(" please add arg in %s \n\n", os.Args[0])
		os.Exit(5)
	}
	t := time.Now()
	dir := os.Args[1]
	fmt.Printf("%-4s %-10s %s\n", "Type", "Size", "Path")
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("read dir failed\n", "err", err)
		os.Exit(5)
	}
	for _, fi := range files {
		if fi.IsDir() {

			wg.Add(1)
			tmpPath := filepath.Join(dir, fi.Name())
			go DirSize(tmpPath)
		} else {
			finfo, _ := fi.Info()
			fmt.Printf(" %-3s %-10s %s\n", "f", humanize.IBytes(uint64(finfo.Size())), fi.Name())
		}
	}
	wg.Wait()
	fmt.Println("\ntime: " + time.Since(t).String())
}
