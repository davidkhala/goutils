# go-utils
 
![Github Actions](https://github.com/davidkhala/goutils/workflows/Github%20Actions/badge.svg)

golang utils, it covers
 - go command utils
 - format utils
 - crypto utils
 - http request utils

## Prerequisite
- golang 1.14


## vendor tool
- [Go Module](./vgo.md)

## Notes
- [test]golang file with suffix `_test.go` is test file
- code security: `securego/gosec`
- [conditional build](https://golang.org/pkg/go/build/#hdr-Build_Constraints)
- Directory and file names that begin with "." or "_" are ignored by the go tool, as are directories named "testdata".
## DONE
- EC public key encryption (refer to ethereum.ECIES but without huge vendor directory)


## TODO
- syndtr/goleveldb
- EC private key generate does not load string seed to io.reader: panic: unexpected EOF
- How to fetch private Golang module in go mod vendor fashion
