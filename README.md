# go-utils
golang utils, it covers
 - go command utils
 - format utils
 - crypto utils
 - http request utils

## Notes
- when `dep ensure` in CICD environment need to access private golang repository as part of vendor, please `cp ./.circleci/.netrc $HOME/`.
  - See in [dep FAQ: consume private repos](https://github.com/golang/dep/blob/master/docs/FAQ.md#how-do-i-get-dep-to-consume-private-git-repos-using-a-github-token)

## DONE
- EC public key encryption (refer to ethereum.ECIES
but without huge vendor directory)

## TODO
- syndtr/goleveldb
- EC private key generate does not load string seed to io.reader: panic: unexpected EOF
