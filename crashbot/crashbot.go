package main

import (
	"fmt"
	"log"

	"github.com/howeyc/fsnotify"
)

func handleEvents(ev *fsnotify.FileEvent) {
	if ev.IsCreate() {
		fmt.Println(ev)
	}
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				handleEvents(ev)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch("/fuzz/jq/out/crashes")
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Watch("/fuzz/jq/out/hangs")
	if err != nil {
		log.Fatal(err)
	}

	// Hang so program doesn't exit
	<-done

	/* ... do stuff ... */
	watcher.Close()
}
