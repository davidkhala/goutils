#!/usr/bin/env bash
set -e

isMacOS() {
	[[ $(uname) == "Darwin" ]]
	return $?
}

isUbuntu20() {
	lsb_release -d | grep "Ubuntu 20."
	return $?
}

latest() {
	if [[ "$1" == "remove" ]]; then
		if isMacOS; then
			brew uninstall go || true
		elif isUbuntu20; then
			sudo snap remove go
		else
			sudo apt-get -y remove golang-go
			sudo add-apt-repository --remove -y ppa:longsleep/golang-backports
		fi

	else
		if isMacOS; then
			brew install go || true
		elif isUbuntu20; then
			sudo snap install go --classic
		else
			sudo add-apt-repository -y ppa:longsleep/golang-backports
			sudo apt update
			sudo apt install -y golang-go
		fi
	fi
}
"$@"
