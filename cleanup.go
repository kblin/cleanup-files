package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func SearchFiles(search_path *string, max_age *int, interval *int) {
	for {
		filepath.Walk(*search_path, func(path string, f os.FileInfo, err error) error {
			fmt.Println("file: ", path)
			age := time.Now().Sub(f.ModTime())
			fmt.Println("age ", age.Minutes())
			return nil
		})
		time.Sleep(time.Duration(*interval) * time.Second)
	}

}

func main() {
	interval := flag.Int("interval", 5, "How often (in seconds) to run the check")
	search_path := flag.String("path", "/some/directory", "A path to look for files to clean up in")
	age := flag.Int("age", 5, "Maximum age (in minutes) a file can have before being cleaned up")

	flag.Parse()

	go SearchFiles(search_path, age, interval)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-ch)
}
