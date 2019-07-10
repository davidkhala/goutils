#!/usr/bin/env bash
set -e
CURRENT=$(cd $(dirname ${BASH_SOURCE}) && pwd)
fcn=$1
remain_params=""
for ((i = 2; i <= ${#}; i++)); do
    j=${!i}
    remain_params="$remain_params $j"
done

get() {
    # support github and linux/unix only
    local reposURL=$1
    local orgName
    local projectName

    if [[ ${reposURL} == github* ]]; then
        echo ...using native go get format
        export GIT_TERMINAL_PROMPT=1
        go get -u -v $1
    elif [[ ${reposURL} == https://* ]]; then
        orgName=$(echo ${reposURL} | cut -d '/' -f 4)
        projectName=$(echo ${reposURL} | cut -d '/' -f 5 | cut -d '.' -f 1)
        get github.com/${projectName}/$orgName
    elif [[ ${reposURL} == git@* ]]; then
        echo ...using SSH
        orgName=$(echo ${reposURL} | cut -d '/' -f 1 | cut -d ':' -f 2)
        projectName=$(echo ${reposURL} | cut -d '/' -f 2 | cut -d '.' -f 1)
        local GOPATH=$(go env GOPATH)
        local orgPath=${GOPATH}/src/github.com/${orgName}
        mkdir -p ${orgPath}
        cd ${orgPath}
        if [[ ! -d ${orgPath}/${projectName} ]]; then
            git clone $1
        else
            cd ${orgPath}/${projectName}
            git pull
        fi
    fi
}

$fcn $remain_params
