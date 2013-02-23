package main

import (
    "bufio"
    "errors"
    "fmt"
    "github.com/howeyc/fsnotify"
    "log"
    "os"
    "path/filepath"
    "strings"
    "sync"
)

type RecursiveWatcher struct {
    *fsnotify.Watcher
    folders     []string
    folderMutex sync.RWMutex
}

func NewRecurisveWatcher(path string) (*RecursiveWatcher, error) {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return nil, err
    }
    folders := Subfolders(path)
    if len(folders) == 0 {
        return nil, errors.New("No folders to watch.")
    }
    rw := &RecursiveWatcher{Watcher: watcher}

    for _, folder := range folders {
        rw.AddFolder(folder)
    }
    return rw, nil
}

func (watcher *RecursiveWatcher) HasFolder(folder string) bool {
    watcher.folderMutex.RLock()
    defer watcher.folderMutex.RUnlock()

    for _, f := range watcher.folders {
        if f == folder {
            return true
        }
    }
    return false
}

func (watcher *RecursiveWatcher) AddFolder(folder string) {
    if watcher.HasFolder(folder) {
        fmt.Printf("Already watching %s\n", folder)
        return
    }

    watcher.folderMutex.Lock()
    defer watcher.folderMutex.Unlock()

    err := watcher.Watch(folder)
    if err != nil {
        log.Println("Error watching: ", folder, err)
    }

    watcher.folders = append(watcher.folders, folder)
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

func watch() {
    watcher, err := NewRecurisveWatcher("./")
    if err != nil {
        log.Fatal(err)
    }

    go func() {
        for {
            select {
            case event := <-watcher.Event:
                fmt.Print("\n", event, "\n> ")

                // create a directory b/c, ./b
                if event.IsCreate() {
                    fi, err := os.Stat(event.Name)
                    if err != nil {
                        log.Println(err)
                    } else if fi.IsDir() {
                        watcher.AddFolder(event.Name)
                    }
                }

                // delete|rename ... can't check if it's a folder... but is it in the directoy list?
                // watcher.RemoveWatch(event.Name)

            case err := <-watcher.Error:
                log.Println("error", err)
            }
        }
    }()

    // do stuff
    for {
        fmt.Print("> ")
        r := bufio.NewReader(os.Stdin)
        in, err := r.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        in = strings.ToLower(strings.TrimSpace(in))
        if in == "exit" {
            break
        }

        fmt.Println(in)
    }

    watcher.Close()
}

func main() {
    fmt.Println("G.A.T. 0.0.1 is watching your files")
    watch()
}
