package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/dustin/go-humanize"
)

var waitGroup sync.WaitGroup
var ch = make(chan struct{}, 255)

func dirents(path string) ([]os.FileInfo, bool) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	return entries, true
}

// 递归计算目录下所有文件
func walkDir(path string, fileSize chan<- int64) {
	defer waitGroup.Done()
	ch <- struct{}{} //限制并发量
	entries, ok := dirents(path)
	<-ch
	if !ok {
		log.Fatal("can not find this dir path!!")
		return
	}
	for _, e := range entries {
		if e.IsDir() {
			waitGroup.Add(1)
			go walkDir(filepath.Join(path, e.Name()), fileSize)
		} else {
			fileSize <- e.Size()
		}
	}
}

func all_file(dir_path string) {

	//文件大小chennel
	fileSize := make(chan int64)
	//文件总大小
	var sizeCount int64
	//文件数目
	var fileCount int

	//计算目录下所有文件占的大小总和
	waitGroup.Add(1)
	go walkDir(dir_path, fileSize)

	go func() {
		defer close(fileSize)
		waitGroup.Wait()
	}()

	for size := range fileSize {
		fileCount++
		sizeCount += size
	}
	fsize := humanize.IBytes(uint64(sizeCount))
	fmt.Printf("%-10s %-8s  %s \n", fsize, fmt.Sprint(fileCount), dir_path)
	// fmt.Printf("size: %.1fM   file count: %d\n", float64(sizeCount)/1024/1024, fileCount)
}
func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf(" please add arg in %s \n\n", os.Args[0])
		os.Exit(5)
	}
	t := time.Now()
	dir_path := os.Args[1]
	fmt.Printf(" %-3s %-10s %-8s  %s\n", "Type", "Size", "Count", "Path")
	files, err := ioutil.ReadDir(dir_path)
	if err != nil {
		log.Fatal(err)
	} else {
		for _, fi := range files {
			if fi.IsDir() {
				fmt.Printf("  %-3s ", "d")
				all_file(filepath.Join(dir_path, fi.Name()))
			} else {
				fsize := humanize.IBytes(uint64(fi.Size()))
				fmt.Printf("  %-3s %-10s %-8s  %s \n", "f", fsize, "1", fi.Name())
			}
		}
	}
	fmt.Println("time: " + time.Since(t).String())
}
