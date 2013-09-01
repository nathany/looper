# Looper

Looper is a development tool for the [Go Programming Language][go]. It automatically runs your tests and hot compiles your code when it detects file system changes.

## Status

[![Build Status](https://drone.io/github.com/gophertown/looper/status.png)](https://drone.io/github.com/gophertown/looper/latest)

This is an *early alpha*. There is still quite a lot to do (Hot Compiles, Growl notifications, and interactions for profiling, benchmarking, etc.). Also, it has only been tested on Mac OS X 10.8.

See the public [Trello board](https://trello.com/b/VvblYiSE) for the Roadmap.

## Get Going

If you are on OS X, you need to first install GNU Readline via [Homebrew](http://mxcl.github.com/homebrew/):

``` console
$ brew install readline
```

To install Looper, or to update your installation, run:

``` console
$ go get -u github.com/gophertown/looper
```

Then run `looper` in your project folder:

``` console
$ looper
Looper 0.2.1 is watching your files
Type help for help.

Watching path ./
```

## Gat (Go Autotest)

Packages are the unit of compilation in Go. By convention, each package has a separate folder, though a single folder may also have a `_test` package for black box testing.

When Looper detects a change to a *.go file, it will build & run the tests for that directory. You can also run all tests against all packages at once.

To setup a Suite definition ([Gocheck][], [PrettyTest][pat]), additional Checkers, or other test helpers, use any test file you like in the same folder (eg. `suite_test.go`).

Gat is inspired by Andrea Fazzi's [PrettyAutoTest][pat].

## Blunderbuss (Hot Compile)

...to be determined...

Blunderbuss is inspired by [shotgun][], the reloading rack development server for Ruby, and [devweb][] by Russ Cox.

## Interactions

* `a`, `all`, `↩`: Run all tests.
* `h`, `help`: Show help.
* `e`, `exit`: Quit Looper

## Thanks

Special thanks to Chris Howey for the [fsnotify][] package.

[go]: http://golang.org/
[fsnotify]: https://github.com/howeyc/fsnotify
[pat]: https://github.com/remogatto/prettytest
[devweb]: http://code.google.com/p/rsc/source/browse/devweb/
[shotgun]: https://rubygems.org/gems/shotgun
[Gocheck]: http://labix.org/gocheck

