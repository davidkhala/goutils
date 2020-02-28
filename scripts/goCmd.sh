#!/usr/bin/env bash
set -e
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
	local GOPATH=$(go env GOPATH)

	if [[ ${reposURL} == github* ]]; then
		get "https://${reposURL}.git"
	elif [[ ${reposURL} == https://* ]]; then
		orgName=$(echo ${reposURL} | cut -d '/' -f 4)
		projectName=$(echo ${reposURL} | cut -d '/' -f 5 | cut -d '.' -f 1)
		local orgPath=${GOPATH}/src/github.com/${orgName}
		mkdir -p ${orgPath}
		cd ${orgPath}
		if [[ ! -d ${orgPath}/${projectName} ]]; then
			git clone $1
		else
			cd ${orgPath}/${projectName}
			git pull
		fi
		echo ${orgPath}/${projectName}
	elif [[ ${reposURL} == git@* ]]; then
		echo ...using SSH
		orgName=$(echo ${reposURL} | cut -d '/' -f 1 | cut -d ':' -f 2)
		projectName=$(echo ${reposURL} | cut -d '/' -f 2 | cut -d '.' -f 1)

		local orgPath=${GOPATH}/src/github.com/${orgName}
		mkdir -p ${orgPath}
		cd ${orgPath}
		if [[ ! -d ${orgPath}/${projectName} ]]; then
			git clone $1
		else
			cd ${orgPath}/${projectName}
			git pull
		fi
		echo ${orgPath}/${projectName}
	else
		exit 1
	fi
}

getAndEnsure() {
	local projectPath
	projectPath=$(get "$1" | tail -1)
	cd "${projectPath}"
	setModuleMode on
	go mod vendor
	cd - >/dev/null
	setModuleMode auto
}

setModuleMode() {
	modes=('off' 'on' 'auto')
	if [[ " ${modes[*]} " == *"$1"* ]]; then
		export GO111MODULE=$1
		echo "export GO111MODULE=$1"
	else
		echo "illegal GO111MODULE [${1}]"
		exit 1
	fi
}

$fcn $remain_params
