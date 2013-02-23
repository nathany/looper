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
)

type RecursiveWatcher struct {
    *fsnotify.Watcher
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
    err := watcher.Watch(folder)
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

func watch() {
    watcher, err := NewRecurisveWatcher("./")
    if err != nil {
        log.Fatal(err)
    }

    go func() {
        for {
            select {
            case event := <-watcher.Event:
                fmt.Println(event)

                // create a directory
                if event.IsCreate() {
                    fi, err := os.Stat(event.Name)
                    if err != nil {
                        log.Println(err)
                    } else if fi.IsDir() {
                        watcher.AddFolder(event.Name)
                    }
                }

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
