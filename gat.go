package main

import (
    "bufio"
    "fmt"
    "github.com/howeyc/fsnotify"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// returns a slice of subfolders (recursive), including the folder passed in
func folders(path string) (paths []string) {
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
    watcher, err := fsnotify.NewWatcher()
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
                        err = watcher.Watch(event.Name)
                        if err != nil {
                            log.Println("error watching: ", err)
                        }
                    }
                }

                // delete|rename ... can't check if it's a folder... but is it in the directoy list?
                // watcher.RemoveWatch(event.Name)

            case err := <-watcher.Error:
                log.Println("error", err)
            }
        }
    }()

    for _, folder := range folders("./") {
        err = watcher.Watch(folder)
        if err != nil {
            log.Println("Error watching: ", folder, err)
        }
        fmt.Printf("Watching path %s\n", folder)
    }

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
