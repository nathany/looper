package gat

import (
    "errors"
    "fmt"
    "github.com/howeyc/fsnotify"
    "log"
    "os"
    "path/filepath"
)

type RecursiveWatcher struct {
    *fsnotify.Watcher
    Files chan string
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
    fmt.Printf("Watching path %s\n", folder)
}

// returns a slice of subfolders (recursive), including the folder passed in
func Subfolders(path string) (paths []string) {
    filepath.Walk(path, func(newPath string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if info.IsDir() {
            name := info.Name()
            // skip folders that begin with a dot
            hidden := filepath.HasPrefix(name, ".") && name != "." && name != ".."
            if hidden {
                return filepath.SkipDir
            } else {
                paths = append(paths, newPath)
            }
        }
        return nil
    })
    return paths
}

func Watch(path string) *RecursiveWatcher {
    watcher, err := NewRecurisveWatcher(path)
    if err != nil {
        log.Fatal(err)
    }

    watcher.Files = make(chan string, 20)

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
