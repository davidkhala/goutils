#!/usr/bin/env bash
set -e
CURRENT=$(cd $(dirname ${BASH_SOURCE}) && pwd)
fcn=$1
remain_params=""
for ((i = 2; i <= ${#}; i++)); do
	j=${!i}
	remain_params="$remain_params $j"
done

if ! go version; then
	echo GO not found, exit...
	exit 1
fi
if [ -z "$GOPATH" ]; then
	export GOPATH=$(go env GOPATH)
fi
if [ -z "$GOPATH" ]; then
	echo GOPATH not found, exit...
	exit 1
fi
function get() {
	export GIT_TERMINAL_PROMPT=1
	go get -u -v $remain_params
}

$fcn $remain_params
