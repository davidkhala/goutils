# go-utils

[![Build Status](https://travis-ci.com/davidkhala/goutils.svg?branch=stable)](https://travis-ci.com/davidkhala/goutils) 
[![Github Actions](https://github.com/davidkhala/goutils/workflows/Github%20Actions/badge.svg?branch=stable)](https://github.com/davidkhala/goutils/actions?query=branch%3Astable)

golang utils, it covers
 - go command utils
 - format utils
 - crypto utils
 - http request utils

## Prerequisite
- golang 1.11.x


## vendor tool
- [Go Module](./vgo.md)

## [Notes](https://github.com/davidkhala/goutils/wiki/Notes)

### Fetch private github vendor packages
- For `dep`, See in [dep FAQ: consume private repos](https://github.com/golang/dep/blob/master/docs/FAQ.md#how-do-i-get-dep-to-consume-private-git-repos-using-a-github-token)
    - content of `.netrc`:
         ```
            machine github.com
             login [YOUR_GITHUB_USERNAME]
             password [GITHUB_TOKEN]
         ```    
- For Go module
  - `git config --global url."https://[YOUR_GITHUB_USERNAME]:[GITHUB_TOKEN]@github.com".insteadOf "https://github.com"` 

## DONE
- EC public key encryption (refer to ethereum.ECIES but without huge vendor directory)


## TODO
- syndtr/goleveldb
- EC private key generate does not load string seed to io.reader: panic: unexpected EOF
