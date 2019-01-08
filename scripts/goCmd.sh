#!/usr/bin/env bash
set -e
CURRENT=$(cd $(dirname ${BASH_SOURCE}) && pwd)
fcn=$1
remain_params=""
for ((i = 2; i <= ${#}; i++)); do
	j=${!i}
	remain_params="$remain_params $j"
done

function get() {
	export GIT_TERMINAL_PROMPT=1
	go get -u -v $remain_params
}

$fcn $remain_params