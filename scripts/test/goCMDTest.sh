#!/usr/bin/env bash
set -e
fcn=$1
remain_params=""
for ((i = 2; i <= ${#}; i++)); do
    j=${!i}
    remain_params="$remain_params $j"
done
CURRENT=$(cd $(dirname ${BASH_SOURCE}) && pwd)
root=$(dirname $CURRENT)
$root/goCmd.sh getAndEnsure github.com/MediConCenHK/go-chaincode-common
$root/goCmd.sh getAndEnsure https://github.com/MediConCenHK/go-chaincode-common.git
$root/goCmd.sh getAndEnsure git@github.com:MediConCenHK/go-chaincode-common.git