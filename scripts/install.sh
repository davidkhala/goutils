#!/usr/bin/env bash
set -e

isMacOS() {
  [[ $(uname) == "Darwin" ]]
  return $?
}

latest() {
  if isMacOS; then
    brew install go || true
  else
    lsb_dist=$(curl -sSL https://raw.githubusercontent.com/davidkhala/linux-utils/main/system.sh | bash -s get_distribution)
    case $lsb_dist in
    ubuntu | debian)
      sudo snap install go --classic
      ;;
    centos | rhel | ol)
      sudo yum install -y golang-bin
      ;;
    esac
  fi
}
uninstall() {
  if isMacOS; then
    brew uninstall go || true
  else
    lsb_dist=$(curl -sSL https://raw.githubusercontent.com/davidkhala/linux-utils/main/system.sh | bash -s get_distribution)
    case $lsb_dist in
    ubuntu | debian)
      sudo snap remove go
      ;;
    centos | rhel | ol)
      sudo yum remove golang-bin
      ;;
    esac
  fi
}
"$@"
