package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/go-fsnotify/fsnotify"
	"github.com/xyproto/recwatch"
)

func Run(watcher *recwatch.RecursiveWatcher, debug bool) {
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				// create a file/directory
				if event.Op&fsnotify.Create == fsnotify.Create {
					fi, err := os.Stat(event.Name)
					if err != nil {
						// eg. stat .subl513.tmp : no such file or directory
						if debug {
							DebugError(err)
						}
					} else if fi.IsDir() {
						if debug {
							DebugMessage("Detected new directory %s", event.Name)
						}
						if !recwatch.ShouldIgnoreFile(filepath.Base(event.Name)) {
							watcher.AddFolder(event.Name)
						}
					} else {
						if debug {
							DebugMessage("Detected new file %s", event.Name)
						}
						watcher.Files <- event.Name // created a file
					}
				}

				if event.Op&fsnotify.Write == fsnotify.Write {
					// modified a file, assuming that you don't modify folders
					if debug {
						DebugMessage("Detected file modification %s", event.Name)
					}
					watcher.Files <- event.Name
				}

			case err := <-watcher.Errors:
				log.Println("error", err)
			}
		}
	}()
}
