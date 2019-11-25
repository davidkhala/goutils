# go module (vgo)

## Before usage
`export GO111MODULE=[off|on|auto]` (default is auto)

- [go.mod] If a module provides packages that are only imported by projects that haven't migrated to modules yet, the module requirement will be marked with an // indirect comment.

## Initialize
`go mod init`  

- creates a new go.mod file and automatically imports dependencies from Godeps.json, Gopkg.lock and etc.

**go: modules disabled inside GOPATH/src by GO111MODULE=auto; see '[go help modules](./vgo_man.md)'**

## Prune
`go mod tidy`
- It adds new module requirements for packages not provided by any known module
- It removes requirements on modules that don't provide any imported packages. 

It is always good practice to run go mod tidy before committing `go.mod` file to version control.

## Install vendor
`go mod vendor`
