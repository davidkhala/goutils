#!/usr/bin/env bash
set -e -x

_clone() {
	local GOPATH=$(go env GOPATH)
	local reposURL=$1
	local orgName=$2
	local projectName=$3
	local branch=$4
	local orgPath=${GOPATH}/src/github.com/${orgName}
	mkdir -p "${orgPath}"
	cd "${orgPath}"
	if [[ ! -d ${orgPath}/${projectName} ]]; then
		git clone "$reposURL"
	fi
	cd "${orgPath}/${projectName}"
	git pull
	if [[ -n $branch ]]; then
		git checkout $branch
	fi
	echo "${orgPath}/${projectName}"
}

get() {
	# support github and linux/unix only
	local reposURL=$1
	local branch=$2
	local orgName
	local projectName

	if [[ ${reposURL} == github* ]]; then
		get "https://${reposURL}.git" $branch
	elif [[ ${reposURL} == https://* ]]; then
		orgName=$(echo ${reposURL} | cut -d '/' -f 4)
		projectName=$(echo ${reposURL} | cut -d '/' -f 5 | cut -d '.' -f 1)
		_clone $reposURL $orgName $projectName $branch
	elif [[ ${reposURL} == git@* ]]; then
		echo ...using SSH
		orgName=$(echo ${reposURL} | cut -d '/' -f 1 | cut -d ':' -f 2)
		projectName=$(echo ${reposURL} | cut -d '/' -f 2 | cut -d '.' -f 1)

		_clone $reposURL $orgName $projectName $branch
	else
		exit 1
	fi
}

getAndEnsure() {
	local currentMode=$GO111MODULE
	local projectPath
	projectPath=$(get "$1" "$2" | tail -1)
	cd "${projectPath}"
	if [[ ! "$currentMode" == "on" ]]; then
		setModuleMode on
	fi

	go mod vendor
	cd - >/dev/null
	if [[ ! "$currentMode" == "on" ]]; then
		setModuleMode $currentMode
	fi

}
sync() {
	go list -m -json all
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

$@
