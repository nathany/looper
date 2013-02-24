# G.A.T.

G.A.T. is a development tool for the [Go Programming Language][go]. It automatically runs your tests and hot compiles your code when it detects file system changes.

## Get Going

If you are on OS X, you need to first install GNU Readline via [Homebrew](http://mxcl.github.com/homebrew/):

``` console
$ brew install readline
```

``` console
$ go get -u github.com/gophertown/gat

$ gat

G.A.T. 0.0.1 is now watching your files.
```

## Autotest

The convention in [Go][] is to use a *sidecar* test file in the same folder. When G.A.T. detects a change to your production code or the test itself, it will run that test.

You can trigger the full test suite by hitting enter.

Go files can be suffixed with an OS and/or architecture (eg.  `file_unix.go`). In this case, G.A.T. will look for both `file_unix_test.go` and `file_test.go`.

## Hot Compiles

...to be determined...

## .gatconfig

...

## Thanks

Inspired by Andrea Fazzi's [PrettyAutoTest][pat] and [devweb][] by Russ Cox. The name is inspired by [shotgun][], the reloading rack development server for Ruby. Special thanks to Chris Howey for the [fsnotify][] package.

[go]: http://golang.org/
[fsnotify]: https://github.com/howeyc/fsnotify
[pat]: https://github.com/remogatto/prettytest
[devweb]: http://code.google.com/p/rsc/source/browse/devweb/
[shotgun]: https://rubygems.org/gems/shotgun


