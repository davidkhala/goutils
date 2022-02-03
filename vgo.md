# go module (vgo)

## Initialize
`go mod init`  

- creates a new `go.mod` file and automatically imports dependencies from Godeps.json, Gopkg.lock and etc.

**see '[go help modules](./vgo_man.md)'**

## Prune
`go mod tidy`
- It adds new module requirements for packages not provided by any known module
- It removes requirements on modules that don't provide any imported packages. 
- If a module provides packages that are only imported by projects that haven't migrated to modules yet, the module requirement will be marked with an // indirect comment.
It is always good practice to run go mod tidy before committing `go.mod` file to version control.

## Install vendor
`go mod vendor`
