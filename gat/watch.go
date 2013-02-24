package gat

import (
    "errors"
    "github.com/howeyc/fsnotify"
    "log"
    "os"
)

type RecursiveWatcher struct {
    *fsnotify.Watcher
    Files   chan string
    Folders chan string
}

func NewRecurisveWatcher(path string) (*RecursiveWatcher, error) {
    folders := Subfolders(path)
    if len(folders) == 0 {
        return nil, errors.New("No folders to watch.")
    }

    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return nil, err
    }
    rw := &RecursiveWatcher{Watcher: watcher}

    rw.Files = make(chan string, 10)
    rw.Folders = make(chan string, len(folders))

    for _, folder := range folders {
        rw.AddFolder(folder)
    }
    return rw, nil
}

func (watcher *RecursiveWatcher) AddFolder(folder string) {
    err := watcher.WatchFlags(folder, fsnotify.FSN_CREATE|fsnotify.FSN_MODIFY)
    if err != nil {
        log.Println("Error watching: ", folder, err)
    }
    watcher.Folders <- folder
}

func Watch(path string) *RecursiveWatcher {
    watcher, err := NewRecurisveWatcher(path)
    if err != nil {
        log.Fatal(err)
    }

    go func() {
        for {
            select {
            case event := <-watcher.Event:
                // create a directory
                if event.IsCreate() {
                    fi, err := os.Stat(event.Name)
                    if err != nil {
                        log.Println(err)
                    } else if fi.IsDir() {
                        watcher.AddFolder(event.Name)
                    } else {
                        watcher.Files <- event.Name // created a file
                    }
                }

                if event.IsModify() {
                    // modified a file, assuming that you don't modify folders
                    watcher.Files <- event.Name
                }

            case err := <-watcher.Error:
                log.Println("error", err)
            }
        }
    }()

    return watcher
}
