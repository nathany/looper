# Looper Changelog

Roadmap & voting at the public [Trello board](https://trello.com/b/VvblYiSE).

## 0.1.2 / 2013-05-16

* Rename to Looper
* Packages are the unit of compilation in Go. Use a package-level granularity for testing.

## 0.1.1 / 2013-04-21

* Fixes "named files must all be in one directory" error [#2]
* Pass through for -tags command line argument. Thanks @jtacoma.

## 0.1.0 / 2013-02-24

* Recursively watches the file system, adding subfolders when created.
* Readline interaction to run all tests or exit.
* ANSI colors to add some flare.
* Focused testing of a single file for a quick TDD loop (subject to change)
